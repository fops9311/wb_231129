package main

import (
	"fmt"
	"sync"
)

type ConcurrentCounter struct {
	v    int
	m    sync.Mutex
	Hook func()
}

func (c *ConcurrentCounter) Inc() {
	c.m.Lock()
	defer c.m.Unlock()
	c.v++
	c.Hook()
}
func (c *ConcurrentCounter) Get() int {
	//применение мьютекса этом месте не обеспечивает синхронный вывод
	//потому что вывод может залочиться до инкримента
	//c.m.Lock()
	//defer c.m.Unlock()
	return c.v
}
func main() {
	//группа не обязательна, она здесь просто для демонстрации синхронного вывода, как и хук
	var wg sync.WaitGroup
	counter := &ConcurrentCounter{
		Hook: func() {
			wg.Done()
		},
	}
	//делаем 500 инкриметов
	for range make([]any, 500) {
		wg.Add(1)
		go counter.Inc()
		//если раскоментировать будет синхронный вывод
		//wg.Wait()
		fmt.Println(counter.Get())
	}
	wg.Wait()
	fmt.Println(counter.Get())
}
