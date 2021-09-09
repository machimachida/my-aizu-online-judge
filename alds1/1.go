package alds1

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func InsersionSort() {
	var N int
	_, err := fmt.Scan(&N)
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
		_, err := fmt.Scan(&A[i])
		if err != nil {
			fmt.Println(err)
			return
		}
		if A[i] < 0 || A[i] > 1000 {
			fmt.Println("invalid A's element range")
			return
		}
	}

	for k := 0; k < N; k++ {
		fmt.Print(A[k])
		if k != N-1 {
			fmt.Print(" ")
		} else {
			fmt.Print("\n")
		}
	}
	for i := 1; i < N; i++ {
		v := A[i]
		j := i - 1
		for j >= 0 && A[j] > v {
			A[j+1] = A[j]
			j--
		}
		A[j+1] = v
		for k := 0; k < N; k++ {
			fmt.Print(A[k])
			if k != N-1 {
				fmt.Print(" ")
			} else {
				fmt.Print("\n")
			}
		}
	}
}

func GreatestCommonDivisor() {
	var a int
	var b int
	_, err := fmt.Scan(&a, &b)
	if err != nil {
		fmt.Println(err)
		return
	}
	if a < 1 || a > 1000000000 || b < 1 || b > 1000000000 {
		fmt.Println("invalid a or b range")
		return
	}
	if a < b {
		a, b = b, a
	}

	for {
		c := a % b
		if c == 0 {
			break
		}
		a, b = b, c
	}
	fmt.Println(b)
}

func PrimeNumbers() {
	var N int
	_, err := fmt.Scan(&N)
	if err != nil {
		fmt.Println(err)
		return
	}
	if N < 1 || N > 10000 {
		fmt.Println("invalid N range")
		return
	}

	var el int
	var cnt int
	for i := 0; i < N; i++ {
		_, err := fmt.Scan(&el)
		if err != nil {
			fmt.Println(err)
			return
		}
		if el < 2 || el > 100000000 {
			fmt.Println("invalid element range")
			return
		}

		if el <= 1 {
			continue
		}
		if el == 2 {
			cnt++
			continue
		}
		if el%2 == 0 {
			continue
		}

		isPrime := true
		for j := 3; j < int(math.Ceil(math.Sqrt(float64(el))))+1; j = j + 2 {
			if el%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			cnt++
		}
	}
	fmt.Println(cnt)
}

func MaximumProfit() {
	sc := bufio.NewScanner(os.Stdin)

	sc.Scan()
	n, err := strconv.Atoi(sc.Text())
	if err != nil {
		fmt.Println(err)
		return
	}
	if n < 2 || n > 200000 {
		fmt.Println("invalid n range")
		return
	}

	var Rt int
	minR := 2000000000
	max := -1000000000
	for i := 0; i < n; i++ {
		sc.Scan()
		Rt, err = strconv.Atoi(sc.Text())
		if err != nil {
			fmt.Println(err)
			return
		}
		if Rt < 1 || Rt > 1000000000 {
			fmt.Println("invalid Rt range")
			return
		}
		if diff := Rt - minR; diff > max {
			max = diff
		}
		if Rt < minR {
			minR = Rt
		}
	}
	fmt.Println(max)
}
