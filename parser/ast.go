package parser

import tok "github.com/flaque/shield/tokenizer"

type Group int

const (
	PROGRAM Group = iota
	CODE_BLOCK
	STATEMENT
	FUNC_DEF
	NIL
)

type AST struct {
	value    *tok.Token
	group    Group
	children []AST
}

func NewASTToken(t *tok.Token) AST {
	return AST{
		value:    t,
		children: []AST{},
		group:    NIL,
	}
}

func NewASTGroup(group Group) AST {
	return AST{
		value:    nil,
		children: []AST{},
		group:    group,
	}
}

func (a *AST) addToken(t *tok.Token) AST {
	ast := NewASTToken(t)
	a.children = append(a.children, ast)
	return ast
}

func (a *AST) addGroup(group Group) AST {
	ast := NewASTGroup(group)
	a.children = append(a.children, ast)
	return ast
}
