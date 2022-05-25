package main

import (
	"fmt"
)

func VirtualCall(a Animal) {
	a.HelloAnimal()
}

func main() {
	// Static call
	A()
	B()
	// Special call
	var c Cat = Cat{Name: "Tom"}
	c.HelloCat()
	var m Mouse = Mouse{Name: "Jerry"}
	m.HelloMouse()
	// Virtual call
	var a1 Animal = Cat{Name: "Tom"}
	var a2 Animal = Mouse{Name: "Jerry"}
	VirtualCall(a1)
	VirtualCall(a2)

	MyTestBranch(true)
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

type Animal interface {
	HelloAnimal()
}

type Cat struct {
	Name string
}

func (c Cat) HelloCat() {
	fmt.Println("I am cat", c.Name)
}
func (c Cat) HelloAnimal() {
	fmt.Println("I am animal cat", c.Name)
}

type Mouse struct {
	Name string
}

func (m Mouse) HelloMouse() {
	fmt.Println("I am mouse", m.Name)
}
func (m Mouse) HelloAnimal() {
	fmt.Println("I am animal mouse", m.Name)
}

func MyTestBranch(flag bool) {
	if flag {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}