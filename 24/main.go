package main

import (
	"fmt"

	point "github.com/tsarkovmi/tasks_l1/24/point"
)

func main() {
	p1 := point.NewPoint(3.4, 7.2) //объявим точки
	p2 := point.NewPoint(1.2, 4.1)

	nP := point.NotPoint{}
	fmt.Printf("distance is %f\n", nP.Distance(p1, p2))
}
