package main

import (
	"fmt"
)

func main() {
	A()
}
func Hello() {
	fmt.Println("Hello")
}
func A() {
	B(func(n int) {
		Hello()
	})
}
func B(fp func(int)) {
	C(fp)
}
func C(fp func(int)) {
	D(fp)
}
func D(fp func(int)) {
	fp(123)
}
