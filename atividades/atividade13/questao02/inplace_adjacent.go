package main

import (
	"fmt"
	"os"
)

func inplaceAdj(str []string) []string {
	if len(str) == 0 {
		return str
	}

	i := 0
	for j := 1; j < len(str); j++ {
		if str[i] != str[j] {
			// anterior e proximo diferentes
			i++
			str[i] = str[j]
		}
	}
	return str[:i+1]
}

func main() {
	str := os.Args

	str = inplaceAdj(str)

	for i := range str {
		fmt.Printf("%s\n", str[i])
	}
}
