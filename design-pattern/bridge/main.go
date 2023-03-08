package main

import "fmt"

type DisplayInterface interface {
	Open()
	Print()
	Close()
	// Display()と命名すると、CountDisplayからDisplayを呼び出す際、
	// Displayフィールドを呼び出してしまうので、本とはメソッド名を変更。
	DisplayAll()
}

type Display struct {
	// DisplayInterface interfaceは値を入れない場合は書かなくても大丈夫
	displayImpl DisplayImpl
}

func NewDisplay(impl DisplayImpl) Display {
	return Display{displayImpl: impl}
}

func (d Display) Open() {
	d.displayImpl.RawOpen()
}

func (d Display) Print() {
	d.displayImpl.RawPrint()
}

func (d Display) Close() {
	d.displayImpl.RawClose()
}

func (d Display) DisplayAll() {
	d.Open()
	d.Print()
	d.Close()
}

type CountDisplay struct {
	Display
}

func NewCountDisplay(impl DisplayImpl) CountDisplay {
	return CountDisplay{Display: NewDisplay(impl)}
}

func (cd CountDisplay) MultiDisplay(times int) {
	cd.Open()
	for i := 0; i < times; i++ {
		cd.Print()
	}
	cd.Close()
}

type DisplayImpl interface {
	RawOpen()
	RawPrint()
	RawClose()
}

type StringDisplayImpl struct {
	//	DisplayImpl
	str   string
	width int
}

func NewStringDisplayImpl(str string) StringDisplayImpl {
	return StringDisplayImpl{str: str, width: len(str)}
}

func (sdi StringDisplayImpl) RawOpen() {
	sdi.printLine()
}

func (sdi StringDisplayImpl) RawPrint() {
	fmt.Println("|" + sdi.str + "|")
}

func (sdi StringDisplayImpl) RawClose() {
	sdi.printLine()
}

func (sdi StringDisplayImpl) printLine() {
	fmt.Print("+")
	for i := 0; i < sdi.width; i++ {
		fmt.Print("-")
	}
	fmt.Println("+")
}

func main() {
	var (
		d1 DisplayInterface = NewDisplay(NewStringDisplayImpl("Hello, Japan."))
		d2 DisplayInterface = NewCountDisplay(NewStringDisplayImpl("Hello, World."))
		d3 CountDisplay     = NewCountDisplay(NewStringDisplayImpl("Hello, Unierse."))
	)
	d1.DisplayAll()
	d2.DisplayAll()
	d3.DisplayAll()
	d3.MultiDisplay(5)
}
