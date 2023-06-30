// Code generated from /home/reiyo/Project/Go/MarionetteEnsemble/Puppeto/grammars/PuppetoParser.g4 by ANTLR 4.12.0. DO NOT EDIT.

package lib // PuppetoParser
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type PuppetoParser struct {
	*antlr.BaseParser
}

var puppetoparserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func puppetoparserParserInit() {
	staticData := &puppetoparserParserStaticData
	staticData.literalNames = []string{
		"", "'...'", "'nan'", "'<'", "'>'", "'<?pup'", "'?>'", "", "", "'=='",
		"'!='", "'<='", "'>='", "'for'", "'='", "'echo'", "'{'", "'}'", "','",
		"'['", "']'", "'fn'", "'('", "')'", "'&&'", "'||'", "':'", "'.'", "'+'",
		"'-'", "'**'", "'%'", "'&'", "'|'", "'^'", "'<<'", "'>>'", "'!'", "'if'",
		"'else'", "'switch'", "'case'", "'break'", "'try'", "'continue'", "'catch'",
		"'default'", "'let'", "'*'", "'/'", "", "", "'nil'",
	}
	staticData.symbolicNames = []string{
		"", "SPREAD", "NAN", "LT", "GT", "PUP_START_TAG", "PUP_END_TAG", "TWO_SIDES_OPERATOR",
		"ONE_SIDE_OPERATOR", "EQ", "NEQ", "LTOE", "GTOE", "FOR", "TO", "ECHO",
		"L_CURLY", "R_CURLY", "COMMA", "L_SQUARE", "R_SQUARE", "FN", "L_PARENTHES",
		"R_PARENTHES", "AND", "OR", "COLON", "DOT", "PLUS", "MINUS", "DOUBLE_STAR",
		"PERCENT", "BAND", "BOR", "BXOR", "BLS", "BRS", "NOT", "IF", "ELSE",
		"SWITCH", "CASE", "BREAK", "TRY", "CONTINUE", "CATCH", "DEFAULT", "LET",
		"STAR", "SLASH", "FLOAT", "INTEGER", "NIL", "BOOLEAN", "IDENTIFIER",
		"STRING", "WS", "SEMICOLON", "HTML",
	}
	staticData.ruleNames = []string{
		"literalValue", "object", "pair", "array", "functionDefinition", "argumentList",
		"argument", "none", "scopeBody", "value", "expression", "logicalExpression",
		"equalityExpression", "comparisonExpression", "additiveExpression",
		"multiplicativeExpression", "powerExpression", "bitwiseExpression",
		"shiftExpression", "unaryExpression", "getProperty", "setProperty",
		"callExpression", "chainExpression", "lhsVariableAssignation", "rhsVariableAssignation",
		"midVariableAssignation", "variableAssignation", "variableDefinition",
		"variableDefinitionBody", "ifStatement", "loopStatement", "switchStatement",
		"switchCase", "breakStatement", "continueStatement", "tryStatement",
		"statementList", "statement", "scope", "scopeWithBreakContinue", "programBody",
		"programBodyList", "echoStatement", "programBodyWithBreakContinue",
		"pupCode", "htmlList", "html", "program",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 58, 477, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46, 2, 47,
		7, 47, 2, 48, 7, 48, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 105, 8,
		1, 10, 1, 12, 1, 108, 9, 1, 3, 1, 110, 8, 1, 1, 1, 3, 1, 113, 8, 1, 1,
		1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 123, 8, 2, 1, 2, 1,
		2, 1, 2, 3, 2, 128, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 134, 8, 3, 10,
		3, 12, 3, 137, 9, 3, 1, 3, 3, 3, 140, 8, 3, 3, 3, 142, 8, 3, 1, 3, 1, 3,
		1, 4, 1, 4, 3, 4, 148, 8, 4, 1, 4, 1, 4, 3, 4, 152, 8, 4, 1, 4, 1, 4, 1,
		4, 1, 5, 1, 5, 1, 5, 5, 5, 160, 8, 5, 10, 5, 12, 5, 163, 9, 5, 1, 5, 3,
		5, 166, 8, 5, 1, 5, 1, 5, 3, 5, 170, 8, 5, 1, 6, 1, 6, 1, 6, 3, 6, 175,
		8, 6, 1, 6, 3, 6, 178, 8, 6, 1, 7, 1, 7, 1, 8, 1, 8, 3, 8, 184, 8, 8, 1,
		9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 194, 8, 9, 1, 10, 1,
		10, 1, 11, 1, 11, 1, 11, 5, 11, 201, 8, 11, 10, 11, 12, 11, 204, 9, 11,
		1, 12, 1, 12, 1, 12, 5, 12, 209, 8, 12, 10, 12, 12, 12, 212, 9, 12, 1,
		13, 1, 13, 1, 13, 5, 13, 217, 8, 13, 10, 13, 12, 13, 220, 9, 13, 1, 14,
		1, 14, 1, 14, 5, 14, 225, 8, 14, 10, 14, 12, 14, 228, 9, 14, 1, 15, 1,
		15, 1, 15, 5, 15, 233, 8, 15, 10, 15, 12, 15, 236, 9, 15, 1, 16, 1, 16,
		1, 16, 5, 16, 241, 8, 16, 10, 16, 12, 16, 244, 9, 16, 1, 17, 1, 17, 1,
		17, 5, 17, 249, 8, 17, 10, 17, 12, 17, 252, 9, 17, 1, 18, 1, 18, 1, 18,
		5, 18, 257, 8, 18, 10, 18, 12, 18, 260, 9, 18, 1, 19, 1, 19, 1, 19, 3,
		19, 265, 8, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 3, 20,
		274, 8, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1,
		22, 5, 22, 285, 8, 22, 10, 22, 12, 22, 288, 9, 22, 1, 22, 3, 22, 291, 8,
		22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 3, 23, 299, 8, 23, 1, 24,
		1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 27, 1,
		27, 1, 27, 4, 27, 314, 8, 27, 11, 27, 12, 27, 315, 1, 28, 1, 28, 1, 28,
		1, 28, 5, 28, 322, 8, 28, 10, 28, 12, 28, 325, 9, 28, 1, 29, 1, 29, 1,
		29, 3, 29, 330, 8, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30,
		1, 30, 5, 30, 340, 8, 30, 10, 30, 12, 30, 343, 9, 30, 1, 30, 1, 30, 3,
		30, 347, 8, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 1, 32,
		5, 32, 357, 8, 32, 10, 32, 12, 32, 360, 9, 32, 1, 32, 1, 32, 1, 32, 5,
		32, 365, 8, 32, 10, 32, 12, 32, 368, 9, 32, 3, 32, 370, 8, 32, 1, 32, 1,
		32, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 34, 1, 34, 1, 35, 1, 35, 1, 36,
		1, 36, 1, 36, 1, 36, 3, 36, 387, 8, 36, 1, 36, 3, 36, 390, 8, 36, 1, 37,
		1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 3, 37, 400, 8, 37, 1,
		38, 1, 38, 1, 38, 5, 38, 405, 8, 38, 10, 38, 12, 38, 408, 9, 38, 1, 38,
		4, 38, 411, 8, 38, 11, 38, 12, 38, 412, 3, 38, 415, 8, 38, 1, 39, 1, 39,
		3, 39, 419, 8, 39, 1, 39, 1, 39, 1, 40, 1, 40, 5, 40, 425, 8, 40, 10, 40,
		12, 40, 428, 9, 40, 1, 40, 1, 40, 1, 41, 1, 41, 1, 41, 3, 41, 435, 8, 41,
		1, 42, 4, 42, 438, 8, 42, 11, 42, 12, 42, 439, 1, 43, 1, 43, 1, 43, 1,
		44, 1, 44, 1, 44, 1, 44, 1, 44, 3, 44, 450, 8, 44, 1, 45, 1, 45, 3, 45,
		454, 8, 45, 1, 45, 1, 45, 1, 46, 1, 46, 1, 47, 1, 47, 4, 47, 462, 8, 47,
		11, 47, 12, 47, 463, 1, 47, 1, 47, 1, 48, 1, 48, 3, 48, 470, 8, 48, 1,
		48, 3, 48, 473, 8, 48, 1, 48, 1, 48, 1, 48, 0, 0, 49, 0, 2, 4, 6, 8, 10,
		12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46,
		48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82,
		84, 86, 88, 90, 92, 94, 96, 0, 11, 2, 0, 2, 2, 50, 55, 1, 0, 24, 25, 1,
		0, 9, 10, 2, 0, 3, 4, 11, 12, 1, 0, 28, 29, 2, 0, 31, 31, 48, 49, 1, 0,
		32, 34, 1, 0, 35, 36, 2, 0, 29, 29, 37, 37, 2, 0, 7, 7, 14, 14, 2, 0, 3,
		4, 7, 58, 499, 0, 98, 1, 0, 0, 0, 2, 100, 1, 0, 0, 0, 4, 127, 1, 0, 0,
		0, 6, 129, 1, 0, 0, 0, 8, 145, 1, 0, 0, 0, 10, 156, 1, 0, 0, 0, 12, 171,
		1, 0, 0, 0, 14, 179, 1, 0, 0, 0, 16, 183, 1, 0, 0, 0, 18, 193, 1, 0, 0,
		0, 20, 195, 1, 0, 0, 0, 22, 197, 1, 0, 0, 0, 24, 205, 1, 0, 0, 0, 26, 213,
		1, 0, 0, 0, 28, 221, 1, 0, 0, 0, 30, 229, 1, 0, 0, 0, 32, 237, 1, 0, 0,
		0, 34, 245, 1, 0, 0, 0, 36, 253, 1, 0, 0, 0, 38, 264, 1, 0, 0, 0, 40, 266,
		1, 0, 0, 0, 42, 275, 1, 0, 0, 0, 44, 279, 1, 0, 0, 0, 46, 298, 1, 0, 0,
		0, 48, 300, 1, 0, 0, 0, 50, 303, 1, 0, 0, 0, 52, 306, 1, 0, 0, 0, 54, 313,
		1, 0, 0, 0, 56, 317, 1, 0, 0, 0, 58, 326, 1, 0, 0, 0, 60, 331, 1, 0, 0,
		0, 62, 348, 1, 0, 0, 0, 64, 352, 1, 0, 0, 0, 66, 373, 1, 0, 0, 0, 68, 378,
		1, 0, 0, 0, 70, 380, 1, 0, 0, 0, 72, 382, 1, 0, 0, 0, 74, 399, 1, 0, 0,
		0, 76, 414, 1, 0, 0, 0, 78, 416, 1, 0, 0, 0, 80, 422, 1, 0, 0, 0, 82, 434,
		1, 0, 0, 0, 84, 437, 1, 0, 0, 0, 86, 441, 1, 0, 0, 0, 88, 449, 1, 0, 0,
		0, 90, 451, 1, 0, 0, 0, 92, 457, 1, 0, 0, 0, 94, 461, 1, 0, 0, 0, 96, 467,
		1, 0, 0, 0, 98, 99, 7, 0, 0, 0, 99, 1, 1, 0, 0, 0, 100, 109, 5, 16, 0,
		0, 101, 106, 3, 4, 2, 0, 102, 103, 5, 18, 0, 0, 103, 105, 3, 4, 2, 0, 104,
		102, 1, 0, 0, 0, 105, 108, 1, 0, 0, 0, 106, 104, 1, 0, 0, 0, 106, 107,
		1, 0, 0, 0, 107, 110, 1, 0, 0, 0, 108, 106, 1, 0, 0, 0, 109, 101, 1, 0,
		0, 0, 109, 110, 1, 0, 0, 0, 110, 112, 1, 0, 0, 0, 111, 113, 5, 18, 0, 0,
		112, 111, 1, 0, 0, 0, 112, 113, 1, 0, 0, 0, 113, 114, 1, 0, 0, 0, 114,
		115, 5, 17, 0, 0, 115, 3, 1, 0, 0, 0, 116, 117, 5, 19, 0, 0, 117, 118,
		3, 20, 10, 0, 118, 119, 5, 20, 0, 0, 119, 123, 1, 0, 0, 0, 120, 123, 5,
		54, 0, 0, 121, 123, 5, 55, 0, 0, 122, 116, 1, 0, 0, 0, 122, 120, 1, 0,
		0, 0, 122, 121, 1, 0, 0, 0, 123, 124, 1, 0, 0, 0, 124, 125, 5, 26, 0, 0,
		125, 128, 3, 20, 10, 0, 126, 128, 5, 54, 0, 0, 127, 122, 1, 0, 0, 0, 127,
		126, 1, 0, 0, 0, 128, 5, 1, 0, 0, 0, 129, 141, 5, 19, 0, 0, 130, 135, 3,
		20, 10, 0, 131, 132, 5, 18, 0, 0, 132, 134, 3, 20, 10, 0, 133, 131, 1,
		0, 0, 0, 134, 137, 1, 0, 0, 0, 135, 133, 1, 0, 0, 0, 135, 136, 1, 0, 0,
		0, 136, 139, 1, 0, 0, 0, 137, 135, 1, 0, 0, 0, 138, 140, 5, 18, 0, 0, 139,
		138, 1, 0, 0, 0, 139, 140, 1, 0, 0, 0, 140, 142, 1, 0, 0, 0, 141, 130,
		1, 0, 0, 0, 141, 142, 1, 0, 0, 0, 142, 143, 1, 0, 0, 0, 143, 144, 5, 20,
		0, 0, 144, 7, 1, 0, 0, 0, 145, 147, 5, 21, 0, 0, 146, 148, 5, 54, 0, 0,
		147, 146, 1, 0, 0, 0, 147, 148, 1, 0, 0, 0, 148, 149, 1, 0, 0, 0, 149,
		151, 5, 22, 0, 0, 150, 152, 3, 10, 5, 0, 151, 150, 1, 0, 0, 0, 151, 152,
		1, 0, 0, 0, 152, 153, 1, 0, 0, 0, 153, 154, 5, 23, 0, 0, 154, 155, 3, 16,
		8, 0, 155, 9, 1, 0, 0, 0, 156, 161, 3, 12, 6, 0, 157, 158, 5, 18, 0, 0,
		158, 160, 3, 12, 6, 0, 159, 157, 1, 0, 0, 0, 160, 163, 1, 0, 0, 0, 161,
		159, 1, 0, 0, 0, 161, 162, 1, 0, 0, 0, 162, 165, 1, 0, 0, 0, 163, 161,
		1, 0, 0, 0, 164, 166, 5, 18, 0, 0, 165, 164, 1, 0, 0, 0, 165, 166, 1, 0,
		0, 0, 166, 169, 1, 0, 0, 0, 167, 168, 5, 1, 0, 0, 168, 170, 5, 54, 0, 0,
		169, 167, 1, 0, 0, 0, 169, 170, 1, 0, 0, 0, 170, 11, 1, 0, 0, 0, 171, 177,
		5, 54, 0, 0, 172, 173, 5, 14, 0, 0, 173, 175, 3, 20, 10, 0, 174, 172, 1,
		0, 0, 0, 174, 175, 1, 0, 0, 0, 175, 178, 1, 0, 0, 0, 176, 178, 3, 14, 7,
		0, 177, 174, 1, 0, 0, 0, 177, 176, 1, 0, 0, 0, 178, 13, 1, 0, 0, 0, 179,
		180, 1, 0, 0, 0, 180, 15, 1, 0, 0, 0, 181, 184, 3, 20, 10, 0, 182, 184,
		3, 78, 39, 0, 183, 181, 1, 0, 0, 0, 183, 182, 1, 0, 0, 0, 184, 17, 1, 0,
		0, 0, 185, 194, 3, 0, 0, 0, 186, 194, 3, 2, 1, 0, 187, 194, 3, 6, 3, 0,
		188, 194, 3, 8, 4, 0, 189, 190, 5, 22, 0, 0, 190, 191, 3, 20, 10, 0, 191,
		192, 5, 23, 0, 0, 192, 194, 1, 0, 0, 0, 193, 185, 1, 0, 0, 0, 193, 186,
		1, 0, 0, 0, 193, 187, 1, 0, 0, 0, 193, 188, 1, 0, 0, 0, 193, 189, 1, 0,
		0, 0, 194, 19, 1, 0, 0, 0, 195, 196, 3, 22, 11, 0, 196, 21, 1, 0, 0, 0,
		197, 202, 3, 24, 12, 0, 198, 199, 7, 1, 0, 0, 199, 201, 3, 24, 12, 0, 200,
		198, 1, 0, 0, 0, 201, 204, 1, 0, 0, 0, 202, 200, 1, 0, 0, 0, 202, 203,
		1, 0, 0, 0, 203, 23, 1, 0, 0, 0, 204, 202, 1, 0, 0, 0, 205, 210, 3, 26,
		13, 0, 206, 207, 7, 2, 0, 0, 207, 209, 3, 26, 13, 0, 208, 206, 1, 0, 0,
		0, 209, 212, 1, 0, 0, 0, 210, 208, 1, 0, 0, 0, 210, 211, 1, 0, 0, 0, 211,
		25, 1, 0, 0, 0, 212, 210, 1, 0, 0, 0, 213, 218, 3, 28, 14, 0, 214, 215,
		7, 3, 0, 0, 215, 217, 3, 28, 14, 0, 216, 214, 1, 0, 0, 0, 217, 220, 1,
		0, 0, 0, 218, 216, 1, 0, 0, 0, 218, 219, 1, 0, 0, 0, 219, 27, 1, 0, 0,
		0, 220, 218, 1, 0, 0, 0, 221, 226, 3, 30, 15, 0, 222, 223, 7, 4, 0, 0,
		223, 225, 3, 30, 15, 0, 224, 222, 1, 0, 0, 0, 225, 228, 1, 0, 0, 0, 226,
		224, 1, 0, 0, 0, 226, 227, 1, 0, 0, 0, 227, 29, 1, 0, 0, 0, 228, 226, 1,
		0, 0, 0, 229, 234, 3, 32, 16, 0, 230, 231, 7, 5, 0, 0, 231, 233, 3, 32,
		16, 0, 232, 230, 1, 0, 0, 0, 233, 236, 1, 0, 0, 0, 234, 232, 1, 0, 0, 0,
		234, 235, 1, 0, 0, 0, 235, 31, 1, 0, 0, 0, 236, 234, 1, 0, 0, 0, 237, 242,
		3, 34, 17, 0, 238, 239, 5, 30, 0, 0, 239, 241, 3, 34, 17, 0, 240, 238,
		1, 0, 0, 0, 241, 244, 1, 0, 0, 0, 242, 240, 1, 0, 0, 0, 242, 243, 1, 0,
		0, 0, 243, 33, 1, 0, 0, 0, 244, 242, 1, 0, 0, 0, 245, 250, 3, 36, 18, 0,
		246, 247, 7, 6, 0, 0, 247, 249, 3, 36, 18, 0, 248, 246, 1, 0, 0, 0, 249,
		252, 1, 0, 0, 0, 250, 248, 1, 0, 0, 0, 250, 251, 1, 0, 0, 0, 251, 35, 1,
		0, 0, 0, 252, 250, 1, 0, 0, 0, 253, 258, 3, 38, 19, 0, 254, 255, 7, 7,
		0, 0, 255, 257, 3, 38, 19, 0, 256, 254, 1, 0, 0, 0, 257, 260, 1, 0, 0,
		0, 258, 256, 1, 0, 0, 0, 258, 259, 1, 0, 0, 0, 259, 37, 1, 0, 0, 0, 260,
		258, 1, 0, 0, 0, 261, 262, 7, 8, 0, 0, 262, 265, 3, 38, 19, 0, 263, 265,
		3, 46, 23, 0, 264, 261, 1, 0, 0, 0, 264, 263, 1, 0, 0, 0, 265, 39, 1, 0,
		0, 0, 266, 273, 3, 18, 9, 0, 267, 268, 5, 19, 0, 0, 268, 269, 3, 20, 10,
		0, 269, 270, 5, 20, 0, 0, 270, 274, 1, 0, 0, 0, 271, 272, 5, 27, 0, 0,
		272, 274, 5, 54, 0, 0, 273, 267, 1, 0, 0, 0, 273, 271, 1, 0, 0, 0, 274,
		41, 1, 0, 0, 0, 275, 276, 3, 40, 20, 0, 276, 277, 7, 9, 0, 0, 277, 278,
		3, 20, 10, 0, 278, 43, 1, 0, 0, 0, 279, 280, 3, 18, 9, 0, 280, 281, 5,
		22, 0, 0, 281, 286, 3, 20, 10, 0, 282, 283, 5, 18, 0, 0, 283, 285, 3, 20,
		10, 0, 284, 282, 1, 0, 0, 0, 285, 288, 1, 0, 0, 0, 286, 284, 1, 0, 0, 0,
		286, 287, 1, 0, 0, 0, 287, 290, 1, 0, 0, 0, 288, 286, 1, 0, 0, 0, 289,
		291, 5, 18, 0, 0, 290, 289, 1, 0, 0, 0, 290, 291, 1, 0, 0, 0, 291, 292,
		1, 0, 0, 0, 292, 293, 5, 23, 0, 0, 293, 45, 1, 0, 0, 0, 294, 299, 3, 40,
		20, 0, 295, 299, 3, 42, 21, 0, 296, 299, 3, 44, 22, 0, 297, 299, 3, 18,
		9, 0, 298, 294, 1, 0, 0, 0, 298, 295, 1, 0, 0, 0, 298, 296, 1, 0, 0, 0,
		298, 297, 1, 0, 0, 0, 299, 47, 1, 0, 0, 0, 300, 301, 5, 8, 0, 0, 301, 302,
		5, 54, 0, 0, 302, 49, 1, 0, 0, 0, 303, 304, 5, 54, 0, 0, 304, 305, 5, 8,
		0, 0, 305, 51, 1, 0, 0, 0, 306, 307, 5, 54, 0, 0, 307, 308, 7, 9, 0, 0,
		308, 309, 3, 20, 10, 0, 309, 53, 1, 0, 0, 0, 310, 314, 3, 52, 26, 0, 311,
		314, 3, 48, 24, 0, 312, 314, 3, 50, 25, 0, 313, 310, 1, 0, 0, 0, 313, 311,
		1, 0, 0, 0, 313, 312, 1, 0, 0, 0, 314, 315, 1, 0, 0, 0, 315, 313, 1, 0,
		0, 0, 315, 316, 1, 0, 0, 0, 316, 55, 1, 0, 0, 0, 317, 318, 5, 47, 0, 0,
		318, 323, 3, 58, 29, 0, 319, 320, 5, 18, 0, 0, 320, 322, 3, 58, 29, 0,
		321, 319, 1, 0, 0, 0, 322, 325, 1, 0, 0, 0, 323, 321, 1, 0, 0, 0, 323,
		324, 1, 0, 0, 0, 324, 57, 1, 0, 0, 0, 325, 323, 1, 0, 0, 0, 326, 329, 5,
		54, 0, 0, 327, 328, 5, 14, 0, 0, 328, 330, 3, 20, 10, 0, 329, 327, 1, 0,
		0, 0, 329, 330, 1, 0, 0, 0, 330, 59, 1, 0, 0, 0, 331, 332, 5, 38, 0, 0,
		332, 333, 3, 20, 10, 0, 333, 341, 3, 16, 8, 0, 334, 335, 5, 39, 0, 0, 335,
		336, 5, 38, 0, 0, 336, 337, 3, 20, 10, 0, 337, 338, 3, 16, 8, 0, 338, 340,
		1, 0, 0, 0, 339, 334, 1, 0, 0, 0, 340, 343, 1, 0, 0, 0, 341, 339, 1, 0,
		0, 0, 341, 342, 1, 0, 0, 0, 342, 346, 1, 0, 0, 0, 343, 341, 1, 0, 0, 0,
		344, 345, 5, 39, 0, 0, 345, 347, 3, 16, 8, 0, 346, 344, 1, 0, 0, 0, 346,
		347, 1, 0, 0, 0, 347, 61, 1, 0, 0, 0, 348, 349, 5, 13, 0, 0, 349, 350,
		3, 20, 10, 0, 350, 351, 3, 88, 44, 0, 351, 63, 1, 0, 0, 0, 352, 353, 5,
		40, 0, 0, 353, 354, 3, 20, 10, 0, 354, 358, 5, 16, 0, 0, 355, 357, 3, 66,
		33, 0, 356, 355, 1, 0, 0, 0, 357, 360, 1, 0, 0, 0, 358, 356, 1, 0, 0, 0,
		358, 359, 1, 0, 0, 0, 359, 369, 1, 0, 0, 0, 360, 358, 1, 0, 0, 0, 361,
		362, 5, 46, 0, 0, 362, 366, 5, 26, 0, 0, 363, 365, 3, 88, 44, 0, 364, 363,
		1, 0, 0, 0, 365, 368, 1, 0, 0, 0, 366, 364, 1, 0, 0, 0, 366, 367, 1, 0,
		0, 0, 367, 370, 1, 0, 0, 0, 368, 366, 1, 0, 0, 0, 369, 361, 1, 0, 0, 0,
		369, 370, 1, 0, 0, 0, 370, 371, 1, 0, 0, 0, 371, 372, 5, 17, 0, 0, 372,
		65, 1, 0, 0, 0, 373, 374, 5, 41, 0, 0, 374, 375, 3, 20, 10, 0, 375, 376,
		5, 26, 0, 0, 376, 377, 3, 88, 44, 0, 377, 67, 1, 0, 0, 0, 378, 379, 5,
		42, 0, 0, 379, 69, 1, 0, 0, 0, 380, 381, 5, 44, 0, 0, 381, 71, 1, 0, 0,
		0, 382, 383, 5, 43, 0, 0, 383, 389, 3, 82, 41, 0, 384, 386, 5, 45, 0, 0,
		385, 387, 5, 54, 0, 0, 386, 385, 1, 0, 0, 0, 386, 387, 1, 0, 0, 0, 387,
		388, 1, 0, 0, 0, 388, 390, 3, 82, 41, 0, 389, 384, 1, 0, 0, 0, 389, 390,
		1, 0, 0, 0, 390, 73, 1, 0, 0, 0, 391, 400, 3, 56, 28, 0, 392, 400, 3, 8,
		4, 0, 393, 400, 3, 60, 30, 0, 394, 400, 3, 62, 31, 0, 395, 400, 3, 64,
		32, 0, 396, 400, 3, 72, 36, 0, 397, 400, 3, 86, 43, 0, 398, 400, 3, 54,
		27, 0, 399, 391, 1, 0, 0, 0, 399, 392, 1, 0, 0, 0, 399, 393, 1, 0, 0, 0,
		399, 394, 1, 0, 0, 0, 399, 395, 1, 0, 0, 0, 399, 396, 1, 0, 0, 0, 399,
		397, 1, 0, 0, 0, 399, 398, 1, 0, 0, 0, 400, 75, 1, 0, 0, 0, 401, 415, 3,
		74, 37, 0, 402, 406, 5, 6, 0, 0, 403, 405, 3, 92, 46, 0, 404, 403, 1, 0,
		0, 0, 405, 408, 1, 0, 0, 0, 406, 404, 1, 0, 0, 0, 406, 407, 1, 0, 0, 0,
		407, 409, 1, 0, 0, 0, 408, 406, 1, 0, 0, 0, 409, 411, 5, 5, 0, 0, 410,
		402, 1, 0, 0, 0, 411, 412, 1, 0, 0, 0, 412, 410, 1, 0, 0, 0, 412, 413,
		1, 0, 0, 0, 413, 415, 1, 0, 0, 0, 414, 401, 1, 0, 0, 0, 414, 410, 1, 0,
		0, 0, 415, 77, 1, 0, 0, 0, 416, 418, 5, 16, 0, 0, 417, 419, 3, 84, 42,
		0, 418, 417, 1, 0, 0, 0, 418, 419, 1, 0, 0, 0, 419, 420, 1, 0, 0, 0, 420,
		421, 5, 17, 0, 0, 421, 79, 1, 0, 0, 0, 422, 426, 5, 16, 0, 0, 423, 425,
		3, 88, 44, 0, 424, 423, 1, 0, 0, 0, 425, 428, 1, 0, 0, 0, 426, 424, 1,
		0, 0, 0, 426, 427, 1, 0, 0, 0, 427, 429, 1, 0, 0, 0, 428, 426, 1, 0, 0,
		0, 429, 430, 5, 17, 0, 0, 430, 81, 1, 0, 0, 0, 431, 435, 3, 78, 39, 0,
		432, 435, 3, 76, 38, 0, 433, 435, 3, 20, 10, 0, 434, 431, 1, 0, 0, 0, 434,
		432, 1, 0, 0, 0, 434, 433, 1, 0, 0, 0, 435, 83, 1, 0, 0, 0, 436, 438, 3,
		82, 41, 0, 437, 436, 1, 0, 0, 0, 438, 439, 1, 0, 0, 0, 439, 437, 1, 0,
		0, 0, 439, 440, 1, 0, 0, 0, 440, 85, 1, 0, 0, 0, 441, 442, 5, 15, 0, 0,
		442, 443, 3, 20, 10, 0, 443, 87, 1, 0, 0, 0, 444, 450, 3, 80, 40, 0, 445,
		450, 3, 76, 38, 0, 446, 450, 3, 68, 34, 0, 447, 450, 3, 70, 35, 0, 448,
		450, 3, 20, 10, 0, 449, 444, 1, 0, 0, 0, 449, 445, 1, 0, 0, 0, 449, 446,
		1, 0, 0, 0, 449, 447, 1, 0, 0, 0, 449, 448, 1, 0, 0, 0, 450, 89, 1, 0,
		0, 0, 451, 453, 5, 5, 0, 0, 452, 454, 3, 84, 42, 0, 453, 452, 1, 0, 0,
		0, 453, 454, 1, 0, 0, 0, 454, 455, 1, 0, 0, 0, 455, 456, 5, 6, 0, 0, 456,
		91, 1, 0, 0, 0, 457, 458, 7, 10, 0, 0, 458, 93, 1, 0, 0, 0, 459, 462, 3,
		90, 45, 0, 460, 462, 3, 92, 46, 0, 461, 459, 1, 0, 0, 0, 461, 460, 1, 0,
		0, 0, 462, 463, 1, 0, 0, 0, 463, 461, 1, 0, 0, 0, 463, 464, 1, 0, 0, 0,
		464, 465, 1, 0, 0, 0, 465, 466, 5, 0, 0, 1, 466, 95, 1, 0, 0, 0, 467, 469,
		5, 5, 0, 0, 468, 470, 3, 84, 42, 0, 469, 468, 1, 0, 0, 0, 469, 470, 1,
		0, 0, 0, 470, 472, 1, 0, 0, 0, 471, 473, 5, 6, 0, 0, 472, 471, 1, 0, 0,
		0, 472, 473, 1, 0, 0, 0, 473, 474, 1, 0, 0, 0, 474, 475, 5, 0, 0, 1, 475,
		97, 1, 0, 0, 0, 55, 106, 109, 112, 122, 127, 135, 139, 141, 147, 151, 161,
		165, 169, 174, 177, 183, 193, 202, 210, 218, 226, 234, 242, 250, 258, 264,
		273, 286, 290, 298, 313, 315, 323, 329, 341, 346, 358, 366, 369, 386, 389,
		399, 406, 412, 414, 418, 426, 434, 439, 449, 453, 461, 463, 469, 472,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// PuppetoParserInit initializes any static state used to implement PuppetoParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewPuppetoParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func PuppetoParserInit() {
	staticData := &puppetoparserParserStaticData
	staticData.once.Do(puppetoparserParserInit)
}

// NewPuppetoParser produces a new parser instance for the optional input antlr.TokenStream.
func NewPuppetoParser(input antlr.TokenStream) *PuppetoParser {
	PuppetoParserInit()
	this := new(PuppetoParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &puppetoparserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "PuppetoParser.g4"

	return this
}

// PuppetoParser tokens.
const (
	PuppetoParserEOF                = antlr.TokenEOF
	PuppetoParserSPREAD             = 1
	PuppetoParserNAN                = 2
	PuppetoParserLT                 = 3
	PuppetoParserGT                 = 4
	PuppetoParserPUP_START_TAG      = 5
	PuppetoParserPUP_END_TAG        = 6
	PuppetoParserTWO_SIDES_OPERATOR = 7
	PuppetoParserONE_SIDE_OPERATOR  = 8
	PuppetoParserEQ                 = 9
	PuppetoParserNEQ                = 10
	PuppetoParserLTOE               = 11
	PuppetoParserGTOE               = 12
	PuppetoParserFOR                = 13
	PuppetoParserTO                 = 14
	PuppetoParserECHO               = 15
	PuppetoParserL_CURLY            = 16
	PuppetoParserR_CURLY            = 17
	PuppetoParserCOMMA              = 18
	PuppetoParserL_SQUARE           = 19
	PuppetoParserR_SQUARE           = 20
	PuppetoParserFN                 = 21
	PuppetoParserL_PARENTHES        = 22
	PuppetoParserR_PARENTHES        = 23
	PuppetoParserAND                = 24
	PuppetoParserOR                 = 25
	PuppetoParserCOLON              = 26
	PuppetoParserDOT                = 27
	PuppetoParserPLUS               = 28
	PuppetoParserMINUS              = 29
	PuppetoParserDOUBLE_STAR        = 30
	PuppetoParserPERCENT            = 31
	PuppetoParserBAND               = 32
	PuppetoParserBOR                = 33
	PuppetoParserBXOR               = 34
	PuppetoParserBLS                = 35
	PuppetoParserBRS                = 36
	PuppetoParserNOT                = 37
	PuppetoParserIF                 = 38
	PuppetoParserELSE               = 39
	PuppetoParserSWITCH             = 40
	PuppetoParserCASE               = 41
	PuppetoParserBREAK              = 42
	PuppetoParserTRY                = 43
	PuppetoParserCONTINUE           = 44
	PuppetoParserCATCH              = 45
	PuppetoParserDEFAULT            = 46
	PuppetoParserLET                = 47
	PuppetoParserSTAR               = 48
	PuppetoParserSLASH              = 49
	PuppetoParserFLOAT              = 50
	PuppetoParserINTEGER            = 51
	PuppetoParserNIL                = 52
	PuppetoParserBOOLEAN            = 53
	PuppetoParserIDENTIFIER         = 54
	PuppetoParserSTRING             = 55
	PuppetoParserWS                 = 56
	PuppetoParserSEMICOLON          = 57
	PuppetoParserHTML               = 58
)

// PuppetoParser rules.
const (
	PuppetoParserRULE_literalValue                 = 0
	PuppetoParserRULE_object                       = 1
	PuppetoParserRULE_pair                         = 2
	PuppetoParserRULE_array                        = 3
	PuppetoParserRULE_functionDefinition           = 4
	PuppetoParserRULE_argumentList                 = 5
	PuppetoParserRULE_argument                     = 6
	PuppetoParserRULE_none                         = 7
	PuppetoParserRULE_scopeBody                    = 8
	PuppetoParserRULE_value                        = 9
	PuppetoParserRULE_expression                   = 10
	PuppetoParserRULE_logicalExpression            = 11
	PuppetoParserRULE_equalityExpression           = 12
	PuppetoParserRULE_comparisonExpression         = 13
	PuppetoParserRULE_additiveExpression           = 14
	PuppetoParserRULE_multiplicativeExpression     = 15
	PuppetoParserRULE_powerExpression              = 16
	PuppetoParserRULE_bitwiseExpression            = 17
	PuppetoParserRULE_shiftExpression              = 18
	PuppetoParserRULE_unaryExpression              = 19
	PuppetoParserRULE_getProperty                  = 20
	PuppetoParserRULE_setProperty                  = 21
	PuppetoParserRULE_callExpression               = 22
	PuppetoParserRULE_chainExpression              = 23
	PuppetoParserRULE_lhsVariableAssignation       = 24
	PuppetoParserRULE_rhsVariableAssignation       = 25
	PuppetoParserRULE_midVariableAssignation       = 26
	PuppetoParserRULE_variableAssignation          = 27
	PuppetoParserRULE_variableDefinition           = 28
	PuppetoParserRULE_variableDefinitionBody       = 29
	PuppetoParserRULE_ifStatement                  = 30
	PuppetoParserRULE_loopStatement                = 31
	PuppetoParserRULE_switchStatement              = 32
	PuppetoParserRULE_switchCase                   = 33
	PuppetoParserRULE_breakStatement               = 34
	PuppetoParserRULE_continueStatement            = 35
	PuppetoParserRULE_tryStatement                 = 36
	PuppetoParserRULE_statementList                = 37
	PuppetoParserRULE_statement                    = 38
	PuppetoParserRULE_scope                        = 39
	PuppetoParserRULE_scopeWithBreakContinue       = 40
	PuppetoParserRULE_programBody                  = 41
	PuppetoParserRULE_programBodyList              = 42
	PuppetoParserRULE_echoStatement                = 43
	PuppetoParserRULE_programBodyWithBreakContinue = 44
	PuppetoParserRULE_pupCode                      = 45
	PuppetoParserRULE_htmlList                     = 46
	PuppetoParserRULE_html                         = 47
	PuppetoParserRULE_program                      = 48
)

// ILiteralValueContext is an interface to support dynamic dispatch.
type ILiteralValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FLOAT() antlr.TerminalNode
	INTEGER() antlr.TerminalNode
	STRING() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	NIL() antlr.TerminalNode
	BOOLEAN() antlr.TerminalNode
	NAN() antlr.TerminalNode

	// IsLiteralValueContext differentiates from other interfaces.
	IsLiteralValueContext()
}

type LiteralValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralValueContext() *LiteralValueContext {
	var p = new(LiteralValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PuppetoParserRULE_literalValue
	return p
}

func (*LiteralValueContext) IsLiteralValueContext() {}

func NewLiteralValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralValueContext {
	var p = new(LiteralValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PuppetoParserRULE_literalValue

	return p
}

func (s *LiteralValueContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralValueContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(PuppetoParserFLOAT, 0)
}

func (s *LiteralValueContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(PuppetoParserINTEGER, 0)
}

func (s *LiteralValueContext) STRING() antlr.TerminalNode {
	return s.GetToken(PuppetoParserSTRING, 0)
}

func (s *LiteralValueContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(PuppetoParserIDENTIFIER, 0)
}

func (s *LiteralValueContext) NIL() antlr.TerminalNode {
	return s.GetToken(PuppetoParserNIL, 0)
}

func (s *LiteralValueContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(PuppetoParserBOOLEAN, 0)
}

func (s *LiteralValueContext) NAN() antlr.TerminalNode {
	return s.GetToken(PuppetoParserNAN, 0)
}

func (s *LiteralValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PuppetoParserVisitor:
		return t.VisitLiteralValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PuppetoParser) LiteralValue() (localctx ILiteralValueContext) {
	this := p
	_ = this

	localctx = NewLiteralValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, PuppetoParserRULE_literalValue)
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
		p.SetState(98)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&70931694131085316) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IObjectContext is an interface to support dynamic dispatch.
type IObjectContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	L_CURLY() antlr.TerminalNode
	R_CURLY() antlr.TerminalNode
	AllPair() []IPairContext
	Pair(i int) IPairContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsObjectContext differentiates from other interfaces.
	IsObjectContext()
}

type ObjectContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjectContext() *ObjectContext {
	var p = new(ObjectContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PuppetoParserRULE_object
	return p
}

func (*ObjectContext) IsObjectContext() {}

func NewObjectContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjectContext {
	var p = new(ObjectContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PuppetoParserRULE_object

	return p
}

func (s *ObjectContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjectContext) L_CURLY() antlr.TerminalNode {
	return s.GetToken(PuppetoParserL_CURLY, 0)
}

func (s *ObjectContext) R_CURLY() antlr.TerminalNode {
	return s.GetToken(PuppetoParserR_CURLY, 0)
}

func (s *ObjectContext) AllPair() []IPairContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPairContext); ok {
			len++
		}
	}

	tst := make([]IPairContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPairContext); ok {
			tst[i] = t.(IPairContext)
			i++
		}
	}

	return tst
}

func (s *ObjectContext) Pair(i int) IPairContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPairContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPairContext)
}

func (s *ObjectContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(PuppetoParserCOMMA)
}

func (s *ObjectContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(PuppetoParserCOMMA, i)
}

func (s *ObjectContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjectContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PuppetoParserVisitor:
		return t.VisitObject(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PuppetoParser) Object() (localctx IObjectContext) {
	this := p
	_ = this

	localctx = NewObjectContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, PuppetoParserRULE_object)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(100)
		p.Match(PuppetoParserL_CURLY)
	}
	p.SetState(109)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&54043195528970240) != 0 {
		{
			p.SetState(101)
			p.Pair()
		}
		p.SetState(106)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(102)
					p.Match(PuppetoParserCOMMA)
				}
				{
					p.SetState(103)
					p.Pair()
				}

			}
			p.SetState(108)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())
		}

	}
	p.SetState(112)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PuppetoParserCOMMA {
		{
			p.SetState(111)
			p.Match(PuppetoParserCOMMA)
		}

	}
	{
		p.SetState(114)
		p.Match(PuppetoParserR_CURLY)
	}

	return localctx
}

// IPairContext is an interface to support dynamic dispatch.
type IPairContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	COLON() antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	L_SQUARE() antlr.TerminalNode
	R_SQUARE() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	STRING() antlr.TerminalNode

	// IsPairContext differentiates from other interfaces.
	IsPairContext()
}

type PairContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPairContext() *PairContext {
	var p = new(PairContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PuppetoParserRULE_pair
	return p
}

func (*PairContext) IsPairContext() {}

func NewPairContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PairContext {
	var p = new(PairContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PuppetoParserRULE_pair

	return p
}

func (s *PairContext) GetParser() antlr.Parser { return s.parser }

func (s *PairContext) COLON() antlr.TerminalNode {
	return s.GetToken(PuppetoParserCOLON, 0)
}

func (s *PairContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *PairContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PairContext) L_SQUARE() antlr.TerminalNode {
	return s.GetToken(PuppetoParserL_SQUARE, 0)
}

func (s *PairContext) R_SQUARE() antlr.TerminalNode {
	return s.GetToken(PuppetoParserR_SQUARE, 0)
}

func (s *PairContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(PuppetoParserIDENTIFIER, 0)
}

func (s *PairContext) STRING() antlr.TerminalNode {
	return s.GetToken(PuppetoParserSTRING, 0)
}

func (s *PairContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PairContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PairContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PuppetoParserVisitor:
		return t.VisitPair(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PuppetoParser) Pair() (localctx IPairContext) {
	this := p
	_ = this

	localctx = NewPairContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, PuppetoParserRULE_pair)

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

	p.SetState(127)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(122)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case PuppetoParserL_SQUARE:
			{
				p.SetState(116)
				p.Match(PuppetoParserL_SQUARE)
			}
			{
				p.SetState(117)
				p.Expression()
			}
			{
				p.SetState(118)
				p.Match(PuppetoParserR_SQUARE)
			}

		case PuppetoParserIDENTIFIER:
			{
				p.SetState(120)
				p.Match(PuppetoParserIDENTIFIER)
			}

		case PuppetoParserSTRING:
			{
				p.SetState(121)
				p.Match(PuppetoParserSTRING)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}
		{
			p.SetState(124)
			p.Match(PuppetoParserCOLON)
		}
		{
			p.SetState(125)
			p.Expression()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(126)
			p.Match(PuppetoParserIDENTIFIER)
		}

	}

	return localctx
}

// IArrayContext is an interface to support dynamic dispatch.
type IArrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	L_SQUARE() antlr.TerminalNode
	R_SQUARE() antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsArrayContext differentiates from other interfaces.
	IsArrayContext()
}

type ArrayContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayContext() *ArrayContext {
	var p = new(ArrayContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PuppetoParserRULE_array
	return p
}

func (*ArrayContext) IsArrayContext() {}

func NewArrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayContext {
	var p = new(ArrayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PuppetoParserRULE_array

	return p
}

func (s *ArrayContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayContext) L_SQUARE() antlr.TerminalNode {
	return s.GetToken(PuppetoParserL_SQUARE, 0)
}

func (s *ArrayContext) R_SQUARE() antlr.TerminalNode {
	return s.GetToken(PuppetoParserR_SQUARE, 0)
}

func (s *ArrayContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ArrayContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ArrayContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(PuppetoParserCOMMA)
}

func (s *ArrayContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(PuppetoParserCOMMA, i)
}

func (s *ArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PuppetoParserVisitor:
		return t.VisitArray(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PuppetoParser) Array() (localctx IArrayContext) {
	this := p
	_ = this

	localctx = NewArrayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, PuppetoParserRULE_array)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(129)
		p.Match(PuppetoParserL_SQUARE)
	}
	p.SetState(141)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&70931832113790980) != 0 {
		{
			p.SetState(130)
			p.Expression()
		}
		p.SetState(135)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(131)
					p.Match(PuppetoParserCOMMA)
				}
				{
					p.SetState(132)
					p.Expression()
				}

			}
			p.SetState(137)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())
		}
		p.SetState(139)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == PuppetoParserCOMMA {
			{
				p.SetState(138)
				p.Match(PuppetoParserCOMMA)
			}

		}

	}
	{
		p.SetState(143)
		p.Match(PuppetoParserR_SQUARE)
	}

	return localctx
}

// IFunctionDefinitionContext is an interface to support dynamic dispatch.
type IFunctionDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// Getter signatures
	FN() antlr.TerminalNode
	L_PARENTHES() antlr.TerminalNode
	R_PARENTHES() antlr.TerminalNode
	ScopeBody() IScopeBodyContext
	ArgumentList() IArgumentListContext
	IDENTIFIER() antlr.TerminalNode

	// IsFunctionDefinitionContext differentiates from other interfaces.
	IsFunctionDefinitionContext()
}

type FunctionDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
}

func NewEmptyFunctionDefinitionContext() *FunctionDefinitionContext {
	var p = new(FunctionDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PuppetoParserRULE_functionDefinition
	return p
}

func (*FunctionDefinitionContext) IsFunctionDefinitionContext() {}

func NewFunctionDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionDefinitionContext {
	var p = new(FunctionDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PuppetoParserRULE_functionDefinition

	return p
}

func (s *FunctionDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionDefinitionContext) GetName() antlr.Token { return s.name }

func (s *FunctionDefinitionContext) SetName(v antlr.Token) { s.name = v }

func (s *FunctionDefinitionContext) FN() antlr.TerminalNode {
	return s.GetToken(PuppetoParserFN, 0)
}

func (s *FunctionDefinitionContext) L_PARENTHES() antlr.TerminalNode {
	return s.GetToken(PuppetoParserL_PARENTHES, 0)
}

func (s *FunctionDefinitionContext) R_PARENTHES() antlr.TerminalNode {
	return s.GetToken(PuppetoParserR_PARENTHES, 0)
}

func (s *FunctionDefinitionContext) ScopeBody() IScopeBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IScopeBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IScopeBodyContext)
}

func (s *FunctionDefinitionContext) ArgumentList() IArgumentListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentListContext)
}

func (s *FunctionDefinitionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(PuppetoParserIDENTIFIER, 0)
}

func (s *FunctionDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PuppetoParserVisitor:
		return t.VisitFunctionDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PuppetoParser) FunctionDefinition() (localctx IFunctionDefinitionContext) {
	this := p
	_ = this

	localctx = NewFunctionDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, PuppetoParserRULE_functionDefinition)
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
		p.SetState(145)
		p.Match(PuppetoParserFN)
	}
	p.SetState(147)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PuppetoParserIDENTIFIER {
		{
			p.SetState(146)

			var _m = p.Match(PuppetoParserIDENTIFIER)

			localctx.(*FunctionDefinitionContext).name = _m
		}

	}
	{
		p.SetState(149)
		p.Match(PuppetoParserL_PARENTHES)
	}
	p.SetState(151)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PuppetoParserIDENTIFIER {
		{
			p.SetState(150)
			p.ArgumentList()
		}

	}
	{
		p.SetState(153)
		p.Match(PuppetoParserR_PARENTHES)
	}
	{
		p.SetState(154)
		p.ScopeBody()
	}

	return localctx
}

// IArgumentListContext is an interface to support dynamic dispatch.
type IArgumentListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetRest returns the rest token.
	GetRest() antlr.Token

	// SetRest sets the rest token.
	SetRest(antlr.Token)

	// Getter signatures
	AllArgument() []IArgumentContext
	Argument(i int) IArgumentContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	SPREAD() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode

	// IsArgumentListContext differentiates from other interfaces.
	IsArgumentListContext()
}

type ArgumentListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	rest   antlr.Token
}

func NewEmptyArgumentListContext() *ArgumentListContext {
	var p = new(ArgumentListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PuppetoParserRULE_argumentList
	return p
}

func (*ArgumentListContext) IsArgumentListContext() {}

func NewArgumentListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentListContext {
	var p = new(ArgumentListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PuppetoParserRULE_argumentList

	return p
}

func (s *ArgumentListContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentListContext) GetRest() antlr.Token { return s.rest }

func (s *ArgumentListContext) SetRest(v antlr.Token) { s.rest = v }

func (s *ArgumentListContext) AllArgument() []IArgumentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IArgumentContext); ok {
			len++
		}
	}

	tst := make([]IArgumentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IArgumentContext); ok {
			tst[i] = t.(IArgumentContext)
			i++
		}
	}

	return tst
}

func (s *ArgumentListContext) Argument(i int) IArgumentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentContext)
}

func (s *ArgumentListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(PuppetoParserCOMMA)
}

func (s *ArgumentListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(PuppetoParserCOMMA, i)
}

func (s *ArgumentListContext) SPREAD() antlr.TerminalNode {
	return s.GetToken(PuppetoParserSPREAD, 0)
}

func (s *ArgumentListContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(PuppetoParserIDENTIFIER, 0)
}

func (s *ArgumentListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PuppetoParserVisitor:
		return t.VisitArgumentList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PuppetoParser) ArgumentList() (localctx IArgumentListContext) {
	this := p
	_ = this

	localctx = NewArgumentListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, PuppetoParserRULE_argumentList)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(156)
		p.Argument()
	}
	p.SetState(161)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(157)
				p.Match(PuppetoParserCOMMA)
			}
			{
				p.SetState(158)
				p.Argument()
			}

		}
		p.SetState(163)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())
	}
	p.SetState(165)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PuppetoParserCOMMA {
		{
			p.SetState(164)
			p.Match(PuppetoParserCOMMA)
		}

	}
	p.SetState(169)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PuppetoParserSPREAD {
		{
			p.SetState(167)
			p.Match(PuppetoParserSPREAD)
		}
		{
			p.SetState(168)

			var _m = p.Match(PuppetoParserIDENTIFIER)

			localctx.(*ArgumentListContext).rest = _m
		}

	}

	return localctx
}

// IArgumentContext is an interface to support dynamic dispatch.
type IArgumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	None() INoneContext
	TO() antlr.TerminalNode
	Expression() IExpressionContext

	// IsArgumentContext differentiates from other interfaces.
	IsArgumentContext()
}

type ArgumentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentContext() *ArgumentContext {
	var p = new(ArgumentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PuppetoParserRULE_argument
	return p
}

func (*ArgumentContext) IsArgumentContext() {}

func NewArgumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentContext {
	var p = new(ArgumentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PuppetoParserRULE_argument

	return p
}

func (s *ArgumentContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(PuppetoParserIDENTIFIER, 0)
}

func (s *ArgumentContext) None() INoneContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INoneContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INoneContext)
}

func (s *ArgumentContext) TO() antlr.TerminalNode {
	return s.GetToken(PuppetoParserTO, 0)
}

func (s *ArgumentContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ArgumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PuppetoParserVisitor:
		return t.VisitArgument(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PuppetoParser) Argument() (localctx IArgumentContext) {
	this := p
	_ = this

	localctx = NewArgumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, PuppetoParserRULE_argument)
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
		p.SetState(171)
		p.Match(PuppetoParserIDENTIFIER)
	}
	p.SetState(177)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		p.SetState(174)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == PuppetoParserTO {
			{
				p.SetState(172)
				p.Match(PuppetoParserTO)
			}
			{
				p.SetState(173)
				p.Expression()
			}

		}

	case 2:
		{
			p.SetState(176)
			p.None()
		}

	}

	return localctx
}

// INoneContext is an interface to support dynamic dispatch.
type INoneContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsNoneContext differentiates from other interfaces.
	IsNoneContext()
}

type NoneContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNoneContext() *NoneContext {
	var p = new(NoneContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PuppetoParserRULE_none
	return p
}

func (*NoneContext) IsNoneContext() {}

func NewNoneContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NoneContext {
	var p = new(NoneContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PuppetoParserRULE_none

	return p
}

func (s *NoneContext) GetParser() antlr.Parser { return s.parser }
func (s *NoneContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NoneContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NoneContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PuppetoParserVisitor:
		return t.VisitNone(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PuppetoParser) None() (localctx INoneContext) {
	this := p
	_ = this

	localctx = NewNoneContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, PuppetoParserRULE_none)

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

	return localctx
}

// IScopeBodyContext is an interface to support dynamic dispatch.
type IScopeBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext
	Scope() IScopeContext

	// IsScopeBodyContext differentiates from other interfaces.
	IsScopeBodyContext()
}

type ScopeBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyScopeBodyContext() *ScopeBodyContext {
	var p = new(ScopeBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PuppetoParserRULE_scopeBody
	return p
}

func (*ScopeBodyContext) IsScopeBodyContext() {}

func NewScopeBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ScopeBodyContext {
	var p = new(ScopeBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PuppetoParserRULE_scopeBody

	return p
}

func (s *ScopeBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *ScopeBodyContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ScopeBodyContext) Scope() IScopeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IScopeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IScopeContext)
}

func (s *ScopeBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScopeBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ScopeBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PuppetoParserVisitor:
		return t.VisitScopeBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PuppetoParser) ScopeBody() (localctx IScopeBodyContext) {
	this := p
	_ = this

	localctx = NewScopeBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, PuppetoParserRULE_scopeBody)

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

	p.SetState(183)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(181)
			p.Expression()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(182)
			p.Scope()
		}

	}

	return localctx
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LiteralValue() ILiteralValueContext
	Object() IObjectContext
	Array() IArrayContext
	FunctionDefinition() IFunctionDefinitionContext
	L_PARENTHES() antlr.TerminalNode
	Expression() IExpressionContext
	R_PARENTHES() antlr.TerminalNode

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PuppetoParserRULE_value
	return p
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PuppetoParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) LiteralValue() ILiteralValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralValueContext)
}

func (s *ValueContext) Object() IObjectContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IObjectContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IObjectContext)
}

func (s *ValueContext) Array() IArrayContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayContext)
}

func (s *ValueContext) FunctionDefinition() IFunctionDefinitionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionDefinitionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionDefinitionContext)
}

func (s *ValueContext) L_PARENTHES() antlr.TerminalNode {
	return s.GetToken(PuppetoParserL_PARENTHES, 0)
}

func (s *ValueContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ValueContext) R_PARENTHES() antlr.TerminalNode {
	return s.GetToken(PuppetoParserR_PARENTHES, 0)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PuppetoParserVisitor:
		return t.VisitValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PuppetoParser) Value() (localctx IValueContext) {
	this := p
	_ = this

	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, PuppetoParserRULE_value)

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

	p.SetState(193)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PuppetoParserNAN, PuppetoParserFLOAT, PuppetoParserINTEGER, PuppetoParserNIL, PuppetoParserBOOLEAN, PuppetoParserIDENTIFIER, PuppetoParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(185)
			p.LiteralValue()
		}

	case PuppetoParserL_CURLY:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(186)
			p.Object()
		}

	case PuppetoParserL_SQUARE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(187)
			p.Array()
		}

	case PuppetoParserFN:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(188)
			p.FunctionDefinition()
		}

	case PuppetoParserL_PARENTHES:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(189)
			p.Match(PuppetoParserL_PARENTHES)
		}
		{
			p.SetState(190)
			p.Expression()
		}
		{
			p.SetState(191)
			p.Match(PuppetoParserR_PARENTHES)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LogicalExpression() ILogicalExpressionContext

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PuppetoParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PuppetoParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) LogicalExpression() ILogicalExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILogicalExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILogicalExpressionContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PuppetoParserVisitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PuppetoParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, PuppetoParserRULE_expression)

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
		p.SetState(195)
		p.LogicalExpression()
	}

	return localctx
}

// ILogicalExpressionContext is an interface to support dynamic dispatch.
type ILogicalExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllEqualityExpression() []IEqualityExpressionContext
	EqualityExpression(i int) IEqualityExpressionContext
	AllAND() []antlr.TerminalNode
	AND(i int) antlr.TerminalNode
	AllOR() []antlr.TerminalNode
	OR(i int) antlr.TerminalNode

	// IsLogicalExpressionContext differentiates from other interfaces.
	IsLogicalExpressionContext()
}

type LogicalExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLogicalExpressionContext() *LogicalExpressionContext {
	var p = new(LogicalExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PuppetoParserRULE_logicalExpression
	return p
}

func (*LogicalExpressionContext) IsLogicalExpressionContext() {}

func NewLogicalExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LogicalExpressionContext {
	var p = new(LogicalExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PuppetoParserRULE_logicalExpression

	return p
}

func (s *LogicalExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *LogicalExpressionContext) AllEqualityExpression() []IEqualityExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEqualityExpressionContext); ok {
			len++
		}
	}

	tst := make([]IEqualityExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEqualityExpressionContext); ok {
			tst[i] = t.(IEqualityExpressionContext)
			i++
		}
	}

	return tst
}

func (s *LogicalExpressionContext) EqualityExpression(i int) IEqualityExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEqualityExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEqualityExpressionContext)
}

func (s *LogicalExpressionContext) AllAND() []antlr.TerminalNode {
	return s.GetTokens(PuppetoParserAND)
}

func (s *LogicalExpressionContext) AND(i int) antlr.TerminalNode {
	return s.GetToken(PuppetoParserAND, i)
}

func (s *LogicalExpressionContext) AllOR() []antlr.TerminalNode {
	return s.GetTokens(PuppetoParserOR)
}

func (s *LogicalExpressionContext) OR(i int) antlr.TerminalNode {
	return s.GetToken(PuppetoParserOR, i)
}

func (s *LogicalExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LogicalExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PuppetoParserVisitor:
		return t.VisitLogicalExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PuppetoParser) LogicalExpression() (localctx ILogicalExpressionContext) {
	this := p
	_ = this

	localctx = NewLogicalExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, PuppetoParserRULE_logicalExpression)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(197)
		p.EqualityExpression()
	}
	p.SetState(202)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(198)
				_la = p.GetTokenStream().LA(1)

				if !(_la == PuppetoParserAND || _la == PuppetoParserOR) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(199)
				p.EqualityExpression()
			}

		}
		p.SetState(204)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext())
	}

	return localctx
}

// IEqualityExpressionContext is an interface to support dynamic dispatch.
type IEqualityExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllComparisonExpression() []IComparisonExpressionContext
	ComparisonExpression(i int) IComparisonExpressionContext
	AllEQ() []antlr.TerminalNode
	EQ(i int) antlr.TerminalNode
	AllNEQ() []antlr.TerminalNode
	NEQ(i int) antlr.TerminalNode

	// IsEqualityExpressionContext differentiates from other interfaces.
	IsEqualityExpressionContext()
}

type EqualityExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEqualityExpressionContext() *EqualityExpressionContext {
	var p = new(EqualityExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PuppetoParserRULE_equalityExpression
	return p
}

func (*EqualityExpressionContext) IsEqualityExpressionContext() {}

func NewEqualityExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EqualityExpressionContext {
	var p = new(EqualityExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PuppetoParserRULE_equalityExpression

	return p
}

func (s *EqualityExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *EqualityExpressionContext) AllComparisonExpression() []IComparisonExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IComparisonExpressionContext); ok {
			len++
		}
	}

	tst := make([]IComparisonExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IComparisonExpressionContext); ok {
			tst[i] = t.(IComparisonExpressionContext)
			i++
		}
	}

	return tst
}

func (s *EqualityExpressionContext) ComparisonExpression(i int) IComparisonExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComparisonExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComparisonExpressionContext)
}

func (s *EqualityExpressionContext) AllEQ() []antlr.TerminalNode {
	return s.GetTokens(PuppetoParserEQ)
}

func (s *EqualityExpressionContext) EQ(i int) antlr.TerminalNode {
	return s.GetToken(PuppetoParserEQ, i)
}

func (s *EqualityExpressionContext) AllNEQ() []antlr.TerminalNode {
	return s.GetTokens(PuppetoParserNEQ)
}

func (s *EqualityExpressionContext) NEQ(i int) antlr.TerminalNode {
	return s.GetToken(PuppetoParserNEQ, i)
}

func (s *EqualityExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualityExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EqualityExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PuppetoParserVisitor:
		return t.VisitEqualityExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PuppetoParser) EqualityExpression() (localctx IEqualityExpressionContext) {
	this := p
	_ = this

	localctx = NewEqualityExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, PuppetoParserRULE_equalityExpression)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(205)
		p.ComparisonExpression()
	}
	p.SetState(210)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(206)
				_la = p.GetTokenStream().LA(1)

				if !(_la == PuppetoParserEQ || _la == PuppetoParserNEQ) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(207)
				p.ComparisonExpression()
			}

		}
		p.SetState(212)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext())
	}

	return localctx
}

// IComparisonExpressionContext is an interface to support dynamic dispatch.
type IComparisonExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllAdditiveExpression() []IAdditiveExpressionContext
	AdditiveExpression(i int) IAdditiveExpressionContext
	AllLT() []antlr.TerminalNode
	LT(i int) antlr.TerminalNode
	AllGT() []antlr.TerminalNode
	GT(i int) antlr.TerminalNode
	AllLTOE() []antlr.TerminalNode
	LTOE(i int) antlr.TerminalNode
	AllGTOE() []antlr.TerminalNode
	GTOE(i int) antlr.TerminalNode

	// IsComparisonExpressionContext differentiates from other interfaces.
	IsComparisonExpressionContext()
}

type ComparisonExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparisonExpressionContext() *ComparisonExpressionContext {
	var p = new(ComparisonExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PuppetoParserRULE_comparisonExpression
	return p
}

func (*ComparisonExpressionContext) IsComparisonExpressionContext() {}

func NewComparisonExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComparisonExpressionContext {
	var p = new(ComparisonExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PuppetoParserRULE_comparisonExpression

	return p
}

func (s *ComparisonExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ComparisonExpressionContext) AllAdditiveExpression() []IAdditiveExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAdditiveExpressionContext); ok {
			len++
		}
	}

	tst := make([]IAdditiveExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAdditiveExpressionContext); ok {
			tst[i] = t.(IAdditiveExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ComparisonExpressionContext) AdditiveExpression(i int) IAdditiveExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAdditiveExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAdditiveExpressionContext)
}

func (s *ComparisonExpressionContext) AllLT() []antlr.TerminalNode {
	return s.GetTokens(PuppetoParserLT)
}

func (s *ComparisonExpressionContext) LT(i int) antlr.TerminalNode {
	return s.GetToken(PuppetoParserLT, i)
}

func (s *ComparisonExpressionContext) AllGT() []antlr.TerminalNode {
	return s.GetTokens(PuppetoParserGT)
}

func (s *ComparisonExpressionContext) GT(i int) antlr.TerminalNode {
	return s.GetToken(PuppetoParserGT, i)
}

func (s *ComparisonExpressionContext) AllLTOE() []antlr.TerminalNode {
	return s.GetTokens(PuppetoParserLTOE)
}

func (s *ComparisonExpressionContext) LTOE(i int) antlr.TerminalNode {
	return s.GetToken(PuppetoParserLTOE, i)
}

func (s *ComparisonExpressionContext) AllGTOE() []antlr.TerminalNode {
	return s.GetTokens(PuppetoParserGTOE)
}

func (s *ComparisonExpressionContext) GTOE(i int) antlr.TerminalNode {
	return s.GetToken(PuppetoParserGTOE, i)
}

func (s *ComparisonExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparisonExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComparisonExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PuppetoParserVisitor:
		return t.VisitComparisonExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PuppetoParser) ComparisonExpression() (localctx IComparisonExpressionContext) {
	this := p
	_ = this

	localctx = NewComparisonExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, PuppetoParserRULE_comparisonExpression)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(213)
		p.AdditiveExpression()
	}
	p.SetState(218)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(214)
				_la = p.GetTokenStream().LA(1)

				if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&6168) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(215)
				p.AdditiveExpression()
			}

		}
		p.SetState(220)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext())
	}

	return localctx
}

// IAdditiveExpressionContext is an interface to support dynamic dispatch.
type IAdditiveExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllMultiplicativeExpression() []IMultiplicativeExpressionContext
	MultiplicativeExpression(i int) IMultiplicativeExpressionContext
	AllPLUS() []antlr.TerminalNode
	PLUS(i int) antlr.TerminalNode
	AllMINUS() []antlr.TerminalNode
	MINUS(i int) antlr.TerminalNode

	// IsAdditiveExpressionContext differentiates from other interfaces.
	IsAdditiveExpressionContext()
}

type AdditiveExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAdditiveExpressionContext() *AdditiveExpressionContext {
	var p = new(AdditiveExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PuppetoParserRULE_additiveExpression
	return p
}

func (*AdditiveExpressionContext) IsAdditiveExpressionContext() {}

func NewAdditiveExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AdditiveExpressionContext {
	var p = new(AdditiveExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PuppetoParserRULE_additiveExpression

	return p
}

func (s *AdditiveExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *AdditiveExpressionContext) AllMultiplicativeExpression() []IMultiplicativeExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMultiplicativeExpressionContext); ok {
			len++
		}
	}

	tst := make([]IMultiplicativeExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMultiplicativeExpressionContext); ok {
			tst[i] = t.(IMultiplicativeExpressionContext)
			i++
		}
	}

	return tst
}

func (s *AdditiveExpressionContext) MultiplicativeExpression(i int) IMultiplicativeExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMultiplicativeExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMultiplicativeExpressionContext)
}

func (s *AdditiveExpressionContext) AllPLUS() []antlr.TerminalNode {
	return s.GetTokens(PuppetoParserPLUS)
}

func (s *AdditiveExpressionContext) PLUS(i int) antlr.TerminalNode {
	return s.GetToken(PuppetoParserPLUS, i)
}

func (s *AdditiveExpressionContext) AllMINUS() []antlr.TerminalNode {
	return s.GetTokens(PuppetoParserMINUS)
}

func (s *AdditiveExpressionContext) MINUS(i int) antlr.TerminalNode {
	return s.GetToken(PuppetoParserMINUS, i)
}

func (s *AdditiveExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdditiveExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AdditiveExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PuppetoParserVisitor:
		return t.VisitAdditiveExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PuppetoParser) AdditiveExpression() (localctx IAdditiveExpressionContext) {
	this := p
	_ = this

	localctx = NewAdditiveExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, PuppetoParserRULE_additiveExpression)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(221)
		p.MultiplicativeExpression()
	}
	p.SetState(226)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(222)
				_la = p.GetTokenStream().LA(1)

				if !(_la == PuppetoParserPLUS || _la == PuppetoParserMINUS) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(223)
				p.MultiplicativeExpression()
			}

		}
		p.SetState(228)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())
	}

	return localctx
}

// IMultiplicativeExpressionContext is an interface to support dynamic dispatch.
type IMultiplicativeExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllPowerExpression() []IPowerExpressionContext
	PowerExpression(i int) IPowerExpressionContext
	AllSTAR() []antlr.TerminalNode
	STAR(i int) antlr.TerminalNode
	AllSLASH() []antlr.TerminalNode
	SLASH(i int) antlr.TerminalNode
	AllPERCENT() []antlr.TerminalNode
	PERCENT(i int) antlr.TerminalNode

	// IsMultiplicativeExpressionContext differentiates from other interfaces.
	IsMultiplicativeExpressionContext()
}

type MultiplicativeExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultiplicativeExpressionContext() *MultiplicativeExpressionContext {
	var p = new(MultiplicativeExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PuppetoParserRULE_multiplicativeExpression
	return p
}

func (*MultiplicativeExpressionContext) IsMultiplicativeExpressionContext() {}

func NewMultiplicativeExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiplicativeExpressionContext {
	var p = new(MultiplicativeExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PuppetoParserRULE_multiplicativeExpression

	return p
}

func (s *MultiplicativeExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *MultiplicativeExpressionContext) AllPowerExpression() []IPowerExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPowerExpressionContext); ok {
			len++
		}
	}

	tst := make([]IPowerExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPowerExpressionContext); ok {
			tst[i] = t.(IPowerExpressionContext)
			i++
		}
	}

	return tst
}

func (s *MultiplicativeExpressionContext) PowerExpression(i int) IPowerExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPowerExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPowerExpressionContext)
}

func (s *MultiplicativeExpressionContext) AllSTAR() []antlr.TerminalNode {
	return s.GetTokens(PuppetoParserSTAR)
}

func (s *MultiplicativeExpressionContext) STAR(i int) antlr.TerminalNode {
	return s.GetToken(PuppetoParserSTAR, i)
}

func (s *MultiplicativeExpressionContext) AllSLASH() []antlr.TerminalNode {
	return s.GetTokens(PuppetoParserSLASH)
}

func (s *MultiplicativeExpressionContext) SLASH(i int) antlr.TerminalNode {
	return s.GetToken(PuppetoParserSLASH, i)
}

func (s *MultiplicativeExpressionContext) AllPERCENT() []antlr.TerminalNode {
	return s.GetTokens(PuppetoParserPERCENT)
}

func (s *MultiplicativeExpressionContext) PERCENT(i int) antlr.TerminalNode {
	return s.GetToken(PuppetoParserPERCENT, i)
}

func (s *MultiplicativeExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplicativeExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultiplicativeExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PuppetoParserVisitor:
		return t.VisitMultiplicativeExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PuppetoParser) MultiplicativeExpression() (localctx IMultiplicativeExpressionContext) {
	this := p
	_ = this

	localctx = NewMultiplicativeExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, PuppetoParserRULE_multiplicativeExpression)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(229)
		p.PowerExpression()
	}
	p.SetState(234)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(230)
				_la = p.GetTokenStream().LA(1)

				if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&844427077615616) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(231)
				p.PowerExpression()
			}

		}
		p.SetState(236)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext())
	}

	return localctx
}

// IPowerExpressionContext is an interface to support dynamic dispatch.
type IPowerExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllBitwiseExpression() []IBitwiseExpressionContext
	BitwiseExpression(i int) IBitwiseExpressionContext
	AllDOUBLE_STAR() []antlr.TerminalNode
	DOUBLE_STAR(i int) antlr.TerminalNode

	// IsPowerExpressionContext differentiates from other interfaces.
	IsPowerExpressionContext()
}

type PowerExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPowerExpressionContext() *PowerExpressionContext {
	var p = new(PowerExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PuppetoParserRULE_powerExpression
	return p
}

func (*PowerExpressionContext) IsPowerExpressionContext() {}

func NewPowerExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PowerExpressionContext {
	var p = new(PowerExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PuppetoParserRULE_powerExpression

	return p
}

func (s *PowerExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *PowerExpressionContext) AllBitwiseExpression() []IBitwiseExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBitwiseExpressionContext); ok {
			len++
		}
	}

	tst := make([]IBitwiseExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBitwiseExpressionContext); ok {
			tst[i] = t.(IBitwiseExpressionContext)
			i++
		}
	}

	return tst
}

func (s *PowerExpressionContext) BitwiseExpression(i int) IBitwiseExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBitwiseExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBitwiseExpressionContext)
}

func (s *PowerExpressionContext) AllDOUBLE_STAR() []antlr.TerminalNode {
	return s.GetTokens(PuppetoParserDOUBLE_STAR)
}

func (s *PowerExpressionContext) DOUBLE_STAR(i int) antlr.TerminalNode {
	return s.GetToken(PuppetoParserDOUBLE_STAR, i)
}

func (s *PowerExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PowerExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PowerExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PuppetoParserVisitor:
		return t.VisitPowerExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PuppetoParser) PowerExpression() (localctx IPowerExpressionContext) {
	this := p
	_ = this

	localctx = NewPowerExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, PuppetoParserRULE_powerExpression)

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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(237)
		p.BitwiseExpression()
	}
	p.SetState(242)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(238)
				p.Match(PuppetoParserDOUBLE_STAR)
			}
			{
				p.SetState(239)
				p.BitwiseExpression()
			}

		}
		p.SetState(244)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext())
	}

	return localctx
}

// IBitwiseExpressionContext is an interface to support dynamic dispatch.
type IBitwiseExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllShiftExpression() []IShiftExpressionContext
	ShiftExpression(i int) IShiftExpressionContext
	AllBAND() []antlr.TerminalNode
	BAND(i int) antlr.TerminalNode
	AllBOR() []antlr.TerminalNode
	BOR(i int) antlr.TerminalNode
	AllBXOR() []antlr.TerminalNode
	BXOR(i int) antlr.TerminalNode

	// IsBitwiseExpressionContext differentiates from other interfaces.
	IsBitwiseExpressionContext()
}

type BitwiseExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBitwiseExpressionContext() *BitwiseExpressionContext {
	var p = new(BitwiseExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PuppetoParserRULE_bitwiseExpression
	return p
}

func (*BitwiseExpressionContext) IsBitwiseExpressionContext() {}

func NewBitwiseExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BitwiseExpressionContext {
	var p = new(BitwiseExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PuppetoParserRULE_bitwiseExpression

	return p
}

func (s *BitwiseExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *BitwiseExpressionContext) AllShiftExpression() []IShiftExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IShiftExpressionContext); ok {
			len++
		}
	}

	tst := make([]IShiftExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IShiftExpressionContext); ok {
			tst[i] = t.(IShiftExpressionContext)
			i++
		}
	}

	return tst
}

func (s *BitwiseExpressionContext) ShiftExpression(i int) IShiftExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IShiftExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IShiftExpressionContext)
}

func (s *BitwiseExpressionContext) AllBAND() []antlr.TerminalNode {
	return s.GetTokens(PuppetoParserBAND)
}

func (s *BitwiseExpressionContext) BAND(i int) antlr.TerminalNode {
	return s.GetToken(PuppetoParserBAND, i)
}

func (s *BitwiseExpressionContext) AllBOR() []antlr.TerminalNode {
	return s.GetTokens(PuppetoParserBOR)
}

func (s *BitwiseExpressionContext) BOR(i int) antlr.TerminalNode {
	return s.GetToken(PuppetoParserBOR, i)
}

func (s *BitwiseExpressionContext) AllBXOR() []antlr.TerminalNode {
	return s.GetTokens(PuppetoParserBXOR)
}

func (s *BitwiseExpressionContext) BXOR(i int) antlr.TerminalNode {
	return s.GetToken(PuppetoParserBXOR, i)
}

func (s *BitwiseExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BitwiseExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BitwiseExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PuppetoParserVisitor:
		return t.VisitBitwiseExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PuppetoParser) BitwiseExpression() (localctx IBitwiseExpressionContext) {
	this := p
	_ = this

	localctx = NewBitwiseExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, PuppetoParserRULE_bitwiseExpression)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(245)
		p.ShiftExpression()
	}
	p.SetState(250)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(246)
				_la = p.GetTokenStream().LA(1)

				if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&30064771072) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(247)
				p.ShiftExpression()
			}

		}
		p.SetState(252)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext())
	}

	return localctx
}

// IShiftExpressionContext is an interface to support dynamic dispatch.
type IShiftExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllUnaryExpression() []IUnaryExpressionContext
	UnaryExpression(i int) IUnaryExpressionContext
	AllBLS() []antlr.TerminalNode
	BLS(i int) antlr.TerminalNode
	AllBRS() []antlr.TerminalNode
	BRS(i int) antlr.TerminalNode

	// IsShiftExpressionContext differentiates from other interfaces.
	IsShiftExpressionContext()
}

type ShiftExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyShiftExpressionContext() *ShiftExpressionContext {
	var p = new(ShiftExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PuppetoParserRULE_shiftExpression
	return p
}

func (*ShiftExpressionContext) IsShiftExpressionContext() {}

func NewShiftExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ShiftExpressionContext {
	var p = new(ShiftExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PuppetoParserRULE_shiftExpression

	return p
}

func (s *ShiftExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ShiftExpressionContext) AllUnaryExpression() []IUnaryExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IUnaryExpressionContext); ok {
			len++
		}
	}

	tst := make([]IUnaryExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IUnaryExpressionContext); ok {
			tst[i] = t.(IUnaryExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ShiftExpressionContext) UnaryExpression(i int) IUnaryExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnaryExpressionContext)
}

func (s *ShiftExpressionContext) AllBLS() []antlr.TerminalNode {
	return s.GetTokens(PuppetoParserBLS)
}

func (s *ShiftExpressionContext) BLS(i int) antlr.TerminalNode {
	return s.GetToken(PuppetoParserBLS, i)
}

func (s *ShiftExpressionContext) AllBRS() []antlr.TerminalNode {
	return s.GetTokens(PuppetoParserBRS)
}

func (s *ShiftExpressionContext) BRS(i int) antlr.TerminalNode {
	return s.GetToken(PuppetoParserBRS, i)
}

func (s *ShiftExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ShiftExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ShiftExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PuppetoParserVisitor:
		return t.VisitShiftExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PuppetoParser) ShiftExpression() (localctx IShiftExpressionContext) {
	this := p
	_ = this

	localctx = NewShiftExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, PuppetoParserRULE_shiftExpression)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(253)
		p.UnaryExpression()
	}
	p.SetState(258)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(254)
				_la = p.GetTokenStream().LA(1)

				if !(_la == PuppetoParserBLS || _la == PuppetoParserBRS) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(255)
				p.UnaryExpression()
			}

		}
		p.SetState(260)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext())
	}

	return localctx
}

// IUnaryExpressionContext is an interface to support dynamic dispatch.
type IUnaryExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	UnaryExpression() IUnaryExpressionContext
	NOT() antlr.TerminalNode
	MINUS() antlr.TerminalNode
	ChainExpression() IChainExpressionContext

	// IsUnaryExpressionContext differentiates from other interfaces.
	IsUnaryExpressionContext()
}

type UnaryExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnaryExpressionContext() *UnaryExpressionContext {
	var p = new(UnaryExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PuppetoParserRULE_unaryExpression
	return p
}

func (*UnaryExpressionContext) IsUnaryExpressionContext() {}

func NewUnaryExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryExpressionContext {
	var p = new(UnaryExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PuppetoParserRULE_unaryExpression

	return p
}

func (s *UnaryExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *UnaryExpressionContext) UnaryExpression() IUnaryExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnaryExpressionContext)
}

func (s *UnaryExpressionContext) NOT() antlr.TerminalNode {
	return s.GetToken(PuppetoParserNOT, 0)
}

func (s *UnaryExpressionContext) MINUS() antlr.TerminalNode {
	return s.GetToken(PuppetoParserMINUS, 0)
}

func (s *UnaryExpressionContext) ChainExpression() IChainExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IChainExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IChainExpressionContext)
}

func (s *UnaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnaryExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PuppetoParserVisitor:
		return t.VisitUnaryExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PuppetoParser) UnaryExpression() (localctx IUnaryExpressionContext) {
	this := p
	_ = this

	localctx = NewUnaryExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, PuppetoParserRULE_unaryExpression)
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

	p.SetState(264)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PuppetoParserMINUS, PuppetoParserNOT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(261)
			_la = p.GetTokenStream().LA(1)

			if !(_la == PuppetoParserMINUS || _la == PuppetoParserNOT) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(262)
			p.UnaryExpression()
		}

	case PuppetoParserNAN, PuppetoParserL_CURLY, PuppetoParserL_SQUARE, PuppetoParserFN, PuppetoParserL_PARENTHES, PuppetoParserFLOAT, PuppetoParserINTEGER, PuppetoParserNIL, PuppetoParserBOOLEAN, PuppetoParserIDENTIFIER, PuppetoParserSTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(263)
			p.ChainExpression()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IGetPropertyContext is an interface to support dynamic dispatch.
type IGetPropertyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Value() IValueContext
	L_SQUARE() antlr.TerminalNode
	Expression() IExpressionContext
	R_SQUARE() antlr.TerminalNode
	DOT() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode

	// IsGetPropertyContext differentiates from other interfaces.
	IsGetPropertyContext()
}

type GetPropertyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGetPropertyContext() *GetPropertyContext {
	var p = new(GetPropertyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PuppetoParserRULE_getProperty
	return p
}

func (*GetPropertyContext) IsGetPropertyContext() {}

func NewGetPropertyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GetPropertyContext {
	var p = new(GetPropertyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PuppetoParserRULE_getProperty

	return p
}

func (s *GetPropertyContext) GetParser() antlr.Parser { return s.parser }

func (s *GetPropertyContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *GetPropertyContext) L_SQUARE() antlr.TerminalNode {
	return s.GetToken(PuppetoParserL_SQUARE, 0)
}

func (s *GetPropertyContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *GetPropertyContext) R_SQUARE() antlr.TerminalNode {
	return s.GetToken(PuppetoParserR_SQUARE, 0)
}

func (s *GetPropertyContext) DOT() antlr.TerminalNode {
	return s.GetToken(PuppetoParserDOT, 0)
}

func (s *GetPropertyContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(PuppetoParserIDENTIFIER, 0)
}

func (s *GetPropertyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GetPropertyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GetPropertyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PuppetoParserVisitor:
		return t.VisitGetProperty(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PuppetoParser) GetProperty() (localctx IGetPropertyContext) {
	this := p
	_ = this

	localctx = NewGetPropertyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, PuppetoParserRULE_getProperty)

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
		p.SetState(266)
		p.Value()
	}
	p.SetState(273)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PuppetoParserL_SQUARE:
		{
			p.SetState(267)
			p.Match(PuppetoParserL_SQUARE)
		}
		{
			p.SetState(268)
			p.Expression()
		}
		{
			p.SetState(269)
			p.Match(PuppetoParserR_SQUARE)
		}

	case PuppetoParserDOT:
		{
			p.SetState(271)
			p.Match(PuppetoParserDOT)
		}
		{
			p.SetState(272)
			p.Match(PuppetoParserIDENTIFIER)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ISetPropertyContext is an interface to support dynamic dispatch.
type ISetPropertyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	GetProperty() IGetPropertyContext
	Expression() IExpressionContext
	TWO_SIDES_OPERATOR() antlr.TerminalNode
	TO() antlr.TerminalNode

	// IsSetPropertyContext differentiates from other interfaces.
	IsSetPropertyContext()
}

type SetPropertyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySetPropertyContext() *SetPropertyContext {
	var p = new(SetPropertyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PuppetoParserRULE_setProperty
	return p
}

func (*SetPropertyContext) IsSetPropertyContext() {}

func NewSetPropertyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SetPropertyContext {
	var p = new(SetPropertyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PuppetoParserRULE_setProperty

	return p
}

func (s *SetPropertyContext) GetParser() antlr.Parser { return s.parser }

func (s *SetPropertyContext) GetProperty() IGetPropertyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGetPropertyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGetPropertyContext)
}

func (s *SetPropertyContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SetPropertyContext) TWO_SIDES_OPERATOR() antlr.TerminalNode {
	return s.GetToken(PuppetoParserTWO_SIDES_OPERATOR, 0)
}

func (s *SetPropertyContext) TO() antlr.TerminalNode {
	return s.GetToken(PuppetoParserTO, 0)
}

func (s *SetPropertyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SetPropertyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SetPropertyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PuppetoParserVisitor:
		return t.VisitSetProperty(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PuppetoParser) SetProperty() (localctx ISetPropertyContext) {
	this := p
	_ = this

	localctx = NewSetPropertyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, PuppetoParserRULE_setProperty)
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
		p.SetState(275)
		p.GetProperty()
	}
	{
		p.SetState(276)
		_la = p.GetTokenStream().LA(1)

		if !(_la == PuppetoParserTWO_SIDES_OPERATOR || _la == PuppetoParserTO) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(277)
		p.Expression()
	}

	return localctx
}

// ICallExpressionContext is an interface to support dynamic dispatch.
type ICallExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Value() IValueContext
	L_PARENTHES() antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	R_PARENTHES() antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsCallExpressionContext differentiates from other interfaces.
	IsCallExpressionContext()
}

type CallExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCallExpressionContext() *CallExpressionContext {
	var p = new(CallExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PuppetoParserRULE_callExpression
	return p
}

func (*CallExpressionContext) IsCallExpressionContext() {}

func NewCallExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CallExpressionContext {
	var p = new(CallExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PuppetoParserRULE_callExpression

	return p
}

func (s *CallExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *CallExpressionContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *CallExpressionContext) L_PARENTHES() antlr.TerminalNode {
	return s.GetToken(PuppetoParserL_PARENTHES, 0)
}

func (s *CallExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *CallExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *CallExpressionContext) R_PARENTHES() antlr.TerminalNode {
	return s.GetToken(PuppetoParserR_PARENTHES, 0)
}

func (s *CallExpressionContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(PuppetoParserCOMMA)
}

func (s *CallExpressionContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(PuppetoParserCOMMA, i)
}

func (s *CallExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CallExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PuppetoParserVisitor:
		return t.VisitCallExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PuppetoParser) CallExpression() (localctx ICallExpressionContext) {
	this := p
	_ = this

	localctx = NewCallExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, PuppetoParserRULE_callExpression)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(279)
		p.Value()
	}
	{
		p.SetState(280)
		p.Match(PuppetoParserL_PARENTHES)
	}
	{
		p.SetState(281)
		p.Expression()
	}
	p.SetState(286)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(282)
				p.Match(PuppetoParserCOMMA)
			}
			{
				p.SetState(283)
				p.Expression()
			}

		}
		p.SetState(288)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext())
	}
	p.SetState(290)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PuppetoParserCOMMA {
		{
			p.SetState(289)
			p.Match(PuppetoParserCOMMA)
		}

	}
	{
		p.SetState(292)
		p.Match(PuppetoParserR_PARENTHES)
	}

	return localctx
}

// IChainExpressionContext is an interface to support dynamic dispatch.
type IChainExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	GetProperty() IGetPropertyContext
	SetProperty() ISetPropertyContext
	CallExpression() ICallExpressionContext
	Value() IValueContext

	// IsChainExpressionContext differentiates from other interfaces.
	IsChainExpressionContext()
}

type ChainExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyChainExpressionContext() *ChainExpressionContext {
	var p = new(ChainExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PuppetoParserRULE_chainExpression
	return p
}

func (*ChainExpressionContext) IsChainExpressionContext() {}

func NewChainExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ChainExpressionContext {
	var p = new(ChainExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PuppetoParserRULE_chainExpression

	return p
}

func (s *ChainExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ChainExpressionContext) GetProperty() IGetPropertyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGetPropertyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGetPropertyContext)
}

func (s *ChainExpressionContext) SetProperty() ISetPropertyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISetPropertyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISetPropertyContext)
}

func (s *ChainExpressionContext) CallExpression() ICallExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICallExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICallExpressionContext)
}

func (s *ChainExpressionContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *ChainExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ChainExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ChainExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PuppetoParserVisitor:
		return t.VisitChainExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PuppetoParser) ChainExpression() (localctx IChainExpressionContext) {
	this := p
	_ = this

	localctx = NewChainExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, PuppetoParserRULE_chainExpression)

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

	p.SetState(298)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(294)
			p.GetProperty()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(295)
			p.SetProperty()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(296)
			p.CallExpression()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(297)
			p.Value()
		}

	}

	return localctx
}

// ILhsVariableAssignationContext is an interface to support dynamic dispatch.
type ILhsVariableAssignationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ONE_SIDE_OPERATOR() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode

	// IsLhsVariableAssignationContext differentiates from other interfaces.
	IsLhsVariableAssignationContext()
}

type LhsVariableAssignationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLhsVariableAssignationContext() *LhsVariableAssignationContext {
	var p = new(LhsVariableAssignationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PuppetoParserRULE_lhsVariableAssignation
	return p
}

func (*LhsVariableAssignationContext) IsLhsVariableAssignationContext() {}

func NewLhsVariableAssignationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LhsVariableAssignationContext {
	var p = new(LhsVariableAssignationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PuppetoParserRULE_lhsVariableAssignation

	return p
}

func (s *LhsVariableAssignationContext) GetParser() antlr.Parser { return s.parser }

func (s *LhsVariableAssignationContext) ONE_SIDE_OPERATOR() antlr.TerminalNode {
	return s.GetToken(PuppetoParserONE_SIDE_OPERATOR, 0)
}

func (s *LhsVariableAssignationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(PuppetoParserIDENTIFIER, 0)
}

func (s *LhsVariableAssignationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LhsVariableAssignationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LhsVariableAssignationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PuppetoParserVisitor:
		return t.VisitLhsVariableAssignation(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PuppetoParser) LhsVariableAssignation() (localctx ILhsVariableAssignationContext) {
	this := p
	_ = this

	localctx = NewLhsVariableAssignationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, PuppetoParserRULE_lhsVariableAssignation)

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
		p.SetState(300)
		p.Match(PuppetoParserONE_SIDE_OPERATOR)
	}
	{
		p.SetState(301)
		p.Match(PuppetoParserIDENTIFIER)
	}

	return localctx
}

// IRhsVariableAssignationContext is an interface to support dynamic dispatch.
type IRhsVariableAssignationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	ONE_SIDE_OPERATOR() antlr.TerminalNode

	// IsRhsVariableAssignationContext differentiates from other interfaces.
	IsRhsVariableAssignationContext()
}

type RhsVariableAssignationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRhsVariableAssignationContext() *RhsVariableAssignationContext {
	var p = new(RhsVariableAssignationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PuppetoParserRULE_rhsVariableAssignation
	return p
}

func (*RhsVariableAssignationContext) IsRhsVariableAssignationContext() {}

func NewRhsVariableAssignationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RhsVariableAssignationContext {
	var p = new(RhsVariableAssignationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PuppetoParserRULE_rhsVariableAssignation

	return p
}

func (s *RhsVariableAssignationContext) GetParser() antlr.Parser { return s.parser }

func (s *RhsVariableAssignationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(PuppetoParserIDENTIFIER, 0)
}

func (s *RhsVariableAssignationContext) ONE_SIDE_OPERATOR() antlr.TerminalNode {
	return s.GetToken(PuppetoParserONE_SIDE_OPERATOR, 0)
}

func (s *RhsVariableAssignationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RhsVariableAssignationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RhsVariableAssignationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PuppetoParserVisitor:
		return t.VisitRhsVariableAssignation(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PuppetoParser) RhsVariableAssignation() (localctx IRhsVariableAssignationContext) {
	this := p
	_ = this

	localctx = NewRhsVariableAssignationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, PuppetoParserRULE_rhsVariableAssignation)

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
		p.SetState(303)
		p.Match(PuppetoParserIDENTIFIER)
	}
	{
		p.SetState(304)
		p.Match(PuppetoParserONE_SIDE_OPERATOR)
	}

	return localctx
}

// IMidVariableAssignationContext is an interface to support dynamic dispatch.
type IMidVariableAssignationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	Expression() IExpressionContext
	TWO_SIDES_OPERATOR() antlr.TerminalNode
	TO() antlr.TerminalNode

	// IsMidVariableAssignationContext differentiates from other interfaces.
	IsMidVariableAssignationContext()
}

type MidVariableAssignationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMidVariableAssignationContext() *MidVariableAssignationContext {
	var p = new(MidVariableAssignationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PuppetoParserRULE_midVariableAssignation
	return p
}

func (*MidVariableAssignationContext) IsMidVariableAssignationContext() {}

func NewMidVariableAssignationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MidVariableAssignationContext {
	var p = new(MidVariableAssignationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PuppetoParserRULE_midVariableAssignation

	return p
}

func (s *MidVariableAssignationContext) GetParser() antlr.Parser { return s.parser }

func (s *MidVariableAssignationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(PuppetoParserIDENTIFIER, 0)
}

func (s *MidVariableAssignationContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *MidVariableAssignationContext) TWO_SIDES_OPERATOR() antlr.TerminalNode {
	return s.GetToken(PuppetoParserTWO_SIDES_OPERATOR, 0)
}

func (s *MidVariableAssignationContext) TO() antlr.TerminalNode {
	return s.GetToken(PuppetoParserTO, 0)
}

func (s *MidVariableAssignationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MidVariableAssignationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MidVariableAssignationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PuppetoParserVisitor:
		return t.VisitMidVariableAssignation(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PuppetoParser) MidVariableAssignation() (localctx IMidVariableAssignationContext) {
	this := p
	_ = this

	localctx = NewMidVariableAssignationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, PuppetoParserRULE_midVariableAssignation)
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
		p.SetState(306)
		p.Match(PuppetoParserIDENTIFIER)
	}
	{
		p.SetState(307)
		_la = p.GetTokenStream().LA(1)

		if !(_la == PuppetoParserTWO_SIDES_OPERATOR || _la == PuppetoParserTO) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(308)
		p.Expression()
	}

	return localctx
}

// IVariableAssignationContext is an interface to support dynamic dispatch.
type IVariableAssignationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllMidVariableAssignation() []IMidVariableAssignationContext
	MidVariableAssignation(i int) IMidVariableAssignationContext
	AllLhsVariableAssignation() []ILhsVariableAssignationContext
	LhsVariableAssignation(i int) ILhsVariableAssignationContext
	AllRhsVariableAssignation() []IRhsVariableAssignationContext
	RhsVariableAssignation(i int) IRhsVariableAssignationContext

	// IsVariableAssignationContext differentiates from other interfaces.
	IsVariableAssignationContext()
}

type VariableAssignationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableAssignationContext() *VariableAssignationContext {
	var p = new(VariableAssignationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PuppetoParserRULE_variableAssignation
	return p
}

func (*VariableAssignationContext) IsVariableAssignationContext() {}

func NewVariableAssignationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableAssignationContext {
	var p = new(VariableAssignationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PuppetoParserRULE_variableAssignation

	return p
}

func (s *VariableAssignationContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableAssignationContext) AllMidVariableAssignation() []IMidVariableAssignationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMidVariableAssignationContext); ok {
			len++
		}
	}

	tst := make([]IMidVariableAssignationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMidVariableAssignationContext); ok {
			tst[i] = t.(IMidVariableAssignationContext)
			i++
		}
	}

	return tst
}

func (s *VariableAssignationContext) MidVariableAssignation(i int) IMidVariableAssignationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMidVariableAssignationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMidVariableAssignationContext)
}

func (s *VariableAssignationContext) AllLhsVariableAssignation() []ILhsVariableAssignationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILhsVariableAssignationContext); ok {
			len++
		}
	}

	tst := make([]ILhsVariableAssignationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILhsVariableAssignationContext); ok {
			tst[i] = t.(ILhsVariableAssignationContext)
			i++
		}
	}

	return tst
}

func (s *VariableAssignationContext) LhsVariableAssignation(i int) ILhsVariableAssignationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILhsVariableAssignationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILhsVariableAssignationContext)
}

func (s *VariableAssignationContext) AllRhsVariableAssignation() []IRhsVariableAssignationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRhsVariableAssignationContext); ok {
			len++
		}
	}

	tst := make([]IRhsVariableAssignationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRhsVariableAssignationContext); ok {
			tst[i] = t.(IRhsVariableAssignationContext)
			i++
		}
	}

	return tst
}

func (s *VariableAssignationContext) RhsVariableAssignation(i int) IRhsVariableAssignationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRhsVariableAssignationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRhsVariableAssignationContext)
}

func (s *VariableAssignationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableAssignationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableAssignationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PuppetoParserVisitor:
		return t.VisitVariableAssignation(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PuppetoParser) VariableAssignation() (localctx IVariableAssignationContext) {
	this := p
	_ = this

	localctx = NewVariableAssignationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, PuppetoParserRULE_variableAssignation)

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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(313)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			p.SetState(313)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext()) {
			case 1:
				{
					p.SetState(310)
					p.MidVariableAssignation()
				}

			case 2:
				{
					p.SetState(311)
					p.LhsVariableAssignation()
				}

			case 3:
				{
					p.SetState(312)
					p.RhsVariableAssignation()
				}

			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(315)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext())
	}

	return localctx
}

// IVariableDefinitionContext is an interface to support dynamic dispatch.
type IVariableDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LET() antlr.TerminalNode
	AllVariableDefinitionBody() []IVariableDefinitionBodyContext
	VariableDefinitionBody(i int) IVariableDefinitionBodyContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsVariableDefinitionContext differentiates from other interfaces.
	IsVariableDefinitionContext()
}

type VariableDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableDefinitionContext() *VariableDefinitionContext {
	var p = new(VariableDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PuppetoParserRULE_variableDefinition
	return p
}

func (*VariableDefinitionContext) IsVariableDefinitionContext() {}

func NewVariableDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDefinitionContext {
	var p = new(VariableDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PuppetoParserRULE_variableDefinition

	return p
}

func (s *VariableDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableDefinitionContext) LET() antlr.TerminalNode {
	return s.GetToken(PuppetoParserLET, 0)
}

func (s *VariableDefinitionContext) AllVariableDefinitionBody() []IVariableDefinitionBodyContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVariableDefinitionBodyContext); ok {
			len++
		}
	}

	tst := make([]IVariableDefinitionBodyContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVariableDefinitionBodyContext); ok {
			tst[i] = t.(IVariableDefinitionBodyContext)
			i++
		}
	}

	return tst
}

func (s *VariableDefinitionContext) VariableDefinitionBody(i int) IVariableDefinitionBodyContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableDefinitionBodyContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableDefinitionBodyContext)
}

func (s *VariableDefinitionContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(PuppetoParserCOMMA)
}

func (s *VariableDefinitionContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(PuppetoParserCOMMA, i)
}

func (s *VariableDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PuppetoParserVisitor:
		return t.VisitVariableDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PuppetoParser) VariableDefinition() (localctx IVariableDefinitionContext) {
	this := p
	_ = this

	localctx = NewVariableDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, PuppetoParserRULE_variableDefinition)
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
		p.SetState(317)
		p.Match(PuppetoParserLET)
	}
	{
		p.SetState(318)
		p.VariableDefinitionBody()
	}
	p.SetState(323)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == PuppetoParserCOMMA {
		{
			p.SetState(319)
			p.Match(PuppetoParserCOMMA)
		}
		{
			p.SetState(320)
			p.VariableDefinitionBody()
		}

		p.SetState(325)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IVariableDefinitionBodyContext is an interface to support dynamic dispatch.
type IVariableDefinitionBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	TO() antlr.TerminalNode
	Expression() IExpressionContext

	// IsVariableDefinitionBodyContext differentiates from other interfaces.
	IsVariableDefinitionBodyContext()
}

type VariableDefinitionBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableDefinitionBodyContext() *VariableDefinitionBodyContext {
	var p = new(VariableDefinitionBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PuppetoParserRULE_variableDefinitionBody
	return p
}

func (*VariableDefinitionBodyContext) IsVariableDefinitionBodyContext() {}

func NewVariableDefinitionBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDefinitionBodyContext {
	var p = new(VariableDefinitionBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PuppetoParserRULE_variableDefinitionBody

	return p
}

func (s *VariableDefinitionBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableDefinitionBodyContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(PuppetoParserIDENTIFIER, 0)
}

func (s *VariableDefinitionBodyContext) TO() antlr.TerminalNode {
	return s.GetToken(PuppetoParserTO, 0)
}

func (s *VariableDefinitionBodyContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *VariableDefinitionBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDefinitionBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableDefinitionBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PuppetoParserVisitor:
		return t.VisitVariableDefinitionBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PuppetoParser) VariableDefinitionBody() (localctx IVariableDefinitionBodyContext) {
	this := p
	_ = this

	localctx = NewVariableDefinitionBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, PuppetoParserRULE_variableDefinitionBody)
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
		p.SetState(326)
		p.Match(PuppetoParserIDENTIFIER)
	}
	p.SetState(329)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PuppetoParserTO {
		{
			p.SetState(327)
			p.Match(PuppetoParserTO)
		}
		{
			p.SetState(328)
			p.Expression()
		}

	}

	return localctx
}

// IIfStatementContext is an interface to support dynamic dispatch.
type IIfStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIF() []antlr.TerminalNode
	IF(i int) antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllScopeBody() []IScopeBodyContext
	ScopeBody(i int) IScopeBodyContext
	AllELSE() []antlr.TerminalNode
	ELSE(i int) antlr.TerminalNode

	// IsIfStatementContext differentiates from other interfaces.
	IsIfStatementContext()
}

type IfStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStatementContext() *IfStatementContext {
	var p = new(IfStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PuppetoParserRULE_ifStatement
	return p
}

func (*IfStatementContext) IsIfStatementContext() {}

func NewIfStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStatementContext {
	var p = new(IfStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PuppetoParserRULE_ifStatement

	return p
}

func (s *IfStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStatementContext) AllIF() []antlr.TerminalNode {
	return s.GetTokens(PuppetoParserIF)
}

func (s *IfStatementContext) IF(i int) antlr.TerminalNode {
	return s.GetToken(PuppetoParserIF, i)
}

func (s *IfStatementContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *IfStatementContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IfStatementContext) AllScopeBody() []IScopeBodyContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IScopeBodyContext); ok {
			len++
		}
	}

	tst := make([]IScopeBodyContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IScopeBodyContext); ok {
			tst[i] = t.(IScopeBodyContext)
			i++
		}
	}

	return tst
}

func (s *IfStatementContext) ScopeBody(i int) IScopeBodyContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IScopeBodyContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IScopeBodyContext)
}

func (s *IfStatementContext) AllELSE() []antlr.TerminalNode {
	return s.GetTokens(PuppetoParserELSE)
}

func (s *IfStatementContext) ELSE(i int) antlr.TerminalNode {
	return s.GetToken(PuppetoParserELSE, i)
}

func (s *IfStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PuppetoParserVisitor:
		return t.VisitIfStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PuppetoParser) IfStatement() (localctx IIfStatementContext) {
	this := p
	_ = this

	localctx = NewIfStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, PuppetoParserRULE_ifStatement)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(331)
		p.Match(PuppetoParserIF)
	}
	{
		p.SetState(332)
		p.Expression()
	}
	{
		p.SetState(333)
		p.ScopeBody()
	}
	p.SetState(341)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 34, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(334)
				p.Match(PuppetoParserELSE)
			}
			{
				p.SetState(335)
				p.Match(PuppetoParserIF)
			}
			{
				p.SetState(336)
				p.Expression()
			}
			{
				p.SetState(337)
				p.ScopeBody()
			}

		}
		p.SetState(343)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 34, p.GetParserRuleContext())
	}
	p.SetState(346)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PuppetoParserELSE {
		{
			p.SetState(344)
			p.Match(PuppetoParserELSE)
		}
		{
			p.SetState(345)
			p.ScopeBody()
		}

	}

	return localctx
}

// ILoopStatementContext is an interface to support dynamic dispatch.
type ILoopStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FOR() antlr.TerminalNode
	Expression() IExpressionContext
	ProgramBodyWithBreakContinue() IProgramBodyWithBreakContinueContext

	// IsLoopStatementContext differentiates from other interfaces.
	IsLoopStatementContext()
}

type LoopStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLoopStatementContext() *LoopStatementContext {
	var p = new(LoopStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PuppetoParserRULE_loopStatement
	return p
}

func (*LoopStatementContext) IsLoopStatementContext() {}

func NewLoopStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LoopStatementContext {
	var p = new(LoopStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PuppetoParserRULE_loopStatement

	return p
}

func (s *LoopStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *LoopStatementContext) FOR() antlr.TerminalNode {
	return s.GetToken(PuppetoParserFOR, 0)
}

func (s *LoopStatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LoopStatementContext) ProgramBodyWithBreakContinue() IProgramBodyWithBreakContinueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProgramBodyWithBreakContinueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProgramBodyWithBreakContinueContext)
}

func (s *LoopStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LoopStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LoopStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PuppetoParserVisitor:
		return t.VisitLoopStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PuppetoParser) LoopStatement() (localctx ILoopStatementContext) {
	this := p
	_ = this

	localctx = NewLoopStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, PuppetoParserRULE_loopStatement)

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
		p.SetState(348)
		p.Match(PuppetoParserFOR)
	}
	{
		p.SetState(349)
		p.Expression()
	}
	{
		p.SetState(350)
		p.ProgramBodyWithBreakContinue()
	}

	return localctx
}

// ISwitchStatementContext is an interface to support dynamic dispatch.
type ISwitchStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SWITCH() antlr.TerminalNode
	Expression() IExpressionContext
	L_CURLY() antlr.TerminalNode
	R_CURLY() antlr.TerminalNode
	AllSwitchCase() []ISwitchCaseContext
	SwitchCase(i int) ISwitchCaseContext
	DEFAULT() antlr.TerminalNode
	COLON() antlr.TerminalNode
	AllProgramBodyWithBreakContinue() []IProgramBodyWithBreakContinueContext
	ProgramBodyWithBreakContinue(i int) IProgramBodyWithBreakContinueContext

	// IsSwitchStatementContext differentiates from other interfaces.
	IsSwitchStatementContext()
}

type SwitchStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitchStatementContext() *SwitchStatementContext {
	var p = new(SwitchStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PuppetoParserRULE_switchStatement
	return p
}

func (*SwitchStatementContext) IsSwitchStatementContext() {}

func NewSwitchStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwitchStatementContext {
	var p = new(SwitchStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PuppetoParserRULE_switchStatement

	return p
}

func (s *SwitchStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *SwitchStatementContext) SWITCH() antlr.TerminalNode {
	return s.GetToken(PuppetoParserSWITCH, 0)
}

func (s *SwitchStatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SwitchStatementContext) L_CURLY() antlr.TerminalNode {
	return s.GetToken(PuppetoParserL_CURLY, 0)
}

func (s *SwitchStatementContext) R_CURLY() antlr.TerminalNode {
	return s.GetToken(PuppetoParserR_CURLY, 0)
}

func (s *SwitchStatementContext) AllSwitchCase() []ISwitchCaseContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISwitchCaseContext); ok {
			len++
		}
	}

	tst := make([]ISwitchCaseContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISwitchCaseContext); ok {
			tst[i] = t.(ISwitchCaseContext)
			i++
		}
	}

	return tst
}

func (s *SwitchStatementContext) SwitchCase(i int) ISwitchCaseContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISwitchCaseContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISwitchCaseContext)
}

func (s *SwitchStatementContext) DEFAULT() antlr.TerminalNode {
	return s.GetToken(PuppetoParserDEFAULT, 0)
}

func (s *SwitchStatementContext) COLON() antlr.TerminalNode {
	return s.GetToken(PuppetoParserCOLON, 0)
}

func (s *SwitchStatementContext) AllProgramBodyWithBreakContinue() []IProgramBodyWithBreakContinueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IProgramBodyWithBreakContinueContext); ok {
			len++
		}
	}

	tst := make([]IProgramBodyWithBreakContinueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IProgramBodyWithBreakContinueContext); ok {
			tst[i] = t.(IProgramBodyWithBreakContinueContext)
			i++
		}
	}

	return tst
}

func (s *SwitchStatementContext) ProgramBodyWithBreakContinue(i int) IProgramBodyWithBreakContinueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProgramBodyWithBreakContinueContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProgramBodyWithBreakContinueContext)
}

func (s *SwitchStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SwitchStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PuppetoParserVisitor:
		return t.VisitSwitchStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PuppetoParser) SwitchStatement() (localctx ISwitchStatementContext) {
	this := p
	_ = this

	localctx = NewSwitchStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, PuppetoParserRULE_switchStatement)
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
		p.SetState(352)
		p.Match(PuppetoParserSWITCH)
	}
	{
		p.SetState(353)
		p.Expression()
	}
	{
		p.SetState(354)
		p.Match(PuppetoParserL_CURLY)
	}
	p.SetState(358)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == PuppetoParserCASE {
		{
			p.SetState(355)
			p.SwitchCase()
		}

		p.SetState(360)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(369)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PuppetoParserDEFAULT {
		{
			p.SetState(361)
			p.Match(PuppetoParserDEFAULT)
		}
		{
			p.SetState(362)
			p.Match(PuppetoParserCOLON)
		}
		p.SetState(366)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&71104730317300036) != 0 {
			{
				p.SetState(363)
				p.ProgramBodyWithBreakContinue()
			}

			p.SetState(368)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(371)
		p.Match(PuppetoParserR_CURLY)
	}

	return localctx
}

// ISwitchCaseContext is an interface to support dynamic dispatch.
type ISwitchCaseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CASE() antlr.TerminalNode
	Expression() IExpressionContext
	COLON() antlr.TerminalNode
	ProgramBodyWithBreakContinue() IProgramBodyWithBreakContinueContext

	// IsSwitchCaseContext differentiates from other interfaces.
	IsSwitchCaseContext()
}

type SwitchCaseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitchCaseContext() *SwitchCaseContext {
	var p = new(SwitchCaseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PuppetoParserRULE_switchCase
	return p
}

func (*SwitchCaseContext) IsSwitchCaseContext() {}

func NewSwitchCaseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwitchCaseContext {
	var p = new(SwitchCaseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PuppetoParserRULE_switchCase

	return p
}

func (s *SwitchCaseContext) GetParser() antlr.Parser { return s.parser }

func (s *SwitchCaseContext) CASE() antlr.TerminalNode {
	return s.GetToken(PuppetoParserCASE, 0)
}

func (s *SwitchCaseContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SwitchCaseContext) COLON() antlr.TerminalNode {
	return s.GetToken(PuppetoParserCOLON, 0)
}

func (s *SwitchCaseContext) ProgramBodyWithBreakContinue() IProgramBodyWithBreakContinueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProgramBodyWithBreakContinueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProgramBodyWithBreakContinueContext)
}

func (s *SwitchCaseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchCaseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SwitchCaseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PuppetoParserVisitor:
		return t.VisitSwitchCase(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PuppetoParser) SwitchCase() (localctx ISwitchCaseContext) {
	this := p
	_ = this

	localctx = NewSwitchCaseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, PuppetoParserRULE_switchCase)

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
		p.SetState(373)
		p.Match(PuppetoParserCASE)
	}
	{
		p.SetState(374)
		p.Expression()
	}
	{
		p.SetState(375)
		p.Match(PuppetoParserCOLON)
	}
	{
		p.SetState(376)
		p.ProgramBodyWithBreakContinue()
	}

	return localctx
}

// IBreakStatementContext is an interface to support dynamic dispatch.
type IBreakStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BREAK() antlr.TerminalNode

	// IsBreakStatementContext differentiates from other interfaces.
	IsBreakStatementContext()
}

type BreakStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBreakStatementContext() *BreakStatementContext {
	var p = new(BreakStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PuppetoParserRULE_breakStatement
	return p
}

func (*BreakStatementContext) IsBreakStatementContext() {}

func NewBreakStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BreakStatementContext {
	var p = new(BreakStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PuppetoParserRULE_breakStatement

	return p
}

func (s *BreakStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *BreakStatementContext) BREAK() antlr.TerminalNode {
	return s.GetToken(PuppetoParserBREAK, 0)
}

func (s *BreakStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BreakStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BreakStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PuppetoParserVisitor:
		return t.VisitBreakStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PuppetoParser) BreakStatement() (localctx IBreakStatementContext) {
	this := p
	_ = this

	localctx = NewBreakStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, PuppetoParserRULE_breakStatement)

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
		p.SetState(378)
		p.Match(PuppetoParserBREAK)
	}

	return localctx
}

// IContinueStatementContext is an interface to support dynamic dispatch.
type IContinueStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CONTINUE() antlr.TerminalNode

	// IsContinueStatementContext differentiates from other interfaces.
	IsContinueStatementContext()
}

type ContinueStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContinueStatementContext() *ContinueStatementContext {
	var p = new(ContinueStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PuppetoParserRULE_continueStatement
	return p
}

func (*ContinueStatementContext) IsContinueStatementContext() {}

func NewContinueStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ContinueStatementContext {
	var p = new(ContinueStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PuppetoParserRULE_continueStatement

	return p
}

func (s *ContinueStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ContinueStatementContext) CONTINUE() antlr.TerminalNode {
	return s.GetToken(PuppetoParserCONTINUE, 0)
}

func (s *ContinueStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContinueStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ContinueStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PuppetoParserVisitor:
		return t.VisitContinueStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PuppetoParser) ContinueStatement() (localctx IContinueStatementContext) {
	this := p
	_ = this

	localctx = NewContinueStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, PuppetoParserRULE_continueStatement)

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
		p.SetState(380)
		p.Match(PuppetoParserCONTINUE)
	}

	return localctx
}

// ITryStatementContext is an interface to support dynamic dispatch.
type ITryStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TRY() antlr.TerminalNode
	AllProgramBody() []IProgramBodyContext
	ProgramBody(i int) IProgramBodyContext
	CATCH() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode

	// IsTryStatementContext differentiates from other interfaces.
	IsTryStatementContext()
}

type TryStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTryStatementContext() *TryStatementContext {
	var p = new(TryStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PuppetoParserRULE_tryStatement
	return p
}

func (*TryStatementContext) IsTryStatementContext() {}

func NewTryStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TryStatementContext {
	var p = new(TryStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PuppetoParserRULE_tryStatement

	return p
}

func (s *TryStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *TryStatementContext) TRY() antlr.TerminalNode {
	return s.GetToken(PuppetoParserTRY, 0)
}

func (s *TryStatementContext) AllProgramBody() []IProgramBodyContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IProgramBodyContext); ok {
			len++
		}
	}

	tst := make([]IProgramBodyContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IProgramBodyContext); ok {
			tst[i] = t.(IProgramBodyContext)
			i++
		}
	}

	return tst
}

func (s *TryStatementContext) ProgramBody(i int) IProgramBodyContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProgramBodyContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProgramBodyContext)
}

func (s *TryStatementContext) CATCH() antlr.TerminalNode {
	return s.GetToken(PuppetoParserCATCH, 0)
}

func (s *TryStatementContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(PuppetoParserIDENTIFIER, 0)
}

func (s *TryStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TryStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TryStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PuppetoParserVisitor:
		return t.VisitTryStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PuppetoParser) TryStatement() (localctx ITryStatementContext) {
	this := p
	_ = this

	localctx = NewTryStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, PuppetoParserRULE_tryStatement)

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
		p.SetState(382)
		p.Match(PuppetoParserTRY)
	}
	{
		p.SetState(383)
		p.ProgramBody()
	}
	p.SetState(389)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 40, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(384)
			p.Match(PuppetoParserCATCH)
		}
		p.SetState(386)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(385)
				p.Match(PuppetoParserIDENTIFIER)
			}

		}
		{
			p.SetState(388)
			p.ProgramBody()
		}

	}

	return localctx
}

// IStatementListContext is an interface to support dynamic dispatch.
type IStatementListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VariableDefinition() IVariableDefinitionContext
	FunctionDefinition() IFunctionDefinitionContext
	IfStatement() IIfStatementContext
	LoopStatement() ILoopStatementContext
	SwitchStatement() ISwitchStatementContext
	TryStatement() ITryStatementContext
	EchoStatement() IEchoStatementContext
	VariableAssignation() IVariableAssignationContext

	// IsStatementListContext differentiates from other interfaces.
	IsStatementListContext()
}

type StatementListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementListContext() *StatementListContext {
	var p = new(StatementListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PuppetoParserRULE_statementList
	return p
}

func (*StatementListContext) IsStatementListContext() {}

func NewStatementListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementListContext {
	var p = new(StatementListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PuppetoParserRULE_statementList

	return p
}

func (s *StatementListContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementListContext) VariableDefinition() IVariableDefinitionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableDefinitionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableDefinitionContext)
}

func (s *StatementListContext) FunctionDefinition() IFunctionDefinitionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionDefinitionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionDefinitionContext)
}

func (s *StatementListContext) IfStatement() IIfStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStatementContext)
}

func (s *StatementListContext) LoopStatement() ILoopStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILoopStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILoopStatementContext)
}

func (s *StatementListContext) SwitchStatement() ISwitchStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISwitchStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISwitchStatementContext)
}

func (s *StatementListContext) TryStatement() ITryStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITryStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITryStatementContext)
}

func (s *StatementListContext) EchoStatement() IEchoStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEchoStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEchoStatementContext)
}

func (s *StatementListContext) VariableAssignation() IVariableAssignationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableAssignationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableAssignationContext)
}

func (s *StatementListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PuppetoParserVisitor:
		return t.VisitStatementList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PuppetoParser) StatementList() (localctx IStatementListContext) {
	this := p
	_ = this

	localctx = NewStatementListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, PuppetoParserRULE_statementList)

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

	p.SetState(399)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PuppetoParserLET:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(391)
			p.VariableDefinition()
		}

	case PuppetoParserFN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(392)
			p.FunctionDefinition()
		}

	case PuppetoParserIF:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(393)
			p.IfStatement()
		}

	case PuppetoParserFOR:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(394)
			p.LoopStatement()
		}

	case PuppetoParserSWITCH:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(395)
			p.SwitchStatement()
		}

	case PuppetoParserTRY:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(396)
			p.TryStatement()
		}

	case PuppetoParserECHO:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(397)
			p.EchoStatement()
		}

	case PuppetoParserONE_SIDE_OPERATOR, PuppetoParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(398)
			p.VariableAssignation()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	StatementList() IStatementListContext
	AllPUP_END_TAG() []antlr.TerminalNode
	PUP_END_TAG(i int) antlr.TerminalNode
	AllPUP_START_TAG() []antlr.TerminalNode
	PUP_START_TAG(i int) antlr.TerminalNode
	AllHtmlList() []IHtmlListContext
	HtmlList(i int) IHtmlListContext

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PuppetoParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PuppetoParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) StatementList() IStatementListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementListContext)
}

func (s *StatementContext) AllPUP_END_TAG() []antlr.TerminalNode {
	return s.GetTokens(PuppetoParserPUP_END_TAG)
}

func (s *StatementContext) PUP_END_TAG(i int) antlr.TerminalNode {
	return s.GetToken(PuppetoParserPUP_END_TAG, i)
}

func (s *StatementContext) AllPUP_START_TAG() []antlr.TerminalNode {
	return s.GetTokens(PuppetoParserPUP_START_TAG)
}

func (s *StatementContext) PUP_START_TAG(i int) antlr.TerminalNode {
	return s.GetToken(PuppetoParserPUP_START_TAG, i)
}

func (s *StatementContext) AllHtmlList() []IHtmlListContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IHtmlListContext); ok {
			len++
		}
	}

	tst := make([]IHtmlListContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IHtmlListContext); ok {
			tst[i] = t.(IHtmlListContext)
			i++
		}
	}

	return tst
}

func (s *StatementContext) HtmlList(i int) IHtmlListContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHtmlListContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IHtmlListContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PuppetoParserVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PuppetoParser) Statement() (localctx IStatementContext) {
	this := p
	_ = this

	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, PuppetoParserRULE_statement)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(414)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PuppetoParserONE_SIDE_OPERATOR, PuppetoParserFOR, PuppetoParserECHO, PuppetoParserFN, PuppetoParserIF, PuppetoParserSWITCH, PuppetoParserTRY, PuppetoParserLET, PuppetoParserIDENTIFIER:
		{
			p.SetState(401)
			p.StatementList()
		}

	case PuppetoParserPUP_END_TAG:
		p.SetState(410)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(402)
					p.Match(PuppetoParserPUP_END_TAG)
				}
				p.SetState(406)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&576460752303423384) != 0 {
					{
						p.SetState(403)
						p.HtmlList()
					}

					p.SetState(408)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)
				}
				{
					p.SetState(409)
					p.Match(PuppetoParserPUP_START_TAG)
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(412)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 43, p.GetParserRuleContext())
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IScopeContext is an interface to support dynamic dispatch.
type IScopeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	L_CURLY() antlr.TerminalNode
	R_CURLY() antlr.TerminalNode
	ProgramBodyList() IProgramBodyListContext

	// IsScopeContext differentiates from other interfaces.
	IsScopeContext()
}

type ScopeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyScopeContext() *ScopeContext {
	var p = new(ScopeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PuppetoParserRULE_scope
	return p
}

func (*ScopeContext) IsScopeContext() {}

func NewScopeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ScopeContext {
	var p = new(ScopeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PuppetoParserRULE_scope

	return p
}

func (s *ScopeContext) GetParser() antlr.Parser { return s.parser }

func (s *ScopeContext) L_CURLY() antlr.TerminalNode {
	return s.GetToken(PuppetoParserL_CURLY, 0)
}

func (s *ScopeContext) R_CURLY() antlr.TerminalNode {
	return s.GetToken(PuppetoParserR_CURLY, 0)
}

func (s *ScopeContext) ProgramBodyList() IProgramBodyListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProgramBodyListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProgramBodyListContext)
}

func (s *ScopeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScopeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ScopeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PuppetoParserVisitor:
		return t.VisitScope(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PuppetoParser) Scope() (localctx IScopeContext) {
	this := p
	_ = this

	localctx = NewScopeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, PuppetoParserRULE_scope)
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
		p.SetState(416)
		p.Match(PuppetoParserL_CURLY)
	}
	p.SetState(418)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&71082740084744516) != 0 {
		{
			p.SetState(417)
			p.ProgramBodyList()
		}

	}
	{
		p.SetState(420)
		p.Match(PuppetoParserR_CURLY)
	}

	return localctx
}

// IScopeWithBreakContinueContext is an interface to support dynamic dispatch.
type IScopeWithBreakContinueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	L_CURLY() antlr.TerminalNode
	R_CURLY() antlr.TerminalNode
	AllProgramBodyWithBreakContinue() []IProgramBodyWithBreakContinueContext
	ProgramBodyWithBreakContinue(i int) IProgramBodyWithBreakContinueContext

	// IsScopeWithBreakContinueContext differentiates from other interfaces.
	IsScopeWithBreakContinueContext()
}

type ScopeWithBreakContinueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyScopeWithBreakContinueContext() *ScopeWithBreakContinueContext {
	var p = new(ScopeWithBreakContinueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PuppetoParserRULE_scopeWithBreakContinue
	return p
}

func (*ScopeWithBreakContinueContext) IsScopeWithBreakContinueContext() {}

func NewScopeWithBreakContinueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ScopeWithBreakContinueContext {
	var p = new(ScopeWithBreakContinueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PuppetoParserRULE_scopeWithBreakContinue

	return p
}

func (s *ScopeWithBreakContinueContext) GetParser() antlr.Parser { return s.parser }

func (s *ScopeWithBreakContinueContext) L_CURLY() antlr.TerminalNode {
	return s.GetToken(PuppetoParserL_CURLY, 0)
}

func (s *ScopeWithBreakContinueContext) R_CURLY() antlr.TerminalNode {
	return s.GetToken(PuppetoParserR_CURLY, 0)
}

func (s *ScopeWithBreakContinueContext) AllProgramBodyWithBreakContinue() []IProgramBodyWithBreakContinueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IProgramBodyWithBreakContinueContext); ok {
			len++
		}
	}

	tst := make([]IProgramBodyWithBreakContinueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IProgramBodyWithBreakContinueContext); ok {
			tst[i] = t.(IProgramBodyWithBreakContinueContext)
			i++
		}
	}

	return tst
}

func (s *ScopeWithBreakContinueContext) ProgramBodyWithBreakContinue(i int) IProgramBodyWithBreakContinueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProgramBodyWithBreakContinueContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProgramBodyWithBreakContinueContext)
}

func (s *ScopeWithBreakContinueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScopeWithBreakContinueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ScopeWithBreakContinueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PuppetoParserVisitor:
		return t.VisitScopeWithBreakContinue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PuppetoParser) ScopeWithBreakContinue() (localctx IScopeWithBreakContinueContext) {
	this := p
	_ = this

	localctx = NewScopeWithBreakContinueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, PuppetoParserRULE_scopeWithBreakContinue)
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
		p.SetState(422)
		p.Match(PuppetoParserL_CURLY)
	}
	p.SetState(426)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&71104730317300036) != 0 {
		{
			p.SetState(423)
			p.ProgramBodyWithBreakContinue()
		}

		p.SetState(428)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(429)
		p.Match(PuppetoParserR_CURLY)
	}

	return localctx
}

// IProgramBodyContext is an interface to support dynamic dispatch.
type IProgramBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Scope() IScopeContext
	Statement() IStatementContext
	Expression() IExpressionContext

	// IsProgramBodyContext differentiates from other interfaces.
	IsProgramBodyContext()
}

type ProgramBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramBodyContext() *ProgramBodyContext {
	var p = new(ProgramBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PuppetoParserRULE_programBody
	return p
}

func (*ProgramBodyContext) IsProgramBodyContext() {}

func NewProgramBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramBodyContext {
	var p = new(ProgramBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PuppetoParserRULE_programBody

	return p
}

func (s *ProgramBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramBodyContext) Scope() IScopeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IScopeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IScopeContext)
}

func (s *ProgramBodyContext) Statement() IStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ProgramBodyContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ProgramBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PuppetoParserVisitor:
		return t.VisitProgramBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PuppetoParser) ProgramBody() (localctx IProgramBodyContext) {
	this := p
	_ = this

	localctx = NewProgramBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, PuppetoParserRULE_programBody)

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

	p.SetState(434)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 47, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(431)
			p.Scope()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(432)
			p.Statement()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(433)
			p.Expression()
		}

	}

	return localctx
}

// IProgramBodyListContext is an interface to support dynamic dispatch.
type IProgramBodyListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllProgramBody() []IProgramBodyContext
	ProgramBody(i int) IProgramBodyContext

	// IsProgramBodyListContext differentiates from other interfaces.
	IsProgramBodyListContext()
}

type ProgramBodyListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramBodyListContext() *ProgramBodyListContext {
	var p = new(ProgramBodyListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PuppetoParserRULE_programBodyList
	return p
}

func (*ProgramBodyListContext) IsProgramBodyListContext() {}

func NewProgramBodyListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramBodyListContext {
	var p = new(ProgramBodyListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PuppetoParserRULE_programBodyList

	return p
}

func (s *ProgramBodyListContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramBodyListContext) AllProgramBody() []IProgramBodyContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IProgramBodyContext); ok {
			len++
		}
	}

	tst := make([]IProgramBodyContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IProgramBodyContext); ok {
			tst[i] = t.(IProgramBodyContext)
			i++
		}
	}

	return tst
}

func (s *ProgramBodyListContext) ProgramBody(i int) IProgramBodyContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProgramBodyContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProgramBodyContext)
}

func (s *ProgramBodyListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramBodyListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramBodyListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PuppetoParserVisitor:
		return t.VisitProgramBodyList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PuppetoParser) ProgramBodyList() (localctx IProgramBodyListContext) {
	this := p
	_ = this

	localctx = NewProgramBodyListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, PuppetoParserRULE_programBodyList)

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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(437)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(436)
				p.ProgramBody()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(439)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 48, p.GetParserRuleContext())
	}

	return localctx
}

// IEchoStatementContext is an interface to support dynamic dispatch.
type IEchoStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ECHO() antlr.TerminalNode
	Expression() IExpressionContext

	// IsEchoStatementContext differentiates from other interfaces.
	IsEchoStatementContext()
}

type EchoStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEchoStatementContext() *EchoStatementContext {
	var p = new(EchoStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PuppetoParserRULE_echoStatement
	return p
}

func (*EchoStatementContext) IsEchoStatementContext() {}

func NewEchoStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EchoStatementContext {
	var p = new(EchoStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PuppetoParserRULE_echoStatement

	return p
}

func (s *EchoStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *EchoStatementContext) ECHO() antlr.TerminalNode {
	return s.GetToken(PuppetoParserECHO, 0)
}

func (s *EchoStatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *EchoStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EchoStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EchoStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PuppetoParserVisitor:
		return t.VisitEchoStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PuppetoParser) EchoStatement() (localctx IEchoStatementContext) {
	this := p
	_ = this

	localctx = NewEchoStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, PuppetoParserRULE_echoStatement)

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
		p.SetState(441)
		p.Match(PuppetoParserECHO)
	}
	{
		p.SetState(442)
		p.Expression()
	}

	return localctx
}

// IProgramBodyWithBreakContinueContext is an interface to support dynamic dispatch.
type IProgramBodyWithBreakContinueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ScopeWithBreakContinue() IScopeWithBreakContinueContext
	Statement() IStatementContext
	BreakStatement() IBreakStatementContext
	ContinueStatement() IContinueStatementContext
	Expression() IExpressionContext

	// IsProgramBodyWithBreakContinueContext differentiates from other interfaces.
	IsProgramBodyWithBreakContinueContext()
}

type ProgramBodyWithBreakContinueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramBodyWithBreakContinueContext() *ProgramBodyWithBreakContinueContext {
	var p = new(ProgramBodyWithBreakContinueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PuppetoParserRULE_programBodyWithBreakContinue
	return p
}

func (*ProgramBodyWithBreakContinueContext) IsProgramBodyWithBreakContinueContext() {}

func NewProgramBodyWithBreakContinueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramBodyWithBreakContinueContext {
	var p = new(ProgramBodyWithBreakContinueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PuppetoParserRULE_programBodyWithBreakContinue

	return p
}

func (s *ProgramBodyWithBreakContinueContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramBodyWithBreakContinueContext) ScopeWithBreakContinue() IScopeWithBreakContinueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IScopeWithBreakContinueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IScopeWithBreakContinueContext)
}

func (s *ProgramBodyWithBreakContinueContext) Statement() IStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ProgramBodyWithBreakContinueContext) BreakStatement() IBreakStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBreakStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBreakStatementContext)
}

func (s *ProgramBodyWithBreakContinueContext) ContinueStatement() IContinueStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IContinueStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IContinueStatementContext)
}

func (s *ProgramBodyWithBreakContinueContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ProgramBodyWithBreakContinueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramBodyWithBreakContinueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramBodyWithBreakContinueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PuppetoParserVisitor:
		return t.VisitProgramBodyWithBreakContinue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PuppetoParser) ProgramBodyWithBreakContinue() (localctx IProgramBodyWithBreakContinueContext) {
	this := p
	_ = this

	localctx = NewProgramBodyWithBreakContinueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, PuppetoParserRULE_programBodyWithBreakContinue)

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

	p.SetState(449)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 49, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(444)
			p.ScopeWithBreakContinue()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(445)
			p.Statement()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(446)
			p.BreakStatement()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(447)
			p.ContinueStatement()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(448)
			p.Expression()
		}

	}

	return localctx
}

// IPupCodeContext is an interface to support dynamic dispatch.
type IPupCodeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PUP_START_TAG() antlr.TerminalNode
	PUP_END_TAG() antlr.TerminalNode
	ProgramBodyList() IProgramBodyListContext

	// IsPupCodeContext differentiates from other interfaces.
	IsPupCodeContext()
}

type PupCodeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPupCodeContext() *PupCodeContext {
	var p = new(PupCodeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PuppetoParserRULE_pupCode
	return p
}

func (*PupCodeContext) IsPupCodeContext() {}

func NewPupCodeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PupCodeContext {
	var p = new(PupCodeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PuppetoParserRULE_pupCode

	return p
}

func (s *PupCodeContext) GetParser() antlr.Parser { return s.parser }

func (s *PupCodeContext) PUP_START_TAG() antlr.TerminalNode {
	return s.GetToken(PuppetoParserPUP_START_TAG, 0)
}

func (s *PupCodeContext) PUP_END_TAG() antlr.TerminalNode {
	return s.GetToken(PuppetoParserPUP_END_TAG, 0)
}

func (s *PupCodeContext) ProgramBodyList() IProgramBodyListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProgramBodyListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProgramBodyListContext)
}

func (s *PupCodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PupCodeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PupCodeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PuppetoParserVisitor:
		return t.VisitPupCode(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PuppetoParser) PupCode() (localctx IPupCodeContext) {
	this := p
	_ = this

	localctx = NewPupCodeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, PuppetoParserRULE_pupCode)

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
		p.SetState(451)
		p.Match(PuppetoParserPUP_START_TAG)
	}
	p.SetState(453)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 50, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(452)
			p.ProgramBodyList()
		}

	}
	{
		p.SetState(455)
		p.Match(PuppetoParserPUP_END_TAG)
	}

	return localctx
}

// IHtmlListContext is an interface to support dynamic dispatch.
type IHtmlListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	WS() antlr.TerminalNode
	LT() antlr.TerminalNode
	HTML() antlr.TerminalNode
	AND() antlr.TerminalNode
	BAND() antlr.TerminalNode
	BLS() antlr.TerminalNode
	BOOLEAN() antlr.TerminalNode
	BOR() antlr.TerminalNode
	BREAK() antlr.TerminalNode
	BRS() antlr.TerminalNode
	BXOR() antlr.TerminalNode
	CASE() antlr.TerminalNode
	CATCH() antlr.TerminalNode
	COLON() antlr.TerminalNode
	COMMA() antlr.TerminalNode
	CONTINUE() antlr.TerminalNode
	DEFAULT() antlr.TerminalNode
	DOT() antlr.TerminalNode
	DOUBLE_STAR() antlr.TerminalNode
	ECHO() antlr.TerminalNode
	ELSE() antlr.TerminalNode
	EQ() antlr.TerminalNode
	FLOAT() antlr.TerminalNode
	FN() antlr.TerminalNode
	FOR() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	IF() antlr.TerminalNode
	INTEGER() antlr.TerminalNode
	L_CURLY() antlr.TerminalNode
	L_PARENTHES() antlr.TerminalNode
	L_SQUARE() antlr.TerminalNode
	LET() antlr.TerminalNode
	LTOE() antlr.TerminalNode
	MINUS() antlr.TerminalNode
	GT() antlr.TerminalNode
	GTOE() antlr.TerminalNode
	NEQ() antlr.TerminalNode
	NIL() antlr.TerminalNode
	NOT() antlr.TerminalNode
	ONE_SIDE_OPERATOR() antlr.TerminalNode
	OR() antlr.TerminalNode
	PERCENT() antlr.TerminalNode
	PLUS() antlr.TerminalNode
	R_CURLY() antlr.TerminalNode
	R_PARENTHES() antlr.TerminalNode
	R_SQUARE() antlr.TerminalNode
	SEMICOLON() antlr.TerminalNode
	SLASH() antlr.TerminalNode
	STAR() antlr.TerminalNode
	STRING() antlr.TerminalNode
	SWITCH() antlr.TerminalNode
	TO() antlr.TerminalNode
	TRY() antlr.TerminalNode
	TWO_SIDES_OPERATOR() antlr.TerminalNode

	// IsHtmlListContext differentiates from other interfaces.
	IsHtmlListContext()
}

type HtmlListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHtmlListContext() *HtmlListContext {
	var p = new(HtmlListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PuppetoParserRULE_htmlList
	return p
}

func (*HtmlListContext) IsHtmlListContext() {}

func NewHtmlListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HtmlListContext {
	var p = new(HtmlListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PuppetoParserRULE_htmlList

	return p
}

func (s *HtmlListContext) GetParser() antlr.Parser { return s.parser }

func (s *HtmlListContext) WS() antlr.TerminalNode {
	return s.GetToken(PuppetoParserWS, 0)
}

func (s *HtmlListContext) LT() antlr.TerminalNode {
	return s.GetToken(PuppetoParserLT, 0)
}

func (s *HtmlListContext) HTML() antlr.TerminalNode {
	return s.GetToken(PuppetoParserHTML, 0)
}

func (s *HtmlListContext) AND() antlr.TerminalNode {
	return s.GetToken(PuppetoParserAND, 0)
}

func (s *HtmlListContext) BAND() antlr.TerminalNode {
	return s.GetToken(PuppetoParserBAND, 0)
}

func (s *HtmlListContext) BLS() antlr.TerminalNode {
	return s.GetToken(PuppetoParserBLS, 0)
}

func (s *HtmlListContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(PuppetoParserBOOLEAN, 0)
}

func (s *HtmlListContext) BOR() antlr.TerminalNode {
	return s.GetToken(PuppetoParserBOR, 0)
}

func (s *HtmlListContext) BREAK() antlr.TerminalNode {
	return s.GetToken(PuppetoParserBREAK, 0)
}

func (s *HtmlListContext) BRS() antlr.TerminalNode {
	return s.GetToken(PuppetoParserBRS, 0)
}

func (s *HtmlListContext) BXOR() antlr.TerminalNode {
	return s.GetToken(PuppetoParserBXOR, 0)
}

func (s *HtmlListContext) CASE() antlr.TerminalNode {
	return s.GetToken(PuppetoParserCASE, 0)
}

func (s *HtmlListContext) CATCH() antlr.TerminalNode {
	return s.GetToken(PuppetoParserCATCH, 0)
}

func (s *HtmlListContext) COLON() antlr.TerminalNode {
	return s.GetToken(PuppetoParserCOLON, 0)
}

func (s *HtmlListContext) COMMA() antlr.TerminalNode {
	return s.GetToken(PuppetoParserCOMMA, 0)
}

func (s *HtmlListContext) CONTINUE() antlr.TerminalNode {
	return s.GetToken(PuppetoParserCONTINUE, 0)
}

func (s *HtmlListContext) DEFAULT() antlr.TerminalNode {
	return s.GetToken(PuppetoParserDEFAULT, 0)
}

func (s *HtmlListContext) DOT() antlr.TerminalNode {
	return s.GetToken(PuppetoParserDOT, 0)
}

func (s *HtmlListContext) DOUBLE_STAR() antlr.TerminalNode {
	return s.GetToken(PuppetoParserDOUBLE_STAR, 0)
}

func (s *HtmlListContext) ECHO() antlr.TerminalNode {
	return s.GetToken(PuppetoParserECHO, 0)
}

func (s *HtmlListContext) ELSE() antlr.TerminalNode {
	return s.GetToken(PuppetoParserELSE, 0)
}

func (s *HtmlListContext) EQ() antlr.TerminalNode {
	return s.GetToken(PuppetoParserEQ, 0)
}

func (s *HtmlListContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(PuppetoParserFLOAT, 0)
}

func (s *HtmlListContext) FN() antlr.TerminalNode {
	return s.GetToken(PuppetoParserFN, 0)
}

func (s *HtmlListContext) FOR() antlr.TerminalNode {
	return s.GetToken(PuppetoParserFOR, 0)
}

func (s *HtmlListContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(PuppetoParserIDENTIFIER, 0)
}

func (s *HtmlListContext) IF() antlr.TerminalNode {
	return s.GetToken(PuppetoParserIF, 0)
}

func (s *HtmlListContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(PuppetoParserINTEGER, 0)
}

func (s *HtmlListContext) L_CURLY() antlr.TerminalNode {
	return s.GetToken(PuppetoParserL_CURLY, 0)
}

func (s *HtmlListContext) L_PARENTHES() antlr.TerminalNode {
	return s.GetToken(PuppetoParserL_PARENTHES, 0)
}

func (s *HtmlListContext) L_SQUARE() antlr.TerminalNode {
	return s.GetToken(PuppetoParserL_SQUARE, 0)
}

func (s *HtmlListContext) LET() antlr.TerminalNode {
	return s.GetToken(PuppetoParserLET, 0)
}

func (s *HtmlListContext) LTOE() antlr.TerminalNode {
	return s.GetToken(PuppetoParserLTOE, 0)
}

func (s *HtmlListContext) MINUS() antlr.TerminalNode {
	return s.GetToken(PuppetoParserMINUS, 0)
}

func (s *HtmlListContext) GT() antlr.TerminalNode {
	return s.GetToken(PuppetoParserGT, 0)
}

func (s *HtmlListContext) GTOE() antlr.TerminalNode {
	return s.GetToken(PuppetoParserGTOE, 0)
}

func (s *HtmlListContext) NEQ() antlr.TerminalNode {
	return s.GetToken(PuppetoParserNEQ, 0)
}

func (s *HtmlListContext) NIL() antlr.TerminalNode {
	return s.GetToken(PuppetoParserNIL, 0)
}

func (s *HtmlListContext) NOT() antlr.TerminalNode {
	return s.GetToken(PuppetoParserNOT, 0)
}

func (s *HtmlListContext) ONE_SIDE_OPERATOR() antlr.TerminalNode {
	return s.GetToken(PuppetoParserONE_SIDE_OPERATOR, 0)
}

func (s *HtmlListContext) OR() antlr.TerminalNode {
	return s.GetToken(PuppetoParserOR, 0)
}

func (s *HtmlListContext) PERCENT() antlr.TerminalNode {
	return s.GetToken(PuppetoParserPERCENT, 0)
}

func (s *HtmlListContext) PLUS() antlr.TerminalNode {
	return s.GetToken(PuppetoParserPLUS, 0)
}

func (s *HtmlListContext) R_CURLY() antlr.TerminalNode {
	return s.GetToken(PuppetoParserR_CURLY, 0)
}

func (s *HtmlListContext) R_PARENTHES() antlr.TerminalNode {
	return s.GetToken(PuppetoParserR_PARENTHES, 0)
}

func (s *HtmlListContext) R_SQUARE() antlr.TerminalNode {
	return s.GetToken(PuppetoParserR_SQUARE, 0)
}

func (s *HtmlListContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(PuppetoParserSEMICOLON, 0)
}

func (s *HtmlListContext) SLASH() antlr.TerminalNode {
	return s.GetToken(PuppetoParserSLASH, 0)
}

func (s *HtmlListContext) STAR() antlr.TerminalNode {
	return s.GetToken(PuppetoParserSTAR, 0)
}

func (s *HtmlListContext) STRING() antlr.TerminalNode {
	return s.GetToken(PuppetoParserSTRING, 0)
}

func (s *HtmlListContext) SWITCH() antlr.TerminalNode {
	return s.GetToken(PuppetoParserSWITCH, 0)
}

func (s *HtmlListContext) TO() antlr.TerminalNode {
	return s.GetToken(PuppetoParserTO, 0)
}

func (s *HtmlListContext) TRY() antlr.TerminalNode {
	return s.GetToken(PuppetoParserTRY, 0)
}

func (s *HtmlListContext) TWO_SIDES_OPERATOR() antlr.TerminalNode {
	return s.GetToken(PuppetoParserTWO_SIDES_OPERATOR, 0)
}

func (s *HtmlListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HtmlListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HtmlListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PuppetoParserVisitor:
		return t.VisitHtmlList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PuppetoParser) HtmlList() (localctx IHtmlListContext) {
	this := p
	_ = this

	localctx = NewHtmlListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, PuppetoParserRULE_htmlList)
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
		p.SetState(457)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&576460752303423384) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IHtmlContext is an interface to support dynamic dispatch.
type IHtmlContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllPupCode() []IPupCodeContext
	PupCode(i int) IPupCodeContext
	AllHtmlList() []IHtmlListContext
	HtmlList(i int) IHtmlListContext

	// IsHtmlContext differentiates from other interfaces.
	IsHtmlContext()
}

type HtmlContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHtmlContext() *HtmlContext {
	var p = new(HtmlContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PuppetoParserRULE_html
	return p
}

func (*HtmlContext) IsHtmlContext() {}

func NewHtmlContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HtmlContext {
	var p = new(HtmlContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PuppetoParserRULE_html

	return p
}

func (s *HtmlContext) GetParser() antlr.Parser { return s.parser }

func (s *HtmlContext) EOF() antlr.TerminalNode {
	return s.GetToken(PuppetoParserEOF, 0)
}

func (s *HtmlContext) AllPupCode() []IPupCodeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPupCodeContext); ok {
			len++
		}
	}

	tst := make([]IPupCodeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPupCodeContext); ok {
			tst[i] = t.(IPupCodeContext)
			i++
		}
	}

	return tst
}

func (s *HtmlContext) PupCode(i int) IPupCodeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPupCodeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPupCodeContext)
}

func (s *HtmlContext) AllHtmlList() []IHtmlListContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IHtmlListContext); ok {
			len++
		}
	}

	tst := make([]IHtmlListContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IHtmlListContext); ok {
			tst[i] = t.(IHtmlListContext)
			i++
		}
	}

	return tst
}

func (s *HtmlContext) HtmlList(i int) IHtmlListContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHtmlListContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IHtmlListContext)
}

func (s *HtmlContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HtmlContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HtmlContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PuppetoParserVisitor:
		return t.VisitHtml(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PuppetoParser) Html() (localctx IHtmlContext) {
	this := p
	_ = this

	localctx = NewHtmlContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, PuppetoParserRULE_html)
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
	p.SetState(461)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&576460752303423416) != 0) {
		p.SetState(461)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case PuppetoParserPUP_START_TAG:
			{
				p.SetState(459)
				p.PupCode()
			}

		case PuppetoParserLT, PuppetoParserGT, PuppetoParserTWO_SIDES_OPERATOR, PuppetoParserONE_SIDE_OPERATOR, PuppetoParserEQ, PuppetoParserNEQ, PuppetoParserLTOE, PuppetoParserGTOE, PuppetoParserFOR, PuppetoParserTO, PuppetoParserECHO, PuppetoParserL_CURLY, PuppetoParserR_CURLY, PuppetoParserCOMMA, PuppetoParserL_SQUARE, PuppetoParserR_SQUARE, PuppetoParserFN, PuppetoParserL_PARENTHES, PuppetoParserR_PARENTHES, PuppetoParserAND, PuppetoParserOR, PuppetoParserCOLON, PuppetoParserDOT, PuppetoParserPLUS, PuppetoParserMINUS, PuppetoParserDOUBLE_STAR, PuppetoParserPERCENT, PuppetoParserBAND, PuppetoParserBOR, PuppetoParserBXOR, PuppetoParserBLS, PuppetoParserBRS, PuppetoParserNOT, PuppetoParserIF, PuppetoParserELSE, PuppetoParserSWITCH, PuppetoParserCASE, PuppetoParserBREAK, PuppetoParserTRY, PuppetoParserCONTINUE, PuppetoParserCATCH, PuppetoParserDEFAULT, PuppetoParserLET, PuppetoParserSTAR, PuppetoParserSLASH, PuppetoParserFLOAT, PuppetoParserINTEGER, PuppetoParserNIL, PuppetoParserBOOLEAN, PuppetoParserIDENTIFIER, PuppetoParserSTRING, PuppetoParserWS, PuppetoParserSEMICOLON, PuppetoParserHTML:
			{
				p.SetState(460)
				p.HtmlList()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(463)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(465)
		p.Match(PuppetoParserEOF)
	}

	return localctx
}

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PUP_START_TAG() antlr.TerminalNode
	EOF() antlr.TerminalNode
	ProgramBodyList() IProgramBodyListContext
	PUP_END_TAG() antlr.TerminalNode

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PuppetoParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PuppetoParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) PUP_START_TAG() antlr.TerminalNode {
	return s.GetToken(PuppetoParserPUP_START_TAG, 0)
}

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(PuppetoParserEOF, 0)
}

func (s *ProgramContext) ProgramBodyList() IProgramBodyListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProgramBodyListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProgramBodyListContext)
}

func (s *ProgramContext) PUP_END_TAG() antlr.TerminalNode {
	return s.GetToken(PuppetoParserPUP_END_TAG, 0)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PuppetoParserVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PuppetoParser) Program() (localctx IProgramContext) {
	this := p
	_ = this

	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, PuppetoParserRULE_program)
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
		p.SetState(467)
		p.Match(PuppetoParserPUP_START_TAG)
	}
	p.SetState(469)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 53, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(468)
			p.ProgramBodyList()
		}

	}
	p.SetState(472)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PuppetoParserPUP_END_TAG {
		{
			p.SetState(471)
			p.Match(PuppetoParserPUP_END_TAG)
		}

	}
	{
		p.SetState(474)
		p.Match(PuppetoParserEOF)
	}

	return localctx
}
