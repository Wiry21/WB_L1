package main

import "fmt"

func search(arr []int, target int) int {
	var l int
	r := len(arr) - 1

	for l <= r {
		mid := (l + r) / 2

		switch {
		case arr[mid] == target:
			return mid
		case arr[mid] > target:
			r = mid - 1
		case arr[mid] < target:
			l = mid + 1
		}
	}
	return -1
}

func main() {
	slice := []int{-1000, -8, -5, 1, 21, 38, 45}
	i := search(slice, 1)
	fmt.Println(i)
}
