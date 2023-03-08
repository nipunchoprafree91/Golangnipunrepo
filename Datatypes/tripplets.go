package Datatypes

import (
	"fmt"
	"runtime"
)

func GetTriplets() [][]int {
	fmt.Println("Hi Lets solve triplet problems....")
	fmt.Println("Memory usage Before the program:")
	PrintMemUsage()

	array1 := []int{5, 1, 7, 6, 3, 9, 4}

	match3n := matchingtripleto3n(array1)
	fmt.Printf("Matching array with complexity O(n3) is: %v\n", match3n)

	out := comparearrays(match3n)
	fmt.Printf("Matching unique array with complexity O(n3) is: %v\n", out)

	// match2n := matchingtripleto2n(array1)
	// fmt.Printf("Matching array with complexity O(n2) is: %v\n", match2n)

	return out

}

func searchsum(sum int, arr []int, elemarr []int) (finalarr []int, midcount int) {
	midcount = 0
	newarr := make([]int, 0)
	for i := 0; i < len(arr); i++ {
		if sum == arr[i] {
			midcount = 1
			elemarr = append(elemarr, arr[i])
			newarr = append(newarr, elemarr...)
			finalarr = newarr
			// fmt.Printf("sum: %v is matched with array element: %v\n", sum, finalarr)
		}
	}
	return finalarr, midcount
}

func matchingtripleto2n(arr []int) (finalarr [][]int) {
	startindex := 0
	matches := 0
	endindex := len(arr)
	var sumel int

	//For iterating to find  first element
	for firstel := startindex; firstel < endindex; firstel++ {

		for nextel := firstel + 1; nextel < endindex; nextel++ {
			elemarr := []int{}
			sumel = arr[firstel] + arr[nextel]

			//append to get the array elements passed to search
			elemarr = append(elemarr, arr[firstel])
			elemarr = append(elemarr, arr[nextel])

			// fmt.Printf("element array : %v\n", elemarr)

			returnarr, match := searchsum(sumel, arr, elemarr)
			// fmt.Printf("final array got in function : %v , matches : %v\n", finalarr, matches)
			if match == 1 {
				matches += 1
				finalarr = append(finalarr, returnarr)
			}

		}
	}
	// fmt.Println("Finalarr: ", finalarr)
	fmt.Println("Total matches: ", matches)
	runtime.GC()
	fmt.Println("Memory usage after GC")
	PrintMemUsage()
	return finalarr

}

func comparearrays(givenarr [][]int) (uniquearr [][]int) {
	lengiven := len(givenarr)
	fmt.Println("Length of given array is: %v", lengiven)
	for i := 0; i < lengiven-1; i++ {
		fmt.Printf("Array at index %v is:%v\n", i, compareelem(givenarr[i], givenarr[i+1]))
	}
	uniquearr = append(uniquearr)
	return uniquearr

}

func compareelem(arr1 []int, arr2 []int) (uniquearr []int) {
	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr2); j++ {
			if arr1[i] == arr2[j] {
				// fmt.Println("Both arrays have equal elements")
				uniquearr = append(uniquearr, arr1[i])
			}
		}

	}
	return uniquearr
}

func matchingtripleto3n(arr []int) (matcharr [][]int) {
	sum := 0
	startindex := 0
	endindex := len(arr)
	fmt.Println("starting sum is : ", sum)
	fmt.Printf("start index is : %v and Length of array is : %v\n", startindex, endindex)
	arrcon := []int{}
	count := 0

	//For iterating over only first element
	for stind := startindex; stind <= endindex-1; stind++ {

		//For second element
		for actstind := startindex; actstind < endindex-1; actstind++ {

			arrcon = []int{}
			//For third element
			for secind := startindex; secind < endindex; secind++ {

				if secind == stind || secind < actstind || actstind == stind || actstind == secind {
					continue
				}
				arrcon = append(arrcon, arr[stind])
				arrcon = append(arrcon, arr[actstind])
				arrcon = append(arrcon, arr[secind])

				// fmt.Printf("Count: %v ,start index: %d  actind index: %d , secind :%v , Array: %v\n",
				//  count, stind, actstind, secind, arrcon)

				if arr[stind]+arr[actstind] == arr[secind] {
					fmt.Printf("%v + %v = %v in %v\n", arr[stind], arr[actstind], arr[secind], arrcon)
					count += 1
					matcharr = append(matcharr, arrcon)
				}
				if arr[actstind]+arr[secind] == arr[stind] {
					fmt.Printf("%v + %v = %v in %v\n", arr[actstind], arr[secind], arr[stind], arrcon)
					count += 1
					matcharr = append(matcharr, arrcon)
				}
				if arr[stind]+arr[secind] == arr[actstind] {
					fmt.Printf("%v + %v = %v in %v\n", arr[stind], arr[secind], arr[actstind], arrcon)
					count += 1
					matcharr = append(matcharr, arrcon)
				}
				arrcon = []int{}
			}
		}
	}
	fmt.Printf("Number of triplets having condition met :%v\n", count)
	runtime.GC()
	fmt.Println("Memory usage after GC")
	PrintMemUsage()
	return matcharr

}

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", m.Alloc)
	fmt.Printf("\tTotalAlloc = %v MiB", m.TotalAlloc)
	fmt.Printf("\tSys = %v MiB", m.Sys)
	fmt.Printf("\tNumGC = %v\n\n", m.NumGC)
}
