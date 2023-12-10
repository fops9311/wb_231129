package main

import "fmt"

var input string = "!\"№%!\"№главрыба mainfish"

// правильная реализация
// если необходимо работать с отдельным символом, нужно использовать преобразование в rune
func ReverseString(s string) string {
	chars := []rune(s) //type rune = int32
	for i := 0; i < len(chars)/2; i++ {
		chars[i], chars[len(chars)-i-1] = chars[len(chars)-i-1], chars[i]
	}
	return string(chars)
}
func ReverseStringWrong(s string) string {
	b := []byte(s) //type byte = uint8
	for i := 0; i < len(b)/2; i++ {
		b[i], b[len(b)-i-1] = b[len(b)-i-1], b[i]
	}
	return string(b)
}
func main() {
	fmt.Println("OK:\t", input, ReverseString(input))
	fmt.Println("NotOK:\t", input, ReverseStringWrong(input))
}
