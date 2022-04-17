package main

import (
	"fmt"
	"time"

)

func main() {
	st := time.Now().Format("2006-01-02")
	fmt.Printf("Type is %T",st)
	// time.Sleep(1 * time.Second)

	// et := time.Now()
	// fmt.Println(et.Format("10_01_2019"))
	// fmt.Println(st)

	// diff := et.Sub(st)
	// fmt.Println(diff)
	// count := Mypoll()
	// fmt.Println("Count is %v", count)
}
// func Mypoll() (count int64) {
// 	Myticker := time.NewTicker(1 * time.Second)
// 	count = 1
// 	for _ = range Myticker.C {
// 		count += 1
// 		fmt.Println("Tick No :", count)
// 		if count == 60 {
// 			fmt.Print("returning and stopping")
// 			return count
// 		}
// 	}
// 	fmt.Println("Tick No :", count)
// 	return 0
// }
