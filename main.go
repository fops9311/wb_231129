package main

import (
	"fmt"
	"os"
	"reflect"
	"sync"
	"time"
)

// устные вопросы
func main() {
	//L1----------------------------------------------------------------
	fmt.Println("|Вопрос 05| Какой размер у структуры struct{}{}?")
	fmt.Println("struct{}{} size:\t\t", reflect.TypeOf(struct{}{}).Size())
	fmt.Println("struct{bool} size:\t\t", reflect.TypeOf(Test{}).Size())
	fmt.Println("struct{interface{}} size:\t", reflect.TypeOf(Test1{}).Size())
	fmt.Println("вывод: интерфейс это не ничто, а вполне даже много чего")
	fmt.Println("")
	fmt.Println("|Вопрос 06| Есть ли в Go перегрузка методов или операторов?")
	fmt.Println("ответ: нет")
	fmt.Println("")
	fmt.Println("|Вопрос 07| В какой последовательности будут выведены элементы map[int]int?")
	fmt.Println("ответ: при итерации по мапе ключи выдаются в случайном порядке")
	fmt.Println("")
	fmt.Println("|Вопрос 08| В чем разница make и new?")
	fmt.Println("ответ: new - выделяет и обнуляет память (возвращает указатель), но не инизиализирует. make - инициализирует и служит только для инициализации каналов срезов и мап (не возвращает указатель) стр 168")
	fmt.Println("")
	fmt.Println("|Вопрос 09|Сколько существует способов задать переменную типа slice или map?")
	fmt.Println("ответ: литералы", map[string]int{"r": 1})
	fmt.Println("ответ: make", make(map[string]int))
	fmt.Println("ответ: new", func() interface{} {
		m := *(new(map[string]int))
		go func() {
			<-time.NewTimer(time.Second).C
			m["TEST1"] = 1
			fmt.Println("ответ: new", m)
		}()
		m = map[string]int{}
		m["TEST"] = 1
		return m
	}(),
	)
	<-time.NewTimer(time.Second * 2).C
	fmt.Println("")
	fmt.Println(`|Вопрос 10| Что выведет данная программа и почему?


	func update(p *int) {
	  b := 2
	  p = &b
	}
	
	func main() {
	  var (
		 a = 1
		 p = &a
	  )
	  fmt.Println(*p)
	  update(p)
	  fmt.Println(*p)
	}
	`)
	fmt.Println("ответ: p = &b не передает новое значение в исходную ссылку, а присваевает переменной с типом ссыки новое значение, поэтому вывод 1 и 1")
	fmt.Println("ответ: *p = b передаёт")
	main1()
	fmt.Println("")
	fmt.Println(`|Вопрос 11| Что выведет данная программа и почему?

	func main() {
		wg := sync.WaitGroup{}
		for i := 0; i < 5; i++ {
		   wg.Add(1)
		   go func(wg sync.WaitGroup, i int) {
			  fmt.Println(i)
			  wg.Done()
		   }(wg, i)
		}
		wg.Wait()
		fmt.Println("exit")
	  }
	`)
	fmt.Println("ответ: deadlock потому что группу надо передавать по ссылке")
	main3()
	fmt.Println("")
	fmt.Println(`|Вопрос 12| Что выведет данная программа и почему?
	func main() {
		n := 0
		if true {
		   n := 1
		   n++
		}
		fmt.Println(n)
	  }	  
	`)
	fmt.Println("ответ: 0 потому что n в новом скоупе")
	main4()
	fmt.Println("")
	fmt.Println(`|Вопрос 13| Что выведет данная программа и почему?
	func someAction(v []int8, b int8) {
		v[0] = 100
		v = append(v, b)
	  }
	  
	  func main() {
		var a = []int8{1, 2, 3, 4, 5}
		someAction(a, 6)
		fmt.Println(a)
	  }	  
	`)
	fmt.Println("ответ: 100, 2, 3, 4, 5 потому что при аппенде создается новый массив на основе прежнего")
	main5()
	fmt.Println("")
	fmt.Println(`|Вопрос 14| Что выведет данная программа и почему?
	func main() {
		slice := []string{"a", "a"}
	  
		func(slice []string) {
		   slice = append(slice, "a")
		   slice[0] = "b"
		   slice[1] = "b"
		   fmt.Print(slice)
		}(slice)
		fmt.Print(slice)
	  }
	`)
	fmt.Println("ответ: b b a \n a a")
	main6()
	//L2-----------------------------------------------------
	fmt.Println()
	fmt.Println("L2-----------------------------------")
	main8()
}

func Foo() error {
	var err *os.PathError = nil
	return err
}

func main8() {
	err := Foo()
	fmt.Println(err)
	var i interface{}
	i = nil
	fmt.Println(&err, &os.PathError{} == nil, i == nil)
}

func test() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return
}

func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
}

func main7() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}

type Test struct {
	//i interface{}
	b bool
}
type Test1 struct {
	//i interface{}
	b interface{}
}

func update(p *int) {
	b := 2
	p = &b
}
func update1(p *int) {
	b := 2
	*p = b
}

type intp struct {
	v int
}

func update2(p *intp) {
	b := 2
	p.v = b
}

func main1() {
	var (
		a  = 1
		p  = &a
		pp = &intp{}
	)
	fmt.Println(p, *p)
	update(p)
	fmt.Println(p, *p)
	update1(p)
	fmt.Println(p, *p)
	update2(pp)
	fmt.Println(pp, *p)
}
func main3() {
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, i int) {
			fmt.Println(i)
			wg.Done()
		}(&wg, i)
	}
	wg.Wait()
	fmt.Println("exit")
}
func main4() {
	n := 0
	if true {
		n := 1
		n++
	}
	fmt.Println(n)
}
func someAction(v []int8, b int8) {
	v[0] = 100
	v = append(v, b)
}

func main5() {
	var a = []int8{1, 2, 3, 4, 5}
	someAction(a, 6)
	fmt.Println(a)
}
func main6() {
	slice := []string{"a", "a"}

	func(slice []string) {
		slice = append(slice, "a")
		slice[0] = "b"
		slice[1] = "b"
		fmt.Print(slice)
	}(slice)
	fmt.Print(slice)
}
