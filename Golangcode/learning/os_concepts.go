package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

)

func main() {
	userDir, _ := os.UserHomeDir()
	fmt.Println(userDir)

	allFiles, _ := ioutil.ReadDir(userDir)
	//fmt.Println(allFiles) //Returns Byte array

	//var wg sync.WaitGroup
	// var lf FileStruct
	fmt.Printf("Type of allFilesis %T", allFiles)
	//wg.Add(len(allFiles))

	st := time.Now()
	for _, fileDet := range allFiles {
		go func(f os.FileInfo) {
			fmt.Println(f.Name())
			//defer wg.Done()
		}(fileDet)
	}
	//wg.Wait()
	endt := time.Now()
	fmt.Println("Time taken is", endt.Sub(st))
}
