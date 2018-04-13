// Code generated from parser/Shield.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Shield

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ShieldListener is a complete listener for a parse tree produced by ShieldParser.
type ShieldListener interface {
	antlr.ParseTreeListener

	// EnterShield is called when entering the shield production.
	EnterShield(c *ShieldContext)

	// EnterEquals is called when entering the equals production.
	EnterEquals(c *EqualsContext)

	// EnterMath is called when entering the math production.
	EnterMath(c *MathContext)

	// EnterTerm is called when entering the term production.
	EnterTerm(c *TermContext)

	// ExitShield is called when exiting the shield production.
	ExitShield(c *ShieldContext)

	// ExitEquals is called when exiting the equals production.
	ExitEquals(c *EqualsContext)

	// ExitMath is called when exiting the math production.
	ExitMath(c *MathContext)

	// ExitTerm is called when exiting the term production.
	ExitTerm(c *TermContext)
}
