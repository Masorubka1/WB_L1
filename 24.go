package main

import (
	"fmt"
	"math"
)

// Структура Point представляет точку с координатами x и y
type Point struct {
	x, y float64
}

// Конструктор для создания новой точки
func NewPoint(x, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

// Метод для вычисления расстояния между двумя точками
func (p *Point) DistanceTo(other *Point) float64 {
	dx := p.x - other.x
	dy := p.y - other.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main24() {
	// Создаем две точки
	point1 := NewPoint(1.0, 2.0)
	point2 := NewPoint(4.0, 6.0)

	// Вычисляем расстояние между точками
	distance := point1.DistanceTo(point2)

	fmt.Printf("Расстояние между точками: %.2f\n", distance)
}
