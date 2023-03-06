package main

import (
	"fmt"
	"os"
	"reflect"
)

type ItemInterface interface {
	MakeHTML() string
}

type Item struct {
	ItemInterface
	caption string
}

func NewItem(caption string) Item {
	return Item{caption: caption}
}

type Link struct {
	Item
	url string
}

func NewLink(caption, url string) Link {
	return Link{Item: NewItem(caption), url: url}
}

type Tray struct {
	Item
	tray []Item
}

func NewTray(caption string) Tray {
	return Tray{Item: NewItem(caption), tray: make([]Item, 0)}
}

func (t *Tray) Add(item Item) {
	t.tray = append(t.tray, item)
}

type PageInterface interface {
	MakeHTML() string
}

type Page struct {
	PageInterface
	title   string
	author  string
	content []Item
}

func NewPage(title, author string) Page {
	return Page{title: title, author: author}
}

func (p *Page) Add(item Item) {
	p.content = append(p.content, item)
}

func (p Page) Output() {
	filename := p.title + ".html"
	writer, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	fmt.Fprint(writer, p.MakeHTML())
	writer.Close()
	fmt.Println(filename + "を作成しました。")
}

type FactoryInterface interface {
	CreateLink(caption, url string) Link
	CreateTray(caption string) Tray
	CreatePage(title, author string) Page
}

type Factory struct {
	FactoryInterface
}

func GetFactory(structType reflect.Type) Factory {
	elem := reflect.New(structType).Elem()
	if !elem.CanInterface() {
		panic(fmt.Sprintf("failed to reflect.Value.CanInterface: %v\n", elem))
	}

	factory, ok := elem.Interface().(Factory)
	if !ok {
		panic(fmt.Sprintf("failed to cast Factory: %v\n", elem))
	}
	return factory
}

func main() {
	a := reflect.New(reflect.TypeOf(Item{}))
	fmt.Println(a, a.IsNil())
	b := a.Interface()
	b1, ok := b.(Item)
	fmt.Println(b1, ok)

	b2, ok := b.(*Item)
	fmt.Println(b2, ok)

	if a.CanAddr() {
		c := a.Addr().Interface()
		c1, ok := c.(Item)
		fmt.Println(c1, ok)

		c2, ok := c.(*Item)
		fmt.Println(c2, ok)
	}

	d := a.Elem().Interface()
	d1, ok := d.(Item)
	fmt.Println(d1, ok)

	d2, ok := d.(*Item)
	fmt.Println(d2, ok)
}
