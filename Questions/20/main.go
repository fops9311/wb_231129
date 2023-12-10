package main

import (
	"fmt"
	"strings"
)

var input string = "snow dog1 dog2 sun"
var input1 string = "snow dog1 sun"

func ReverseWords(s string) string {
	words := strings.Split(s, " ")

	for i := 0; i < len(words)/2; i++ {
		words[i], words[len(words)-i-1] = words[len(words)-i-1], words[i]
	}
	return strings.Join(words, " ")
}

func main() {
	fmt.Println("OK:\t", input, "|", ReverseWords(input))
	fmt.Println("OK:\t", input1, "|", ReverseWords(input1))
}
