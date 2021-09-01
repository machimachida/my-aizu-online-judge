package support

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"runtime"
)

func Stdin(f func(), s string) error {
	r, w, err := os.Pipe()
	if err != nil {
		return err
	}
	w.Write([]byte(s))
	w.Close()
	stdin := os.Stdin
	os.Stdin = r
	f()
	os.Stdin = stdin
	return nil
}

/*
fn(function name): aizu/support.Submission
*/
func Submission(f func()) (string, error) {
	file, line := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).FileLine(reflect.ValueOf(f).Pointer())
	fmt.Println(file)
	fmt.Println(line)

	fp, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	var c int
	for scanner.Scan() {
		c++
	}
	return "", nil // TODO
}
