package main

import (
	"flag"
	"fmt"
	"sync"
	"time"
)

func main() {
	//количество
	var seconds int
	//парсим флаги
	flag.IntVar(&seconds, "w", 5, "seconds to run")
	flag.Parse()
	//запускаем таймер на N секунд
	t := time.NewTimer(time.Second * time.Duration(seconds))
	//канал в который всё будет писаться
	c := make(chan int)

	var wg sync.WaitGroup
	wg.Add(1)
	//читатель
	go func() {
		//читаем пока канал открыт
		for i := range c {
			fmt.Println(i)
		}
		fmt.Println("DONE READING")
		//уведомляем группу что закончили
		wg.Done()
	}()
	//писатель
	func() {
		for {
			select {
			case <-t.C:
				//закрывает канал и завершаем писать
				close(c)
				return
			default:
				//если таймер не истек то пишем в канал
				c <- int(time.Now().UnixNano())
			}
		}
	}()
	// ждем пока читатель закончит
	wg.Wait()
	fmt.Println("END")
}
