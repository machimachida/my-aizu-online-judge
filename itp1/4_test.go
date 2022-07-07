package itp1

import (
	"fmt"
	"practice-go/support"
)

func ExampleABProblem() {
	err := support.Stdin(ABProblem, "3 2")
	if err != nil {
		fmt.Printf("failed: %#v\n", err)
	}
	// Output: 1 1 1.50000
}

func ExampleCircle() {
	samples := []string{"2", "3"}
	for _, s := range samples {
		err := support.Stdin(Circle, s)
		if err != nil {
			fmt.Printf("failed: %#v\n", err)
		}
	}
	// Output: 12.566371 12.566371
	// 28.274334 18.849556
}

func ExampleSimpleCalculator() {
	err := support.Stdin(SimpleCalculator, "1 + 2\n56 - 18\n13 * 2\n100 / 10\n27 + 81\n0 ? 0")
	if err != nil {
		fmt.Printf("failed: %#v\n", err)
	}
	// Output: 3
	// 38
	// 26
	// 10
	// 108
}

func ExampleMinMaxAndSum() {
	err := support.Stdin(MinMaxAndSum, "5\n10 1 5 4 17")
	if err != nil {
		fmt.Printf("failed: %#v\n", err)
	}
	// Output: 1 17 37
}
