package main

import "fmt"

// Принимает два массива и возвращает массив.
// Сначала из более короткого массива делает map - элемент = true.
// Затем сравниваем второй массив с map, если элемент есть, то добавляем в массив c.
func Intersection(a, b []int) (c []int) {
	m := make(map[int]bool)

	if len(a) > len(b) {
		a, b = b, a
	}

	for _, item := range a {
		m[item] = true
	}

	for _, item := range b {
		if _, ok := m[item]; ok {
			c = append(c, item)
		}
	}
	return
}

func main() {
	a := []int{4, 6, 5, 1, 3, 2}
	b := []int{5, 6, 4}
	result := Intersection(a, b)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(result)
}
