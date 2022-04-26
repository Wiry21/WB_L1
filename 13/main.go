package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("This is int ", v)
	case string:
		fmt.Println("This is string", v)
	case bool:
		fmt.Println("this is bool", v)
	case chan int:
		fmt.Println("This is channel", v)
	default:
		fmt.Println("this is not a int, string, bool or channel", v)
	}
}

func main() {
	//ch1 := make(chan int)
	do(124)
}
