package main

import "fmt"

func somePrint1() {
	for i := 1; i <= 1000; i++ {
		go func() {
			fmt.Println("go")
		}()
	}
	go func() {
		fmt.Print("end")
	}()
}

func somePrint2() {
	for i := 1; i <= 100; i++ {
		go func() {
			fmt.Println("go")
		}()
	}
	go func() {
		fmt.Print("end")
	}()
}
