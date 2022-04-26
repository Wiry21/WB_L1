package main

import (
	"fmt"
	"sync"
)

type countStruct struct {
	mx sync.Mutex
	wg sync.WaitGroup
	i  int
}

func newStruct() countStruct {
	return countStruct{}
}

func count(one *countStruct) {
	one.mx.Lock()
	one.i++
	one.mx.Unlock()
	one.wg.Done()
}

func main() {
	counter := newStruct()
	fmt.Println("Значение счетчика до вычислений", counter.i)

	for i := 0; i < 10; i++ {
		counter.wg.Add(1)
		go count(&counter)
	}
	counter.wg.Wait()
	fmt.Println("Значение счетчика после вычислений", counter.i)
}
