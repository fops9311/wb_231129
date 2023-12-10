package main

import "fmt"

func DeleteN[T any](s []T, n int) []T {
	if n < 0 {
		return s
	}
	if n >= len(s) {
		return s
	}
	res := make([]T, len(s)-1)
	offset := 0
	for i := range s {
		if i == n {
			offset--
			continue
		}
		res[i+offset] = s[i]
	}
	return res
}

func main() {
	fmt.Println(DeleteN([]int{1, 2, 3, 4}, 2))
	fmt.Println(DeleteN([]int{1, 2, 3, 4}, -2))
	fmt.Println(DeleteN([]int{1, 2, 3, 4}, 0))
	fmt.Println(DeleteN([]int{}, 0))
	fmt.Println(DeleteN([]int{1, 2, 3, 4}, 500))
}
