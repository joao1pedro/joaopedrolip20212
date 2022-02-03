// Correção: 0,8. Contou "quatro", "Quatro" e "quatro;" como palavras diferentes.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)

	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)

	for in.Scan() {
		counts[in.Text()]++
	}

	err := in.Err()
	if err != nil {
		fmt.Fprintf(os.Stderr, "wordfreq: %v\n", err)
		os.Exit(1)
	}

	for key, value := range counts {
		fmt.Println(key, value)
	}
}
