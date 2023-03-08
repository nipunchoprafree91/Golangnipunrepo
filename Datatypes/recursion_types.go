// Recursion is a concept where a function calls itself
// by direct or indirect means. Each call to the recursive
// function is a smaller version so that it converges at some point.
// Every recursive function has a base case or base condition which
// is the final executable statement in recursion and halts further calls.
// There are different types of recursion as explained in the following examples:

package Datatypes

import (
	"fmt"
)

//Direct - factorial
func DirectRecursion(number int) int {
	var fact int
	if number == 0 || number == 1 {
		fact = 1
		return fact
	}
	if number < 0 {
		fact = -1
		fmt.Println("Invalid number")
		return fact
	}

	fact = number * DirectRecursion(number-1)
	fmt.Println(number)
	return fact

}
