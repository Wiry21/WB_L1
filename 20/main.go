package main

import (
	"fmt"
	"strings"
)

// Проходим по массиву слов с конца.
func main() {
	str := "snow dog sun"
	tempstr := strings.Split(str, " ")
	arr := make([]string, 0)
	for i := len(tempstr) - 1; i >= 0; i-- {
		arr = append(arr, tempstr[i])

	}
	fmt.Println("Строка до переворота: ", str)
	fmt.Println("Строка после переворота: ", strings.Join(arr, " "))
}
