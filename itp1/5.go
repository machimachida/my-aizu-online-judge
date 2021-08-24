package itp1

import (
	"errors"
	"strconv"
	"strings"
)

func PrintARectangle(str string) (string, error) {
	ss := strings.Split(str, "\n")
	res := ""
	for _, s := range ss {
		if s == "0 0" {
			break
		}
		els := strings.Split(s, " ")
		if len(els) != 2 {
			return "", errors.New("invalid input")
		}

		is := make([]int, 2)
		for i, el := range els {
			n, err := strconv.Atoi(el)
			if err != nil {
				return "", err
			}
			if n < 1 || n > 300 {
				var v string
				if i == 0 {
					v = "H"
				} else {
					v = "W"
				}
				return "", errors.New(v + " is not match 1 <= " + v + " <= 300. " + v + " = " + el)
			}
			is[i] = n
		}

		for i := 0; i < is[0]; i++ {
			for j := 0; j < is[1]; j++ {
				res += "#"
			}
			res += "\n"
		}
		res += "\n"
	}
	return res, nil
}

func PrintAFrame(str string) (string, error) {
	ss := strings.Split(str, "\n")
	res := ""
	for _, s := range ss {
		if s == "0 0" {
			break
		}
		els := strings.Split(s, " ")
		if len(els) != 2 {
			return "", errors.New("invalid input")
		}

		is := make([]int, 2)
		for i, el := range els {
			n, err := strconv.Atoi(el)
			if err != nil {
				return "", err
			}
			if n < 3 || n > 300 {
				var v string
				if i == 0 {
					v = "H"
				} else {
					v = "W"
				}
				return "", errors.New(v + " is not match 3 <= " + v + " <= 300. " + v + " = " + el)
			}
			is[i] = n
		}

		H := is[0] - 1
		for i := 0; i < is[0]; i++ {
			W := is[1] - 1
			for j := 0; j < is[1]; j++ {
				if i == 0 || i == H || j == 0 || j == W {
					res += "#"
				} else {
					res += "."
				}
			}
			res += "\n"
		}
		res += "\n"
	}
	return res, nil
}

func PrintAChessboard(str string) (string, error) {
	ss := strings.Split(str, "\n")
	res := ""
	for _, s := range ss {
		if s == "0 0" {
			break
		}
		els := strings.Split(s, " ")
		if len(els) != 2 {
			return "", errors.New("invalid input")
		}

		is := make([]int, 2)
		for i, el := range els {
			n, err := strconv.Atoi(el)
			if err != nil {
				return "", err
			}
			if n < 1 || n > 300 {
				var v string
				if i == 0 {
					v = "H"
				} else {
					v = "W"
				}
				return "", errors.New(v + " is not match 1 <= " + v + " <= 300. " + v + " = " + el)
			}
			is[i] = n
		}

		for i := 0; i < is[0]; i++ {
			isHash := i%2 == 0
			for j := 0; j < is[1]; j++ {
				if isHash {
					res += "#"
				} else {
					res += "."
				}
				isHash = !isHash
			}
			res += "\n"
		}
		res += "\n"
	}
	return res, nil
}

func StructuredProgramming(str string) (string, error) {
	n, err := strconv.Atoi(str)
	if err != nil {
		return "", err
	}
	if n < 3 || n > 10000 {
		return "", errors.New("n is not match 1 <= n <= 300. n = " + str)
	}

	res := ""
	for i := 1; i <= n; i++ {
		x := i
		if x%3 == 0 {
			res += " " + strconv.Itoa(i)
		} else {
			for x != 0 {
				if x%10 == 3 {
					res += " " + strconv.Itoa(i)
					break
				}
				x = x / 10
			}
		}
	}
	return res, nil
}
