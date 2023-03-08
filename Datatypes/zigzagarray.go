/*
Given an array of distinct elements of size N,
the task is to rearrange the elements of the array
 in a zig-zag fashion so that the converted array
 should be in the below form:

arr[0] < arr[1]  > arr[2] < arr[3] > arr[4] < . . . . arr[n-2] < arr[n-1] > arr[n].

Input:  N = 7 , arr[] = {4, 3, 7, 8, 6, 2, 1}
Output: N = 7 , arr[] = {3, 7, 4, 8, 2, 6, 1}

*/

package Datatypes

import (
	"fmt"
)

func Setzigzagarray(arr []int) {

	for elem := 0; elem < len(arr)-1; elem++ {

		if elem%2 == 1 {
			if arr[elem] < arr[elem+1] {
				fmt.Printf("Switching as element was less: %v , arr: %v\n", elem, arr)
				arr[elem], arr[elem+1] = arr[elem+1], arr[elem]
			}

		} else {
			if arr[elem] > arr[elem+1] {
				fmt.Printf("Switching as element was more: %v , arr: %v\n", elem, arr)
				arr[elem], arr[elem+1] = arr[elem+1], arr[elem]
			}
		}
	}
	fmt.Println("Array at the end is: ", arr)

}
