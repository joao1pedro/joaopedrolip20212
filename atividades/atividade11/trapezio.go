package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func f(x float64) (result float64) {
	result = math.Sqrt(1 + math.Sin(x)*math.Sin(x))
	return result
}

func main() {
	var sum float64
	sum = 0

	a, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println(err)
	}
	b, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil {
		fmt.Println(err)
	}
	n, err := strconv.ParseInt(os.Args[3], 10, 64)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Limite inferior: ", a)
	fmt.Println("Limite superior: ", b)

	b = (5*a)/2 - b
	h := (b - a) / float64(n)

	// fmt.Println("Valor de h: ", h)

	for i := 1; i < int(n); i++ {
		sum = sum + (f(a)+f(a+h))*h/2
		a = a + h
	}

	fmt.Println("Número de partições: ", n)
	fmt.Println("Valor da integral: ", sum)
}
