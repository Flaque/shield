// Code generated from parser/Shield.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Shield

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseShieldListener is a complete listener for a parse tree produced by ShieldParser.
type BaseShieldListener struct{}

var _ ShieldListener = &BaseShieldListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseShieldListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseShieldListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseShieldListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseShieldListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterShield is called when production shield is entered.
func (s *BaseShieldListener) EnterShield(ctx *ShieldContext) {}

// ExitShield is called when production shield is exited.
func (s *BaseShieldListener) ExitShield(ctx *ShieldContext) {}

// EnterEquals is called when production equals is entered.
func (s *BaseShieldListener) EnterEquals(ctx *EqualsContext) {}

// ExitEquals is called when production equals is exited.
func (s *BaseShieldListener) ExitEquals(ctx *EqualsContext) {}

// EnterMath is called when production math is entered.
func (s *BaseShieldListener) EnterMath(ctx *MathContext) {}

// ExitMath is called when production math is exited.
func (s *BaseShieldListener) ExitMath(ctx *MathContext) {}

// EnterTerm is called when production term is entered.
func (s *BaseShieldListener) EnterTerm(ctx *TermContext) {}

// ExitTerm is called when production term is exited.
func (s *BaseShieldListener) ExitTerm(ctx *TermContext) {}
