// FUNÇÕES NORMAIS DE FUNÇÕES COMO VALORES

package main

import (
	"fmt"
	"math"
)

// Função normal
func soma(x, y float64) float64 {
	return x + y
}

// Retornando listas
func bhaskara(a, b float64, delta float64) [2]float64 {
	return [2]float64{
		(-b + math.Sqrt(delta)/2*a),
		(-b - math.Sqrt(delta)/2*a),
	}
}

// Recursiva
func eleva(num float64, cont int) float64 {
	if cont <= 1 {
		return num
	}
	cont--
	return num * eleva(num, cont)
}

func main() {
	// Função como valor
	var soma2 = func(x, y int) int {
		return x + y
	}
	fmt.Println("5 + 5 =", soma(5, 5))
	fmt.Println("5 + 5 =", soma2(5, 5))

	fmt.Printf("\n---------------\n")

	var a, b, c float64 = 1, 8, -9
	var delta = func(a, b, c float64) float64 {
		return float64(eleva(b, 2) - 4*a*c)
	}
	x := bhaskara(a, b, delta(a, b, c))
	fmt.Println("X'  =", x[0])
	fmt.Println("X\"  =", x[1])
}
