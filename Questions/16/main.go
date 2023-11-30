package main

import (
	"fmt"
	"math/rand"
)

func createHugeRandomSliceInt(n int) []int {
	res := make([]int, n)
	for i := range res {
		res[i] = rand.Intn(10000)
	}
	return res
}
func main() {
	var a []int
	for i := range make([]int, 15) {
		a = createHugeRandomSliceInt(i)
		var b []int = make([]int, len(a))
		copy(b, a)
		qsort(a, 0, len(a)-1)
		if !isSorted(a) {
			debug = true
			qsort(b, 0, len(b)-1)
			fmt.Println("sorted ", isSorted(b), " \t", b)
			return
		}
		fmt.Println("sorted ", isSorted(a), " \t", i)
	}

}
func isSorted(arr []int) bool {
	if len(arr) < 2 {
		return true
	}

	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			return false
		}
	}

	return true
}

var debug = false

func qsort(a []int, low, high int) {

	middle := (high-low)/2 + low
	if debug {
		fmt.Printf("_________________________________\n")
		fmt.Printf("a  : %v,\t low: %d,\t high: %d,\t middle: %d\n", a, low, high, middle)
	}

	if high == low {
		return
	}
	if (high - low) == 1 {
		if a[high] < a[low] {
			a[high], a[low] = a[low], a[high]
		}
		return
	}
	for range a[low : high+1] {
		bigi := getBiggerIndex(a, middle, low)
		smalli := getSmallerIndex(a, middle, high)
		if bigi == middle && smalli == middle {
			if debug {
				fmt.Printf("a r: %v,\t smalli: %d,\t bigi: %d,\t middle: %d\n", a, smalli, bigi, middle)
				fmt.Printf("_________________________________\n")
			}
			qsort(a, low, middle)
			qsort(a, middle, high)
			return
		}
		if bigi == middle {

			if debug {
				fmt.Printf("a -: %v,\t smalli: %d,\t bigi: %d,\t middle: %d\n", a, smalli, bigi, middle)
				//a[middle], a[middle+1] = a[middle+1], a[middle]
			}
			a[smalli], a[middle] = a[middle], a[smalli]
			//middle++

			if debug {
				fmt.Printf("a--: %v,\t smalli: %d,\t bigi: %d,\t middle: %d\n", a, smalli, bigi, middle)
			}
			continue
		}
		if smalli == middle {

			if debug {
				fmt.Printf("a +: %v,\t smalli: %d,\t bigi: %d,\t middle: %d\n", a, smalli, bigi, middle)
			}
			//a[middle], a[middle-1] = a[middle-1], a[middle]
			a[bigi], a[middle] = a[middle], a[bigi]
			//middle--
			if debug {
				fmt.Printf("a++: %v,\t smalli: %d,\t bigi: %d,\t middle: %d\n", a, smalli, bigi, middle)
			}
			continue
		}
		a[bigi], a[smalli] = a[smalli], a[bigi]
		if debug {
			fmt.Printf("a s: %v,\t smalli: %d,\t bigi: %d,\t middle: %d\n", a, smalli, bigi, middle)
		}
	}
}
func getBiggerIndex(a []int, middle int, low int) int {
	v := a[middle]
	for i := range a[low : middle+1] {
		if a[i] > v {
			return i
		}
	}
	return middle
}
func getSmallerIndex(a []int, middle int, high int) int {
	v := a[middle]

	s := a[middle : high+1]
	for i := range s {
		if s[i] < v {
			return i + middle
		}
	}
	return middle
}
