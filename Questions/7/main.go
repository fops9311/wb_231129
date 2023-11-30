package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	//параллельная запись в мапу приведет к панике, поэтому для конкурентного доступа к мапе нужно пользоваться методами синхронизации
	//такими как mutex
	m := make(map[string]int)
	var mu sync.Mutex
	write := func(key string, val int) {
		mu.Lock()
		defer mu.Unlock()
		m[key] = val
		wg.Done()
	}
	wg.Add(4)
	go write("1", 1)
	go write("2", 2)
	go write("3", 3)
	go write("4", 4)
	wg.Wait()
	//или канал
	c := make(chan KeyVal)
	go func() {
		for v := range c {
			m[v.Key] = v.Val
		}
		wg.Done()
	}()
	wg.Add(1)
	c <- KeyVal{Key: "11", Val: 1}
	c <- KeyVal{Key: "12", Val: 2}
	c <- KeyVal{Key: "13", Val: 3}
	c <- KeyVal{Key: "14", Val: 4}
	close(c)
	wg.Wait()
	fmt.Println(m)
}

type KeyVal struct {
	Key string
	Val int
}
