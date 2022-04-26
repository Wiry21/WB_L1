package main

import "fmt"

func in(arr []int, ch1 chan int) {
	defer close(ch1)
	for _, val := range arr {
		ch1 <- val
	}
}

func out(ch2 chan int, ch1 chan int) {
	defer close(ch2)
	for val := range ch1 {
		ch2 <- val * 2
	}
}

func main() {
	// Создаем каналы ch1 и ch2 и массив, из которого будем брать числа.
	ch1 := make(chan int)
	ch2 := make(chan int)
	arr := []int{1, 2, 3, 4, 5}

	// Запускаем горутины и печатаем содержимое второго канала.
	go in(arr, ch1)
	go out(ch2, ch1)

	for val := range ch2 {
		fmt.Print(val, " ")
	}
}
