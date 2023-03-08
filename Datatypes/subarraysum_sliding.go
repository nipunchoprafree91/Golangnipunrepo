package Datatypes

import "fmt"

func Mymatchingsubarray() []int {
	fmt.Println("Hi Lets solve some array problems....")

	array1 := []int{5, 1, 4, 7, 6, 3, 9}
	sumexpected := 18
	match := Matchingsubarray(array1, sumexpected)
	fmt.Printf("matching array with a sum of %v is: %v\n", sumexpected, match)

	return match

}

func Matchingsubarray(arr []int, sumexpected int) (matcharr []int) {
	sum := 0
	startindex := 0
	endindex := len(arr)
	fmt.Println("Length of array is : ", endindex)

	sl := []int{}
	fmt.Printf("Start index is %v\n", startindex)
	fmt.Printf("End index is %v\n", endindex)
	fmt.Printf("Starting sum is %v\n", sum)
	fmt.Printf("Sum Expected is %v\n", sumexpected)

	// Printing all the subarrays in the array
	for stind := startindex; stind <= endindex; stind++ {
		// fmt.Println()
		for actend := endindex; actend >= 0; actend-- {
			subslicesum := 0
			if actend < stind || actend == stind {
				continue
			}
			sl = arr[stind:actend]
			fmt.Printf("start index: %d  end index: %d  array: %v\n", stind, actend-1, sl)

			for i := range sl {
				subslicesum += sl[i]
			}

			if subslicesum == 18 {
				fmt.Printf("sum at at start index :%v and end index : %v is 18\n", stind, actend-1)
				matcharr = append(matcharr, sl...)
				return matcharr
			}

		}
	}
	return matcharr

}
