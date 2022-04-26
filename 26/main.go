package main

import (
	"fmt"
	"strings"
)

// Принимает на вход строку s, возвращает true, если уникальные, иначе false.
func checkUnique(s string) bool {
	s = strings.ToLower(s)
	m := make(map[string]int)
	for _, value := range s {
		if _, ok := m[string(value)]; ok {
			return false
		}
		m[string(value)]++
	}
	return true
}

func main() {
	str1 := "abcde"
	str2 := "abCdefAaff"
	str3 := "абвгдее"
	fmt.Println(str1, "-", checkUnique(str1))
	fmt.Println(str2, "-", checkUnique(str2))
	fmt.Println(str3, "-", checkUnique(str3))
}
