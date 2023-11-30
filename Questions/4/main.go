package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {
	//количество воркеров задается во флаге
	var workerNum int
	//парсим флаги
	flag.IntVar(&workerNum, "w", 5, "worker count")
	flag.Parse()

	var mainChan chan int = make(chan int)
	//функция, которая управляет записью в канал пока активен контекст
	contextHandlerDeamon := func(ctx context.Context, c chan int) {
		for {
			select {
			case <-ctx.Done(): //получили сигнал прерывания, и закрыли канал
				fmt.Println("\ngot interrupt")
				close(c)
				return
			case <-time.After(100 * time.Millisecond): //подождали 100 милисекунд и записали в канал значение
				c <- int(time.Now().UnixMicro())
			}
		}
	}
	//привязка контекста к сигналу прерывания
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	//освобождает контекст
	defer cancel()
	//запуск горутины
	go contextHandlerDeamon(ctx, mainChan)

	//группа для ожидания заверения всех воркеров
	var wg sync.WaitGroup
	//определение воркера
	work := func(c chan int, workerId int) {
		//пока канал открыт ждем сообщения
		for i := range c {
			//сообщение получено
			fmt.Printf("worker %d recieved %d\n", workerId, i)
			//имитация бурной деятельности
			<-time.NewTimer(time.Millisecond * 2000).C
			//вывод результата работы в консоль
			fmt.Printf("worker %d proc-ed %d\n", workerId, i)
		}
		//после закрытия канала завешаем работу воркера
		fmt.Printf("end of worker %d\n", workerId)
		//уведомляем группу
		wg.Done()
	}
	for i := range make([]interface{}, workerNum) {
		//добавляем в группу
		wg.Add(1)
		//запускаем воркера
		go work(mainChan, i)
	}
	//ждем пока все воркеры не закончат работу
	wg.Wait()
	//завершаем программу
	fmt.Printf("end of main\n")
}
