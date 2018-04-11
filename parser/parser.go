package parser

import (
	"fmt"

	tok "github.com/flaque/shield/tokenizer"
)

type Parser struct {
	toks    []tok.Token
	current tok.Token
	next    tok.Token
	err     error
}

func NewParser() *Parser {
	return &Parser{
		toks: []tok.Token{},
		err:  nil,
	}
}

// Run cycles through the program and checks for errors
func (p *Parser) Run(toks []tok.Token) error {
	p.toks = toks
	p.program()
	return p.err
}

func (p *Parser) unexpected(t tok.Token) {
	p.err = fmt.Errorf("Unexpected Token %s", t)
}

func (p *Parser) unexpectedEOF() {
	p.err = fmt.Errorf("Unexpected End of File!")
}

func (p *Parser) peak() {
	if len(p.toks) != 0 {
		p.next = p.toks[0]
	} else {
		p.unexpectedEOF()
	}
}

func (p *Parser) eat(s tok.Symbol) {

	if len(p.toks) == 0 {
		p.unexpectedEOF()
		return
	}

	p.current, p.toks = p.toks[0], p.toks[1:]
	p.peak()

	if p.current.Sym != s {
		p.err = fmt.Errorf("Error, Shield expected a %s \n", s)
	}
}

func (p *Parser) program() {
	p.peak()
	for p.next.Sym != tok.EOF && p.err == nil {
		p.code_block()
	}
}

func (p *Parser) code_block() {
	if p.next.Sym == tok.FUNC {
		p.func_def()
	}

	p.statement()
}

func (p *Parser) statement() {
	if p.next.Sym == tok.RETURN {
		p.returner()
		return
	}

	if p.next.Sym == tok.LABEL {
		p.eat(tok.LABEL)

		if p.next.Sym == tok.ASSIGN {
			p.assignment()
			return
		} else {
			p.paren_comma_term()
			return
		}
	}
}

func (p *Parser) func_def() {
	p.eat(tok.FUNC)
	p.eat(tok.LABEL)
	p.paren_index()
	p.eat(tok.DOES)
	p.equals()
	p.eat(tok.BY)
	p.statement()
	p.eat(tok.END)
}

func (p *Parser) paren_comma_term() {
	p.eat(tok.L_PAREN)
	p.comma_term()
	p.eat(tok.R_PAREN)
}

func (p *Parser) equals() {
	p.term()
	p.eat(tok.IS)
	p.term()
}

func (p *Parser) paren_index() {
	p.eat(tok.L_PAREN)
	p.index()
	p.eat(tok.R_PAREN)
}

func (p *Parser) index() {
	p.eat(tok.LABEL)
	if p.next.Sym == tok.COMMA {
		p.eat(tok.COMMA)
		p.index()
	}
}

func (p *Parser) comma_term() {
	p.term()
	if p.next.Sym == tok.COMMA {
		p.eat(tok.COMMA)
		p.comma_term()
	}
}

func (p *Parser) returner() {
	p.eat(tok.RETURN)
	p.term()
}

func (p *Parser) assignment() {
	p.eat(tok.ASSIGN)
	p.term()
}

func (p *Parser) math() {
	p.eat(tok.PLUS) // Check for operators here
	p.term()
}

func (p *Parser) term() {

	// label or func call
	if p.next.Sym == tok.LABEL {
		p.eat(tok.LABEL)

		// function call
		if p.next.Sym == tok.L_PAREN {
			p.paren_comma_term()
		}

		// Math
		if p.next.Sym == tok.PLUS {
			p.math()
		}
		return
	}

	// literal
	p.literal()
}

func (p *Parser) literal() {
	if p.next.Sym == tok.NUM {
		p.eat(tok.NUM)
		return
	}

	p.unexpected(p.next)
}
