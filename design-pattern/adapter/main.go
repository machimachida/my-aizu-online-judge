package main

import "fmt"

// 移譲のパターンはPrint型変数にPrintBanner型変数を代入できないので無理そう

type Banner struct {
	string string
}

func (b Banner) showWithParen() {
	fmt.Println("(" + b.string + ")")
}

func (b Banner) showWithAster() {
	fmt.Println("*" + b.string + "*")
}

type Print interface {
	printWeek()
	printStrong()
}

type PrintBanner struct {
	Banner
	Print
}

func (pb PrintBanner) printWeek() {
	pb.showWithParen()
}

func (pb PrintBanner) printStrong() {
	pb.showWithAster()
}

func main() {
	var p Print = PrintBanner{Banner: Banner{string: "Hello"}}
	p.printWeek()
	p.printStrong()
}
