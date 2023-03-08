package Datatypes

import (
	"fmt"
)

/*
Find an element in array such that sum of left array is equal to sum of right array
input: arr[] = [1, 2, 3, 4, 5, 6, 7, 8, 9,10,12], K = 4
Output: 3, 2, 1, 6, 5, 4, 9, 8, 7
*/

func RevArrGrp(arr []int, groupsize int) {

	maxlen := len(arr)
	numgroups := maxlen / groupsize
	totalelemtraverse := groupsize * numgroups

	fmt.Println("max and group : ", maxlen, numgroups, totalelemtraverse)

	for i := 0; i < len(arr); i++ {
		tempend := i + groupsize - 1
		if i == totalelemtraverse {
			break
		}
		for start := i; start < tempend; start++ {
			if start > tempend {
				break
			}

			fmt.Printf("Inside : start: %v, Tempend: %v , arr %v\n", start, tempend, arr)
			arr[start], arr[tempend] = arr[tempend], arr[start]
			tempend -= 1
			start += 1

		}

		i = tempend + 1

	}
	fmt.Println("reverse group array is :", arr)

}
