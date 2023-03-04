package main

import "testing"

// 練習問題1-1のためのテスト

// append()では内部で再アロケートすると思うが、明示的に書いたものとの速度の違いを確認
func BenchmarkAppendBook_Append(b *testing.B) {
	bookShelf := NewBookShelf(4)
	AppendBook := func(bs *BookShelf, b *Book) {
		if bs.last == len(bs.books) {
			bs.books = append(bs.books, b)
			bs.last++
			return
		}

		bs.books[bs.last] = b
		bs.last++
	}

	for i := 0; i < b.N; i++ {
		AppendBook(bookShelf, &Book{"Example"})
	}
}

// TODO: 要調査
//
// めちゃくちゃ遅い。メモリ開放されずに再定義しているから？
func BenchmarkAppendBook_Reallocate(b *testing.B) {
	bookShelf := NewBookShelf(4)
	AppendBook := func(bs *BookShelf, b *Book) {
		if bs.last == len(bs.books) {
			bs.books = make([]*Book, bs.last+1)
			bs.books = append(bs.books, b)
			bs.last++
			return
		}

		bs.books[bs.last] = b
		bs.last++
	}

	for i := 0; i < b.N; i++ {
		AppendBook(bookShelf, &Book{"Example"})
	}
}
