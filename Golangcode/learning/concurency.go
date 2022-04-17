package main

import (
	"fmt"
	"runtime"
	"time"

)

func main(){
	var result int

	processor := runtime.NumCPU()

	fmt.Println("No of Procs: ", processor)
	fmt.Println("Result is : " , result)

	for i:=0 ; i < processor ; i++{
		go func(){
			for {result ++}
		}()
	}
	time.Sleep(time.Second)
	fmt.Println("Result=",result)
}