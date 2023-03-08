package Multithreading

import (
	"fmt"
	"sync"
)

func MyMain() {
	// Create two channels to allow communication between the goroutines
	c1 := make(chan int)
	c2 := make(chan []int)

	// Create a WaitGroup to wait for all goroutines to complete
	wg := &sync.WaitGroup{}

	// Add 2 to the WaitGroup to wait for 2 goroutines (sub1 and sub2)
	wg.Add(2)

	// Launch the sub1 goroutine
	go sub1(c1, c2, wg)

	// Send a value to c1
	c1 <- 10

	// Wait for an array to be received on c2
	result := <-c2

	// Print the result
	fmt.Println("Result:", result)

	// Wait for all goroutines to complete
	wg.Wait()
}

func sub1(c1 chan int, c2 chan []int, wg *sync.WaitGroup) {
	// Create a channel to receive the values from sub2
	c3 := make(chan int)

	// Launch the sub2 goroutine
	wg.Add(1)
	go sub2(c1, c3, wg)

	// Wait for a value to be received on c1
	value := <-c1
	fmt.Print(value)

	// Receive the values from sub2 through c3 and append them to an array
	var arr []int
	for i := range c3 {
		arr = append(arr, i)
	}

	// Send the array to main through c2
	c2 <- arr

	// Notify the WaitGroup that the sub1 goroutine has completed
	wg.Done()
}

func sub2(c1 chan int, c3 chan []int, wg *sync.WaitGroup) {
	// Wait for a value to be received on c1
	value := <-c1

	// Append the values to an array and send the array to c3
	var arr []int
	for i := 0; i < value; i++ {
		arr = append(arr, i)
	}
	c3 <- arr

	// Notify the WaitGroup that the sub2 goroutine has completed
	wg.Done()
}

/*

In this implementation, we create a separate channel c3 to receive

the values from sub2, append them to an array, and send the array

back to sub1 only once. We launch sub2 with c1, c3, and wg as parameters

and wait for a value to be received on c1. Once we receive a value,

we append the values to an array and send the array to c3. In sub1,

we launch sub2 with c1, c3, and wg as parameters and wait for a value

to be received on c1. Once we receive a value, we send it to sub2 through c3.

We then receive the values from sub2 through c3, append them to an array, and

send the array back to the main goroutine through c2. Finally, we notify wg

that the sub1 goroutine has completed.

Note that in sub2, we notify wg using wg.Done() once we have finished sending the array to c3.

*/
