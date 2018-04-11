package parser

import (
	"testing"

	tok "github.com/flaque/shield/tokenizer"
	"github.com/stretchr/testify/assert"
)

func tokList(symbols ...tok.Symbol) []tok.Token {
	toks := []tok.Token{}

	for _, sym := range symbols {
		toks = append(toks, tok.NewToken(sym))
	}
	return toks
}

func parserWith(symbols ...tok.Symbol) *Parser {
	p := NewParser()
	p.toks = tokList(symbols...)
	p.peak()
	return p
}

func TestLiteral(t *testing.T) {
	p := parserWith(tok.NUM, tok.EOF)

	p.literal()
	assert.Nil(t, p.err)
}

func TestMath(t *testing.T) {
	p := parserWith(tok.PLUS, tok.NUM, tok.EOF)

	p.math()
	assert.Nil(t, p.err)
}

func TestTermAsLiteral(t *testing.T) {
	p := parserWith(tok.NUM, tok.EOF)

	p.term()
	assert.Nil(t, p.err)
}

func TestTermAsFuncCall(t *testing.T) {
	p := parserWith(tok.LABEL, tok.L_PAREN, tok.NUM, tok.R_PAREN, tok.EOF)

	p.term()
	assert.Nil(t, p.err)
}

func TestTermAsCommaSeperatedFuncCall(t *testing.T) {
	p := parserWith(tok.LABEL, tok.L_PAREN, tok.NUM, tok.COMMA, tok.NUM, tok.R_PAREN, tok.EOF)

	p.term()
	assert.Nil(t, p.err)
}

func TestReturn(t *testing.T) {
	p := parserWith(tok.RETURN, tok.NUM, tok.EOF)

	p.returner()
	assert.Nil(t, p.err)
}

func TestStatementAsReturn(t *testing.T) {
	p := parserWith(tok.RETURN, tok.NUM, tok.EOF)

	p.statement()
	assert.Nil(t, p.err)
}

func TestStatementAsFuncCall(t *testing.T) {
	p := parserWith(tok.LABEL, tok.L_PAREN, tok.LABEL, tok.R_PAREN, tok.EOF)

	p.statement()
	assert.Nil(t, p.err)
}

func TestStatementAsAssignment(t *testing.T) {
	p := parserWith(tok.LABEL, tok.ASSIGN, tok.NUM, tok.EOF)

	p.statement()
	assert.Nil(t, p.err)
}

func TestParenCommaTermWithOneArg(t *testing.T) {
	p := parserWith(tok.L_PAREN, tok.NUM, tok.R_PAREN, tok.EOF)

	p.paren_comma_term()
	assert.Nil(t, p.err)
}

func TestParenCommaTermWithMultiArgs(t *testing.T) {
	p := parserWith(tok.L_PAREN, tok.NUM, tok.COMMA, tok.NUM, tok.R_PAREN, tok.EOF)

	p.paren_comma_term()
	assert.Nil(t, p.err)
}

func TestParenIndexWithOneArg(t *testing.T) {
	p := parserWith(tok.L_PAREN, tok.LABEL, tok.R_PAREN, tok.EOF)

	p.paren_index()
	assert.Nil(t, p.err)
}

func TestParenIndexWithMultiArgs(t *testing.T) {
	p := parserWith(
		tok.L_PAREN,
		tok.LABEL,
		tok.COMMA,
		tok.LABEL,
		tok.R_PAREN,
		tok.EOF,
	)

	p.paren_index()
	assert.Nil(t, p.err)
}

func TestFuncEquals(t *testing.T) {
	p := parserWith(tok.NUM, tok.IS, tok.NUM, tok.EOF)

	p.equals()
	assert.Nil(t, p.err)
}

func TestFuncDef(t *testing.T) {

	// A whole darn function!
	p := parserWith(
		tok.FUNC,
		tok.LABEL,
		tok.L_PAREN,
		tok.LABEL,
		tok.COMMA,
		tok.LABEL,
		tok.R_PAREN,
		tok.DOES,
		tok.LABEL,
		tok.L_PAREN,
		tok.NUM,
		tok.COMMA,
		tok.NUM,
		tok.R_PAREN,
		tok.IS,
		tok.NUM,
		tok.BY,
		tok.RETURN,
		tok.LABEL,
		tok.PLUS,
		tok.LABEL,
		tok.END,
		tok.EOF,
	)

	p.func_def()
	assert.Nil(t, p.err)
}
