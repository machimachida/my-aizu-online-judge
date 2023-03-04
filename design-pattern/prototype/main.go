package main

import "practice-go/design-pattern/prototype/framework"

type MessageBox struct {
	decochar rune
	*framework.Manager
}

func NewMessageBox(decochar rune) *MessageBox {
	return &MessageBox{decochar: decochar}
}

func (mb MessageBox) Use(s string) {

}
