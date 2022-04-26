package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) *Point {
	newStr := &Point{
		x: x,
		y: y,
	}
	return newStr
}

func GetDistance(f, s *Point) float64 {
	return math.Sqrt((s.x-f.x)*(s.x-f.x) + (s.y-f.y)*(s.y-f.y))
}

func main() {
	a := NewPoint(2, 2)
	b := NewPoint(5, 6)
	fmt.Println("Расстояние между точками  равно:", GetDistance(a, b))
}
