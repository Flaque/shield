package tokenizer

type Tokenizer struct {
	toks    []Token
	current string
	recents string
	input   string
}

func NewTokenizer(input string) *Tokenizer {
	return &Tokenizer{
		toks:  []Token{},
		input: input,
	}
}

func (t *Tokenizer) eat() {
	t.current, t.input = string(t.input[0]), string(t.input[1:])
}

func (t *Tokenizer) add(tok Token) {
	t.toks = append(t.toks, tok)
}

func (t *Tokenizer) alpha() {
	if t.recents == "" {
		return
	}

	t.add(keywords(t.recents))
	t.recents = ""
}

func (t *Tokenizer) listen() {

	// check for parens
	if val, ok := delims(t.current); ok {
		t.alpha()
		t.add(val)
		return
	}

	// whitespace
	if IsWhitespace(t.current) || t.input == "" {
		t.alpha()
	} else {
		t.recents += t.current
	}
}

func (t *Tokenizer) Run() []Token {
	for t.input != "" {
		t.eat()
		t.listen()
	}

	if t.toks[len(t.toks)-1].Sym != EOF {
		t.toks = append(t.toks, NewToken(EOF))
	}

	return t.toks
}
