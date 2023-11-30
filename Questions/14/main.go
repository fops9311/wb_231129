package main

import "fmt"

func main() {
	for _, i := range []interface{}{true, 1, "1", make(chan bool)} {
		//чтобы просто определить тип можно воспользоваться таким методом
		fmt.Printf("type is %T\n", i)
		//а чтобы использовать переменную как конкретный тип нужно использовать type assert
		switch v := i.(type) {
		case bool:
			fmt.Printf("v == true %v\n", v == true)
		case int:
			fmt.Printf("v == 12 %v\n", v == 12)
		case chan bool:
			close(v)
			fmt.Printf("%v closed\n", v)
		case string:
			fmt.Printf("| %s |\n", v)

		default:
			fmt.Printf("dunno what to do")

		}
	}
}
