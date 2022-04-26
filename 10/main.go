package main

import (
	"fmt"
	"math"
)

// Принимает на вход исходный массив, возвращает map с ключами ввиде названия подмножества и
// значениями ввиде данных подмножеств.
func group(arr []float64) map[int][]float64 {
	m := make(map[int][]float64)
	// Из каждого значения массива берём ключ с разницей в 10 и по этому ключу добавляем значение.
	for _, value := range arr {
		if math.Abs(value) > 9 {
			group := int(value/10) * 10
			m[group] = append(m[group], value)
		} else {
			if value >= 0 {
				m[0] = append(m[0], value)
			} else {
				m[-1] = append(m[-1], value)
			}
		}
	}
	return m
}

func main() {
	arr := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, 3.1, -2.9}
	fmt.Println(group(arr))

}
