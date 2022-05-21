package main

import (
	"fmt"

	"github.com/kiankw/gotmp"
)

type Animal interface {
	Hello()
}
type Cat struct {
	Name string
}

func (c Cat) Hello() {
	fmt.Println("I am cat", c.Name)
}

type Mouse struct {
	Name string
}

func (m Mouse) Hello() {
	fmt.Println("I am mouse", m.Name)
}

func main() {
	A()

	var a1 Animal = Cat{Name: "Tom"}
	var a2 Animal = Mouse{Name: "Jerry"}
	a1.Hello()
	a2.Hello()
}

func A() {
	fmt.Println("Hello")
	A1()
	A2()
	AB()
	ABC()
	ABCD()
	AHello()
}

func B() {
	fmt.Println("Hello")
	AB()
	ABC()
	ABCD()
}

func C() {
	fmt.Println("Hello")
	ABC()
	ABCD()
}

func D() {
	fmt.Println("Hello")
	ABCD()
}

func A1() {
	fmt.Println("Hello")
}
func A2() {
	fmt.Println("Hello")
}

func AB() {
	fmt.Println("Hello")
}

func ABC() {
	fmt.Println("Hello")

}

func ABCD() {
	fmt.Println("Hello")
}

func AHello() {
	fmt.Println("Hello")
	gotmp.PrintHello()
}
