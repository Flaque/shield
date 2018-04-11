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

func (p *Parser) unexpected(t tok.Token) {
	p.err = fmt.Errorf("Unexpected Token %s", t)
}

func (p *Parser) peak() {
	if len(p.toks) != 0 {
		p.next = p.toks[0]
	} else {
		p.next = tok.NewToken(tok.EOF) //TODO Maybe put this in the tokenizer
	}
}

func (p *Parser) eat(s tok.Symbol) {
	p.current, p.toks = p.toks[0], p.toks[1:]

	if p.current.Sym != s {
		p.err = fmt.Errorf("Error, Shield expected a %s \n", s.Sym)
	}
}

func (p *Parser) program() {
	p.peak()
	for len(p.toks) != 0 && p.err != nil {
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

}

func (p *Parser) func_def() {

}

func (p *Parser) func_call() {

}

func (p *Parser) paren_expr() {

}

func (p *Parser) paren_index() {

}

func (p *Parser) paren_comma_term() {

}

func (p *Parser) index() {

}

func (p *Parser) comma_term() {

}

func (p *Parser) returner() {

}

func (p *Parser) expr() {
	p.eat(tok.LABEL)
	p.eat(tok.EQUALS)
	p.term()
}

func (p *Parser) math() {
	p.term()
	p.eat(tok.PLUS) // Check for operators here
	p.term()
}

func (p *Parser) term() {

	// label or func call
	if p.next.Sym == tok.LABEL {
		p.eat(tok.LABEL)
		if p.next.Sym == tok.LABEL {
			p.func_call()
		} else {
			p.unexpected(p.next)
		}
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
