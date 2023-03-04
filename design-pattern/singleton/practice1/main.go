package main

import (
	"fmt"
	"sync"
)

type TickMaker struct {
	ticket int
	sync.Mutex
}

var tickMaker = &TickMaker{ticket: 1000}

func GetNextTicketNumber() int {
	tickMaker.Lock()
	n := tickMaker.ticket
	tickMaker.ticket++
	tickMaker.Unlock()
	return n
}

func main() {
	m := map[int]struct{}{}
	wg := sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			n := GetNextTicketNumber()
			fmt.Println(n)
			if _, ok := m[n]; ok {
				fmt.Println("Duplicate on ", n)
			} else {
				m[n] = struct{}{}
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
