package main

import (
	"fmt"
	"reflect"

)

type mystruct struct {
	a          int64
	b          string
    isEngineer bool
}

type mystruct2 struct{
    struct2 *mystruct
}

func main() {
	struct1 := mystruct{
		a:          10,
		b:          "Nipun",
		isEngineer: true,
    }

    mynewstruct := mystruct2{
        struct2 : &mystruct{
            a:          20,
		    b:          "chopra",
		    isEngineer: true,
        },
    }
	fmt.Printf("MY Struct is: %v\n\n", struct1)
    fmt.Printf("My new struct is %v\n", *mynewstruct.struct2)

	s1Type := reflect.TypeOf(struct1)
	s1value := reflect.ValueOf(struct1)

	fmt.Printf("Value of struct field is: %v\n", s1value)
	fmt.Printf("Type of struct is :%v\n\n", s1Type)

	for i := 0; i < s1Type.NumField(); i++ {
		s1Type := s1Type.Field(i)
		fmt.Printf("Field Name: %v\n", s1Type.Name)
		fmt.Printf("Field Value: %v\n\n", s1value.Field(i))
    }

    s2Type := reflect.TypeOf(*mynewstruct.struct2)
	s2value := reflect.ValueOf(*mynewstruct.struct2)

	fmt.Printf("Value of struct field is: %v\n", s2value)
	fmt.Printf("Type of struct is :%v\n\n", s2Type)

	for i := 0; i < s2Type.NumField(); i++ {
		s2Type := s2Type.Field(i)
		fmt.Printf("Field Name: %v\n", s2Type.Name)
		fmt.Printf("Field Value: %v\n\n", s2value.Field(i))
	}
}
