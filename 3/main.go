package main

import (
	"fmt"
	"sync"
)

func main() {
	//объявим и инициализируем необходимый массив
	array := []int{2, 4, 6, 8, 10}
	var sum int
	wg := new(sync.WaitGroup) //создадим группу горутин, которые должны выполняться вместе как одна группа
	mu := new(sync.Mutex)     //используем тип мьютекс

	for _, v := range array {
		wg.Add(1)                                              //увеличиваем счётчик горутин в группе
		go func(num int, wg *sync.WaitGroup, mu *sync.Mutex) { //вызываем анонимную функцию в отдельной горутине
			defer wg.Done() // после работы функции используем метод Done, чтобы сигнализировать, что горутина закончила работу
			mu.Lock()       //заблокировали горутину
			sum += num * num
			mu.Unlock() //разблокировали горутину
		}(v, wg, mu)
	}
	wg.Wait() // метод Wait блокирует выполнение функции main до завершения работы всех горутин группы wg
	fmt.Println(sum)
}
