package main

import (
	"fmt"
)

func main() {
	A()
}
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
	fp(123)
}
