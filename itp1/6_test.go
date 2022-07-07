package itp1

import (
	"fmt"
	"practice-go/support"
)

func ExampleReversingNumbers() {
	samples := []string{"5\n1 2 3 4 5", "8\n3 3 4 4 5 8 7 9"}
	for _, s := range samples {
		err := support.Stdin(ReversingNumbers, s)
		if err != nil {
			fmt.Printf("failed: %#v\n", err)
		}
	}
	// Output: 5 4 3 2 1
	// 9 7 8 5 4 4 3 3
}

func ExampleFindingMissingCards() {
	err := support.Stdin(FindingMissingCards, "47\nS 10\nS 11\nS 12\nS 13\nH 1\nH 2\nS 6\nS 7\nS 8\nS 9\nH 6\nH 8\nH 9\nH 10\nH 11\nH 4\nH 5\nS 2\nS 3\nS 4\nS 5\nH 12\nH 13\nC 1\nC 2\nD 1\nD 2\nD 3\nD 4\nD 5\nD 6\nD 7\nC 3\nC 4\nC 5\nC 6\nC 7\nC 8\nC 9\nC 10\nC 11\nC 13\nD 9\nD 10\nD 11\nD 12\nD 13")
	if err != nil {
		fmt.Printf("failed: %#v\n", err)
	}
	// Output: S 1
	// H 3
	// H 7
	// C 12
	// D 8
}

func ExampleOfficialHouse() {
	err := support.Stdin(OfficialHouse, "3\n1 1 3 8\n3 2 2 7\n4 3 8 1")
	if err != nil {
		fmt.Printf("failed: %#v\n", err)
	}
	// Output:  0 0 8 0 0 0 0 0 0 0
	//  0 0 0 0 0 0 0 0 0 0
	//  0 0 0 0 0 0 0 0 0 0
	// ####################
	//  0 0 0 0 0 0 0 0 0 0
	//  0 0 0 0 0 0 0 0 0 0
	//  0 0 0 0 0 0 0 0 0 0
	// ####################
	//  0 0 0 0 0 0 0 0 0 0
	//  0 7 0 0 0 0 0 0 0 0
	//  0 0 0 0 0 0 0 0 0 0
	// ####################
	//  0 0 0 0 0 0 0 0 0 0
	//  0 0 0 0 0 0 0 0 0 0
	//  0 0 0 0 0 0 0 1 0 0
}

func ExampleMatrixVectorMultiplication() {
	err := support.Stdin(MatrixVectorMultiplication, "3 4\n1 2 0 1\n0 3 0 1\n4 1 1 0\n1\n2\n3\n0")
	if err != nil {
		fmt.Printf("failed: %#v\n", err)
	}
	// Output: 5
	// 6
	// 9
}
