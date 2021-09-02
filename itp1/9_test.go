package itp1

import (
	"aizu/support"
	"fmt"
)

func ExampleFindingAWord() {
	err := support.Stdin(FindingAWord, "computer\nNurtures computer scientists and highly skilled computer engineers\nwho will create and exploit knowledge for the new era\nProvides an outstanding computer environment\nEND_OF_TEXT")
	if err != nil {
		fmt.Printf("failed: %#v\n", err)
	}
	// Output: 3
}

func ExampleShuffle() {
	err := support.Stdin(Shuffle, "aabc\n3\n1\n2\n1\nvwxyz\n2\n3\n4\n-")
	if err != nil {
		fmt.Printf("failed: %#v\n", err)
	}
	// Output: aabc
	// xyzvw
}

func ExampleCardGame() {
	err := support.Stdin(CardGame, "3\ncat dog\nfish fish\nlion tiger")
	if err != nil {
		fmt.Printf("failed: %#v\n", err)
	}
	// Output: 1 7
}

func ExampleTransformation() {
	samples := []string{
		"abcde\n3\nreplace 1 3 xyz\nreverse 0 2\nprint 1 4",
		"xyz\n3\nprint 0 2\nreplace 0 2 abc\nprint 0 2",
	}
	for _, s := range samples {
		err := support.Stdin(Transformation, s)
		if err != nil {
			fmt.Printf("failed: %#v\n", err)
		}
	}
	// Output: xaze
	// xyz
	// abc
}
