package main

import (
	"fmt"
	"sync"
)

// Принимает число "i", сумму s и waitgroup wg, считает квадрат числа и выводит.
func square(i int, s *int, wg *sync.WaitGroup) {
	*s += i * i
	wg.Done()
}

func main() {
	// Создаем массив int и переменную s, в которую запишем сумму.
	ints := []int{2, 4, 6, 8, 10}
	var s int
	var w sync.WaitGroup

	for _, v := range ints { // Выполняем горутину.
		w.Add(1)
		go square(v, &s, &w)
	}
	w.Wait()
	fmt.Println(s)
}
