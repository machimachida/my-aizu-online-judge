package alds1

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func BubbleSort() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	sc.Scan()
	N, err := strconv.Atoi(sc.Text())
	if err != nil {
		fmt.Println(err)
		return
	}
	if N < 1 || N > 100 {
		fmt.Println("invalid N range")
		return
	}

	A := make([]int, N)
	for i := 0; i < N; i++ {
		sc.Scan()
		A[i], err = strconv.Atoi(sc.Text())
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	var swp int
	for i := 0; i < N; i++ {
		for j := N - 1; j > i; j-- {
			if A[j] < A[j-1] {
				A[j], A[j-1] = A[j-1], A[j]
				swp++
			}
		}
	}
	for i := 0; i < N; i++ {
		fmt.Print(A[i])
		if i != N-1 {
			fmt.Print(" ")
		} else {
			fmt.Print("\n")
		}
	}
	fmt.Println(swp)
}

func SelectionSort() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	sc.Scan()
	N, err := strconv.Atoi(sc.Text())
	if err != nil {
		fmt.Println(err)
		return
	}
	if N < 1 || N > 100 {
		fmt.Println("invalid N range")
		return
	}

	A := make([]int, N)
	for i := 0; i < N; i++ {
		sc.Scan()
		A[i], err = strconv.Atoi(sc.Text())
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	var swp int
	for i := 0; i < N; i++ {
		minI := i
		for j := i; j < N; j++ {
			if A[j] < A[minI] {
				minI = j
			}
		}
		if i != minI {
			A[i], A[minI] = A[minI], A[i]
			swp++
		}
	}
	for i := 0; i < N; i++ {
		fmt.Print(A[i])
		if i != N-1 {
			fmt.Print(" ")
		} else {
			fmt.Print("\n")
		}
	}
	fmt.Println(swp)
}

func StableSort() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	sc.Scan()
	N, err := strconv.Atoi(sc.Text())
	if err != nil {
		fmt.Println(err)
		return
	}
	if N < 1 || N > 36 {
		fmt.Println("invalid N range")
		return
	}

	type card struct {
		suit       rune
		value      int
		inputOrder int
	}

	A := make([]card, N)
	for i := 0; i < N; i++ {
		sc.Scan()
		c := sc.Text()
		A[i].suit = rune(c[0])
		A[i].value, err = strconv.Atoi(string(c[1]))
		A[i].inputOrder = i
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	execSortAndPrint := func(sorting func(C []card)) {
		C := make([]card, N)
		_ = copy(C, A)
		sorting(C)
		isStable := true
		for i := 0; i < N; i++ {
			fmt.Printf("%c%d", C[i].suit, C[i].value)
			if i != N-1 {
				fmt.Print(" ")
			} else {
				fmt.Print("\n")
			}
			if i != 0 && C[i].value == C[i-1].value && C[i].inputOrder < C[i-1].inputOrder {
				isStable = false
			}
		}
		if isStable {
			fmt.Println("Stable")
		} else {
			fmt.Println("Not stable")
		}
	}

	execSortAndPrint(func(C []card) {
		for i := 0; i < N; i++ {
			for j := N - 1; j > i; j-- {
				if C[j].value < C[j-1].value {
					C[j], C[j-1] = C[j-1], C[j]
				}
			}
		}
	})

	execSortAndPrint(func(C []card) {
		for i := 0; i < N; i++ {
			minI := i
			for j := i; j < N; j++ {
				if C[j].value < C[minI].value {
					minI = j
				}
			}
			if i != minI {
				C[i], C[minI] = C[minI], C[i]
			}
		}
	})
}

func ShellSort() {
	sc := bufio.NewScanner(os.Stdin)

	sc.Scan()
	n, err := strconv.Atoi(sc.Text())
	if err != nil {
		fmt.Println(err)
		return
	}
	if n < 1 || n > 1000000 {
		fmt.Println("invalid n range")
		return
	}

	A := make([]int, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		A[i], err = strconv.Atoi(sc.Text())
		if err != nil {
			fmt.Println(err)
			return
		}
		if A[i] < 0 || A[i] > 1000000000 {
			fmt.Println("invalid Ai range")
			return
		}
	}

	var cnt int
	var m int
	for i := 1; i <= n; i = i*3 + 1 {
		m++
	}
	G := make([]int, m)
	for i := 0; i < m; i++ {
		G[i] = int((math.Pow(3, float64(m-i)) - 1) / 2)
	}
	for i := 0; i < m; i++ {
		for j := G[i]; j < n; j++ {
			v := A[j]
			k := j - G[i]
			for k >= 0 && A[k] > v {
				A[k+G[i]] = A[k]
				k = k - G[i]
				cnt++
			}
			A[k+G[i]] = v
		}
	}

	fmt.Println(m)
	for i := 0; i < m; i++ {
		fmt.Print(G[i])
		if i != m-1 {
			fmt.Print(" ")
		} else {
			fmt.Print("\n")
		}
	}
	fmt.Println(cnt)
	for i := 0; i < n; i++ {
		fmt.Println(A[i])
	}
}
