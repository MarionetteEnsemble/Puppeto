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
		"", "'...'", "'return'", "'nan'", "'<'", "'>'", "'<?pup'", "'?>'", "'<?='",
		"'=?>'", "", "", "'=='", "'!='", "'<='", "'>='", "'for'", "'='", "'echo'",
		"'{'", "'}'", "','", "'['", "']'", "'fn'", "'('", "')'", "'&&'", "'||'",
		"':'", "'.'", "'+'", "'-'", "'**'", "'%'", "'&'", "'|'", "'^'", "'<<'",
		"'>>'", "'!'", "'if'", "'else'", "'switch'", "'case'", "'break'", "'try'",
		"'continue'", "'catch'", "'default'", "'let'", "'*'", "'/'", "", "",
		"'nil'",
	}
	staticData.symbolicNames = []string{
		"", "SPREAD", "RETURN", "NAN", "LT", "GT", "PUP_START_TAG", "PUP_END_TAG",
		"PUP_SHORT_START_TAG", "PUP_SHORT_END_TAG", "TWO_SIDES_OPERATOR", "ONE_SIDE_OPERATOR",
		"EQ", "NEQ", "LTOE", "GTOE", "FOR", "TO", "ECHO", "L_CURLY", "R_CURLY",
		"COMMA", "L_SQUARE", "R_SQUARE", "FN", "L_PARENTHES", "R_PARENTHES",
		"AND", "OR", "COLON", "DOT", "PLUS", "MINUS", "DOUBLE_STAR", "PERCENT",
		"BAND", "BOR", "BXOR", "BLS", "BRS", "NOT", "IF", "ELSE", "SWITCH",
		"CASE", "BREAK", "TRY", "CONTINUE", "CATCH", "DEFAULT", "LET", "STAR",
		"SLASH", "FLOAT", "INTEGER", "NIL", "BOOLEAN", "IDENTIFIER", "STRING",
		"WS", "SEMICOLON", "HTML",
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
		"returnStatement", "statementList", "statement", "scope", "scopeWithBreakContinue",
		"programBody", "programBodyList", "echoStatement", "programBodyWithBreakContinue",
		"pupShortCode", "pupCode", "htmlList", "html", "program",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 61, 503, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46, 2, 47,
		7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 1, 0, 1, 0, 1, 1, 1, 1,
		1, 1, 1, 1, 5, 1, 109, 8, 1, 10, 1, 12, 1, 112, 9, 1, 3, 1, 114, 8, 1,
		1, 1, 3, 1, 117, 8, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		3, 2, 127, 8, 2, 1, 2, 1, 2, 1, 2, 3, 2, 132, 8, 2, 1, 3, 1, 3, 1, 3, 1,
		3, 5, 3, 138, 8, 3, 10, 3, 12, 3, 141, 9, 3, 1, 3, 3, 3, 144, 8, 3, 3,
		3, 146, 8, 3, 1, 3, 1, 3, 1, 4, 1, 4, 3, 4, 152, 8, 4, 1, 4, 1, 4, 3, 4,
		156, 8, 4, 1, 4, 1, 4, 3, 4, 160, 8, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1,
		5, 5, 5, 168, 8, 5, 10, 5, 12, 5, 171, 9, 5, 1, 5, 3, 5, 174, 8, 5, 1,
		6, 1, 6, 1, 6, 3, 6, 179, 8, 6, 1, 6, 3, 6, 182, 8, 6, 1, 7, 1, 7, 1, 8,
		1, 8, 3, 8, 188, 8, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		3, 9, 198, 8, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 5, 11, 205, 8, 11,
		10, 11, 12, 11, 208, 9, 11, 1, 12, 1, 12, 1, 12, 5, 12, 213, 8, 12, 10,
		12, 12, 12, 216, 9, 12, 1, 13, 1, 13, 1, 13, 5, 13, 221, 8, 13, 10, 13,
		12, 13, 224, 9, 13, 1, 14, 1, 14, 1, 14, 5, 14, 229, 8, 14, 10, 14, 12,
		14, 232, 9, 14, 1, 15, 1, 15, 1, 15, 5, 15, 237, 8, 15, 10, 15, 12, 15,
		240, 9, 15, 1, 16, 1, 16, 1, 16, 5, 16, 245, 8, 16, 10, 16, 12, 16, 248,
		9, 16, 1, 17, 1, 17, 1, 17, 5, 17, 253, 8, 17, 10, 17, 12, 17, 256, 9,
		17, 1, 18, 1, 18, 1, 18, 5, 18, 261, 8, 18, 10, 18, 12, 18, 264, 9, 18,
		1, 19, 1, 19, 1, 19, 3, 19, 269, 8, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1,
		20, 1, 20, 1, 20, 3, 20, 278, 8, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22,
		1, 22, 1, 22, 1, 22, 1, 22, 5, 22, 289, 8, 22, 10, 22, 12, 22, 292, 9,
		22, 1, 22, 3, 22, 295, 8, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23,
		3, 23, 303, 8, 23, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 26, 1,
		26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 3, 27, 318, 8, 27, 1, 27, 1, 27,
		1, 27, 1, 27, 3, 27, 324, 8, 27, 5, 27, 326, 8, 27, 10, 27, 12, 27, 329,
		9, 27, 1, 27, 3, 27, 332, 8, 27, 1, 28, 1, 28, 1, 28, 1, 28, 5, 28, 338,
		8, 28, 10, 28, 12, 28, 341, 9, 28, 1, 29, 1, 29, 1, 29, 3, 29, 346, 8,
		29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 5, 30, 356,
		8, 30, 10, 30, 12, 30, 359, 9, 30, 1, 30, 1, 30, 3, 30, 363, 8, 30, 1,
		31, 1, 31, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 1, 32, 5, 32, 373, 8, 32,
		10, 32, 12, 32, 376, 9, 32, 1, 32, 1, 32, 1, 32, 5, 32, 381, 8, 32, 10,
		32, 12, 32, 384, 9, 32, 3, 32, 386, 8, 32, 1, 32, 1, 32, 1, 33, 1, 33,
		1, 33, 1, 33, 1, 33, 1, 34, 1, 34, 1, 35, 1, 35, 1, 36, 1, 36, 1, 36, 1,
		36, 3, 36, 403, 8, 36, 1, 36, 3, 36, 406, 8, 36, 1, 37, 1, 37, 3, 37, 410,
		8, 37, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 3,
		38, 421, 8, 38, 1, 39, 1, 39, 1, 39, 5, 39, 426, 8, 39, 10, 39, 12, 39,
		429, 9, 39, 1, 39, 4, 39, 432, 8, 39, 11, 39, 12, 39, 433, 3, 39, 436,
		8, 39, 1, 40, 1, 40, 3, 40, 440, 8, 40, 1, 40, 1, 40, 1, 41, 1, 41, 5,
		41, 446, 8, 41, 10, 41, 12, 41, 449, 9, 41, 1, 41, 1, 41, 1, 42, 1, 42,
		1, 42, 3, 42, 456, 8, 42, 1, 43, 4, 43, 459, 8, 43, 11, 43, 12, 43, 460,
		1, 44, 1, 44, 1, 44, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 3, 45, 471, 8,
		45, 1, 46, 1, 46, 1, 46, 1, 46, 1, 47, 1, 47, 3, 47, 479, 8, 47, 1, 47,
		1, 47, 1, 48, 1, 48, 1, 49, 1, 49, 1, 49, 4, 49, 488, 8, 49, 11, 49, 12,
		49, 489, 1, 49, 1, 49, 1, 50, 1, 50, 3, 50, 496, 8, 50, 1, 50, 3, 50, 499,
		8, 50, 1, 50, 1, 50, 1, 50, 0, 0, 51, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18,
		20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54,
		56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 88, 90,
		92, 94, 96, 98, 100, 0, 11, 2, 0, 3, 3, 53, 58, 1, 0, 27, 28, 1, 0, 12,
		13, 2, 0, 4, 5, 14, 15, 1, 0, 31, 32, 2, 0, 34, 34, 51, 52, 1, 0, 35, 37,
		1, 0, 38, 39, 2, 0, 32, 32, 40, 40, 2, 0, 10, 10, 17, 17, 2, 0, 4, 5, 10,
		61, 529, 0, 102, 1, 0, 0, 0, 2, 104, 1, 0, 0, 0, 4, 131, 1, 0, 0, 0, 6,
		133, 1, 0, 0, 0, 8, 149, 1, 0, 0, 0, 10, 164, 1, 0, 0, 0, 12, 175, 1, 0,
		0, 0, 14, 183, 1, 0, 0, 0, 16, 187, 1, 0, 0, 0, 18, 197, 1, 0, 0, 0, 20,
		199, 1, 0, 0, 0, 22, 201, 1, 0, 0, 0, 24, 209, 1, 0, 0, 0, 26, 217, 1,
		0, 0, 0, 28, 225, 1, 0, 0, 0, 30, 233, 1, 0, 0, 0, 32, 241, 1, 0, 0, 0,
		34, 249, 1, 0, 0, 0, 36, 257, 1, 0, 0, 0, 38, 268, 1, 0, 0, 0, 40, 270,
		1, 0, 0, 0, 42, 279, 1, 0, 0, 0, 44, 283, 1, 0, 0, 0, 46, 302, 1, 0, 0,
		0, 48, 304, 1, 0, 0, 0, 50, 307, 1, 0, 0, 0, 52, 310, 1, 0, 0, 0, 54, 317,
		1, 0, 0, 0, 56, 333, 1, 0, 0, 0, 58, 342, 1, 0, 0, 0, 60, 347, 1, 0, 0,
		0, 62, 364, 1, 0, 0, 0, 64, 368, 1, 0, 0, 0, 66, 389, 1, 0, 0, 0, 68, 394,
		1, 0, 0, 0, 70, 396, 1, 0, 0, 0, 72, 398, 1, 0, 0, 0, 74, 407, 1, 0, 0,
		0, 76, 420, 1, 0, 0, 0, 78, 435, 1, 0, 0, 0, 80, 437, 1, 0, 0, 0, 82, 443,
		1, 0, 0, 0, 84, 455, 1, 0, 0, 0, 86, 458, 1, 0, 0, 0, 88, 462, 1, 0, 0,
		0, 90, 470, 1, 0, 0, 0, 92, 472, 1, 0, 0, 0, 94, 476, 1, 0, 0, 0, 96, 482,
		1, 0, 0, 0, 98, 487, 1, 0, 0, 0, 100, 493, 1, 0, 0, 0, 102, 103, 7, 0,
		0, 0, 103, 1, 1, 0, 0, 0, 104, 113, 5, 19, 0, 0, 105, 110, 3, 4, 2, 0,
		106, 107, 5, 21, 0, 0, 107, 109, 3, 4, 2, 0, 108, 106, 1, 0, 0, 0, 109,
		112, 1, 0, 0, 0, 110, 108, 1, 0, 0, 0, 110, 111, 1, 0, 0, 0, 111, 114,
		1, 0, 0, 0, 112, 110, 1, 0, 0, 0, 113, 105, 1, 0, 0, 0, 113, 114, 1, 0,
		0, 0, 114, 116, 1, 0, 0, 0, 115, 117, 5, 21, 0, 0, 116, 115, 1, 0, 0, 0,
		116, 117, 1, 0, 0, 0, 117, 118, 1, 0, 0, 0, 118, 119, 5, 20, 0, 0, 119,
		3, 1, 0, 0, 0, 120, 121, 5, 22, 0, 0, 121, 122, 3, 20, 10, 0, 122, 123,
		5, 23, 0, 0, 123, 127, 1, 0, 0, 0, 124, 127, 5, 57, 0, 0, 125, 127, 5,
		58, 0, 0, 126, 120, 1, 0, 0, 0, 126, 124, 1, 0, 0, 0, 126, 125, 1, 0, 0,
		0, 127, 128, 1, 0, 0, 0, 128, 129, 5, 29, 0, 0, 129, 132, 3, 20, 10, 0,
		130, 132, 5, 57, 0, 0, 131, 126, 1, 0, 0, 0, 131, 130, 1, 0, 0, 0, 132,
		5, 1, 0, 0, 0, 133, 145, 5, 22, 0, 0, 134, 139, 3, 20, 10, 0, 135, 136,
		5, 21, 0, 0, 136, 138, 3, 20, 10, 0, 137, 135, 1, 0, 0, 0, 138, 141, 1,
		0, 0, 0, 139, 137, 1, 0, 0, 0, 139, 140, 1, 0, 0, 0, 140, 143, 1, 0, 0,
		0, 141, 139, 1, 0, 0, 0, 142, 144, 5, 21, 0, 0, 143, 142, 1, 0, 0, 0, 143,
		144, 1, 0, 0, 0, 144, 146, 1, 0, 0, 0, 145, 134, 1, 0, 0, 0, 145, 146,
		1, 0, 0, 0, 146, 147, 1, 0, 0, 0, 147, 148, 5, 23, 0, 0, 148, 7, 1, 0,
		0, 0, 149, 151, 5, 24, 0, 0, 150, 152, 5, 57, 0, 0, 151, 150, 1, 0, 0,
		0, 151, 152, 1, 0, 0, 0, 152, 153, 1, 0, 0, 0, 153, 155, 5, 25, 0, 0, 154,
		156, 3, 10, 5, 0, 155, 154, 1, 0, 0, 0, 155, 156, 1, 0, 0, 0, 156, 159,
		1, 0, 0, 0, 157, 158, 5, 1, 0, 0, 158, 160, 5, 57, 0, 0, 159, 157, 1, 0,
		0, 0, 159, 160, 1, 0, 0, 0, 160, 161, 1, 0, 0, 0, 161, 162, 5, 26, 0, 0,
		162, 163, 3, 16, 8, 0, 163, 9, 1, 0, 0, 0, 164, 169, 3, 12, 6, 0, 165,
		166, 5, 21, 0, 0, 166, 168, 3, 12, 6, 0, 167, 165, 1, 0, 0, 0, 168, 171,
		1, 0, 0, 0, 169, 167, 1, 0, 0, 0, 169, 170, 1, 0, 0, 0, 170, 173, 1, 0,
		0, 0, 171, 169, 1, 0, 0, 0, 172, 174, 5, 21, 0, 0, 173, 172, 1, 0, 0, 0,
		173, 174, 1, 0, 0, 0, 174, 11, 1, 0, 0, 0, 175, 181, 5, 57, 0, 0, 176,
		177, 5, 17, 0, 0, 177, 179, 3, 20, 10, 0, 178, 176, 1, 0, 0, 0, 178, 179,
		1, 0, 0, 0, 179, 182, 1, 0, 0, 0, 180, 182, 3, 14, 7, 0, 181, 178, 1, 0,
		0, 0, 181, 180, 1, 0, 0, 0, 182, 13, 1, 0, 0, 0, 183, 184, 1, 0, 0, 0,
		184, 15, 1, 0, 0, 0, 185, 188, 3, 20, 10, 0, 186, 188, 3, 80, 40, 0, 187,
		185, 1, 0, 0, 0, 187, 186, 1, 0, 0, 0, 188, 17, 1, 0, 0, 0, 189, 198, 3,
		0, 0, 0, 190, 198, 3, 2, 1, 0, 191, 198, 3, 6, 3, 0, 192, 198, 3, 8, 4,
		0, 193, 194, 5, 25, 0, 0, 194, 195, 3, 20, 10, 0, 195, 196, 5, 26, 0, 0,
		196, 198, 1, 0, 0, 0, 197, 189, 1, 0, 0, 0, 197, 190, 1, 0, 0, 0, 197,
		191, 1, 0, 0, 0, 197, 192, 1, 0, 0, 0, 197, 193, 1, 0, 0, 0, 198, 19, 1,
		0, 0, 0, 199, 200, 3, 22, 11, 0, 200, 21, 1, 0, 0, 0, 201, 206, 3, 24,
		12, 0, 202, 203, 7, 1, 0, 0, 203, 205, 3, 24, 12, 0, 204, 202, 1, 0, 0,
		0, 205, 208, 1, 0, 0, 0, 206, 204, 1, 0, 0, 0, 206, 207, 1, 0, 0, 0, 207,
		23, 1, 0, 0, 0, 208, 206, 1, 0, 0, 0, 209, 214, 3, 26, 13, 0, 210, 211,
		7, 2, 0, 0, 211, 213, 3, 26, 13, 0, 212, 210, 1, 0, 0, 0, 213, 216, 1,
		0, 0, 0, 214, 212, 1, 0, 0, 0, 214, 215, 1, 0, 0, 0, 215, 25, 1, 0, 0,
		0, 216, 214, 1, 0, 0, 0, 217, 222, 3, 28, 14, 0, 218, 219, 7, 3, 0, 0,
		219, 221, 3, 28, 14, 0, 220, 218, 1, 0, 0, 0, 221, 224, 1, 0, 0, 0, 222,
		220, 1, 0, 0, 0, 222, 223, 1, 0, 0, 0, 223, 27, 1, 0, 0, 0, 224, 222, 1,
		0, 0, 0, 225, 230, 3, 30, 15, 0, 226, 227, 7, 4, 0, 0, 227, 229, 3, 30,
		15, 0, 228, 226, 1, 0, 0, 0, 229, 232, 1, 0, 0, 0, 230, 228, 1, 0, 0, 0,
		230, 231, 1, 0, 0, 0, 231, 29, 1, 0, 0, 0, 232, 230, 1, 0, 0, 0, 233, 238,
		3, 32, 16, 0, 234, 235, 7, 5, 0, 0, 235, 237, 3, 32, 16, 0, 236, 234, 1,
		0, 0, 0, 237, 240, 1, 0, 0, 0, 238, 236, 1, 0, 0, 0, 238, 239, 1, 0, 0,
		0, 239, 31, 1, 0, 0, 0, 240, 238, 1, 0, 0, 0, 241, 246, 3, 34, 17, 0, 242,
		243, 5, 33, 0, 0, 243, 245, 3, 34, 17, 0, 244, 242, 1, 0, 0, 0, 245, 248,
		1, 0, 0, 0, 246, 244, 1, 0, 0, 0, 246, 247, 1, 0, 0, 0, 247, 33, 1, 0,
		0, 0, 248, 246, 1, 0, 0, 0, 249, 254, 3, 36, 18, 0, 250, 251, 7, 6, 0,
		0, 251, 253, 3, 36, 18, 0, 252, 250, 1, 0, 0, 0, 253, 256, 1, 0, 0, 0,
		254, 252, 1, 0, 0, 0, 254, 255, 1, 0, 0, 0, 255, 35, 1, 0, 0, 0, 256, 254,
		1, 0, 0, 0, 257, 262, 3, 38, 19, 0, 258, 259, 7, 7, 0, 0, 259, 261, 3,
		38, 19, 0, 260, 258, 1, 0, 0, 0, 261, 264, 1, 0, 0, 0, 262, 260, 1, 0,
		0, 0, 262, 263, 1, 0, 0, 0, 263, 37, 1, 0, 0, 0, 264, 262, 1, 0, 0, 0,
		265, 266, 7, 8, 0, 0, 266, 269, 3, 38, 19, 0, 267, 269, 3, 46, 23, 0, 268,
		265, 1, 0, 0, 0, 268, 267, 1, 0, 0, 0, 269, 39, 1, 0, 0, 0, 270, 277, 3,
		18, 9, 0, 271, 272, 5, 22, 0, 0, 272, 273, 3, 20, 10, 0, 273, 274, 5, 23,
		0, 0, 274, 278, 1, 0, 0, 0, 275, 276, 5, 30, 0, 0, 276, 278, 5, 57, 0,
		0, 277, 271, 1, 0, 0, 0, 277, 275, 1, 0, 0, 0, 278, 41, 1, 0, 0, 0, 279,
		280, 3, 40, 20, 0, 280, 281, 7, 9, 0, 0, 281, 282, 3, 20, 10, 0, 282, 43,
		1, 0, 0, 0, 283, 284, 3, 18, 9, 0, 284, 285, 5, 25, 0, 0, 285, 290, 3,
		20, 10, 0, 286, 287, 5, 21, 0, 0, 287, 289, 3, 20, 10, 0, 288, 286, 1,
		0, 0, 0, 289, 292, 1, 0, 0, 0, 290, 288, 1, 0, 0, 0, 290, 291, 1, 0, 0,
		0, 291, 294, 1, 0, 0, 0, 292, 290, 1, 0, 0, 0, 293, 295, 5, 21, 0, 0, 294,
		293, 1, 0, 0, 0, 294, 295, 1, 0, 0, 0, 295, 296, 1, 0, 0, 0, 296, 297,
		5, 26, 0, 0, 297, 45, 1, 0, 0, 0, 298, 303, 3, 40, 20, 0, 299, 303, 3,
		42, 21, 0, 300, 303, 3, 44, 22, 0, 301, 303, 3, 18, 9, 0, 302, 298, 1,
		0, 0, 0, 302, 299, 1, 0, 0, 0, 302, 300, 1, 0, 0, 0, 302, 301, 1, 0, 0,
		0, 303, 47, 1, 0, 0, 0, 304, 305, 5, 11, 0, 0, 305, 306, 5, 57, 0, 0, 306,
		49, 1, 0, 0, 0, 307, 308, 5, 57, 0, 0, 308, 309, 5, 11, 0, 0, 309, 51,
		1, 0, 0, 0, 310, 311, 5, 57, 0, 0, 311, 312, 7, 9, 0, 0, 312, 313, 3, 20,
		10, 0, 313, 53, 1, 0, 0, 0, 314, 318, 3, 52, 26, 0, 315, 318, 3, 48, 24,
		0, 316, 318, 3, 50, 25, 0, 317, 314, 1, 0, 0, 0, 317, 315, 1, 0, 0, 0,
		317, 316, 1, 0, 0, 0, 318, 327, 1, 0, 0, 0, 319, 323, 5, 21, 0, 0, 320,
		324, 3, 52, 26, 0, 321, 324, 3, 48, 24, 0, 322, 324, 3, 50, 25, 0, 323,
		320, 1, 0, 0, 0, 323, 321, 1, 0, 0, 0, 323, 322, 1, 0, 0, 0, 324, 326,
		1, 0, 0, 0, 325, 319, 1, 0, 0, 0, 326, 329, 1, 0, 0, 0, 327, 325, 1, 0,
		0, 0, 327, 328, 1, 0, 0, 0, 328, 331, 1, 0, 0, 0, 329, 327, 1, 0, 0, 0,
		330, 332, 5, 21, 0, 0, 331, 330, 1, 0, 0, 0, 331, 332, 1, 0, 0, 0, 332,
		55, 1, 0, 0, 0, 333, 334, 5, 50, 0, 0, 334, 339, 3, 58, 29, 0, 335, 336,
		5, 21, 0, 0, 336, 338, 3, 58, 29, 0, 337, 335, 1, 0, 0, 0, 338, 341, 1,
		0, 0, 0, 339, 337, 1, 0, 0, 0, 339, 340, 1, 0, 0, 0, 340, 57, 1, 0, 0,
		0, 341, 339, 1, 0, 0, 0, 342, 345, 5, 57, 0, 0, 343, 344, 5, 17, 0, 0,
		344, 346, 3, 20, 10, 0, 345, 343, 1, 0, 0, 0, 345, 346, 1, 0, 0, 0, 346,
		59, 1, 0, 0, 0, 347, 348, 5, 41, 0, 0, 348, 349, 3, 20, 10, 0, 349, 357,
		3, 16, 8, 0, 350, 351, 5, 42, 0, 0, 351, 352, 5, 41, 0, 0, 352, 353, 3,
		20, 10, 0, 353, 354, 3, 16, 8, 0, 354, 356, 1, 0, 0, 0, 355, 350, 1, 0,
		0, 0, 356, 359, 1, 0, 0, 0, 357, 355, 1, 0, 0, 0, 357, 358, 1, 0, 0, 0,
		358, 362, 1, 0, 0, 0, 359, 357, 1, 0, 0, 0, 360, 361, 5, 42, 0, 0, 361,
		363, 3, 16, 8, 0, 362, 360, 1, 0, 0, 0, 362, 363, 1, 0, 0, 0, 363, 61,
		1, 0, 0, 0, 364, 365, 5, 16, 0, 0, 365, 366, 3, 20, 10, 0, 366, 367, 3,
		90, 45, 0, 367, 63, 1, 0, 0, 0, 368, 369, 5, 43, 0, 0, 369, 370, 3, 20,
		10, 0, 370, 374, 5, 19, 0, 0, 371, 373, 3, 66, 33, 0, 372, 371, 1, 0, 0,
		0, 373, 376, 1, 0, 0, 0, 374, 372, 1, 0, 0, 0, 374, 375, 1, 0, 0, 0, 375,
		385, 1, 0, 0, 0, 376, 374, 1, 0, 0, 0, 377, 378, 5, 49, 0, 0, 378, 382,
		5, 29, 0, 0, 379, 381, 3, 90, 45, 0, 380, 379, 1, 0, 0, 0, 381, 384, 1,
		0, 0, 0, 382, 380, 1, 0, 0, 0, 382, 383, 1, 0, 0, 0, 383, 386, 1, 0, 0,
		0, 384, 382, 1, 0, 0, 0, 385, 377, 1, 0, 0, 0, 385, 386, 1, 0, 0, 0, 386,
		387, 1, 0, 0, 0, 387, 388, 5, 20, 0, 0, 388, 65, 1, 0, 0, 0, 389, 390,
		5, 44, 0, 0, 390, 391, 3, 20, 10, 0, 391, 392, 5, 29, 0, 0, 392, 393, 3,
		90, 45, 0, 393, 67, 1, 0, 0, 0, 394, 395, 5, 45, 0, 0, 395, 69, 1, 0, 0,
		0, 396, 397, 5, 47, 0, 0, 397, 71, 1, 0, 0, 0, 398, 399, 5, 46, 0, 0, 399,
		405, 3, 84, 42, 0, 400, 402, 5, 48, 0, 0, 401, 403, 5, 57, 0, 0, 402, 401,
		1, 0, 0, 0, 402, 403, 1, 0, 0, 0, 403, 404, 1, 0, 0, 0, 404, 406, 3, 84,
		42, 0, 405, 400, 1, 0, 0, 0, 405, 406, 1, 0, 0, 0, 406, 73, 1, 0, 0, 0,
		407, 409, 5, 2, 0, 0, 408, 410, 3, 20, 10, 0, 409, 408, 1, 0, 0, 0, 409,
		410, 1, 0, 0, 0, 410, 75, 1, 0, 0, 0, 411, 421, 3, 56, 28, 0, 412, 421,
		3, 8, 4, 0, 413, 421, 3, 60, 30, 0, 414, 421, 3, 62, 31, 0, 415, 421, 3,
		64, 32, 0, 416, 421, 3, 72, 36, 0, 417, 421, 3, 88, 44, 0, 418, 421, 3,
		54, 27, 0, 419, 421, 3, 74, 37, 0, 420, 411, 1, 0, 0, 0, 420, 412, 1, 0,
		0, 0, 420, 413, 1, 0, 0, 0, 420, 414, 1, 0, 0, 0, 420, 415, 1, 0, 0, 0,
		420, 416, 1, 0, 0, 0, 420, 417, 1, 0, 0, 0, 420, 418, 1, 0, 0, 0, 420,
		419, 1, 0, 0, 0, 421, 77, 1, 0, 0, 0, 422, 436, 3, 76, 38, 0, 423, 427,
		5, 7, 0, 0, 424, 426, 3, 96, 48, 0, 425, 424, 1, 0, 0, 0, 426, 429, 1,
		0, 0, 0, 427, 425, 1, 0, 0, 0, 427, 428, 1, 0, 0, 0, 428, 430, 1, 0, 0,
		0, 429, 427, 1, 0, 0, 0, 430, 432, 5, 6, 0, 0, 431, 423, 1, 0, 0, 0, 432,
		433, 1, 0, 0, 0, 433, 431, 1, 0, 0, 0, 433, 434, 1, 0, 0, 0, 434, 436,
		1, 0, 0, 0, 435, 422, 1, 0, 0, 0, 435, 431, 1, 0, 0, 0, 436, 79, 1, 0,
		0, 0, 437, 439, 5, 19, 0, 0, 438, 440, 3, 86, 43, 0, 439, 438, 1, 0, 0,
		0, 439, 440, 1, 0, 0, 0, 440, 441, 1, 0, 0, 0, 441, 442, 5, 20, 0, 0, 442,
		81, 1, 0, 0, 0, 443, 447, 5, 19, 0, 0, 444, 446, 3, 90, 45, 0, 445, 444,
		1, 0, 0, 0, 446, 449, 1, 0, 0, 0, 447, 445, 1, 0, 0, 0, 447, 448, 1, 0,
		0, 0, 448, 450, 1, 0, 0, 0, 449, 447, 1, 0, 0, 0, 450, 451, 5, 20, 0, 0,
		451, 83, 1, 0, 0, 0, 452, 456, 3, 80, 40, 0, 453, 456, 3, 78, 39, 0, 454,
		456, 3, 20, 10, 0, 455, 452, 1, 0, 0, 0, 455, 453, 1, 0, 0, 0, 455, 454,
		1, 0, 0, 0, 456, 85, 1, 0, 0, 0, 457, 459, 3, 84, 42, 0, 458, 457, 1, 0,
		0, 0, 459, 460, 1, 0, 0, 0, 460, 458, 1, 0, 0, 0, 460, 461, 1, 0, 0, 0,
		461, 87, 1, 0, 0, 0, 462, 463, 5, 18, 0, 0, 463, 464, 3, 20, 10, 0, 464,
		89, 1, 0, 0, 0, 465, 471, 3, 82, 41, 0, 466, 471, 3, 78, 39, 0, 467, 471,
		3, 68, 34, 0, 468, 471, 3, 70, 35, 0, 469, 471, 3, 20, 10, 0, 470, 465,
		1, 0, 0, 0, 470, 466, 1, 0, 0, 0, 470, 467, 1, 0, 0, 0, 470, 468, 1, 0,
		0, 0, 470, 469, 1, 0, 0, 0, 471, 91, 1, 0, 0, 0, 472, 473, 5, 8, 0, 0,
		473, 474, 3, 20, 10, 0, 474, 475, 5, 9, 0, 0, 475, 93, 1, 0, 0, 0, 476,
		478, 5, 6, 0, 0, 477, 479, 3, 86, 43, 0, 478, 477, 1, 0, 0, 0, 478, 479,
		1, 0, 0, 0, 479, 480, 1, 0, 0, 0, 480, 481, 5, 7, 0, 0, 481, 95, 1, 0,
		0, 0, 482, 483, 7, 10, 0, 0, 483, 97, 1, 0, 0, 0, 484, 488, 3, 94, 47,
		0, 485, 488, 3, 92, 46, 0, 486, 488, 3, 96, 48, 0, 487, 484, 1, 0, 0, 0,
		487, 485, 1, 0, 0, 0, 487, 486, 1, 0, 0, 0, 488, 489, 1, 0, 0, 0, 489,
		487, 1, 0, 0, 0, 489, 490, 1, 0, 0, 0, 490, 491, 1, 0, 0, 0, 491, 492,
		5, 0, 0, 1, 492, 99, 1, 0, 0, 0, 493, 495, 5, 6, 0, 0, 494, 496, 3, 86,
		43, 0, 495, 494, 1, 0, 0, 0, 495, 496, 1, 0, 0, 0, 496, 498, 1, 0, 0, 0,
		497, 499, 5, 7, 0, 0, 498, 497, 1, 0, 0, 0, 498, 499, 1, 0, 0, 0, 499,
		500, 1, 0, 0, 0, 500, 501, 5, 0, 0, 1, 501, 101, 1, 0, 0, 0, 58, 110, 113,
		116, 126, 131, 139, 143, 145, 151, 155, 159, 169, 173, 178, 181, 187, 197,
		206, 214, 222, 230, 238, 246, 254, 262, 268, 277, 290, 294, 302, 317, 323,
		327, 331, 339, 345, 357, 362, 374, 382, 385, 402, 405, 409, 420, 427, 433,
		435, 439, 447, 455, 460, 470, 478, 487, 489, 495, 498,
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
	PuppetoParserEOF                 = antlr.TokenEOF
	PuppetoParserSPREAD              = 1
	PuppetoParserRETURN              = 2
	PuppetoParserNAN                 = 3
	PuppetoParserLT                  = 4
	PuppetoParserGT                  = 5
	PuppetoParserPUP_START_TAG       = 6
	PuppetoParserPUP_END_TAG         = 7
	PuppetoParserPUP_SHORT_START_TAG = 8
	PuppetoParserPUP_SHORT_END_TAG   = 9
	PuppetoParserTWO_SIDES_OPERATOR  = 10
	PuppetoParserONE_SIDE_OPERATOR   = 11
	PuppetoParserEQ                  = 12
	PuppetoParserNEQ                 = 13
	PuppetoParserLTOE                = 14
	PuppetoParserGTOE                = 15
	PuppetoParserFOR                 = 16
	PuppetoParserTO                  = 17
	PuppetoParserECHO                = 18
	PuppetoParserL_CURLY             = 19
	PuppetoParserR_CURLY             = 20
	PuppetoParserCOMMA               = 21
	PuppetoParserL_SQUARE            = 22
	PuppetoParserR_SQUARE            = 23
	PuppetoParserFN                  = 24
	PuppetoParserL_PARENTHES         = 25
	PuppetoParserR_PARENTHES         = 26
	PuppetoParserAND                 = 27
	PuppetoParserOR                  = 28
	PuppetoParserCOLON               = 29
	PuppetoParserDOT                 = 30
	PuppetoParserPLUS                = 31
	PuppetoParserMINUS               = 32
	PuppetoParserDOUBLE_STAR         = 33
	PuppetoParserPERCENT             = 34
	PuppetoParserBAND                = 35
	PuppetoParserBOR                 = 36
	PuppetoParserBXOR                = 37
	PuppetoParserBLS                 = 38
	PuppetoParserBRS                 = 39
	PuppetoParserNOT                 = 40
	PuppetoParserIF                  = 41
	PuppetoParserELSE                = 42
	PuppetoParserSWITCH              = 43
	PuppetoParserCASE                = 44
	PuppetoParserBREAK               = 45
	PuppetoParserTRY                 = 46
	PuppetoParserCONTINUE            = 47
	PuppetoParserCATCH               = 48
	PuppetoParserDEFAULT             = 49
	PuppetoParserLET                 = 50
	PuppetoParserSTAR                = 51
	PuppetoParserSLASH               = 52
	PuppetoParserFLOAT               = 53
	PuppetoParserINTEGER             = 54
	PuppetoParserNIL                 = 55
	PuppetoParserBOOLEAN             = 56
	PuppetoParserIDENTIFIER          = 57
	PuppetoParserSTRING              = 58
	PuppetoParserWS                  = 59
	PuppetoParserSEMICOLON           = 60
	PuppetoParserHTML                = 61
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
	PuppetoParserRULE_returnStatement              = 37
	PuppetoParserRULE_statementList                = 38
	PuppetoParserRULE_statement                    = 39
	PuppetoParserRULE_scope                        = 40
	PuppetoParserRULE_scopeWithBreakContinue       = 41
	PuppetoParserRULE_programBody                  = 42
	PuppetoParserRULE_programBodyList              = 43
	PuppetoParserRULE_echoStatement                = 44
	PuppetoParserRULE_programBodyWithBreakContinue = 45
	PuppetoParserRULE_pupShortCode                 = 46
	PuppetoParserRULE_pupCode                      = 47
	PuppetoParserRULE_htmlList                     = 48
	PuppetoParserRULE_html                         = 49
	PuppetoParserRULE_program                      = 50
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
		p.SetState(102)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&567453553048682504) != 0) {
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
		p.SetState(104)
		p.Match(PuppetoParserL_CURLY)
	}
	p.SetState(113)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&432345564231761920) != 0 {
		{
			p.SetState(105)
			p.Pair()
		}
		p.SetState(110)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(106)
					p.Match(PuppetoParserCOMMA)
				}
				{
					p.SetState(107)
					p.Pair()
				}

			}
			p.SetState(112)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())
		}

	}
	p.SetState(116)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PuppetoParserCOMMA {
		{
			p.SetState(115)
			p.Match(PuppetoParserCOMMA)
		}

	}
	{
		p.SetState(118)
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

	p.SetState(131)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(126)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case PuppetoParserL_SQUARE:
			{
				p.SetState(120)
				p.Match(PuppetoParserL_SQUARE)
			}
			{
				p.SetState(121)
				p.Expression()
			}
			{
				p.SetState(122)
				p.Match(PuppetoParserR_SQUARE)
			}

		case PuppetoParserIDENTIFIER:
			{
				p.SetState(124)
				p.Match(PuppetoParserIDENTIFIER)
			}

		case PuppetoParserSTRING:
			{
				p.SetState(125)
				p.Match(PuppetoParserSTRING)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}
		{
			p.SetState(128)
			p.Match(PuppetoParserCOLON)
		}
		{
			p.SetState(129)
			p.Expression()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(130)
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
		p.SetState(133)
		p.Match(PuppetoParserL_SQUARE)
	}
	p.SetState(145)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&567454656910327816) != 0 {
		{
			p.SetState(134)
			p.Expression()
		}
		p.SetState(139)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(135)
					p.Match(PuppetoParserCOMMA)
				}
				{
					p.SetState(136)
					p.Expression()
				}

			}
			p.SetState(141)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())
		}
		p.SetState(143)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == PuppetoParserCOMMA {
			{
				p.SetState(142)
				p.Match(PuppetoParserCOMMA)
			}

		}

	}
	{
		p.SetState(147)
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

	// GetRest returns the rest token.
	GetRest() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// SetRest sets the rest token.
	SetRest(antlr.Token)

	// Getter signatures
	FN() antlr.TerminalNode
	L_PARENTHES() antlr.TerminalNode
	R_PARENTHES() antlr.TerminalNode
	ScopeBody() IScopeBodyContext
	ArgumentList() IArgumentListContext
	SPREAD() antlr.TerminalNode
	AllIDENTIFIER() []antlr.TerminalNode
	IDENTIFIER(i int) antlr.TerminalNode

	// IsFunctionDefinitionContext differentiates from other interfaces.
	IsFunctionDefinitionContext()
}

type FunctionDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
	rest   antlr.Token
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

func (s *FunctionDefinitionContext) GetRest() antlr.Token { return s.rest }

func (s *FunctionDefinitionContext) SetName(v antlr.Token) { s.name = v }

func (s *FunctionDefinitionContext) SetRest(v antlr.Token) { s.rest = v }

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

func (s *FunctionDefinitionContext) SPREAD() antlr.TerminalNode {
	return s.GetToken(PuppetoParserSPREAD, 0)
}

func (s *FunctionDefinitionContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(PuppetoParserIDENTIFIER)
}

func (s *FunctionDefinitionContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(PuppetoParserIDENTIFIER, i)
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
		p.SetState(149)
		p.Match(PuppetoParserFN)
	}
	p.SetState(151)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PuppetoParserIDENTIFIER {
		{
			p.SetState(150)

			var _m = p.Match(PuppetoParserIDENTIFIER)

			localctx.(*FunctionDefinitionContext).name = _m
		}

	}
	{
		p.SetState(153)
		p.Match(PuppetoParserL_PARENTHES)
	}
	p.SetState(155)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PuppetoParserIDENTIFIER {
		{
			p.SetState(154)
			p.ArgumentList()
		}

	}
	p.SetState(159)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PuppetoParserSPREAD {
		{
			p.SetState(157)
			p.Match(PuppetoParserSPREAD)
		}
		{
			p.SetState(158)

			var _m = p.Match(PuppetoParserIDENTIFIER)

			localctx.(*FunctionDefinitionContext).rest = _m
		}

	}
	{
		p.SetState(161)
		p.Match(PuppetoParserR_PARENTHES)
	}
	{
		p.SetState(162)
		p.ScopeBody()
	}

	return localctx
}

// IArgumentListContext is an interface to support dynamic dispatch.
type IArgumentListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllArgument() []IArgumentContext
	Argument(i int) IArgumentContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsArgumentListContext differentiates from other interfaces.
	IsArgumentListContext()
}

type ArgumentListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
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
		p.SetState(164)
		p.Argument()
	}
	p.SetState(169)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(165)
				p.Match(PuppetoParserCOMMA)
			}
			{
				p.SetState(166)
				p.Argument()
			}

		}
		p.SetState(171)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())
	}
	p.SetState(173)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PuppetoParserCOMMA {
		{
			p.SetState(172)
			p.Match(PuppetoParserCOMMA)
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
		p.SetState(175)
		p.Match(PuppetoParserIDENTIFIER)
	}
	p.SetState(181)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		p.SetState(178)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == PuppetoParserTO {
			{
				p.SetState(176)
				p.Match(PuppetoParserTO)
			}
			{
				p.SetState(177)
				p.Expression()
			}

		}

	case 2:
		{
			p.SetState(180)
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

	p.SetState(187)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(185)
			p.Expression()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(186)
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

	p.SetState(197)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PuppetoParserNAN, PuppetoParserFLOAT, PuppetoParserINTEGER, PuppetoParserNIL, PuppetoParserBOOLEAN, PuppetoParserIDENTIFIER, PuppetoParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(189)
			p.LiteralValue()
		}

	case PuppetoParserL_CURLY:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(190)
			p.Object()
		}

	case PuppetoParserL_SQUARE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(191)
			p.Array()
		}

	case PuppetoParserFN:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(192)
			p.FunctionDefinition()
		}

	case PuppetoParserL_PARENTHES:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(193)
			p.Match(PuppetoParserL_PARENTHES)
		}
		{
			p.SetState(194)
			p.Expression()
		}
		{
			p.SetState(195)
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
		p.SetState(199)
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
		p.SetState(201)
		p.EqualityExpression()
	}
	p.SetState(206)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(202)
				_la = p.GetTokenStream().LA(1)

				if !(_la == PuppetoParserAND || _la == PuppetoParserOR) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(203)
				p.EqualityExpression()
			}

		}
		p.SetState(208)
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
		p.SetState(209)
		p.ComparisonExpression()
	}
	p.SetState(214)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(210)
				_la = p.GetTokenStream().LA(1)

				if !(_la == PuppetoParserEQ || _la == PuppetoParserNEQ) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(211)
				p.ComparisonExpression()
			}

		}
		p.SetState(216)
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
		p.SetState(217)
		p.AdditiveExpression()
	}
	p.SetState(222)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(218)
				_la = p.GetTokenStream().LA(1)

				if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&49200) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(219)
				p.AdditiveExpression()
			}

		}
		p.SetState(224)
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
		p.SetState(225)
		p.MultiplicativeExpression()
	}
	p.SetState(230)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(226)
				_la = p.GetTokenStream().LA(1)

				if !(_la == PuppetoParserPLUS || _la == PuppetoParserMINUS) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(227)
				p.MultiplicativeExpression()
			}

		}
		p.SetState(232)
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
		p.SetState(233)
		p.PowerExpression()
	}
	p.SetState(238)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(234)
				_la = p.GetTokenStream().LA(1)

				if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&6755416620924928) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(235)
				p.PowerExpression()
			}

		}
		p.SetState(240)
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
		p.SetState(241)
		p.BitwiseExpression()
	}
	p.SetState(246)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(242)
				p.Match(PuppetoParserDOUBLE_STAR)
			}
			{
				p.SetState(243)
				p.BitwiseExpression()
			}

		}
		p.SetState(248)
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
		p.SetState(249)
		p.ShiftExpression()
	}
	p.SetState(254)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(250)
				_la = p.GetTokenStream().LA(1)

				if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&240518168576) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(251)
				p.ShiftExpression()
			}

		}
		p.SetState(256)
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
		p.SetState(257)
		p.UnaryExpression()
	}
	p.SetState(262)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(258)
				_la = p.GetTokenStream().LA(1)

				if !(_la == PuppetoParserBLS || _la == PuppetoParserBRS) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(259)
				p.UnaryExpression()
			}

		}
		p.SetState(264)
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

	p.SetState(268)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PuppetoParserMINUS, PuppetoParserNOT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(265)
			_la = p.GetTokenStream().LA(1)

			if !(_la == PuppetoParserMINUS || _la == PuppetoParserNOT) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(266)
			p.UnaryExpression()
		}

	case PuppetoParserNAN, PuppetoParserL_CURLY, PuppetoParserL_SQUARE, PuppetoParserFN, PuppetoParserL_PARENTHES, PuppetoParserFLOAT, PuppetoParserINTEGER, PuppetoParserNIL, PuppetoParserBOOLEAN, PuppetoParserIDENTIFIER, PuppetoParserSTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(267)
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
		p.SetState(270)
		p.Value()
	}
	p.SetState(277)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PuppetoParserL_SQUARE:
		{
			p.SetState(271)
			p.Match(PuppetoParserL_SQUARE)
		}
		{
			p.SetState(272)
			p.Expression()
		}
		{
			p.SetState(273)
			p.Match(PuppetoParserR_SQUARE)
		}

	case PuppetoParserDOT:
		{
			p.SetState(275)
			p.Match(PuppetoParserDOT)
		}
		{
			p.SetState(276)
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
		p.SetState(279)
		p.GetProperty()
	}
	{
		p.SetState(280)
		_la = p.GetTokenStream().LA(1)

		if !(_la == PuppetoParserTWO_SIDES_OPERATOR || _la == PuppetoParserTO) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(281)
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
		p.SetState(283)
		p.Value()
	}
	{
		p.SetState(284)
		p.Match(PuppetoParserL_PARENTHES)
	}
	{
		p.SetState(285)
		p.Expression()
	}
	p.SetState(290)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(286)
				p.Match(PuppetoParserCOMMA)
			}
			{
				p.SetState(287)
				p.Expression()
			}

		}
		p.SetState(292)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext())
	}
	p.SetState(294)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PuppetoParserCOMMA {
		{
			p.SetState(293)
			p.Match(PuppetoParserCOMMA)
		}

	}
	{
		p.SetState(296)
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

	p.SetState(302)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(298)
			p.GetProperty()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(299)
			p.SetProperty()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(300)
			p.CallExpression()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(301)
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
		p.SetState(304)
		p.Match(PuppetoParserONE_SIDE_OPERATOR)
	}
	{
		p.SetState(305)
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
		p.SetState(307)
		p.Match(PuppetoParserIDENTIFIER)
	}
	{
		p.SetState(308)
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
		p.SetState(310)
		p.Match(PuppetoParserIDENTIFIER)
	}
	{
		p.SetState(311)
		_la = p.GetTokenStream().LA(1)

		if !(_la == PuppetoParserTWO_SIDES_OPERATOR || _la == PuppetoParserTO) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(312)
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
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

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

func (s *VariableAssignationContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(PuppetoParserCOMMA)
}

func (s *VariableAssignationContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(PuppetoParserCOMMA, i)
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
	p.SetState(317)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(314)
			p.MidVariableAssignation()
		}

	case 2:
		{
			p.SetState(315)
			p.LhsVariableAssignation()
		}

	case 3:
		{
			p.SetState(316)
			p.RhsVariableAssignation()
		}

	}
	p.SetState(327)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(319)
				p.Match(PuppetoParserCOMMA)
			}
			p.SetState(323)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext()) {
			case 1:
				{
					p.SetState(320)
					p.MidVariableAssignation()
				}

			case 2:
				{
					p.SetState(321)
					p.LhsVariableAssignation()
				}

			case 3:
				{
					p.SetState(322)
					p.RhsVariableAssignation()
				}

			}

		}
		p.SetState(329)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext())
	}
	p.SetState(331)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PuppetoParserCOMMA {
		{
			p.SetState(330)
			p.Match(PuppetoParserCOMMA)
		}

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
		p.SetState(333)
		p.Match(PuppetoParserLET)
	}
	{
		p.SetState(334)
		p.VariableDefinitionBody()
	}
	p.SetState(339)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == PuppetoParserCOMMA {
		{
			p.SetState(335)
			p.Match(PuppetoParserCOMMA)
		}
		{
			p.SetState(336)
			p.VariableDefinitionBody()
		}

		p.SetState(341)
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
		p.SetState(342)
		p.Match(PuppetoParserIDENTIFIER)
	}
	p.SetState(345)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PuppetoParserTO {
		{
			p.SetState(343)
			p.Match(PuppetoParserTO)
		}
		{
			p.SetState(344)
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
		p.SetState(347)
		p.Match(PuppetoParserIF)
	}
	{
		p.SetState(348)
		p.Expression()
	}
	{
		p.SetState(349)
		p.ScopeBody()
	}
	p.SetState(357)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 36, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(350)
				p.Match(PuppetoParserELSE)
			}
			{
				p.SetState(351)
				p.Match(PuppetoParserIF)
			}
			{
				p.SetState(352)
				p.Expression()
			}
			{
				p.SetState(353)
				p.ScopeBody()
			}

		}
		p.SetState(359)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 36, p.GetParserRuleContext())
	}
	p.SetState(362)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PuppetoParserELSE {
		{
			p.SetState(360)
			p.Match(PuppetoParserELSE)
		}
		{
			p.SetState(361)
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
		p.SetState(364)
		p.Match(PuppetoParserFOR)
	}
	{
		p.SetState(365)
		p.Expression()
	}
	{
		p.SetState(366)
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
		p.SetState(368)
		p.Match(PuppetoParserSWITCH)
	}
	{
		p.SetState(369)
		p.Expression()
	}
	{
		p.SetState(370)
		p.Match(PuppetoParserL_CURLY)
	}
	p.SetState(374)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == PuppetoParserCASE {
		{
			p.SetState(371)
			p.SwitchCase()
		}

		p.SetState(376)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(385)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PuppetoParserDEFAULT {
		{
			p.SetState(377)
			p.Match(PuppetoParserDEFAULT)
		}
		{
			p.SetState(378)
			p.Match(PuppetoParserCOLON)
		}
		p.SetState(382)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&568837842538399884) != 0 {
			{
				p.SetState(379)
				p.ProgramBodyWithBreakContinue()
			}

			p.SetState(384)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(387)
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
		p.SetState(389)
		p.Match(PuppetoParserCASE)
	}
	{
		p.SetState(390)
		p.Expression()
	}
	{
		p.SetState(391)
		p.Match(PuppetoParserCOLON)
	}
	{
		p.SetState(392)
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
		p.SetState(394)
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
		p.SetState(396)
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
		p.SetState(398)
		p.Match(PuppetoParserTRY)
	}
	{
		p.SetState(399)
		p.ProgramBody()
	}
	p.SetState(405)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 42, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(400)
			p.Match(PuppetoParserCATCH)
		}
		p.SetState(402)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 41, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(401)
				p.Match(PuppetoParserIDENTIFIER)
			}

		}
		{
			p.SetState(404)
			p.ProgramBody()
		}

	}

	return localctx
}

// IReturnStatementContext is an interface to support dynamic dispatch.
type IReturnStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RETURN() antlr.TerminalNode
	Expression() IExpressionContext

	// IsReturnStatementContext differentiates from other interfaces.
	IsReturnStatementContext()
}

type ReturnStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnStatementContext() *ReturnStatementContext {
	var p = new(ReturnStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PuppetoParserRULE_returnStatement
	return p
}

func (*ReturnStatementContext) IsReturnStatementContext() {}

func NewReturnStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnStatementContext {
	var p = new(ReturnStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PuppetoParserRULE_returnStatement

	return p
}

func (s *ReturnStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnStatementContext) RETURN() antlr.TerminalNode {
	return s.GetToken(PuppetoParserRETURN, 0)
}

func (s *ReturnStatementContext) Expression() IExpressionContext {
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

func (s *ReturnStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PuppetoParserVisitor:
		return t.VisitReturnStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PuppetoParser) ReturnStatement() (localctx IReturnStatementContext) {
	this := p
	_ = this

	localctx = NewReturnStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, PuppetoParserRULE_returnStatement)

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
		p.SetState(407)
		p.Match(PuppetoParserRETURN)
	}
	p.SetState(409)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 43, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(408)
			p.Expression()
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
	ReturnStatement() IReturnStatementContext

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

func (s *StatementListContext) ReturnStatement() IReturnStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturnStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturnStatementContext)
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
	p.EnterRule(localctx, 76, PuppetoParserRULE_statementList)

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

	p.SetState(420)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PuppetoParserLET:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(411)
			p.VariableDefinition()
		}

	case PuppetoParserFN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(412)
			p.FunctionDefinition()
		}

	case PuppetoParserIF:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(413)
			p.IfStatement()
		}

	case PuppetoParserFOR:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(414)
			p.LoopStatement()
		}

	case PuppetoParserSWITCH:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(415)
			p.SwitchStatement()
		}

	case PuppetoParserTRY:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(416)
			p.TryStatement()
		}

	case PuppetoParserECHO:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(417)
			p.EchoStatement()
		}

	case PuppetoParserONE_SIDE_OPERATOR, PuppetoParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(418)
			p.VariableAssignation()
		}

	case PuppetoParserRETURN:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(419)
			p.ReturnStatement()
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
	p.EnterRule(localctx, 78, PuppetoParserRULE_statement)
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
	p.SetState(435)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PuppetoParserRETURN, PuppetoParserONE_SIDE_OPERATOR, PuppetoParserFOR, PuppetoParserECHO, PuppetoParserFN, PuppetoParserIF, PuppetoParserSWITCH, PuppetoParserTRY, PuppetoParserLET, PuppetoParserIDENTIFIER:
		{
			p.SetState(422)
			p.StatementList()
		}

	case PuppetoParserPUP_END_TAG:
		p.SetState(431)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(423)
					p.Match(PuppetoParserPUP_END_TAG)
				}
				p.SetState(427)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4611686018427386928) != 0 {
					{
						p.SetState(424)
						p.HtmlList()
					}

					p.SetState(429)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)
				}
				{
					p.SetState(430)
					p.Match(PuppetoParserPUP_START_TAG)
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(433)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 46, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 80, PuppetoParserRULE_scope)
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
		p.SetState(437)
		p.Match(PuppetoParserL_CURLY)
	}
	p.SetState(439)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&568661920677955724) != 0 {
		{
			p.SetState(438)
			p.ProgramBodyList()
		}

	}
	{
		p.SetState(441)
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
	p.EnterRule(localctx, 82, PuppetoParserRULE_scopeWithBreakContinue)
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
		p.SetState(443)
		p.Match(PuppetoParserL_CURLY)
	}
	p.SetState(447)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&568837842538399884) != 0 {
		{
			p.SetState(444)
			p.ProgramBodyWithBreakContinue()
		}

		p.SetState(449)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(450)
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
	p.EnterRule(localctx, 84, PuppetoParserRULE_programBody)

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

	p.SetState(455)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 50, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(452)
			p.Scope()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(453)
			p.Statement()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(454)
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
	p.EnterRule(localctx, 86, PuppetoParserRULE_programBodyList)

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
	p.SetState(458)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(457)
				p.ProgramBody()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(460)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 51, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 88, PuppetoParserRULE_echoStatement)

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
		p.SetState(462)
		p.Match(PuppetoParserECHO)
	}
	{
		p.SetState(463)
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
	p.EnterRule(localctx, 90, PuppetoParserRULE_programBodyWithBreakContinue)

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

	p.SetState(470)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 52, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(465)
			p.ScopeWithBreakContinue()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(466)
			p.Statement()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(467)
			p.BreakStatement()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(468)
			p.ContinueStatement()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(469)
			p.Expression()
		}

	}

	return localctx
}

// IPupShortCodeContext is an interface to support dynamic dispatch.
type IPupShortCodeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PUP_SHORT_START_TAG() antlr.TerminalNode
	Expression() IExpressionContext
	PUP_SHORT_END_TAG() antlr.TerminalNode

	// IsPupShortCodeContext differentiates from other interfaces.
	IsPupShortCodeContext()
}

type PupShortCodeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPupShortCodeContext() *PupShortCodeContext {
	var p = new(PupShortCodeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PuppetoParserRULE_pupShortCode
	return p
}

func (*PupShortCodeContext) IsPupShortCodeContext() {}

func NewPupShortCodeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PupShortCodeContext {
	var p = new(PupShortCodeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PuppetoParserRULE_pupShortCode

	return p
}

func (s *PupShortCodeContext) GetParser() antlr.Parser { return s.parser }

func (s *PupShortCodeContext) PUP_SHORT_START_TAG() antlr.TerminalNode {
	return s.GetToken(PuppetoParserPUP_SHORT_START_TAG, 0)
}

func (s *PupShortCodeContext) Expression() IExpressionContext {
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

func (s *PupShortCodeContext) PUP_SHORT_END_TAG() antlr.TerminalNode {
	return s.GetToken(PuppetoParserPUP_SHORT_END_TAG, 0)
}

func (s *PupShortCodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PupShortCodeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PupShortCodeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PuppetoParserVisitor:
		return t.VisitPupShortCode(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PuppetoParser) PupShortCode() (localctx IPupShortCodeContext) {
	this := p
	_ = this

	localctx = NewPupShortCodeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, PuppetoParserRULE_pupShortCode)

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
		p.SetState(472)
		p.Match(PuppetoParserPUP_SHORT_START_TAG)
	}
	{
		p.SetState(473)
		p.Expression()
	}
	{
		p.SetState(474)
		p.Match(PuppetoParserPUP_SHORT_END_TAG)
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
	p.EnterRule(localctx, 94, PuppetoParserRULE_pupCode)

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
		p.SetState(476)
		p.Match(PuppetoParserPUP_START_TAG)
	}
	p.SetState(478)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 53, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(477)
			p.ProgramBodyList()
		}

	}
	{
		p.SetState(480)
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
	p.EnterRule(localctx, 96, PuppetoParserRULE_htmlList)
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
		p.SetState(482)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4611686018427386928) != 0) {
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
	AllPupShortCode() []IPupShortCodeContext
	PupShortCode(i int) IPupShortCodeContext
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

func (s *HtmlContext) AllPupShortCode() []IPupShortCodeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPupShortCodeContext); ok {
			len++
		}
	}

	tst := make([]IPupShortCodeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPupShortCodeContext); ok {
			tst[i] = t.(IPupShortCodeContext)
			i++
		}
	}

	return tst
}

func (s *HtmlContext) PupShortCode(i int) IPupShortCodeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPupShortCodeContext); ok {
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

	return t.(IPupShortCodeContext)
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
	p.EnterRule(localctx, 98, PuppetoParserRULE_html)
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
	p.SetState(487)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4611686018427387248) != 0) {
		p.SetState(487)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case PuppetoParserPUP_START_TAG:
			{
				p.SetState(484)
				p.PupCode()
			}

		case PuppetoParserPUP_SHORT_START_TAG:
			{
				p.SetState(485)
				p.PupShortCode()
			}

		case PuppetoParserLT, PuppetoParserGT, PuppetoParserTWO_SIDES_OPERATOR, PuppetoParserONE_SIDE_OPERATOR, PuppetoParserEQ, PuppetoParserNEQ, PuppetoParserLTOE, PuppetoParserGTOE, PuppetoParserFOR, PuppetoParserTO, PuppetoParserECHO, PuppetoParserL_CURLY, PuppetoParserR_CURLY, PuppetoParserCOMMA, PuppetoParserL_SQUARE, PuppetoParserR_SQUARE, PuppetoParserFN, PuppetoParserL_PARENTHES, PuppetoParserR_PARENTHES, PuppetoParserAND, PuppetoParserOR, PuppetoParserCOLON, PuppetoParserDOT, PuppetoParserPLUS, PuppetoParserMINUS, PuppetoParserDOUBLE_STAR, PuppetoParserPERCENT, PuppetoParserBAND, PuppetoParserBOR, PuppetoParserBXOR, PuppetoParserBLS, PuppetoParserBRS, PuppetoParserNOT, PuppetoParserIF, PuppetoParserELSE, PuppetoParserSWITCH, PuppetoParserCASE, PuppetoParserBREAK, PuppetoParserTRY, PuppetoParserCONTINUE, PuppetoParserCATCH, PuppetoParserDEFAULT, PuppetoParserLET, PuppetoParserSTAR, PuppetoParserSLASH, PuppetoParserFLOAT, PuppetoParserINTEGER, PuppetoParserNIL, PuppetoParserBOOLEAN, PuppetoParserIDENTIFIER, PuppetoParserSTRING, PuppetoParserWS, PuppetoParserSEMICOLON, PuppetoParserHTML:
			{
				p.SetState(486)
				p.HtmlList()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(489)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(491)
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
	p.EnterRule(localctx, 100, PuppetoParserRULE_program)
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
		p.SetState(493)
		p.Match(PuppetoParserPUP_START_TAG)
	}
	p.SetState(495)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 56, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(494)
			p.ProgramBodyList()
		}

	}
	p.SetState(498)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PuppetoParserPUP_END_TAG {
		{
			p.SetState(497)
			p.Match(PuppetoParserPUP_END_TAG)
		}

	}
	{
		p.SetState(500)
		p.Match(PuppetoParserEOF)
	}

	return localctx
}
