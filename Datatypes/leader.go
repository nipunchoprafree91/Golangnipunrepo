package Datatypes

import (
	"fmt"
)

/*
 Write a program to print all the LEADERS in the array.
 An element is a leader if it is greater than all the elements to its right side.
 And the rightmost element is always a leader.
*/
func Leader(arr []int) {

	for elem := 0; elem < len(arr); elem++ {

		for secelem := elem + 1; secelem < len(arr); secelem++ {

			if arr[elem] <= arr[secelem] {
				break
			}

			if secelem == len(arr)-1 {
				fmt.Println("Leaders ", arr[elem])
			}

		}
	}
}
