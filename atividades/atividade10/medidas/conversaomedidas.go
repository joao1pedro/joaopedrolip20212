// Correção: 1,0
package main

import (
	"fmt"
	"medidas/medconv"
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		m := medconv.Metro(t)
		p := medconv.Pe(t)
		k := medconv.Quilograma(t)
		l := medconv.Libra(t)
		fmt.Printf("%s = %s, %s = %s, %s = %s, %s = %s", m, medconv.MToPe(m), p, medconv.PeToM(p), k, medconv.KgToLb(k), l, medconv.LbToKg(l))
	}
}
