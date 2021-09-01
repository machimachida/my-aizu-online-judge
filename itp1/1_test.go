package itp1

import (
	"aizu/support"
	"fmt"
)

func ExampleHelloWorld() {
	HelloWorld()
	// Output: Hello World
}

func ExampleXCubic() {
	samples := []string{"2", "3"}
	for _, s := range samples {
		err := support.Stdin(XCubic, s)
		if err != nil {
			fmt.Printf("failed: %#v\n", err)
		}
	}
	// Output: 8
	// 27
}

func ExampleRectangle() {
	err := support.Stdin(Rectangle, "3 5")
	if err != nil {
		fmt.Printf("failed: %#v\n", err)
	}
	// Output: 15 16
}

func ExampleWatch() {
	err := support.Stdin(Watch, "46979")
	if err != nil {
		fmt.Printf("failed: %#v\n", err)
	}
	// Output: 13:2:59
}
