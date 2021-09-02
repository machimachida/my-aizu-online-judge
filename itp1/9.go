package itp1

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func FindingAWord() {
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		fmt.Println("invalid input")
		return
	}
	W := scanner.Text()
	if len(W) > 10 {
		fmt.Println("invalid length of W")
		return
	}
	for _, c := range W {
		if c >= 'A' && c <= 'Z' {
			fmt.Println("invalid W")
			return
		}
	}

	var count int
	for scanner.Scan() {
		T := scanner.Text()
		if T == "END_OF_TEXT" {
			break
		}
		if len(T) >= 1000 {
			fmt.Println("invalid length of the input string")
			return
		}

		T = strings.ToLower(T)
		var i int
		var scan int
		for {
			i = strings.Index(T[scan:], W)
			if i == -1 {
				break
			}
			i += scan

			if (i == 0 || T[i-1] == ' ') && (i+len(W) == len(T) || T[i+len(W)] == ' ') {
				count++
			}
			scan = i + len(W)
		}
	}
	fmt.Println(count)
}

func Shuffle() {
	var str string
	var m int
	var h int
	for i := 0; i < 10; i++ {
		_, err := fmt.Scan(&str)
		if err != nil {
			fmt.Println(err)
			return
		}
		if str == "-" {
			return
		}
		l := len(str)
		if l > 200 {
			fmt.Println("invalid length of input string")
			return
		}

		_, err = fmt.Scan(&m)
		if err != nil {
			fmt.Println(err)
			return
		}
		if m < 1 || m > 100 {
			fmt.Println("invalid m range")
			return
		}

		for j := 0; j < m; j++ {
			_, err = fmt.Scan(&h)
			if err != nil {
				fmt.Println(err)
				return
			}
			if h < 1 || h >= l {
				fmt.Println("invalid h range")
				return
			}
			str = str[h:] + str[:h]
		}
		fmt.Println(str)
	}
}

func CardGame() {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println(err)
		return
	}
	if n < 1 || n > 1000 {
		fmt.Println("invalid n range")
	}

	ps := make([]int, 2)
	strs := make([]string, 2)
	for i := 0; i < n; i++ {
		_, err := fmt.Scan(&strs[0], &strs[1])
		if err != nil {
			fmt.Println(err)
			return
		}
		for i := range strs {
			if len(strs[i]) > 100 {
				fmt.Println("invalid string length")
				return
			}
			for _, c := range strs[i] {
				if c >= 'A' && c <= 'Z' {
					fmt.Println("invalid string letter")
					return
				}
			}
		}
		if strs[0] == strs[1] {
			ps[0]++
			ps[1]++
			continue
		}

		sp := &ps[0]
		lp := &ps[1]
		if len(strs[0]) > len(strs[1]) {
			strs[0], strs[1] = strs[1], strs[0]
			sp, lp = lp, sp
		}
		l := len(strs[0])
		for i, c := range strs[0] {
			r := rune(strs[1][i])
			if c == r {
				if i == l-1 {
					*lp += 3
				}
				continue
			} else if c > r {
				*sp += 3
				break
			} else if c < rune(strs[1][i]) {
				*lp += 3
				break
			}
		}
	}
	fmt.Println(ps[0], ps[1])
}

func Transformation() {
	var str string
	_, err := fmt.Scan(&str)
	if err != nil {
		fmt.Println(err)
		return
	}
	l := len(str)
	if l < 1 || l > 1000 {
		fmt.Println("invalid str length")
		return
	}

	var q int
	_, err = fmt.Scan(&q)
	if err != nil {
		fmt.Println(err)
		return
	}
	if q < 1 || q > 1000 {
		fmt.Println("invalid q range")
		return
	}

	var op string
	var a int
	var b int
	for i := 0; i < q; i++ {
		_, err = fmt.Scan(&op, &a, &b)
		if err != nil {
			fmt.Println(err)
			return
		}
		if a < 0 || a > b {
			fmt.Println("invalid a range")
			return
		}
		if b >= l {
			fmt.Println("invalid b range")
			return
		}

		switch op {
		case "replace":
			var r string
			_, err := fmt.Scan(&r)
			if err != nil {
				fmt.Println(err)
				return
			}
			if len(r) != b-a+1 {
				fmt.Println("invalid r length")
				return
			}
			str = str[:a] + r + str[b+1:]
		case "reverse":
			var rv string
			for _, c := range str[a : b+1] {
				rv = string(c) + rv
			}
			str = str[:a] + rv + str[b+1:]
		case "print":
			fmt.Println(str[a : b+1])
		default:
			fmt.Println("invalid operation")
			return
		}
	}
}
