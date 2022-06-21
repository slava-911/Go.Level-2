package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func main() {

	// Задание 1
	var wg sync.WaitGroup
	var x int64
	var pool = make(chan struct{}, runtime.NumCPU())
	for i := 1; i <= 1000; i++ {
		pool <- struct{}{}
		wg.Add(1)
		go func() {
			defer func() {
				wg.Done()
				<-pool
			}()
			atomic.AddInt64(&x, 1)
		}()
	}
	wg.Wait()
	fmt.Println("x =", x)

	// Задание 2
	sigCh := make(chan string)
	done := make(chan bool)
	go func() {
		defer fmt.Println("Выход из записывающей горутины")
		sigCh <- "SIG1"
		time.Sleep(1 * time.Second)
		sigCh <- "SIG2"
		time.Sleep(1 * time.Second)
		rand.Seed(time.Now().Unix())
		if rand.Float64() > 0.5 {
			sigCh <- "SIGTERM"
			time.Sleep(1 * time.Second)
		}
		sigCh <- "SIG3"
		time.Sleep(1 * time.Second)
		sigCh <- "SIG4"
		time.Sleep(3 * time.Second)
		close(sigCh)
	}()

	go func() {
		defer fmt.Println("Выход из читающей горутины")
		for sig := range sigCh {
			fmt.Println(sig)
			if sig == "SIGTERM" {
				done <- true
			}
		}
		close(done)
	}()

	_, ok := <-done
	if ok {
		<-time.After(1 * time.Second)
		fmt.Println("Выход по таймауту")
		return
	}

	fmt.Println("next")

}
