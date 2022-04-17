package main

import (
	"Package1"
	"fmt"

)

func main() {
	fmt.Println("pckg1 struct", Package1.Mystr1())
	fmt.Println("pckg2 struct", Package1.Mystr1())
}
