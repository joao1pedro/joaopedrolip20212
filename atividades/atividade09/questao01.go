package main

import (
	"fmt"
	"math"
)

var pi float64 = 22 / 7

var raio = 1.984

func main() {
	fmt.Println("Comprimento: ", 2*pi*raio)     // comprimento
	fmt.Println("Área: ", pi*math.Pow(2, raio)) // área
}
