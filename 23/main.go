package main

import "fmt"

func deleteElemSlice(s []int, n int) []int {
	s = append(s[:n], s[n+1:]...)
	return s
}

func main() {
	s1 := []int{1, 2, 3, 4, 5}
	fmt.Println(deleteElemSlice(s1, 1))
}
