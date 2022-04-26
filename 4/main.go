package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Принимает номер воркера number и канал ch, данные читаются из канала и записываются в data, а затем выводятся.
func work(number int, ch chan string) {
	for {
		data := <-ch
		fmt.Printf("Воркер №%d %s \n", number, data)
	}
}

func main() {
	// Создаем переменную n, в которую запишем число воркеров.
	var n int
	fmt.Print("Введите число воркеров: ")
	_, err := fmt.Scan(&n)
	if err != nil {
		log.Println(err.Error())
	}

	// Создаем канал для получения данных, отправляемых воркерам.
	// Запускаем n воркеров через горутины, передавая номер воркера i и канал ch.
	ch := make(chan string)
	for i := 1; i <= n; i++ {
		go work(i, ch)
	}

	// Создаем канал для получения уведомлений в виде os.Signal.
	// signal.Notify регистрирует данный канал для получения уведомлений об указанных сигналах.
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT)

	func() {
		word := "воркает."
		for {
			select {
			case <-sigs:
				return
			default:
				time.Sleep(time.Second * 2)
				ch <- word
			}
		}
	}()
	close(ch)
	close(sigs)
	fmt.Print("exiting")
}
