package main

import (
	"fmt"
	"strings"
)

func isUnique(str string) bool {
	str = strings.ToLower(str)   //переведём все символы в нижний регистр
	m := make(map[rune]struct{}) //создадим пустую мапу

	for _, r := range str {
		if _, ok := m[r]; ok { //ищем записи в мапе по ключу r
			return false //если такие записи есть, то строка не уникальна
		}
		m[r] = struct{}{} //в ином случае, запишем по новому ключу пустую структуру
	}

	return true
}

func main() {
	fmt.Printf("abcd - %v\n", isUnique("abcd"))
	fmt.Printf("abCdefAaf - %v\n", isUnique("abCdefAaf"))
	fmt.Printf("aabcd - %v\n", isUnique("aabcd"))
}
