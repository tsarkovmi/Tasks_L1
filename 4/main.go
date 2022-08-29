/*
Реализовать постоянную запись данных в канал (главный поток). Реализовать
набор из N воркеров, которые читают произвольные данные из канала и
выводят в stdout. Необходима возможность выбора количества воркеров при
старте.
Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать
способ завершения работы всех воркеров.
*/

package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func WorkersPool(ctx context.Context, wg *sync.WaitGroup, number int, in <-chan int) {
	defer wg.Done() //после работы горутины, даём знать группе, что работа оконченна

	for {
		select {
		case n := <-in: //получаем данные из канала
			fmt.Printf("worker number %d got: %d\n", number, n)
		case <-ctx.Done(): //ожидаем сигнала из контекста
			fmt.Printf("exiting from worker №%d\n", number)
			return
		}

	}

}

func main() {

	rand.Seed(time.Now().UnixNano()) //Задаём случайные данные для генератора случайный чисел
	var NumberOfWorkers int          //создаем переменную, которая будет хранить количество воркеров

	wg := new(sync.WaitGroup) //создадим группу горутин

	fmt.Print("Please enter the number of workers no more than 1000:")

	_, err := fmt.Scan(&NumberOfWorkers) //считываем количество необходимых воркеров
	//и проверям на ошибку (делаю колличество воркеров не более 1000)
	if err != nil || NumberOfWorkers > 1000 {
		fmt.Printf("Input Error %v. Or working Limit Exceeded ", err.Error())
	}

	channel := make(chan int, NumberOfWorkers) //создаём канал с буфером равным количеству воркеров
	defer close(channel)                       //закрываем каналы, когда работа программы завершается

	ctx, cancel := context.WithCancel(context.Background()) //создаю новый контекст, с помощью которого буду закрывать воркеры

	wg.Add(NumberOfWorkers) // добавляю в группу необходимое количество горутин
	//запускаю в работу воркеров
	for i := 0; i < NumberOfWorkers; i++ {
		go WorkersPool(ctx, wg, i, channel)
	}

	wg.Add(1) //добавляю в группу ещё одну горутину
	//анонимная функция, которая генерирует и записывает случайные данные в канал
	go func() {
		//wg.Add(1)
		defer wg.Done()
		tick := time.NewTicker(time.Millisecond * 400) //таймер для отсрочки отправки данных в канал
		for {
			select { //конструкция select-case для работы с каналами
			case <-tick.C: // С-канал, который возвращает значение через заданное время
				a := rand.Int() //генерируем случайное число
				channel <- a    //записываем число в обрабатываемый воркерами канал
				fmt.Printf("writer sent: %d\n", a)
			case <-ctx.Done(): //прекращение работы конструкции после получаения сигнала из Контекста
				fmt.Println("exiting from writer")
				return

			}
		}
	}()

	gracefulShutdown := make(chan os.Signal, 1) //создадим канал для перехватывания сигнало завершени работы

	//SIGTERM является общим сигналом завершения программы,
	//SIGINT передается, когда пользователь вводит сигнал прерывания (например, Ctrl+C)
	//SIGINT аналогичен SIGTERM, но предназначен для пользовательских событий
	signal.Notify(gracefulShutdown, syscall.SIGINT, syscall.SIGTERM)

	<-gracefulShutdown //слушаем канал, пока не получим необходимый сигнал
	fmt.Println("interrupt signal")
	cancel()  // Вызов cancel() позволяет освободить ресурсы, связанные с определенным контекстом. (Завершаются все воркеры)
	wg.Wait() //ожидаем завершения всех вызванных горутин

}
