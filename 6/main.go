package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Принимает канал exit, работает до тех пор, пока данные не придут по каналу.
func classic(exit chan bool) {
	for {
		select {
		case <-exit:
			fmt.Println("Остановка горутины")
			return
		default:
			fmt.Println("Горутина 1 выполняется")
			time.Sleep(time.Second)
		}
	}
}

func main() {
	// Создаем канал exit и запускаем горутину, через 3 секунды передаем данные и горутина останавливается.
	exit := make(chan bool)
	go classic(exit)
	time.Sleep(time.Second * 3)
	exit <- true

	// Создаем WaitGroup wg и канал ch.
	var wg sync.WaitGroup
	ch := make(chan bool)
	// Запускаем горутину, передавая в нее данные через канал ch, затем закроем горутину.
	wg.Add(1)
	go func() {
		defer wg.Done()
		for b := range ch {
			fmt.Printf("Hello %t\n", b)
		}
	}()
	ch <- true
	ch <- true
	close(ch)
	wg.Wait()

	// Создаем канал forever и context ctx, при выполнение функции cancel закроется канал.
	forever := make(chan bool)
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done(): // Если cancel() выполнилось.
				// Посылаем сигнал горутине, которая должна завершиться.
				// Ждем выхода из горутины, затем return.
				forever <- true
				return
			default:
				fmt.Println("for loop")
			}

			time.Sleep(500 * time.Millisecond)
		}
	}(ctx)

	go func() {
		time.Sleep(3 * time.Second)
		cancel()
	}()

	<-forever
	fmt.Println("finish")
}
