package main

import (
	"fmt"
	"math/big"
)

func main() {
	//воспользуемся библиотекой big и создадим указатель на число
	a, ok := big.NewInt(0).SetString("5456513215645649854321321564878413215648649", 10) //передадим наше число как строку, аргумент 10 означает СС
	if !ok {                                                                            //проверим на ошибку
		fmt.Println("не удалось создать первое число")
		return
	}
	b, ok := big.NewInt(0).SetString("4587132156798711316877984521256498789978441", 10)
	if !ok {
		fmt.Println("не удалось создать второе число")
		return
	}

	value := big.NewInt(0)
	fmt.Println("*:", value.Mul(a, b)) //воспользуемся существующими методами
	fmt.Println("/:", value.Div(a, b))
	fmt.Println("+:", value.Add(a, b))
	fmt.Println("-:", value.Sub(a, b))
}
