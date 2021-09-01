package itp1

import (
	"aizu/support"
	"fmt"
)

func ExamplePrintManyHelloWorld() {
	PrintManyHelloWorld() // TODO: write test code
	// Output: Hello World
}

func ExamplePrintTestCases() {
	err := support.Stdin(PrintTestCases, "3\n5\n11\n7\n8\n19\n0")
	if err != nil {
		fmt.Printf("failed: %#v\n", err)
	}
	// Output: Case 1: 3
	// Case 2: 5
	// Case 3: 11
	// Case 4: 7
	// Case 5: 8
	// Case 6: 19
}

func ExampleSwappingTwoNumbers() {
	err := support.Stdin(SwappingTwoNumbers, "3 2\n2 2\n5 3\n0 0")
	if err != nil {
		fmt.Printf("failed: %#v\n", err)
	}
	// Output: 2 3
	// 2 2
	// 3 5
}

func ExampleHowManyDivisors() {
	err := support.Stdin(HowManyDivisors, "5 14 80")
	if err != nil {
		fmt.Printf("failed: %#v\n", err)
	}
	// Output: 3
}
