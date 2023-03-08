package Datatypes

import (
	"fmt"
)

func Middletype(arr []int) {
	len1 := len(arr)
	fmt.Println("length of array is: ", len1)

	if len1%2 == 0 {
		middle := arr[len1/2]
		fmt.Println("Even array length of array is: ", middle)
	} else {
		middle := len1 / 2
		middle1 := arr[middle : middle+2]
		fmt.Println("Odd Array length of array is: ", middle1)
	}
}
