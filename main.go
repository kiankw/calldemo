package main

import "fmt"

func AF(s string) {
	fmt.Print("A: ")
	fmt.Println(s)
}

func BF(b bool) {
	fmt.Print("B: ")
	fmt.Println(b)
}

func CF(i int) {
	fmt.Print("C: ")
	fmt.Println(i)
}

func main() {
	A(AF)

	p := People{Name: "Tom"}
	p.B(BF)

	var a Animal
	a = Cat{Name: "Tome"}
	a.C(CF)
}

func A(fp func(string)) {
	fp("Hello")
}

type People struct {
	Name string
}

func (p People) B(fp func(bool)) {
	fp(true)
}

type Animal interface {
	C(fp func(int))
}

type Cat struct {
	Name string
}

func (c Cat) C(fp func(int)) {
	fp(123)
}
