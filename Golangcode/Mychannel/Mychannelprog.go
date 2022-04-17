package Mychannel

import (
	"fmt"
	"sync"
)

func Mychannelfunc(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Starting Execute of the Channel function")

}
func Mychannelfunc1(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Starting Execute of the Channel function1")
}
func Mychannelfunc2(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Starting Execute of the Channel function2")
}
