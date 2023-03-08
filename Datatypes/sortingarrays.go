package Datatypes

import (
	"fmt"
)

//compare each element against each other and this doe not sort in
//First attempt we have to attempt multiple times
//check after every sort if  there is some element in unsorted order

// https://dev.to/adnanbabakan/sorting-algorithms-in-go-725

func Bubblesort(arr []int) {
	arrlength := len(arr)
	count := 0
	everyelementcheked := false
	for !everyelementcheked {
		everyelementcheked = true
		for index := 0; index < arrlength-1; index++ {
			if arr[index] > arr[index+1] {
				tmp := arr[index+1]
				arr[index+1] = arr[index]
				arr[index] = tmp
				everyelementcheked = false
			}
			fmt.Printf("COUNT: %v, Array while Bubble sorting is : %v\n", count, arr)
			count++
		}
	}
	fmt.Println("Array after sorting is : ", arr)
}

func Bubblesortwithrecursion(arr []int, size, count int) (arr1 []int, cnt int) {
	if size == 1 {
		fmt.Printf("Count: %v,Size: %v,  Array while Bubble sorting is : %v\n", count+1, size, arr)
		return arr, count
	}
	for index := 0; index < size-1; index++ {
		if arr[index] > arr[index+1] {
			tmp := arr[index+1]
			arr[index+1] = arr[index]
			arr[index] = tmp
		}
		count += 1
		fmt.Printf("Count: %v,Size: %v, Array while Bubblesort with recursion is : %v\n", count, size, arr)
	}
	fmt.Println(size)
	Bubblesortwithrecursion(arr, size-1, count)

	return arr, count
}

/*
Insertion sort is a famous sorting algorithm that works like
 sorting a deck of cards in hand. As you proceed through the elements of
 an array you would move it back until it is in the correct place.
*/
func InsertionSort(arr []int) {
	elements := len(arr)
	count := 0
	for compareelem := 1; compareelem <= elements-1; compareelem++ {

		for allelem := 0; allelem < compareelem; allelem++ {

			if arr[compareelem] < arr[allelem] {
				arr[allelem], arr[compareelem] = arr[compareelem], arr[allelem]
			}
		}
		count++
		fmt.Printf("Count: %v, Array while InsertionSort sorting is : %v\n", count, arr)
	}

}

/*
Selection sort is quite interesting and yet simple
What you need to do is simply replace the current element
in iteration by the lowest value on the right. As you go
 further the left part of the array is sorted, and you
  don't need to check for the last element.
*/
func Selectionsort(arr []int) {
	elements := len(arr)
	fmt.Println("Array before switching is:  ", arr)
	var comp, small_index int

	for elem := 0; elem < elements-1; elem++ {
		small_index = elem
		for comp = elem + 1; comp < elements; comp++ {

			if arr[comp] < arr[small_index] {
				small_index = comp
			}

		}

		arr[elem], arr[small_index] = arr[small_index], arr[elem]

		fmt.Println("array while switching is :  ", arr)
	}
	fmt.Println("array after sorting is :    ", arr)
}

//Merge sort functions start
/*
Merge is a quite fast sorting algorithm. In a merge sort
, you should utilize the divide-and-conquer practice.
 First, you'll divide the passed array by 2 recursively,
  until you reach the length of 1, then you should merge them.
   To merge two arrays that we know are already sorted themselves
   you can implement a simple function called merge.
*/
func Mergesortarray(fp, sp []int) []int {
	var n = make([]int, len(fp)+len(sp))

	fpindex := 0
	spindex := 0
	nindex := 0

	for fpindex < len(fp) && spindex < len(sp) {

		if fp[fpindex] < sp[spindex] {
			n[nindex] = fp[fpindex]
			fpindex++
		} else {
			n[nindex] = sp[spindex]
			spindex++
		}
		fmt.Println("End Index : ", n)
		nindex++
	}

	for fpindex < len(fp) {
		n[nindex] = fp[fpindex]

		fpindex++
		nindex++
		fmt.Println("End Index : ", n)
	}

	for spindex < len(sp) {
		n[nindex] = sp[spindex]

		spindex++
		nindex++
		fmt.Println("End Index : ", n)
	}

	return n
}

func Mergesort(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}

	var fp = Mergesort(arr[0 : len(arr)/2])
	var sp = Mergesort(arr[len(arr)/2:])

	return Mergesortarray(fp, sp)
}
