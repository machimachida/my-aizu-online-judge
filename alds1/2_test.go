package alds1

import (
	"fmt"
	"practice-go/support"
)

func ExampleBubbleSort() {
	samples := []string{"5\n5 3 2 4 1", "6\n5 2 4 6 1 3"}
	for _, s := range samples {
		err := support.Stdin(BubbleSort, s)
		if err != nil {
			fmt.Printf("failed: %#v\n", err)
		}
	}
	// Output: 1 2 3 4 5
	// 8
	// 1 2 3 4 5 6
	// 9
}

func ExampleSelectionSort() {
	samples := []string{"6\n5 6 4 2 1 3", "6\n5 2 4 6 1 3"}
	for _, s := range samples {
		err := support.Stdin(SelectionSort, s)
		if err != nil {
			fmt.Printf("failed: %#v\n", err)
		}
	}
	// Output: 1 2 3 4 5 6
	// 4
	// 1 2 3 4 5 6
	// 3
}

func ExampleStableSort() {
	samples := []string{"5\nH4 C9 S4 D2 C3", "2\nS1 H1"}
	for _, s := range samples {
		err := support.Stdin(StableSort, s)
		if err != nil {
			fmt.Printf("failed: %#v\n", err)
		}
	}
	// Output: D2 C3 H4 S4 C9
	// Stable
	// D2 C3 S4 H4 C9
	// Not stable
	// S1 H1
	// Stable
	// S1 H1
	// Stable
}

func ExampleShellSort() {
	samples := []string{"5\n5\n1\n4\n3\n2", "3\n3\n2\n1"}
	for _, s := range samples {
		err := support.Stdin(ShellSort, s)
		if err != nil {
			fmt.Printf("failed: %#v\n", err)
		}
	}
	// Output: 2
	// 4 1
	// 3
	// 1
	// 2
	// 3
	// 4
	// 5
	// 1
	// 1
	// 3
	// 1
	// 2
	// 3
}
