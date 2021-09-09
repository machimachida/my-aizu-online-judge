package alds1

import (
	"aizu/support"
	"fmt"
)

func ExampleInsersionSort() {
	samples := []string{
		"6\n5 2 4 6 1 3",
		"3\n1 2 3",
	}
	for _, s := range samples {
		err := support.Stdin(InsersionSort, s)
		if err != nil {
			fmt.Printf("failed: %#v\n", err)
		}
	}
	// Output: 5 2 4 6 1 3
	// 2 5 4 6 1 3
	// 2 4 5 6 1 3
	// 2 4 5 6 1 3
	// 1 2 4 5 6 3
	// 1 2 3 4 5 6
	// 1 2 3
	// 1 2 3
	// 1 2 3
}

func ExampleGreatestCommonDivisor() {
	samples := []string{"54 20", "147 105"}
	for _, s := range samples {
		err := support.Stdin(GreatestCommonDivisor, s)
		if err != nil {
			fmt.Printf("failed: %#v\n", err)
		}
	}
	// Output: 2
	// 21
}

func ExamplePrimeNumbers() {
	samples := []string{"5\n2\n3\n4\n5\n6", "11\n7\n8\n9\n10\n11\n12\n13\n14\n15\n16\n17"}
	for _, s := range samples {
		err := support.Stdin(PrimeNumbers, s)
		if err != nil {
			fmt.Printf("failed: %#v\n", err)
		}
	}
	// Output: 3
	// 4
}

func ExampleMaximumProfit() {
	samples := []string{"6\n5\n3\n1\n3\n4\n3", "3\n4\n3\n2"}
	for _, s := range samples {
		err := support.Stdin(MaximumProfit, s)
		if err != nil {
			fmt.Printf("failed: %#v\n", err)
		}
	}
	// Output: 3
	// -1
}
