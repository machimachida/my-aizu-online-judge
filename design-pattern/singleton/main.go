package main

import (
	"fmt"
)

type Singleton struct{}

var singleton = &Singleton{}

func GetInstance() *Singleton {
	return singleton
}

func main() {
	fmt.Println("Start.")
	obj1 := GetInstance()
	obj2 := GetInstance()
	fmt.Println(obj1 == obj2)
	fmt.Println("End.")
}
