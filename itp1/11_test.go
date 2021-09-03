package itp1

import (
	"aizu/support"
	"fmt"
)

func ExampleDice1() {
	samples := []string{
		"1 2 4 8 16 32\nSE",
		"1 2 4 8 16 32\nEESWN",
	}
	for _, s := range samples {
		err := support.Stdin(Dice1, s)
		if err != nil {
			fmt.Printf("failed: %#v\n", err)
		}
	}
	// Output: 8
	// 32
}

func ExampleDice2() {
	err := support.Stdin(Dice2, "1 2 3 4 5 6\n3\n6 5\n1 3\n3 2")
	if err != nil {
		fmt.Printf("failed: %#v\n", err)
	}
	// Output: 3
	// 5
	// 6
}

func ExampleDice3() {
	samples := []string{
		"1 2 3 4 5 6\n6 2 4 3 5 1",
		"1 2 3 4 5 6\n6 5 4 3 2 1",
	}
	for _, s := range samples {
		err := support.Stdin(Dice3, s)
		if err != nil {
			fmt.Printf("failed: %#v\n", err)
		}
	}
	// Output: Yes
	// No
}

func ExampleDice4() {
	samples := []string{
		"3\n1 2 3 4 5 6\n6 2 4 3 5 1\n6 5 4 3 2 1",
		"3\n1 2 3 4 5 6\n6 5 4 3 2 1\n5 4 3 2 1 6",
	}
	for _, s := range samples {
		err := support.Stdin(Dice4, s)
		if err != nil {
			fmt.Printf("failed: %#v\n", err)
		}
	}
	// Output: No
	// Yes
}
