package main

import (
	"fmt"
	"practice-go/design-pattern/prototype/framework"
)

type MessageBox struct {
	framework.Product
	decochar rune
}

func NewMessageBox(decochar rune) MessageBox {
	return MessageBox{decochar: decochar}
}

func (mb MessageBox) Use(s string) {
	length := len(s)
	for i := 0; i < length+4; i++ {
		fmt.Print(string(mb.decochar))
	}
	fmt.Println()
	fmt.Println(string(mb.decochar) + " " + s + " " + string(mb.decochar))
	for i := 0; i < length+4; i++ {
		fmt.Print(string(mb.decochar))
	}
	fmt.Println()
}

func (mb MessageBox) CreateClone() framework.Product {
	var p framework.Product = mb
	return p
}

type UnderlinePen struct {
	framework.Product
	ulchar rune
}

func NewUnderlinePen(ulchar rune) UnderlinePen {
	return UnderlinePen{ulchar: ulchar}
}

func (up UnderlinePen) Use(s string) {
	length := len(s)
	fmt.Println("\"" + s + "\"")
	fmt.Print(" ")
	for i := 0; i < length; i++ {
		fmt.Print(string(up.ulchar))
	}
	fmt.Println()
}

func (up UnderlinePen) CreateClone() framework.Product {
	var p framework.Product = up
	return p
}

func main() {
	manager := framework.NewManager()
	upen := NewUnderlinePen('~')
	mbox := NewMessageBox('*')
	sbox := NewMessageBox('/')
	manager.Register("strong message", upen)
	manager.Register("warning box", mbox)
	manager.Register("slash box", sbox)

	p1 := manager.Create("strong message")
	if p1 != nil {
		p1.Use("Hello, world.")
	}
	p2 := manager.Create("warning box")
	if p2 != nil {
		p2.Use("Hello, world.")
	}
	p3 := manager.Create("slash box")
	if p3 != nil {
		p3.Use("Hello, world.")
	}
}
