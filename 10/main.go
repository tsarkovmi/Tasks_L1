package main

import (
	"fmt"
)

func main() {
	array := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5} //создали заданный массив

	group := make(map[int][]float64) // создали мапу для хранения температур по группам
	for _, v := range array {        //цикл по массиву
		key := int(v/10) * 10              //создаём ключ для подгруппы исходя из температуры (приводим float к типу int и округляем)
		group[key] = append(group[key], v) //в мапу с необходимым ключом записываем нашу температуру
	}

	fmt.Printf("temperatures in the groups: %v\n", group) //выводим полученные результаты
}
