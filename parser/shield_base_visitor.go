// Code generated from parser/Shield.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Shield

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseShieldVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseShieldVisitor) VisitShield(ctx *ShieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShieldVisitor) VisitEquals(ctx *EqualsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShieldVisitor) VisitMath(ctx *MathContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShieldVisitor) VisitTerm(ctx *TermContext) interface{} {
	return v.VisitChildren(ctx)
}
