package main

import "fmt"

// Что выведет код и почему?
type impl struct{}

type I interface {
	C()
}

func (*impl) C() {}

func A() I {
	return nil

}
func B() I {
	var ret *impl
	return ret
}

func main() {
	a := A()
	b := B()
	fmt.Println(a == b)
}

/* Данный код выведет false.
Так получится потому что функция А - действительно nil интерфейс(и тип и значение = nil),
а вот функция В - возвращает переменную ret, которая будет иметь тип *impl и значение nil. Поэтому они не равны.
Чтоб сделать вывод задачи true, надо просто сделать возвращаемое значение функции В - nil
*/
