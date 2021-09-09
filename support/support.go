package support

import (
	"os"
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
