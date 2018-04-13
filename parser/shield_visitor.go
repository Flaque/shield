// Code generated from parser/Shield.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Shield

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by ShieldParser.
type ShieldVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by ShieldParser#shield.
	VisitShield(ctx *ShieldContext) interface{}

	// Visit a parse tree produced by ShieldParser#equals.
	VisitEquals(ctx *EqualsContext) interface{}

	// Visit a parse tree produced by ShieldParser#math.
	VisitMath(ctx *MathContext) interface{}

	// Visit a parse tree produced by ShieldParser#term.
	VisitTerm(ctx *TermContext) interface{}
}
