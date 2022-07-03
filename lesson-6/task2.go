package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
)

var quit chan int = make(chan int)

func print10to19() {
	fmt.Println("Start******print10to19 ")
	for i := 10; i < 20; i++ {
		runtime.Gosched()
		fmt.Println("******10to19: ", i)
	}
	fmt.Println("End******print10to19 ")
	quit <- 1
}

func print20to29() {
	fmt.Println("Start======print20to29 ")
	for i := 20; i < 30; i++ {
		runtime.Gosched()
		fmt.Println("======20to29: ", i)
	}
	fmt.Println("End======print20to29 ")
	quit <- 2
}

func print30to39() {
	fmt.Println("Start######print30to39 ")
	for i := 30; i < 40; i++ {
		runtime.Gosched()
		fmt.Println("######30to39: ", i)
	}
	fmt.Println("End######print30to39 ")
	quit <- 0
}

func main() {
	trace.Start(os.Stderr)
	defer trace.Stop()

	runtime.GOMAXPROCS(3)
	fmt.Println("start ---")
	go print10to19()
	go print20to29()
	go print30to39()
	fmt.Println("start ===")
	for i := 0; i < 3; i++ {
		sc := <-quit
		fmt.Println("sc:", sc)
	}
	fmt.Println("end")
}
