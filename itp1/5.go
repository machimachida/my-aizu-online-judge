package itp1

import (
	"fmt"
	"strings"
)

func PrintARectangle() {
	HW := make([]int, 2)
	for {
		_, err := fmt.Scan(&HW[0], &HW[1])
		if err != nil {
			fmt.Println(err)
			return
		}
		if HW[0] == 0 && HW[1] == 0 {
			return
		}
		for i := range HW {
			if HW[i] < 1 || HW[i] > 300 {
				fmt.Println("invalid H or W range")
				return
			}
		}

		for i := 0; i < HW[0]; i++ {
			fmt.Println(strings.Repeat("#", HW[1]))
		}
		fmt.Print("\n")
	}
}

func PrintAFrame() {
	HW := make([]int, 2)
	for {
		_, err := fmt.Scan(&HW[0], &HW[1])
		if err != nil {
			fmt.Println(err)
			return
		}
		if HW[0] == 0 && HW[1] == 0 {
			return
		}
		for i := range HW {
			if HW[i] < 3 || HW[i] > 300 {
				fmt.Println("invalid H or W range")
				return
			}
		}

		for i := 0; i < HW[0]; i++ {
			if i == 0 || i == HW[0]-1 {
				fmt.Println(strings.Repeat("#", HW[1]))
			} else {
				fmt.Printf("#%s#\n", strings.Repeat(".", HW[1]-2))
			}
		}
		fmt.Print("\n")
	}
}

func PrintAChessboard() {
	HW := make([]int, 2)
	for {
		_, err := fmt.Scan(&HW[0], &HW[1])
		if err != nil {
			fmt.Println(err)
			return
		}
		if HW[0] == 0 && HW[1] == 0 {
			return
		}
		for i := range HW {
			if HW[i] < 1 || HW[i] > 300 {
				fmt.Println("invalid H or W range")
				return
			}
		}

		for i := 0; i < HW[0]; i++ {
			isHash := i%2 == 0
			for j := 0; j < HW[1]; j++ {
				if isHash {
					fmt.Print("#")
				} else {
					fmt.Print(".")
				}
				isHash = !isHash
			}
			fmt.Print("\n")
		}
		fmt.Print("\n")
	}
}

func StructuredProgramming() {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println(err)
		return
	}
	if n < 3 || n > 10000 {
		fmt.Println("invalid n range")
	}

	for i := 1; i <= n; i++ {
		x := i
		if x%3 == 0 {
			fmt.Printf(" %d", i)
		} else {
			for x != 0 {
				if x%10 == 3 {
					fmt.Printf(" %d", i)
					break
				}
				x = x / 10
			}
		}
	}
	fmt.Print("\n")
}
