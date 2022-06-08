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
	C(func(n int) {
		fp(n)
	})
}
func C(fp func(int)) {
	fp(123)
}
