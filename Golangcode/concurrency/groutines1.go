package main

import (
	"fmt"
	"time"
)

//Goroutines1 to print numbers in golang
func Printnum(num int) {
	for i := 0; i <= num; i++ {
		fmt.Printf("%v", i)
	}
}
func main() {
	fmt.Println("Printing golang go-routines Chapter 1........")

	time_st := time.Now()
	go Printnum(10)
	time.Sleep(10 * time.Millisecond)
	time_end := time.Now()

	fmt.Printf("\n")
	//Anonymous goroutines in GOlang
	go func(name string) {
		for _, c := range name {
			fmt.Printf("%c", c)
		}
	}("nipun")
	time.Sleep(10 * time.Millisecond)

	time_diff := time_end.Sub(time_st).Seconds()
	fmt.Printf("\nTime taken to print 10 numbers sequentially is: %v\n", time_diff)
	fmt.Printf("go main execution stopped")
}
