package documentationExample

import "fmt"

// https://pkg.go.dev/testing#hdr-Examples.
// Function name begins with Example.
// Compile time checking when running go test -v ./... by adding // output:
// Empty output represented by // output:
// One output per line
func ExampleSum() {
	fmt.Println(Sum(1, 2))
	fmt.Println(Sum(1, 3))
	fmt.Println(Sum(1, 4))
	// output: 3
	// 4
	// 5
}
