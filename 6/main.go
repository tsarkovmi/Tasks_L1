package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//пишем случайные значения в канал
func Writer(ctx context.Context, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	tick := time.NewTicker(time.Millisecond * 700) //таймер для отсрочки отправки данных в канал
	for {
		select { //конструкция select-case для работы с каналами
		case <-tick.C: // С-канал, который возвращает значение через заданное время
			a := rand.Int() //генерируем случайное число
			ch <- a         //записываем число в обрабатываемый воркерами канал
			fmt.Printf("writer sent: %d\n", a)
		case <-ctx.Done(): //прекращение работы конструкции после получаения сигнала из Контекста
			fmt.Println("exiting from writer")
			tick.Stop() //прекращаем отсчет тиков
			return

		}
	}

}

//прекращение работы горутины по сигналу канала
func GoroutineWithCloseChannel(wg *sync.WaitGroup, ch <-chan int) {
	defer wg.Done()
	for {
		data, ok := <-ch //считываем данные из канала в переменную и проверяем на успех
		if !ok {         //если считать данные из канала не удалось, печатаем,  что канал закрыт
			fmt.Printf("channel close. goroutine close.\n")
			return
		}
		fmt.Printf("Goroutine got: %d\n", data) //в ином случае печатаем полученные данные
	}
}

func GoroutineWithRangeChannel(wg *sync.WaitGroup, ch <-chan int) {
	defer wg.Done()

	for in := range ch { //цикл работает до тех пор, пока канал открыт и пишутся данные
		fmt.Printf("Goroutine got: %d\n", in)
	}
	fmt.Println("Goroutine range close.") //как только канал закрывается цикл прекращает работу
}

func GoroutineWithStopChannel(wg *sync.WaitGroup, ch <-chan int, stop <-chan struct{}) {
	defer wg.Done()

	for {
		select { //используем дополнительный канал для закрытия горутины
		case in := <-ch:
			fmt.Printf("Goroutine got: %d\n", in)
		case <-stop: //как только этот канал что то передаст, горутина закроется
			fmt.Println("goroutine got stop signal. Goroutine close")
			return
		}
	}
}

func GoroutineWithContext(wg *sync.WaitGroup, ch <-chan int, ctx context.Context) {
	defer wg.Done()

	for {
		select { //ждём сигнал контекста
		case in := <-ch:
			fmt.Printf("Goroutine got: %d\n", in)
		case <-ctx.Done():
			fmt.Println("goroutine got context signal. Goroutine close")
			return
		}
	}
}

func GoroutineWithContextTimeout(wg *sync.WaitGroup, ch <-chan int, ctx context.Context) {
	defer wg.Done()

	for {
		select { //аналогично ждём сигнал контекста, такая же конструкция как и выше, только контекст срабатывает сам, по таймауту
		case in := <-ch:
			fmt.Printf("Goroutine got: %d\n", in)
		case <-ctx.Done():
			fmt.Println("goroutine got timeout context signal. Goroutine close")
			return
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	wg := new(sync.WaitGroup)
	ctx, cancel := context.WithCancel(context.Background()) //создаём новые контекст для подачи сигнала отмены

	channel := make(chan int) //создаём новый канал

	//Запустим первую горутину
	wg.Add(2)                                 //добавим в группу ожидания две горутины
	go Writer(ctx, channel, wg)               //горутина, которая пишет в канал произвольные данные
	go GoroutineWithCloseChannel(wg, channel) //горутина, которая их считывает, закрывается сразу, как закрывается канал

	time.Sleep(2 * time.Second) //таймеры для отсрочки закрытия горутины
	cancel()                    //подаем сигнал контексту (в данном случае закрывается только пишущая горутина)
	time.Sleep(time.Second)     //таймер для синхронизации
	close(channel)              //закрываем канал (в данном случае канал закрывается и закрывается читающая горутина)
	wg.Wait()                   //ожидаем выполнения всех деферов

	//Запуск новой горутины с закрытием в цикле
	fmt.Println("\nStart new Goroutine With Range")
	ctx, cancel = context.WithCancel(context.Background()) //обновляем наш контекст
	channel = make(chan int)                               //и канал
	wg.Add(2)
	go Writer(ctx, channel, wg)
	go GoroutineWithRangeChannel(wg, channel) //у этой горутины почти такой же смысл, как и у предыдущей, она закрывается вместе с каналом

	time.Sleep(2 * time.Second)
	cancel()
	time.Sleep(time.Second)
	close(channel)

	wg.Wait()

	//Запуск новой горутины с закрытием по сигналу канала
	fmt.Println("\nStart new Goroutine With Stop signal")
	ctx, cancel = context.WithCancel(context.Background())
	channel = make(chan int)
	stop := make(chan struct{}) //канал для подачи сигнала закрытия горутины
	wg.Add(2)
	go Writer(ctx, channel, wg)
	go GoroutineWithStopChannel(wg, channel, stop) //горутина закрывается, как только в канал stop подаются данные

	time.Sleep(2 * time.Second)
	cancel() //закрываем writer
	time.Sleep(time.Second)
	stop <- struct{}{} //подаём сигнал стоп, горутина останавливается
	close(channel)     //закрываем открытые каналы
	close(stop)

	wg.Wait()

	//Запуск новой горутины с закрытием по сигналу контекста
	fmt.Println("\nStart new Goroutine With Ctx signal")
	ctx, cancel = context.WithCancel(context.Background())
	channel = make(chan int)
	wg.Add(2)
	go Writer(ctx, channel, wg)
	go GoroutineWithContext(wg, channel, ctx) //горутина закрывается по сигналу контекста, так же как горутина writer

	time.Sleep(2 * time.Second)
	cancel()
	time.Sleep(time.Second)
	close(channel)

	wg.Wait()

	//Запуск новой горутины с закрытие по сигналу контекста с таймаутом
	fmt.Println("\nStart new Goroutine With Ctx Timeout signal")
	ctx, cancel = context.WithCancel(context.Background())
	timeout, _ := context.WithTimeout(context.Background(), time.Second*3) //делаем новый контекст, которые подаёт сигнал по с течением времени
	channel = make(chan int)
	wg.Add(2)
	go Writer(ctx, channel, wg)
	go GoroutineWithContextTimeout(wg, channel, timeout) //передаём в горутину контекст таймаута, закрывается после 3 секунд работы

	time.Sleep(2 * time.Second)
	cancel()
	time.Sleep(time.Second)
	close(channel)

	wg.Wait()
}
