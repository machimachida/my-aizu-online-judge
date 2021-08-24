package itp1

import (
	"errors"
	"strconv"
	"strings"
)

func ReversingNumbers(str string) (string, error) {
	ss := strings.Split(str, "\n")
	n, err := strconv.Atoi(ss[0])
	if err != nil {
		return "", err
	}
	if n > 100 {
		return "", errors.New("n is not match n <= 100. n = " + ss[0])
	}

	els := strings.Split(ss[1], " ")
	if len(els) != n {
		return "", errors.New("invalid input")
	}

	is := make([]int, n)
	for i, el := range els {
		a, err := strconv.Atoi(el)
		if err != nil {
			return "", err
		}
		if a < 0 || a >= 1000 {
			return "", errors.New("a is not match 0 <= a < 1000. a = " + el)
		}
		is[i] = a
	}

	res := ""
	for i := n - 1; i >= 0; i-- {
		res += strconv.Itoa(is[i]) + " "
	}
	return res[:len(res)-1], nil
}

func FindingMissingCards(str string) (string, error) {
	ss := strings.Split(str, "\n")
	n, err := strconv.Atoi(ss[0])
	if err != nil {
		return "", err
	}

	if len(ss)-1 != n {
		return "", errors.New("invalid n. lines = " + strconv.Itoa(len(ss)) + " n = " + strconv.Itoa(n))
	}

	exs := make([][]bool, 4)
	for i := 0; i < 4; i++ {
		exs[i] = make([]bool, 13)
	}

	for _, s := range ss[1:] {
		els := strings.Split(s, " ")
		if len(els) != 2 {
			return "", errors.New("invalid input")
		}

		a, err := strconv.Atoi(els[1])
		if err != nil {
			return "", err
		}
		if a < 1 || a > 13 {
			return "", errors.New("invalid rank")
		}

		switch els[0] {
		case "S":
			exs[0][a-1] = true
		case "H":
			exs[1][a-1] = true
		case "C":
			exs[2][a-1] = true
		case "D":
			exs[3][a-1] = true
		default:
			return "", errors.New("invalid suit")
		}
	}
	res := ""
	for i, ex := range exs {
		var suit string
		switch i {
		case 0:
			suit = "S "
		case 1:
			suit = "H "
		case 2:
			suit = "C "
		case 3:
			suit = "D "
		}
		for j, isExist := range ex {
			if !isExist {
				res += suit + strconv.Itoa(j+1) + "\n"
			}
		}
	}
	return res, nil
}

func OfficialHouse(str string) (string, error) {
	ss := strings.Split(str, "\n")
	n, err := strconv.Atoi(ss[0])
	if err != nil {
		return "", err
	}

	if len(ss)-1 != n {
		return "", errors.New("invalid n. lines = " + strconv.Itoa(len(ss)) + " n = " + strconv.Itoa(n))
	}

	oh := make([][]int, 12)
	for i := 0; i < 12; i++ {
		oh[i] = make([]int, 10)
	}

	for _, s := range ss[1:] {
		els := strings.Split(s, " ")
		if len(els) != 4 {
			return "", errors.New("invalid line")
		}

		b, err := strconv.Atoi(els[0])
		if err != nil {
			return "", err
		}
		if b < 1 || b > 4 {
			return "", errors.New("incorrect building number: " + els[0])
		}

		f, err := strconv.Atoi(els[1])
		if err != nil {
			return "", err
		}
		if f < 1 || f > 3 {
			return "", errors.New("incorrect floor number" + els[1])
		}

		r, err := strconv.Atoi(els[2])
		if err != nil {
			return "", err
		}
		if r < 1 || r > 10 {
			return "", errors.New("incorrect room number" + els[2])
		}

		v, err := strconv.Atoi(els[3])
		if err != nil {
			return "", err
		}
		if v < 0 || v > 9 {
			return "", errors.New("the number of tenants is not match 0 <= v <= 9. v = " + els[2])
		}

		oh[(b-1)*3+(f-1)][r-1] += v
	}

	res := ""
	l := len(oh)
	for i, f := range oh {
		for _, r := range f {
			res += " " + strconv.Itoa(r)
		}
		res += "\n"
		if i%3 == 2 && i != l-1 {
			res += "####################\n"
		}
	}

	return res[:len(res)-1], nil
}

func MatrixVectorMultiplication(str string) (string, error) {
	ss := strings.Split(str, "\n")
	nms := strings.Split(ss[0], " ")
	if len(nms) != 2 {
		return "", errors.New("invalid number of integer in first line")
	}

	nm := make([]int, 2)
	for i, s := range nms {
		n, err := strconv.Atoi(s)
		if err != nil {
			return "", err
		}
		if n < 1 || n > 100 {
			return "", errors.New("invalid n or m")
		}
		nm[i] = n
	}

	if len(ss) != nm[0]+nm[1]+1 {
		return "", errors.New("invalid number of lines")
	}

	a := make([][]int, nm[0])
	for i := 0; i < nm[0]; i++ {
		a[i] = make([]int, nm[1])
		ns := strings.Split(ss[i+1], " ")
		if len(ns) != nm[1] {
			return "", errors.New("invalid number of integer")
		}
		for j, s := range ns {
			n, err := strconv.Atoi(s)
			if err != nil {
				return "", err
			}
			if n < 0 || n > 1000 {
				return "", errors.New("invalid a")
			}

			a[i][j] = n
		}
	}
	b := make([]int, nm[1])
	for i, s := range ss[nm[0]+1:] {
		n, err := strconv.Atoi(s)
		if err != nil {
			return "", err
		}
		if n < 0 || n > 1000 {
			return "", errors.New("invalid b")
		}

		b[i] = n
	}

	res := ""
	for i := 0; i < nm[0]; i++ {
		c := 0
		for j := 0; j < nm[1]; j++ {
			c += a[i][j] * b[j]
		}
		res += strconv.Itoa(c) + "\n"
	}
	return res[:len(res)-1], nil
}
