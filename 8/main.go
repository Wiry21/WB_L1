package main

import (
	"fmt"
	"log"
)

func main() {
	// Создаем переменные n, i, valBite в которые запишем число в котором хотим поменять бит, номер бита
	// и во что установить бит.
	var x int64
	var i, valBite int
	fmt.Print("Введите число в котором хотите поменять бит, номер бита и во что установить - 1 или 0: ")
	_, err := fmt.Scan(&x, &i, &valBite)
	if err != nil {
		log.Println(err.Error())
	}

	fmt.Printf("Число: %d в двоичной системе до изменения %b\n", x, x)
	// Если хотим установить то в 1, то используем bitwise OR, если в 0, то AND NOT.
	if valBite == 1 {
		x |= 1 << i
	} else {
		x &^= 1 << i
	}

	fmt.Printf("Число: %d в двоичной системе после изменения %b\n", x, x)
}
