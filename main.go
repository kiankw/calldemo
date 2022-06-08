package main

import (
	"fmt"
)

func Hello(n int) {
	fmt.Println(n)
}

func main() {
	A(func(n int) {
		Hello(n)
	})
}
func A(fp func(int)) {
	B(func(n int) {
		fp(n)
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
