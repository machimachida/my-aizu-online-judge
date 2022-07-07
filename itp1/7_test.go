package itp1

import (
	"fmt"
	"practice-go/support"
)

func ExampleGrading() {
	err := support.Stdin(Grading, "40 42 -1\n20 30 -1\n0 2 -1\n-1 -1 -1")
	if err != nil {
		fmt.Printf("failed: %#v\n", err)
	}
	// Output: A
	// C
	// F
}

func ExampleHowManyWays() {
	err := support.Stdin(HowManyWays, "5 9\n0 0")
	if err != nil {
		fmt.Printf("failed: %#v\n", err)
	}
	// Output: 2
}

func ExampleSpreadsheet() {
	err := support.Stdin(Spreadsheet, "4 5\n1 1 3 4 5\n2 2 2 4 5\n3 3 0 1 1\n2 3 4 4 6")
	if err != nil {
		fmt.Printf("failed: %#v\n", err)
	}
	// 	Output: 1 1 3 4 5 14
	// 2 2 2 4 5 15
	// 3 3 0 1 1 8
	// 2 3 4 4 6 19
	// 8 9 9 13 17 56
}

func ExampleMatrixMultiplication() {
	err := support.Stdin(MatrixMultiplication, "3 2 3\n1 2\n0 3\n4 5\n1 2 1\n0 3 2")
	if err != nil {
		fmt.Printf("failed: %#v\n", err)
	}
	// Output: 1 8 5
	// 0 9 6
	// 4 23 14
}
