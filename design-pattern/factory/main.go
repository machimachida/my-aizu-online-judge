package main

import (
	"fmt"
	"practice-go/design-pattern/factory/framework"
	"practice-go/design-pattern/factory/idcard"
)

func main() {
	var (
		factory framework.Factory = idcard.NewIDCardFactory()
		card1   framework.Product = factory.Create("aa")
		card2   framework.Product = factory.Create("bb")
		card3   framework.Product = factory.Create("cc")
	)
	card1.Use()
	card2.Use()
	card3.Use()
	fmt.Println(factory.(*idcard.IDCardFactory).GetOwners())
}
