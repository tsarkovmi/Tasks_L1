package main

import (
	"fmt"
)

//получаем срез, и индекс элемента, который нужно удалить
func deleteElem[T any](arr []T, i int) []T {
	return append(arr[:i], arr[i+1:]...) //возвращаем результат работы функции append
	//дописываю к первому аргументу, второй
	//arr[:i] - все элемент до i-го
	//arr[i+1:] - и все элементы после i-го
}

func deleteElemNoOrder[T any](arr []T, i int) []T {
	arr[i] = arr[len(arr)-1] //вместо i-го элемента записываю последний элемент
	arr = arr[:len(arr)-1]   //обрезаю срез, удаляя ранее перенесённый элемент
	return arr               //возвращаю полученный срез
}

func deleteElemGC[T any](arr []*T, i int) []*T {
	if i < len(arr)-1 {
		copy(arr[i:], arr[i+1:]) //на место удаляемого элемента, крепится оставшаяся последовательность
	}
	arr[len(arr)-1] = nil   //лишний указатель удаляется (обнулсяется, что одно и тоже)
	return arr[:len(arr)-1] //вовращаем без обнулённого указателя
}

func main() {
	// удаление элемента с помощью append
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} //инициализируем срез
	fmt.Printf("before slice: %v\n", a)
	a = deleteElem(a, 2) //вызываем функцию удаления 3 го элемента
	fmt.Printf(" after slice: %v\n\n", a)

	// удаление элемента без сохранения порядка
	a = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("before slice: %v\n", a)
	a = deleteElemNoOrder(a, 2)
	fmt.Printf(" after slice: %v\n\n", a)

	// удаление элемента из массива указателей (устанавливаем удаленный элемент в nil)
	a = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := make([]*int, 10)
	for i := range b {
		b[i] = &a[i]
	}
	fmt.Printf("before slice: %v\n", b)
	b = deleteElemGC(b, 2)
	fmt.Printf(" after slice: %v\n\n", b)
}
