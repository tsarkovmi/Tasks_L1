package main

import "fmt"

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"} //инициализировали необходимый массив

	m := make(map[string]struct{}) //создали пустую мапу
	res := make([]string, 0)       //создали срез строк

	for _, v := range arr { //идём по срезу строк, читаем в v, значения
		if _, ok := m[v]; ok { //если в мапе, по данному ключу уже есть данные, то пропускаем
			continue
		}
		res = append(res, v) //если нет, то делаем запись в срез строк
		m[v] = struct{}{}    //делаем метку, что ключ добавлен
	}

	fmt.Printf("subset: %v\n", res)
}
