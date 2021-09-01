package itp1

import (
	"aizu/support"
	"fmt"
)

func ExampleSmallLargeEqual() {
	samples := []string{"1 2", "4 3", "5 5"}
	for _, s := range samples {
		err := support.Stdin(SmallLargeEqual, s)
		if err != nil {
			fmt.Printf("failed: %#v\n", err)
		}
	}
	// Output: a < b
	// a > b
	// a == b
}

func ExampleRange() {
	samples := []string{"1 3 8", "3 8 1"}
	for _, s := range samples {
		err := support.Stdin(Range, s)
		if err != nil {
			fmt.Printf("failed: %#v\n", err)
		}
	}
	// Output: Yes
	// No
}

func ExampleSortingThreeNumbers() {
	err := support.Stdin(SortingThreeNumbers, "3 8 1")
	if err != nil {
		fmt.Printf("failed: %#v\n", err)
	}
	// Output: 1 3 8
}

func ExampleCircleInARectangle() {
	samples := []string{"5 4 2 2 1", "5 4 2 4 1"}
	for _, s := range samples {
		err := support.Stdin(CircleInARectangle, s)
		if err != nil {
			fmt.Printf("failed: %#v\n", err)
		}
	}
	// Output: Yes
	// No
}
