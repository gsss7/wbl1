package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) Point {
	return Point{x: x, y: y}
}

func (p Point) Distance(other Point) float64 {
	dx := p.x - other.x
	dy := p.y - other.y
	return math.Sqrt(dx*dx + dy*dy)
}

func (p Point) X() float64 {
	return p.x
}

func (p Point) Y() float64 {
	return p.y
}

func (p Point) Equal(other Point) bool {
	return p.x == other.x && p.y == other.y
}

func main() {
	p1 := NewPoint(1, 2500.32)
	p2 := NewPoint(94334223422.54, 1234.45)
	p3 := NewPoint(0, 0)
	p4 := NewPoint(3, 4)

	dist1 := p1.Distance(p2)
	dist2 := p3.Distance(p4)

	fmt.Printf("Точка 1: %s\n", p1)
	fmt.Printf("Точка 2: %s\n", p2)
	fmt.Printf("Расстояние между точками: %.2f\n\n", dist1)

	fmt.Printf("Точка 3: %s\n", p3)
	fmt.Printf("Точка 4: %s\n", p4)
	fmt.Printf("Расстояние между точками: %.2f\n\n", dist2)

	fmt.Printf("Координаты точки 1: x=%.2f, y=%.2f\n", p1.X(), p1.Y())

	p5 := NewPoint(1, 2500.32)
	fmt.Printf("Точка 1 равна точке 5? %v\n", p1.Equal(p5))
}
