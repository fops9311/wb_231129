package main

import (
	"fmt"
	"reflect"
	"strings"
)

func IsUniqueOnly(s string) bool {
	s = strings.ToUpper(s)
	chars := []rune(s)
	m := make(map[rune]struct{})
	for _, c := range chars {
		if _, ok := m[c]; ok {
			return false
		}
		m[c] = struct{}{}
	}
	return true
}

func main() {
	fmt.Println(IsUniqueOnly("abcd"))
	fmt.Println(IsUniqueOnly("abCdefAaf"))
	fmt.Println(IsUniqueOnly("aabcd"))
	fmt.Println(IsUniqueOnly("Aabcd"))
	fmt.Println(reflect.TypeOf(struct{}{}).Size())
}
