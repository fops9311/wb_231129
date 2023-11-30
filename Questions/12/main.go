package main

import "fmt"

type Set[T comparable] struct {
	s []T
}

func NewSet[T comparable](s []T) Set[T] {
	res := Set[T]{}
	m := make(map[T]interface{})
	for _, v := range s {
		m[v] = nil
	}

	for k := range m {
		res.s = append(res.s, k)
	}
	return res
}
func (s Set[comparable]) Len() int {
	return len(s.s)
}

// множества равны если они содержат одинаковые элементы
func (s1 Set[comparable]) Equal(s2 Set[comparable]) bool {
	if s1.Len() != s2.Len() {
		return false
	}
	s3 := s1.Intersect(s2)
	if s3.Len() != s1.Len() {
		return false
	}
	return true
}
func (s1 Set[comparable]) Intersect(s2 Set[comparable]) (s3 Set[comparable]) {
	set1 := s1.s
	set2 := s2.s
	//мапа для хранения уникальных значений
	m := make(map[comparable]interface{})
	//результирующий слайс
	result := make([]comparable, 0)
	//канал для передачи уникальных значений
	r := make(chan comparable)
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
	s3.s = result
	return s3
}

func main() {
	//создали три множества на основе слайсов
	s1 := (NewSet([]string{"cat", "cat", "dog", "cat", "tree"}))
	s2 := (NewSet([]string{"dog", "cat", "tree"}))
	s3 := (NewSet([]string{"cat", "cat", "dog", "cat", "tree", "opossum"}))
	fmt.Println("s1", s1)
	fmt.Println("s2", s2)
	fmt.Println("s3", s3)
	//сравнили
	fmt.Println("s1.Equal(s2)", s1.Equal(s2))
	fmt.Println("s2.Equal(s3)", s2.Equal(s3))
	fmt.Println("s1.Equal(s3)", s1.Equal(s3))
}
