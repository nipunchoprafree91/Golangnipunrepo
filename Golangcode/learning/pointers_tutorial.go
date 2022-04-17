package main

import (
	"fmt"
)

func changeval(p *int) {
	fmt.Println("Caliing changevalue function.....")
	*p = 30
}

type emp struct {
	name string
	id int32
}

type diag struct{
	Endp *emp
}

// type common struct {
// 	Endp emp `json:"Endp,omitempty"`
// }

func main() {
	a := 10
	fmt.Printf("value of a is  : %v\n", a)
	fmt.Printf("Address of a is: %v\n\n", &a)

	//Assigning a new pointer to address of old value
	b := &a
	fmt.Printf("value at address %v is   : %v\n", b, *b)
	fmt.Printf("Address of Variable b is : %v\n\n", &b)

	//Modifying the value of original variable
	*b = 20
	fmt.Printf("New value of a is         : %v\n", a)
	fmt.Printf("New value B is pointing is: %v\n\n", *b)

	//use of new built-in function to allocate the value
	a = 10
	fmt.Printf("value of a is  : %v\n", a)
	fmt.Printf("Address of a is: %v\n\n", &a)

	c := new(int)
	//Not pointing to anything yet
	fmt.Printf("value at mem %v is: %v\n", c, *c)
	c = &a
	fmt.Printf("value at mem %v is: %v\n", c, *c)
	fmt.Printf("Address of c is         : %v\n\n", &c)

	a = 10
	fmt.Printf("value of a is  : %v\n", a)
	fmt.Printf("Address of a is: %v\n\n", &a)

	//changing value by changeValue func
	d := &a
	changeval(d)
	fmt.Printf("value at mem %v is: %v\n", d, *d)
	fmt.Printf("Address of d is: %v\n\n", &d)

	empVar := emp{"nipun", 20}
	empPtr := &emp{"nipun",30}

	fmt.Println("empvar name : ", empVar.name)
	fmt.Println("empvar id : ", empVar.id)

	fmt.Println("empp name : ", empPtr.name)
	fmt.Println("empp id : ", empPtr.id)


	// commonStr := common{"common",40}
	diagStr := diag{}
	diagStr.Endp.name = "diagstruct initialized"

	// fmt.Println("Name common  : ",  commonStr.name)
	fmt.Println("Name Diag : ",     *diagStr.Endp.name)

}
