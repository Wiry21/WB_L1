package main

import "fmt"

func main() {
	a := 2
	b := 7
	fmt.Println("a =", a, "b =", b)
	a = a + b
	b = a - b
	a = a - b
	fmt.Println("a =", a, "b =", b)
	a, b = b, a
	fmt.Println("a =", a, "b =", b)
}
