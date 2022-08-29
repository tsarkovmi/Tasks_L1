package point

import "math"

type Point struct { //создали необходимую структуру для хранения точки
	x float64
	y float64
}

type NotPoint struct { //создадим родительскую структуру для инкапсуляции параметров структуруы Поинт
	Point
}

func NewPoint(x float64, y float64) Point { //конструктор
	return Point{
		x: x,
		y: y,
	}
}

func (nP NotPoint) Distance(p1 Point, p2 Point) float64 { //вычисление дистанции по известным формулам
	return math.Sqrt((p2.x-p1.x)*(p2.x-p1.x) + (p2.y-p1.y)*(p2.y-p1.y))
}
