package alds1

import (
	"fmt"
	"os"
	"practice-go/support"
	"testing"
)

func ExampleStack() {
	samples := []string{"1 2 +", "1 2 + 3 4 - *"}
	for _, s := range samples {
		err := support.Stdin(Stack, s)
		if err != nil {
			fmt.Printf("failed: %#v\n", err)
		}
	}
	// Output: 3
	// -3
}

func BenchmarkStack_slice(b *testing.B) {
	null, err := os.Open(os.DevNull)
	if err != nil {
		fmt.Println(err)
	}
	stdout := os.Stdout
	os.Stdout = null
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := support.Stdin(Stack, "1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + +")
		if err != nil {
			fmt.Printf("failed: %#v\n", err)
		}
	}
	b.StopTimer()
	os.Stdout = stdout
	null.Close()
}

func BenchmarkStack_stack(b *testing.B) {
	null, err := os.Open(os.DevNull)
	if err != nil {
		fmt.Println(err)
	}
	stdout := os.Stdout
	os.Stdout = null
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := support.Stdin(stackStack, "1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + + +")
		if err != nil {
			fmt.Printf("failed: %#v\n", err)
		}
	}
	b.StopTimer()
	os.Stdout = stdout
	null.Close()
}

func ExampleQueue() {
	err := support.Stdin(Queue, "5 100\np1 150\np2 80\np3 200\np4 350\np5 20")
	if err != nil {
		fmt.Printf("failed: %#v\n", err)
	}
	// Output: p2 180
	// p5 400
	// p1 450
	// p3 550
	// p4 800
}

func BenchmarkQueue_slice(b *testing.B) {
	null, err := os.Open(os.DevNull)
	if err != nil {
		fmt.Println(err)
	}
	stdout := os.Stdout
	os.Stdout = null
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := support.Stdin(Queue, "5 100\np1 150\np2 80\np3 200\np4 350\np5 20")
		if err != nil {
			fmt.Printf("failed: %#v\n", err)
		}
	}
	b.StopTimer()
	os.Stdout = stdout
	null.Close()
}

func BenchmarkQueue_map(b *testing.B) {
	null, err := os.Open(os.DevNull)
	if err != nil {
		fmt.Println(err)
	}
	stdout := os.Stdout
	os.Stdout = null
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := support.Stdin(QueueMap, "5 100\np1 150\np2 80\np3 200\np4 350\np5 20")
		if err != nil {
			fmt.Printf("failed: %#v\n", err)
		}
	}
	b.StopTimer()
	os.Stdout = stdout
	null.Close()
}

func ExampleDoublyLinkedList() {
	samples := []string{
		"7\ninsert 5\ninsert 2\ninsert 3\ninsert 1\ndelete 3\ninsert 6\ndelete 5",
		"9\ninsert 5\ninsert 2\ninsert 3\ninsert 1\ndelete 3\ninsert 6\ndelete 5\ndeleteFirst\ndeleteLast",
		`21
insert 1
insert 2
insert 3
insert 4
insert 5
insert 6
insert 7
insert 8
delete 5
delete 7
insert 9
deleteLast
deleteLast
insert 10
deleteFirst
deleteFirst
delete 8
insert 7
insert 8
delete 4
insert 1`, `39
insert 1
insert 2
insert 1
insert 1
insert 1
insert 2
insert 2
insert 1
delete 2
deleteLast
deleteFirst
insert 2
insert 2
insert 2
insert 2
insert 2
insert 2
insert 2
insert 2
insert 2
insert 2
deleteFirst
deleteFirst
delete 2
delete 2
deleteFirst
delete 1
insert 3
insert 1
insert 1
delete 2
delete 2
delete 2
deleteLast
insert 4
delete 2
deleteLast
deleteFirst
delete 2
`, `114
insert 8
insert 7
insert 7
insert 8
insert 7
insert 8
insert 1
insert 1
insert 1
insert 1
insert 1
insert 1
insert 1
insert 1
insert 8
insert 8
insert 1
insert 1
insert 8
deleteLast
insert 5
insert 1
insert 1
insert 7
insert 8
deleteFirst
insert 8
insert 7
insert 8
insert 1
insert 1
insert 1
insert 1
delete 8
delete 1
delete 7
insert 1
insert 1
insert 7
insert 1
insert 1
insert 1
delete 1
insert 1
insert 8
insert 5
delete 1
delete 1
insert 7
delete 8
delete 7
insert 8
insert 1
insert 7
insert 5
insert 7
insert 1
insert 1
insert 1
delete 1
insert 1
insert 8
deleteFirst
insert 8
insert 1
insert 1
insert 5
insert 7
insert 1
insert 1
insert 8
insert 1
deleteFirst
insert 7
insert 8
delete 1
insert 7
delete 8
delete 7
insert 1
insert 1
insert 8
insert 7
insert 8
insert 5
insert 7
deleteLast
deleteLast
insert 8
deleteFirst
insert 8
delete 1
insert 7
delete 8
insert 8
insert 5
deleteLast
insert 7
insert 1
insert 1
insert 8
insert 13
deleteLast
deleteLast
insert 8
delete 8
insert 13
delete 1
delete 1
deleteLast
insert 1
insert 1
delete 1
delete 1
`,
	}
	for _, s := range samples {
		err := support.Stdin(DoublyLinkedList, s)
		if err != nil {
			fmt.Printf("failed: %#v\n", err)
		}
	}
	// Output: 6 1 2
	// 1
	// 1 8 7 6 3
	// 1 1 3 2 1
	// 13 13 8 7 5 8 7 7 5 8 7 8 1 7 8 1 7 5 1 1 8 1 1 1 7 5 7 1 8 5 1 7 1 1 1 1 1 8 7 1 1 5 8 1 1 8 8 1 1 1 1 1 1 1
}
