package main

import (
	"fmt"
	"math"

	"golang.org/x/exp/utf8string"
)

type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func area(w *Wheel) float64 {
	utf8string.NewString()
	return math.Pi * math.Pow(float64(w.Radius), float64(2))
}

func main() {
	var w Wheel
	w.X = 8
	w.Y = 5
	w.Radius = 3
	w.Spokes = 10

	fmt.Printf("Area of wheel is= %.2f", area(&w))
}
