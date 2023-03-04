package idcard

import (
	"fmt"
	"practice-go/design-pattern/factory/framework"
	"strconv"
)

type IDCard struct {
	framework.Product
	owner  string
	serial int
}

func NewIDCard(owner string, serial int) IDCard {
	fmt.Println(owner + "(" + strconv.Itoa(serial) + ")" + "のカードを作ります。")
	c := IDCard{owner: owner, serial: serial}
	c.Product = c
	return c
}

func (c IDCard) Use() {
	fmt.Println(c.owner + "(" + strconv.Itoa(c.serial) + ")のカードを使います。")
}

func (c IDCard) GetOwner() string {
	return c.owner
}

func (c IDCard) GetSerial() int {
	return c.serial
}

type IDCardFactory struct {
	framework.FactoryClass
	owners map[int]string
	serial int
}

func NewIDCardFactory() *IDCardFactory {
	cf := &IDCardFactory{owners: map[int]string{}, serial: 100}
	cf.Factory = cf
	return cf
}

func (cf *IDCardFactory) CreateProduct(owner string) framework.Product {
	c := NewIDCard(owner, cf.serial)
	cf.serial++
	return c
}

func (cf *IDCardFactory) RegisterProduct(product framework.Product) {
	cf.owners[product.(IDCard).serial] = product.(IDCard).GetOwner()
}

func (cf IDCardFactory) GetOwners() map[int]string {
	return cf.owners
}
