package itp1

import (
	"aizu/support"
	"fmt"
)

func ExampleTogglingCases() {
	err := support.Stdin(TogglingCases, "fAIR, LATER, OCCASIONALLY CLOUDY.")
	if err != nil {
		fmt.Printf("failed: %#v\n", err)
	}
	// Output: Fair, later, occasionally cloudy.
}

func ExampleSumOfNumbers() {
	err := support.Stdin(SumOfNumbers, "123\n55\n1000\n0")
	if err != nil {
		fmt.Printf("failed: %#v\n", err)
	}
	// Output: 6
	// 10
	// 1
}

func ExampleCountingCharacters() {
	err := support.Stdin(CountingCharacters, "This is a pen.")
	if err != nil {
		fmt.Printf("failed: %#v\n", err)
	}
	// Output: a : 1
	// b : 0
	// c : 0
	// d : 0
	// e : 1
	// f : 0
	// g : 0
	// h : 1
	// i : 2
	// j : 0
	// k : 0
	// l : 0
	// m : 0
	// n : 1
	// o : 0
	// p : 1
	// q : 0
	// r : 0
	// s : 2
	// t : 1
	// u : 0
	// v : 0
	// w : 0
	// x : 0
	// y : 0
	// z : 0
}

func ExampleRing() {
	samples := []string{"vanceknowledgetoad\nadvance", "vanceknowledgetoad\nadvanced"}
	for _, s := range samples {
		err := support.Stdin(Ring, s)
		if err != nil {
			fmt.Printf("failed: %#v\n", err)
		}
	}
	// Output: Yes
	// No
}
