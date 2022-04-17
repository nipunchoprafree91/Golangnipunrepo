package Mytestinglib

import (
	"fmt"
	"github.com/nipunchoprafree91/Golangcode/Mychannel"
)

func Mytestinglibfunc() {
	fmt.Println("Starting Execute of the Mytesting lib function")
	fmt.Println("Calling My channel Code From Testing library")
	go Mychannel.Mychannelfunc()

}
