package Datatypes

import (
	"fmt"
)

/* Print the elements of an array in the decreasing
 frequency if 2 numbers have the same frequency then
 print the one which came first

Input:  arr[] = {2, 5, 2, 8, 5, 5, 6, 8, 8, 8}
2=2
5=2
8=3
6=1
Output: arr[] = {8, 8, 8, 8, 5, 5, 5, 2, 2,6}
*/

func SortByFrequency(arr []int) {

	fmt.Println("Program to sort by frequency")

	freqcountarr := make([]int, len(arr))
	// elements of this array will be frequency of element
	// in the previous array ast the same index

	for elem := 0; elem < len(arr); elem++ {

		count := 0

		for nextelem := 0; nextelem < len(arr); nextelem++ {

			if arr[elem] == arr[nextelem] {
				count += 1
			}
		}
		//After this point we will know the frequency of every element at an index
		freqcountarr[elem] = count

	}
	fmt.Printf("Frequency array of elem is : %v\n", freqcountarr)

	sortByFrequencyArray(arr, freqcountarr)
}

// Use insertion sort to insert the  element of an array
//make sure to replace both array and freqrry otherwise sorting will not occur
func sortByFrequencyArray(arr, freqarr []int) {
	// at this point we have two arrays
	// [2, 5, 2, 8, 5, 5, 6, 8, 8, 8]        unsorted

	// [0 ,1, 2, 3, 4, 5, 6, 7, 8, 9]        unsorted
	// [2  3  2  4  3  3  1  4  4  4]        frequency
	fmt.Println("unsorted array: ", arr)
	fmt.Println("Frequency array:", freqarr)

	fmt.Println("Array before switching is:  ", arr)

	elements := len(arr)
	count := 0
	for compareelem := 1; compareelem <= elements-1; compareelem++ {

		for allelem := 0; allelem < compareelem; allelem++ {

			if freqarr[compareelem] < freqarr[allelem] {
				freqarr[allelem], freqarr[compareelem] = freqarr[compareelem], freqarr[allelem]
				arr[allelem], arr[compareelem] = arr[compareelem], arr[allelem]
			}
		}
		count++
		fmt.Printf("Count: %v, Array while InsertionSort sorting is : %v\n", count, arr)
	}

}
