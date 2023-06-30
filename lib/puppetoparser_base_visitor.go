// Code generated from /home/reiyo/Project/Go/MarionetteEnsemble/Puppeto/grammars/PuppetoParser.g4 by ANTLR 4.12.0. DO NOT EDIT.

package lib // PuppetoParser
import "github.com/antlr4-go/antlr"

type BasePuppetoParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasePuppetoParserVisitor) VisitLiteralValue(ctx *LiteralValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePuppetoParserVisitor) VisitObject(ctx *ObjectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePuppetoParserVisitor) VisitPair(ctx *PairContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePuppetoParserVisitor) VisitArray(ctx *ArrayContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePuppetoParserVisitor) VisitFunctionDefinition(ctx *FunctionDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePuppetoParserVisitor) VisitArgumentList(ctx *ArgumentListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePuppetoParserVisitor) VisitArgument(ctx *ArgumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePuppetoParserVisitor) VisitNone(ctx *NoneContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePuppetoParserVisitor) VisitScopeBody(ctx *ScopeBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePuppetoParserVisitor) VisitValue(ctx *ValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePuppetoParserVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePuppetoParserVisitor) VisitGetProperty(ctx *GetPropertyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePuppetoParserVisitor) VisitSetProperty(ctx *SetPropertyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePuppetoParserVisitor) VisitCallExpression(ctx *CallExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePuppetoParserVisitor) VisitChainExpression(ctx *ChainExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePuppetoParserVisitor) VisitLogicalExpression(ctx *LogicalExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePuppetoParserVisitor) VisitEqualityExpression(ctx *EqualityExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePuppetoParserVisitor) VisitComparisonExpression(ctx *ComparisonExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePuppetoParserVisitor) VisitAdditiveExpression(ctx *AdditiveExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePuppetoParserVisitor) VisitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePuppetoParserVisitor) VisitPowerExpression(ctx *PowerExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePuppetoParserVisitor) VisitBitwiseExpression(ctx *BitwiseExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePuppetoParserVisitor) VisitShiftExpression(ctx *ShiftExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePuppetoParserVisitor) VisitUnaryExpression(ctx *UnaryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePuppetoParserVisitor) VisitLhsVariableAssignation(ctx *LhsVariableAssignationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePuppetoParserVisitor) VisitRhsVariableAssignation(ctx *RhsVariableAssignationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePuppetoParserVisitor) VisitMidVariableAssignation(ctx *MidVariableAssignationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePuppetoParserVisitor) VisitVariableAssignation(ctx *VariableAssignationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePuppetoParserVisitor) VisitVariableDefinition(ctx *VariableDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePuppetoParserVisitor) VisitVariableDefinitionBody(ctx *VariableDefinitionBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePuppetoParserVisitor) VisitIfStatement(ctx *IfStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePuppetoParserVisitor) VisitLoopStatement(ctx *LoopStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePuppetoParserVisitor) VisitSwitchStatement(ctx *SwitchStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePuppetoParserVisitor) VisitSwitchCase(ctx *SwitchCaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePuppetoParserVisitor) VisitBreakStatement(ctx *BreakStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePuppetoParserVisitor) VisitContinueStatement(ctx *ContinueStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePuppetoParserVisitor) VisitTryStatement(ctx *TryStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePuppetoParserVisitor) VisitStatementList(ctx *StatementListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePuppetoParserVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePuppetoParserVisitor) VisitScope(ctx *ScopeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePuppetoParserVisitor) VisitScopeWithBreakContinue(ctx *ScopeWithBreakContinueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePuppetoParserVisitor) VisitProgramBody(ctx *ProgramBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePuppetoParserVisitor) VisitProgramBodyList(ctx *ProgramBodyListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePuppetoParserVisitor) VisitEchoStatement(ctx *EchoStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePuppetoParserVisitor) VisitProgramBodyWithBreakContinue(ctx *ProgramBodyWithBreakContinueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePuppetoParserVisitor) VisitPupCode(ctx *PupCodeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePuppetoParserVisitor) VisitHtmlList(ctx *HtmlListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePuppetoParserVisitor) VisitHtml(ctx *HtmlContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePuppetoParserVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}
