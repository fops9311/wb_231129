package main

import (
	"fmt"
	"sync"
)

func main() {

	in := []int{2, 4, 6, 8, 10}
	//можно запустить выполнение функции, которая будет умножать и выводить в консоль в горутине
	//порядок вывода при этом будет случайным
	//для ограничения количество горутин, можно воспользоваться паттерном worker, или периодически вызывать wg.Wait() в цикле инициализации гоурутин
	fmt.Println("---------------random---------------")
	var wg sync.WaitGroup
	for _, i := range in {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println(i * i)
		}(i)
	}
	wg.Wait()

	//а можно разделить функцию на две, одна будет умножать
	//а вторая выводить на экран, образуя таким образом конвейер
	//при этом выполнение тоже происходит конкурентно, но порядок вывода не нарушается
	fmt.Println("---------------in order---------------")
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
			fmt.Println(v)
			wg1.Done()
		}
	}()
	for _, i := range in {
		wg1.Add(1)
		square <- i
	}
	wg1.Wait()
}
