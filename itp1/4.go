package itp1

import (
	"fmt"
	"math"
)

func ABProblem() {
	ab := make([]int, 2)
	_, err := fmt.Scan(&ab[0], &ab[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := range ab {
		if ab[i] < 1 || ab[i] > 1000000000 {
			fmt.Println("invalid a or b range")
			return
		}
	}
	fmt.Printf("%d %d %.5f\n", ab[0]/ab[1], ab[0]%ab[1], float64(ab[0])/float64(ab[1]))
}

func Circle() {
	var r float64
	_, err := fmt.Scan(&r)
	if err != nil {
		fmt.Println(err)
		return
	}
	if r <= 0 || r >= 10000 {
		fmt.Println("invalid r range")
		return
	}
	fmt.Printf("%.6f %.6f\n", r*r*math.Pi, 2*r*math.Pi)
}

func SimpleCalculator() {
	var op string
	ab := make([]int, 2)
	for {
		_, err := fmt.Scan(&ab[0], &op, &ab[1])
		if err != nil {
			fmt.Println(err)
			return
		}
		if op == "?" {
			return
		}
		for i := range ab {
			if ab[i] < 0 || ab[i] > 20000 {
				fmt.Println("invalid a or b range")
				return
			}
		}

		switch op {
		case "+":
			fmt.Printf("%d\n", ab[0]+ab[1])
		case "-":
			fmt.Printf("%d\n", ab[0]-ab[1])
		case "*":
			fmt.Printf("%d\n", ab[0]*ab[1])
		case "/":
			if ab[1] == 0 {
				fmt.Println("A division by zero are given.")
				return
			}
			fmt.Printf("%d\n", ab[0]/ab[1])
		}
	}
}

func MinMaxAndSum() {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println(err)
		return
	}
	if n <= 0 || n > 10000 {
		fmt.Println("invalid n range")
		return
	}

	var ai int
	var sum int
	max := -1000000
	min := 1000000
	for i := 0; i < n; i++ {
		_, err := fmt.Scan(&ai)
		if err != nil {
			fmt.Println(err)
			return
		}
		if ai < -1000000 || ai > 1000000 {
			fmt.Println("invalid ai range")
			return
		}
		sum += ai
		if ai > max {
			max = ai
		}
		if ai < min {
			min = ai
		}
	}
	fmt.Printf("%d %d %d\n", min, max, sum)
}
