package engine

import (
	"fmt"
	"math"
	"strings"

	"github.com/MarionetteEnsemble/Puppeto/lib"
	"github.com/antlr4-go/antlr"
)

type PuppetoInterpreter struct {
	File    string
	Scope   *Scope
	Stack   Stack
	OutFunc OutFuncType
	Parser  *lib.PuppetoParser
}

type PuppetoErrorListener struct {
	*antlr.DefaultErrorListener
	CurrentError error
}

func NewPuppetoErrorListener() *PuppetoErrorListener {
	return new(PuppetoErrorListener)
}

func (c *PuppetoErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	c.CurrentError = fmt.Errorf("Syntax Error at line %d, column %d: %s", line, column, msg)
}

type OutFuncType = func(string)

func DefaultOutFunc(s string) {
	fmt.Println(s)
}

func NewPuppetoInterpreter(parser *lib.PuppetoParser, file string, scope *Scope, stack Stack, outFunc OutFuncType) *PuppetoInterpreter {
	return &PuppetoInterpreter{
		Parser:  parser,
		File:    file,
		Scope:   scope,
		Stack:   stack,
		OutFunc: outFunc,
	}
}

func (t *PuppetoInterpreter) TokenToPosition(tok antlr.Token) Position {
	return NewPosition(int64(tok.GetLine()), int64(tok.GetColumn()), t.File)
}

func (t *PuppetoInterpreter) VisitLiteralValue(ctx *lib.LiteralValueContext) (PValue, PError) {
	if ctx.FLOAT() != nil {
		return StrToF(ctx.FLOAT().GetText()), nil
	} else if ctx.INTEGER() != nil {
		return StrToInt(ctx.INTEGER().GetText()), nil
	} else if ctx.STRING() != nil {
		str := ctx.STRING().GetText()
		s := Unquote(str)

		return s, nil
	} else if ctx.IDENTIFIER() != nil {
		ident := ctx.IDENTIFIER()
		v, e := t.Scope.Get(ident.GetText())

		if e != nil {
			return nil, NewPError(TypeError, e.Error(), append(Stack{t.TokenToPosition(ident.GetSymbol())}, t.Stack...))
		}

		return v, nil
	} else if ctx.NIL() != nil {
		return nil, nil
	} else if ctx.BOOLEAN() != nil {
		if ctx.BOOLEAN().GetText() == "true" {
			return true, nil
		}

		return false, nil
	} else if ctx.NAN() != nil {
		return math.NaN(), nil
	}

	panic("LiteralValue")
}

func (t *PuppetoInterpreter) VisitPair(ctx *lib.PairContext) (PValue, PValue, PError) {
	if ctx.GetChildCount() == 1 {
		key := ctx.IDENTIFIER().GetText()

		return key, key, nil
	}

	exps := ctx.AllExpression()
	var e PError
	var key, val PValue

	if len(exps) > 1 {
		key, e = t.VisitExpression(exps[0].(*lib.ExpressionContext))

		if e != nil {
			return nil, nil, e
		}

		val, e = t.VisitExpression(exps[1].(*lib.ExpressionContext))

		if e != nil {
			return nil, nil, e
		}
	} else {
		identI := ctx.IDENTIFIER()

		if identI != nil {
			key = identI.GetText()
		} else {
			key = Unquote(ctx.STRING().GetText())
		}

		val, e = t.VisitExpression(exps[0].(*lib.ExpressionContext))

		if e != nil {
			return nil, nil, e
		}
	}

	return key, val, nil
}

func (t *PuppetoInterpreter) VisitObject(ctx *lib.ObjectContext) (PValue, PError) {
	val := PObject{}

	for _, pairI := range ctx.AllPair() {
		k, v, e := t.VisitPair(pairI.(*lib.PairContext))

		if e != nil {
			return nil, e
		}

		val[ToStr(k)] = v
	}

	return val, nil
}

func (t *PuppetoInterpreter) VisitArray(ctx *lib.ArrayContext) (PValue, PError) {
	a := PArray{}

	for _, expI := range ctx.AllExpression() {
		v, e := t.VisitExpression(expI.(*lib.ExpressionContext))

		if e != nil {
			return nil, e
		}

		a = append(a, v)
	}

	return a, nil
}

func (t *PuppetoInterpreter) VisitFunctionDefinition(ctx *lib.FunctionDefinitionContext, isExpression bool) (PValue, PError) {
	c := ctx.ScopeBody()
	var fn PFunction = func(stack Stack, args []PValue) (PValue, PValue) {
		t := NewPuppetoInterpreter(t.Parser, t.File, NewScope(t.Scope), append(Stack{t.TokenToPosition(ctx.GetStart())}, t.Stack...), t.OutFunc)
		argsListI := ctx.ArgumentList()
		var dargs []Argument = []Argument{}

		if argsListI != nil {
			dargs = t.VisitArgumentList(argsListI.(*lib.ArgumentListContext))

			for i, arg := range dargs {
				if len(args) > i {
					_ = t.Scope.Define(arg.Name, args[i])
				} else {
					v, e := arg.GetDefaultValue()

					if e != nil {
						return nil, e
					}

					_ = t.Scope.Define(arg.Name, v)
				}
			}
		}

		if ctx.GetRest() != nil && len(dargs) < len(args) {
			_ = t.Scope.Define(ctx.GetRest().GetText(), args[len(dargs):])
		}

		_, v, e := t.VisitScopeBody(c.(*lib.ScopeBodyContext))

		return v, e
	}

	if !isExpression {
		name := ctx.GetName().GetText()
		e := t.Scope.Define(name, fn)

		if e != nil {
			return nil, e
		}
	}

	return fn, nil
}

func (t *PuppetoInterpreter) VisitArgumentList(ctx *lib.ArgumentListContext) []Argument {
	argsD := []Argument{}

	for _, argI := range ctx.AllArgument() {
		argsD = append(argsD, t.VisitArgument(argI.(*lib.ArgumentContext)))
	}

	return argsD
}

type Argument struct {
	Name            string
	GetDefaultValue func() (PValue, PValue)
}

func (t *PuppetoInterpreter) VisitArgument(ctx *lib.ArgumentContext) Argument {
	// Implementation for VisitArgsList method
	return Argument{ctx.IDENTIFIER().GetText(), func() (PValue, PValue) {
		exp := ctx.Expression()

		if exp == nil {
			return nil, nil
		}

		return t.VisitExpression(exp.(*lib.ExpressionContext))
	}}
}

func (t *PuppetoInterpreter) VisitScopeBody(ctx *lib.ScopeBodyContext) (bool, PValue, PError) {
	if ctx.Expression() != nil {
		v, e := t.VisitExpression(ctx.Expression().(*lib.ExpressionContext))

		return true, v, e
	} else if ctx.Scope() != nil {
		return t.VisitScope(ctx.Scope().(*lib.ScopeContext))
	}

	panic("ScopeBody")
}

func (t *PuppetoInterpreter) VisitValue(ctx *lib.ValueContext) (PValue, PError) {
	if ctx.LiteralValue() != nil {
		return t.VisitLiteralValue(ctx.LiteralValue().(*lib.LiteralValueContext))
	} else if ctx.Expression() != nil {
		return t.VisitExpression(ctx.Expression().(*lib.ExpressionContext))
	} else if ctx.Object() != nil {
		return t.VisitObject(ctx.Object().(*lib.ObjectContext))
	} else if ctx.Array() != nil {
		return t.VisitArray(ctx.Array().(*lib.ArrayContext))
	} else if ctx.FunctionDefinition() != nil {
		return t.VisitFunctionDefinition(ctx.FunctionDefinition().(*lib.FunctionDefinitionContext), true)
	}

	panic("Value")
}

func (t *PuppetoInterpreter) VisitExpression(ctx *lib.ExpressionContext) (PValue, PError) {
	return t.VisitLogicalExpression(ctx.LogicalExpression().(*lib.LogicalExpressionContext))
}

func (t *PuppetoInterpreter) VisitLogicalExpression(ctx *lib.LogicalExpressionContext) (PValue, PError) {
	exps := ctx.AllEqualityExpression()
	lhs, e := t.VisitEqualityExpression(exps[0].(*lib.EqualityExpressionContext))

	if e != nil {
		return nil, e
	}

	if len(exps) > 1 {
		for i, rhsI := range exps[1:] {
			rhs, e := t.VisitEqualityExpression(rhsI.(*lib.EqualityExpressionContext))

			if e != nil {
				return nil, e
			}

			switch ctx.GetChild((i+1)*2 - 1).(*antlr.TerminalNodeImpl).GetText() {
			case "&&":
				lhs = And(lhs, rhs)
			case "||":
				lhs = Or(lhs, rhs)
			}
		}
	}

	return lhs, nil
}

func (t *PuppetoInterpreter) VisitEqualityExpression(ctx *lib.EqualityExpressionContext) (PValue, PError) {
	exps := ctx.AllComparisonExpression()
	lhs, e := t.VisitComparisonExpression(exps[0].(*lib.ComparisonExpressionContext))

	if e != nil {
		return nil, e
	}

	if len(exps) > 1 {
		for i, rhsI := range exps[1:] {
			rhs, e := t.VisitComparisonExpression(rhsI.(*lib.ComparisonExpressionContext))

			if e != nil {
				return nil, e
			}

			switch ctx.GetChild((i+1)*2 - 1).(*antlr.TerminalNodeImpl).GetText() {
			case "==":
				lhs = Eq(lhs, rhs)
			case "!=":
				lhs = Neq(lhs, rhs)
			}
		}
	}

	return lhs, nil
}

func (t *PuppetoInterpreter) VisitComparisonExpression(ctx *lib.ComparisonExpressionContext) (PValue, PError) {
	exps := ctx.AllAdditiveExpression()
	lhs, e := t.VisitAdditiveExpression(exps[0].(*lib.AdditiveExpressionContext))

	if e != nil {
		return nil, e
	}

	if len(exps) > 1 {
		for i, rhsI := range exps[1:] {
			rhs, e := t.VisitAdditiveExpression(rhsI.(*lib.AdditiveExpressionContext))

			if e != nil {
				return nil, e
			}

			switch ctx.GetChild((i+1)*2 - 1).(*antlr.TerminalNodeImpl).GetText() {
			case ">=":
				lhs = Gtoe(lhs, rhs)
			case "<=":
				lhs = Ltoe(lhs, rhs)
			case ">":
				lhs = Gt(lhs, rhs)
			case "<":
				lhs = Lt(lhs, rhs)
			}
		}
	}

	return lhs, nil
}

func (t *PuppetoInterpreter) VisitAdditiveExpression(ctx *lib.AdditiveExpressionContext) (PValue, PError) {
	exps := ctx.AllMultiplicativeExpression()
	lhs, e := t.VisitMultiplicativeExpression(exps[0].(*lib.MultiplicativeExpressionContext))

	if e != nil {
		return nil, e
	}

	if len(exps) > 1 {
		for i, rhsI := range exps[1:] {
			rhs, e := t.VisitMultiplicativeExpression(rhsI.(*lib.MultiplicativeExpressionContext))

			if e != nil {
				return nil, e
			}

			switch ctx.GetChild((i+1)*2 - 1).(*antlr.TerminalNodeImpl).GetText() {
			case "+":
				lhs = Sum(lhs, rhs)
			case "-":
				lhs = Sub(lhs, rhs)
			}
		}
	}

	return lhs, nil
}

func (t *PuppetoInterpreter) VisitMultiplicativeExpression(ctx *lib.MultiplicativeExpressionContext) (PValue, PError) {
	exps := ctx.AllPowerExpression()
	lhs, e := t.VisitPowerExpression(exps[0].(*lib.PowerExpressionContext))

	if e != nil {
		return nil, e
	}

	if len(exps) > 1 {
		for i, rhsI := range exps[1:] {
			rhs, e := t.VisitPowerExpression(rhsI.(*lib.PowerExpressionContext))

			if e != nil {
				return nil, e
			}

			switch ctx.GetChild((i+1)*2 - 1).(*antlr.TerminalNodeImpl).GetText() {
			case "*":
				lhs = Mul(lhs, rhs)
			case "/":
				lhs = Div(lhs, rhs)
			case "%":
				lhs = Mod(lhs, rhs)
			}
		}
	}

	return lhs, nil
}

func (t *PuppetoInterpreter) VisitPowerExpression(ctx *lib.PowerExpressionContext) (PValue, PError) {
	exps := ctx.AllBitwiseExpression()
	lhs, e := t.VisitBitwiseExpression(exps[0].(*lib.BitwiseExpressionContext))

	if e != nil {
		return nil, e
	}

	if len(exps) > 1 {
		for i, rhsI := range exps[1:] {
			rhs, e := t.VisitBitwiseExpression(rhsI.(*lib.BitwiseExpressionContext))

			if e != nil {
				return nil, e
			}

			switch ctx.GetChild((i+1)*2 - 1).(*antlr.TerminalNodeImpl).GetText() {
			case "**":
				lhs = Pow(lhs, rhs)
			}
		}
	}

	return lhs, nil
}

func (t *PuppetoInterpreter) VisitBitwiseExpression(ctx *lib.BitwiseExpressionContext) (PValue, PError) {
	exps := ctx.AllShiftExpression()
	lhs, e := t.VisitShiftExpression(exps[0].(*lib.ShiftExpressionContext))

	if e != nil {
		return nil, e
	}

	if len(exps) > 1 {
		for i, rhsI := range exps[1:] {
			rhs, e := t.VisitShiftExpression(rhsI.(*lib.ShiftExpressionContext))

			if e != nil {
				return nil, e
			}

			switch ctx.GetChild((i+1)*2 - 1).(*antlr.TerminalNodeImpl).GetText() {
			case "&":
				lhs = Band(lhs, rhs)
			case "|":
				lhs = Bor(lhs, rhs)
			case "^":
				lhs = Bxor(lhs, rhs)
			}
		}
	}

	return lhs, nil
}

func (t *PuppetoInterpreter) VisitShiftExpression(ctx *lib.ShiftExpressionContext) (PValue, PError) {
	exps := ctx.AllUnaryExpression()
	lhs, e := t.VisitUnaryExpression(exps[0].(*lib.UnaryExpressionContext))

	if e != nil {
		return nil, e
	}

	if len(exps) > 1 {
		for i, rhsI := range exps[1:] {
			rhs, e := t.VisitUnaryExpression(rhsI.(*lib.UnaryExpressionContext))

			if e != nil {
				return nil, e
			}

			switch ctx.GetChild((i+1)*2 - 1).(*antlr.TerminalNodeImpl).GetText() {
			case "<<":
				lhs = Bls(lhs, rhs)
			case ">>":
				lhs = Brs(lhs, rhs)
			}
		}
	}

	return lhs, nil
}

func (t *PuppetoInterpreter) VisitUnaryExpression(ctx *lib.UnaryExpressionContext) (PValue, PError) {
	exp := ctx.UnaryExpression()

	if exp != nil {
		lhs, e := t.VisitUnaryExpression(exp.(*lib.UnaryExpressionContext))

		if e != nil {
			return nil, e
		}

		switch ctx.GetChild(1).(*antlr.TerminalNodeImpl).GetText() {
		case "!":
			return Not(lhs), nil
		case "-":
			return Neg(lhs), nil
		}

		return lhs, nil
	}

	return t.VisitChainExpression(ctx.ChainExpression().(*lib.ChainExpressionContext))

}

func (t *PuppetoInterpreter) VisitGetProperty(ctx *lib.GetPropertyContext) (PValue, PError) {
	v, e := t.VisitValue(ctx.Value().(*lib.ValueContext))

	if e != nil {
		return nil, e
	}

	k, e := t.VisitExpression(ctx.Expression().(*lib.ExpressionContext))

	if e != nil {
		return nil, e
	}

	return Get(v, k)
}

func (t *PuppetoInterpreter) VisitSetProperty(ctx *lib.SetPropertyContext) (PValue, PError) {
	getPropC := ctx.GetProperty().(*lib.GetPropertyContext)
	var op string

	if ctx.TWO_SIDES_OPERATOR() != nil {
		op = ctx.TWO_SIDES_OPERATOR().GetText()
	} else {
		op = "="
	}

	v, e := t.VisitValue(getPropC.Value().(*lib.ValueContext))

	if e != nil {
		return nil, e
	}

	k, e := t.VisitExpression(getPropC.Expression().(*lib.ExpressionContext))

	if e != nil {
		return nil, e
	}

	to, e := t.VisitExpression(ctx.Expression().(*lib.ExpressionContext))

	if e != nil {
		return nil, e
	}

	switch op {
	//case "=":
	case "+=":
		v, e := Get(v, k)

		if e != nil {
			return nil, e
		}

		to = Sum(to, v)
	case "-=":
		v, e := Get(v, k)

		if e != nil {
			return nil, e
		}

		to = Sub(to, v)
	case "*=":
		v, e := Get(v, k)

		if e != nil {
			return nil, e
		}

		to = Mul(to, v)
	case "/=":
		v, e := Get(v, k)

		if e != nil {
			return nil, e
		}

		to = Div(to, v)
	case "&=":
		v, e := Get(v, k)

		if e != nil {
			return nil, e
		}

		to = Band(to, v)
	case "|=":
		v, e := Get(v, k)

		if e != nil {
			return nil, e
		}

		to = Bor(to, v)
	case "**=":
		v, e := Get(v, k)

		if e != nil {
			return nil, e
		}

		to = Pow(to, v)
	case "^=":
		v, e := Get(v, k)

		if e != nil {
			return nil, e
		}

		to = Bxor(to, v)
	}

	return Set(v, k, to)
}

func (t *PuppetoInterpreter) VisitCallExpression(ctx *lib.CallExpressionContext) (PValue, PError) {
	c := ctx.Value().(*lib.ValueContext)
	fn, e := t.VisitValue(c)

	if e != nil {
		return nil, e
	}

	f, ok := fn.(PFunction)

	if !ok {
		return nil, NewPError(TypeError, "Expected function", append(Stack{t.TokenToPosition(ctx.GetStop())}, t.Stack...))
	}

	args := PArray{}

	for _, ctx := range ctx.AllExpression() {
		v, e := t.VisitExpression(ctx.(*lib.ExpressionContext))

		if e != nil {
			return nil, e
		}

		args = append(args, v)
	}

	return f(
		append(
			Stack{t.TokenToPosition(ctx.GetStart())},
			t.Stack...,
		),
		args,
	)
}

func (t *PuppetoInterpreter) VisitChainExpression(ctx *lib.ChainExpressionContext) (PValue, PError) {
	if ctx.Value() != nil {
		return t.VisitValue(ctx.Value().(*lib.ValueContext))
	} else if ctx.GetProperty() != nil {
		return t.VisitGetProperty(ctx.GetProperty().(*lib.GetPropertyContext))
	} else if ctx.SetProperty() != nil {
		return t.VisitSetProperty(ctx.SetProperty().(*lib.SetPropertyContext))
	} else if ctx.CallExpression() != nil {
		return t.VisitCallExpression(ctx.CallExpression().(*lib.CallExpressionContext))
	}

	panic("ChainExpression")
}

func (t *PuppetoInterpreter) VisitLhsVariableAssignation(ctx *lib.LhsVariableAssignationContext) (PValue, PError) {

	op := ctx.ONE_SIDE_OPERATOR().GetText()
	name := ctx.IDENTIFIER().GetText()
	v, e := t.Scope.Get(name)

	if e != nil {
		return nil, e
	}

	switch op {
	case "++":
		v = Sum(v, int64(1))
		e = t.Scope.Set(name, v)
	case "--":
		v = Sub(v, int64(1))
		e = t.Scope.Set(name, v)
	}

	if e != nil {
		return nil, e
	}

	return v, nil
}

func (t *PuppetoInterpreter) VisitRhsVariableAssignation(ctx *lib.RhsVariableAssignationContext) (PValue, PError) {
	op := ctx.ONE_SIDE_OPERATOR().GetText()
	name := ctx.IDENTIFIER().GetText()
	v, e := t.Scope.Get(name)

	if e != nil {
		return nil, e
	}

	switch op {
	case "++":
		e = t.Scope.Set(name, Sum(v, int64(1)))
	case "--":
		e = t.Scope.Set(name, Sub(v, int64(1)))
	}

	if e != nil {
		return nil, e
	}

	return v, nil
}

func (t *PuppetoInterpreter) VisitMidVariableAssignation(ctx *lib.MidVariableAssignationContext) (PValue, PError) {
	op := ctx.TWO_SIDES_OPERATOR().GetText()
	name := ctx.IDENTIFIER().GetText()
	v, e := t.VisitExpression(ctx.Expression().(*lib.ExpressionContext))
	var s any

	if e != nil {
		return nil, e
	}
	switch op {
	case "=":
		e = t.Scope.Set(name, v)
	case "+=":
		s, e = t.Scope.Get(name)

		if e != nil {
			return nil, e
		}

		e = t.Scope.Set(name, Sum(s, v))
	case "-=":
		s, e = t.Scope.Get(name)

		if e != nil {
			return nil, e
		}

		e = t.Scope.Set(name, Sub(s, v))
	case "*=":
		s, e = t.Scope.Get(name)

		if e != nil {
			return nil, e
		}

		e = t.Scope.Set(name, Mul(s, v))
	case "/=":
		s, e = t.Scope.Get(name)

		if e != nil {
			return nil, e
		}

		e = t.Scope.Set(name, Div(s, v))
	case "&=":
		s, e = t.Scope.Get(name)

		if e != nil {
			return nil, e
		}

		e = t.Scope.Set(name, Band(s, v))
	case "|=":
		s, e = t.Scope.Get(name)

		if e != nil {
			return nil, e
		}

		e = t.Scope.Set(name, Bor(s, v))
	case "**=":
		s, e = t.Scope.Get(name)

		if e != nil {
			return nil, e
		}

		e = t.Scope.Set(name, Pow(s, v))
	case "^=":
		s, e = t.Scope.Get(name)

		if e != nil {
			return nil, e
		}

		e = t.Scope.Set(name, Bxor(s, v))
	}
	if e != nil {
		return nil, e
	}

	return v, nil
}

func (t *PuppetoInterpreter) VisitVariableAssignation(ctx *lib.VariableAssignationContext) (PValue, PError) {
	var v, e any

	for _, node := range ctx.GetChildren() {
		switch node := node.(type) {
		case *lib.LhsVariableAssignationContext:
			v, e = t.VisitLhsVariableAssignation(node)
		case *lib.RhsVariableAssignationContext:
			v, e = t.VisitRhsVariableAssignation(node)
		case *lib.MidVariableAssignationContext:
			v, e = t.VisitMidVariableAssignation(node)
		}

		if e != nil {
			return nil, e
		}
	}

	return v, nil
}

func (t *PuppetoInterpreter) VisitVariableDefinition(ctx *lib.VariableDefinitionContext) PError {
	for _, ctx := range ctx.AllVariableDefinitionBody() {
		e := t.VisitVariableDefinitionBody(ctx.(*lib.VariableDefinitionBodyContext))

		if e != nil {
			return e
		}
	}

	return nil
}

func (t *PuppetoInterpreter) VisitVariableDefinitionBody(ctx *lib.VariableDefinitionBodyContext) PError {
	name := ctx.IDENTIFIER().GetText()
	exp := ctx.Expression()
	var v any

	if exp != nil {
		var e any

		v, e = t.VisitExpression(exp.(*lib.ExpressionContext))

		if e != nil {
			return e
		}
	} else {
		v = nil
	}

	return t.Scope.Define(name, v)
}

func (t *PuppetoInterpreter) VisitIfStatement(ctx *lib.IfStatementContext) (bool, PValue, PError) {
	scopeBodiesI := ctx.AllScopeBody()
	expsI := ctx.AllExpression()

	for i, expI := range expsI {
		v, e := t.VisitExpression(expI.(*lib.ExpressionContext))

		if e != nil {
			return true, nil, e
		}

		if ToBool(v) {
			return t.VisitScopeBody(scopeBodiesI[i].(*lib.ScopeBodyContext))

		}
	}

	if len(scopeBodiesI) > len(expsI) {
		return t.VisitScopeBody(scopeBodiesI[len(scopeBodiesI)-1].(*lib.ScopeBodyContext))
	}

	return true, nil, nil
}

func (t *PuppetoInterpreter) VisitProgramBodyWithBreakContinue(ctx *lib.ProgramBodyWithBreakContinueContext) (bool, PValue, PError, bool) {
	var e any

	node := ctx.GetChild(0)
	switch node := node.(type) {
	case *lib.ScopeWithBreakContinueContext:
		ok, v, e, ok2 := t.VisitScopeWithBreakContinue(node)

		if e != nil || !ok2 || !ok {
			return ok, v, e, ok2
		}
	case *lib.BreakStatementContext:
		return true, nil, nil, false
	case *lib.ContinueStatementContext:
		return true, nil, nil, false
	case *lib.ExpressionContext:
		_, e = t.VisitExpression(node)
	default:
		ok, v, e := t.VisitStatement(node.(*lib.StatementContext))

		return ok, v, e, true
	}

	if e != nil {
		return true, nil, e, true
	}

	return true, nil, nil, true
}

func (t *PuppetoInterpreter) VisitScopeWithBreakContinue(ctx *lib.ScopeWithBreakContinueContext) (bool, PValue, PError, bool) {
	for _, pb := range ctx.AllProgramBodyWithBreakContinue() {
		ok, v, e, ok2 := t.VisitProgramBodyWithBreakContinue(pb.(*lib.ProgramBodyWithBreakContinueContext))

		if e != nil || !ok2 {
			return ok, nil, e, ok2
		}

		if !ok {
			return ok, v, e, ok2
		}
	}

	return true, nil, nil, true
}

func (t *PuppetoInterpreter) VisitLoopStatement(ctx *lib.LoopStatementContext) (bool, PValue, PError) {
	exp := ctx.Expression().(*lib.ExpressionContext)
	programBody := ctx.ProgramBodyWithBreakContinue().(*lib.ProgramBodyWithBreakContinueContext)
	interpreter := NewPuppetoInterpreter(t.Parser, t.File, NewScope(t.Scope), t.Stack, t.OutFunc)

	for {
		v, e := t.VisitExpression(exp)

		if e != nil {
			return true, v, e
		}

		if !ToBool(v) {
			break
		}

		ok, v, e, ok2 := interpreter.VisitProgramBodyWithBreakContinue(programBody)

		if !ok2 || !ok {
			return ok, v, nil
		}

		if e != nil {
			return ok, nil, e
		}
	}

	return true, nil, nil
}

func (t *PuppetoInterpreter) VisitSwitchStatement(ctx *lib.SwitchStatementContext) (bool, PValue, PError) {
	v, e := t.VisitExpression(ctx.Expression().(*lib.ExpressionContext))

	if e != nil {
		return true, v, e
	}

	for _, scI := range ctx.AllSwitchCase() {
		interpreter := NewPuppetoInterpreter(t.Parser, t.File, NewScope(t.Scope), t.Stack, t.OutFunc)
		ok, v, e, ok2 := interpreter.VisitSwitchCase(scI.(*lib.SwitchCaseContext), v)

		if e != nil {
			return ok, nil, e
		}

		if !ok2 || !ok {
			return ok, v, nil
		}
	}

	for _, pb := range ctx.AllProgramBodyWithBreakContinue() {
		ok, v, e, ok2 := t.VisitProgramBodyWithBreakContinue(pb.(*lib.ProgramBodyWithBreakContinueContext))

		if e != nil {
			return true, nil, e
		}
		if !ok2 || !ok {
			return ok, v, nil
		}
	}

	return true, nil, nil
}

func (t *PuppetoInterpreter) VisitSwitchCase(ctx *lib.SwitchCaseContext, value any) (bool, PValue, PError, bool) {
	v, e := t.VisitExpression(ctx.Expression().(*lib.ExpressionContext))

	if e != nil {
		return true, nil, e, false
	}

	if v == value {
		return t.VisitProgramBodyWithBreakContinue(ctx.ProgramBodyWithBreakContinue().(*lib.ProgramBodyWithBreakContinueContext))
	}

	return true, nil, nil, false
}

func (t *PuppetoInterpreter) VisitBreakStatement(ctx *lib.BreakStatementContext) PError {
	panic("not necessary")
}

func (t *PuppetoInterpreter) VisitContinueStatement(ctx *lib.ContinueStatementContext) PError {
	panic("not necessary")
}

func (t *PuppetoInterpreter) VisitTryStatement(ctx *lib.TryStatementContext) (bool, PValue, PError) {
	programBodies := ctx.AllProgramBody()
	ok, v, e := t.VisitProgramBody(programBodies[0].(*lib.ProgramBodyContext))

	if e != nil && len(programBodies) > 1 {
		if ctx.IDENTIFIER() != nil {
			t = NewPuppetoInterpreter(t.Parser, t.File, NewScope(t.Scope), t.Stack, t.OutFunc)
			_ = t.Scope.Define(ctx.IDENTIFIER().GetText(), e)
		}

		ok, v, e := t.VisitProgramBody(programBodies[1].(*lib.ProgramBodyContext))

		if e != nil {
			return ok, nil, e
		}
		if !ok {
			return ok, v, nil
		}
	}
	if !ok {
		return ok, v, nil
	}

	return true, nil, nil
}

func (t *PuppetoInterpreter) VisitEchoStatement(ctx *lib.EchoStatementContext) PError {
	v, e := t.VisitExpression(ctx.Expression().(*lib.ExpressionContext))

	if e != nil {
		return e
	}

	t.OutFunc(ToStr(v))

	return nil
}

func (t *PuppetoInterpreter) VisitStatement(ctx *lib.StatementContext) (bool, PValue, PError) {
	stGtList := ctx.StatementList()

	if stGtList != nil {
		ok, v, e := t.VisitStatementList(stGtList.(*lib.StatementListContext))

		if e != nil {
			return ok, nil, e
		}
		if !ok {
			return ok, v, nil
		}
	} else {
		htmlLists := ctx.AllHtmlList()

		if len(htmlLists) > 0 {
			t.OutFunc(
				t.Parser.GetTokenStream().GetTextFromTokens(
					htmlLists[0].GetChild(0).(*antlr.TerminalNodeImpl).GetSymbol(),
					htmlLists[len(htmlLists)-1].GetChild(0).(*antlr.TerminalNodeImpl).GetSymbol(),
				),
			)
		}
	}

	return true, nil, nil
}

func (t *PuppetoInterpreter) VisitReturnStatement(ctx *lib.ReturnStatementContext) (PValue, PError) {
	if ctx.Expression() != nil {
		return t.VisitExpression(ctx.Expression().(*lib.ExpressionContext))
	}

	return nil, nil
}

func (t *PuppetoInterpreter) VisitStatementList(ctx *lib.StatementListContext) (bool, PValue, PError) {
	switch node := ctx.GetChild(0).(type) {
	case *lib.VariableDefinitionContext:
		return true, nil, t.VisitVariableDefinition(node)
	case *lib.FunctionDefinitionContext:
		v, e := t.VisitFunctionDefinition(node, false)

		return true, v, e
	case *lib.ReturnStatementContext:
		v, e := t.VisitReturnStatement(node)

		return false, v, e
	case *lib.IfStatementContext:
		return t.VisitIfStatement(node)
	case *lib.SwitchStatementContext:
		return t.VisitSwitchStatement(node)
	case *lib.TryStatementContext:
		return t.VisitTryStatement(node)
	case *lib.LoopStatementContext:
		return t.VisitLoopStatement(node)
	case *lib.EchoStatementContext:
		return true, nil, t.VisitEchoStatement(node)
	case *lib.VariableAssignationContext:
		_, e := t.VisitVariableAssignation(node)

		return true, nil, e
	}

	panic("Statement")
}

func (t *PuppetoInterpreter) VisitScope(ctx *lib.ScopeContext) (bool, PValue, PError) {
	return t.VisitProgramBodyList(ctx.ProgramBodyList().(*lib.ProgramBodyListContext))
}

func (t *PuppetoInterpreter) VisitProgramBody(ctx *lib.ProgramBodyContext) (bool, PValue, PError) {
	if ctx.Statement() != nil {
		return t.VisitStatement(ctx.Statement().(*lib.StatementContext))
	} else if ctx.Scope() != nil {
		return t.VisitScope(ctx.Scope().(*lib.ScopeContext))
	} else if ctx.Expression() != nil {
		_, e := t.VisitExpression(ctx.Expression().(*lib.ExpressionContext))

		return true, nil, e
	}

	panic("ProgramBody")
}

func (t *PuppetoInterpreter) VisitProgramBodyList(ctx *lib.ProgramBodyListContext) (bool, PValue, PError) {
	var v any
	var e any
	var ok bool = true

	for _, programBody := range ctx.AllProgramBody() {
		ok, v, e = t.VisitProgramBody(programBody.(*lib.ProgramBodyContext))

		if e != nil {
			return ok, nil, e
		}
		if !ok {
			return ok, v, e
		}
	}

	return ok, v, e
}

func (t *PuppetoInterpreter) VisitProgram(ctx *lib.ProgramContext) (PValue, PError) {
	if ctx.ProgramBodyList() == nil {
		return nil, nil
	}

	_, v, e := t.VisitProgramBodyList(ctx.ProgramBodyList().(*lib.ProgramBodyListContext))

	return v, e
}

func (t *PuppetoInterpreter) VisitPupCode(ctx *lib.PupCodeContext) (bool, PValue, PError) {
	programBodyList := ctx.ProgramBodyList()

	if programBodyList == nil {
		return true, nil, nil
	}

	return t.VisitProgramBodyList(programBodyList.(*lib.ProgramBodyListContext))
}

func (t *PuppetoInterpreter) VisitPupShortCode(ctx *lib.PupShortCodeContext) PError {
	expI := ctx.Expression()

	if expI == nil {
		return nil
	}

	v, e := t.VisitExpression(expI.(*lib.ExpressionContext))

	if e != nil {
		return e
	}

	t.OutFunc(ToStr(v))

	return nil
}

func (t *PuppetoInterpreter) VisitHtml(ctx *lib.HtmlContext) error {
	txts := []antlr.Token{}
	charStream := t.Parser.GetTokenStream()
	min := 0

	for _, node := range ctx.GetChildren() {
		switch node := node.(type) {
		case *antlr.TerminalNodeImpl:
			txts = append(txts, node.GetSymbol())
		case *antlr.ErrorNodeImpl:
			txts = append(txts, node.GetSymbol())
		case *lib.PupShortCodeContext:
			t.VisitPupShortCode(node)
		case *lib.HtmlListContext:
			txts = append(txts, node.GetChild(0).(*antlr.TerminalNodeImpl).GetSymbol())
		case *lib.PupCodeContext:
			if len(txts) > 0 {
				tx := charStream.GetTextFromTokens(txts[0], txts[len(txts)-1])
				t.OutFunc(tx[min:])

				txts = []antlr.Token{}
			}

			txts = append(txts, node.GetStop())
			min = len(node.GetStop().GetText())

			ok, _, e := t.VisitPupCode(node)

			if e != nil {
				return fmt.Errorf(StringifyError(append(Stack{t.TokenToPosition(node.GetStart())}, t.Stack...), e))
			}
			if !ok {
				return nil
			}
		}
	}

	if len(txts) > 0 {
		tx := charStream.GetTextFromTokens(txts[0], txts[len(txts)-1])
		t.OutFunc(tx[min:])
	}

	return nil
}

func DefaultReadFileFunc(stack Stack, specifier string) (string, any) {
	panic("s")
}

func NewGlobalScope(rootDir string, readFile ReadFileFunc) *Scope {
	scope := NewScope(nil)

	_ = scope.Define("typeof", func(_ Stack, args []any) (any, any) {
		return TypeOf(args[0]), nil
	})
	_ = scope.Define("is_nan", func(_ Stack, args []any) (any, any) {
		f, ok := args[0].(float64)

		if !ok {
			return false, nil
		}

		return math.IsNaN(f), nil
	})
	_ = scope.Define("inspect", func(_ Stack, args []any) (any, any) {
		return Inspect(args[0]), nil
	})
	_ = scope.Define("to_string", func(_ Stack, args []any) (any, any) {
		return ToStr(args[0]), nil
	})
	_ = scope.Define("to_integer", func(_ Stack, args []any) (any, any) {
		return ToInt(args[0]), nil
	})
	_ = scope.Define("to_boolean", func(_ Stack, args []any) (any, any) {
		return ToBool(args[0]), nil
	})
	_ = scope.Define("parse_float", func(_ Stack, args []any) (any, any) {
		return StrToF(ToStr(args[0])), nil
	})
	_ = scope.Define("parse_integer", func(_ Stack, args []any) (any, any) {
		return StrToInt(ToStr(args[0])), nil
	})
	_ = scope.Define("parse_error", func(stack Stack, args []any) (any, any) {
		return StringifyError(stack, args[0]), nil
	})
	_ = scope.Define("import_string_as_file", func(_ Stack, args []any) (any, any) {
		newScope := NewScope(scope)
		return ImportStr(newScope, JoinPath(rootDir, ToStr(args[0])), ToStr(args[1]))
	})
	_ = scope.Define("import_file", func(stack Stack, args []any) (any, any) {
		newScope := NewScope(scope)
		return ImportFileStr(newScope, JoinPath(rootDir, ToStr(args[0])), readFile, stack)
	})
	_ = scope.Define("array_join", func(stack Stack, args []any) (any, any) {
		ss := []string{}
		arr, ok := args[0].(PArray)

		if !ok {
			return nil, NewPError(TypeError, "Expected array #0", stack)
		}

		sep, ok := args[1].(string)

		if !ok {
			return nil, NewPError(TypeError, "Expected string #1", stack)
		}

		for _, v := range arr {
			ss = append(ss, ToStr(v))
		}

		return strings.Join(ss, sep), nil
	})
	_ = scope.Define("array_get_length", func(stack Stack, args []any) (any, any) {
		arr, ok := args[0].(PArray)

		if !ok {
			return nil, NewPError(TypeError, "Expected array #0", stack)
		}

		return int64(len(arr)), nil
	})
	_ = scope.Define("object_get_entries", func(stack Stack, args []any) (any, any) {
		o, ok := args[0].(PObject)

		if !ok {
			return nil, NewPError(TypeError, "Expected object #0", stack)
		}

		a := [][]any{}

		for k, v := range o {
			a = append(a, []any{k, v})
		}

		return a, nil
	})
	_ = scope.Define("object_get_keys", func(stack Stack, args []any) (any, any) {
		o, ok := args[0].(PObject)

		if !ok {
			return nil, NewPError(TypeError, "Expected object #0", stack)
		}

		a := []string{}

		for k := range o {
			a = append(a, k)
		}

		return a, nil
	})
	_ = scope.Define("object_get_values", func(stack Stack, args []any) (any, any) {
		o, ok := args[0].(PObject)

		if !ok {
			return nil, NewPError(TypeError, "Expected object #0", stack)
		}

		a := []any{}

		for _, v := range o {
			a = append(a, v)
		}

		return a, nil
	})

	return scope
}

func RunString(scope *Scope, file string, code string) (PValue, PError) {
	input := antlr.NewInputStream(code)
	lexer := lib.NewPuppetoLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := lib.NewPuppetoParser(stream)
	errorListener := NewPuppetoErrorListener()
	p.RemoveErrorListeners()
	p.AddErrorListener(errorListener)
	p.BuildParseTrees = true
	treeI := p.Program()

	if errorListener.CurrentError != nil {
		return nil, errorListener.CurrentError.Error()
	}

	interpreter := NewPuppetoInterpreter(p, file, scope, Stack{}, DefaultOutFunc)
	return interpreter.VisitProgram(treeI.(*lib.ProgramContext))

}

func RunFile(scope *Scope, file string) (PValue, PError) {
	input, e := antlr.NewFileStream(file)

	if e != nil {
		return nil, NewPError(TypeError, e.Error(), Stack{})
	}

	lexer := lib.NewPuppetoLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := lib.NewPuppetoParser(stream)
	errorListener := NewPuppetoErrorListener()
	p.RemoveErrorListeners()
	p.AddErrorListener(errorListener)
	p.BuildParseTrees = true
	treeI := p.Program()

	if errorListener.CurrentError != nil {
		return nil, errorListener.CurrentError.Error()
	}

	interpreter := NewPuppetoInterpreter(p, file, scope, Stack{}, DefaultOutFunc)
	return interpreter.VisitProgram(treeI.(*lib.ProgramContext))

}

func RunHTMLString(scope *Scope, file string, code string) (string, error) {
	input := antlr.NewInputStream(code)
	lexer := lib.NewPuppetoLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := lib.NewPuppetoParser(stream)
	errorListener := NewPuppetoErrorListener()
	p.RemoveErrorListeners()
	p.AddErrorListener(errorListener)
	p.BuildParseTrees = true
	o := ""
	htmlI := p.Html()

	if errorListener.CurrentError != nil {
		return "", errorListener.CurrentError
	}

	interpreter := NewPuppetoInterpreter(p, file, scope, Stack{}, func(s string) {
		o += s
	})
	e := interpreter.VisitHtml(htmlI.(*lib.HtmlContext))

	return o, e
}

func RunHTMLFile(scope *Scope, file string) (string, error) {
	input, e := antlr.NewFileStream(file)

	if e != nil {
		return "", e
	}

	lexer := lib.NewPuppetoLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := lib.NewPuppetoParser(stream)
	errorListener := NewPuppetoErrorListener()
	p.RemoveErrorListeners()
	p.AddErrorListener(errorListener)
	p.BuildParseTrees = true
	o := ""
	htmlI := p.Html()

	if errorListener.CurrentError != nil {
		return "", errorListener.CurrentError
	}

	interpreter := NewPuppetoInterpreter(p, file, scope, Stack{}, func(s string) {
		o += s
	})
	e = interpreter.VisitHtml(htmlI.(*lib.HtmlContext))

	return o, e
}
