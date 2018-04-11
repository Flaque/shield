package tokenizer

import "strconv"

type Symbol int

const (
	FUNC Symbol = iota
	LABEL
	DOES
	BY
	COLON
	DONE
	L_PAREN
	R_PAREN
	L_SQUIG
	R_SQUIG
	PLUS
	MINUS
	EQUALS
	ASSIGN
	RETURN
	IS
	NUM
	COMMA
	END
	EOF
)

var keys = map[string]Symbol{
	"func":   FUNC,
	"does":   DOES,
	"by":     BY,
	"done":   DONE,
	"return": RETURN,
	"is":     IS,
	"end":    END,
}

var dels = map[string]Symbol{
	"{": L_SQUIG,
	"}": R_SQUIG,
	"(": L_PAREN,
	")": R_PAREN,
	"+": PLUS,
	"-": MINUS,
	",": COMMA,
}

type Token struct {
	Sym   Symbol
	Value string
}

func NewToken(s Symbol) Token {
	return Token{
		Sym:   s,
		Value: "",
	}
}

func (t Token) String() string {
	if t.Value != "" {
		return t.Sym.String() + "(" + t.Value + ")"
	} else {
		return t.Sym.String()
	}
}

func keywords(str string) Token {
	val, ok := keys[str]
	if ok {
		return Token{
			Value: "",
			Sym:   val,
		}
	}

	// Num
	if _, err := strconv.ParseInt(str, 10, 64); err == nil {
		return Token{
			Value: string(str),
			Sym:   NUM,
		}
	}

	// Label
	return Token{
		Value: string(str),
		Sym:   LABEL,
	}
}

func delims(str string) (Token, bool) {
	val, ok := dels[str]

	return Token{
		Value: "",
		Sym:   val,
	}, ok
}
