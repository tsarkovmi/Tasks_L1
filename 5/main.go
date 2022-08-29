package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func Worker(wg *sync.WaitGroup, in <-chan int) {
	defer wg.Done() //после работы горутины, даём знать группе, что работа оконченна

	for i := range in { //цикл работает пока канал активен
		fmt.Printf("worker got:%d\n", i)
	}
	fmt.Println("worker shuts down")

}

func main() {
	rand.Seed(time.Now().UnixNano()) //рандомизируем генерируемые значения
	wg := new(sync.WaitGroup)        //создадим группу горутин
	chanel := make(chan int)         //создадим канал для обмена данными горутин

	wg.Add(2)             //добавим в группу две горутины
	go Worker(wg, chanel) //запустим первую горутину

	go func() { //запустим вторую горутину
		defer wg.Done()
		tick := time.NewTicker(time.Millisecond * 200) //создадим тикер для отправки данных с определённой переодичностью
		defer tick.Stop()
		timer := time.NewTimer(time.Second * 5) //создадим таймер работы этой горутины
		defer timer.Stop()

		for {
			select {
			case <-tick.C: //С-канал, который возвращает значение через заданное время
				in := rand.Int()
				chanel <- in
				fmt.Printf("writer send %d\n", in)
			case <-timer.C: //С-канал, который возвращает значение через заданное время
				fmt.Println("writer stopped. time is over")
				defer close(chanel) //закрываем канал, тем самым закрываем воркер
				return
			}
		}
	}()

	wg.Wait()

}
