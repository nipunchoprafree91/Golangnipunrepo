package Datatypes

import (
	"fmt"
)

/*
Find an element in array such that sum of left array is equal to sum of right array
*/

func EquilibriumPoint(arr []int) {

	for elem := 0; elem < len(arr); elem++ {

		leftarr := arr[:elem]
		rightarr := arr[elem+1:]

		secarraysum := 0
		for secarrelem := 0; secarrelem < len(leftarr); secarrelem++ {
			secarraysum += leftarr[secarrelem]
		}

		thirdarraysum := 0
		for thirdarray := 0; thirdarray < len(rightarr); thirdarray++ {
			thirdarraysum += rightarr[thirdarray]
		}

		if secarraysum == thirdarraysum {
			fmt.Printf("sum of left array : %v and right array %v is same.", leftarr, rightarr)
			fmt.Printf("Element is %v at index : %v", elem, arr[elem])
		}

	}
}
