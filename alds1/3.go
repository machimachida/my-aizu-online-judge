package alds1

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Stack() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	inp := strings.Split(sc.Text(), " ")
	if len(inp) < 3 {
		fmt.Println("invalid number of values")
		return
	}
	var stack []int
	var idx int
	var numOprd int
	var numOprt int

	push := func(s string) error {
		i, err := strconv.Atoi(s)
		if err != nil {
			return err
		}
		if i < 1 || i > 1000000 {
			return errors.New("invalid operand")
		}
		if idx < len(stack) {
			stack[idx] = i
		} else {
			stack = append(stack, i)
		}
		idx++
		numOprd++
		return nil
	}

	for _, s := range inp {
		if idx > 1 {
			switch s {
			case "+":
				stack[idx-2] = stack[idx-2] + stack[idx-1]
				numOprt++
				idx--
			case "-":
				stack[idx-2] = stack[idx-2] - stack[idx-1]
				numOprt++
				idx--
			case "*":
				stack[idx-2] = stack[idx-2] * stack[idx-1]
				numOprt++
				idx--
			default:
				err := push(s)
				if err != nil {
					fmt.Println(err)
					return
				}
			}
		} else {
			err := push(s)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
	if idx != 1 || len(stack) == 0 {
		fmt.Println("invalid answer")
		return
	}
	if numOprd < 2 || numOprd > 100 {
		fmt.Println("invalid number of operands")
		return
	}
	if numOprt < 1 || numOprd > 99 {
		fmt.Println("invalid number of operators")
		return
	}
	fmt.Println(stack[0])
}

/* this function makes many allocs but is as fast as Stack() */
func stackStack() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	inp := strings.Split(sc.Text(), " ")
	if len(inp) < 3 {
		fmt.Println("invalid number of values")
		return
	}
	type (
		node struct {
			value int
			prev  *node
		}
		stack struct {
			top    *node
			length int
		}
	)

	// View the top item on the stack
	peek := func(this *stack) int {
		if this.length == 0 {
			return 0
		}
		return this.top.value
	}
	// Pop the top item of the stack and return it
	pop := func(this *stack) int {
		if this.length == 0 {
			return 0
		}

		n := this.top
		this.top = n.prev
		this.length--
		return n.value
	}
	// Push a value onto the top of the stack
	push := func(this *stack, value int) {
		n := &node{value, this.top}
		this.top = n
		this.length++
	}

	var stk stack
	for _, s := range inp {
		switch s {
		case "+":
			push(&stk, pop(&stk)+pop(&stk))
		case "-":
			push(&stk, pop(&stk)-pop(&stk))
		case "*":
			push(&stk, pop(&stk)*pop(&stk))
		default:
			i, err := strconv.Atoi(s)
			if err != nil {
				fmt.Println(err)
				return
			}
			push(&stk, i)
		}
	}
	fmt.Println(peek(&stk))
}

func Queue() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	nq := strings.Split(sc.Text(), " ")
	if len(nq) != 2 {
		fmt.Println("invalid number of first values")
		return
	}
	n, err := strconv.Atoi(nq[0])
	if err != nil {
		fmt.Println(err)
		return
	}
	if n < 1 || n > 100000 {
		fmt.Println("invalid n range")
		return
	}
	q, err := strconv.Atoi(nq[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	if q < 1 || q > 1000 {
		fmt.Println("invalid q range")
		return
	}

	type schedule struct {
		name string
		time int
	}
	schds := make([]schedule, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		nt := strings.Split(sc.Text(), " ")
		if len(nt) != 2 {
			fmt.Println("invalid number of first values")
			return
		}
		if len(nt[0]) == 0 || len(nt[0]) > 10 {
			fmt.Println("invalid name length")
			return
		}
		schds[i].name = nt[0]
		schds[i].time, err = strconv.Atoi(nt[1])
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	// round-robin scheduling
	var now int
	for len(schds) != 0 {
		if schds[0].time <= q {
			now += schds[0].time
			fmt.Println(schds[0].name, now)
			schds = schds[1:]
		} else {
			now += q
			schds[0].time -= q
			schds = append(schds[1:], schds[0])
		}
	}
}

// huge inputs make this function slow
func QueueMap() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	nq := strings.Split(sc.Text(), " ")
	if len(nq) != 2 {
		fmt.Println("invalid number of first values")
		return
	}
	n, err := strconv.Atoi(nq[0])
	if err != nil {
		fmt.Println(err)
		return
	}
	if n < 1 || n > 100000 {
		fmt.Println("invalid n range")
		return
	}
	q, err := strconv.Atoi(nq[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	if q < 1 || q > 1000 {
		fmt.Println("invalid q range")
		return
	}

	schds := make(map[string]int, n)
	odr := make([]string, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		nt := strings.Split(sc.Text(), " ")
		if len(nt) != 2 {
			fmt.Println("invalid number of first values")
			return
		}
		if len(nt[0]) == 0 || len(nt[0]) > 10 {
			fmt.Println("invalid name length")
			return
		}
		schds[nt[0]], err = strconv.Atoi(nt[1])
		if err != nil {
			fmt.Println(err)
			return
		}
		odr[i] = nt[0]
	}
	// round-robin scheduling
	var now int
	var cnt int
	for i := 0; i < n; i++ {
		if schds[odr[i]] == 0 {
			if i == n-1 {
				i = -1
			}
			continue
		}
		if schds[odr[i]] <= q {
			now += schds[odr[i]]
			schds[odr[i]] = 0
			fmt.Println(odr[i], now)
			if cnt == n-1 {
				break
			} else {
				cnt++
			}
		} else {
			now += q
			schds[odr[i]] -= q
		}
		if i == n-1 {
			i = -1
		}
	}
}

func DoublyLinkedList() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	n, err := strconv.Atoi(sc.Text())
	if err != nil {
		fmt.Println(err)
		return
	}
	if n < 0 || n > 2000000 {
		fmt.Println("invalid number of operation")
		return
	}

	type doublyLinkedListNode struct {
		key  int
		next *doublyLinkedListNode
		prev *doublyLinkedListNode
	}

	bottom := &doublyLinkedListNode{}
	top := &doublyLinkedListNode{}
	bottom.prev = top
	top.next = bottom

	var listSize int
	var numOfDel int
	for i := 0; i < n; i++ {
		if numOfDel > 20 {
			fmt.Println("invalid number of delete operations")
			return
		}
		if listSize > 1000000 {
			fmt.Println("invalid list length")
			return
		}

		sc.Scan()
		opk := strings.Split(sc.Text(), " ")
		if opk[0] == "deleteFirst" {
			if listSize == 0 {
				fmt.Println("invalid deleteFirst operation")
				return
			}

			lastNode := top.next
			top.next = lastNode.next
			lastNode.next.prev = top
			lastNode = nil

			listSize--
			continue
		} else if opk[0] == "deleteLast" {
			if listSize == 0 {
				fmt.Println("invalid deleteLast operation")
				return
			}

			firstNode := bottom.prev
			bottom.prev = firstNode.prev
			firstNode.prev.next = bottom
			firstNode = nil

			listSize--
			continue
		}

		if len(opk) != 2 {
			fmt.Println("invalid operation and key")
			return
		}
		k, err := strconv.Atoi(opk[1])
		if err != nil {
			fmt.Println(err)
			return
		}
		if k < 0 { // || k > 1000000000 { 制約にkeyは10^9以下って書いているのに10^9以上のインプットが問題に出てくるのでコメントアウト
			fmt.Println("invalid value of a key")
			return
		}

		switch opk[0] {
		case "insert":
			node := &doublyLinkedListNode{
				key:  k,
				next: top.next,
				prev: top,
			}
			node.next.prev = node
			top.next = node

			listSize++
		case "delete":
			numOfDel++
			if listSize == 0 {
				fmt.Println("invalid delete operation")
				return
			}

			for n := top.next; n != bottom; n = n.next {
				if n.key == k {
					n.next.prev = n.prev
					n.prev.next = n.next
					n = nil
					break
				}
			}

			listSize--
		default:
			fmt.Println("invalid operation")
			return
		}
	}

	for n := top.next; n != bottom; n = n.next {
		fmt.Print(n.key)
		if n.next != bottom {
			fmt.Print(" ")
		}
	}
	fmt.Println("")
}

// スライスで実装した場合
func DoublyLinkedListNotDLLButSlice() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	n, err := strconv.Atoi(sc.Text())
	if err != nil {
		fmt.Println(err)
		return
	}
	if n < 0 || n > 2000000 {
		fmt.Println("invalid number of operation")
		return
	}

	list := make([]int, 0, n)
	var numOfDel int
	for i := 0; i < n; i++ {
		if numOfDel > 20 {
			fmt.Println("invalid number of delete operations")
			return
		}
		if len(list) > 1000000 {
			fmt.Println("invalid list length")
			return
		}

		sc.Scan()
		opk := strings.Split(sc.Text(), " ")
		if opk[0] == "deleteFirst" {
			if len(list) == 0 {
				fmt.Println("invalid deleteFirst operation")
				return
			}
			list = list[:len(list)-1]
			continue
		} else if opk[0] == "deleteLast" {
			if len(list) == 0 {
				fmt.Println("invalid deleteLast operation")
				return
			}
			list = list[1:]
			continue
		}

		if len(opk) != 2 {
			fmt.Println("invalid operation and key")
			return
		}
		k, err := strconv.Atoi(opk[1])
		if err != nil {
			fmt.Println(err)
			return
		}
		if k < 0 { // || k > 1000000000 { 制約にkeyは10^9以下って書いているのに10^9以上のインプットが問題に出てくるのでコメントアウト
			fmt.Println("invalid value of a key")
			return
		}

		switch opk[0] {
		case "insert":
			list = append(list, k)
		case "delete":
			if len(list) == 0 {
				fmt.Println("invalid delete operation")
				return
			}
			for j := len(list) - 1; j >= 0; j-- {
				if k == list[j] {
					list = append(list[:j], list[j+1:]...)
					break
				}
			}
			numOfDel++
		default:
			fmt.Println("invalid operation")
			return
		}
	}

	for i := len(list) - 1; i >= 0; i-- {
		fmt.Print(list[i])
		if i != 0 {
			fmt.Print(" ")
		}
	}
	fmt.Println("")
}

func AreasOnTheCrossSectionDiaglam() {

}
