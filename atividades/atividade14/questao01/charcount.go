package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)
	var utflen [utf8.UTFMax + 1]int
	invalid := 0

	in := bufio.NewReader(os.Stdin)

	var isLetter, isDigit, isNumber, isSymbol, isSpace, isPunct int

	for {
		r, n, err := in.ReadRune()

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}

		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}

		counts[r]++
		utflen[n]++

		if unicode.IsLetter(r) {
			isLetter++
		}

		if unicode.IsDigit(r) {
			isDigit++
		}

		if unicode.IsNumber(r) {
			isNumber++
		}
		if unicode.IsSymbol(r) {
			isSymbol++
		}
		if unicode.IsSpace(r) {
			isSpace++
		}
		if unicode.IsPunct(r) {
			isPunct++
		}
	}

	fmt.Printf(("tipo\tcount\n"))
	if isLetter > 0 {
		fmt.Printf("Letter\t%d\n", isLetter)
	}
	if isDigit > 0 {
		fmt.Printf("Digit\t%d\n", isDigit)
	}
	if isNumber > 0 {
		fmt.Printf("Number\t%d\n", isNumber)
	}
	if isSymbol > 0 {
		fmt.Printf("Symbol\t%d\n", isSymbol)
	}
	if isSpace > 0 {
		fmt.Printf("Space\t%d\n", isSpace)
	}
	if isPunct > 0 {
		fmt.Printf("Punct\t%d\n", isPunct)
	}

	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Printf("len\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}

	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
