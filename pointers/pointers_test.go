package pointers

import (
	"fmt"
)

func Example_pointers() {
	var x int
	var p *int                            // *int == pointer to int
	p = &x                                // &x == address of x
	fmt.Print(fmt.Sprintf("%v", p) != "") // *p == value at address
	// Output: true

	x = 1
	p = &x // p, of type *int, points to x
	fmt.Println(*p)
	// 1

	*p = 2 // equivalent to x = 2
	fmt.Println(x)
	// 2
}
