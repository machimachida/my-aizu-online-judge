package itp1

import (
	"errors"
	"sort"
	"strconv"
	"strings"
)

func SmallLargeEqual(s string) (string, error) {
	ss := strings.Split(s, " ")
	if len(ss) != 2 {
		return "", errors.New("invalid input")
	}
	is := make([]int, 2)
	for i, s := range ss {
		n, err := strconv.Atoi(s)
		if err != nil {
			return "", err
		}
		if n < -1000 || n > 1000 {
			var v string
			if i == 0 {
				v = "a"
			} else {
				v = "b"
			}
			return "", errors.New(v + " is not match -1000 <= " + v + " <= 1000. " + v + " = " + strconv.Itoa(n))
		}
		is[i] = n
	}

	if is[0] < is[1] {
		return "a < b", nil
	} else if is[0] > is[1] {
		return "a > b", nil
	} else {
		return "a == b", nil
	}
}

func Range(s string) (string, error) {
	ss := strings.Split(s, " ")
	if len(ss) != 3 {
		return "", errors.New("invalid input")
	}
	is := make([]int, 3)
	for i, s := range ss {
		n, err := strconv.Atoi(s)
		if err != nil {
			return "", err
		}
		if n < 0 || n > 100 {
			var v string
			if i == 0 {
				v = "a"
			} else if i == 1 {
				v = "b"
			} else {
				v = "c"
			}
			return "", errors.New(v + " is not match 0 <= " + v + " <= 100. " + v + " = " + strconv.Itoa(n))
		}
		is[i] = n
	}

	if is[0] < is[1] && is[1] < is[2] {
		return "Yes", nil
	} else {
		return "No", nil
	}
}

func SortingThreeNumbers(s string) (string, error) {
	ss := strings.Split(s, " ")
	if len(ss) != 3 {
		return "", errors.New("invalid input")
	}
	is := make([]int, 3)
	for i, s := range ss {
		n, err := strconv.Atoi(s)
		if err != nil {
			return "", err
		}
		if n < 0 || n > 100 {
			var v string
			if i == 0 {
				v = "a"
			} else if i == 1 {
				v = "b"
			} else {
				v = "c"
			}
			return "", errors.New(v + " is not match 0 <= " + v + " <= 100. " + v + " = " + strconv.Itoa(n))
		}
		is[i] = n
	}
	sort.Ints(is)
	return strconv.Itoa(is[0]) + " " + strconv.Itoa(is[1]) + " " + strconv.Itoa(is[2]), nil
}

func CircleInARectangle(s string) (string, error) {
	ss := strings.Split(s, " ")
	if len(ss) != 5 {
		return "", errors.New("invalid input")
	}
	is := make([]int, 5)
	for i, s := range ss {
		n, err := strconv.Atoi(s)
		if err != nil {
			return "", err
		}
		if i == 1 || i == 2 {
			if n < -100 || n > 100 {
				var v string
				if i == 1 {
					v = "x"
				} else {
					v = "y"
				}
				return "", errors.New(v + " is not match -100 <= " + v + " <= 100. " + v + " = " + strconv.Itoa(n))
			}
		} else {
			if n <= 0 || n > 100 {
				var v string
				if i == 0 {
					v = "W"
				} else if i == 1 {
					v = "H"
				} else {
					v = "r"
				}
				return "", errors.New(v + " is not match 0 < " + v + " <= 100. " + v + " = " + strconv.Itoa(n))
			}
		}
		is[i] = n
	}

	if is[2]-is[4] < 0 ||
		is[3]-is[4] < 0 ||
		is[2]+is[4] > is[0] ||
		is[3]+is[4] > is[1] {
		return "No", nil
	} else {
		return "Yes", nil
	}
}
