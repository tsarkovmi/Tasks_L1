package main

import (
	"fmt"
	"strings"
)

func reverseString(str string) string { //получаем на вход строку и возвращаем строку
	runes := []rune(str)  //делам срез рун из исходной строки, для операций с юникодом
	var b strings.Builder //пользуемся методом Builder, возвращаем структуру содержащую срез байт и указатель на структуру

	for i := len(runes) - 1; i >= 0; i-- { //идём по строке с обратной стороны
		b.WriteRune(runes[i]) //пишем в срез байт сразу же учитывая символы юникода (то есть руны)
	}

	return b.String() //преобразует срез байт в структуре к строке, возвращает строку
}

func main() {
	fmt.Println(reverseString("главрыба"))
}
