package main

import "fmt"

type Helloer interface {
	Hello() string
}

type Cat struct {
	ID int
}

func (c Cat) Hello() string {
	return "Cat Hello"
}

type Mouse struct {
	Name string
}

func (m Mouse) Hello() string {
	return "Mouse Hello"
}

func Test(h Helloer) {
	fmt.Println(h.Hello())
}

func main() {
	c := Cat{ID: 1}
	Test(c)
}
