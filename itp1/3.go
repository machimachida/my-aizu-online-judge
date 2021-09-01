package itp1

import (
	"fmt"
)

func PrintManyHelloWorld() {
	for i := 0; i < 1000; i++ {
		fmt.Println("Hello World")
	}
}

func PrintTestCases() {
	var x int
	for i := 1; i <= 10000; i++ {
		_, err := fmt.Scan(&x)
		if err != nil {
			fmt.Println(err)
			return
		}
		if x == 0 {
			return
		}
		if x < 1 || x > 10000 {
			fmt.Println("invalid x range")
			return
		}
		fmt.Printf("Case %d: %d\n", i, x)
	}
	// 終端記号"0"が10001行目に存在する可能性がある
	_, err := fmt.Scan(&x)
	if err != nil {
		fmt.Println(err)
		return
	}
	if x != 0 {
		fmt.Println("invalid number of datasets")
	}
}

func SwappingTwoNumbers() {
	xy := make([]int, 2)
	for i := 1; i <= 3000; i++ {
		_, err := fmt.Scan(&xy[0], &xy[1])
		if err != nil {
			fmt.Println(err)
			return
		}
		if xy[0] == 0 && xy[1] == 0 {
			return
		}
		for j := range xy {
			if xy[j] < 0 || xy[j] > 10000 {
				fmt.Println("invalid x or y range")
				return
			}
		}

		if xy[0] > xy[1] {
			xy[0], xy[1] = xy[1], xy[0]
		}
		fmt.Printf("%d %d\n", xy[0], xy[1])
	}
	// 終端記号"0 0"が10001行目に存在する可能性がある
	_, err := fmt.Scan(&xy[0], &xy[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	if xy[0] != 0 || xy[1] != 0 {
		fmt.Println("invalid number of datasets")
	}
}

func HowManyDivisors() {
	abc := make([]int, 3)
	_, err := fmt.Scan(&abc[0], &abc[1], &abc[2])
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := range abc {
		if abc[i] < 1 || abc[i] > 10000 {
			fmt.Println("invalid a or b or c range")
			return
		}
	}

	if abc[0] > abc[1] {
		fmt.Println("a > b is invalid")
		return
	}

	nd := 0
	for i := abc[0]; i <= abc[1]; i++ {
		if abc[2]%i == 0 {
			nd++
		}
	}
	fmt.Println(nd)
}
