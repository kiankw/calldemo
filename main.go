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
	var m Mouse = Mouse{Name: "Jerry"}
	m.HelloMouse()
	// Virtual call
	var a Animal = Cat{Name: "Tom"}
	a.HelloAnimal()

	// MyTestBranch(true)

	MyPrint(HelloName)
}

func HelloName(name string) string {
	return "Hello World!" + name
}

func MyPrint(f func(name string) string) {
	fmt.Println(f("Tom"))
}

func C() {
	A()
	B()
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

type Dog struct {
	Name string
}

func (d Dog) HelloAnimal() {
	fmt.Println("I am animal dog", d.Name)
}

func MyTestBranch(flag bool) {
	if flag {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}
