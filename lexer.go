package foo

import (
	"fmt"
	"unicode"
)

var keywords = map[string]Symbol{
	"func": FUNC,
}

// State of the lexer
var tok Symbol
var current rune
var input string
var pos int
var collector string

func peak(str string) rune {
	return rune(str[0])
}

func pop(str string) (rune, string) {
	return rune(str[0]), string(str[1:]) // val, rest of string
}

func setTok(s Symbol) {
	tok = s
	fmt.Println(s)
}

func eat() bool {
	if input == "" {
		return false
	}

	current, input = pop(input)
	return true
}

func labelOrKeyword(value string) {
	if word, ok := keywords[value]; ok {
		setTok(word)
	} else {
		setTok(LABEL)
	}
}

func lparen() {

}

func alpha() {
	if !eat() {
		labelOrKeyword(collector)
		return
	}

	if !unicode.IsLetter(current) {
		labelOrKeyword(collector)
		return
	}

	// collect and continue
	collector += string(current)
	alpha()
}

func eof() bool {
	if input == "" {
		setTok(END)
		return true
	}
	return false
}

func next() {

	// EOF
	if eof() {
		return
	}

	// pop off next char
	ahead := peak(input)

	// Space delim
	if unicode.IsSpace(ahead) {
		next()
		return
	}

	// Alpha chars
	if unicode.IsLetter(ahead) {
		alpha()
		next()
		return
	}

	// Something we didn't recognize
	fmt.Println("Error: Unrecognized symbol")
}

func accept(s Symbol) bool {
	if tok == s {
		next()
		return true
	}
	return false
}
