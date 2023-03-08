package Datatypes

import (
	"fmt"
)

func Getuniquesubstring() map[string]int {
	for _, c := range "go is awesome" {
		fmt.Printf("%c", c)
	}
	fmt.Println()
	string1 := "geeksforgeeks"

	maps1 := make(map[string]int)

	stind := 0
	endind := 0

	max_len := 0

	fmt.Printf("string to be : %v\n", string1)
	fmt.Printf("string to be : %v\n", maps1)
	fmt.Printf("string to be : %v\n", stind)
	fmt.Printf("string to be : %v\n", endind)
	fmt.Printf("string to be : %v\n", max_len)

	for stind := range string1 {
		//fmt.Printf("char at Index at : %v is  %c \n", stind, string1[stind])

		val, ok := maps1[string(string1[stind])]

		if !ok {
			maps1[string(string1[stind])] = val + 1
			// fmt.Printf("Map when key is present is %v\n", maps1)
		} else {
			maps1[string(string1[stind])]++
			// fmt.Printf("Map when key is not %v\n", maps1)
		}

	}
	fmt.Printf("Map %v\n", maps1)
	return maps1
}
