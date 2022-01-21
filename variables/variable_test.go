package variables

import (
	"fmt"
)

// Default values for uninitialised variables
func Example_variables() {
	// string
	var str string
	fmt.Println(str == "")
	// Output: true

	// numbers
	var n int
	fmt.Println(n == 0)
	// true

	// boolean
	var b bool
	fmt.Println(b == false)
	// true

	// interface
	type interfaceType interface {
		Foo()
	}
	var i interfaceType
	fmt.Println(i == nil)
	// true

	// slice
	var slice []int
	fmt.Println(slice == nil)
	// true

	// pointer
	var p *int
	fmt.Println(p == nil)
	// true

	// channel
	var ch chan int
	fmt.Println(ch == nil)
	// true

	// map
	var m map[string]string
	fmt.Println(m == nil)
	// true

	// string
	var f func(s string)
	fmt.Println(f == nil)
	// true
}
