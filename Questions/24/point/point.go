package point

import "math"

// тип точка с инкапсулированными параметрами
type Point struct {
	x float64
	y float64
}

// функция для определения расстояния между двумя точками
func (p Point) DistanceTo(pt Point) float64 {
	return math.Sqrt(math.Pow(p.x-pt.x, 2) + math.Pow(p.y-pt.y, 2))
}

// конструктор точки с инкапсулированными параметрами
func NewPoint(x float64, y float64) Point {
	return Point{
		x: x,
		y: y,
	}
}
