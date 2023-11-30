package main

import (
	"fmt"
	"sync"
)

func main() {
	in := []int{2, 4, 6, 8, 10}

	//решение в лоб
	//без использования мьютекса возможна гонка данных,
	//хотя у меня ни разу не получилось увидеть её эффект, вероятно из-за маленького набора данных
	fmt.Println("---------------face on---------------")
	var wg sync.WaitGroup
	var m sync.Mutex
	var sum = 0
	for _, i := range in {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			//вот он, маленький кусочек конкурентного кода
			sq := i * i
			//вот он, маленький кусочек конкурентного кода
			m.Lock()
			defer m.Unlock()
			sum += sq
		}(i)
	}
	wg.Wait()
	fmt.Println("sum:", sum)

	//решение через каналы
	//гонки данных не возникает, т.к. нет параллельного доступа к переменной sum. Она записывается только в одной горутине
	fmt.Println("---------------face off---------------")
	sum = 0
	var square chan int = make(chan int)
	var out chan int = make(chan int)

	var wg1 sync.WaitGroup
	go func() {
		for v := range square {
			out <- v * v
		}
	}()
	go func() {
		for v := range out {
			sum += v
			wg1.Done()
		}
	}()
	for _, i := range in {
		wg1.Add(1)
		square <- i
	}
	wg1.Wait()
	fmt.Println("sum:", sum)
}
