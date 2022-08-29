package main

import (
	"fmt"
	"sync"
)

//создадим структуру мапы, где дополнительно используется мьютекс
type Counters struct {
	mx sync.RWMutex //RWMutex для блокировки чтения и записи мапы
	m  map[string]int
}

func NewCounters() *Counters { //конструктор для нашей структуры
	return &Counters{
		m: make(map[string]int),
	}
}

func (c *Counters) Get(key string) (int, bool) { // метод для получения значения из мапы
	c.mx.RLock()         //блокируем структуру на запись, однако другие горутины могут получать доступ к чтению мапы
	defer c.mx.RUnlock() //разблокируем структуру после работы метода

	value, ok := c.m[key] //считываем данные из мапы
	return value, ok      //возвращаем значение и успех
}

func (c *Counters) Put(key string, value int) { //метод для записи в мапу некоторых данных
	c.mx.Lock() //Блокируем мапу для других горутин полностью, и на запись и на чтение
	defer c.mx.Unlock()

	c.m[key] = value //записываем в мапу необходимые данные
}

func (c *Counters) Delete(key string) { //метод для удаления записей из мапы
	c.mx.Lock() //Блокируем мапу для других горутин полностью, и на запись и на чтение
	defer c.mx.Unlock()
	delete(c.m, key) //удаляем значение из мапы с помощью функции delete
}

func main() {
	mp := NewCounters() //с помощью конструктора получаем указатель на новую структуру
	mp.Put("one", 1)    //используем метод для вставки в мапу нового значения

	check, ok := mp.Get("one") //проверим получение из мапы данных из мапы
	if !ok {                   //обработаем ошибку
		fmt.Printf("Error get value.\n")
		return
	}
	fmt.Printf("map[one]=%d\n", check) //напечатаем полученное значения

	mp.Delete("one") //попробуем удалить запись из мапы по некоторому ключу

	_, ok = mp.Get("two") //снова считаем значение из мапы
	if !ok {              //обработаем ошибку
		fmt.Printf("Error get value.\n")
		return
	}

}
