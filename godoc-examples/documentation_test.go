package godoc_examples

import "fmt"

// Compile time checking when running go test -v ./... by adding // output:
// http://localhost:6060/pkg/github.com/valantonini/go-kaizen/documentationExample/#pkg-examples
// https://pkg.go.dev/testing#hdr-Examples.

// Function name begins with Example. *Must* include an output comment
func ExampleSum() {
	// One output check per line
	fmt.Println(Sum(1, 2))
	fmt.Println(Sum(1, 3))
	fmt.Println(Sum(1, 4))
	// output: 3
	// 4
	// 5
}

// for multiple, append with _second _third
func ExampleSum_second() {
	// no output
	Sum(1, 2)
	// output:
}
