package itp1

import (
	"aizu/support"
	"fmt"
)

func ExamplePrintARectangle() {
	err := support.Stdin(PrintARectangle, "3 4\n5 6\n2 2\n0 0")
	if err != nil {
		fmt.Printf("failed: %#v\n", err)
	}
	// Output: ####
	// ####
	// ####
	//
	// ######
	// ######
	// ######
	// ######
	// ######
	//
	// ##
	// ##
	//
}

func ExamplePrintAFrame() {
	err := support.Stdin(PrintAFrame, "3 4\n5 6\n3 3\n0 0")
	if err != nil {
		fmt.Printf("failed: %#v\n", err)
	}
	// Output: 	####
	// #..#
	// ####
	//
	// ######
	// #....#
	// #....#
	// #....#
	// ######
	//
	// ###
	// #.#
	// ###
	//
}

func ExamplePrintAChessboard() {
	err := support.Stdin(PrintAChessboard, "3 4\n5 6\n3 3\n2 2\n1 1\n0 0")
	if err != nil {
		fmt.Printf("failed: %#v\n", err)
	}
	// Output: #.#.
	// .#.#
	// #.#.
	//
	// #.#.#.
	// .#.#.#
	// #.#.#.
	// .#.#.#
	// #.#.#.
	//
	// #.#
	// .#.
	// #.#
	//
	// #.
	// .#
	//
	// #
	//
}

func ExampleStructuredProgramming() {
	err := support.Stdin(StructuredProgramming, "30")
	if err != nil {
		fmt.Printf("failed: %#v\n", err)
	}
	// Output:  3 6 9 12 13 15 18 21 23 24 27 30
}
