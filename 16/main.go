package main

import "fmt"

func quickSort(arr []int) []int {
	if len(arr) < 2 { // в случае, если массив меньше двух, то есть 1 элемент, сортировать нечего, возвращаем массив
		return arr
	}

	pivot := arr[0]                         //опорный элемент
	less := make([]int, 0, 1+len(arr)/2)    //создадим отрезок для левой части
	greater := make([]int, 0, 1+len(arr)/2) //создадим отрезок для правой части

	for _, val := range arr[1:] { //идём по массиву
		if val <= pivot { //если текущий элемент меньше либо равен нулевому
			less = append(less, val) //то он записывается в левый отрезок (который меньший)
		} else { //в ином случае записываем в правый отрезок
			greater = append(greater, val)
		}
	}

	arr = append([]int{}, quickSort(less)...) //запишем в чистый массив arr рекурсивно вызванную функцию quickSort (Сортируем левую часть)
	arr = append(arr, pivot)                  //к отсортированной последовательности допишем наш опорный элемент
	arr = append(arr, quickSort(greater)...)  //к полученной последовательности допишем отсортированную правую часть среза
	return arr
}

func main() {
	fmt.Println(quickSort([]int{1, 3, 7, 5, 1, 4, 3, 2, 1, 9}))
}
