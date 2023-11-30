package main

import (
	"fmt"
	"math/rand"
	"time"
)

func createHugeRandomSliceInt(n int) []int {
	res := make([]int, n)
	for i := range res {
		res[i] = rand.Intn(10000)
	}
	return res
}

func main() {

	test1 := createHugeRandomSliceInt(300000) //[]int{1000, 4, 300, 0, 1, 1, 2, 1, 0}
	test := test1
	//9846 5799 7751 4324 3312
	//test := []int{9846, 5799, 7751, 4324, 3312}
	t := time.Now()
	qsort(test)
	d1 := time.Now().Sub(t)
	fmt.Println("test:", "size:", len(test), "time:", d1)
	fmt.Println("isSorted:", isSorted(test))

	test = test1
	t = time.Now()
	quickSort(test, 0, len(test)-1)
	d2 := time.Now().Sub(t)
	fmt.Println("test:", "size:", len(test), "time:", d2)
	fmt.Println("isSorted:", isSorted(test))

	fmt.Printf("mine is faster x%.2f times\n", float64(d2)/float64(d1))
}

func quickSort(arr []int, low int, high int) {
	if low < high {
		pi := partion(arr, low, high)

		// Recursively sort elements before partition and after partition
		quickSort(arr, low, pi-1)
		quickSort(arr, pi+1, high)
	}
}

func partion(arr []int, low int, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++

			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
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

var itrecount = 0

func qsort(s []int) int {

	if len(s) < 2 {
		return 0
	}
	itrecount++
	middleId := len(s) / 2

	//fmt.Println(itrecount, "|iteration len:", len(s), "slice:", s, "middle:", middleId)
	smaller := func(i int) bool {
		if i < s[middleId] {
			return true
		}
		return false
	}
	bigger := func(i int) bool {
		if i > s[middleId] {
			return true
		}
		return false
	}

	leftIds := Find(s, 0, middleId, bigger)
	rightIds := Find(s, middleId, len(s)-1, smaller)
	//fmt.Println(itrecount, "|left ids:", leftIds, "right ids:", rightIds)
	middleId = Swap(s, leftIds, rightIds, middleId)

	//fmt.Println(itrecount, "|iteration end len:", len(s), "slice:", s, "middle:", middleId)

	qsort(s[0:middleId])
	if middleId < len(s) {
		qsort(s[middleId+1:])
	}
	return middleId
}

func Find(s []int, low int, high int, f func(int) bool) []int {
	res := make([]int, 0)
	for i := range s[low : high+1] {
		id := i + low
		if f(s[id]) {
			res = append(res, id)
		}
	}
	return res
}
func Swap(s []int, leftIds []int, rightIds []int, middleId int) int {
	//fmt.Print("Swap")
	if (len(leftIds) == 0) && (len(rightIds) == 0) {
		//fmt.Println("0")
		return middleId
	}

	//fmt.Println("s:", s)
	//fmt.Println("e")
	if len(leftIds) == len(rightIds) {
		for i := range leftIds {
			//fmt.Printf("s[%d], s[%d] = s[%d], s[%d]\n", leftIds[i], rightIds[i], rightIds[i], leftIds[i])
			s[leftIds[i]], s[rightIds[i]] = s[rightIds[i]], s[leftIds[i]]

			//fmt.Println("s:", s)
		}
		return middleId
	}
	//fmt.Println("l")
	if len(leftIds) < len(rightIds) {
		for i := range leftIds {
			s[leftIds[i]], s[rightIds[i]] = s[rightIds[i]], s[leftIds[i]]
		}
		leftovers := rightIds[len(leftIds):]
		//fmt.Println("right leftovers:", leftovers)

		return Reflect(s, leftovers, middleId)
	}
	//fmt.Println("r")
	if len(rightIds) < len(leftIds) {
		for i := range rightIds {
			s[leftIds[i]], s[rightIds[i]] = s[rightIds[i]], s[leftIds[i]]
		}
		leftovers := leftIds[len(rightIds):]
		//fmt.Println("left leftovers:", leftovers, "s:", s, "middle:", middleId)

		return Reflect(s, leftovers, middleId)
	}
	return middleId
}
func Reflect(s []int, Ids []int, middleId int) int {
	if len(Ids) == 0 {
		return middleId
	}
	lefty := middleId > Ids[0]
	iter := func(i int) {
		//fmt.Println("Reflect iter start:", s, "middle:", middleId, "id:", i)
		if middleId == i+1 {
			s[middleId], s[i] = s[i], s[middleId]
			middleId--
			//fmt.Println("1Reflect iter end:", s, "middle:", middleId, "id:", i)
			return
		}
		if middleId == i-1 {
			s[middleId], s[i] = s[i], s[middleId]
			middleId++
			//fmt.Println("2Reflect iter end:", s, "middle:", middleId, "id:", i)
			return
		}
		if i > middleId {
			s[middleId], s[middleId+1] = s[middleId+1], s[middleId]

			s[middleId], s[i] = s[i], s[middleId]
			middleId++
			//fmt.Println("3Reflect iter end:", s, "middle:", middleId, "id:", i)
			return
		}
		if i < middleId {
			s[middleId], s[middleId-1] = s[middleId-1], s[middleId]

			s[middleId], s[i] = s[i], s[middleId]
			middleId--
			//fmt.Println("4Reflect iter end:", s, "middle:", middleId, "id:", i)
			return
		}
		//fmt.Println("5Reflect iter end:")

	}
	if lefty {
		for i := len(Ids) - 1; i >= 0; i-- {
			v := Ids[i]
			iter(v)
		}
	} else {
		for _, i := range Ids {
			iter(i)
		}
	}
	return middleId
}
