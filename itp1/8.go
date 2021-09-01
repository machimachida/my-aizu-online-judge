package itp1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func TogglingCases() { // TODO AOJの動作環境が古いため、strings.Builderが使えない
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		fmt.Println("invalid input")
		return
	}
	str := scanner.Text()
	if len(str) >= 1200 {
		fmt.Println("invalid length of the input string")
		return
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
	fmt.Println(b.String())
}

func SumOfNumbers() {
	var line string
	for i := 0; i < 1000; i++ {
		_, err := fmt.Scan(&line)
		if err != nil {
			fmt.Println(err)
			return
		}
		if line == "0" {
			break
		}

		var sum int
		for i := range line {
			n, err := strconv.Atoi(string(line[i]))
			if err != nil {
				fmt.Println(err)
				return
			}
			sum += n
		}
		fmt.Println(sum)
	}
}

func CountingCharacters() {
	scanner := bufio.NewScanner(os.Stdin)
	n := make(map[rune]int, 26)
	for scanner.Scan() {
		str := scanner.Text()
		if len(str) >= 1200 {
			fmt.Println("invalid length of the input string")
			return
		}

		for _, c := range str {
			if c >= 'A' && c <= 'Z' {
				c += 'a' - 'A'
			}
			n[c]++
		}
	}
	for c := 'a'; c <= 'z'; c++ {
		fmt.Println(string(c), ":", n[c])
	}
}

func Ring() {
	sp := make([]string, 2)
	_, err := fmt.Scan(&sp[0], &sp[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	sl := len(sp[0])
	pl := len(sp[1])
	if pl == 0 || pl > sl || pl > 100 || sl > 100 {
		fmt.Println("invalid p or s length")
		return
	}

	for i := range sp[0] {
		for j, c := range sp[1] {
			index := i + j
			if i+j >= sl {
				index -= sl
			}
			if c != rune(sp[0][index]) {
				break
			}
			if j == pl-1 {
				fmt.Println("Yes")
				return
			}
		}
	}
	fmt.Println("No")
}
