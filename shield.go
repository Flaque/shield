package main

import (
	"fmt"
	"os"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/flaque/shield/parser"
)

type TreeShapeListener struct {
	*parser.BaseShieldListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (this *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(ctx.GetText())
}

func main() {
	input, _ := antlr.NewFileStream(os.Args[1])
	lexer := parser.NewShieldLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewShieldParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

	p.BuildParseTrees = true
	tree := p.Math()

	antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)
}
