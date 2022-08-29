package main

import "fmt"

type Human struct { //объявили нашу родительскую структуру
	Name       string
	SecondName string
	Age        int
	Action     // которая содержит две дочерние
	Act        Action
}

func (h *Human) SetName(name string) { //несколько методов для родительской структуры
	h.Name = name
}

func (h *Human) SetSecondName(SecondName string) {
	h.SecondName = SecondName
}

func (h *Human) SetAge(Age int) {
	h.Age = Age
}

type Action struct { //объявили дочернию структуру
	Name     string
	Walk     bool
	Work     bool
	Location string
}

func (a *Action) SetWalk(Walk bool) { //также несколько методов для дочерней структуры
	a.Walk = Walk
}

func (a *Action) SetWork(Work bool) {
	a.Work = Work
}

func (a *Action) SetLocation(Location string) {
	a.Location = Location
}

func (a *Action) SetName(Name string) {
	a.Name = Name
}

func main() {
	a := Human{ //проинициализируем родительcкую структуру
		Name:       "Mark",
		SecondName: "Zuckerberg",
		Age:        19,
		Action: Action{
			Name:     "Eliot",
			Walk:     true,
			Work:     false,
			Location: "USA",
		},
		Act: Action{
			Walk:     false,
			Work:     false,
			Location: "New-York",
		},
	}

	//обратимся к полю Walk дочерней структуры явно указывая поле дочерней структуры при обращении
	fmt.Println(a.Action.Walk) //true

	//Так как поле уникальное мы можем обратиться к нему не указывая дочернию структуру
	fmt.Println(a.Walk) //true

	//Если поле не уникальное, приоритет за родительским полем
	fmt.Println(a.Name) //Mark

	//Чтобы обратиться к одноименному полю дочерней структуры нужно указать необходимое поле при обращении
	fmt.Println(a.Action.Name) //Eliot

	/*
		Методы
	*/

	//обратиться к методу дочерней структуры можно двумя способами. Явно указывая, что обращаемся к методу дочерней структуры
	a.Action.SetWalk(false)
	//Или сразу же вызывая метод дочерней структуры. Это возможно, потому что методы наследуются родительской структурой
	a.SetWalk(true)
	fmt.Println(a.Walk) //true

	a.SetLocation("Russia")
	a.Act.SetLocation("USA")
	fmt.Printf("%v, %v\n", a.Location, a.Act.Location) // Russia, USA

	//Если существует два одинаковых метода у дочерней и у родительской структур, то необходимо обращаться к методу подробно

	//в таком случае, имя присвоится родительской структуре, так как она выше уровнем
	a.SetName("Bob")
	fmt.Println(a.Name)        //Bob
	fmt.Println(a.Action.Name) //Eliot
	/*
		чтобы присвоить имя дочерней структуре, или структуре ниже уровнем,
		необходимо указать поле или тип структуры, для которого реализован необходимый метод
	*/
	a.Action.SetName("Jhon")
	fmt.Println(a.Name)        //Bob
	fmt.Println(a.Action.Name) //Jhon
}
