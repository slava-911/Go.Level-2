package main

import (
	"fmt"
	"sync"
	"time"
)

var mutex sync.Mutex

func main() {

	// Задание 1
	var wg sync.WaitGroup
	n := 100
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Printf("%d goroutine working \n", i)
			time.Sleep(1 * time.Second)
		}(i)
	}
	wg.Wait()
	fmt.Println("All goroutines completed")

	// Задание 2
	criticalSection()
	criticalSectionWithPanic()
	time.Sleep(1 * time.Second)
}

func criticalSection() {
	mutex.Lock()
	fmt.Println("critical section")
	defer mutex.Unlock()
}

func criticalSectionWithPanic() {
	defer func() {
		fmt.Println("recovered", recover())
		mutex.Unlock()
	}()
	mutex.Lock()
	fmt.Println("critical section with panic")
	panic("AAA!")
}
