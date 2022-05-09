package main

import (
	"fmt"
	"sync"
)

// Принимает число "i" и waitgroup wg, считает квадрат числа и выводит.
func square(i int, wg *sync.WaitGroup) {
	fmt.Println(i * i)
	wg.Done()
}

func main() {
	// Создаем массив int и waitgroup w.
	ints := []int{2, 4, 6, 8, 10}
	var w sync.WaitGroup
	// Выполняем горутину.
	for _, v := range ints {
		w.Add(1)
		go square(v, &w)
	}
	w.Wait()
}
