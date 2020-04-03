// Code generated from STIXPattern.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // STIXPattern

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 56, 235,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	7, 3, 48, 10, 3, 12, 3, 14, 3, 51, 11, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 7, 4, 59, 10, 4, 12, 4, 14, 4, 62, 11, 4, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 7, 5, 70, 10, 5, 12, 5, 14, 5, 73, 11, 5, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 84, 10, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 7, 6, 92, 10, 6, 12, 6, 14, 6, 95, 11, 6, 3, 7, 3, 7,
	3, 7, 3, 7, 3, 7, 3, 7, 7, 7, 103, 10, 7, 12, 7, 14, 7, 106, 11, 7, 3,
	8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 7, 8, 114, 10, 8, 12, 8, 14, 8, 117, 11,
	8, 3, 9, 3, 9, 5, 9, 121, 10, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 128,
	10, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 135, 10, 9, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 9, 5, 9, 142, 10, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 149,
	10, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 156, 10, 9, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 9, 5, 9, 163, 10, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 9, 5, 9, 174, 10, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 11,
	3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3,
	13, 5, 13, 193, 10, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16,
	3, 16, 3, 16, 3, 16, 5, 16, 205, 10, 16, 3, 16, 3, 16, 7, 16, 209, 10,
	16, 12, 16, 14, 16, 212, 11, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3,
	17, 7, 17, 220, 10, 17, 12, 17, 14, 17, 223, 11, 17, 3, 17, 3, 17, 5, 17,
	227, 10, 17, 3, 18, 3, 18, 5, 18, 231, 10, 18, 3, 19, 3, 19, 3, 19, 2,
	9, 4, 6, 8, 10, 12, 14, 30, 20, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22,
	24, 26, 28, 30, 32, 34, 36, 2, 9, 3, 2, 33, 34, 3, 2, 35, 38, 4, 2, 4,
	4, 6, 6, 3, 2, 31, 32, 4, 2, 9, 9, 31, 31, 4, 2, 3, 4, 52, 52, 4, 2, 3,
	9, 11, 11, 2, 246, 2, 38, 3, 2, 2, 2, 4, 41, 3, 2, 2, 2, 6, 52, 3, 2, 2,
	2, 8, 63, 3, 2, 2, 2, 10, 83, 3, 2, 2, 2, 12, 96, 3, 2, 2, 2, 14, 107,
	3, 2, 2, 2, 16, 173, 3, 2, 2, 2, 18, 175, 3, 2, 2, 2, 20, 180, 3, 2, 2,
	2, 22, 184, 3, 2, 2, 2, 24, 188, 3, 2, 2, 2, 26, 194, 3, 2, 2, 2, 28, 196,
	3, 2, 2, 2, 30, 204, 3, 2, 2, 2, 32, 226, 3, 2, 2, 2, 34, 230, 3, 2, 2,
	2, 36, 232, 3, 2, 2, 2, 38, 39, 5, 4, 3, 2, 39, 40, 7, 2, 2, 3, 40, 3,
	3, 2, 2, 2, 41, 42, 8, 3, 1, 2, 42, 43, 5, 6, 4, 2, 43, 49, 3, 2, 2, 2,
	44, 45, 12, 4, 2, 2, 45, 46, 7, 15, 2, 2, 46, 48, 5, 4, 3, 5, 47, 44, 3,
	2, 2, 2, 48, 51, 3, 2, 2, 2, 49, 47, 3, 2, 2, 2, 49, 50, 3, 2, 2, 2, 50,
	5, 3, 2, 2, 2, 51, 49, 3, 2, 2, 2, 52, 53, 8, 4, 1, 2, 53, 54, 5, 8, 5,
	2, 54, 60, 3, 2, 2, 2, 55, 56, 12, 4, 2, 2, 56, 57, 7, 13, 2, 2, 57, 59,
	5, 6, 4, 5, 58, 55, 3, 2, 2, 2, 59, 62, 3, 2, 2, 2, 60, 58, 3, 2, 2, 2,
	60, 61, 3, 2, 2, 2, 61, 7, 3, 2, 2, 2, 62, 60, 3, 2, 2, 2, 63, 64, 8, 5,
	1, 2, 64, 65, 5, 10, 6, 2, 65, 71, 3, 2, 2, 2, 66, 67, 12, 4, 2, 2, 67,
	68, 7, 12, 2, 2, 68, 70, 5, 8, 5, 5, 69, 66, 3, 2, 2, 2, 70, 73, 3, 2,
	2, 2, 71, 69, 3, 2, 2, 2, 71, 72, 3, 2, 2, 2, 72, 9, 3, 2, 2, 2, 73, 71,
	3, 2, 2, 2, 74, 75, 8, 6, 1, 2, 75, 76, 7, 46, 2, 2, 76, 77, 5, 12, 7,
	2, 77, 78, 7, 45, 2, 2, 78, 84, 3, 2, 2, 2, 79, 80, 7, 44, 2, 2, 80, 81,
	5, 4, 3, 2, 81, 82, 7, 43, 2, 2, 82, 84, 3, 2, 2, 2, 83, 74, 3, 2, 2, 2,
	83, 79, 3, 2, 2, 2, 84, 93, 3, 2, 2, 2, 85, 86, 12, 5, 2, 2, 86, 92, 5,
	18, 10, 2, 87, 88, 12, 4, 2, 2, 88, 92, 5, 20, 11, 2, 89, 90, 12, 3, 2,
	2, 90, 92, 5, 22, 12, 2, 91, 85, 3, 2, 2, 2, 91, 87, 3, 2, 2, 2, 91, 89,
	3, 2, 2, 2, 92, 95, 3, 2, 2, 2, 93, 91, 3, 2, 2, 2, 93, 94, 3, 2, 2, 2,
	94, 11, 3, 2, 2, 2, 95, 93, 3, 2, 2, 2, 96, 97, 8, 7, 1, 2, 97, 98, 5,
	14, 8, 2, 98, 104, 3, 2, 2, 2, 99, 100, 12, 4, 2, 2, 100, 101, 7, 13, 2,
	2, 101, 103, 5, 12, 7, 5, 102, 99, 3, 2, 2, 2, 103, 106, 3, 2, 2, 2, 104,
	102, 3, 2, 2, 2, 104, 105, 3, 2, 2, 2, 105, 13, 3, 2, 2, 2, 106, 104, 3,
	2, 2, 2, 107, 108, 8, 8, 1, 2, 108, 109, 5, 16, 9, 2, 109, 115, 3, 2, 2,
	2, 110, 111, 12, 4, 2, 2, 111, 112, 7, 12, 2, 2, 112, 114, 5, 14, 8, 5,
	113, 110, 3, 2, 2, 2, 114, 117, 3, 2, 2, 2, 115, 113, 3, 2, 2, 2, 115,
	116, 3, 2, 2, 2, 116, 15, 3, 2, 2, 2, 117, 115, 3, 2, 2, 2, 118, 120, 5,
	24, 13, 2, 119, 121, 7, 14, 2, 2, 120, 119, 3, 2, 2, 2, 120, 121, 3, 2,
	2, 2, 121, 122, 3, 2, 2, 2, 122, 123, 9, 2, 2, 2, 123, 124, 5, 34, 18,
	2, 124, 174, 3, 2, 2, 2, 125, 127, 5, 24, 13, 2, 126, 128, 7, 14, 2, 2,
	127, 126, 3, 2, 2, 2, 127, 128, 3, 2, 2, 2, 128, 129, 3, 2, 2, 2, 129,
	130, 9, 3, 2, 2, 130, 131, 5, 36, 19, 2, 131, 174, 3, 2, 2, 2, 132, 134,
	5, 24, 13, 2, 133, 135, 7, 14, 2, 2, 134, 133, 3, 2, 2, 2, 134, 135, 3,
	2, 2, 2, 135, 136, 3, 2, 2, 2, 136, 137, 7, 22, 2, 2, 137, 138, 5, 32,
	17, 2, 138, 174, 3, 2, 2, 2, 139, 141, 5, 24, 13, 2, 140, 142, 7, 14, 2,
	2, 141, 140, 3, 2, 2, 2, 141, 142, 3, 2, 2, 2, 142, 143, 3, 2, 2, 2, 143,
	144, 7, 16, 2, 2, 144, 145, 7, 9, 2, 2, 145, 174, 3, 2, 2, 2, 146, 148,
	5, 24, 13, 2, 147, 149, 7, 14, 2, 2, 148, 147, 3, 2, 2, 2, 148, 149, 3,
	2, 2, 2, 149, 150, 3, 2, 2, 2, 150, 151, 7, 17, 2, 2, 151, 152, 7, 9, 2,
	2, 152, 174, 3, 2, 2, 2, 153, 155, 5, 24, 13, 2, 154, 156, 7, 14, 2, 2,
	155, 154, 3, 2, 2, 2, 155, 156, 3, 2, 2, 2, 156, 157, 3, 2, 2, 2, 157,
	158, 7, 19, 2, 2, 158, 159, 7, 9, 2, 2, 159, 174, 3, 2, 2, 2, 160, 162,
	5, 24, 13, 2, 161, 163, 7, 14, 2, 2, 162, 161, 3, 2, 2, 2, 162, 163, 3,
	2, 2, 2, 163, 164, 3, 2, 2, 2, 164, 165, 7, 18, 2, 2, 165, 166, 7, 9, 2,
	2, 166, 174, 3, 2, 2, 2, 167, 168, 7, 44, 2, 2, 168, 169, 5, 12, 7, 2,
	169, 170, 7, 43, 2, 2, 170, 174, 3, 2, 2, 2, 171, 172, 7, 20, 2, 2, 172,
	174, 5, 24, 13, 2, 173, 118, 3, 2, 2, 2, 173, 125, 3, 2, 2, 2, 173, 132,
	3, 2, 2, 2, 173, 139, 3, 2, 2, 2, 173, 146, 3, 2, 2, 2, 173, 153, 3, 2,
	2, 2, 173, 160, 3, 2, 2, 2, 173, 167, 3, 2, 2, 2, 173, 171, 3, 2, 2, 2,
	174, 17, 3, 2, 2, 2, 175, 176, 7, 23, 2, 2, 176, 177, 7, 11, 2, 2, 177,
	178, 7, 24, 2, 2, 178, 179, 7, 11, 2, 2, 179, 19, 3, 2, 2, 2, 180, 181,
	7, 28, 2, 2, 181, 182, 9, 4, 2, 2, 182, 183, 7, 25, 2, 2, 183, 21, 3, 2,
	2, 2, 184, 185, 7, 29, 2, 2, 185, 186, 7, 4, 2, 2, 186, 187, 7, 30, 2,
	2, 187, 23, 3, 2, 2, 2, 188, 189, 5, 26, 14, 2, 189, 190, 7, 40, 2, 2,
	190, 192, 5, 28, 15, 2, 191, 193, 5, 30, 16, 2, 192, 191, 3, 2, 2, 2, 192,
	193, 3, 2, 2, 2, 193, 25, 3, 2, 2, 2, 194, 195, 9, 5, 2, 2, 195, 27, 3,
	2, 2, 2, 196, 197, 9, 6, 2, 2, 197, 29, 3, 2, 2, 2, 198, 199, 8, 16, 1,
	2, 199, 200, 7, 41, 2, 2, 200, 205, 9, 6, 2, 2, 201, 202, 7, 46, 2, 2,
	202, 203, 9, 7, 2, 2, 203, 205, 7, 45, 2, 2, 204, 198, 3, 2, 2, 2, 204,
	201, 3, 2, 2, 2, 205, 210, 3, 2, 2, 2, 206, 207, 12, 5, 2, 2, 207, 209,
	5, 30, 16, 6, 208, 206, 3, 2, 2, 2, 209, 212, 3, 2, 2, 2, 210, 208, 3,
	2, 2, 2, 210, 211, 3, 2, 2, 2, 211, 31, 3, 2, 2, 2, 212, 210, 3, 2, 2,
	2, 213, 214, 7, 44, 2, 2, 214, 227, 7, 43, 2, 2, 215, 216, 7, 44, 2, 2,
	216, 221, 5, 34, 18, 2, 217, 218, 7, 42, 2, 2, 218, 220, 5, 34, 18, 2,
	219, 217, 3, 2, 2, 2, 220, 223, 3, 2, 2, 2, 221, 219, 3, 2, 2, 2, 221,
	222, 3, 2, 2, 2, 222, 224, 3, 2, 2, 2, 223, 221, 3, 2, 2, 2, 224, 225,
	7, 43, 2, 2, 225, 227, 3, 2, 2, 2, 226, 213, 3, 2, 2, 2, 226, 215, 3, 2,
	2, 2, 227, 33, 3, 2, 2, 2, 228, 231, 5, 36, 19, 2, 229, 231, 7, 10, 2,
	2, 230, 228, 3, 2, 2, 2, 230, 229, 3, 2, 2, 2, 231, 35, 3, 2, 2, 2, 232,
	233, 9, 8, 2, 2, 233, 37, 3, 2, 2, 2, 24, 49, 60, 71, 83, 91, 93, 104,
	115, 120, 127, 134, 141, 148, 155, 162, 173, 192, 204, 210, 221, 226, 230,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "", "", "", "", "", "", "", "", "", "'AND'", "'OR'", "'NOT'", "'FOLLOWEDBY'",
	"'LIKE'", "'MATCHES'", "'ISSUPERSET'", "'ISSUBSET'", "'EXISTS'", "'LAST'",
	"'IN'", "'START'", "'STOP'", "'SECONDS'", "'true'", "'false'", "'WITHIN'",
	"'REPEATS'", "'TIMES'", "", "", "", "", "'<'", "'<='", "'>'", "'>='", "'''",
	"':'", "'.'", "','", "')'", "'('", "']'", "'['", "'+'", "", "'-'", "'^'",
	"'/'", "'*'",
}
var symbolicNames = []string{
	"", "IntNegLiteral", "IntPosLiteral", "FloatNegLiteral", "FloatPosLiteral",
	"HexLiteral", "BinaryLiteral", "StringLiteral", "BoolLiteral", "TimestampLiteral",
	"AND", "OR", "NOT", "FOLLOWEDBY", "LIKE", "MATCHES", "ISSUPERSET", "ISSUBSET",
	"EXISTS", "LAST", "IN", "START", "STOP", "SECONDS", "TRUE", "FALSE", "WITHIN",
	"REPEATS", "TIMES", "IdentifierWithoutHyphen", "IdentifierWithHyphen",
	"EQ", "NEQ", "LT", "LE", "GT", "GE", "QUOTE", "COLON", "DOT", "COMMA",
	"RPAREN", "LPAREN", "RBRACK", "LBRACK", "PLUS", "HYPHEN", "MINUS", "POWER_OP",
	"DIVIDE", "ASTERISK", "WS", "COMMENT", "LINE_COMMENT", "InvalidCharacter",
}

var ruleNames = []string{
	"pattern", "observationExpressions", "observationExpressionOr", "observationExpressionAnd",
	"observationExpression", "comparisonExpression", "comparisonExpressionAnd",
	"propTest", "startStopQualifier", "withinQualifier", "repeatedQualifier",
	"objectPath", "objectType", "firstPathComponent", "objectPathComponent",
	"setLiteral", "primitiveLiteral", "orderableLiteral",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type STIXPatternParser struct {
	*antlr.BaseParser
}

func NewSTIXPatternParser(input antlr.TokenStream) *STIXPatternParser {
	this := new(STIXPatternParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "STIXPattern.g4"

	return this
}

// STIXPatternParser tokens.
const (
	STIXPatternParserEOF                     = antlr.TokenEOF
	STIXPatternParserIntNegLiteral           = 1
	STIXPatternParserIntPosLiteral           = 2
	STIXPatternParserFloatNegLiteral         = 3
	STIXPatternParserFloatPosLiteral         = 4
	STIXPatternParserHexLiteral              = 5
	STIXPatternParserBinaryLiteral           = 6
	STIXPatternParserStringLiteral           = 7
	STIXPatternParserBoolLiteral             = 8
	STIXPatternParserTimestampLiteral        = 9
	STIXPatternParserAND                     = 10
	STIXPatternParserOR                      = 11
	STIXPatternParserNOT                     = 12
	STIXPatternParserFOLLOWEDBY              = 13
	STIXPatternParserLIKE                    = 14
	STIXPatternParserMATCHES                 = 15
	STIXPatternParserISSUPERSET              = 16
	STIXPatternParserISSUBSET                = 17
	STIXPatternParserEXISTS                  = 18
	STIXPatternParserLAST                    = 19
	STIXPatternParserIN                      = 20
	STIXPatternParserSTART                   = 21
	STIXPatternParserSTOP                    = 22
	STIXPatternParserSECONDS                 = 23
	STIXPatternParserTRUE                    = 24
	STIXPatternParserFALSE                   = 25
	STIXPatternParserWITHIN                  = 26
	STIXPatternParserREPEATS                 = 27
	STIXPatternParserTIMES                   = 28
	STIXPatternParserIdentifierWithoutHyphen = 29
	STIXPatternParserIdentifierWithHyphen    = 30
	STIXPatternParserEQ                      = 31
	STIXPatternParserNEQ                     = 32
	STIXPatternParserLT                      = 33
	STIXPatternParserLE                      = 34
	STIXPatternParserGT                      = 35
	STIXPatternParserGE                      = 36
	STIXPatternParserQUOTE                   = 37
	STIXPatternParserCOLON                   = 38
	STIXPatternParserDOT                     = 39
	STIXPatternParserCOMMA                   = 40
	STIXPatternParserRPAREN                  = 41
	STIXPatternParserLPAREN                  = 42
	STIXPatternParserRBRACK                  = 43
	STIXPatternParserLBRACK                  = 44
	STIXPatternParserPLUS                    = 45
	STIXPatternParserHYPHEN                  = 46
	STIXPatternParserMINUS                   = 47
	STIXPatternParserPOWER_OP                = 48
	STIXPatternParserDIVIDE                  = 49
	STIXPatternParserASTERISK                = 50
	STIXPatternParserWS                      = 51
	STIXPatternParserCOMMENT                 = 52
	STIXPatternParserLINE_COMMENT            = 53
	STIXPatternParserInvalidCharacter        = 54
)

// STIXPatternParser rules.
const (
	STIXPatternParserRULE_pattern                  = 0
	STIXPatternParserRULE_observationExpressions   = 1
	STIXPatternParserRULE_observationExpressionOr  = 2
	STIXPatternParserRULE_observationExpressionAnd = 3
	STIXPatternParserRULE_observationExpression    = 4
	STIXPatternParserRULE_comparisonExpression     = 5
	STIXPatternParserRULE_comparisonExpressionAnd  = 6
	STIXPatternParserRULE_propTest                 = 7
	STIXPatternParserRULE_startStopQualifier       = 8
	STIXPatternParserRULE_withinQualifier          = 9
	STIXPatternParserRULE_repeatedQualifier        = 10
	STIXPatternParserRULE_objectPath               = 11
	STIXPatternParserRULE_objectType               = 12
	STIXPatternParserRULE_firstPathComponent       = 13
	STIXPatternParserRULE_objectPathComponent      = 14
	STIXPatternParserRULE_setLiteral               = 15
	STIXPatternParserRULE_primitiveLiteral         = 16
	STIXPatternParserRULE_orderableLiteral         = 17
)

// IPatternContext is an interface to support dynamic dispatch.
type IPatternContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPatternContext differentiates from other interfaces.
	IsPatternContext()
}

type PatternContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPatternContext() *PatternContext {
	var p = new(PatternContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = STIXPatternParserRULE_pattern
	return p
}

func (*PatternContext) IsPatternContext() {}

func NewPatternContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PatternContext {
	var p = new(PatternContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = STIXPatternParserRULE_pattern

	return p
}

func (s *PatternContext) GetParser() antlr.Parser { return s.parser }

func (s *PatternContext) ObservationExpressions() IObservationExpressionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObservationExpressionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObservationExpressionsContext)
}

func (s *PatternContext) EOF() antlr.TerminalNode {
	return s.GetToken(STIXPatternParserEOF, 0)
}

func (s *PatternContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PatternContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PatternContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STIXPatternListener); ok {
		listenerT.EnterPattern(s)
	}
}

func (s *PatternContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STIXPatternListener); ok {
		listenerT.ExitPattern(s)
	}
}

func (p *STIXPatternParser) Pattern() (localctx IPatternContext) {
	localctx = NewPatternContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, STIXPatternParserRULE_pattern)

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
		p.SetState(36)
		p.observationExpressions(0)
	}
	{
		p.SetState(37)
		p.Match(STIXPatternParserEOF)
	}

	return localctx
}

// IObservationExpressionsContext is an interface to support dynamic dispatch.
type IObservationExpressionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObservationExpressionsContext differentiates from other interfaces.
	IsObservationExpressionsContext()
}

type ObservationExpressionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObservationExpressionsContext() *ObservationExpressionsContext {
	var p = new(ObservationExpressionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = STIXPatternParserRULE_observationExpressions
	return p
}

func (*ObservationExpressionsContext) IsObservationExpressionsContext() {}

func NewObservationExpressionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObservationExpressionsContext {
	var p = new(ObservationExpressionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = STIXPatternParserRULE_observationExpressions

	return p
}

func (s *ObservationExpressionsContext) GetParser() antlr.Parser { return s.parser }

func (s *ObservationExpressionsContext) ObservationExpressionOr() IObservationExpressionOrContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObservationExpressionOrContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObservationExpressionOrContext)
}

func (s *ObservationExpressionsContext) AllObservationExpressions() []IObservationExpressionsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IObservationExpressionsContext)(nil)).Elem())
	var tst = make([]IObservationExpressionsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IObservationExpressionsContext)
		}
	}

	return tst
}

func (s *ObservationExpressionsContext) ObservationExpressions(i int) IObservationExpressionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObservationExpressionsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IObservationExpressionsContext)
}

func (s *ObservationExpressionsContext) FOLLOWEDBY() antlr.TerminalNode {
	return s.GetToken(STIXPatternParserFOLLOWEDBY, 0)
}

func (s *ObservationExpressionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObservationExpressionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObservationExpressionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STIXPatternListener); ok {
		listenerT.EnterObservationExpressions(s)
	}
}

func (s *ObservationExpressionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STIXPatternListener); ok {
		listenerT.ExitObservationExpressions(s)
	}
}

func (p *STIXPatternParser) ObservationExpressions() (localctx IObservationExpressionsContext) {
	return p.observationExpressions(0)
}

func (p *STIXPatternParser) observationExpressions(_p int) (localctx IObservationExpressionsContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewObservationExpressionsContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IObservationExpressionsContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, STIXPatternParserRULE_observationExpressions, _p)

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
		p.SetState(40)
		p.observationExpressionOr(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(47)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewObservationExpressionsContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, STIXPatternParserRULE_observationExpressions)
			p.SetState(42)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(43)
				p.Match(STIXPatternParserFOLLOWEDBY)
			}
			{
				p.SetState(44)
				p.observationExpressions(3)
			}

		}
		p.SetState(49)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())
	}

	return localctx
}

// IObservationExpressionOrContext is an interface to support dynamic dispatch.
type IObservationExpressionOrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObservationExpressionOrContext differentiates from other interfaces.
	IsObservationExpressionOrContext()
}

type ObservationExpressionOrContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObservationExpressionOrContext() *ObservationExpressionOrContext {
	var p = new(ObservationExpressionOrContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = STIXPatternParserRULE_observationExpressionOr
	return p
}

func (*ObservationExpressionOrContext) IsObservationExpressionOrContext() {}

func NewObservationExpressionOrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObservationExpressionOrContext {
	var p = new(ObservationExpressionOrContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = STIXPatternParserRULE_observationExpressionOr

	return p
}

func (s *ObservationExpressionOrContext) GetParser() antlr.Parser { return s.parser }

func (s *ObservationExpressionOrContext) ObservationExpressionAnd() IObservationExpressionAndContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObservationExpressionAndContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObservationExpressionAndContext)
}

func (s *ObservationExpressionOrContext) AllObservationExpressionOr() []IObservationExpressionOrContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IObservationExpressionOrContext)(nil)).Elem())
	var tst = make([]IObservationExpressionOrContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IObservationExpressionOrContext)
		}
	}

	return tst
}

func (s *ObservationExpressionOrContext) ObservationExpressionOr(i int) IObservationExpressionOrContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObservationExpressionOrContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IObservationExpressionOrContext)
}

func (s *ObservationExpressionOrContext) OR() antlr.TerminalNode {
	return s.GetToken(STIXPatternParserOR, 0)
}

func (s *ObservationExpressionOrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObservationExpressionOrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObservationExpressionOrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STIXPatternListener); ok {
		listenerT.EnterObservationExpressionOr(s)
	}
}

func (s *ObservationExpressionOrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STIXPatternListener); ok {
		listenerT.ExitObservationExpressionOr(s)
	}
}

func (p *STIXPatternParser) ObservationExpressionOr() (localctx IObservationExpressionOrContext) {
	return p.observationExpressionOr(0)
}

func (p *STIXPatternParser) observationExpressionOr(_p int) (localctx IObservationExpressionOrContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewObservationExpressionOrContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IObservationExpressionOrContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 4
	p.EnterRecursionRule(localctx, 4, STIXPatternParserRULE_observationExpressionOr, _p)

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
		p.SetState(51)
		p.observationExpressionAnd(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(58)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewObservationExpressionOrContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, STIXPatternParserRULE_observationExpressionOr)
			p.SetState(53)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(54)
				p.Match(STIXPatternParserOR)
			}
			{
				p.SetState(55)
				p.observationExpressionOr(3)
			}

		}
		p.SetState(60)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())
	}

	return localctx
}

// IObservationExpressionAndContext is an interface to support dynamic dispatch.
type IObservationExpressionAndContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObservationExpressionAndContext differentiates from other interfaces.
	IsObservationExpressionAndContext()
}

type ObservationExpressionAndContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObservationExpressionAndContext() *ObservationExpressionAndContext {
	var p = new(ObservationExpressionAndContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = STIXPatternParserRULE_observationExpressionAnd
	return p
}

func (*ObservationExpressionAndContext) IsObservationExpressionAndContext() {}

func NewObservationExpressionAndContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObservationExpressionAndContext {
	var p = new(ObservationExpressionAndContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = STIXPatternParserRULE_observationExpressionAnd

	return p
}

func (s *ObservationExpressionAndContext) GetParser() antlr.Parser { return s.parser }

func (s *ObservationExpressionAndContext) ObservationExpression() IObservationExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObservationExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObservationExpressionContext)
}

func (s *ObservationExpressionAndContext) AllObservationExpressionAnd() []IObservationExpressionAndContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IObservationExpressionAndContext)(nil)).Elem())
	var tst = make([]IObservationExpressionAndContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IObservationExpressionAndContext)
		}
	}

	return tst
}

func (s *ObservationExpressionAndContext) ObservationExpressionAnd(i int) IObservationExpressionAndContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObservationExpressionAndContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IObservationExpressionAndContext)
}

func (s *ObservationExpressionAndContext) AND() antlr.TerminalNode {
	return s.GetToken(STIXPatternParserAND, 0)
}

func (s *ObservationExpressionAndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObservationExpressionAndContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObservationExpressionAndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STIXPatternListener); ok {
		listenerT.EnterObservationExpressionAnd(s)
	}
}

func (s *ObservationExpressionAndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STIXPatternListener); ok {
		listenerT.ExitObservationExpressionAnd(s)
	}
}

func (p *STIXPatternParser) ObservationExpressionAnd() (localctx IObservationExpressionAndContext) {
	return p.observationExpressionAnd(0)
}

func (p *STIXPatternParser) observationExpressionAnd(_p int) (localctx IObservationExpressionAndContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewObservationExpressionAndContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IObservationExpressionAndContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 6
	p.EnterRecursionRule(localctx, 6, STIXPatternParserRULE_observationExpressionAnd, _p)

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
		p.SetState(62)
		p.observationExpression(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(69)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewObservationExpressionAndContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, STIXPatternParserRULE_observationExpressionAnd)
			p.SetState(64)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(65)
				p.Match(STIXPatternParserAND)
			}
			{
				p.SetState(66)
				p.observationExpressionAnd(3)
			}

		}
		p.SetState(71)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}

	return localctx
}

// IObservationExpressionContext is an interface to support dynamic dispatch.
type IObservationExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObservationExpressionContext differentiates from other interfaces.
	IsObservationExpressionContext()
}

type ObservationExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObservationExpressionContext() *ObservationExpressionContext {
	var p = new(ObservationExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = STIXPatternParserRULE_observationExpression
	return p
}

func (*ObservationExpressionContext) IsObservationExpressionContext() {}

func NewObservationExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObservationExpressionContext {
	var p = new(ObservationExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = STIXPatternParserRULE_observationExpression

	return p
}

func (s *ObservationExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ObservationExpressionContext) CopyFrom(ctx *ObservationExpressionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ObservationExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObservationExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ObservationExpressionRepeatedContext struct {
	*ObservationExpressionContext
}

func NewObservationExpressionRepeatedContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ObservationExpressionRepeatedContext {
	var p = new(ObservationExpressionRepeatedContext)

	p.ObservationExpressionContext = NewEmptyObservationExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ObservationExpressionContext))

	return p
}

func (s *ObservationExpressionRepeatedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObservationExpressionRepeatedContext) ObservationExpression() IObservationExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObservationExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObservationExpressionContext)
}

func (s *ObservationExpressionRepeatedContext) RepeatedQualifier() IRepeatedQualifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRepeatedQualifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRepeatedQualifierContext)
}

func (s *ObservationExpressionRepeatedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STIXPatternListener); ok {
		listenerT.EnterObservationExpressionRepeated(s)
	}
}

func (s *ObservationExpressionRepeatedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STIXPatternListener); ok {
		listenerT.ExitObservationExpressionRepeated(s)
	}
}

type ObservationExpressionSimpleContext struct {
	*ObservationExpressionContext
}

func NewObservationExpressionSimpleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ObservationExpressionSimpleContext {
	var p = new(ObservationExpressionSimpleContext)

	p.ObservationExpressionContext = NewEmptyObservationExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ObservationExpressionContext))

	return p
}

func (s *ObservationExpressionSimpleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObservationExpressionSimpleContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(STIXPatternParserLBRACK, 0)
}

func (s *ObservationExpressionSimpleContext) ComparisonExpression() IComparisonExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComparisonExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComparisonExpressionContext)
}

func (s *ObservationExpressionSimpleContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(STIXPatternParserRBRACK, 0)
}

func (s *ObservationExpressionSimpleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STIXPatternListener); ok {
		listenerT.EnterObservationExpressionSimple(s)
	}
}

func (s *ObservationExpressionSimpleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STIXPatternListener); ok {
		listenerT.ExitObservationExpressionSimple(s)
	}
}

type ObservationExpressionCompoundContext struct {
	*ObservationExpressionContext
}

func NewObservationExpressionCompoundContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ObservationExpressionCompoundContext {
	var p = new(ObservationExpressionCompoundContext)

	p.ObservationExpressionContext = NewEmptyObservationExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ObservationExpressionContext))

	return p
}

func (s *ObservationExpressionCompoundContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObservationExpressionCompoundContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(STIXPatternParserLPAREN, 0)
}

func (s *ObservationExpressionCompoundContext) ObservationExpressions() IObservationExpressionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObservationExpressionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObservationExpressionsContext)
}

func (s *ObservationExpressionCompoundContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(STIXPatternParserRPAREN, 0)
}

func (s *ObservationExpressionCompoundContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STIXPatternListener); ok {
		listenerT.EnterObservationExpressionCompound(s)
	}
}

func (s *ObservationExpressionCompoundContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STIXPatternListener); ok {
		listenerT.ExitObservationExpressionCompound(s)
	}
}

type ObservationExpressionWithinContext struct {
	*ObservationExpressionContext
}

func NewObservationExpressionWithinContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ObservationExpressionWithinContext {
	var p = new(ObservationExpressionWithinContext)

	p.ObservationExpressionContext = NewEmptyObservationExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ObservationExpressionContext))

	return p
}

func (s *ObservationExpressionWithinContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObservationExpressionWithinContext) ObservationExpression() IObservationExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObservationExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObservationExpressionContext)
}

func (s *ObservationExpressionWithinContext) WithinQualifier() IWithinQualifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWithinQualifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWithinQualifierContext)
}

func (s *ObservationExpressionWithinContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STIXPatternListener); ok {
		listenerT.EnterObservationExpressionWithin(s)
	}
}

func (s *ObservationExpressionWithinContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STIXPatternListener); ok {
		listenerT.ExitObservationExpressionWithin(s)
	}
}

type ObservationExpressionStartStopContext struct {
	*ObservationExpressionContext
}

func NewObservationExpressionStartStopContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ObservationExpressionStartStopContext {
	var p = new(ObservationExpressionStartStopContext)

	p.ObservationExpressionContext = NewEmptyObservationExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ObservationExpressionContext))

	return p
}

func (s *ObservationExpressionStartStopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObservationExpressionStartStopContext) ObservationExpression() IObservationExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObservationExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObservationExpressionContext)
}

func (s *ObservationExpressionStartStopContext) StartStopQualifier() IStartStopQualifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStartStopQualifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStartStopQualifierContext)
}

func (s *ObservationExpressionStartStopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STIXPatternListener); ok {
		listenerT.EnterObservationExpressionStartStop(s)
	}
}

func (s *ObservationExpressionStartStopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STIXPatternListener); ok {
		listenerT.ExitObservationExpressionStartStop(s)
	}
}

func (p *STIXPatternParser) ObservationExpression() (localctx IObservationExpressionContext) {
	return p.observationExpression(0)
}

func (p *STIXPatternParser) observationExpression(_p int) (localctx IObservationExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewObservationExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IObservationExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 8
	p.EnterRecursionRule(localctx, 8, STIXPatternParserRULE_observationExpression, _p)

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
	p.SetState(81)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case STIXPatternParserLBRACK:
		localctx = NewObservationExpressionSimpleContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(73)
			p.Match(STIXPatternParserLBRACK)
		}
		{
			p.SetState(74)
			p.comparisonExpression(0)
		}
		{
			p.SetState(75)
			p.Match(STIXPatternParserRBRACK)
		}

	case STIXPatternParserLPAREN:
		localctx = NewObservationExpressionCompoundContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(77)
			p.Match(STIXPatternParserLPAREN)
		}
		{
			p.SetState(78)
			p.observationExpressions(0)
		}
		{
			p.SetState(79)
			p.Match(STIXPatternParserRPAREN)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(91)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(89)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
			case 1:
				localctx = NewObservationExpressionStartStopContext(p, NewObservationExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, STIXPatternParserRULE_observationExpression)
				p.SetState(83)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(84)
					p.StartStopQualifier()
				}

			case 2:
				localctx = NewObservationExpressionWithinContext(p, NewObservationExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, STIXPatternParserRULE_observationExpression)
				p.SetState(85)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(86)
					p.WithinQualifier()
				}

			case 3:
				localctx = NewObservationExpressionRepeatedContext(p, NewObservationExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, STIXPatternParserRULE_observationExpression)
				p.SetState(87)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(88)
					p.RepeatedQualifier()
				}

			}

		}
		p.SetState(93)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())
	}

	return localctx
}

// IComparisonExpressionContext is an interface to support dynamic dispatch.
type IComparisonExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

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
	p.RuleIndex = STIXPatternParserRULE_comparisonExpression
	return p
}

func (*ComparisonExpressionContext) IsComparisonExpressionContext() {}

func NewComparisonExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComparisonExpressionContext {
	var p = new(ComparisonExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = STIXPatternParserRULE_comparisonExpression

	return p
}

func (s *ComparisonExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ComparisonExpressionContext) ComparisonExpressionAnd() IComparisonExpressionAndContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComparisonExpressionAndContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComparisonExpressionAndContext)
}

func (s *ComparisonExpressionContext) AllComparisonExpression() []IComparisonExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IComparisonExpressionContext)(nil)).Elem())
	var tst = make([]IComparisonExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IComparisonExpressionContext)
		}
	}

	return tst
}

func (s *ComparisonExpressionContext) ComparisonExpression(i int) IComparisonExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComparisonExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IComparisonExpressionContext)
}

func (s *ComparisonExpressionContext) OR() antlr.TerminalNode {
	return s.GetToken(STIXPatternParserOR, 0)
}

func (s *ComparisonExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparisonExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComparisonExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STIXPatternListener); ok {
		listenerT.EnterComparisonExpression(s)
	}
}

func (s *ComparisonExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STIXPatternListener); ok {
		listenerT.ExitComparisonExpression(s)
	}
}

func (p *STIXPatternParser) ComparisonExpression() (localctx IComparisonExpressionContext) {
	return p.comparisonExpression(0)
}

func (p *STIXPatternParser) comparisonExpression(_p int) (localctx IComparisonExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewComparisonExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IComparisonExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 10
	p.EnterRecursionRule(localctx, 10, STIXPatternParserRULE_comparisonExpression, _p)

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
		p.SetState(95)
		p.comparisonExpressionAnd(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(102)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewComparisonExpressionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, STIXPatternParserRULE_comparisonExpression)
			p.SetState(97)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(98)
				p.Match(STIXPatternParserOR)
			}
			{
				p.SetState(99)
				p.comparisonExpression(3)
			}

		}
		p.SetState(104)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
	}

	return localctx
}

// IComparisonExpressionAndContext is an interface to support dynamic dispatch.
type IComparisonExpressionAndContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsComparisonExpressionAndContext differentiates from other interfaces.
	IsComparisonExpressionAndContext()
}

type ComparisonExpressionAndContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparisonExpressionAndContext() *ComparisonExpressionAndContext {
	var p = new(ComparisonExpressionAndContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = STIXPatternParserRULE_comparisonExpressionAnd
	return p
}

func (*ComparisonExpressionAndContext) IsComparisonExpressionAndContext() {}

func NewComparisonExpressionAndContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComparisonExpressionAndContext {
	var p = new(ComparisonExpressionAndContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = STIXPatternParserRULE_comparisonExpressionAnd

	return p
}

func (s *ComparisonExpressionAndContext) GetParser() antlr.Parser { return s.parser }

func (s *ComparisonExpressionAndContext) PropTest() IPropTestContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPropTestContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPropTestContext)
}

func (s *ComparisonExpressionAndContext) AllComparisonExpressionAnd() []IComparisonExpressionAndContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IComparisonExpressionAndContext)(nil)).Elem())
	var tst = make([]IComparisonExpressionAndContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IComparisonExpressionAndContext)
		}
	}

	return tst
}

func (s *ComparisonExpressionAndContext) ComparisonExpressionAnd(i int) IComparisonExpressionAndContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComparisonExpressionAndContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IComparisonExpressionAndContext)
}

func (s *ComparisonExpressionAndContext) AND() antlr.TerminalNode {
	return s.GetToken(STIXPatternParserAND, 0)
}

func (s *ComparisonExpressionAndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparisonExpressionAndContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComparisonExpressionAndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STIXPatternListener); ok {
		listenerT.EnterComparisonExpressionAnd(s)
	}
}

func (s *ComparisonExpressionAndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STIXPatternListener); ok {
		listenerT.ExitComparisonExpressionAnd(s)
	}
}

func (p *STIXPatternParser) ComparisonExpressionAnd() (localctx IComparisonExpressionAndContext) {
	return p.comparisonExpressionAnd(0)
}

func (p *STIXPatternParser) comparisonExpressionAnd(_p int) (localctx IComparisonExpressionAndContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewComparisonExpressionAndContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IComparisonExpressionAndContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 12
	p.EnterRecursionRule(localctx, 12, STIXPatternParserRULE_comparisonExpressionAnd, _p)

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
		p.SetState(106)
		p.PropTest()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(113)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewComparisonExpressionAndContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, STIXPatternParserRULE_comparisonExpressionAnd)
			p.SetState(108)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(109)
				p.Match(STIXPatternParserAND)
			}
			{
				p.SetState(110)
				p.comparisonExpressionAnd(3)
			}

		}
		p.SetState(115)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())
	}

	return localctx
}

// IPropTestContext is an interface to support dynamic dispatch.
type IPropTestContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPropTestContext differentiates from other interfaces.
	IsPropTestContext()
}

type PropTestContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPropTestContext() *PropTestContext {
	var p = new(PropTestContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = STIXPatternParserRULE_propTest
	return p
}

func (*PropTestContext) IsPropTestContext() {}

func NewPropTestContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PropTestContext {
	var p = new(PropTestContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = STIXPatternParserRULE_propTest

	return p
}

func (s *PropTestContext) GetParser() antlr.Parser { return s.parser }

func (s *PropTestContext) CopyFrom(ctx *PropTestContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *PropTestContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropTestContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type PropTestExistsContext struct {
	*PropTestContext
}

func NewPropTestExistsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PropTestExistsContext {
	var p = new(PropTestExistsContext)

	p.PropTestContext = NewEmptyPropTestContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PropTestContext))

	return p
}

func (s *PropTestExistsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropTestExistsContext) EXISTS() antlr.TerminalNode {
	return s.GetToken(STIXPatternParserEXISTS, 0)
}

func (s *PropTestExistsContext) ObjectPath() IObjectPathContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObjectPathContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObjectPathContext)
}

func (s *PropTestExistsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STIXPatternListener); ok {
		listenerT.EnterPropTestExists(s)
	}
}

func (s *PropTestExistsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STIXPatternListener); ok {
		listenerT.ExitPropTestExists(s)
	}
}

type PropTestRegexContext struct {
	*PropTestContext
}

func NewPropTestRegexContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PropTestRegexContext {
	var p = new(PropTestRegexContext)

	p.PropTestContext = NewEmptyPropTestContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PropTestContext))

	return p
}

func (s *PropTestRegexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropTestRegexContext) ObjectPath() IObjectPathContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObjectPathContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObjectPathContext)
}

func (s *PropTestRegexContext) MATCHES() antlr.TerminalNode {
	return s.GetToken(STIXPatternParserMATCHES, 0)
}

func (s *PropTestRegexContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(STIXPatternParserStringLiteral, 0)
}

func (s *PropTestRegexContext) NOT() antlr.TerminalNode {
	return s.GetToken(STIXPatternParserNOT, 0)
}

func (s *PropTestRegexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STIXPatternListener); ok {
		listenerT.EnterPropTestRegex(s)
	}
}

func (s *PropTestRegexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STIXPatternListener); ok {
		listenerT.ExitPropTestRegex(s)
	}
}

type PropTestOrderContext struct {
	*PropTestContext
}

func NewPropTestOrderContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PropTestOrderContext {
	var p = new(PropTestOrderContext)

	p.PropTestContext = NewEmptyPropTestContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PropTestContext))

	return p
}

func (s *PropTestOrderContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropTestOrderContext) ObjectPath() IObjectPathContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObjectPathContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObjectPathContext)
}

func (s *PropTestOrderContext) OrderableLiteral() IOrderableLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOrderableLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOrderableLiteralContext)
}

func (s *PropTestOrderContext) GT() antlr.TerminalNode {
	return s.GetToken(STIXPatternParserGT, 0)
}

func (s *PropTestOrderContext) LT() antlr.TerminalNode {
	return s.GetToken(STIXPatternParserLT, 0)
}

func (s *PropTestOrderContext) GE() antlr.TerminalNode {
	return s.GetToken(STIXPatternParserGE, 0)
}

func (s *PropTestOrderContext) LE() antlr.TerminalNode {
	return s.GetToken(STIXPatternParserLE, 0)
}

func (s *PropTestOrderContext) NOT() antlr.TerminalNode {
	return s.GetToken(STIXPatternParserNOT, 0)
}

func (s *PropTestOrderContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STIXPatternListener); ok {
		listenerT.EnterPropTestOrder(s)
	}
}

func (s *PropTestOrderContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STIXPatternListener); ok {
		listenerT.ExitPropTestOrder(s)
	}
}

type PropTestLikeContext struct {
	*PropTestContext
}

func NewPropTestLikeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PropTestLikeContext {
	var p = new(PropTestLikeContext)

	p.PropTestContext = NewEmptyPropTestContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PropTestContext))

	return p
}

func (s *PropTestLikeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropTestLikeContext) ObjectPath() IObjectPathContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObjectPathContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObjectPathContext)
}

func (s *PropTestLikeContext) LIKE() antlr.TerminalNode {
	return s.GetToken(STIXPatternParserLIKE, 0)
}

func (s *PropTestLikeContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(STIXPatternParserStringLiteral, 0)
}

func (s *PropTestLikeContext) NOT() antlr.TerminalNode {
	return s.GetToken(STIXPatternParserNOT, 0)
}

func (s *PropTestLikeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STIXPatternListener); ok {
		listenerT.EnterPropTestLike(s)
	}
}

func (s *PropTestLikeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STIXPatternListener); ok {
		listenerT.ExitPropTestLike(s)
	}
}

type PropTestEqualContext struct {
	*PropTestContext
}

func NewPropTestEqualContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PropTestEqualContext {
	var p = new(PropTestEqualContext)

	p.PropTestContext = NewEmptyPropTestContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PropTestContext))

	return p
}

func (s *PropTestEqualContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropTestEqualContext) ObjectPath() IObjectPathContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObjectPathContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObjectPathContext)
}

func (s *PropTestEqualContext) PrimitiveLiteral() IPrimitiveLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimitiveLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimitiveLiteralContext)
}

func (s *PropTestEqualContext) EQ() antlr.TerminalNode {
	return s.GetToken(STIXPatternParserEQ, 0)
}

func (s *PropTestEqualContext) NEQ() antlr.TerminalNode {
	return s.GetToken(STIXPatternParserNEQ, 0)
}

func (s *PropTestEqualContext) NOT() antlr.TerminalNode {
	return s.GetToken(STIXPatternParserNOT, 0)
}

func (s *PropTestEqualContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STIXPatternListener); ok {
		listenerT.EnterPropTestEqual(s)
	}
}

func (s *PropTestEqualContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STIXPatternListener); ok {
		listenerT.ExitPropTestEqual(s)
	}
}

type PropTestSetContext struct {
	*PropTestContext
}

func NewPropTestSetContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PropTestSetContext {
	var p = new(PropTestSetContext)

	p.PropTestContext = NewEmptyPropTestContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PropTestContext))

	return p
}

func (s *PropTestSetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropTestSetContext) ObjectPath() IObjectPathContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObjectPathContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObjectPathContext)
}

func (s *PropTestSetContext) IN() antlr.TerminalNode {
	return s.GetToken(STIXPatternParserIN, 0)
}

func (s *PropTestSetContext) SetLiteral() ISetLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISetLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISetLiteralContext)
}

func (s *PropTestSetContext) NOT() antlr.TerminalNode {
	return s.GetToken(STIXPatternParserNOT, 0)
}

func (s *PropTestSetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STIXPatternListener); ok {
		listenerT.EnterPropTestSet(s)
	}
}

func (s *PropTestSetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STIXPatternListener); ok {
		listenerT.ExitPropTestSet(s)
	}
}

type PropTestIsSubsetContext struct {
	*PropTestContext
}

func NewPropTestIsSubsetContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PropTestIsSubsetContext {
	var p = new(PropTestIsSubsetContext)

	p.PropTestContext = NewEmptyPropTestContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PropTestContext))

	return p
}

func (s *PropTestIsSubsetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropTestIsSubsetContext) ObjectPath() IObjectPathContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObjectPathContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObjectPathContext)
}

func (s *PropTestIsSubsetContext) ISSUBSET() antlr.TerminalNode {
	return s.GetToken(STIXPatternParserISSUBSET, 0)
}

func (s *PropTestIsSubsetContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(STIXPatternParserStringLiteral, 0)
}

func (s *PropTestIsSubsetContext) NOT() antlr.TerminalNode {
	return s.GetToken(STIXPatternParserNOT, 0)
}

func (s *PropTestIsSubsetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STIXPatternListener); ok {
		listenerT.EnterPropTestIsSubset(s)
	}
}

func (s *PropTestIsSubsetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STIXPatternListener); ok {
		listenerT.ExitPropTestIsSubset(s)
	}
}

type PropTestParenContext struct {
	*PropTestContext
}

func NewPropTestParenContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PropTestParenContext {
	var p = new(PropTestParenContext)

	p.PropTestContext = NewEmptyPropTestContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PropTestContext))

	return p
}

func (s *PropTestParenContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropTestParenContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(STIXPatternParserLPAREN, 0)
}

func (s *PropTestParenContext) ComparisonExpression() IComparisonExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComparisonExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComparisonExpressionContext)
}

func (s *PropTestParenContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(STIXPatternParserRPAREN, 0)
}

func (s *PropTestParenContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STIXPatternListener); ok {
		listenerT.EnterPropTestParen(s)
	}
}

func (s *PropTestParenContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STIXPatternListener); ok {
		listenerT.ExitPropTestParen(s)
	}
}

type PropTestIsSupersetContext struct {
	*PropTestContext
}

func NewPropTestIsSupersetContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PropTestIsSupersetContext {
	var p = new(PropTestIsSupersetContext)

	p.PropTestContext = NewEmptyPropTestContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PropTestContext))

	return p
}

func (s *PropTestIsSupersetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropTestIsSupersetContext) ObjectPath() IObjectPathContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObjectPathContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObjectPathContext)
}

func (s *PropTestIsSupersetContext) ISSUPERSET() antlr.TerminalNode {
	return s.GetToken(STIXPatternParserISSUPERSET, 0)
}

func (s *PropTestIsSupersetContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(STIXPatternParserStringLiteral, 0)
}

func (s *PropTestIsSupersetContext) NOT() antlr.TerminalNode {
	return s.GetToken(STIXPatternParserNOT, 0)
}

func (s *PropTestIsSupersetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STIXPatternListener); ok {
		listenerT.EnterPropTestIsSuperset(s)
	}
}

func (s *PropTestIsSupersetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STIXPatternListener); ok {
		listenerT.ExitPropTestIsSuperset(s)
	}
}

func (p *STIXPatternParser) PropTest() (localctx IPropTestContext) {
	localctx = NewPropTestContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, STIXPatternParserRULE_propTest)
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

	p.SetState(171)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		localctx = NewPropTestEqualContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(116)
			p.ObjectPath()
		}
		p.SetState(118)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == STIXPatternParserNOT {
			{
				p.SetState(117)
				p.Match(STIXPatternParserNOT)
			}

		}
		{
			p.SetState(120)
			_la = p.GetTokenStream().LA(1)

			if !(_la == STIXPatternParserEQ || _la == STIXPatternParserNEQ) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(121)
			p.PrimitiveLiteral()
		}

	case 2:
		localctx = NewPropTestOrderContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(123)
			p.ObjectPath()
		}
		p.SetState(125)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == STIXPatternParserNOT {
			{
				p.SetState(124)
				p.Match(STIXPatternParserNOT)
			}

		}
		{
			p.SetState(127)
			_la = p.GetTokenStream().LA(1)

			if !(((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(STIXPatternParserLT-33))|(1<<(STIXPatternParserLE-33))|(1<<(STIXPatternParserGT-33))|(1<<(STIXPatternParserGE-33)))) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(128)
			p.OrderableLiteral()
		}

	case 3:
		localctx = NewPropTestSetContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(130)
			p.ObjectPath()
		}
		p.SetState(132)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == STIXPatternParserNOT {
			{
				p.SetState(131)
				p.Match(STIXPatternParserNOT)
			}

		}
		{
			p.SetState(134)
			p.Match(STIXPatternParserIN)
		}
		{
			p.SetState(135)
			p.SetLiteral()
		}

	case 4:
		localctx = NewPropTestLikeContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(137)
			p.ObjectPath()
		}
		p.SetState(139)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == STIXPatternParserNOT {
			{
				p.SetState(138)
				p.Match(STIXPatternParserNOT)
			}

		}
		{
			p.SetState(141)
			p.Match(STIXPatternParserLIKE)
		}
		{
			p.SetState(142)
			p.Match(STIXPatternParserStringLiteral)
		}

	case 5:
		localctx = NewPropTestRegexContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(144)
			p.ObjectPath()
		}
		p.SetState(146)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == STIXPatternParserNOT {
			{
				p.SetState(145)
				p.Match(STIXPatternParserNOT)
			}

		}
		{
			p.SetState(148)
			p.Match(STIXPatternParserMATCHES)
		}
		{
			p.SetState(149)
			p.Match(STIXPatternParserStringLiteral)
		}

	case 6:
		localctx = NewPropTestIsSubsetContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(151)
			p.ObjectPath()
		}
		p.SetState(153)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == STIXPatternParserNOT {
			{
				p.SetState(152)
				p.Match(STIXPatternParserNOT)
			}

		}
		{
			p.SetState(155)
			p.Match(STIXPatternParserISSUBSET)
		}
		{
			p.SetState(156)
			p.Match(STIXPatternParserStringLiteral)
		}

	case 7:
		localctx = NewPropTestIsSupersetContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(158)
			p.ObjectPath()
		}
		p.SetState(160)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == STIXPatternParserNOT {
			{
				p.SetState(159)
				p.Match(STIXPatternParserNOT)
			}

		}
		{
			p.SetState(162)
			p.Match(STIXPatternParserISSUPERSET)
		}
		{
			p.SetState(163)
			p.Match(STIXPatternParserStringLiteral)
		}

	case 8:
		localctx = NewPropTestParenContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(165)
			p.Match(STIXPatternParserLPAREN)
		}
		{
			p.SetState(166)
			p.comparisonExpression(0)
		}
		{
			p.SetState(167)
			p.Match(STIXPatternParserRPAREN)
		}

	case 9:
		localctx = NewPropTestExistsContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(169)
			p.Match(STIXPatternParserEXISTS)
		}
		{
			p.SetState(170)
			p.ObjectPath()
		}

	}

	return localctx
}

// IStartStopQualifierContext is an interface to support dynamic dispatch.
type IStartStopQualifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStartStopQualifierContext differentiates from other interfaces.
	IsStartStopQualifierContext()
}

type StartStopQualifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartStopQualifierContext() *StartStopQualifierContext {
	var p = new(StartStopQualifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = STIXPatternParserRULE_startStopQualifier
	return p
}

func (*StartStopQualifierContext) IsStartStopQualifierContext() {}

func NewStartStopQualifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartStopQualifierContext {
	var p = new(StartStopQualifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = STIXPatternParserRULE_startStopQualifier

	return p
}

func (s *StartStopQualifierContext) GetParser() antlr.Parser { return s.parser }

func (s *StartStopQualifierContext) START() antlr.TerminalNode {
	return s.GetToken(STIXPatternParserSTART, 0)
}

func (s *StartStopQualifierContext) AllTimestampLiteral() []antlr.TerminalNode {
	return s.GetTokens(STIXPatternParserTimestampLiteral)
}

func (s *StartStopQualifierContext) TimestampLiteral(i int) antlr.TerminalNode {
	return s.GetToken(STIXPatternParserTimestampLiteral, i)
}

func (s *StartStopQualifierContext) STOP() antlr.TerminalNode {
	return s.GetToken(STIXPatternParserSTOP, 0)
}

func (s *StartStopQualifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartStopQualifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartStopQualifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STIXPatternListener); ok {
		listenerT.EnterStartStopQualifier(s)
	}
}

func (s *StartStopQualifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STIXPatternListener); ok {
		listenerT.ExitStartStopQualifier(s)
	}
}

func (p *STIXPatternParser) StartStopQualifier() (localctx IStartStopQualifierContext) {
	localctx = NewStartStopQualifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, STIXPatternParserRULE_startStopQualifier)

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
		p.SetState(173)
		p.Match(STIXPatternParserSTART)
	}
	{
		p.SetState(174)
		p.Match(STIXPatternParserTimestampLiteral)
	}
	{
		p.SetState(175)
		p.Match(STIXPatternParserSTOP)
	}
	{
		p.SetState(176)
		p.Match(STIXPatternParserTimestampLiteral)
	}

	return localctx
}

// IWithinQualifierContext is an interface to support dynamic dispatch.
type IWithinQualifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWithinQualifierContext differentiates from other interfaces.
	IsWithinQualifierContext()
}

type WithinQualifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWithinQualifierContext() *WithinQualifierContext {
	var p = new(WithinQualifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = STIXPatternParserRULE_withinQualifier
	return p
}

func (*WithinQualifierContext) IsWithinQualifierContext() {}

func NewWithinQualifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WithinQualifierContext {
	var p = new(WithinQualifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = STIXPatternParserRULE_withinQualifier

	return p
}

func (s *WithinQualifierContext) GetParser() antlr.Parser { return s.parser }

func (s *WithinQualifierContext) WITHIN() antlr.TerminalNode {
	return s.GetToken(STIXPatternParserWITHIN, 0)
}

func (s *WithinQualifierContext) SECONDS() antlr.TerminalNode {
	return s.GetToken(STIXPatternParserSECONDS, 0)
}

func (s *WithinQualifierContext) IntPosLiteral() antlr.TerminalNode {
	return s.GetToken(STIXPatternParserIntPosLiteral, 0)
}

func (s *WithinQualifierContext) FloatPosLiteral() antlr.TerminalNode {
	return s.GetToken(STIXPatternParserFloatPosLiteral, 0)
}

func (s *WithinQualifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WithinQualifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WithinQualifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STIXPatternListener); ok {
		listenerT.EnterWithinQualifier(s)
	}
}

func (s *WithinQualifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STIXPatternListener); ok {
		listenerT.ExitWithinQualifier(s)
	}
}

func (p *STIXPatternParser) WithinQualifier() (localctx IWithinQualifierContext) {
	localctx = NewWithinQualifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, STIXPatternParserRULE_withinQualifier)
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
		p.SetState(178)
		p.Match(STIXPatternParserWITHIN)
	}
	{
		p.SetState(179)
		_la = p.GetTokenStream().LA(1)

		if !(_la == STIXPatternParserIntPosLiteral || _la == STIXPatternParserFloatPosLiteral) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(180)
		p.Match(STIXPatternParserSECONDS)
	}

	return localctx
}

// IRepeatedQualifierContext is an interface to support dynamic dispatch.
type IRepeatedQualifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRepeatedQualifierContext differentiates from other interfaces.
	IsRepeatedQualifierContext()
}

type RepeatedQualifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRepeatedQualifierContext() *RepeatedQualifierContext {
	var p = new(RepeatedQualifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = STIXPatternParserRULE_repeatedQualifier
	return p
}

func (*RepeatedQualifierContext) IsRepeatedQualifierContext() {}

func NewRepeatedQualifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RepeatedQualifierContext {
	var p = new(RepeatedQualifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = STIXPatternParserRULE_repeatedQualifier

	return p
}

func (s *RepeatedQualifierContext) GetParser() antlr.Parser { return s.parser }

func (s *RepeatedQualifierContext) REPEATS() antlr.TerminalNode {
	return s.GetToken(STIXPatternParserREPEATS, 0)
}

func (s *RepeatedQualifierContext) IntPosLiteral() antlr.TerminalNode {
	return s.GetToken(STIXPatternParserIntPosLiteral, 0)
}

func (s *RepeatedQualifierContext) TIMES() antlr.TerminalNode {
	return s.GetToken(STIXPatternParserTIMES, 0)
}

func (s *RepeatedQualifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RepeatedQualifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RepeatedQualifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STIXPatternListener); ok {
		listenerT.EnterRepeatedQualifier(s)
	}
}

func (s *RepeatedQualifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STIXPatternListener); ok {
		listenerT.ExitRepeatedQualifier(s)
	}
}

func (p *STIXPatternParser) RepeatedQualifier() (localctx IRepeatedQualifierContext) {
	localctx = NewRepeatedQualifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, STIXPatternParserRULE_repeatedQualifier)

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
		p.SetState(182)
		p.Match(STIXPatternParserREPEATS)
	}
	{
		p.SetState(183)
		p.Match(STIXPatternParserIntPosLiteral)
	}
	{
		p.SetState(184)
		p.Match(STIXPatternParserTIMES)
	}

	return localctx
}

// IObjectPathContext is an interface to support dynamic dispatch.
type IObjectPathContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObjectPathContext differentiates from other interfaces.
	IsObjectPathContext()
}

type ObjectPathContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjectPathContext() *ObjectPathContext {
	var p = new(ObjectPathContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = STIXPatternParserRULE_objectPath
	return p
}

func (*ObjectPathContext) IsObjectPathContext() {}

func NewObjectPathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjectPathContext {
	var p = new(ObjectPathContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = STIXPatternParserRULE_objectPath

	return p
}

func (s *ObjectPathContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjectPathContext) ObjectType() IObjectTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObjectTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObjectTypeContext)
}

func (s *ObjectPathContext) COLON() antlr.TerminalNode {
	return s.GetToken(STIXPatternParserCOLON, 0)
}

func (s *ObjectPathContext) FirstPathComponent() IFirstPathComponentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFirstPathComponentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFirstPathComponentContext)
}

func (s *ObjectPathContext) ObjectPathComponent() IObjectPathComponentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObjectPathComponentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObjectPathComponentContext)
}

func (s *ObjectPathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectPathContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjectPathContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STIXPatternListener); ok {
		listenerT.EnterObjectPath(s)
	}
}

func (s *ObjectPathContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STIXPatternListener); ok {
		listenerT.ExitObjectPath(s)
	}
}

func (p *STIXPatternParser) ObjectPath() (localctx IObjectPathContext) {
	localctx = NewObjectPathContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, STIXPatternParserRULE_objectPath)

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
		p.SetState(186)
		p.ObjectType()
	}
	{
		p.SetState(187)
		p.Match(STIXPatternParserCOLON)
	}
	{
		p.SetState(188)
		p.FirstPathComponent()
	}
	p.SetState(190)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(189)
			p.objectPathComponent(0)
		}

	}

	return localctx
}

// IObjectTypeContext is an interface to support dynamic dispatch.
type IObjectTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObjectTypeContext differentiates from other interfaces.
	IsObjectTypeContext()
}

type ObjectTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjectTypeContext() *ObjectTypeContext {
	var p = new(ObjectTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = STIXPatternParserRULE_objectType
	return p
}

func (*ObjectTypeContext) IsObjectTypeContext() {}

func NewObjectTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjectTypeContext {
	var p = new(ObjectTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = STIXPatternParserRULE_objectType

	return p
}

func (s *ObjectTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjectTypeContext) IdentifierWithoutHyphen() antlr.TerminalNode {
	return s.GetToken(STIXPatternParserIdentifierWithoutHyphen, 0)
}

func (s *ObjectTypeContext) IdentifierWithHyphen() antlr.TerminalNode {
	return s.GetToken(STIXPatternParserIdentifierWithHyphen, 0)
}

func (s *ObjectTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjectTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STIXPatternListener); ok {
		listenerT.EnterObjectType(s)
	}
}

func (s *ObjectTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STIXPatternListener); ok {
		listenerT.ExitObjectType(s)
	}
}

func (p *STIXPatternParser) ObjectType() (localctx IObjectTypeContext) {
	localctx = NewObjectTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, STIXPatternParserRULE_objectType)
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
		p.SetState(192)
		_la = p.GetTokenStream().LA(1)

		if !(_la == STIXPatternParserIdentifierWithoutHyphen || _la == STIXPatternParserIdentifierWithHyphen) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IFirstPathComponentContext is an interface to support dynamic dispatch.
type IFirstPathComponentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFirstPathComponentContext differentiates from other interfaces.
	IsFirstPathComponentContext()
}

type FirstPathComponentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFirstPathComponentContext() *FirstPathComponentContext {
	var p = new(FirstPathComponentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = STIXPatternParserRULE_firstPathComponent
	return p
}

func (*FirstPathComponentContext) IsFirstPathComponentContext() {}

func NewFirstPathComponentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FirstPathComponentContext {
	var p = new(FirstPathComponentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = STIXPatternParserRULE_firstPathComponent

	return p
}

func (s *FirstPathComponentContext) GetParser() antlr.Parser { return s.parser }

func (s *FirstPathComponentContext) IdentifierWithoutHyphen() antlr.TerminalNode {
	return s.GetToken(STIXPatternParserIdentifierWithoutHyphen, 0)
}

func (s *FirstPathComponentContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(STIXPatternParserStringLiteral, 0)
}

func (s *FirstPathComponentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FirstPathComponentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FirstPathComponentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STIXPatternListener); ok {
		listenerT.EnterFirstPathComponent(s)
	}
}

func (s *FirstPathComponentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STIXPatternListener); ok {
		listenerT.ExitFirstPathComponent(s)
	}
}

func (p *STIXPatternParser) FirstPathComponent() (localctx IFirstPathComponentContext) {
	localctx = NewFirstPathComponentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, STIXPatternParserRULE_firstPathComponent)
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
		p.SetState(194)
		_la = p.GetTokenStream().LA(1)

		if !(_la == STIXPatternParserStringLiteral || _la == STIXPatternParserIdentifierWithoutHyphen) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IObjectPathComponentContext is an interface to support dynamic dispatch.
type IObjectPathComponentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObjectPathComponentContext differentiates from other interfaces.
	IsObjectPathComponentContext()
}

type ObjectPathComponentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjectPathComponentContext() *ObjectPathComponentContext {
	var p = new(ObjectPathComponentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = STIXPatternParserRULE_objectPathComponent
	return p
}

func (*ObjectPathComponentContext) IsObjectPathComponentContext() {}

func NewObjectPathComponentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjectPathComponentContext {
	var p = new(ObjectPathComponentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = STIXPatternParserRULE_objectPathComponent

	return p
}

func (s *ObjectPathComponentContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjectPathComponentContext) CopyFrom(ctx *ObjectPathComponentContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ObjectPathComponentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectPathComponentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IndexPathStepContext struct {
	*ObjectPathComponentContext
}

func NewIndexPathStepContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IndexPathStepContext {
	var p = new(IndexPathStepContext)

	p.ObjectPathComponentContext = NewEmptyObjectPathComponentContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ObjectPathComponentContext))

	return p
}

func (s *IndexPathStepContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexPathStepContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(STIXPatternParserLBRACK, 0)
}

func (s *IndexPathStepContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(STIXPatternParserRBRACK, 0)
}

func (s *IndexPathStepContext) IntPosLiteral() antlr.TerminalNode {
	return s.GetToken(STIXPatternParserIntPosLiteral, 0)
}

func (s *IndexPathStepContext) IntNegLiteral() antlr.TerminalNode {
	return s.GetToken(STIXPatternParserIntNegLiteral, 0)
}

func (s *IndexPathStepContext) ASTERISK() antlr.TerminalNode {
	return s.GetToken(STIXPatternParserASTERISK, 0)
}

func (s *IndexPathStepContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STIXPatternListener); ok {
		listenerT.EnterIndexPathStep(s)
	}
}

func (s *IndexPathStepContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STIXPatternListener); ok {
		listenerT.ExitIndexPathStep(s)
	}
}

type PathStepContext struct {
	*ObjectPathComponentContext
}

func NewPathStepContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PathStepContext {
	var p = new(PathStepContext)

	p.ObjectPathComponentContext = NewEmptyObjectPathComponentContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ObjectPathComponentContext))

	return p
}

func (s *PathStepContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PathStepContext) AllObjectPathComponent() []IObjectPathComponentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IObjectPathComponentContext)(nil)).Elem())
	var tst = make([]IObjectPathComponentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IObjectPathComponentContext)
		}
	}

	return tst
}

func (s *PathStepContext) ObjectPathComponent(i int) IObjectPathComponentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObjectPathComponentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IObjectPathComponentContext)
}

func (s *PathStepContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STIXPatternListener); ok {
		listenerT.EnterPathStep(s)
	}
}

func (s *PathStepContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STIXPatternListener); ok {
		listenerT.ExitPathStep(s)
	}
}

type KeyPathStepContext struct {
	*ObjectPathComponentContext
}

func NewKeyPathStepContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *KeyPathStepContext {
	var p = new(KeyPathStepContext)

	p.ObjectPathComponentContext = NewEmptyObjectPathComponentContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ObjectPathComponentContext))

	return p
}

func (s *KeyPathStepContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeyPathStepContext) IdentifierWithoutHyphen() antlr.TerminalNode {
	return s.GetToken(STIXPatternParserIdentifierWithoutHyphen, 0)
}

func (s *KeyPathStepContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(STIXPatternParserStringLiteral, 0)
}

func (s *KeyPathStepContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STIXPatternListener); ok {
		listenerT.EnterKeyPathStep(s)
	}
}

func (s *KeyPathStepContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STIXPatternListener); ok {
		listenerT.ExitKeyPathStep(s)
	}
}

func (p *STIXPatternParser) ObjectPathComponent() (localctx IObjectPathComponentContext) {
	return p.objectPathComponent(0)
}

func (p *STIXPatternParser) objectPathComponent(_p int) (localctx IObjectPathComponentContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewObjectPathComponentContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IObjectPathComponentContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 28
	p.EnterRecursionRule(localctx, 28, STIXPatternParserRULE_objectPathComponent, _p)
	var _la int

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
	p.SetState(202)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case STIXPatternParserDOT:
		localctx = NewKeyPathStepContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(197)
			p.Match(STIXPatternParserDOT)
		}
		{
			p.SetState(198)
			_la = p.GetTokenStream().LA(1)

			if !(_la == STIXPatternParserStringLiteral || _la == STIXPatternParserIdentifierWithoutHyphen) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case STIXPatternParserLBRACK:
		localctx = NewIndexPathStepContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(199)
			p.Match(STIXPatternParserLBRACK)
		}
		{
			p.SetState(200)
			_la = p.GetTokenStream().LA(1)

			if !(_la == STIXPatternParserIntNegLiteral || _la == STIXPatternParserIntPosLiteral || _la == STIXPatternParserASTERISK) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(201)
			p.Match(STIXPatternParserRBRACK)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(208)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewPathStepContext(p, NewObjectPathComponentContext(p, _parentctx, _parentState))
			p.PushNewRecursionContext(localctx, _startState, STIXPatternParserRULE_objectPathComponent)
			p.SetState(204)

			if !(p.Precpred(p.GetParserRuleContext(), 3)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
			}
			{
				p.SetState(205)
				p.objectPathComponent(4)
			}

		}
		p.SetState(210)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext())
	}

	return localctx
}

// ISetLiteralContext is an interface to support dynamic dispatch.
type ISetLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSetLiteralContext differentiates from other interfaces.
	IsSetLiteralContext()
}

type SetLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySetLiteralContext() *SetLiteralContext {
	var p = new(SetLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = STIXPatternParserRULE_setLiteral
	return p
}

func (*SetLiteralContext) IsSetLiteralContext() {}

func NewSetLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SetLiteralContext {
	var p = new(SetLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = STIXPatternParserRULE_setLiteral

	return p
}

func (s *SetLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *SetLiteralContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(STIXPatternParserLPAREN, 0)
}

func (s *SetLiteralContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(STIXPatternParserRPAREN, 0)
}

func (s *SetLiteralContext) AllPrimitiveLiteral() []IPrimitiveLiteralContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPrimitiveLiteralContext)(nil)).Elem())
	var tst = make([]IPrimitiveLiteralContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPrimitiveLiteralContext)
		}
	}

	return tst
}

func (s *SetLiteralContext) PrimitiveLiteral(i int) IPrimitiveLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimitiveLiteralContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPrimitiveLiteralContext)
}

func (s *SetLiteralContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(STIXPatternParserCOMMA)
}

func (s *SetLiteralContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(STIXPatternParserCOMMA, i)
}

func (s *SetLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SetLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SetLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STIXPatternListener); ok {
		listenerT.EnterSetLiteral(s)
	}
}

func (s *SetLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STIXPatternListener); ok {
		listenerT.ExitSetLiteral(s)
	}
}

func (p *STIXPatternParser) SetLiteral() (localctx ISetLiteralContext) {
	localctx = NewSetLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, STIXPatternParserRULE_setLiteral)
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

	p.SetState(224)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(211)
			p.Match(STIXPatternParserLPAREN)
		}
		{
			p.SetState(212)
			p.Match(STIXPatternParserRPAREN)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(213)
			p.Match(STIXPatternParserLPAREN)
		}
		{
			p.SetState(214)
			p.PrimitiveLiteral()
		}
		p.SetState(219)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == STIXPatternParserCOMMA {
			{
				p.SetState(215)
				p.Match(STIXPatternParserCOMMA)
			}
			{
				p.SetState(216)
				p.PrimitiveLiteral()
			}

			p.SetState(221)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(222)
			p.Match(STIXPatternParserRPAREN)
		}

	}

	return localctx
}

// IPrimitiveLiteralContext is an interface to support dynamic dispatch.
type IPrimitiveLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrimitiveLiteralContext differentiates from other interfaces.
	IsPrimitiveLiteralContext()
}

type PrimitiveLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimitiveLiteralContext() *PrimitiveLiteralContext {
	var p = new(PrimitiveLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = STIXPatternParserRULE_primitiveLiteral
	return p
}

func (*PrimitiveLiteralContext) IsPrimitiveLiteralContext() {}

func NewPrimitiveLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimitiveLiteralContext {
	var p = new(PrimitiveLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = STIXPatternParserRULE_primitiveLiteral

	return p
}

func (s *PrimitiveLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimitiveLiteralContext) OrderableLiteral() IOrderableLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOrderableLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOrderableLiteralContext)
}

func (s *PrimitiveLiteralContext) BoolLiteral() antlr.TerminalNode {
	return s.GetToken(STIXPatternParserBoolLiteral, 0)
}

func (s *PrimitiveLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimitiveLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimitiveLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STIXPatternListener); ok {
		listenerT.EnterPrimitiveLiteral(s)
	}
}

func (s *PrimitiveLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STIXPatternListener); ok {
		listenerT.ExitPrimitiveLiteral(s)
	}
}

func (p *STIXPatternParser) PrimitiveLiteral() (localctx IPrimitiveLiteralContext) {
	localctx = NewPrimitiveLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, STIXPatternParserRULE_primitiveLiteral)

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

	p.SetState(228)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case STIXPatternParserIntNegLiteral, STIXPatternParserIntPosLiteral, STIXPatternParserFloatNegLiteral, STIXPatternParserFloatPosLiteral, STIXPatternParserHexLiteral, STIXPatternParserBinaryLiteral, STIXPatternParserStringLiteral, STIXPatternParserTimestampLiteral:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(226)
			p.OrderableLiteral()
		}

	case STIXPatternParserBoolLiteral:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(227)
			p.Match(STIXPatternParserBoolLiteral)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IOrderableLiteralContext is an interface to support dynamic dispatch.
type IOrderableLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOrderableLiteralContext differentiates from other interfaces.
	IsOrderableLiteralContext()
}

type OrderableLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOrderableLiteralContext() *OrderableLiteralContext {
	var p = new(OrderableLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = STIXPatternParserRULE_orderableLiteral
	return p
}

func (*OrderableLiteralContext) IsOrderableLiteralContext() {}

func NewOrderableLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrderableLiteralContext {
	var p = new(OrderableLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = STIXPatternParserRULE_orderableLiteral

	return p
}

func (s *OrderableLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *OrderableLiteralContext) IntPosLiteral() antlr.TerminalNode {
	return s.GetToken(STIXPatternParserIntPosLiteral, 0)
}

func (s *OrderableLiteralContext) IntNegLiteral() antlr.TerminalNode {
	return s.GetToken(STIXPatternParserIntNegLiteral, 0)
}

func (s *OrderableLiteralContext) FloatPosLiteral() antlr.TerminalNode {
	return s.GetToken(STIXPatternParserFloatPosLiteral, 0)
}

func (s *OrderableLiteralContext) FloatNegLiteral() antlr.TerminalNode {
	return s.GetToken(STIXPatternParserFloatNegLiteral, 0)
}

func (s *OrderableLiteralContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(STIXPatternParserStringLiteral, 0)
}

func (s *OrderableLiteralContext) BinaryLiteral() antlr.TerminalNode {
	return s.GetToken(STIXPatternParserBinaryLiteral, 0)
}

func (s *OrderableLiteralContext) HexLiteral() antlr.TerminalNode {
	return s.GetToken(STIXPatternParserHexLiteral, 0)
}

func (s *OrderableLiteralContext) TimestampLiteral() antlr.TerminalNode {
	return s.GetToken(STIXPatternParserTimestampLiteral, 0)
}

func (s *OrderableLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrderableLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OrderableLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STIXPatternListener); ok {
		listenerT.EnterOrderableLiteral(s)
	}
}

func (s *OrderableLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STIXPatternListener); ok {
		listenerT.ExitOrderableLiteral(s)
	}
}

func (p *STIXPatternParser) OrderableLiteral() (localctx IOrderableLiteralContext) {
	localctx = NewOrderableLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, STIXPatternParserRULE_orderableLiteral)
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
		p.SetState(230)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<STIXPatternParserIntNegLiteral)|(1<<STIXPatternParserIntPosLiteral)|(1<<STIXPatternParserFloatNegLiteral)|(1<<STIXPatternParserFloatPosLiteral)|(1<<STIXPatternParserHexLiteral)|(1<<STIXPatternParserBinaryLiteral)|(1<<STIXPatternParserStringLiteral)|(1<<STIXPatternParserTimestampLiteral))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

func (p *STIXPatternParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *ObservationExpressionsContext = nil
		if localctx != nil {
			t = localctx.(*ObservationExpressionsContext)
		}
		return p.ObservationExpressions_Sempred(t, predIndex)

	case 2:
		var t *ObservationExpressionOrContext = nil
		if localctx != nil {
			t = localctx.(*ObservationExpressionOrContext)
		}
		return p.ObservationExpressionOr_Sempred(t, predIndex)

	case 3:
		var t *ObservationExpressionAndContext = nil
		if localctx != nil {
			t = localctx.(*ObservationExpressionAndContext)
		}
		return p.ObservationExpressionAnd_Sempred(t, predIndex)

	case 4:
		var t *ObservationExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ObservationExpressionContext)
		}
		return p.ObservationExpression_Sempred(t, predIndex)

	case 5:
		var t *ComparisonExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ComparisonExpressionContext)
		}
		return p.ComparisonExpression_Sempred(t, predIndex)

	case 6:
		var t *ComparisonExpressionAndContext = nil
		if localctx != nil {
			t = localctx.(*ComparisonExpressionAndContext)
		}
		return p.ComparisonExpressionAnd_Sempred(t, predIndex)

	case 14:
		var t *ObjectPathComponentContext = nil
		if localctx != nil {
			t = localctx.(*ObjectPathComponentContext)
		}
		return p.ObjectPathComponent_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *STIXPatternParser) ObservationExpressions_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *STIXPatternParser) ObservationExpressionOr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *STIXPatternParser) ObservationExpressionAnd_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *STIXPatternParser) ObservationExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 3:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *STIXPatternParser) ComparisonExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 6:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *STIXPatternParser) ComparisonExpressionAnd_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 7:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *STIXPatternParser) ObjectPathComponent_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 8:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
