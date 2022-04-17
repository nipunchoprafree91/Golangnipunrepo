package main

import (
	"fmt"
	"time"
)

func squares(c chan int) {
	for i := 0; i <= 3; i++ {
		num := <-c
		fmt.Println(num * num)
	}

}
func sender(c chan int) {
	c <- 1
	c <- 2
	c <- 3
	c <- 4
	close(c)
}

func main() {
	fmt.Println("starting Main execution")
	c := make(chan int, 3)

	go sender(c)

	for val := range c {
		fmt.Printf("length is %v and capacity of the channel is:%v and value is %v\n", len(c), cap(c), val)
		time.Sleep(100 * time.Millisecond)
	}

	d := make(chan int, 3)
	d <- 1
	d <- 2
	d <- 3
	close(d)

	for val1 := range d {
		fmt.Printf("length is %v and capacity of the channel is:%v and value is %v\n", len(d), cap(d), val1)
	}

	fmt.Println("stopping Main execution")
}
