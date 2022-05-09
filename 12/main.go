package main

import "fmt"

// Принимает на вход исходный массив и возвращает множество.
// Если ключ(значение массива) не равен уже существующему значению во множестве, тогда добавляе.
// Если равен, то проверяем следующий элемент.
func removeDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if _, ok := allKeys[item]; !ok {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	res := removeDuplicateStr(arr)
	fmt.Println(res)
}
