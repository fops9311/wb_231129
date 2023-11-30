package main

import (
	"fmt"
	"math/rand"
	"runtime"
)

var justString string

// проблема была в том, что мы сохраняем ссылку на исходных массив,
// который хранится в памяти всегда, пока мы не переопределим justString
// и GC только тогда сможет освободить память под изначальную строку "v"
func someFunc() {
	v := createHugeString(1 << 20)

	//вот это бургер
	//он берет слайс [:100] преобразует его в массив [100]byte, потом всё это дело преобразует обратно в слайс
	// и потом в строку
	//вся соль здесь - *(*
	var a [100]byte = (*(*[100]byte)([]byte(v[:100])))
	justString = string(a[:])
	//justString = v[:100]

	PrintMemUsage()
}
func createHugeString(n int) string {
	const letters = "fghwefcqexdcsgdfcsfdbs"
	res := make([]byte, n)
	for i := range res {
		res[i] = letters[rand.Intn(len(letters))]
	}
	return string(res)
}
func main() {

	PrintMemUsage()
	//теперь после вызова этой функции GC должен собрать v := createHugeString(1 << 20) вот эту большую строку
	someFunc()
	PrintMemUsage()
	fmt.Println(justString)
}

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
	runtime.GC()
}
func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
