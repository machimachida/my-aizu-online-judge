package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type FileIO interface {
	readFromFile(filename string) error
	writeToFile(filename string) error
	setValue(key, value string)
	getValue(key string) string
}

type FileProperties struct {
	FileIO
	property map[string]string
	year     string
	month    string
	day      string
}

func NewFileProperties() FileProperties {
	return FileProperties{property: map[string]string{}}
}

func (fp FileProperties) readFromFile(filename string) error {
	filep, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer filep.Close()

	scanner := bufio.NewScanner(filep)
	for scanner.Scan() {
		pair := strings.Split(scanner.Text(), "=")
		if len(pair) != 2 {
			return errors.New(fmt.Sprintln("Invalid input property", pair, filename))
		}
		fp.setValue(pair[0], pair[1])
	}

	return nil
}

func (fp FileProperties) writeToFile(filename string) error {
	filep, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer filep.Close()

	writer := bufio.NewWriter(filep)
	header := "#written by FileProperties\n#" + time.Now().Format(time.UnixDate) + "\n"
	var props string
	for k, v := range fp.property {
		props += k + "=" + v + "\n"
	}
	if _, err := writer.Write([]byte(header + props)); err != nil {
		return err
	}

	writer.Flush()
	return nil
}

func (fp FileProperties) setValue(key, value string) {
	fp.property[key] = value
}

func (fp FileProperties) getValue(key string) string {
	return fp.property[key]
}

func main() {
	var f FileIO = NewFileProperties()

	err := f.readFromFile("file.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	f.setValue("year", "2004")
	f.setValue("month", "4")
	f.setValue("day", "21")
	err = f.writeToFile("newfile.txt")
	if err != nil {
		fmt.Println(err)
	}
}
