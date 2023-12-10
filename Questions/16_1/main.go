package main

import (
	"fmt"
	"math/rand"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func createHugeRandomSliceInt(n int) []int {
	res := make([]int, n)
	for i := range res {
		res[i] = rand.Intn(10000)
	}
	return res
}

func main() {
	pts := make(plotter.XYs, 0)
	pts1 := make(plotter.XYs, 0)
	times := 0
	var total time.Duration
	var total2 time.Duration
	for i := 140000; i < 150000; i = i + 10 {
		times++
		test := []int{}
		test1 := createHugeRandomSliceInt(i) //[]int{1000, 4, 300, 0, 1, 1, 2, 1, 0}
		var t = time.Time{}

		copy(test, test1)
		var leftIds []int
		var rightIds []int
		copy(leftIds, test1)
		copy(rightIds, test1)
		t = time.Now()
		qsort(test, leftIds, rightIds)
		d1 := time.Now().Sub(t)
		total += (d1)
		if !isSorted(test) {
			return
		}
		//fmt.Println("test:", "size:", len(test), "time:", d1)

		copy(test, createHugeRandomSliceInt(i))
		t = time.Now()
		quickSort(test, 0, len(test)-1)
		d2 := time.Now().Sub(t)
		if !isSorted(test) {
			return
		}

		total2 += (d2)
		//fmt.Println("test:", "size:", len(test), "time:", d2)
		//fmt.Println("isSorted:", isSorted(test))

		//fmt.Printf("mine is faster x%.2f times\n", float64(d2)/float64(d1))
		//fmt.Println(i)
		//pts = append(pts, plotter.XY{X: float64(i), Y: float64(d2) / float64(d1)})
		pts1 = append(pts1, plotter.XY{X: float64(i), Y: 1})
	}
	fmt.Println(total / time.Duration(times))
	fmt.Println(total2 / time.Duration(times))
	p := plot.New()

	p.Title.Text = "Plotutil example"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"
	err := plotutil.AddLinePoints(p, "compare", pts, "line", pts1)
	if err != nil {
		panic(err)
	}

	// Save the plot to a PNG file.
	if err := p.Save(40*vg.Inch, 40*vg.Inch, "points.png"); err != nil {
		panic(err)
	}
}

// мой улучшенный алгоритм квиксорт
func qsort(s []int, leftIds []int, rightIds []int) {
	//массив из нуля и 1 элемента уже отсортирован
	if len(s) < 2 {
		return
	}
	//нахождение среднего индекса массива
	middleId := len(s) / 2
	//функции сравнения (замыкание)
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
	//поиск значений больше середины слева
	Find(s, leftIds, 0, middleId, bigger)
	//поиск значений меньше середины справа
	Find(s, rightIds, middleId, len(s)-1, smaller)
	//замены значений
	middleId = Swap(s, leftIds, rightIds, middleId)
	//рекурсивный вызов для двух оставшихся половин

	if middleId < len(s) && middleId > 0 {
		qsort(s[0:middleId], leftIds, rightIds)
		qsort(s[middleId+1:], leftIds, rightIds)
	}
	return
}

func Find[T any](s []T, res []int, low int, high int, f func(T) bool) {
	//res := make([]int, 0)
	for i := range s[low : high+1] {
		id := i + low
		if f(s[id]) {
			res[i] = id
		}
	}
	res = res[:len(s[low:high+1])]
	return
}
func Swap[T any](s []T, leftIds []int, rightIds []int, middleId int) int {
	if (len(leftIds) == 0) && (len(rightIds) == 0) {
		return middleId
	}
	if len(leftIds) == len(rightIds) {
		for i := range leftIds {
			s[leftIds[i]], s[rightIds[i]] = s[rightIds[i]], s[leftIds[i]]
		}
		return middleId
	}
	if len(leftIds) < len(rightIds) {
		for i := range leftIds {
			s[leftIds[i]], s[rightIds[i]] = s[rightIds[i]], s[leftIds[i]]
		}
		leftovers := rightIds[len(leftIds):]
		return Reflect(s, leftovers, middleId)
	}
	if len(rightIds) < len(leftIds) {
		for i := range rightIds {
			s[leftIds[i]], s[rightIds[i]] = s[rightIds[i]], s[leftIds[i]]
		}
		leftovers := leftIds[len(rightIds):]

		return Reflect(s, leftovers, middleId)
	}
	return middleId
}
func Reflect[T any](s []T, Ids []int, middleId int) int {
	if len(Ids) == 0 {
		return middleId
	}
	lefty := middleId > Ids[0]
	iter := func(i int) {
		if middleId == i+1 {
			s[middleId], s[i] = s[i], s[middleId]
			middleId--
			return
		}
		if middleId == i-1 {
			s[middleId], s[i] = s[i], s[middleId]
			middleId++
			return
		}
		if i > middleId {
			s[middleId], s[middleId+1] = s[middleId+1], s[middleId]

			s[middleId], s[i] = s[i], s[middleId]
			middleId++
			return
		}
		if i < middleId {
			s[middleId], s[middleId-1] = s[middleId-1], s[middleId]

			s[middleId], s[i] = s[i], s[middleId]
			middleId--
			return
		}
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

// первая ссылка гугла. реализация в лоб
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

// выдано YandexGPT2 :)
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
