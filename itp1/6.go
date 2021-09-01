package itp1

import (
	"fmt"
)

func ReversingNumbers() {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println(err)
		return
	}
	if n > 100 {
		fmt.Println("invalid n range")
		return
	}

	an := make([]int, n)
	for i := 0; i < n; i++ {
		_, err := fmt.Scan(&an[i])
		if err != nil {
			fmt.Println(err)
			return
		}
		if an[i] < 0 || an[i] >= 1000 {
			fmt.Println("invalid ai range")
			return
		}
	}

	for i := n - 1; i >= 0; i-- {
		fmt.Print(an[i])
		if i != 0 {
			fmt.Print(" ")
		}
	}
	fmt.Print("\n")
}

func FindingMissingCards() {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println(err)
		return
	}
	if n > 52 {
		fmt.Println("invalid n range")
		return
	}

	exs := make(map[string][]bool, 4)
	suits := []string{"S", "H", "C", "D"}
	for i := range suits {
		exs[suits[i]] = make([]bool, 13)
	}

	for i := 0; i < n; i++ {
		var s string
		var r int
		_, err := fmt.Scan(&s, &r)
		if err != nil {
			fmt.Println(err)
			return
		}
		if s != "S" && s != "H" && s != "C" && s != "D" {
			fmt.Println("invalid suit")
			return
		}
		if r < 1 || r > 13 {
			fmt.Println("invalid rank")
			return
		}
		exs[s][r-1] = true
	}

	for _, s := range suits {
		for i, b := range exs[s] {
			if !b {
				fmt.Printf("%s %d\n", s, i+1)
			}
		}
	}
}

func OfficialHouse() {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println(err)
		return
	}

	oh := make([][]int, 12)
	for i := 0; i < 12; i++ {
		oh[i] = make([]int, 10)
	}
	var b int
	var f int
	var r int
	var t int
	for i := 0; i < n; i++ {
		_, err := fmt.Scan(&b, &f, &r, &t)
		if err != nil {
			fmt.Println(err)
			return
		}
		if b < 1 || b > 4 {
			fmt.Println("invalid building number")
			return
		}
		if f < 1 || f > 3 {
			fmt.Println("incorrect floor number")
			return
		}
		if r < 1 || r > 10 {
			fmt.Println("incorrect room number")
			return
		}
		bf := (b-1)*3 + f - 1
		oh[bf][r-1] += t
		if oh[bf][r-1] < 0 || oh[bf][r-1] > 9 {
			fmt.Println("incorrect number of tenants")
			return
		}
	}

	for i := range oh {
		for j := range oh[i] {
			fmt.Print(" ", oh[i][j])
		}
		fmt.Print("\n")
		if i%3 == 2 && i != len(oh)-1 {
			fmt.Println("####################")
		}
	}
}

func MatrixVectorMultiplication() {
	nm := make([]int, 2)
	_, err := fmt.Scan(&nm[0], &nm[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := range nm {
		if nm[i] < 1 || nm[i] > 100 {
			fmt.Println("invalid n or m range")
		}
	}

	a := make([][]int, nm[0])
	for i := 0; i < nm[0]; i++ {
		a[i] = make([]int, nm[1])
		for j := 0; j < nm[1]; j++ {
			_, err := fmt.Scan(&a[i][j])
			if err != nil {
				fmt.Println(err)
				return
			}
			if a[i][j] < 0 || a[i][j] > 1000 {
				fmt.Println("invalid a range")
				return
			}
		}
	}
	b := make([]int, nm[1])
	for i := 0; i < nm[1]; i++ {
		_, err := fmt.Scan(&b[i])
		if err != nil {
			fmt.Println(err)
			return
		}
		if b[i] < 0 || b[i] > 1000 {
			fmt.Println("invalid b range")
			return
		}
	}

	for i := 0; i < nm[0]; i++ {
		c := 0
		for j := 0; j < nm[1]; j++ {
			c += a[i][j] * b[j]
		}
		fmt.Printf("%d\n", c)
	}
}
