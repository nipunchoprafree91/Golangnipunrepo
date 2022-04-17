package Mytestinglib

import (
	"fmt"
	"github.com/nipunchoprafree91/Golangcode/Mychannel"
	"time"
)

func Mytestinglibfunc() {
	fmt.Println("Starting Execute of the Mytesting lib function")
	fmt.Println("Calling My channel Code From Testing library")
	go Mychannel.Mychannelfunc()
	time.Sleep(1 * time.Second)
	go Mychannel.Mychannelfunc1()
	time.Sleep(1 * time.Second)
	go Mychannel.Mychannelfunc2()
	time.Sleep(1 * time.Second)

}
