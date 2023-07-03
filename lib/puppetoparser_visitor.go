// Code generated from /home/reiyo/Project/Go/MarionetteEnsemble/Puppeto/grammars/PuppetoParser.g4 by ANTLR 4.12.0. DO NOT EDIT.

package lib // PuppetoParser
import "github.com/antlr4-go/antlr"

// A complete Visitor for a parse tree produced by PuppetoParser.
type PuppetoParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by PuppetoParser#literalValue.
	VisitLiteralValue(ctx *LiteralValueContext) interface{}

	// Visit a parse tree produced by PuppetoParser#object.
	VisitObject(ctx *ObjectContext) interface{}

	// Visit a parse tree produced by PuppetoParser#pair.
	VisitPair(ctx *PairContext) interface{}

	// Visit a parse tree produced by PuppetoParser#array.
	VisitArray(ctx *ArrayContext) interface{}

	// Visit a parse tree produced by PuppetoParser#functionDefinition.
	VisitFunctionDefinition(ctx *FunctionDefinitionContext) interface{}

	// Visit a parse tree produced by PuppetoParser#argumentList.
	VisitArgumentList(ctx *ArgumentListContext) interface{}

	// Visit a parse tree produced by PuppetoParser#argument.
	VisitArgument(ctx *ArgumentContext) interface{}

	// Visit a parse tree produced by PuppetoParser#none.
	VisitNone(ctx *NoneContext) interface{}

	// Visit a parse tree produced by PuppetoParser#scopeBody.
	VisitScopeBody(ctx *ScopeBodyContext) interface{}

	// Visit a parse tree produced by PuppetoParser#value.
	VisitValue(ctx *ValueContext) interface{}

	// Visit a parse tree produced by PuppetoParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by PuppetoParser#logicalExpression.
	VisitLogicalExpression(ctx *LogicalExpressionContext) interface{}

	// Visit a parse tree produced by PuppetoParser#equalityExpression.
	VisitEqualityExpression(ctx *EqualityExpressionContext) interface{}

	// Visit a parse tree produced by PuppetoParser#comparisonExpression.
	VisitComparisonExpression(ctx *ComparisonExpressionContext) interface{}

	// Visit a parse tree produced by PuppetoParser#additiveExpression.
	VisitAdditiveExpression(ctx *AdditiveExpressionContext) interface{}

	// Visit a parse tree produced by PuppetoParser#multiplicativeExpression.
	VisitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) interface{}

	// Visit a parse tree produced by PuppetoParser#powerExpression.
	VisitPowerExpression(ctx *PowerExpressionContext) interface{}

	// Visit a parse tree produced by PuppetoParser#bitwiseExpression.
	VisitBitwiseExpression(ctx *BitwiseExpressionContext) interface{}

	// Visit a parse tree produced by PuppetoParser#shiftExpression.
	VisitShiftExpression(ctx *ShiftExpressionContext) interface{}

	// Visit a parse tree produced by PuppetoParser#unaryExpression.
	VisitUnaryExpression(ctx *UnaryExpressionContext) interface{}

	// Visit a parse tree produced by PuppetoParser#getProperty.
	VisitGetProperty(ctx *GetPropertyContext) interface{}

	// Visit a parse tree produced by PuppetoParser#setProperty.
	VisitSetProperty(ctx *SetPropertyContext) interface{}

	// Visit a parse tree produced by PuppetoParser#callExpression.
	VisitCallExpression(ctx *CallExpressionContext) interface{}

	// Visit a parse tree produced by PuppetoParser#chainExpression.
	VisitChainExpression(ctx *ChainExpressionContext) interface{}

	// Visit a parse tree produced by PuppetoParser#lhsVariableAssignation.
	VisitLhsVariableAssignation(ctx *LhsVariableAssignationContext) interface{}

	// Visit a parse tree produced by PuppetoParser#rhsVariableAssignation.
	VisitRhsVariableAssignation(ctx *RhsVariableAssignationContext) interface{}

	// Visit a parse tree produced by PuppetoParser#midVariableAssignation.
	VisitMidVariableAssignation(ctx *MidVariableAssignationContext) interface{}

	// Visit a parse tree produced by PuppetoParser#variableAssignation.
	VisitVariableAssignation(ctx *VariableAssignationContext) interface{}

	// Visit a parse tree produced by PuppetoParser#variableDefinition.
	VisitVariableDefinition(ctx *VariableDefinitionContext) interface{}

	// Visit a parse tree produced by PuppetoParser#variableDefinitionBody.
	VisitVariableDefinitionBody(ctx *VariableDefinitionBodyContext) interface{}

	// Visit a parse tree produced by PuppetoParser#ifStatement.
	VisitIfStatement(ctx *IfStatementContext) interface{}

	// Visit a parse tree produced by PuppetoParser#loopStatement.
	VisitLoopStatement(ctx *LoopStatementContext) interface{}

	// Visit a parse tree produced by PuppetoParser#switchStatement.
	VisitSwitchStatement(ctx *SwitchStatementContext) interface{}

	// Visit a parse tree produced by PuppetoParser#switchCase.
	VisitSwitchCase(ctx *SwitchCaseContext) interface{}

	// Visit a parse tree produced by PuppetoParser#breakStatement.
	VisitBreakStatement(ctx *BreakStatementContext) interface{}

	// Visit a parse tree produced by PuppetoParser#continueStatement.
	VisitContinueStatement(ctx *ContinueStatementContext) interface{}

	// Visit a parse tree produced by PuppetoParser#tryStatement.
	VisitTryStatement(ctx *TryStatementContext) interface{}

	// Visit a parse tree produced by PuppetoParser#returnStatement.
	VisitReturnStatement(ctx *ReturnStatementContext) interface{}

	// Visit a parse tree produced by PuppetoParser#statementList.
	VisitStatementList(ctx *StatementListContext) interface{}

	// Visit a parse tree produced by PuppetoParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by PuppetoParser#scope.
	VisitScope(ctx *ScopeContext) interface{}

	// Visit a parse tree produced by PuppetoParser#scopeWithBreakContinue.
	VisitScopeWithBreakContinue(ctx *ScopeWithBreakContinueContext) interface{}

	// Visit a parse tree produced by PuppetoParser#programBody.
	VisitProgramBody(ctx *ProgramBodyContext) interface{}

	// Visit a parse tree produced by PuppetoParser#programBodyList.
	VisitProgramBodyList(ctx *ProgramBodyListContext) interface{}

	// Visit a parse tree produced by PuppetoParser#echoStatement.
	VisitEchoStatement(ctx *EchoStatementContext) interface{}

	// Visit a parse tree produced by PuppetoParser#programBodyWithBreakContinue.
	VisitProgramBodyWithBreakContinue(ctx *ProgramBodyWithBreakContinueContext) interface{}

	// Visit a parse tree produced by PuppetoParser#pupShortCode.
	VisitPupShortCode(ctx *PupShortCodeContext) interface{}

	// Visit a parse tree produced by PuppetoParser#pupCode.
	VisitPupCode(ctx *PupCodeContext) interface{}

	// Visit a parse tree produced by PuppetoParser#htmlList.
	VisitHtmlList(ctx *HtmlListContext) interface{}

	// Visit a parse tree produced by PuppetoParser#html.
	VisitHtml(ctx *HtmlContext) interface{}

	// Visit a parse tree produced by PuppetoParser#program.
	VisitProgram(ctx *ProgramContext) interface{}
}
