package main

import "fmt"

type Aggregate interface {
	Iterator
}

type Iterator interface {
	HasNext() bool
	Next() any
}

type Book struct {
	name string
}

func (b *Book) GetName() string {
	return b.name
}

type BookShelf struct {
	Aggregate
	books []*Book
	last  int
}

func NewBookShelf(maxsize int) *BookShelf {
	bs := new(BookShelf)
	bs.books = make([]*Book, maxsize)
	return bs
}

func (bs *BookShelf) GetBookAt(index int) *Book {
	return bs.books[index]
}

func (bs *BookShelf) AppendBook(book *Book) {
	bs.books[bs.last] = book
	bs.last++
}

func (bs *BookShelf) GetLength() int {
	return bs.last
}

func (bs *BookShelf) Iterator() Iterator {
	return NewBookShelfIterator(bs)
}

type BookShelfIterator struct {
	Iterator
	bookShelf *BookShelf
	index     int
}

func NewBookShelfIterator(bs *BookShelf) *BookShelfIterator {
	return &BookShelfIterator{bookShelf: bs, index: 0}
}

func (bsi *BookShelfIterator) HasNext() bool {
	return bsi.index < bsi.bookShelf.GetLength()
}

func (bsi *BookShelfIterator) Next() any {
	book := bsi.bookShelf.GetBookAt(bsi.index)
	bsi.index++
	return book
}

func main() {
	bookShelf := NewBookShelf(4)
	bookShelf.AppendBook(&Book{"Around the World in 80 Days"})
	bookShelf.AppendBook(&Book{"Bible"})
	bookShelf.AppendBook(&Book{"Cinderella"})
	bookShelf.AppendBook(&Book{"Daddy-Long-Legs"})

	it := bookShelf.Iterator()
	for it.HasNext() {
		book := it.Next().(*Book)
		fmt.Println(book.GetName())
	}
}
