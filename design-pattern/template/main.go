package main

import "fmt"

type AbstractDisplay interface {
	open()
	print()
	close()
	display()
}

type Abstract5Display struct {
	AbstractDisplay
}

func (a5d Abstract5Display) display() {
	a5d.AbstractDisplay.open()
	for i := 0; i < 5; i++ {
		a5d.AbstractDisplay.print()
	}
	a5d.AbstractDisplay.close()
}

type CharDisplay struct {
	Abstract5Display
	ch rune
}

func NewCharDisplay(ch rune) CharDisplay {
	cd := CharDisplay{ch: ch}
	cd.AbstractDisplay = cd
	return cd
}

func (cd CharDisplay) open() {
	fmt.Print("<<")
}

func (cd CharDisplay) print() {
	fmt.Print(string(cd.ch))
}

func (cd CharDisplay) close() {
	fmt.Println(">>")
}

type StringDisplay struct {
	Abstract5Display
	string string
	width  int
}

func NewStringDisplay(string string) StringDisplay {
	sd := StringDisplay{
		string: string,
		width:  len(string),
	}
	sd.AbstractDisplay = sd
	return sd
}

func (sd StringDisplay) open() {
	sd.printLine()
}

func (sd StringDisplay) print() {
	fmt.Println("|" + sd.string + "|")
}

func (sd StringDisplay) close() {
	sd.printLine()
}

func (sd StringDisplay) printLine() {
	fmt.Print("+")
	for i := 0; i < sd.width; i++ {
		fmt.Print("-")
	}
	fmt.Println("+")
}

func main() {
	var d1 AbstractDisplay = NewCharDisplay('H')
	var d2 AbstractDisplay = NewStringDisplay("Hello, world.")
	d1.display()
	d2.display()
}
