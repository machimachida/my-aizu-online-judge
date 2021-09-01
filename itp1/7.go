package itp1

import (
	"fmt"
)

func Grading() {
	mfr := make([]int, 3)
	for i := 0; i < 50; i++ {
		_, err := fmt.Scan(&mfr[0], &mfr[1], &mfr[2])
		if err != nil {
			fmt.Println(err)
			return
		}
		if mfr[0] == -1 && mfr[1] == -1 && mfr[2] == -1 {
			return
		}

		if mfr[0] == -1 || mfr[1] == -1 {
			fmt.Println("F")
		} else {
			for j := range mfr[:1] {
				if mfr[j] < -1 || mfr[j] > 50 {
					fmt.Println("invalid m or f range")
					return
				}
			}
			switch t := mfr[0] + mfr[1]; {
			case t >= 80:
				fmt.Println("A")
			case t >= 65 && t < 80:
				fmt.Println("B")
			case t >= 50 && t < 65:
				fmt.Println("C")
			case t >= 30 && t < 50:
				if mfr[2] < -1 || mfr[2] > 100 {
					fmt.Println("invalid r range")
				} else if mfr[2] >= 50 {
					fmt.Println("C")
				} else {
					fmt.Println("D")
				}
			default:
				fmt.Println("F")
			}
		}
	}
}

func HowManyWays() {
	var n int
	var x int
	for {
		_, err := fmt.Scan(&n, &x)
		if err != nil {
			fmt.Println(err)
			return
		}
		if n == 0 && x == 0 {
			return
		}
		if n < 3 || n > 100 {
			fmt.Println("invalid n range")
			return
		}
		if x < 0 || x > 300 {
			fmt.Println("invalid x range")
			return
		}

		var c int
		for i := 1; i <= n; i++ {
			for j := i + 1; j <= n; j++ {
				ij := i + j
				if ij >= x {
					continue
				}
				for k := j + 1; k <= n; k++ {
					if ij+k == x {
						c++
					}
				}
			}
		}
		fmt.Println(c)
	}
}

func Spreadsheet() {
	rc := make([]int, 2)
	_, err := fmt.Scan(&rc[0], &rc[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := range rc {
		if rc[i] < 1 || rc[i] > 100 {
			fmt.Println("invalid r or c range")
		}
	}

	cs := make([]int, rc[1])
	for i := 0; i < rc[0]; i++ {
		var el int
		var sum int
		for j := 0; j < rc[1]; j++ {
			_, err := fmt.Scan(&el)
			if err != nil {
				fmt.Println(err)
				return
			}
			if el < 0 || el > 100 {
				fmt.Println("invalid element range")
			}
			sum += el
			cs[j] += el
			fmt.Print(el, " ")
		}
		fmt.Println(sum)
	}
	var sum int
	for i := range cs {
		fmt.Print(cs[i], " ")
		sum += cs[i]
	}
	fmt.Println(sum)
}

func MatrixMultiplication() {
	nml := make([]int, 3)
	_, err := fmt.Scan(&nml[0], &nml[1], &nml[2])
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := range nml {
		if nml[i] < 1 || nml[i] > 100 {
			fmt.Println("invalid n or m or l range")
		}
	}

	a := make([][]int, nml[0])
	for i := 0; i < nml[0]; i++ {
		a[i] = make([]int, nml[1])
		for j := 0; j < nml[1]; j++ {
			_, err := fmt.Scan(&a[i][j])
			if err != nil {
				fmt.Println(err)
				return
			}
			if a[i][j] < 0 || a[i][j] > 10000 {
				fmt.Println("invalid aij range")
				return
			}
		}
	}
	b := make([][]int, nml[1])
	for i := 0; i < nml[1]; i++ {
		b[i] = make([]int, nml[2])
		for j := 0; j < nml[2]; j++ {
			_, err := fmt.Scan(&b[i][j])
			if err != nil {
				fmt.Println(err)
				return
			}
			if b[i][j] < 0 || b[i][j] > 10000 {
				fmt.Println("invalid bij range")
				return
			}
		}
	}

	for i := 0; i < nml[0]; i++ {
		cr := make([]int, nml[2])
		for k := 0; k < nml[1]; k++ {
			for j := 0; j < nml[2]; j++ {
				cr[j] += a[i][k] * b[k][j]
			}
		}
		for j := 0; j < nml[2]; j++ {
			fmt.Print(cr[j])
			if j != nml[2]-1 {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
}
