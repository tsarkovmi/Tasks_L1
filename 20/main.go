package main

import (
	"fmt"
	"strings"
)

func reversedWords(str string) string {
	arr := strings.Split(str, " ") //воспользуемся методом Сплит, чтобы разбить строку на отдельные слова идущие через пробел

	var b strings.Builder //создадим структуру для строки
	b.WriteString(arr[len(arr)-1])

	for i := len(arr) - 2; i >= 0; i-- { //цикл по длинне строки-2, т.к. последний элемент уже записан, идём с обратной стороны
		b.WriteString(" ")    //между элементами ставим пробел
		b.WriteString(arr[i]) //пишем следующее слово

	}

	return b.String()
}

func main() {
	fmt.Printf("reversed string: %s\n", reversedWords("snow dog sun"))
}
