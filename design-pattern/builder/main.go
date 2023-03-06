package main

import (
	"fmt"
	"os"
)

// Golangでは抽象クラスはinterfaceでしか実装できない。
type Builder interface {
	ExistTitle() bool
	MakeTitle(title string)
	MakeString(str string)
	MakeItems(item []string)
	Close()
}

type Director struct {
	builder Builder
}

func NewDirector(builder Builder) Director {
	return Director{builder}
}

func (d Director) Construct() {
	d.builder.MakeTitle("Greeting")
	d.builder.MakeString("朝から昼にかけて")
	d.builder.MakeItems([]string{"おはようございます。", "こんにちは。"})
	d.builder.MakeString("夜に")
	d.builder.MakeItems([]string{"こんばんは。", "おやすみなさい。", "さようなら。"})
	d.builder.Close()
}

type TextBuilder struct {
	Builder
	buffer   []byte
	hasTitle bool
}

func NewTextBuilder() (tb TextBuilder) {
	return TextBuilder{buffer: make([]byte, 0)}
}

func (tb TextBuilder) ExistTitle() bool {
	return tb.hasTitle
}

func (tb *TextBuilder) MakeTitle(title string) {
	tb.buffer = append(tb.buffer, []byte("======================\n")...)
	tb.buffer = append(tb.buffer, []byte("「"+title+"」\n")...)
	tb.buffer = append(tb.buffer, []byte("\n")...)
	tb.hasTitle = true
}

func (tb *TextBuilder) MakeString(str string) {
	if !tb.hasTitle {
		return
	}
	tb.buffer = append(tb.buffer, []byte("□"+str+"\n")...)
	tb.buffer = append(tb.buffer, []byte("\n")...)
}

func (tb *TextBuilder) MakeItems(items []string) {
	if !tb.hasTitle {
		return
	}
	for _, item := range items {
		tb.buffer = append(tb.buffer, []byte("　・"+item+"\n")...)
	}
	tb.buffer = append(tb.buffer, []byte("\n")...)
}

func (tb *TextBuilder) Close() {
	if !tb.hasTitle {
		return
	}
	tb.buffer = append(tb.buffer, []byte("======================\n")...)
}

func (tb TextBuilder) GetResult() string {
	return string(tb.buffer)
}

type HTMLBuilder struct {
	Builder
	filename string
	writer   *os.File
	hasTitle bool
}

func (hb HTMLBuilder) ExistTitle() bool {
	return hb.hasTitle
}

func (hb *HTMLBuilder) MakeTitle(title string) {
	hb.filename = title + ".html"
	var err error
	hb.writer, err = os.OpenFile(hb.filename, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	fmt.Fprintln(hb.writer, "<html><head><title>"+title+"</title></head><body>")
	fmt.Fprintln(hb.writer, "<h1>"+title+"</h1>")
	hb.hasTitle = true
}

func (hb *HTMLBuilder) MakeString(str string) {
	if !hb.hasTitle {
		return
	}
	fmt.Fprintln(hb.writer, "<p>"+str+"</p>")
}

func (hb *HTMLBuilder) MakeItems(items []string) {
	if !hb.hasTitle {
		return
	}
	fmt.Fprintln(hb.writer, "<ul>")
	for _, item := range items {
		fmt.Fprintln(hb.writer, "<li>"+item+"</li>")

	}
	fmt.Fprintln(hb.writer, "</ul>")
}

func (hb *HTMLBuilder) Close() {
	if !hb.hasTitle {
		return
	}
	fmt.Fprintln(hb.writer, "</body></html>")
	hb.writer.Close()
}

func (hb HTMLBuilder) GetResult() string {
	return hb.filename
}

func main() {
	fmt.Println(os.Args)
	if len(os.Args) < 2 {
		usage()
		return
	}

	switch os.Args[1] {
	case "plain":
		tb := NewTextBuilder()
		d := NewDirector(&tb)
		d.Construct()
		result := tb.GetResult()
		fmt.Println(result)
	case "html":
		hb := HTMLBuilder{}
		d := NewDirector(&hb)
		d.Construct()
		filename := hb.GetResult()
		fmt.Println(filename + "が作成されました。")
	default:
		usage()
	}
}

func usage() {
	fmt.Println("Usage: go run ./design-pattern/builder/main.go plain プレーンテキストで文書作成")
	fmt.Println("Usage: go run ./design-pattern/builder/main.go html  HTMLで文書作成")
}
