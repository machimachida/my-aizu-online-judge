package itp1

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func PrintManyHelloWorld() {
	for i := 0; i < 1000; i++ {
		fmt.Println("Hello World")
	}
}

func PrintTestCases(str string) (string, error) {
	ss := strings.Split(str, "\n")
	l := len(ss)
	if l > 1000 {
		return "", errors.New("The number of datasets > 1000")
	}

	res := ""
	for i, s := range ss {
		if s == "0" {
			res = res[:len(res)-1]
			break
		}

		x, err := strconv.Atoi(s)
		if err != nil {
			return "", err
		}

		if x < 1 || x > 10000 {
			return "", errors.New("x is not match 1 <= x <= 10000. x = " + strconv.Itoa(x))
		}
		res += "Case " + strconv.Itoa(i+1) + ": " + s + "\n"
	}

	return res, nil
}

func SwappingTwoNumbers(str string) (string, error) {
	ss := strings.Split(str, "\n")
	if len(ss) > 3000 {
		return "", errors.New("The number of datasets > 3000")
	}

	res := ""
	for _, s := range ss {
		if s == "0 0" {
			res = res[:len(res)-1]
			break
		}

		ns := strings.Split(s, " ")
		if len(ns) != 2 {
			return "", errors.New("invalid input")
		}

		is := make([]int, 2)
		for i, n := range ns {
			x, err := strconv.Atoi(n)
			if err != nil {
				return "", err
			}
			if x < 0 || x > 10000 {
				var v string
				if i == 0 {
					v = "x"
				} else {
					v = "y"
				}
				return "", errors.New(v + " is not match 0 <= " + v + " <= 10000. " + v + " = " + strconv.Itoa(x))
			}
			is[i] = x
		}

		if is[0] > is[1] {
			ns[0], ns[1] = ns[1], ns[0]
		}
		res += ns[0] + " " + ns[1] + "\n"
	}

	return res, nil
}

func HowManyDivisors(str string) (string, error) {
	ss := strings.Split(str, " ")
	if len(ss) != 3 {
		return "", errors.New("The number of datasets > 3000")
	}

	is := make([]int, 3)
	for i, s := range ss {
		x, err := strconv.Atoi(s)
		if err != nil {
			return "", err
		}
		if x < 1 || x > 10000 {
			var v string
			if i == 0 {
				v = "a"
			} else if i == 1 {
				v = "b"
			} else {
				v = "c"
			}
			return "", errors.New(v + " is not match 1 <= " + v + " <= 10000. " + v + " = " + s)
		}
		is[i] = x
	}

	if is[0] > is[1] {
		return "", errors.New("Not match a <= b. a = " + ss[0] + " b = " + ss[1])
	}

	c := 0
	for i := is[0]; i <= is[1]; i++ {
		if is[2]%i == 0 {
			c++
		}
	}

	return strconv.Itoa(c), nil
}
