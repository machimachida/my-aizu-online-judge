package itp1

import (
	"errors"
	"fmt"
)

type dice struct {
	nums []int
}

func newDice() (dice, error) {
	var d dice
	d.nums = make([]int, 6)
	_, err := fmt.Scan(&d.nums[0], &d.nums[1], &d.nums[2], &d.nums[3], &d.nums[4], &d.nums[5])
	if err != nil {
		return d, err
	}
	for _, n := range d.nums {
		if n < 0 || n > 100 {
			return d, errors.New("invald dice number range")
		}
	}
	return d, nil
}

func (d dice) roll(r rune) {
	switch r {
	case 'E':
		d.nums[0], d.nums[3], d.nums[5], d.nums[2] = d.nums[3], d.nums[5], d.nums[2], d.nums[0]
	case 'N':
		d.nums[0], d.nums[1], d.nums[5], d.nums[4] = d.nums[1], d.nums[5], d.nums[4], d.nums[0]
	case 'S':
		d.nums[0], d.nums[1], d.nums[5], d.nums[4] = d.nums[4], d.nums[0], d.nums[1], d.nums[5]
	case 'W':
		d.nums[0], d.nums[3], d.nums[5], d.nums[2] = d.nums[2], d.nums[0], d.nums[3], d.nums[5]
	}
}

func (d dice) isUnique() bool {
	m := make(map[int]bool, 6)
	for _, n := range d.nums {
		if m[n] {
			return false
		} else {
			m[n] = true
		}
	}
	return true
}

func (d1 dice) equal(d2 dice) bool {
	rs := []rune{' ', 'N', 'W', 'W', 'W', 'N'} // 0, 1, 2, 4, 3, 5
	for _, r := range rs {
		d2.roll(r)
		if d1.nums[0] != d2.nums[0] || d1.nums[5] != d2.nums[5] {
			continue
		}
		for i := 0; i < 4; i++ {
			d2.roll('W')
			d2.roll('N')
			d2.roll('E')
			if d1.nums[1] == d2.nums[1] &&
				d1.nums[2] == d2.nums[2] &&
				d1.nums[3] == d2.nums[3] &&
				d1.nums[4] == d2.nums[4] {
				return true
			}
		}
	}
	return false
}

func Dice1() {
	d, err := newDice()
	if err != nil {
		fmt.Println(err)
		return
	}

	var op string
	_, err = fmt.Scan(&op)
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(op) > 100 {
		fmt.Println("invalid operation length")
		return
	}
	for _, r := range op {
		d.roll(r)
	}
	fmt.Println(d.nums[0])
}

func Dice2() {
	d, err := newDice()
	if err != nil {
		fmt.Println(err)
		return
	}

	if !d.isUnique() {
		fmt.Println("dice nums are not unique")
		return
	}
	var q int
	_, err = fmt.Scan(&q)
	if err != nil {
		fmt.Println(err)
		return
	}
	if q < 1 || q > 24 {
		fmt.Println("invalid q range")
		return
	}

	var t int
	var f int
	for i := 0; i < q; i++ {
		_, err = fmt.Scan(&t, &f)
		if err != nil {
			fmt.Println(err)
			return
		}
		if t == d.nums[1] {
			d.roll('N')
		} else if t == d.nums[2] {
			d.roll('W')
		} else if t == d.nums[3] {
			d.roll('E')
		} else if t == d.nums[4] {
			d.roll('S')
		} else if t == d.nums[5] {
			d.roll('N')
			d.roll('N')
		} else if t != d.nums[0] {
			fmt.Println("invalid t")
		}

		if f == d.nums[2] {
			d.roll('E')
			d.roll('N')
			d.roll('W')
		} else if f == d.nums[3] {
			d.roll('W')
			d.roll('N')
			d.roll('E')
		} else if f == d.nums[4] {
			d.roll('E')
			d.roll('E')
			d.roll('N')
			d.roll('N')
		} else if f != d.nums[1] {
			fmt.Println("invalid f")
		}
		fmt.Println(d.nums[2])
	}
}

func Dice3() {
	d1, err := newDice()
	if err != nil {
		fmt.Println(err)
		return
	}
	d2, err := newDice()
	if err != nil {
		fmt.Println(err)
		return
	}

	if d1.equal(d2) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func Dice4() {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println(err)
		return
	}
	if n < 2 || n > 100 {
		fmt.Println("invalid n range")
		return
	}

	ds := make([]dice, n)
	for i := 0; i < n; i++ {
		ds[i], err = newDice()
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if ds[i].equal(ds[j]) {
				fmt.Println("No")
				return
			}
		}
	}
	fmt.Println("Yes")
}
