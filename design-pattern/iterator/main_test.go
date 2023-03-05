package main

import "testing"

// 練習問題1-1のためのテスト

// append()では内部で再アロケートすると思うが、自前で書いたものとの速度の違いを確認
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

// 毎度毎度Reallocateすると、めちゃくちゃ遅い。
func BenchmarkAppendBook_Reallocate(b *testing.B) {
	bookShelf := NewBookShelf(4)
	AppendBook := func(bs *BookShelf, b *Book) {
		if bs.last == len(bs.books) {
			newbs := make([]*Book, bs.last+1)
			copy(newbs, bs.books)
			newbs[bs.last] = b
			bs.books = newbs
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

// appendの実装と同じCapacityの伸ばし方をしている。appendより速い。
func BenchmarkAppendBook_ReallocateDoubleCapacityAndLength(b *testing.B) {
	bookShelf := NewBookShelf(4)
	AppendBook := func(bs *BookShelf, b *Book) {
		if bs.last == cap(bs.books) {
			newbs := make([]*Book, bs.last*2)
			copy(newbs, bs.books)
			newbs[bs.last] = b
			bs.books = newbs
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

// こちらはlenとcapの増やし方をappendと揃えたもの。
// こちらもappendより速いが、よっぽどのことがない限りはappendを素直に使うべきだと思う。
func BenchmarkAppendBook_ReallocateDoubleCapacity(b *testing.B) {
	bookShelf := NewBookShelf(4)
	AppendBook := func(bs *BookShelf, b *Book) {
		if bs.last == cap(bs.books) {
			newbs := make([]*Book, bs.last+1, bs.last*2)
			copy(newbs, bs.books)
			newbs[bs.last] = b
			bs.books = newbs
			bs.last++
			return
		} else if bs.last == len(bs.books) {
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
