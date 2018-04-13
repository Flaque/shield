// Code generated from parser/Shield.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Shield

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 11, 36, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 3, 2, 3, 2, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 5, 3, 18, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 7, 4, 29, 10, 4, 12, 4, 14, 4, 32, 11, 4, 3, 5, 3, 5, 3,
	5, 2, 3, 6, 6, 2, 4, 6, 8, 2, 3, 3, 2, 7, 9, 2, 34, 2, 10, 3, 2, 2, 2,
	4, 17, 3, 2, 2, 2, 6, 19, 3, 2, 2, 2, 8, 33, 3, 2, 2, 2, 10, 11, 5, 6,
	4, 2, 11, 3, 3, 2, 2, 2, 12, 18, 5, 6, 4, 2, 13, 14, 5, 6, 4, 2, 14, 15,
	7, 11, 2, 2, 15, 16, 5, 6, 4, 2, 16, 18, 3, 2, 2, 2, 17, 12, 3, 2, 2, 2,
	17, 13, 3, 2, 2, 2, 18, 5, 3, 2, 2, 2, 19, 20, 8, 4, 1, 2, 20, 21, 5, 8,
	5, 2, 21, 30, 3, 2, 2, 2, 22, 23, 12, 4, 2, 2, 23, 24, 7, 3, 2, 2, 24,
	29, 5, 8, 5, 2, 25, 26, 12, 3, 2, 2, 26, 27, 7, 4, 2, 2, 27, 29, 5, 8,
	5, 2, 28, 22, 3, 2, 2, 2, 28, 25, 3, 2, 2, 2, 29, 32, 3, 2, 2, 2, 30, 28,
	3, 2, 2, 2, 30, 31, 3, 2, 2, 2, 31, 7, 3, 2, 2, 2, 32, 30, 3, 2, 2, 2,
	33, 34, 9, 2, 2, 2, 34, 9, 3, 2, 2, 2, 5, 17, 28, 30,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'+'", "'-'", "", "' '", "", "", "", "", "'is'",
}
var symbolicNames = []string{
	"", "", "", "NEWLINE", "WHITESPACE", "STRING", "LABEL", "INT", "ALPHA",
	"IS",
}

var ruleNames = []string{
	"shield", "equals", "math", "term",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type ShieldParser struct {
	*antlr.BaseParser
}

func NewShieldParser(input antlr.TokenStream) *ShieldParser {
	this := new(ShieldParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Shield.g4"

	return this
}

// ShieldParser tokens.
const (
	ShieldParserEOF        = antlr.TokenEOF
	ShieldParserT__0       = 1
	ShieldParserT__1       = 2
	ShieldParserNEWLINE    = 3
	ShieldParserWHITESPACE = 4
	ShieldParserSTRING     = 5
	ShieldParserLABEL      = 6
	ShieldParserINT        = 7
	ShieldParserALPHA      = 8
	ShieldParserIS         = 9
)

// ShieldParser rules.
const (
	ShieldParserRULE_shield = 0
	ShieldParserRULE_equals = 1
	ShieldParserRULE_math   = 2
	ShieldParserRULE_term   = 3
)

// IShieldContext is an interface to support dynamic dispatch.
type IShieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsShieldContext differentiates from other interfaces.
	IsShieldContext()
}

type ShieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyShieldContext() *ShieldContext {
	var p = new(ShieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ShieldParserRULE_shield
	return p
}

func (*ShieldContext) IsShieldContext() {}

func NewShieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ShieldContext {
	var p = new(ShieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ShieldParserRULE_shield

	return p
}

func (s *ShieldContext) GetParser() antlr.Parser { return s.parser }

func (s *ShieldContext) Math() IMathContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMathContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMathContext)
}

func (s *ShieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ShieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ShieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShieldListener); ok {
		listenerT.EnterShield(s)
	}
}

func (s *ShieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShieldListener); ok {
		listenerT.ExitShield(s)
	}
}

func (s *ShieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ShieldVisitor:
		return t.VisitShield(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ShieldParser) Shield() (localctx IShieldContext) {
	localctx = NewShieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ShieldParserRULE_shield)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(8)
		p.math(0)
	}

	return localctx
}

// IEqualsContext is an interface to support dynamic dispatch.
type IEqualsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEqualsContext differentiates from other interfaces.
	IsEqualsContext()
}

type EqualsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEqualsContext() *EqualsContext {
	var p = new(EqualsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ShieldParserRULE_equals
	return p
}

func (*EqualsContext) IsEqualsContext() {}

func NewEqualsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EqualsContext {
	var p = new(EqualsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ShieldParserRULE_equals

	return p
}

func (s *EqualsContext) GetParser() antlr.Parser { return s.parser }

func (s *EqualsContext) AllMath() []IMathContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMathContext)(nil)).Elem())
	var tst = make([]IMathContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMathContext)
		}
	}

	return tst
}

func (s *EqualsContext) Math(i int) IMathContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMathContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMathContext)
}

func (s *EqualsContext) IS() antlr.TerminalNode {
	return s.GetToken(ShieldParserIS, 0)
}

func (s *EqualsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EqualsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShieldListener); ok {
		listenerT.EnterEquals(s)
	}
}

func (s *EqualsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShieldListener); ok {
		listenerT.ExitEquals(s)
	}
}

func (s *EqualsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ShieldVisitor:
		return t.VisitEquals(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ShieldParser) Equals() (localctx IEqualsContext) {
	localctx = NewEqualsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ShieldParserRULE_equals)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(15)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(10)
			p.math(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(11)
			p.math(0)
		}
		{
			p.SetState(12)
			p.Match(ShieldParserIS)
		}
		{
			p.SetState(13)
			p.math(0)
		}

	}

	return localctx
}

// IMathContext is an interface to support dynamic dispatch.
type IMathContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMathContext differentiates from other interfaces.
	IsMathContext()
}

type MathContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMathContext() *MathContext {
	var p = new(MathContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ShieldParserRULE_math
	return p
}

func (*MathContext) IsMathContext() {}

func NewMathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MathContext {
	var p = new(MathContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ShieldParserRULE_math

	return p
}

func (s *MathContext) GetParser() antlr.Parser { return s.parser }

func (s *MathContext) Term() ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *MathContext) Math() IMathContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMathContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMathContext)
}

func (s *MathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MathContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MathContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShieldListener); ok {
		listenerT.EnterMath(s)
	}
}

func (s *MathContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShieldListener); ok {
		listenerT.ExitMath(s)
	}
}

func (s *MathContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ShieldVisitor:
		return t.VisitMath(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ShieldParser) Math() (localctx IMathContext) {
	return p.math(0)
}

func (p *ShieldParser) math(_p int) (localctx IMathContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewMathContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IMathContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 4
	p.EnterRecursionRule(localctx, 4, ShieldParserRULE_math, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(18)
		p.Term()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(28)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(26)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMathContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, ShieldParserRULE_math)
				p.SetState(20)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(21)
					p.Match(ShieldParserT__0)
				}
				{
					p.SetState(22)
					p.Term()
				}

			case 2:
				localctx = NewMathContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, ShieldParserRULE_math)
				p.SetState(23)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(24)
					p.Match(ShieldParserT__1)
				}
				{
					p.SetState(25)
					p.Term()
				}

			}

		}
		p.SetState(30)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}

	return localctx
}

// ITermContext is an interface to support dynamic dispatch.
type ITermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTermContext differentiates from other interfaces.
	IsTermContext()
}

type TermContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTermContext() *TermContext {
	var p = new(TermContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ShieldParserRULE_term
	return p
}

func (*TermContext) IsTermContext() {}

func NewTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermContext {
	var p = new(TermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ShieldParserRULE_term

	return p
}

func (s *TermContext) GetParser() antlr.Parser { return s.parser }

func (s *TermContext) LABEL() antlr.TerminalNode {
	return s.GetToken(ShieldParserLABEL, 0)
}

func (s *TermContext) INT() antlr.TerminalNode {
	return s.GetToken(ShieldParserINT, 0)
}

func (s *TermContext) STRING() antlr.TerminalNode {
	return s.GetToken(ShieldParserSTRING, 0)
}

func (s *TermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShieldListener); ok {
		listenerT.EnterTerm(s)
	}
}

func (s *TermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShieldListener); ok {
		listenerT.ExitTerm(s)
	}
}

func (s *TermContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ShieldVisitor:
		return t.VisitTerm(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ShieldParser) Term() (localctx ITermContext) {
	localctx = NewTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ShieldParserRULE_term)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(31)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ShieldParserSTRING)|(1<<ShieldParserLABEL)|(1<<ShieldParserINT))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

func (p *ShieldParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 2:
		var t *MathContext = nil
		if localctx != nil {
			t = localctx.(*MathContext)
		}
		return p.Math_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *ShieldParser) Math_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
