package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func setIntersetcion[T constraints.Ordered](arr1 []T, arr2 []T) []T {
	m := make(map[T]int) //создадим мапу с ключём передаваемого типа

	for _, v := range arr1 {
		m[v] += 1
	}
	for _, v := range arr2 {
		m[v] += 1
	}

	res := make([]T, 0)   //создали пустой срез необходимого типа
	for k, v := range m { //цикл пока есть ключи в мапе
		if v > 1 { //находим пересечение двух множеств
			res = append(res, k) //если оно есть, добавляем в массив число, которое пересекается
		}
	}

	return res //возвращаем массив
}

func main() {
	set1 := []int{1, 2, 3, 4, 5, 10} //первое множетсво, представленно как срез интов
	set2 := []int{6, 4, 3, 8}        //второе множество

	res := setIntersetcion(set1, set2)
	fmt.Printf("intersection: %v\n", res)
}
