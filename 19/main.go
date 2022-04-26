package main

import "fmt"

// Итерируемся по строке сначала до конца, к новому символу прибавляем те, что уже прошли.
func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}

func main() {
	str := "главрыба"
	strRev := reverse(str)
	fmt.Println(str)
	fmt.Println(strRev)
}
