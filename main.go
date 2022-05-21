package main

import (
	"fmt"
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
	// Static call
	A()
	B()
	// Special call
	var c Cat = Cat{Name: "Tom"}
	c.Hello()
	var m Mouse = Mouse{Name: "Jerry"}
	m.Hello()
	// Virtual call
	var a1 Animal = Cat{Name: "Tom"}
	var a2 Animal = Mouse{Name: "Jerry"}
	a1.Hello()
	a2.Hello()
}

func A() {
	A1()
	AB()
	fmt.Println("Hello A")
}
func B() {
	B1()
	AB()
	fmt.Println("Hello B")
}
func A1() {
	fmt.Println("Hello A1")
}
func AB() {
	fmt.Println("Hello AB")
}
func B1() {
	fmt.Println("Hello B1")
}
