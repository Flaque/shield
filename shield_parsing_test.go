package main

import (
	"fmt"
	"testing"

	"github.com/flaque/shield/parser"
	"github.com/flaque/shield/tokenizer"
	"github.com/stretchr/testify/assert"
)

func check(code string) error {
	toko := tokenizer.NewTokenizer(code)
	tokens := toko.Run()

	fmt.Println(tokens)

	return parser.NewParser().Run(tokens)
}

func assertBreaks(t *testing.T, code string) {
	assert.NotNil(t, check(code))
}

func assertWorks(t *testing.T, code string) {
	assert.Nil(t, check(code))
}

func TestProperFunction(t *testing.T) {
	code := `
	func add(a, b) does 
		add(1, 2) is 3
	by 
		return a + b
	end
	`

	assertWorks(t, code)
}

func TestScriptSyntax(t *testing.T) {
	code := `
	num = add(a, b)
	print(num)
	`

	assertWorks(t, code)
}

func TestMultiFunctions(t *testing.T) {
	code := `
	func add(a, b) does 
		add(1, 2) is 3
	by 
		return a + b
	end

	func minus(a, b) does 
		minus(1, 2) is 3
	by 
		return a + b
	end
	`

	assertWorks(t, code)
}

func TestImproperFunctionDef(t *testing.T) {
	code := `
	function add(a, b) does
		add(1, 2) is 3
	by 
		return a + b
	end
	`

	assertBreaks(t, code)
}

func TestLiteralInFuncDefinition(t *testing.T) {
	code := `
	func add(a, 3, b) does 
		add(1, 2) is 3
	by 
		return a + b
	end
	`

	assertBreaks(t, code)
}

func TestNoBy(t *testing.T) {
	code := `
	func add(a, b) does 
		return a + b
	end
	`

	assertBreaks(t, code)
}
