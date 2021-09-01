package itp1

import (
	"fmt"
	"sort"
)

func SmallLargeEqual() {
	ab := make([]int, 2)
	_, err := fmt.Scan(&ab[0], &ab[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := range ab {
		if ab[i] < -1000 || ab[i] > 1000 {
			fmt.Println("invalid a or b range")
			return
		}
	}

	if ab[0] < ab[1] {
		fmt.Println("a < b")
	} else if ab[0] > ab[1] {
		fmt.Println("a > b")
	} else {
		fmt.Println("a == b")
	}
}

func Range() {
	abc := make([]int, 3)
	_, err := fmt.Scan(&abc[0], &abc[1], &abc[2])
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := range abc {
		if abc[i] < 0 || abc[i] > 100 {
			fmt.Println("invalid a or b or c range")
		}
	}

	if abc[0] < abc[1] && abc[1] < abc[2] {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func SortingThreeNumbers() {
	is := make([]int, 3)
	_, err := fmt.Scan(&is[0], &is[1], &is[2])
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := range is {
		if is[i] < 1 || is[i] > 10000 {
			fmt.Println("invalid integer range")
		}
	}
	sort.Ints(is)
	fmt.Printf("%d %d %d\n", is[0], is[1], is[2])
}

func CircleInARectangle() {
	is := make([]int, 5)
	_, err := fmt.Scan(&is[0], &is[1], &is[2], &is[3], &is[4])
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := range is {
		if i == 2 || i == 3 {
			if is[i] < -100 || is[i] > 100 {
				fmt.Println("invalid x or y range")
			}
		} else {
			if is[i] <= 0 || is[i] > 100 {
				fmt.Println("invalid W or H or r range")
			}
		}
	}

	if is[2]-is[4] >= 0 &&
		is[3]-is[4] >= 0 &&
		is[2]+is[4] <= is[0] &&
		is[3]+is[4] <= is[1] {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
