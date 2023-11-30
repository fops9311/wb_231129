package main

import "fmt"

func main() {
	var (
		in          int64
		mask        int64
		setBit      int64
		clearBit    int64
		setBitNum   int64 = 3
		clearBitNum int64 = 2
	)

	fmt.Printf("-----------setBit------------\n")
	in = int64(101)
	mask = int64(1 << setBitNum)
	setBit = in | mask
	fmt.Printf("in %b\nmask %b\nsetBit %b\n", in, mask, setBit)
	fmt.Printf("----------clearBit-----------\n")
	in = int64(101)
	mask = int64(1 << clearBitNum)
	clearBit = in &^ mask
	fmt.Printf("in %b\nmask %b\nclearBit %b\n", in, mask, clearBit)

	fmt.Printf("------------func-------------\n")
	fmt.Println("AssignBit(in, setBitNum, true):", AssignBit(in, setBitNum, true))
	fmt.Println("AssignBit(in, clearBitNum, false):", AssignBit(in, clearBitNum, false))
	fmt.Println("AssignBit(in, 66, true):", AssignBit(in, 66, true))
	fmt.Println("AssignBit(in, 66, false):", AssignBit(in, 66, false))
	fmt.Println("AssignBit(in, 0, false):", AssignBit(in, 0, false))
	fmt.Println("AssignBit(in, -1, false):", AssignBit(in, -1, false))
}

// функция которая устанавливает определенный бит в int64 в заданное значение
//
// в случае выхода за пределы n ничего не происходит
func AssignBit(i int64, n int64, v bool) int64 {
	if n < 0 {
		return i
	}
	//создаем маску в определенном бите
	mask := int64(1 << n)
	if v {
		//если нужно установить побитовое или
		return i | mask
	} else {
		//если нужно сбросить побитовое и с инвертированной маской
		return i &^ mask
	}
}
