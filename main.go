package main

import "fmt"

func main() {
	A(func(s string) {
		fmt.Print("A: ")
		fmt.Println(s)
	})

	p := People{Name: "Tom"}
	p.B(func(b bool) {
		fmt.Print("B: ")
		fmt.Println(b)
	})

	var a Animal
	a = Cat{Name: "Tome"}
	a.C(func(i int) {
		fmt.Print("C: ")
		fmt.Println(i)
	})
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
