// Code generated from parser/Shield.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 11, 53, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4,
	6, 4, 27, 10, 4, 13, 4, 14, 4, 28, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 6, 8, 42, 10, 8, 13, 8, 14, 8, 43, 3, 9,
	6, 9, 47, 10, 9, 13, 9, 14, 9, 48, 3, 10, 3, 10, 3, 10, 2, 2, 11, 3, 3,
	5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 3, 2, 5, 4, 2, 12,
	12, 15, 15, 3, 2, 50, 59, 4, 2, 67, 92, 99, 124, 2, 55, 2, 3, 3, 2, 2,
	2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2,
	2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2,
	2, 2, 3, 21, 3, 2, 2, 2, 5, 23, 3, 2, 2, 2, 7, 26, 3, 2, 2, 2, 9, 30, 3,
	2, 2, 2, 11, 34, 3, 2, 2, 2, 13, 38, 3, 2, 2, 2, 15, 41, 3, 2, 2, 2, 17,
	46, 3, 2, 2, 2, 19, 50, 3, 2, 2, 2, 21, 22, 7, 45, 2, 2, 22, 4, 3, 2, 2,
	2, 23, 24, 7, 47, 2, 2, 24, 6, 3, 2, 2, 2, 25, 27, 9, 2, 2, 2, 26, 25,
	3, 2, 2, 2, 27, 28, 3, 2, 2, 2, 28, 26, 3, 2, 2, 2, 28, 29, 3, 2, 2, 2,
	29, 8, 3, 2, 2, 2, 30, 31, 7, 34, 2, 2, 31, 32, 3, 2, 2, 2, 32, 33, 8,
	5, 2, 2, 33, 10, 3, 2, 2, 2, 34, 35, 7, 36, 2, 2, 35, 36, 5, 17, 9, 2,
	36, 37, 7, 36, 2, 2, 37, 12, 3, 2, 2, 2, 38, 39, 5, 17, 9, 2, 39, 14, 3,
	2, 2, 2, 40, 42, 9, 3, 2, 2, 41, 40, 3, 2, 2, 2, 42, 43, 3, 2, 2, 2, 43,
	41, 3, 2, 2, 2, 43, 44, 3, 2, 2, 2, 44, 16, 3, 2, 2, 2, 45, 47, 9, 4, 2,
	2, 46, 45, 3, 2, 2, 2, 47, 48, 3, 2, 2, 2, 48, 46, 3, 2, 2, 2, 48, 49,
	3, 2, 2, 2, 49, 18, 3, 2, 2, 2, 50, 51, 7, 107, 2, 2, 51, 52, 7, 117, 2,
	2, 52, 20, 3, 2, 2, 2, 6, 2, 28, 43, 48, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'+'", "'-'", "", "' '", "", "", "", "", "'is'",
}

var lexerSymbolicNames = []string{
	"", "", "", "NEWLINE", "WHITESPACE", "STRING", "LABEL", "INT", "ALPHA",
	"IS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "NEWLINE", "WHITESPACE", "STRING", "LABEL", "INT", "ALPHA",
	"IS",
}

type ShieldLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewShieldLexer(input antlr.CharStream) *ShieldLexer {

	l := new(ShieldLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Shield.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// ShieldLexer tokens.
const (
	ShieldLexerT__0       = 1
	ShieldLexerT__1       = 2
	ShieldLexerNEWLINE    = 3
	ShieldLexerWHITESPACE = 4
	ShieldLexerSTRING     = 5
	ShieldLexerLABEL      = 6
	ShieldLexerINT        = 7
	ShieldLexerALPHA      = 8
	ShieldLexerIS         = 9
)