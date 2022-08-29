package main

import (
	"fmt"
	"time"
)

func sleep(d time.Duration) {
	<-time.After(d) //Канал заблокирован, и ждёт необходимое время. Затем возвращает текущее время, и функция завершается
}

func main() {
	start := time.Now()            //запишем время начала работы программы
	sleep(5 * time.Second)         //уснём на 5 секунд
	fmt.Println(time.Since(start)) //напечатаем сколько спали
}
