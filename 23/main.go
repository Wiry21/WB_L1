package main

import "fmt"

// Принимает исходный массив и номер удаляемого элемента.
func deleteElemSlice(s []int, n int) []int {
	s = append(s[:n], s[n+1:]...)
	return s
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(deleteElemSlice(s, 1))
}
