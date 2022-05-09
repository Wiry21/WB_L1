package main

import (
	"fmt"
)

// Go кодирует строку как массив байтов. Произвольные байты. Индексация строк дает байты, а не символы.
// Поэтому при индексации по массиву строки, мы можем получить не набор символов, как ожидали.
// Чтобы этого не было, нужно использовать rune.

var justString string

// Создание строки длиной size.
func createHugeString(size int) (s string) {
	for i := 0; i < size; i++ {
		s += "А"
	}
	return
}

func somefunc() {
	s := createHugeString(1 << 10)
	r := []rune(s)

	justString = string(s[:5])
	justStringRune := string(r[:5])

	fmt.Println(justString)
	fmt.Println(justStringRune)
}

func main() {
	somefunc()
}
