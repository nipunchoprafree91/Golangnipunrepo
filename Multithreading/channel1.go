package Multithreading

import (
	"fmt"
	"os/exec"
	"time"
)

func Firstgo(c chan string) {
	fmt.Println("Function to be called")
	cmd := exec.Command("nproc")
	stdoutStderr, _ := cmd.CombinedOutput()
	c <- string(stdoutStderr)

}

func Secondgo() (output chan map[int]string) {
	fmt.Println("Callee function")
	c := make(chan string)
	starttimeint := 0
	ticker := time.NewTicker(5 * time.Second)
	maxtime := time.Now().Add(50 * time.Second)

	output1 := make(map[int]string)
	for tt := range ticker.C {
		starttimeint += 5
		go Firstgo(c)
		output1[starttimeint] = <-c
		output <- output1
		if tt.After(maxtime) {
			fmt.Println("max time reached")
			break
		}
	}

	return output
}
