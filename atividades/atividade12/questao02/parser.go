// Correção: 0,5. Mesmo problema do anterior.
package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var nextToken int
var lexLen int
var filename string = "front.in"

var lexeme [100]string
var nextChar string
var token int
var charClass int

var digitCheck = regexp.MustCompile(`^[0-9]+$`)
var IsLetter = regexp.MustCompile(`^[a-zA-Z]+$`)

const (
	// CHARACTER CLASSES
	LETTER  = 0
	DIGIT   = 1
	UNKNOWN = 99
	// TOKEN CODES
	INT_LIT     = 10
	IDENT       = 11
	ASSIGN_OP   = 20
	ADD_OP      = 21
	SUB_OP      = 22
	MULT_OP     = 23
	DIV_OP      = 24
	LEFT_PAREN  = 25
	RIGHT_PAREN = 26
	//
	EOF = -1
)

func getc(filename string) string {
	var ch string
	filebuffer, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	inputdata := string(filebuffer)
	data := bufio.NewScanner(strings.NewReader(inputdata))
	data.Split(bufio.ScanRunes)

	for data.Scan() {
		ch = data.Text()
	}
	return ch
}

func lookup(ch string) int {
	switch ch {
	case "(":
		addChar()
		nextToken = LEFT_PAREN
		break
	case ")":
		addChar()
		nextToken = RIGHT_PAREN
		break
	case "+":
		addChar()
		nextToken = ADD_OP
		break
	case "-":
		addChar()
		nextToken = SUB_OP
		break
	case "*":
		addChar()
		nextToken = MULT_OP
		break
	case "/":
		addChar()
		nextToken = DIV_OP
		break
	default:
		addChar()
		nextToken = EOF
		break
	}
	return nextToken
}

func addChar() {
	if lexLen <= 98 {
		aux := lexLen
		lexLen++
		lexeme[lexLen] = nextChar
		lexeme[aux] = strconv.Itoa(0)
	} else {
		fmt.Print("Error - lexeme is too long \n")
	}
}

func getChar() {
	nextChar = getc(filename)

	if nextChar != strconv.Itoa(EOF) {
		if IsLetter.MatchString(nextChar) {
			charClass = LETTER
		} else if digitCheck.MatchString(nextChar) {
			charClass = DIGIT
		} else {
			charClass = UNKNOWN
		}

	} else {
		charClass = EOF
	}
}

func getNonBlank() {
	for digitCheck.MatchString(nextChar) {
		getChar()
	}

}

func lex() int {
	lexLen = 0
	getNonBlank()

	switch charClass {
	/* Parse identifiers */
	case LETTER:
		addChar()
		getChar()
		for ok := true; ok; ok = charClass == LETTER || charClass == DIGIT {
			addChar()
			getChar()
		}
		nextToken = IDENT
		break
	/* Parse integer literals */
	case DIGIT:
		addChar()
		getChar()
		for ok := true; ok; ok = charClass == DIGIT {
			addChar()
			getChar()
		}
		nextToken = INT_LIT
		break
	/* Parentheses and operators */
	case UNKNOWN:
		lookup(nextChar)
		getChar()
		break
	/* EOF */
	case EOF:
		nextToken = EOF
		lexeme[0] = "E"
		lexeme[1] = "O"
		lexeme[2] = "F"
		lexeme[3] = strconv.Itoa(0)
		break
	}
	/* End of switch */
	fmt.Printf("Next token is: %d, Next lexeme is %s\n", nextToken, lexeme)
	return nextToken
}

func factor() {
	fmt.Print("Enter <factor>\n")

	if nextToken == IDENT || nextToken == INT_LIT {
		lex()
	} else {
		if nextToken == LEFT_PAREN {
			lex()
			expr()
			if nextToken == RIGHT_PAREN {
				lex()
			} else {
				error()
			}
		} else {
			error()
		}
	}
	fmt.Print("Exit <factor>\n")
}

func term() {
	fmt.Print("Enter <term>\n")
	factor()

	for ok := true; ok; ok = nextToken == MULT_OP || nextToken == DIV_OP {
		lex()
		factor()
	}
	fmt.Print("Exit <term>\n")
}

func expr() {
	fmt.Print("Enter <expr>\n")

	term()

	for ok := true; ok; ok = nextToken == ADD_OP || nextToken == SUB_OP {
		lex()
		term()
	}

	fmt.Print("Exit <expr>\n")

}

func error() {
	fmt.Println("Error!")
}

func main() {
	if len(filename) == 0 {
		fmt.Println("ERROR - cannot open front.in")
	} else {
		getChar()
		for ok := true; ok; ok = nextToken != EOF {
			lex()
			expr()
		}
	}
}
