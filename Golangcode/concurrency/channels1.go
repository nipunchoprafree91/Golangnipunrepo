package main

import (
	"fmt"
	"time"
)

func greet(c chan string) {
	fmt.Println("Hello " + <-c + "!")
}

func main() {
	fmt.Println("starting Main execution")
	c := make(chan string)

	fmt.Printf("Channel Value:  %v\n", c)
	fmt.Printf("Channel Type:  %T\n", c)

	go greet(c)
	time.Sleep(1 * time.Second)

	//Passing value into the channel
	c <- "Nipun"
	//closing the channel
	close(c)

	//checking if channel is closed for operations
	_, ok := <-c
	if ok == true {
		fmt.Println("Channel is Open")
	} else {
		fmt.Println("Channel is closed for operations")
	}

	fmt.Println("stopping Main execution")

}
