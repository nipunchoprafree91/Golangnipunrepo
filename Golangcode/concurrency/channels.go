package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Printing golang go-routines Chapter 2.....")

	//Anonymous goroutines in GOlang
	go func(name string) {
		for _, c := range name {
			fmt.Printf("%c", c)
		}
	}("nipun")
	time.Sleep(10 * time.Millisecond)

	//Anonymous goroutines in GOlang
	go func(num int) {
		for i := 0; i <= num; i++ {
			fmt.Printf("%d", i)
		}
	}(5)
	time.Sleep(10 * time.Millisecond)

	fmt.Println("\n")
	fmt.Println("Ending main goroutine in chapter-2")
}
