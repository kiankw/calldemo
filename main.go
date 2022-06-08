package main

import (
	"fmt"
)

func Hello(n int) {
	fmt.Println(n)
}
func A() {
	B(func(n int) {
		Hello(n)
	})
}
func B(fp func(int)) {
	C(func(n int) {
		fp(n)
	})
}
func C(fp func(int)) {
	D(func(n int) {
		fp(n)
	})
}
func D(fp func(int)) {
	fp(123)
}
func main() {
	A()
}
