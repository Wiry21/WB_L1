package main

import "fmt"

// Принимает на вход интерфейс i, печатает тип данных.
func checkType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("This is int -", v)
	case string:
		fmt.Println("This is string -", v)
	case bool:
		fmt.Println("This is bool -", v)
	case chan int:
		fmt.Println("This is channel -", v)
	default:
		fmt.Println("I don't know about type -", v)
	}
}

func main() {
	ch := make(chan int)
	checkType(5)
	checkType("hello")
	checkType(true)
	checkType(ch)
	checkType(5.25)
}
