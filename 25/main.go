package main

import (
	"fmt"
	"time"
)

func Sleep(i int) {
	<-time.After(time.Duration(i) * time.Second)
}

func main() {
	i := 3
	fmt.Println("Hello")
	Sleep(i)
	fmt.Println("world")
}
