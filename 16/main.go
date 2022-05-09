package main

import (
	"fmt"
	"math/rand"
)

// Рандомно выбираем элемент и помещаем его в конец массива. Сравниваем весь массив с последним элементом,
// тем самым разделив массив на 2 части, одна меньше последнего элемента, другая больше.
func quicksort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivot := rand.Int() % len(a)

	a[pivot], a[right] = a[right], a[pivot]

	for i := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]

			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	quicksort(a[:left])
	quicksort(a[left+1:])

	return a
}

func main() {
	slice := []int{21, 45, 38, -8, -1000, 1, -5}
	fmt.Println("Unsorted ", slice)
	quicksort(slice)
	fmt.Println("Sorted ", slice)
}
