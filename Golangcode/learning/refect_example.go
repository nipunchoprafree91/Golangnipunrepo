package main

import (
	"fmt"
	"reflect"

)

type Rectangle struct {
	Width  int
	Height int
}

func main() {
	r := Rectangle{Width: 5, Height: 10}
	rType := reflect.TypeOf(r)
	rKind := rType.Kind()
	rValue := reflect.ValueOf(r)

	fmt.Println(rType)
	fmt.Println(rKind)
	fmt.Println(rValue)
	fmt.Println(rValue.Interface())
}
