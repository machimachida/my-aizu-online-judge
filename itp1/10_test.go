package itp1

import (
	"fmt"
	"practice-go/support"
)

func ExampleDistance() {
	err := support.Stdin(Distance, "0 0 1 1")
	if err != nil {
		fmt.Printf("failed: %#v\n", err)
	}
	// Output: 1.41421356
}

func ExampleTriangle() {
	err := support.Stdin(Triangle, "4 3 90")
	if err != nil {
		fmt.Printf("failed: %#v\n", err)
	}
	// Output: 6.00000000
	// 12.00000000
	// 3.00000000
}

func ExampleStandardDeviation() {
	err := support.Stdin(StandardDeviation, "5\n70 80 100 90 20\n3\n80 80 80\n0")
	if err != nil {
		fmt.Printf("failed: %#v\n", err)
	}
	// Output: 27.85677655
	// 0.00000000
}

func ExampleDistance2() {
	err := support.Stdin(Distance2, "3\n1 2 3\n2 0 4")
	if err != nil {
		fmt.Printf("failed: %#v\n", err)
	}
	// Output: 4.000000
	// 2.449490
	// 2.154435
	// 2.000000
}
