package Datatypes

import (
	"fmt"
	"reflect"
	"unsafe"
)

func Calculate_powers(num int64, exp int64) int64 {
	//want to multiply 2 number of time the exponent and result is -1
	var result int64
	result = 1

	if exp > 0 {
		result = num * (Calculate_powers(num, exp-1))
	}
	return result
}

func Rundatatypes() {
	fmt.Println("Learning all data Types")
	var a uint8
	a = uint8(Calculate_powers(2, 8)) - 1
	fmt.Println("Type of a is : ", reflect.TypeOf(a))
	fmt.Printf("Max Value of uint8 is : %d\n", a)
	fmt.Printf("bytes in a is : %d\n", unsafe.Sizeof(a))

	var b uint16
	b = uint16(Calculate_powers(2, 16)) - 1
	fmt.Println("Type of b is : ", reflect.TypeOf(b))
	fmt.Printf("Max Value of uint16 is : %d\n", b)
	fmt.Printf("bytes in b is : %d\n", unsafe.Sizeof(b))

	var c uint32
	c = uint32(Calculate_powers(2, 32)) - 1
	fmt.Println("Type of c is : ", reflect.TypeOf(c))
	fmt.Printf(" Max Value of uint32 c is : %d\n", c)
	fmt.Printf("bytes in c is : %d\n", unsafe.Sizeof(c))

	var d uint64
	d = uint64(Calculate_powers(2, 64)) - 1
	fmt.Println("Type of d is : ", reflect.TypeOf(d))
	fmt.Printf("Max Value of uint64 is : %d\n", d)
	fmt.Printf("bytes in d is : %d\n", unsafe.Sizeof(d))
}
