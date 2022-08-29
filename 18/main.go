package main

import (
	"fmt"
	"sync"
)

type CounterMutext struct { //создали структуру с счётчиком и мьютексом
	Counter int64
	mx      *sync.RWMutex
}

func NewCounterMutex() *CounterMutext { //конструктор нашей структуры
	return &CounterMutext{
		Counter: 0,
		mx:      new(sync.RWMutex),
	}
}

func (c *CounterMutext) Add() { //операция инкрементирования
	c.mx.Lock()         //Блокируем структуру на запись и на чтение
	defer c.mx.Unlock() //после работы метода разблокируем
	c.Counter++         //инкрементируем
}

func (c *CounterMutext) Get() { //операция получения значения
	c.mx.RLock()                                     //блокируем структуру на запись
	defer c.mx.RUnlock()                             //после работы метода разблокируем
	fmt.Printf("structure counter is:%d", c.Counter) //выводим итоговое значение
}

func (c *CounterMutext) Workers(wg *sync.WaitGroup, CountNumber int) { //пул воркеров
	defer wg.Done() //группа ожидания
	for i := 0; i < CountNumber; i++ {
		c.Add() //инкрементирует значение структуры с помощью метода
	}
}

func main() {
	NumberOfIterations := 1000 //количество итераций для кажого воркера
	NumberOfWorkers := 100     //количество воркеров
	wg := new(sync.WaitGroup)  //объявление новой группы

	counter := NewCounterMutex() //создание экземпляра структуры

	wg.Add(NumberOfWorkers) //добавление в группу ожидания необходимого количества воркеров
	for i := 0; i < NumberOfWorkers; i++ {
		go counter.Workers(wg, NumberOfIterations) //запускаем метод для работы воркеров в горутине
	}

	wg.Wait()     //ожидаем завершения работы каждой из горутин
	counter.Get() //получаем итоговое значние с помощью метода

}
