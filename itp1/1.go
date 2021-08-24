package itp1

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func HelloWorld() {
	fmt.Println("Hello World")
}

func XCubic(s string) (string, error) {
	x, err := strconv.Atoi(s)
	if err != nil {
		return "", errors.New("invalid input")
	}
	if x < 1 || x > 100 {
		return "", errors.New("x is not match 1 <= x <= 100")
	}
	return strconv.Itoa(x * x * x), nil
}

func Rectangle(s string) (string, error) {
	ss := strings.Split(s, " ")
	if len(ss) != 2 {
		return "", errors.New("invalid input")
	}
	is := make([]int, 2)
	for i, v := range ss {
		n, err := strconv.Atoi(v)
		if err != nil {
			return "", err
		}
		if n < 1 || n > 100 {
			return "", errors.New("value is not match 1 <= value <= 100. value = " + strconv.Itoa(n))
		}
		is[i] = n
	}
	return strconv.Itoa(is[0]*is[1]) + " " + strconv.Itoa(is[0]*2+is[1]*2), nil
}

func Watch(s string) (string, error) {
	n, err := strconv.Atoi(s)
	if err != nil {
		return "", errors.New("invalid input")
	}
	if n < 0 || n > 86400 {
		return "", errors.New("value is not match 0 <= value <= 86400. value = " + strconv.Itoa(n))
	}
	se := n % 60
	n = n / 60
	mi := n % 60
	ho := n / 60
	return strconv.Itoa(ho) + ":" + strconv.Itoa(mi) + ":" + strconv.Itoa(se), nil
}
