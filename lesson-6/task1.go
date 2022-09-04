package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"sync"
)

func main() {
	trace.Start(os.Stderr)
	defer trace.Stop()

	var (
		wg  sync.WaitGroup
		mux sync.Mutex
		x   int64
		n   int = 1000
	)
	wg.Add(n)
	for i := 1; i <= n; i++ {
		go func() {
			defer wg.Done()
			mux.Lock()
			x += 1
			mux.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println("x =", x)
}
