package itp1

import (
	"fmt"
)

func HelloWorld() {
	fmt.Println("Hello World")
}

func XCubic() {
	var x int
	_, err := fmt.Scan(&x)
	if err != nil {
		fmt.Println(err)
		return
	}
	if x < 1 || x > 100 {
		fmt.Println("invalid x range")
		return
	}
	fmt.Println(x * x * x)
}

func Rectangle() {
	ab := make([]int, 2)
	_, err := fmt.Scan(&ab[0], &ab[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := range ab {
		if ab[i] < 1 || ab[i] > 100 {
			fmt.Println("invalid a or b range")
			return
		}
	}
	fmt.Printf("%d %d\n", ab[0]*ab[1], (ab[0]+ab[1])*2)
}

func Watch() {
	var S int
	_, err := fmt.Scan(&S)
	if err != nil {
		fmt.Println(err)
		return
	}
	if S < 0 || S > 86400 {
		fmt.Println("invalid S range")
		return
	}
	se := S % 60
	S = S / 60
	mi := S % 60
	ho := S / 60
	fmt.Printf("%d:%d:%d\n", ho, mi, se)
}
