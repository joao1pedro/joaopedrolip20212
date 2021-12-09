// Correção: 0,4
// Faça suas funções recebendo parâmetros
package main

import (
	"fmt"
	"math"
)

var pi float64 = 22 / 7

var raio float64 = 1.984

func calcularComprimento() float64 {
	var comprimento float64 = 2 * pi * raio

	return comprimento
}

func calcularArea() float64 {
	var area float64 = pi * math.Pow(2, raio)

	return area
}

func main() {
	var comprimento float64 = calcularComprimento()
	var area float64 = calcularArea()

	fmt.Println("Comprimento: ", comprimento)
	fmt.Println("Área: ", area)
}
