package itp1

import (
	"errors"
	"strconv"
	"strings"
)

func TogglingCases(str string) (string, error) {
	if len(str) >= 1200 {
		return "", errors.New("invalid length of the input string")
	}

	var b strings.Builder
	b.Grow(len(str))
	for _, c := range str {
		if c >= 'A' && c <= 'Z' {
			c += 'a' - 'A'
		} else if c >= 'a' && c <= 'z' {
			c -= 'a' - 'A'
		}
		b.WriteRune(c)
	}

	return b.String(), nil
}

func SumOfNumbers(str string) (string, error) {
	lines := strings.Split(str, "\n")
	res := ""
	for _, line := range lines {
		if line == "0" {
			break
		}

		sum := 0
		for _, c := range line {
			n, err := strconv.Atoi(string(c))
			if err != nil {
				return "", err
			}
			sum += n
			if sum > 1000 {
				return "", errors.New("the nunber of digits in each line number does not exceed 1000")
			}
		}
		res += strconv.Itoa(sum) + "\n"
	}
	return res[:len(res)-1], nil
}

func CountingCharacters(str string) (string, error) {
	if len(str) > 1200 {
		return "", errors.New("invalid input length")
	}
	n := make(map[rune]int, 26)
	for _, c := range str {
		if c >= 'A' && c <= 'Z' {
			c += 'a' - 'A'
		}
		n[c]++
	}
	var res string
	for c := 'a'; c <= 'z'; c++ {
		res += string(c) + " : " + strconv.Itoa(n[c]) + "\n"
	}
	return res[:len(res)-1], nil
}

func Ring(str string) (string, error) {
	inputs := strings.Split(str, "\n")
	if len(inputs) != 2 {
		return "", errors.New("invalid input")
	}

	s := inputs[0]
	sl := len(s)
	p := inputs[1]
	pl := len(p)
	if pl == 0 || p > s || pl > 100 || sl > 100 {
		return "", errors.New("invalid p or s length")
	}

	for i := range s {
		for j, c := range p {
			index := i + j
			if i+j >= sl {
				index -= sl
			}
			if c != rune(s[index]) {
				break
			}
			if j == pl-1 {
				return "Yes", nil
			}
		}
	}
	return "No", nil
}
