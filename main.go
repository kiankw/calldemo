package main

import (
	"fmt"

	"github.com/kiankw/gotmp"
)

func main() {
	A()
	B()
	C()
	D()
	gotmp.PrintHello()
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
