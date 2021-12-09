// Correção: 0,3
package main

import (
	"fmt"
	"math"
)
// Aqui, mesmo pi sendo float64, 22 / 7 será definido como inteiro, para depois ser atribuída a variável.
var pi float64 = 22 / 7

var raio = 1.984

func main() {
	fmt.Println("Comprimento: ", 2*pi*raio)     // comprimento
	fmt.Println("Área: ", pi*math.Pow(2, raio)) // área
}
