package Multithreading

import (
	"fmt"
	"sync"
)

// we should start a listening goroutine and then pass the value
// int the channel so that goroutines can read from that channel
// as when you pass that value into the channel it expects
//another goroutine which can read , if not then it will be a deadlock

// Channel operations are also blocking in nature. When some data
// is written to the channel, goroutine is blocked until some other
// goroutine reads it from that channel

func Basicchannel() {
	fmt.Println("Program started")
	wg := new(sync.WaitGroup)

	c := make(chan string)
	d := make(chan int)

	wg.Add(1)
	go Printstr(wg, c)
	c <- "Nipun"
	wg.Wait()
	close(c)

	wg.Add(1)
	go Printint(wg, d)
	d <- 1
	wg.Wait()
	close(d)

	fmt.Println("\n Printing Squares......\n")

	square := make(chan int)
	go Printsquare(wg, square)
	wg.Wait()
	for val := range square {
		wg.Add(1)
		fmt.Println(val)
	}

	//same as above but preferred as we can loop over channels
	// for {
	// 	val, ok := <-square
	// 	if !ok {
	// 		break
	// 	} else {
	// 		fmt.Println(val)
	// 	}
	// }
	fmt.Println("Program stopped")

}
func Printsquare(wg *sync.WaitGroup, square chan int) {
	defer wg.Done()
	for num := 1; num <= 9; num++ {
		square <- num * num
	}
	close(square)

}

func Printstr(wg *sync.WaitGroup, c chan string) {
	fmt.Println("Hello " + <-c + "!")
	wg.Done()
}

func Printint(wg *sync.WaitGroup, d chan int) {
	fmt.Println(<-d)
	wg.Done()
}
