package main

import (
	"fmt"
)

type figures int

const (
	square   figures = iota // квадрат
	circle                  // круг
	triangle                // равносторонний треугольник
)

func area(fig figures) (func(float64) float64, bool) {

	switch fig {
	case square:
		{
			return func(x float64) float64 { return float64(x * x) }, true
		}
	case circle:
		{
			return func(x float64) float64 { return float64(3.14 * x * x) }, true
		}
	case triangle:
		{
			return func(x float64) float64 { return float64(x * x * 0.433gi) }, true
		}
	default:
		return nil, false
	}
}

func main() {
	var myFigure figures
	myFigure = 4

	ar, ok := area(myFigure)
	if !ok {
		fmt.Println("Ошибка")
		return
	}
	x := float64(10)
	myArea := ar(x)
	fmt.Println(myArea)
}
