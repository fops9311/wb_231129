package main

import "fmt"

var set1 []int = []int{4, 2, 6, 5, 3, 1, 3, 4, 5}
var set2 []int = []int{1, 8, 61, 5, 33, 7, 2, 42, 9}

func main() {
	//мапа для хранения уникальных значений
	m := make(map[int]interface{})
	//результирующий слайс
	result := make([]int, 0)
	//канал для передачи уникальных значений
	r := make(chan int)
	go func() {
		//сначала инизиализируем мапу значениями из одного сета
		for _, v := range set1 {
			m[v] = nil
		}
		//потом для всех значений из второго сета, проверяем есть ли такие значения уже в мапе
		for _, v := range set2 {
			if _, ok := m[v]; ok {
				r <- v
			}
		}
		//закрываем канал
		close(r)
	}()
	//читаем все из канала, в нем только пересечения
	for i := range r {
		result = append(result, i)
	}
	fmt.Println(result)
}
