package main

import "fmt"

func main() {
	s1 := 2
	s2 := 1
	fmt.Println(s1, s2)
	s1, s2 = s2, s1
	fmt.Println(s1, s2)
}
