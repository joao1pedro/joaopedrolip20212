# Correção: 1,8
# Quase lá, na trave! O problema é a maneira que você definiu lexeme dificulta tratar variáveis com nome de mais de um caractere.
# Veja o front.in que coloqui no seu repositório.

# declaração de variaveis
lexeme = [None]*100
in_fp = open('front.in', 'r')
charClass = None
lexLen = None
nextChar = None
nextToken = None


def getc(in_fp):
    while True:
        char = in_fp.read(1)

        if not char:
            break
        return char

    in_fp.close()


def lookup(ch):
    global nextToken
    if ch == '(':
        addChar()
        nextToken = 25  # LEFT_PAREN
    elif ch == ')':
        addChar()
        nextToken = 26  # RIGHT_PARENT
    elif ch == '+':
        addChar()
        nextToken = 21  # ADD_OP
    elif ch == '-':
        addChar()
        nextToken = 22  # SUB_OP
    elif ch == '*':
        addChar()
        nextToken = 23  # MULT_OP
    elif ch == '/':
        addChar()
        nextToken = 24  # DIV_OP
    else:
        addChar()
        nextToken = -1  # EOF
    return nextToken


def addChar():
    global nextChar
    global lexLen

    if lexLen <= 98:
        lexeme[lexLen+1] = nextChar
        lexeme[lexLen] = 0
    else:
        print("Error - lexeme is too long \n")


def getChar():
    global charClass
    global nextChar
    global in_fp

    nextChar = getc(in_fp)

    if nextChar != -1 and nextChar is not None:
        if nextChar.isalpha():
            charClass = 0  # LETTER
        elif nextChar.isdigit():
            charClass = 1  # DIGIT
        else:
            charClass = 99  # UNKNOW
    else:
        charClass = -1  # EOF


def getNonBlank():
    global nextChar
    if nextChar is not None:
        while nextChar.isspace():
            getChar()


def lex():
    global lexLen
    global charClass
    global nextToken
    global lexeme

    lexLen = 0
    getNonBlank()

    if charClass == 0:
        # LETTER
        addChar()
        getChar()
        while charClass == 0 or charClass == 1:
            addChar()
            getChar()
        nextToken = 11  # INDENT
    elif charClass == 1:
        # DIGIT
        addChar()
        getChar()
        while charClass == 1:
            addChar()
            getChar()
        nextToken = 10  # INT_LIT
    elif charClass == 99:
        # UNKNOWN
        lookup(nextChar)
        getChar()
    elif charClass == -1:
        nextToken = -1
        lexeme[0] = 'E'
        lexeme[1] = 'O'
        lexeme[2] = 'F'
        lexeme[3] = 0
    print(
        f"Next token is: {nextToken}, Next lexeme is {[x for x in lexeme if x is not None]}\n")
    return nextToken


def main():
    global in_fp
    global nextToken
    if in_fp is None:
        print("ERROR - cannot open front.in \n")
    else:
        getChar()
        while True:
            lex()
            if not nextToken != -1:
                break


if __name__ == '__main__':
    main()
