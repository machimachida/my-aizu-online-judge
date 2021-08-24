package itp1

import (
	"errors"
	"math"
	"sort"
	"strconv"
	"strings"
)

func ABProblem(str string) (string, error) {
	ss := strings.Split(str, " ")
	if len(ss) != 2 {
		return "", errors.New("The number of datasets > 3000")
	}

	is := make([]int, 2)
	for i, s := range ss {
		x, err := strconv.Atoi(s)
		if err != nil {
			return "", err
		}
		if x < 1 || x > 1000000000 {
			var v string
			if i == 0 {
				v = "a"
			} else {
				v = "b"
			}
			return "", errors.New(v + " is not match 1 <= " + v + " <= 10^9. " + v + " = " + s)
		}
		is[i] = x
	}

	return strconv.Itoa(is[0]/is[1]) + " " + strconv.Itoa(is[0]/is[1]) + " " + strconv.FormatFloat(float64(is[0])/float64(is[1]), 'f', 5, 64), nil
}

func Circle(str string) (string, error) {
	r, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return "", err
	}
	if r <= 0 || r >= 10000 {
		return "", errors.New("r is not match 0 < r < 10000. r = " + str)
	}

	return strconv.FormatFloat(r*r*math.Pi, 'f', 6, 64) + " " + strconv.FormatFloat(2*r*math.Pi, 'f', 6, 64), nil
}

func SimpleCalculator(str string) (string, error) {
	ss := strings.Split(str, "\n")
	res := ""
	for _, s := range ss {
		exp := strings.Split(s, " ")
		if len(exp) != 3 {
			return "", errors.New("invalid input")
		}

		if exp[1] == "?" {
			if len(res) > 0 {
				res = res[:len(res)-1]
			}
			break
		}

		is := make([]int, 2)
		for i := 0; i < 3; i = i + 2 {
			n, err := strconv.Atoi(exp[i])
			if err != nil {
				return "", err
			}
			if n < 0 || n > 20000 {
				var v string
				if i == 0 {
					v = "a"
				} else {
					v = "b"
				}
				return "", errors.New(v + " is not match 0 <= " + v + " <= 20000. " + v + " = " + exp[i])
			}
			is[i/2] = n
		}

		switch exp[1] {
		case "+":
			res += strconv.Itoa(is[0]+is[1]) + "\n"
		case "-":
			res += strconv.Itoa(is[0]-is[1]) + "\n"
		case "*":
			res += strconv.Itoa(is[0]*is[1]) + "\n"
		case "/":
			if is[1] == 0 {
				return "", errors.New("A division by zero are given.")
			}
			res += strconv.Itoa(is[0]/is[1]) + "\n"
		}
	}
	return res, nil
}

func MinMaxAndSum(str string) (string, error) {
	ss := strings.Split(str, "\n")
	if len(ss) != 2 {
		return "", errors.New("invalid lines")
	}

	n, err := strconv.Atoi(ss[0])
	if err != nil {
		return "", err
	}
	if n <= 0 || n > 10000 {
		return "", errors.New("n is not match 0 < n <= 10000. n = " + ss[0])
	}

	ns := strings.Split(ss[1], " ")
	if len(ns) != n {
		return "", errors.New("invalid input")
	}

	is := make([]int, n)
	sum := 0
	for i, s := range ns {
		a, err := strconv.Atoi(s)
		if err != nil {
			return "", err
		}
		if a < -1000000 || a > 1000000 {
			return "", errors.New("a is not match -1000000 <= a <= 1000000. a = " + s)
		}
		is[i] = a
		sum += a
	}

	sort.Ints(is)
	return strconv.Itoa(is[0]) + " " + strconv.Itoa(is[n-1]) + " " + strconv.Itoa(sum), nil
}
