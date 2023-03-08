package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

const (
	HANDVALUE_GUU int = 0
	HANDVALUE_CHO int = 1
	HANDVALUE_PAA int = 2
)

var (
	Hands [3]Hand   = [3]Hand{{HANDVALUE_GUU}, {HANDVALUE_CHO}, {HANDVALUE_PAA}}
	Names [3]string = [3]string{"グー", "チョキ", "パー"}
)

type Hand struct {
	handvalue int
}

func (h Hand) IsStrongerThan(hand Hand) bool {
	return h.fight(hand) == 1
}

func (h Hand) IsWeakerThan(hand Hand) bool {
	return h.fight(hand) == -1
}

func (h Hand) fight(hand Hand) int {
	if h == hand {
		return 0
	} else if (h.handvalue+1)%3 == h.handvalue {
		return 1
	} else {
		return -1
	}
}

func (h Hand) String() string {
	return Names[h.handvalue]
}

type Strategy interface {
	NextHand() Hand
	Study(win bool)
}

type WinningStrategy struct {
	won      bool
	prevHand Hand
}

func NewWinningStrategy() *WinningStrategy {
	return &WinningStrategy{}
}

func (ws *WinningStrategy) NextHand() Hand {
	if !ws.won {
		ws.prevHand = Hands[rand.Intn(3)]
	}
	return ws.prevHand
}

func (ws *WinningStrategy) Study(win bool) {
	ws.won = win
}

type ProbStrategy struct {
	prevHandValue    int
	currentHandValue int
	history          [3][3]int
}

func NewProbStrategy() *ProbStrategy {
	return &ProbStrategy{
		prevHandValue:    0,
		currentHandValue: 0,
		history: [3][3]int{
			{1, 1, 1}, {1, 1, 1}, {1, 1, 1},
		},
	}
}

func (ps *ProbStrategy) NextHand() Hand {
	bet := rand.Intn(ps.getSum(ps.currentHandValue))
	var handvalue int
	if bet < ps.history[ps.currentHandValue][0] {
		handvalue = 0
	} else if bet < ps.history[ps.currentHandValue][0]+ps.history[ps.currentHandValue][1] {
		handvalue = 1
	} else {
		handvalue = 2
	}
	ps.prevHandValue = ps.currentHandValue
	ps.currentHandValue = handvalue
	return Hands[handvalue]
}

func (ps ProbStrategy) getSum(hv int) int {
	sum := 0
	for i := 0; i < 3; i++ {
		sum += ps.history[hv][i]
	}
	return sum
}

func (ps *ProbStrategy) Study(win bool) {
	if win {
		ps.history[ps.prevHandValue][ps.currentHandValue]++
	} else {
		ps.history[ps.prevHandValue][(ps.currentHandValue+1)%3]++
		ps.history[ps.prevHandValue][(ps.currentHandValue+2)%3]++
	}
}

type Player struct {
	name      string
	strategy  Strategy
	winCount  int
	loseCount int
	gameCount int
}

func NewPlayer(name string, strategy Strategy) *Player {
	return &Player{name: name, strategy: strategy, winCount: 0, loseCount: 0, gameCount: 0}
}

func (p Player) NextHand() Hand {
	return p.strategy.NextHand()
}

func (p *Player) Win() {
	p.strategy.Study(true)
	p.winCount++
	p.gameCount++
}

func (p *Player) Lose() {
	p.strategy.Study(false)
	p.loseCount++
	p.gameCount++
}

func (p *Player) Even() {
	p.gameCount++
}

func (p *Player) String() string {
	return "[" + p.name + ":" + strconv.Itoa(p.gameCount) + " games, " +
		strconv.Itoa(p.winCount) + " win, " + strconv.Itoa(p.loseCount) + " lose" + "]"
}

func main() {
	player1 := NewPlayer("Taro", NewWinningStrategy())
	player2 := NewPlayer("Hana", NewProbStrategy())
	for i := 0; i < 10000; i++ {
		nextHand1 := player1.NextHand()
		nextHand2 := player2.NextHand()
		if nextHand1.IsStrongerThan(nextHand2) {
			fmt.Println("Winner:", player1)
			player1.Win()
			player2.Lose()
		} else if nextHand2.IsStrongerThan(nextHand1) {
			fmt.Println("Winner:", player2)
			player1.Lose()
			player2.Win()
		} else {
			fmt.Println("Even...")
			player1.Even()
			player2.Even()
		}
	}
	fmt.Println("Total result:")
	fmt.Println(player1)
	fmt.Println(player2)
}
