package factory

import (
	"fmt"
	"os"
	iterator "practice-go/design-pattern/abstract-factory/Iterator"
)

type ItemInterface interface {
	MakeHTML() string
}

type Item struct {
	ItemInterface
	Caption string
}

func NewItem(caption string) Item {
	return Item{Caption: caption}
}

type LinkInterface interface {
	ItemInterface
}

type Link struct {
	LinkInterface
	Item
	URL string
}

func NewLink(caption, url string) Link {
	return Link{Item: NewItem(caption), URL: url}
}

type TrayInterface interface {
	ItemInterface
	iterator.Iterator
	Add(item ItemInterface)
}

type Tray struct {
	TrayInterface
	Item
	tray []ItemInterface
}

func NewTray(caption string) Tray {
	return Tray{Item: NewItem(caption), tray: make([]ItemInterface, 0)}
}

func (t *Tray) Add(item ItemInterface) {
	t.tray = append(t.tray, item)
}

func (t Tray) Iterator() iterator.Iterator {
	return NewTrayIterator(t)
}

type TrayIterator struct {
	iterator.Iterator
	tray  *Tray
	index int
}

func NewTrayIterator(t Tray) *TrayIterator {
	return &TrayIterator{tray: &t, index: 0}
}

func (ti TrayIterator) HasNext() bool {
	return ti.index < len(ti.tray.tray)
}

func (ti *TrayIterator) Next() any {
	tray := ti.tray.tray[ti.index]
	ti.index++
	return tray
}

type PageInterface interface {
	MakeHTML() string
	Add(item ItemInterface)
	Output()
}

type Page struct {
	PageInterface
	Title   string
	Author  string
	content []ItemInterface
}

func NewPage(title, author string) Page {
	return Page{Title: title, Author: author, content: make([]ItemInterface, 0)}
}

func (p *Page) Add(item ItemInterface) {
	p.content = append(p.content, item)
}

func (p Page) Output() {
	filename := p.Title + ".html"
	writer, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	fmt.Fprint(writer, p.MakeHTML())
	writer.Close()
	fmt.Println(filename + "を作成しました。")
}

func (p Page) Iterator() iterator.Iterator {
	return NewPageIterator(p)
}

type PageIterator struct {
	iterator.Iterator
	page  *Page
	index int
}

func NewPageIterator(p Page) *PageIterator {
	return &PageIterator{page: &p, index: 0}
}

func (pi PageIterator) HasNext() bool {
	return pi.index < len(pi.page.content)
}

func (pi *PageIterator) Next() any {
	page := pi.page.content[pi.index]
	pi.index++
	return page
}

type Factory interface {
	CreateLink(caption, url string) LinkInterface
	CreateTray(caption string) TrayInterface
	CreatePage(title, author string) PageInterface
}
