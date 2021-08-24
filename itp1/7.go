package itp1

import (
	"errors"
	"strconv"
	"strings"
)

func Grading(str string) (string, error) {
	ss := strings.Split(str, "\n")
	if len(ss) > 50 {
		return "", errors.New("too many lines")
	}

	res := ""
	for _, s := range ss {
		if s == "-1 -1 -1" {
			break
		}
		ns := strings.Split(s, " ")
		if len(ns) != 3 {
			return "", errors.New("invalid number of scores")
		}

		m, err := strconv.Atoi(ns[0])
		if err != nil {
			return "", err
		}
		if m < -1 || m > 50 {
			return "", errors.New("invalid m")
		}

		f, err := strconv.Atoi(ns[1])
		if err != nil {
			return "", err
		}
		if f < -1 || f > 50 {
			return "", errors.New("invalid f")
		}

		switch t := m + f; {
		case t >= 80:
			res += "A\n"
		case t >= 65 && t < 80:
			res += "B\n"
		case t >= 50 && t < 65:
			res += "C\n"
		case t >= 30 && t < 50:
			r, err := strconv.Atoi(ns[2])
			if err != nil {
				return "", err
			}
			if r < -1 || r > 100 {
				return "", errors.New("invalid r")
			} else if r >= 50 {
				res += "C\n"
			} else {
				res += "D\n"
			}
		default:
			res += "F\n"
		}
	}
	return res[:len(res)-1], nil
}

func HowManyWays(str string) (string, error) {
	ss := strings.Split(str, "\n")

	count := 0
	for _, s := range ss {
		if s == "0 0" {
			break
		}
		ns := strings.Split(s, " ")
		if len(ns) != 2 {
			return "", errors.New("invalid number of integers")
		}

		n, err := strconv.Atoi(ns[0])
		if err != nil {
			return "", err
		}
		if n < 3 || n > 100 {
			return "", errors.New("invalid n")
		}

		x, err := strconv.Atoi(ns[1])
		if err != nil {
			return "", err
		}
		if x < 0 || x > 300 {
			return "", errors.New("invalid x")
		}

		for i := 1; i <= n; i++ {
			for j := i + 1; j <= n; j++ {
				ij := i + j
				if ij >= x {
					continue
				}
				for k := j + 1; k <= n; k++ {
					if ij+k == x {
						count++
					}
				}
			}
		}
	}
	return strconv.Itoa(count), nil
}

func Spreadsheet(str string) (string, error) {
	ss := strings.Split(str, "\n")
	rc := strings.Split(ss[0], " ")
	if len(rc) != 2 {
		return "", errors.New("invalid row and column number")
	}
	r, err := strconv.Atoi(rc[0])
	if err != nil {
		return "", err
	}
	if r < 1 || r > 100 {
		return "", errors.New("invalid r")
	}

	c, err := strconv.Atoi(rc[1])
	if err != nil {
		return "", err
	}
	if c < 1 || c > 100 {
		return "", errors.New("invalid c")
	}

	res := ""
	cs := make([]int, 5)
	for _, row := range ss[1:] {
		els := strings.Split(row, " ")
		if len(els) != c {
			return "", errors.New("invalid columns number")
		}
		var sum int
		for j, el := range els {
			i, err := strconv.Atoi(el)
			if err != nil {
				return "", err
			}
			if i < 0 || i > 100 {
				return "", errors.New("invalid element")
			}
			sum += i
			cs[j] += i
		}

		res += row + " " + strconv.Itoa(sum) + "\n"
	}

	var sum int
	for _, c := range cs {
		res += strconv.Itoa(c) + " "
		sum += c
	}
	return res + strconv.Itoa(sum), nil
}

func MatrixMultiplication(str string) (string, error) {
	ss := strings.Split(str, "\n")
	nml := strings.Split(ss[0], " ")
	if len(nml) != 3 {
		return "", errors.New("invalid n m l")
	}

	n, err := strconv.Atoi(nml[0])
	if err != nil {
		return "", err
	}
	if n < 1 || n > 100 {
		return "", errors.New("invalid r")
	}

	m, err := strconv.Atoi(nml[1])
	if err != nil {
		return "", err
	}
	if m < 1 || m > 100 {
		return "", errors.New("invalid m")
	}

	l, err := strconv.Atoi(nml[2])
	if err != nil {
		return "", err
	}
	if l < 1 || l > 100 {
		return "", errors.New("invalid l")
	}

	if len(ss) != n+m+1 {
		return "", errors.New("invalid input")
	}

	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int, m)
		els := strings.Split(ss[i+1], " ")
		if len(els) != m {
			return "", errors.New("invalid row a")
		}
		for j := 0; j < m; j++ {
			a[i][j], err = strconv.Atoi(els[j])
			if err != nil {
				return "", err
			}
		}
	}
	b := make([][]int, m)
	for i := 0; i < m; i++ {
		b[i] = make([]int, l)
		els := strings.Split(ss[i+n+1], " ")
		if len(els) != l {
			return "", errors.New("invalid row b")
		}
		for j := 0; j < l; j++ {
			b[i][j], err = strconv.Atoi(els[j])
			if err != nil {
				return "", err
			}
		}

	}

	c := make([][]int, n)
	for i := 0; i < n; i++ {
		c[i] = make([]int, l)
		for k := 0; k < m; k++ {
			for j := 0; j < l; j++ {
				c[i][j] += a[i][k] * b[k][j]
			}
		}
	}

	var res string
	for i := 0; i < n; i++ {
		for j := 0; j < l; j++ {
			res += strconv.Itoa(c[i][j])
			if j != l-1 {
				res += " "
			}
		}
		if i != n-1 {
			res += "\n"
		}
	}

	return res, nil
}
