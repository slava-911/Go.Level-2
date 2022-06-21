package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"sync/atomic"
	"syscall"
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
	sigCh := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigCh
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal")
	<-done
	<-time.After(1 * time.Second)
	fmt.Println("Exit by timeout")

}
