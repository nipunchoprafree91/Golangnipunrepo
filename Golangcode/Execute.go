// Learning Packages Golang
// Author: Nipun chopra
// Copyright 2022: Public

package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/nipunchoprafree91/Golangcode/Mytestinglib"
)

func main() {
	fmt.Println("Starting Execution of the Main function")
	fmt.Println("Calling Mytestinglib Code From Testing library")
	Mytestinglib.Mytestinglibfunc()
	_ = colly.NewCollector()
}
