package main

import (
	"fmt"
)

//функция для описания переменной типа interface, просто печатает содержимое переменной и её тип через запятую
func describe(i interface{}) {
	fmt.Printf("%v, %T\n", i, i)
}

func main() {

	var i interface{} //объявим переменную типа interface
	i = 50            //присвоим ей значение типа int
	describe(i)       //распечатаем описание переменной

	i = "string_type" //присвоим значение типа string
	describe(i)

	i = true //происвоим значение типа bool
	describe(i)

	i = make(<-chan int) //присвоим значение типа channel (в выводе указатель на канал и его тип)
	describe(i)

}
