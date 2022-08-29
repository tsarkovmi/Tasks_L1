package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func BinarySearch[T constraints.Ordered](list []T, item T) int {
	var low, mid int //создадим переменные для хранения нижней границы и середины массива
	var guess T
	high := len(list) - 1 //инициализируем переменную для верхней границы массива

	for low <= high { //идём по массиву, пока нижняя граница меньше верхней
		mid = (low + high) / 2 //вычисляем середину не хитрым образом
		guess = list[mid]      //предполагаем, что искомый элемент находится в середение
		if guess == item {     //если искомый элемент совпадает с предполагаемым
			return mid // то он искомый элемент найден, возвращаем индекс в массиве
		}
		if guess > item { //если предполагаемый элемент больше чем искомый
			high = mid - 1 //смещаем верхнюю границу к центру (делим пополам текущий отрезок)
		} else { // в ином случае смещаем нижную границу к центру (делим пополам текущий отрезок)
			low = mid + 1
		}
	}
	return -1 //если элемент не найден - возвращаем -1
}

func main() {
	myList := []int{1, 3, 5, 7, 9, 13, 15, 17, 19, 20, 25, 34, 44, 45, 55, 56, 67, 78, 90, 100}

	fmt.Println(BinarySearch(myList, 9))
	fmt.Println(BinarySearch(myList, -1))

	fmt.Println(BinarySearch(myList, 7))
}
