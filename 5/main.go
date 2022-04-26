package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	// Создаем переменную n, в которую запишем время работы в секундах.
	var n int
	fmt.Print("Введите время работы в секундах: ")
	_, err := fmt.Scan(&n)
	if err != nil {
		log.Println(err.Error())
	}

	// Создаем канал timer, в котором запускаем "таймер" на n секунд.
	timer := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * time.Duration(n))
		timer <- "exiting"
	}()

	// Создаем канал ch, в котором раз в секунду будем писать данные в канал.
	ch := make(chan int, 1)
	go func() {
		var i int
		for {
			i++
			time.Sleep(time.Second)
			ch <- i
		}
	}()

	// Запускаем вечный цикл, пока не придут данные по каналу timer.
	func() {
		for {
			select {
			case <-timer:
				return
			default:
				value := <-ch
				fmt.Println(value)
			}
		}
	}()
	close(timer)
	close(ch)
	fmt.Print("exiting")
}
