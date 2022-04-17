package Mytestinglib

import (
	"fmt"
	"github.com/nipunchoprafree91/Golangcode/Mychannel"
	"sync"
)

func Mytestinglibfunc() {
	fmt.Println("Starting Execute of the Mytesting lib function")
	fmt.Println("Calling My channel Code From Testing library")

	wg := new(sync.WaitGroup)
	wg.Add(3)

	go Mychannel.Mychannelfunc(wg)
	go Mychannel.Mychannelfunc1(wg)
	go Mychannel.Mychannelfunc2(wg)

	wg.Wait()

}
