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
}

func NewBookShelf(maxsize int) *BookShelf {
	bs := new(BookShelf)
	bs.books = make([]*Book, 0, maxsize)
	return bs
}

func (bs *BookShelf) GetBookAt(index int) *Book {
	if index >= bs.GetLength() {
		return nil
	}
	return bs.books[index]
}

func (bs *BookShelf) AppendBook(book *Book) {
	bs.books = append(bs.books, book)
}

func (bs *BookShelf) GetLength() int {
	return len(bs.books)
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

	// 練習問題1-1のため記述
	bookShelf.AppendBook(&Book{"Etc."})
	bookShelf.AppendBook(&Book{"Etc."})

	it := bookShelf.Iterator()
	for it.HasNext() {
		book := it.Next().(*Book)
		fmt.Println(book.GetName())
	}
}
