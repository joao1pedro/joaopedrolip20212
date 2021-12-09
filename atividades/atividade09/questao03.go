// Correção: 0,8
// Mesmo problema da primeira.
package main

import (
	"flag"
	"fmt"
	"math"
)

var pi float64 = 22 / 7

var raio float64

func calcularComprimento(r float64) float64 {
	var comprimento float64 = 2 * pi * r

	return comprimento
}

func calcularArea(r float64) float64 {
	var area float64 = pi * math.Pow(2, raio)

	return area
}

func main() {
	var calcComprimento bool
	var calcArea bool

	flag.Float64Var(&raio, "r", 1.984, "defina o raio")
	flag.BoolVar(&calcComprimento, "c", false, "calcular comprimento?")
	flag.BoolVar(&calcArea, "a", false, "calcular comprimento?")

	flag.Parse()

	if calcComprimento == true {
		var comprimento float64 = calcularComprimento(raio)
		fmt.Println("Comprimento: ", comprimento)
	}

	if calcArea == true {
		var area float64 = calcularArea(raio)
		fmt.Println("Área: ", area)
	}

}
