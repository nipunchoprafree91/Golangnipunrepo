package main

import (
	"fmt"
	"strings"

)

func main() {
	str := "Perf_xfs_1vmdk_low_density_5"
	vm := strings.Split(str, "_")
	fmt.Println(vm)
	fmt.Println(len(vm))

	arr1 := [...]int{1, 2, 3, 4, 5}
	fmt.Printf("value of arr1 is: %v\n", arr1)
	fmt.Printf("Type of the var is : %T\n", arr1)
	for _, value := range arr1 {
		fmt.Printf("Address of value %v is: %v\n", value, &value)
	}
	//Before change to array
	fmt.Printf("\n\nBefore change\n")
	var sl []int
	sl = arr1[:]
	fmt.Printf("Type of the var is : %T\n", sl)
	for _, value := range sl {
		fmt.Printf("Address of value %v is: %v\n", value, &value)
	}
	fmt.Printf("Type of the var is : %T\n", sl)

	arr1[0] = 10
	arr1[1] = 20
	arr1[2] = 30

	fmt.Printf("\n\nAfter Change\n")
	fmt.Printf("Type of the var is : %T\n", sl)
	for _, value := range sl {
		fmt.Printf("Address of value %v is: %v\n", value, &value)
	}
	fmt.Printf("Type of the var is : %T\n", sl)

	sl[0] = 1
	sl[1] = 2
	sl[2] = 3

	fmt.Printf("\n\nAfter Change to slice, we Print the slices to confirm \n")
	for _, value := range sl {
		fmt.Printf("Address of value %v is: %v\n", value, &value)
	}
	fmt.Printf("\n\nAfter Change to slice, we Print the Array\n")
	fmt.Printf("value of arr1 is: %v\n", arr1)
	fmt.Printf("Type of the var is : %T\n", arr1)
	for _, value := range arr1 {
		fmt.Printf("Address of value %v is: %v\n", value, &value)
	}

	arr2 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("Array before calling function is: %v\n", arr2)

	//slices are passed as a value but as they reference the underneath array they are
	//passwed as a reference, passing an array does not change the value as array are
	//passed as a value type.
	//makesquare(sl1)
	sl1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//sl1 = sl1[1:5]
	makesquare(sl1)
	fmt.Printf("Array after calling function is: %v\n", arr2)
	fmt.Printf("Slice after calling function is: %v\n", sl1)
}
func makesquare(sl1 []int) {
	sl1 = sl1[1:5]
	for ind, num := range sl1 {
		sl1[ind] = num * num
	}
}
