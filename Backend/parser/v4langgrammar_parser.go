// Code generated from V4LangGrammar.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // V4LangGrammar

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type V4LangGrammarParser struct {
	*antlr.BaseParser
}

var V4LangGrammarParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func v4langgrammarParserInit() {
	staticData := &V4LangGrammarParserStaticData
	staticData.LiteralNames = []string{
		"", "','", "'type'", "'switch'", "'continue'", "'return'", "'case'",
		"':'", "'default'", "'-'", "'!'", "'*'", "'/'", "'+'", "'<'", "'>'",
		"'<='", "'>='", "'%'", "'=='", "'!='", "'+='", "'-='", "'*='", "'/='",
		"'&&'", "'||'", "'new'", "'++'", "'--'", "'.'", "'int'", "'bool'", "'float64'",
		"'string'", "'rune'", "'mut'", "'fmt.Println'", "'if'", "'else'", "'append'",
		"'len'", "'slices.Index'", "'strings.Join'", "'strconv.Atoi'", "'strconv.ParseFloat'",
		"'reflect.TypeOf'", "'for'", "'range'", "'break'", "'nil'", "'func'",
		"'struct'", "'{'", "'}'", "'('", "')'", "'['", "']'", "':='", "'='",
		"';'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "MUT", "PRINT", "IF", "ELSE", "APPEND", "LEN", "SLICE_INDEX",
		"JOIN", "ATOI", "PARSEFLOAT", "TYPEOF", "FOR", "RANGE", "BREAK", "NIL",
		"FUNC", "STRUCT", "L_CURLY", "R_CURLY", "L_PAREN", "R_PAREN", "L_BRACKET",
		"R_BRACKET", "DECLARE_ASSIGN", "ASSIGN", "SEMICOLON", "INT", "BOOL",
		"FLOAT", "STRING", "RUNE", "WS", "ID", "COMMENT", "MULTILINE_COMMENT",
	}
	staticData.RuleNames = []string{
		"program", "dcl", "varDcl", "funcDcl", "params", "structDcl", "structBody",
		"receiver", "sliceDcl", "sliceValues", "matrixDcl", "matrixValues",
		"statement", "elseIfStmt", "elseStmt", "switchCase", "defaultCase",
		"args", "expr", "call", "instanceValues", "type",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 70, 509, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 1, 0, 5, 0, 46, 8, 0, 10, 0, 12, 0, 49, 9, 0, 1, 1, 1, 1, 3,
		1, 53, 8, 1, 1, 1, 1, 1, 3, 1, 57, 8, 1, 1, 1, 1, 1, 3, 1, 61, 8, 1, 1,
		1, 1, 1, 3, 1, 65, 8, 1, 1, 1, 1, 1, 3, 1, 69, 8, 1, 1, 1, 1, 1, 3, 1,
		73, 8, 1, 3, 1, 75, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 3, 2, 98, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 104, 8, 3, 1, 3,
		1, 3, 3, 3, 108, 8, 3, 1, 3, 1, 3, 5, 3, 112, 8, 3, 10, 3, 12, 3, 115,
		9, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 123, 8, 3, 1, 3, 1, 3,
		3, 3, 127, 8, 3, 1, 3, 1, 3, 5, 3, 131, 8, 3, 10, 3, 12, 3, 134, 9, 3,
		1, 3, 1, 3, 3, 3, 138, 8, 3, 1, 4, 1, 4, 1, 4, 3, 4, 143, 8, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 3, 4, 150, 8, 4, 1, 4, 5, 4, 153, 8, 4, 10, 4, 12,
		4, 156, 9, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 5, 5, 163, 8, 5, 10, 5, 12,
		5, 166, 9, 5, 1, 5, 1, 5, 1, 6, 1, 6, 3, 6, 172, 8, 6, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 3, 8, 191, 8, 8, 1, 9, 1, 9, 1, 9, 5, 9, 196, 8, 9, 10, 9,
		12, 9, 199, 9, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1,
		11, 5, 11, 220, 8, 11, 10, 11, 12, 11, 223, 9, 11, 1, 11, 3, 11, 226, 8,
		11, 1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 232, 8, 12, 1, 12, 1, 12, 1, 12,
		5, 12, 237, 8, 12, 10, 12, 12, 12, 240, 9, 12, 1, 12, 1, 12, 1, 12, 1,
		12, 1, 12, 5, 12, 247, 8, 12, 10, 12, 12, 12, 250, 9, 12, 1, 12, 3, 12,
		253, 8, 12, 1, 12, 1, 12, 1, 12, 1, 12, 5, 12, 259, 8, 12, 10, 12, 12,
		12, 262, 9, 12, 1, 12, 3, 12, 265, 8, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1,
		12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12,
		1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1,
		12, 3, 12, 293, 8, 12, 3, 12, 295, 8, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1,
		13, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 5, 15, 309, 8, 15,
		10, 15, 12, 15, 312, 9, 15, 1, 16, 1, 16, 1, 16, 5, 16, 317, 8, 16, 10,
		16, 12, 16, 320, 9, 16, 1, 17, 1, 17, 1, 17, 5, 17, 325, 8, 17, 10, 17,
		12, 17, 328, 9, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1,
		18, 1, 18, 1, 18, 1, 18, 1, 18, 3, 18, 342, 8, 18, 1, 18, 1, 18, 1, 18,
		1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1,
		18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 3, 18,
		367, 8, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1,
		18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18,
		1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1,
		18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18,
		1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1,
		18, 1, 18, 1, 18, 1, 18, 1, 18, 5, 18, 424, 8, 18, 10, 18, 12, 18, 427,
		9, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1,
		18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18,
		3, 18, 449, 8, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1,
		18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18,
		1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 4, 18, 474, 8, 18, 11, 18, 12, 18, 475,
		5, 18, 478, 8, 18, 10, 18, 12, 18, 481, 9, 18, 1, 19, 1, 19, 3, 19, 485,
		8, 19, 1, 19, 1, 19, 1, 19, 3, 19, 490, 8, 19, 1, 20, 1, 20, 1, 20, 1,
		20, 1, 20, 1, 20, 1, 20, 5, 20, 499, 8, 20, 10, 20, 12, 20, 502, 9, 20,
		1, 20, 3, 20, 505, 8, 20, 1, 21, 1, 21, 1, 21, 0, 1, 36, 22, 0, 2, 4, 6,
		8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42,
		0, 8, 1, 0, 21, 24, 1, 0, 28, 29, 1, 0, 11, 12, 2, 0, 9, 9, 13, 13, 1,
		0, 14, 17, 1, 0, 19, 20, 1, 0, 25, 26, 1, 0, 31, 35, 581, 0, 47, 1, 0,
		0, 0, 2, 74, 1, 0, 0, 0, 4, 97, 1, 0, 0, 0, 6, 137, 1, 0, 0, 0, 8, 139,
		1, 0, 0, 0, 10, 157, 1, 0, 0, 0, 12, 171, 1, 0, 0, 0, 14, 173, 1, 0, 0,
		0, 16, 190, 1, 0, 0, 0, 18, 192, 1, 0, 0, 0, 20, 200, 1, 0, 0, 0, 22, 211,
		1, 0, 0, 0, 24, 294, 1, 0, 0, 0, 26, 296, 1, 0, 0, 0, 28, 301, 1, 0, 0,
		0, 30, 304, 1, 0, 0, 0, 32, 313, 1, 0, 0, 0, 34, 321, 1, 0, 0, 0, 36, 448,
		1, 0, 0, 0, 38, 489, 1, 0, 0, 0, 40, 491, 1, 0, 0, 0, 42, 506, 1, 0, 0,
		0, 44, 46, 3, 2, 1, 0, 45, 44, 1, 0, 0, 0, 46, 49, 1, 0, 0, 0, 47, 45,
		1, 0, 0, 0, 47, 48, 1, 0, 0, 0, 48, 1, 1, 0, 0, 0, 49, 47, 1, 0, 0, 0,
		50, 52, 3, 4, 2, 0, 51, 53, 5, 61, 0, 0, 52, 51, 1, 0, 0, 0, 52, 53, 1,
		0, 0, 0, 53, 75, 1, 0, 0, 0, 54, 56, 3, 20, 10, 0, 55, 57, 5, 61, 0, 0,
		56, 55, 1, 0, 0, 0, 56, 57, 1, 0, 0, 0, 57, 75, 1, 0, 0, 0, 58, 60, 3,
		16, 8, 0, 59, 61, 5, 61, 0, 0, 60, 59, 1, 0, 0, 0, 60, 61, 1, 0, 0, 0,
		61, 75, 1, 0, 0, 0, 62, 64, 3, 24, 12, 0, 63, 65, 5, 61, 0, 0, 64, 63,
		1, 0, 0, 0, 64, 65, 1, 0, 0, 0, 65, 75, 1, 0, 0, 0, 66, 68, 3, 6, 3, 0,
		67, 69, 5, 61, 0, 0, 68, 67, 1, 0, 0, 0, 68, 69, 1, 0, 0, 0, 69, 75, 1,
		0, 0, 0, 70, 72, 3, 10, 5, 0, 71, 73, 5, 61, 0, 0, 72, 71, 1, 0, 0, 0,
		72, 73, 1, 0, 0, 0, 73, 75, 1, 0, 0, 0, 74, 50, 1, 0, 0, 0, 74, 54, 1,
		0, 0, 0, 74, 58, 1, 0, 0, 0, 74, 62, 1, 0, 0, 0, 74, 66, 1, 0, 0, 0, 74,
		70, 1, 0, 0, 0, 75, 3, 1, 0, 0, 0, 76, 77, 5, 36, 0, 0, 77, 78, 5, 68,
		0, 0, 78, 79, 3, 42, 21, 0, 79, 80, 5, 60, 0, 0, 80, 81, 3, 36, 18, 0,
		81, 98, 1, 0, 0, 0, 82, 83, 5, 36, 0, 0, 83, 84, 5, 68, 0, 0, 84, 98, 3,
		42, 21, 0, 85, 86, 5, 36, 0, 0, 86, 87, 5, 68, 0, 0, 87, 88, 5, 53, 0,
		0, 88, 89, 5, 54, 0, 0, 89, 98, 3, 42, 21, 0, 90, 91, 5, 68, 0, 0, 91,
		92, 5, 59, 0, 0, 92, 98, 3, 36, 18, 0, 93, 94, 5, 68, 0, 0, 94, 98, 3,
		42, 21, 0, 95, 96, 5, 68, 0, 0, 96, 98, 5, 68, 0, 0, 97, 76, 1, 0, 0, 0,
		97, 82, 1, 0, 0, 0, 97, 85, 1, 0, 0, 0, 97, 90, 1, 0, 0, 0, 97, 93, 1,
		0, 0, 0, 97, 95, 1, 0, 0, 0, 98, 5, 1, 0, 0, 0, 99, 100, 5, 51, 0, 0, 100,
		101, 5, 68, 0, 0, 101, 103, 5, 55, 0, 0, 102, 104, 3, 8, 4, 0, 103, 102,
		1, 0, 0, 0, 103, 104, 1, 0, 0, 0, 104, 105, 1, 0, 0, 0, 105, 107, 5, 56,
		0, 0, 106, 108, 3, 42, 21, 0, 107, 106, 1, 0, 0, 0, 107, 108, 1, 0, 0,
		0, 108, 109, 1, 0, 0, 0, 109, 113, 5, 53, 0, 0, 110, 112, 3, 2, 1, 0, 111,
		110, 1, 0, 0, 0, 112, 115, 1, 0, 0, 0, 113, 111, 1, 0, 0, 0, 113, 114,
		1, 0, 0, 0, 114, 116, 1, 0, 0, 0, 115, 113, 1, 0, 0, 0, 116, 138, 5, 54,
		0, 0, 117, 118, 5, 51, 0, 0, 118, 119, 3, 14, 7, 0, 119, 120, 5, 68, 0,
		0, 120, 122, 5, 55, 0, 0, 121, 123, 3, 8, 4, 0, 122, 121, 1, 0, 0, 0, 122,
		123, 1, 0, 0, 0, 123, 124, 1, 0, 0, 0, 124, 126, 5, 56, 0, 0, 125, 127,
		3, 42, 21, 0, 126, 125, 1, 0, 0, 0, 126, 127, 1, 0, 0, 0, 127, 128, 1,
		0, 0, 0, 128, 132, 5, 53, 0, 0, 129, 131, 3, 2, 1, 0, 130, 129, 1, 0, 0,
		0, 131, 134, 1, 0, 0, 0, 132, 130, 1, 0, 0, 0, 132, 133, 1, 0, 0, 0, 133,
		135, 1, 0, 0, 0, 134, 132, 1, 0, 0, 0, 135, 136, 5, 54, 0, 0, 136, 138,
		1, 0, 0, 0, 137, 99, 1, 0, 0, 0, 137, 117, 1, 0, 0, 0, 138, 7, 1, 0, 0,
		0, 139, 142, 5, 68, 0, 0, 140, 141, 5, 57, 0, 0, 141, 143, 5, 58, 0, 0,
		142, 140, 1, 0, 0, 0, 142, 143, 1, 0, 0, 0, 143, 144, 1, 0, 0, 0, 144,
		154, 3, 42, 21, 0, 145, 146, 5, 1, 0, 0, 146, 149, 5, 68, 0, 0, 147, 148,
		5, 57, 0, 0, 148, 150, 5, 58, 0, 0, 149, 147, 1, 0, 0, 0, 149, 150, 1,
		0, 0, 0, 150, 151, 1, 0, 0, 0, 151, 153, 3, 42, 21, 0, 152, 145, 1, 0,
		0, 0, 153, 156, 1, 0, 0, 0, 154, 152, 1, 0, 0, 0, 154, 155, 1, 0, 0, 0,
		155, 9, 1, 0, 0, 0, 156, 154, 1, 0, 0, 0, 157, 158, 5, 2, 0, 0, 158, 159,
		5, 68, 0, 0, 159, 160, 5, 52, 0, 0, 160, 164, 5, 53, 0, 0, 161, 163, 3,
		12, 6, 0, 162, 161, 1, 0, 0, 0, 163, 166, 1, 0, 0, 0, 164, 162, 1, 0, 0,
		0, 164, 165, 1, 0, 0, 0, 165, 167, 1, 0, 0, 0, 166, 164, 1, 0, 0, 0, 167,
		168, 5, 54, 0, 0, 168, 11, 1, 0, 0, 0, 169, 172, 3, 4, 2, 0, 170, 172,
		3, 6, 3, 0, 171, 169, 1, 0, 0, 0, 171, 170, 1, 0, 0, 0, 172, 13, 1, 0,
		0, 0, 173, 174, 5, 55, 0, 0, 174, 175, 5, 68, 0, 0, 175, 176, 5, 68, 0,
		0, 176, 177, 5, 56, 0, 0, 177, 15, 1, 0, 0, 0, 178, 179, 5, 36, 0, 0, 179,
		180, 5, 68, 0, 0, 180, 181, 5, 59, 0, 0, 181, 182, 5, 57, 0, 0, 182, 183,
		5, 58, 0, 0, 183, 184, 3, 42, 21, 0, 184, 185, 5, 53, 0, 0, 185, 186, 3,
		18, 9, 0, 186, 187, 5, 54, 0, 0, 187, 191, 1, 0, 0, 0, 188, 189, 5, 68,
		0, 0, 189, 191, 3, 42, 21, 0, 190, 178, 1, 0, 0, 0, 190, 188, 1, 0, 0,
		0, 191, 17, 1, 0, 0, 0, 192, 197, 3, 36, 18, 0, 193, 194, 5, 1, 0, 0, 194,
		196, 3, 36, 18, 0, 195, 193, 1, 0, 0, 0, 196, 199, 1, 0, 0, 0, 197, 195,
		1, 0, 0, 0, 197, 198, 1, 0, 0, 0, 198, 19, 1, 0, 0, 0, 199, 197, 1, 0,
		0, 0, 200, 201, 5, 68, 0, 0, 201, 202, 5, 59, 0, 0, 202, 203, 5, 57, 0,
		0, 203, 204, 5, 58, 0, 0, 204, 205, 5, 57, 0, 0, 205, 206, 5, 58, 0, 0,
		206, 207, 3, 42, 21, 0, 207, 208, 5, 53, 0, 0, 208, 209, 3, 22, 11, 0,
		209, 210, 5, 54, 0, 0, 210, 21, 1, 0, 0, 0, 211, 212, 5, 53, 0, 0, 212,
		213, 3, 18, 9, 0, 213, 221, 5, 54, 0, 0, 214, 215, 5, 1, 0, 0, 215, 216,
		5, 53, 0, 0, 216, 217, 3, 18, 9, 0, 217, 218, 5, 54, 0, 0, 218, 220, 1,
		0, 0, 0, 219, 214, 1, 0, 0, 0, 220, 223, 1, 0, 0, 0, 221, 219, 1, 0, 0,
		0, 221, 222, 1, 0, 0, 0, 222, 225, 1, 0, 0, 0, 223, 221, 1, 0, 0, 0, 224,
		226, 5, 1, 0, 0, 225, 224, 1, 0, 0, 0, 225, 226, 1, 0, 0, 0, 226, 23, 1,
		0, 0, 0, 227, 295, 3, 36, 18, 0, 228, 229, 5, 37, 0, 0, 229, 231, 5, 55,
		0, 0, 230, 232, 3, 34, 17, 0, 231, 230, 1, 0, 0, 0, 231, 232, 1, 0, 0,
		0, 232, 233, 1, 0, 0, 0, 233, 295, 5, 56, 0, 0, 234, 238, 5, 53, 0, 0,
		235, 237, 3, 2, 1, 0, 236, 235, 1, 0, 0, 0, 237, 240, 1, 0, 0, 0, 238,
		236, 1, 0, 0, 0, 238, 239, 1, 0, 0, 0, 239, 241, 1, 0, 0, 0, 240, 238,
		1, 0, 0, 0, 241, 295, 5, 54, 0, 0, 242, 243, 5, 38, 0, 0, 243, 244, 3,
		36, 18, 0, 244, 248, 3, 24, 12, 0, 245, 247, 3, 26, 13, 0, 246, 245, 1,
		0, 0, 0, 247, 250, 1, 0, 0, 0, 248, 246, 1, 0, 0, 0, 248, 249, 1, 0, 0,
		0, 249, 252, 1, 0, 0, 0, 250, 248, 1, 0, 0, 0, 251, 253, 3, 28, 14, 0,
		252, 251, 1, 0, 0, 0, 252, 253, 1, 0, 0, 0, 253, 295, 1, 0, 0, 0, 254,
		255, 5, 3, 0, 0, 255, 256, 3, 36, 18, 0, 256, 260, 5, 53, 0, 0, 257, 259,
		3, 30, 15, 0, 258, 257, 1, 0, 0, 0, 259, 262, 1, 0, 0, 0, 260, 258, 1,
		0, 0, 0, 260, 261, 1, 0, 0, 0, 261, 264, 1, 0, 0, 0, 262, 260, 1, 0, 0,
		0, 263, 265, 3, 32, 16, 0, 264, 263, 1, 0, 0, 0, 264, 265, 1, 0, 0, 0,
		265, 266, 1, 0, 0, 0, 266, 267, 5, 54, 0, 0, 267, 295, 1, 0, 0, 0, 268,
		269, 5, 47, 0, 0, 269, 270, 3, 4, 2, 0, 270, 271, 5, 61, 0, 0, 271, 272,
		3, 36, 18, 0, 272, 273, 5, 61, 0, 0, 273, 274, 3, 36, 18, 0, 274, 275,
		3, 24, 12, 0, 275, 295, 1, 0, 0, 0, 276, 277, 5, 47, 0, 0, 277, 278, 3,
		36, 18, 0, 278, 279, 3, 24, 12, 0, 279, 295, 1, 0, 0, 0, 280, 281, 5, 47,
		0, 0, 281, 282, 5, 68, 0, 0, 282, 283, 5, 1, 0, 0, 283, 284, 5, 68, 0,
		0, 284, 285, 5, 59, 0, 0, 285, 286, 5, 48, 0, 0, 286, 287, 5, 68, 0, 0,
		287, 295, 3, 24, 12, 0, 288, 295, 5, 49, 0, 0, 289, 295, 5, 4, 0, 0, 290,
		292, 5, 5, 0, 0, 291, 293, 3, 36, 18, 0, 292, 291, 1, 0, 0, 0, 292, 293,
		1, 0, 0, 0, 293, 295, 1, 0, 0, 0, 294, 227, 1, 0, 0, 0, 294, 228, 1, 0,
		0, 0, 294, 234, 1, 0, 0, 0, 294, 242, 1, 0, 0, 0, 294, 254, 1, 0, 0, 0,
		294, 268, 1, 0, 0, 0, 294, 276, 1, 0, 0, 0, 294, 280, 1, 0, 0, 0, 294,
		288, 1, 0, 0, 0, 294, 289, 1, 0, 0, 0, 294, 290, 1, 0, 0, 0, 295, 25, 1,
		0, 0, 0, 296, 297, 5, 39, 0, 0, 297, 298, 5, 38, 0, 0, 298, 299, 3, 36,
		18, 0, 299, 300, 3, 24, 12, 0, 300, 27, 1, 0, 0, 0, 301, 302, 5, 39, 0,
		0, 302, 303, 3, 24, 12, 0, 303, 29, 1, 0, 0, 0, 304, 305, 5, 6, 0, 0, 305,
		306, 3, 36, 18, 0, 306, 310, 5, 7, 0, 0, 307, 309, 3, 2, 1, 0, 308, 307,
		1, 0, 0, 0, 309, 312, 1, 0, 0, 0, 310, 308, 1, 0, 0, 0, 310, 311, 1, 0,
		0, 0, 311, 31, 1, 0, 0, 0, 312, 310, 1, 0, 0, 0, 313, 314, 5, 8, 0, 0,
		314, 318, 5, 7, 0, 0, 315, 317, 3, 2, 1, 0, 316, 315, 1, 0, 0, 0, 317,
		320, 1, 0, 0, 0, 318, 316, 1, 0, 0, 0, 318, 319, 1, 0, 0, 0, 319, 33, 1,
		0, 0, 0, 320, 318, 1, 0, 0, 0, 321, 326, 3, 36, 18, 0, 322, 323, 5, 1,
		0, 0, 323, 325, 3, 36, 18, 0, 324, 322, 1, 0, 0, 0, 325, 328, 1, 0, 0,
		0, 326, 324, 1, 0, 0, 0, 326, 327, 1, 0, 0, 0, 327, 35, 1, 0, 0, 0, 328,
		326, 1, 0, 0, 0, 329, 330, 6, 18, -1, 0, 330, 331, 5, 9, 0, 0, 331, 449,
		3, 36, 18, 35, 332, 333, 5, 10, 0, 0, 333, 449, 3, 36, 18, 34, 334, 335,
		5, 68, 0, 0, 335, 336, 7, 0, 0, 0, 336, 449, 3, 36, 18, 27, 337, 338, 5,
		27, 0, 0, 338, 339, 5, 68, 0, 0, 339, 341, 5, 55, 0, 0, 340, 342, 3, 34,
		17, 0, 341, 340, 1, 0, 0, 0, 341, 342, 1, 0, 0, 0, 342, 343, 1, 0, 0, 0,
		343, 449, 5, 56, 0, 0, 344, 345, 5, 68, 0, 0, 345, 346, 5, 53, 0, 0, 346,
		347, 3, 40, 20, 0, 347, 348, 5, 54, 0, 0, 348, 449, 1, 0, 0, 0, 349, 350,
		5, 68, 0, 0, 350, 351, 5, 57, 0, 0, 351, 352, 3, 36, 18, 0, 352, 353, 5,
		58, 0, 0, 353, 354, 5, 57, 0, 0, 354, 355, 3, 36, 18, 0, 355, 356, 5, 58,
		0, 0, 356, 357, 5, 60, 0, 0, 357, 358, 3, 36, 18, 23, 358, 449, 1, 0, 0,
		0, 359, 360, 5, 68, 0, 0, 360, 361, 5, 57, 0, 0, 361, 362, 3, 36, 18, 0,
		362, 363, 5, 58, 0, 0, 363, 364, 5, 57, 0, 0, 364, 366, 3, 36, 18, 0, 365,
		367, 5, 58, 0, 0, 366, 365, 1, 0, 0, 0, 366, 367, 1, 0, 0, 0, 367, 449,
		1, 0, 0, 0, 368, 369, 5, 68, 0, 0, 369, 370, 5, 57, 0, 0, 370, 371, 3,
		36, 18, 0, 371, 372, 5, 58, 0, 0, 372, 373, 5, 60, 0, 0, 373, 374, 3, 36,
		18, 21, 374, 449, 1, 0, 0, 0, 375, 376, 5, 68, 0, 0, 376, 377, 5, 57, 0,
		0, 377, 378, 3, 36, 18, 0, 378, 379, 5, 58, 0, 0, 379, 449, 1, 0, 0, 0,
		380, 381, 5, 41, 0, 0, 381, 382, 5, 55, 0, 0, 382, 383, 3, 36, 18, 0, 383,
		384, 5, 56, 0, 0, 384, 449, 1, 0, 0, 0, 385, 386, 5, 43, 0, 0, 386, 387,
		5, 55, 0, 0, 387, 388, 5, 68, 0, 0, 388, 389, 5, 1, 0, 0, 389, 390, 3,
		36, 18, 0, 390, 391, 5, 56, 0, 0, 391, 449, 1, 0, 0, 0, 392, 393, 5, 44,
		0, 0, 393, 394, 5, 55, 0, 0, 394, 395, 3, 36, 18, 0, 395, 396, 5, 56, 0,
		0, 396, 449, 1, 0, 0, 0, 397, 398, 5, 45, 0, 0, 398, 399, 5, 55, 0, 0,
		399, 400, 3, 36, 18, 0, 400, 401, 5, 56, 0, 0, 401, 449, 1, 0, 0, 0, 402,
		403, 5, 46, 0, 0, 403, 404, 5, 55, 0, 0, 404, 405, 5, 68, 0, 0, 405, 449,
		5, 56, 0, 0, 406, 407, 5, 42, 0, 0, 407, 408, 5, 55, 0, 0, 408, 409, 5,
		68, 0, 0, 409, 410, 5, 1, 0, 0, 410, 411, 3, 36, 18, 0, 411, 412, 5, 56,
		0, 0, 412, 449, 1, 0, 0, 0, 413, 414, 5, 68, 0, 0, 414, 449, 7, 1, 0, 0,
		415, 449, 5, 68, 0, 0, 416, 449, 5, 66, 0, 0, 417, 449, 5, 63, 0, 0, 418,
		449, 5, 64, 0, 0, 419, 449, 5, 65, 0, 0, 420, 449, 5, 62, 0, 0, 421, 425,
		5, 53, 0, 0, 422, 424, 3, 18, 9, 0, 423, 422, 1, 0, 0, 0, 424, 427, 1,
		0, 0, 0, 425, 423, 1, 0, 0, 0, 425, 426, 1, 0, 0, 0, 426, 428, 1, 0, 0,
		0, 427, 425, 1, 0, 0, 0, 428, 449, 5, 54, 0, 0, 429, 430, 5, 57, 0, 0,
		430, 431, 5, 58, 0, 0, 431, 432, 3, 42, 21, 0, 432, 433, 5, 53, 0, 0, 433,
		434, 3, 18, 9, 0, 434, 435, 5, 54, 0, 0, 435, 449, 1, 0, 0, 0, 436, 437,
		5, 40, 0, 0, 437, 438, 5, 55, 0, 0, 438, 439, 5, 68, 0, 0, 439, 440, 5,
		1, 0, 0, 440, 441, 3, 36, 18, 0, 441, 442, 5, 56, 0, 0, 442, 449, 1, 0,
		0, 0, 443, 449, 5, 50, 0, 0, 444, 445, 5, 55, 0, 0, 445, 446, 3, 36, 18,
		0, 446, 447, 5, 56, 0, 0, 447, 449, 1, 0, 0, 0, 448, 329, 1, 0, 0, 0, 448,
		332, 1, 0, 0, 0, 448, 334, 1, 0, 0, 0, 448, 337, 1, 0, 0, 0, 448, 344,
		1, 0, 0, 0, 448, 349, 1, 0, 0, 0, 448, 359, 1, 0, 0, 0, 448, 368, 1, 0,
		0, 0, 448, 375, 1, 0, 0, 0, 448, 380, 1, 0, 0, 0, 448, 385, 1, 0, 0, 0,
		448, 392, 1, 0, 0, 0, 448, 397, 1, 0, 0, 0, 448, 402, 1, 0, 0, 0, 448,
		406, 1, 0, 0, 0, 448, 413, 1, 0, 0, 0, 448, 415, 1, 0, 0, 0, 448, 416,
		1, 0, 0, 0, 448, 417, 1, 0, 0, 0, 448, 418, 1, 0, 0, 0, 448, 419, 1, 0,
		0, 0, 448, 420, 1, 0, 0, 0, 448, 421, 1, 0, 0, 0, 448, 429, 1, 0, 0, 0,
		448, 436, 1, 0, 0, 0, 448, 443, 1, 0, 0, 0, 448, 444, 1, 0, 0, 0, 449,
		479, 1, 0, 0, 0, 450, 451, 10, 32, 0, 0, 451, 452, 7, 2, 0, 0, 452, 478,
		3, 36, 18, 33, 453, 454, 10, 31, 0, 0, 454, 455, 7, 3, 0, 0, 455, 478,
		3, 36, 18, 32, 456, 457, 10, 30, 0, 0, 457, 458, 7, 4, 0, 0, 458, 478,
		3, 36, 18, 31, 459, 460, 10, 29, 0, 0, 460, 461, 5, 18, 0, 0, 461, 478,
		3, 36, 18, 30, 462, 463, 10, 28, 0, 0, 463, 464, 7, 5, 0, 0, 464, 478,
		3, 36, 18, 29, 465, 466, 10, 26, 0, 0, 466, 467, 7, 6, 0, 0, 467, 478,
		3, 36, 18, 27, 468, 469, 10, 13, 0, 0, 469, 470, 5, 60, 0, 0, 470, 478,
		3, 36, 18, 14, 471, 473, 10, 33, 0, 0, 472, 474, 3, 38, 19, 0, 473, 472,
		1, 0, 0, 0, 474, 475, 1, 0, 0, 0, 475, 473, 1, 0, 0, 0, 475, 476, 1, 0,
		0, 0, 476, 478, 1, 0, 0, 0, 477, 450, 1, 0, 0, 0, 477, 453, 1, 0, 0, 0,
		477, 456, 1, 0, 0, 0, 477, 459, 1, 0, 0, 0, 477, 462, 1, 0, 0, 0, 477,
		465, 1, 0, 0, 0, 477, 468, 1, 0, 0, 0, 477, 471, 1, 0, 0, 0, 478, 481,
		1, 0, 0, 0, 479, 477, 1, 0, 0, 0, 479, 480, 1, 0, 0, 0, 480, 37, 1, 0,
		0, 0, 481, 479, 1, 0, 0, 0, 482, 484, 5, 55, 0, 0, 483, 485, 3, 34, 17,
		0, 484, 483, 1, 0, 0, 0, 484, 485, 1, 0, 0, 0, 485, 486, 1, 0, 0, 0, 486,
		490, 5, 56, 0, 0, 487, 488, 5, 30, 0, 0, 488, 490, 5, 68, 0, 0, 489, 482,
		1, 0, 0, 0, 489, 487, 1, 0, 0, 0, 490, 39, 1, 0, 0, 0, 491, 492, 5, 68,
		0, 0, 492, 493, 5, 7, 0, 0, 493, 500, 3, 36, 18, 0, 494, 495, 5, 1, 0,
		0, 495, 496, 5, 68, 0, 0, 496, 497, 5, 7, 0, 0, 497, 499, 3, 36, 18, 0,
		498, 494, 1, 0, 0, 0, 499, 502, 1, 0, 0, 0, 500, 498, 1, 0, 0, 0, 500,
		501, 1, 0, 0, 0, 501, 504, 1, 0, 0, 0, 502, 500, 1, 0, 0, 0, 503, 505,
		5, 1, 0, 0, 504, 503, 1, 0, 0, 0, 504, 505, 1, 0, 0, 0, 505, 41, 1, 0,
		0, 0, 506, 507, 7, 7, 0, 0, 507, 43, 1, 0, 0, 0, 47, 47, 52, 56, 60, 64,
		68, 72, 74, 97, 103, 107, 113, 122, 126, 132, 137, 142, 149, 154, 164,
		171, 190, 197, 221, 225, 231, 238, 248, 252, 260, 264, 292, 294, 310, 318,
		326, 341, 366, 425, 448, 475, 477, 479, 484, 489, 500, 504,
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

// V4LangGrammarParserInit initializes any static state used to implement V4LangGrammarParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewV4LangGrammarParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func V4LangGrammarParserInit() {
	staticData := &V4LangGrammarParserStaticData
	staticData.once.Do(v4langgrammarParserInit)
}

// NewV4LangGrammarParser produces a new parser instance for the optional input antlr.TokenStream.
func NewV4LangGrammarParser(input antlr.TokenStream) *V4LangGrammarParser {
	V4LangGrammarParserInit()
	this := new(V4LangGrammarParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &V4LangGrammarParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "V4LangGrammar.g4"

	return this
}

// V4LangGrammarParser tokens.
const (
	V4LangGrammarParserEOF               = antlr.TokenEOF
	V4LangGrammarParserT__0              = 1
	V4LangGrammarParserT__1              = 2
	V4LangGrammarParserT__2              = 3
	V4LangGrammarParserT__3              = 4
	V4LangGrammarParserT__4              = 5
	V4LangGrammarParserT__5              = 6
	V4LangGrammarParserT__6              = 7
	V4LangGrammarParserT__7              = 8
	V4LangGrammarParserT__8              = 9
	V4LangGrammarParserT__9              = 10
	V4LangGrammarParserT__10             = 11
	V4LangGrammarParserT__11             = 12
	V4LangGrammarParserT__12             = 13
	V4LangGrammarParserT__13             = 14
	V4LangGrammarParserT__14             = 15
	V4LangGrammarParserT__15             = 16
	V4LangGrammarParserT__16             = 17
	V4LangGrammarParserT__17             = 18
	V4LangGrammarParserT__18             = 19
	V4LangGrammarParserT__19             = 20
	V4LangGrammarParserT__20             = 21
	V4LangGrammarParserT__21             = 22
	V4LangGrammarParserT__22             = 23
	V4LangGrammarParserT__23             = 24
	V4LangGrammarParserT__24             = 25
	V4LangGrammarParserT__25             = 26
	V4LangGrammarParserT__26             = 27
	V4LangGrammarParserT__27             = 28
	V4LangGrammarParserT__28             = 29
	V4LangGrammarParserT__29             = 30
	V4LangGrammarParserT__30             = 31
	V4LangGrammarParserT__31             = 32
	V4LangGrammarParserT__32             = 33
	V4LangGrammarParserT__33             = 34
	V4LangGrammarParserT__34             = 35
	V4LangGrammarParserMUT               = 36
	V4LangGrammarParserPRINT             = 37
	V4LangGrammarParserIF                = 38
	V4LangGrammarParserELSE              = 39
	V4LangGrammarParserAPPEND            = 40
	V4LangGrammarParserLEN               = 41
	V4LangGrammarParserSLICE_INDEX       = 42
	V4LangGrammarParserJOIN              = 43
	V4LangGrammarParserATOI              = 44
	V4LangGrammarParserPARSEFLOAT        = 45
	V4LangGrammarParserTYPEOF            = 46
	V4LangGrammarParserFOR               = 47
	V4LangGrammarParserRANGE             = 48
	V4LangGrammarParserBREAK             = 49
	V4LangGrammarParserNIL               = 50
	V4LangGrammarParserFUNC              = 51
	V4LangGrammarParserSTRUCT            = 52
	V4LangGrammarParserL_CURLY           = 53
	V4LangGrammarParserR_CURLY           = 54
	V4LangGrammarParserL_PAREN           = 55
	V4LangGrammarParserR_PAREN           = 56
	V4LangGrammarParserL_BRACKET         = 57
	V4LangGrammarParserR_BRACKET         = 58
	V4LangGrammarParserDECLARE_ASSIGN    = 59
	V4LangGrammarParserASSIGN            = 60
	V4LangGrammarParserSEMICOLON         = 61
	V4LangGrammarParserINT               = 62
	V4LangGrammarParserBOOL              = 63
	V4LangGrammarParserFLOAT             = 64
	V4LangGrammarParserSTRING            = 65
	V4LangGrammarParserRUNE              = 66
	V4LangGrammarParserWS                = 67
	V4LangGrammarParserID                = 68
	V4LangGrammarParserCOMMENT           = 69
	V4LangGrammarParserMULTILINE_COMMENT = 70
)

// V4LangGrammarParser rules.
const (
	V4LangGrammarParserRULE_program        = 0
	V4LangGrammarParserRULE_dcl            = 1
	V4LangGrammarParserRULE_varDcl         = 2
	V4LangGrammarParserRULE_funcDcl        = 3
	V4LangGrammarParserRULE_params         = 4
	V4LangGrammarParserRULE_structDcl      = 5
	V4LangGrammarParserRULE_structBody     = 6
	V4LangGrammarParserRULE_receiver       = 7
	V4LangGrammarParserRULE_sliceDcl       = 8
	V4LangGrammarParserRULE_sliceValues    = 9
	V4LangGrammarParserRULE_matrixDcl      = 10
	V4LangGrammarParserRULE_matrixValues   = 11
	V4LangGrammarParserRULE_statement      = 12
	V4LangGrammarParserRULE_elseIfStmt     = 13
	V4LangGrammarParserRULE_elseStmt       = 14
	V4LangGrammarParserRULE_switchCase     = 15
	V4LangGrammarParserRULE_defaultCase    = 16
	V4LangGrammarParserRULE_args           = 17
	V4LangGrammarParserRULE_expr           = 18
	V4LangGrammarParserRULE_call           = 19
	V4LangGrammarParserRULE_instanceValues = 20
	V4LangGrammarParserRULE_type           = 21
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllDcl() []IDclContext
	Dcl(i int) IDclContext

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V4LangGrammarParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V4LangGrammarParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = V4LangGrammarParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) AllDcl() []IDclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDclContext); ok {
			len++
		}
	}

	tst := make([]IDclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDclContext); ok {
			tst[i] = t.(IDclContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Dcl(i int) IDclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDclContext); ok {
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

	return t.(IDclContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V4LangGrammarVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V4LangGrammarParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, V4LangGrammarParserRULE_program)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(47)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-4418313327768238532) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&23) != 0) {
		{
			p.SetState(44)
			p.Dcl()
		}

		p.SetState(49)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDclContext is an interface to support dynamic dispatch.
type IDclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VarDcl() IVarDclContext
	SEMICOLON() antlr.TerminalNode
	MatrixDcl() IMatrixDclContext
	SliceDcl() ISliceDclContext
	Statement() IStatementContext
	FuncDcl() IFuncDclContext
	StructDcl() IStructDclContext

	// IsDclContext differentiates from other interfaces.
	IsDclContext()
}

type DclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDclContext() *DclContext {
	var p = new(DclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V4LangGrammarParserRULE_dcl
	return p
}

func InitEmptyDclContext(p *DclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V4LangGrammarParserRULE_dcl
}

func (*DclContext) IsDclContext() {}

func NewDclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DclContext {
	var p = new(DclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = V4LangGrammarParserRULE_dcl

	return p
}

func (s *DclContext) GetParser() antlr.Parser { return s.parser }

func (s *DclContext) VarDcl() IVarDclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarDclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarDclContext)
}

func (s *DclContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserSEMICOLON, 0)
}

func (s *DclContext) MatrixDcl() IMatrixDclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMatrixDclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMatrixDclContext)
}

func (s *DclContext) SliceDcl() ISliceDclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISliceDclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISliceDclContext)
}

func (s *DclContext) Statement() IStatementContext {
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

func (s *DclContext) FuncDcl() IFuncDclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncDclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncDclContext)
}

func (s *DclContext) StructDcl() IStructDclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructDclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructDclContext)
}

func (s *DclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.EnterDcl(s)
	}
}

func (s *DclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.ExitDcl(s)
	}
}

func (s *DclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V4LangGrammarVisitor:
		return t.VisitDcl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V4LangGrammarParser) Dcl() (localctx IDclContext) {
	localctx = NewDclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, V4LangGrammarParserRULE_dcl)
	var _la int

	p.SetState(74)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(50)
			p.VarDcl()
		}
		p.SetState(52)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == V4LangGrammarParserSEMICOLON {
			{
				p.SetState(51)
				p.Match(V4LangGrammarParserSEMICOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(54)
			p.MatrixDcl()
		}
		p.SetState(56)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == V4LangGrammarParserSEMICOLON {
			{
				p.SetState(55)
				p.Match(V4LangGrammarParserSEMICOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(58)
			p.SliceDcl()
		}
		p.SetState(60)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == V4LangGrammarParserSEMICOLON {
			{
				p.SetState(59)
				p.Match(V4LangGrammarParserSEMICOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(62)
			p.Statement()
		}
		p.SetState(64)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == V4LangGrammarParserSEMICOLON {
			{
				p.SetState(63)
				p.Match(V4LangGrammarParserSEMICOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(66)
			p.FuncDcl()
		}
		p.SetState(68)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == V4LangGrammarParserSEMICOLON {
			{
				p.SetState(67)
				p.Match(V4LangGrammarParserSEMICOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(70)
			p.StructDcl()
		}
		p.SetState(72)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == V4LangGrammarParserSEMICOLON {
			{
				p.SetState(71)
				p.Match(V4LangGrammarParserSEMICOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVarDclContext is an interface to support dynamic dispatch.
type IVarDclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MUT() antlr.TerminalNode
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	Type_() ITypeContext
	ASSIGN() antlr.TerminalNode
	Expr() IExprContext
	L_CURLY() antlr.TerminalNode
	R_CURLY() antlr.TerminalNode
	DECLARE_ASSIGN() antlr.TerminalNode

	// IsVarDclContext differentiates from other interfaces.
	IsVarDclContext()
}

type VarDclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarDclContext() *VarDclContext {
	var p = new(VarDclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V4LangGrammarParserRULE_varDcl
	return p
}

func InitEmptyVarDclContext(p *VarDclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V4LangGrammarParserRULE_varDcl
}

func (*VarDclContext) IsVarDclContext() {}

func NewVarDclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarDclContext {
	var p = new(VarDclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = V4LangGrammarParserRULE_varDcl

	return p
}

func (s *VarDclContext) GetParser() antlr.Parser { return s.parser }

func (s *VarDclContext) MUT() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserMUT, 0)
}

func (s *VarDclContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(V4LangGrammarParserID)
}

func (s *VarDclContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserID, i)
}

func (s *VarDclContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *VarDclContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserASSIGN, 0)
}

func (s *VarDclContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *VarDclContext) L_CURLY() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserL_CURLY, 0)
}

func (s *VarDclContext) R_CURLY() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserR_CURLY, 0)
}

func (s *VarDclContext) DECLARE_ASSIGN() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserDECLARE_ASSIGN, 0)
}

func (s *VarDclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarDclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarDclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.EnterVarDcl(s)
	}
}

func (s *VarDclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.ExitVarDcl(s)
	}
}

func (s *VarDclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V4LangGrammarVisitor:
		return t.VisitVarDcl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V4LangGrammarParser) VarDcl() (localctx IVarDclContext) {
	localctx = NewVarDclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, V4LangGrammarParserRULE_varDcl)
	p.SetState(97)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(76)
			p.Match(V4LangGrammarParserMUT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(77)
			p.Match(V4LangGrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(78)
			p.Type_()
		}
		{
			p.SetState(79)
			p.Match(V4LangGrammarParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(80)
			p.expr(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(82)
			p.Match(V4LangGrammarParserMUT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(83)
			p.Match(V4LangGrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(84)
			p.Type_()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(85)
			p.Match(V4LangGrammarParserMUT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(86)
			p.Match(V4LangGrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(87)
			p.Match(V4LangGrammarParserL_CURLY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(88)
			p.Match(V4LangGrammarParserR_CURLY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(89)
			p.Type_()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(90)
			p.Match(V4LangGrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(91)
			p.Match(V4LangGrammarParserDECLARE_ASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(92)
			p.expr(0)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(93)
			p.Match(V4LangGrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(94)
			p.Type_()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(95)
			p.Match(V4LangGrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(96)
			p.Match(V4LangGrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncDclContext is an interface to support dynamic dispatch.
type IFuncDclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FUNC() antlr.TerminalNode
	ID() antlr.TerminalNode
	L_PAREN() antlr.TerminalNode
	R_PAREN() antlr.TerminalNode
	L_CURLY() antlr.TerminalNode
	R_CURLY() antlr.TerminalNode
	Params() IParamsContext
	Type_() ITypeContext
	AllDcl() []IDclContext
	Dcl(i int) IDclContext
	Receiver() IReceiverContext

	// IsFuncDclContext differentiates from other interfaces.
	IsFuncDclContext()
}

type FuncDclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncDclContext() *FuncDclContext {
	var p = new(FuncDclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V4LangGrammarParserRULE_funcDcl
	return p
}

func InitEmptyFuncDclContext(p *FuncDclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V4LangGrammarParserRULE_funcDcl
}

func (*FuncDclContext) IsFuncDclContext() {}

func NewFuncDclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncDclContext {
	var p = new(FuncDclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = V4LangGrammarParserRULE_funcDcl

	return p
}

func (s *FuncDclContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncDclContext) FUNC() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserFUNC, 0)
}

func (s *FuncDclContext) ID() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserID, 0)
}

func (s *FuncDclContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserL_PAREN, 0)
}

func (s *FuncDclContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserR_PAREN, 0)
}

func (s *FuncDclContext) L_CURLY() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserL_CURLY, 0)
}

func (s *FuncDclContext) R_CURLY() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserR_CURLY, 0)
}

func (s *FuncDclContext) Params() IParamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamsContext)
}

func (s *FuncDclContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *FuncDclContext) AllDcl() []IDclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDclContext); ok {
			len++
		}
	}

	tst := make([]IDclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDclContext); ok {
			tst[i] = t.(IDclContext)
			i++
		}
	}

	return tst
}

func (s *FuncDclContext) Dcl(i int) IDclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDclContext); ok {
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

	return t.(IDclContext)
}

func (s *FuncDclContext) Receiver() IReceiverContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReceiverContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReceiverContext)
}

func (s *FuncDclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncDclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncDclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.EnterFuncDcl(s)
	}
}

func (s *FuncDclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.ExitFuncDcl(s)
	}
}

func (s *FuncDclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V4LangGrammarVisitor:
		return t.VisitFuncDcl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V4LangGrammarParser) FuncDcl() (localctx IFuncDclContext) {
	localctx = NewFuncDclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, V4LangGrammarParserRULE_funcDcl)
	var _la int

	p.SetState(137)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(99)
			p.Match(V4LangGrammarParserFUNC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(100)
			p.Match(V4LangGrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(101)
			p.Match(V4LangGrammarParserL_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(103)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == V4LangGrammarParserID {
			{
				p.SetState(102)
				p.Params()
			}

		}
		{
			p.SetState(105)
			p.Match(V4LangGrammarParserR_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(107)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&66571993088) != 0 {
			{
				p.SetState(106)
				p.Type_()
			}

		}
		{
			p.SetState(109)
			p.Match(V4LangGrammarParserL_CURLY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(113)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-4418313327768238532) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&23) != 0) {
			{
				p.SetState(110)
				p.Dcl()
			}

			p.SetState(115)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(116)
			p.Match(V4LangGrammarParserR_CURLY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(117)
			p.Match(V4LangGrammarParserFUNC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(118)
			p.Receiver()
		}
		{
			p.SetState(119)
			p.Match(V4LangGrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(120)
			p.Match(V4LangGrammarParserL_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(122)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == V4LangGrammarParserID {
			{
				p.SetState(121)
				p.Params()
			}

		}
		{
			p.SetState(124)
			p.Match(V4LangGrammarParserR_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(126)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&66571993088) != 0 {
			{
				p.SetState(125)
				p.Type_()
			}

		}
		{
			p.SetState(128)
			p.Match(V4LangGrammarParserL_CURLY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(132)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-4418313327768238532) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&23) != 0) {
			{
				p.SetState(129)
				p.Dcl()
			}

			p.SetState(134)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(135)
			p.Match(V4LangGrammarParserR_CURLY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParamsContext is an interface to support dynamic dispatch.
type IParamsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	AllType_() []ITypeContext
	Type_(i int) ITypeContext
	AllL_BRACKET() []antlr.TerminalNode
	L_BRACKET(i int) antlr.TerminalNode
	AllR_BRACKET() []antlr.TerminalNode
	R_BRACKET(i int) antlr.TerminalNode

	// IsParamsContext differentiates from other interfaces.
	IsParamsContext()
}

type ParamsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamsContext() *ParamsContext {
	var p = new(ParamsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V4LangGrammarParserRULE_params
	return p
}

func InitEmptyParamsContext(p *ParamsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V4LangGrammarParserRULE_params
}

func (*ParamsContext) IsParamsContext() {}

func NewParamsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamsContext {
	var p = new(ParamsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = V4LangGrammarParserRULE_params

	return p
}

func (s *ParamsContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamsContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(V4LangGrammarParserID)
}

func (s *ParamsContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserID, i)
}

func (s *ParamsContext) AllType_() []ITypeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITypeContext); ok {
			len++
		}
	}

	tst := make([]ITypeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITypeContext); ok {
			tst[i] = t.(ITypeContext)
			i++
		}
	}

	return tst
}

func (s *ParamsContext) Type_(i int) ITypeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
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

	return t.(ITypeContext)
}

func (s *ParamsContext) AllL_BRACKET() []antlr.TerminalNode {
	return s.GetTokens(V4LangGrammarParserL_BRACKET)
}

func (s *ParamsContext) L_BRACKET(i int) antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserL_BRACKET, i)
}

func (s *ParamsContext) AllR_BRACKET() []antlr.TerminalNode {
	return s.GetTokens(V4LangGrammarParserR_BRACKET)
}

func (s *ParamsContext) R_BRACKET(i int) antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserR_BRACKET, i)
}

func (s *ParamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.EnterParams(s)
	}
}

func (s *ParamsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.ExitParams(s)
	}
}

func (s *ParamsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V4LangGrammarVisitor:
		return t.VisitParams(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V4LangGrammarParser) Params() (localctx IParamsContext) {
	localctx = NewParamsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, V4LangGrammarParserRULE_params)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(139)
		p.Match(V4LangGrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(142)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == V4LangGrammarParserL_BRACKET {
		{
			p.SetState(140)
			p.Match(V4LangGrammarParserL_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(141)
			p.Match(V4LangGrammarParserR_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(144)
		p.Type_()
	}
	p.SetState(154)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == V4LangGrammarParserT__0 {
		{
			p.SetState(145)
			p.Match(V4LangGrammarParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(146)
			p.Match(V4LangGrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(149)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == V4LangGrammarParserL_BRACKET {
			{
				p.SetState(147)
				p.Match(V4LangGrammarParserL_BRACKET)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(148)
				p.Match(V4LangGrammarParserR_BRACKET)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(151)
			p.Type_()
		}

		p.SetState(156)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStructDclContext is an interface to support dynamic dispatch.
type IStructDclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	STRUCT() antlr.TerminalNode
	L_CURLY() antlr.TerminalNode
	R_CURLY() antlr.TerminalNode
	AllStructBody() []IStructBodyContext
	StructBody(i int) IStructBodyContext

	// IsStructDclContext differentiates from other interfaces.
	IsStructDclContext()
}

type StructDclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructDclContext() *StructDclContext {
	var p = new(StructDclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V4LangGrammarParserRULE_structDcl
	return p
}

func InitEmptyStructDclContext(p *StructDclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V4LangGrammarParserRULE_structDcl
}

func (*StructDclContext) IsStructDclContext() {}

func NewStructDclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructDclContext {
	var p = new(StructDclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = V4LangGrammarParserRULE_structDcl

	return p
}

func (s *StructDclContext) GetParser() antlr.Parser { return s.parser }

func (s *StructDclContext) ID() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserID, 0)
}

func (s *StructDclContext) STRUCT() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserSTRUCT, 0)
}

func (s *StructDclContext) L_CURLY() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserL_CURLY, 0)
}

func (s *StructDclContext) R_CURLY() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserR_CURLY, 0)
}

func (s *StructDclContext) AllStructBody() []IStructBodyContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStructBodyContext); ok {
			len++
		}
	}

	tst := make([]IStructBodyContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStructBodyContext); ok {
			tst[i] = t.(IStructBodyContext)
			i++
		}
	}

	return tst
}

func (s *StructDclContext) StructBody(i int) IStructBodyContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructBodyContext); ok {
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

	return t.(IStructBodyContext)
}

func (s *StructDclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructDclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructDclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.EnterStructDcl(s)
	}
}

func (s *StructDclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.ExitStructDcl(s)
	}
}

func (s *StructDclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V4LangGrammarVisitor:
		return t.VisitStructDcl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V4LangGrammarParser) StructDcl() (localctx IStructDclContext) {
	localctx = NewStructDclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, V4LangGrammarParserRULE_structDcl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(157)
		p.Match(V4LangGrammarParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(158)
		p.Match(V4LangGrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(159)
		p.Match(V4LangGrammarParserSTRUCT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(160)
		p.Match(V4LangGrammarParserL_CURLY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(164)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64((_la-36)) & ^0x3f) == 0 && ((int64(1)<<(_la-36))&4295000065) != 0 {
		{
			p.SetState(161)
			p.StructBody()
		}

		p.SetState(166)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(167)
		p.Match(V4LangGrammarParserR_CURLY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStructBodyContext is an interface to support dynamic dispatch.
type IStructBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VarDcl() IVarDclContext
	FuncDcl() IFuncDclContext

	// IsStructBodyContext differentiates from other interfaces.
	IsStructBodyContext()
}

type StructBodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructBodyContext() *StructBodyContext {
	var p = new(StructBodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V4LangGrammarParserRULE_structBody
	return p
}

func InitEmptyStructBodyContext(p *StructBodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V4LangGrammarParserRULE_structBody
}

func (*StructBodyContext) IsStructBodyContext() {}

func NewStructBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructBodyContext {
	var p = new(StructBodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = V4LangGrammarParserRULE_structBody

	return p
}

func (s *StructBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *StructBodyContext) VarDcl() IVarDclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarDclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarDclContext)
}

func (s *StructBodyContext) FuncDcl() IFuncDclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncDclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncDclContext)
}

func (s *StructBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.EnterStructBody(s)
	}
}

func (s *StructBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.ExitStructBody(s)
	}
}

func (s *StructBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V4LangGrammarVisitor:
		return t.VisitStructBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V4LangGrammarParser) StructBody() (localctx IStructBodyContext) {
	localctx = NewStructBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, V4LangGrammarParserRULE_structBody)
	p.SetState(171)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case V4LangGrammarParserMUT, V4LangGrammarParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(169)
			p.VarDcl()
		}

	case V4LangGrammarParserFUNC:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(170)
			p.FuncDcl()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IReceiverContext is an interface to support dynamic dispatch.
type IReceiverContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetThis returns the this token.
	GetThis() antlr.Token

	// GetStructName returns the structName token.
	GetStructName() antlr.Token

	// SetThis sets the this token.
	SetThis(antlr.Token)

	// SetStructName sets the structName token.
	SetStructName(antlr.Token)

	// Getter signatures
	L_PAREN() antlr.TerminalNode
	R_PAREN() antlr.TerminalNode
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode

	// IsReceiverContext differentiates from other interfaces.
	IsReceiverContext()
}

type ReceiverContext struct {
	antlr.BaseParserRuleContext
	parser     antlr.Parser
	this       antlr.Token
	structName antlr.Token
}

func NewEmptyReceiverContext() *ReceiverContext {
	var p = new(ReceiverContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V4LangGrammarParserRULE_receiver
	return p
}

func InitEmptyReceiverContext(p *ReceiverContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V4LangGrammarParserRULE_receiver
}

func (*ReceiverContext) IsReceiverContext() {}

func NewReceiverContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReceiverContext {
	var p = new(ReceiverContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = V4LangGrammarParserRULE_receiver

	return p
}

func (s *ReceiverContext) GetParser() antlr.Parser { return s.parser }

func (s *ReceiverContext) GetThis() antlr.Token { return s.this }

func (s *ReceiverContext) GetStructName() antlr.Token { return s.structName }

func (s *ReceiverContext) SetThis(v antlr.Token) { s.this = v }

func (s *ReceiverContext) SetStructName(v antlr.Token) { s.structName = v }

func (s *ReceiverContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserL_PAREN, 0)
}

func (s *ReceiverContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserR_PAREN, 0)
}

func (s *ReceiverContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(V4LangGrammarParserID)
}

func (s *ReceiverContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserID, i)
}

func (s *ReceiverContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReceiverContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReceiverContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.EnterReceiver(s)
	}
}

func (s *ReceiverContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.ExitReceiver(s)
	}
}

func (s *ReceiverContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V4LangGrammarVisitor:
		return t.VisitReceiver(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V4LangGrammarParser) Receiver() (localctx IReceiverContext) {
	localctx = NewReceiverContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, V4LangGrammarParserRULE_receiver)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(173)
		p.Match(V4LangGrammarParserL_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(174)

		var _m = p.Match(V4LangGrammarParserID)

		localctx.(*ReceiverContext).this = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(175)

		var _m = p.Match(V4LangGrammarParserID)

		localctx.(*ReceiverContext).structName = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(176)
		p.Match(V4LangGrammarParserR_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISliceDclContext is an interface to support dynamic dispatch.
type ISliceDclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MUT() antlr.TerminalNode
	ID() antlr.TerminalNode
	DECLARE_ASSIGN() antlr.TerminalNode
	L_BRACKET() antlr.TerminalNode
	R_BRACKET() antlr.TerminalNode
	Type_() ITypeContext
	L_CURLY() antlr.TerminalNode
	SliceValues() ISliceValuesContext
	R_CURLY() antlr.TerminalNode

	// IsSliceDclContext differentiates from other interfaces.
	IsSliceDclContext()
}

type SliceDclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySliceDclContext() *SliceDclContext {
	var p = new(SliceDclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V4LangGrammarParserRULE_sliceDcl
	return p
}

func InitEmptySliceDclContext(p *SliceDclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V4LangGrammarParserRULE_sliceDcl
}

func (*SliceDclContext) IsSliceDclContext() {}

func NewSliceDclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SliceDclContext {
	var p = new(SliceDclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = V4LangGrammarParserRULE_sliceDcl

	return p
}

func (s *SliceDclContext) GetParser() antlr.Parser { return s.parser }

func (s *SliceDclContext) MUT() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserMUT, 0)
}

func (s *SliceDclContext) ID() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserID, 0)
}

func (s *SliceDclContext) DECLARE_ASSIGN() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserDECLARE_ASSIGN, 0)
}

func (s *SliceDclContext) L_BRACKET() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserL_BRACKET, 0)
}

func (s *SliceDclContext) R_BRACKET() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserR_BRACKET, 0)
}

func (s *SliceDclContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *SliceDclContext) L_CURLY() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserL_CURLY, 0)
}

func (s *SliceDclContext) SliceValues() ISliceValuesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISliceValuesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISliceValuesContext)
}

func (s *SliceDclContext) R_CURLY() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserR_CURLY, 0)
}

func (s *SliceDclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SliceDclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SliceDclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.EnterSliceDcl(s)
	}
}

func (s *SliceDclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.ExitSliceDcl(s)
	}
}

func (s *SliceDclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V4LangGrammarVisitor:
		return t.VisitSliceDcl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V4LangGrammarParser) SliceDcl() (localctx ISliceDclContext) {
	localctx = NewSliceDclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, V4LangGrammarParserRULE_sliceDcl)
	p.SetState(190)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case V4LangGrammarParserMUT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(178)
			p.Match(V4LangGrammarParserMUT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(179)
			p.Match(V4LangGrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(180)
			p.Match(V4LangGrammarParserDECLARE_ASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(181)
			p.Match(V4LangGrammarParserL_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(182)
			p.Match(V4LangGrammarParserR_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(183)
			p.Type_()
		}
		{
			p.SetState(184)
			p.Match(V4LangGrammarParserL_CURLY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(185)
			p.SliceValues()
		}
		{
			p.SetState(186)
			p.Match(V4LangGrammarParserR_CURLY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case V4LangGrammarParserID:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(188)
			p.Match(V4LangGrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(189)
			p.Type_()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISliceValuesContext is an interface to support dynamic dispatch.
type ISliceValuesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpr() []IExprContext
	Expr(i int) IExprContext

	// IsSliceValuesContext differentiates from other interfaces.
	IsSliceValuesContext()
}

type SliceValuesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySliceValuesContext() *SliceValuesContext {
	var p = new(SliceValuesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V4LangGrammarParserRULE_sliceValues
	return p
}

func InitEmptySliceValuesContext(p *SliceValuesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V4LangGrammarParserRULE_sliceValues
}

func (*SliceValuesContext) IsSliceValuesContext() {}

func NewSliceValuesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SliceValuesContext {
	var p = new(SliceValuesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = V4LangGrammarParserRULE_sliceValues

	return p
}

func (s *SliceValuesContext) GetParser() antlr.Parser { return s.parser }

func (s *SliceValuesContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *SliceValuesContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *SliceValuesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SliceValuesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SliceValuesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.EnterSliceValues(s)
	}
}

func (s *SliceValuesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.ExitSliceValues(s)
	}
}

func (s *SliceValuesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V4LangGrammarVisitor:
		return t.VisitSliceValues(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V4LangGrammarParser) SliceValues() (localctx ISliceValuesContext) {
	localctx = NewSliceValuesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, V4LangGrammarParserRULE_sliceValues)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(192)
		p.expr(0)
	}
	p.SetState(197)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == V4LangGrammarParserT__0 {
		{
			p.SetState(193)
			p.Match(V4LangGrammarParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(194)
			p.expr(0)
		}

		p.SetState(199)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMatrixDclContext is an interface to support dynamic dispatch.
type IMatrixDclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	DECLARE_ASSIGN() antlr.TerminalNode
	AllL_BRACKET() []antlr.TerminalNode
	L_BRACKET(i int) antlr.TerminalNode
	AllR_BRACKET() []antlr.TerminalNode
	R_BRACKET(i int) antlr.TerminalNode
	Type_() ITypeContext
	L_CURLY() antlr.TerminalNode
	MatrixValues() IMatrixValuesContext
	R_CURLY() antlr.TerminalNode

	// IsMatrixDclContext differentiates from other interfaces.
	IsMatrixDclContext()
}

type MatrixDclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMatrixDclContext() *MatrixDclContext {
	var p = new(MatrixDclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V4LangGrammarParserRULE_matrixDcl
	return p
}

func InitEmptyMatrixDclContext(p *MatrixDclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V4LangGrammarParserRULE_matrixDcl
}

func (*MatrixDclContext) IsMatrixDclContext() {}

func NewMatrixDclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MatrixDclContext {
	var p = new(MatrixDclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = V4LangGrammarParserRULE_matrixDcl

	return p
}

func (s *MatrixDclContext) GetParser() antlr.Parser { return s.parser }

func (s *MatrixDclContext) ID() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserID, 0)
}

func (s *MatrixDclContext) DECLARE_ASSIGN() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserDECLARE_ASSIGN, 0)
}

func (s *MatrixDclContext) AllL_BRACKET() []antlr.TerminalNode {
	return s.GetTokens(V4LangGrammarParserL_BRACKET)
}

func (s *MatrixDclContext) L_BRACKET(i int) antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserL_BRACKET, i)
}

func (s *MatrixDclContext) AllR_BRACKET() []antlr.TerminalNode {
	return s.GetTokens(V4LangGrammarParserR_BRACKET)
}

func (s *MatrixDclContext) R_BRACKET(i int) antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserR_BRACKET, i)
}

func (s *MatrixDclContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *MatrixDclContext) L_CURLY() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserL_CURLY, 0)
}

func (s *MatrixDclContext) MatrixValues() IMatrixValuesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMatrixValuesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMatrixValuesContext)
}

func (s *MatrixDclContext) R_CURLY() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserR_CURLY, 0)
}

func (s *MatrixDclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MatrixDclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MatrixDclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.EnterMatrixDcl(s)
	}
}

func (s *MatrixDclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.ExitMatrixDcl(s)
	}
}

func (s *MatrixDclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V4LangGrammarVisitor:
		return t.VisitMatrixDcl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V4LangGrammarParser) MatrixDcl() (localctx IMatrixDclContext) {
	localctx = NewMatrixDclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, V4LangGrammarParserRULE_matrixDcl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(200)
		p.Match(V4LangGrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(201)
		p.Match(V4LangGrammarParserDECLARE_ASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(202)
		p.Match(V4LangGrammarParserL_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(203)
		p.Match(V4LangGrammarParserR_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(204)
		p.Match(V4LangGrammarParserL_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(205)
		p.Match(V4LangGrammarParserR_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(206)
		p.Type_()
	}
	{
		p.SetState(207)
		p.Match(V4LangGrammarParserL_CURLY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(208)
		p.MatrixValues()
	}
	{
		p.SetState(209)
		p.Match(V4LangGrammarParserR_CURLY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMatrixValuesContext is an interface to support dynamic dispatch.
type IMatrixValuesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllL_CURLY() []antlr.TerminalNode
	L_CURLY(i int) antlr.TerminalNode
	AllSliceValues() []ISliceValuesContext
	SliceValues(i int) ISliceValuesContext
	AllR_CURLY() []antlr.TerminalNode
	R_CURLY(i int) antlr.TerminalNode

	// IsMatrixValuesContext differentiates from other interfaces.
	IsMatrixValuesContext()
}

type MatrixValuesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMatrixValuesContext() *MatrixValuesContext {
	var p = new(MatrixValuesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V4LangGrammarParserRULE_matrixValues
	return p
}

func InitEmptyMatrixValuesContext(p *MatrixValuesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V4LangGrammarParserRULE_matrixValues
}

func (*MatrixValuesContext) IsMatrixValuesContext() {}

func NewMatrixValuesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MatrixValuesContext {
	var p = new(MatrixValuesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = V4LangGrammarParserRULE_matrixValues

	return p
}

func (s *MatrixValuesContext) GetParser() antlr.Parser { return s.parser }

func (s *MatrixValuesContext) AllL_CURLY() []antlr.TerminalNode {
	return s.GetTokens(V4LangGrammarParserL_CURLY)
}

func (s *MatrixValuesContext) L_CURLY(i int) antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserL_CURLY, i)
}

func (s *MatrixValuesContext) AllSliceValues() []ISliceValuesContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISliceValuesContext); ok {
			len++
		}
	}

	tst := make([]ISliceValuesContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISliceValuesContext); ok {
			tst[i] = t.(ISliceValuesContext)
			i++
		}
	}

	return tst
}

func (s *MatrixValuesContext) SliceValues(i int) ISliceValuesContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISliceValuesContext); ok {
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

	return t.(ISliceValuesContext)
}

func (s *MatrixValuesContext) AllR_CURLY() []antlr.TerminalNode {
	return s.GetTokens(V4LangGrammarParserR_CURLY)
}

func (s *MatrixValuesContext) R_CURLY(i int) antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserR_CURLY, i)
}

func (s *MatrixValuesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MatrixValuesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MatrixValuesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.EnterMatrixValues(s)
	}
}

func (s *MatrixValuesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.ExitMatrixValues(s)
	}
}

func (s *MatrixValuesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V4LangGrammarVisitor:
		return t.VisitMatrixValues(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V4LangGrammarParser) MatrixValues() (localctx IMatrixValuesContext) {
	localctx = NewMatrixValuesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, V4LangGrammarParserRULE_matrixValues)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(211)
		p.Match(V4LangGrammarParserL_CURLY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(212)
		p.SliceValues()
	}
	{
		p.SetState(213)
		p.Match(V4LangGrammarParserR_CURLY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(221)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 23, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(214)
				p.Match(V4LangGrammarParserT__0)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(215)
				p.Match(V4LangGrammarParserL_CURLY)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(216)
				p.SliceValues()
			}
			{
				p.SetState(217)
				p.Match(V4LangGrammarParserR_CURLY)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(223)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 23, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(225)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == V4LangGrammarParserT__0 {
		{
			p.SetState(224)
			p.Match(V4LangGrammarParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V4LangGrammarParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V4LangGrammarParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = V4LangGrammarParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) CopyAll(ctx *StatementContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ContinueStmtContext struct {
	StatementContext
}

func NewContinueStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ContinueStmtContext {
	var p = new(ContinueStmtContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *ContinueStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContinueStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.EnterContinueStmt(s)
	}
}

func (s *ContinueStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.ExitContinueStmt(s)
	}
}

func (s *ContinueStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V4LangGrammarVisitor:
		return t.VisitContinueStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

type SwitchStmtContext struct {
	StatementContext
}

func NewSwitchStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SwitchStmtContext {
	var p = new(SwitchStmtContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *SwitchStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SwitchStmtContext) L_CURLY() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserL_CURLY, 0)
}

func (s *SwitchStmtContext) R_CURLY() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserR_CURLY, 0)
}

func (s *SwitchStmtContext) AllSwitchCase() []ISwitchCaseContext {
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

func (s *SwitchStmtContext) SwitchCase(i int) ISwitchCaseContext {
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

func (s *SwitchStmtContext) DefaultCase() IDefaultCaseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefaultCaseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDefaultCaseContext)
}

func (s *SwitchStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.EnterSwitchStmt(s)
	}
}

func (s *SwitchStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.ExitSwitchStmt(s)
	}
}

func (s *SwitchStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V4LangGrammarVisitor:
		return t.VisitSwitchStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

type PrintStmtContext struct {
	StatementContext
}

func NewPrintStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrintStmtContext {
	var p = new(PrintStmtContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *PrintStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintStmtContext) PRINT() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserPRINT, 0)
}

func (s *PrintStmtContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserL_PAREN, 0)
}

func (s *PrintStmtContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserR_PAREN, 0)
}

func (s *PrintStmtContext) Args() IArgsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgsContext)
}

func (s *PrintStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.EnterPrintStmt(s)
	}
}

func (s *PrintStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.ExitPrintStmt(s)
	}
}

func (s *PrintStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V4LangGrammarVisitor:
		return t.VisitPrintStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

type IfStmtContext struct {
	StatementContext
}

func NewIfStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfStmtContext {
	var p = new(IfStmtContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *IfStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStmtContext) IF() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserIF, 0)
}

func (s *IfStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IfStmtContext) Statement() IStatementContext {
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

func (s *IfStmtContext) AllElseIfStmt() []IElseIfStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IElseIfStmtContext); ok {
			len++
		}
	}

	tst := make([]IElseIfStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IElseIfStmtContext); ok {
			tst[i] = t.(IElseIfStmtContext)
			i++
		}
	}

	return tst
}

func (s *IfStmtContext) ElseIfStmt(i int) IElseIfStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElseIfStmtContext); ok {
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

	return t.(IElseIfStmtContext)
}

func (s *IfStmtContext) ElseStmt() IElseStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElseStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElseStmtContext)
}

func (s *IfStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.EnterIfStmt(s)
	}
}

func (s *IfStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.ExitIfStmt(s)
	}
}

func (s *IfStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V4LangGrammarVisitor:
		return t.VisitIfStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprStmtContext struct {
	StatementContext
}

func NewExprStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprStmtContext {
	var p = new(ExprStmtContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *ExprStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.EnterExprStmt(s)
	}
}

func (s *ExprStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.ExitExprStmt(s)
	}
}

func (s *ExprStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V4LangGrammarVisitor:
		return t.VisitExprStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

type BreakStmtContext struct {
	StatementContext
}

func NewBreakStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BreakStmtContext {
	var p = new(BreakStmtContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *BreakStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BreakStmtContext) BREAK() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserBREAK, 0)
}

func (s *BreakStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.EnterBreakStmt(s)
	}
}

func (s *BreakStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.ExitBreakStmt(s)
	}
}

func (s *BreakStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V4LangGrammarVisitor:
		return t.VisitBreakStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

type BlockStmtContext struct {
	StatementContext
}

func NewBlockStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BlockStmtContext {
	var p = new(BlockStmtContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *BlockStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockStmtContext) L_CURLY() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserL_CURLY, 0)
}

func (s *BlockStmtContext) R_CURLY() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserR_CURLY, 0)
}

func (s *BlockStmtContext) AllDcl() []IDclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDclContext); ok {
			len++
		}
	}

	tst := make([]IDclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDclContext); ok {
			tst[i] = t.(IDclContext)
			i++
		}
	}

	return tst
}

func (s *BlockStmtContext) Dcl(i int) IDclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDclContext); ok {
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

	return t.(IDclContext)
}

func (s *BlockStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.EnterBlockStmt(s)
	}
}

func (s *BlockStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.ExitBlockStmt(s)
	}
}

func (s *BlockStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V4LangGrammarVisitor:
		return t.VisitBlockStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

type ForRangeStmtContext struct {
	StatementContext
}

func NewForRangeStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForRangeStmtContext {
	var p = new(ForRangeStmtContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *ForRangeStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForRangeStmtContext) FOR() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserFOR, 0)
}

func (s *ForRangeStmtContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(V4LangGrammarParserID)
}

func (s *ForRangeStmtContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserID, i)
}

func (s *ForRangeStmtContext) DECLARE_ASSIGN() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserDECLARE_ASSIGN, 0)
}

func (s *ForRangeStmtContext) RANGE() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserRANGE, 0)
}

func (s *ForRangeStmtContext) Statement() IStatementContext {
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

func (s *ForRangeStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.EnterForRangeStmt(s)
	}
}

func (s *ForRangeStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.ExitForRangeStmt(s)
	}
}

func (s *ForRangeStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V4LangGrammarVisitor:
		return t.VisitForRangeStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

type ForIncrementStmtContext struct {
	StatementContext
}

func NewForIncrementStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForIncrementStmtContext {
	var p = new(ForIncrementStmtContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *ForIncrementStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForIncrementStmtContext) FOR() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserFOR, 0)
}

func (s *ForIncrementStmtContext) VarDcl() IVarDclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarDclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarDclContext)
}

func (s *ForIncrementStmtContext) AllSEMICOLON() []antlr.TerminalNode {
	return s.GetTokens(V4LangGrammarParserSEMICOLON)
}

func (s *ForIncrementStmtContext) SEMICOLON(i int) antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserSEMICOLON, i)
}

func (s *ForIncrementStmtContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ForIncrementStmtContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *ForIncrementStmtContext) Statement() IStatementContext {
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

func (s *ForIncrementStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.EnterForIncrementStmt(s)
	}
}

func (s *ForIncrementStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.ExitForIncrementStmt(s)
	}
}

func (s *ForIncrementStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V4LangGrammarVisitor:
		return t.VisitForIncrementStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

type ForStmtContext struct {
	StatementContext
}

func NewForStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForStmtContext {
	var p = new(ForStmtContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *ForStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForStmtContext) FOR() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserFOR, 0)
}

func (s *ForStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ForStmtContext) Statement() IStatementContext {
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

func (s *ForStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.EnterForStmt(s)
	}
}

func (s *ForStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.ExitForStmt(s)
	}
}

func (s *ForStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V4LangGrammarVisitor:
		return t.VisitForStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

type ReturnStmtContext struct {
	StatementContext
}

func NewReturnStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ReturnStmtContext {
	var p = new(ReturnStmtContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *ReturnStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ReturnStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.EnterReturnStmt(s)
	}
}

func (s *ReturnStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.ExitReturnStmt(s)
	}
}

func (s *ReturnStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V4LangGrammarVisitor:
		return t.VisitReturnStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V4LangGrammarParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, V4LangGrammarParserRULE_statement)
	var _la int

	var _alt int

	p.SetState(294)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 32, p.GetParserRuleContext()) {
	case 1:
		localctx = NewExprStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(227)
			p.expr(0)
		}

	case 2:
		localctx = NewPrintStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(228)
			p.Match(V4LangGrammarParserPRINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(229)
			p.Match(V4LangGrammarParserL_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(231)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64((_la-9)) & ^0x3f) == 0 && ((int64(1)<<(_la-9))&856055836861267971) != 0 {
			{
				p.SetState(230)
				p.Args()
			}

		}
		{
			p.SetState(233)
			p.Match(V4LangGrammarParserR_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewBlockStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(234)
			p.Match(V4LangGrammarParserL_CURLY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(238)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-4418313327768238532) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&23) != 0) {
			{
				p.SetState(235)
				p.Dcl()
			}

			p.SetState(240)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(241)
			p.Match(V4LangGrammarParserR_CURLY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewIfStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(242)
			p.Match(V4LangGrammarParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(243)
			p.expr(0)
		}
		{
			p.SetState(244)
			p.Statement()
		}
		p.SetState(248)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(245)
					p.ElseIfStmt()
				}

			}
			p.SetState(250)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		p.SetState(252)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(251)
				p.ElseStmt()
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case 5:
		localctx = NewSwitchStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(254)
			p.Match(V4LangGrammarParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(255)
			p.expr(0)
		}
		{
			p.SetState(256)
			p.Match(V4LangGrammarParserL_CURLY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(260)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == V4LangGrammarParserT__5 {
			{
				p.SetState(257)
				p.SwitchCase()
			}

			p.SetState(262)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(264)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == V4LangGrammarParserT__7 {
			{
				p.SetState(263)
				p.DefaultCase()
			}

		}
		{
			p.SetState(266)
			p.Match(V4LangGrammarParserR_CURLY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		localctx = NewForIncrementStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(268)
			p.Match(V4LangGrammarParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(269)
			p.VarDcl()
		}
		{
			p.SetState(270)
			p.Match(V4LangGrammarParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(271)
			p.expr(0)
		}
		{
			p.SetState(272)
			p.Match(V4LangGrammarParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(273)
			p.expr(0)
		}
		{
			p.SetState(274)
			p.Statement()
		}

	case 7:
		localctx = NewForStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(276)
			p.Match(V4LangGrammarParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(277)
			p.expr(0)
		}
		{
			p.SetState(278)
			p.Statement()
		}

	case 8:
		localctx = NewForRangeStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(280)
			p.Match(V4LangGrammarParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(281)
			p.Match(V4LangGrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(282)
			p.Match(V4LangGrammarParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(283)
			p.Match(V4LangGrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(284)
			p.Match(V4LangGrammarParserDECLARE_ASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(285)
			p.Match(V4LangGrammarParserRANGE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(286)
			p.Match(V4LangGrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(287)
			p.Statement()
		}

	case 9:
		localctx = NewBreakStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(288)
			p.Match(V4LangGrammarParserBREAK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 10:
		localctx = NewContinueStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(289)
			p.Match(V4LangGrammarParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 11:
		localctx = NewReturnStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(290)
			p.Match(V4LangGrammarParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(292)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(291)
				p.expr(0)
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IElseIfStmtContext is an interface to support dynamic dispatch.
type IElseIfStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ELSE() antlr.TerminalNode
	IF() antlr.TerminalNode
	Expr() IExprContext
	Statement() IStatementContext

	// IsElseIfStmtContext differentiates from other interfaces.
	IsElseIfStmtContext()
}

type ElseIfStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseIfStmtContext() *ElseIfStmtContext {
	var p = new(ElseIfStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V4LangGrammarParserRULE_elseIfStmt
	return p
}

func InitEmptyElseIfStmtContext(p *ElseIfStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V4LangGrammarParserRULE_elseIfStmt
}

func (*ElseIfStmtContext) IsElseIfStmtContext() {}

func NewElseIfStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseIfStmtContext {
	var p = new(ElseIfStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = V4LangGrammarParserRULE_elseIfStmt

	return p
}

func (s *ElseIfStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseIfStmtContext) ELSE() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserELSE, 0)
}

func (s *ElseIfStmtContext) IF() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserIF, 0)
}

func (s *ElseIfStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ElseIfStmtContext) Statement() IStatementContext {
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

func (s *ElseIfStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseIfStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElseIfStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.EnterElseIfStmt(s)
	}
}

func (s *ElseIfStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.ExitElseIfStmt(s)
	}
}

func (s *ElseIfStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V4LangGrammarVisitor:
		return t.VisitElseIfStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V4LangGrammarParser) ElseIfStmt() (localctx IElseIfStmtContext) {
	localctx = NewElseIfStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, V4LangGrammarParserRULE_elseIfStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(296)
		p.Match(V4LangGrammarParserELSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(297)
		p.Match(V4LangGrammarParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(298)
		p.expr(0)
	}
	{
		p.SetState(299)
		p.Statement()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IElseStmtContext is an interface to support dynamic dispatch.
type IElseStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ELSE() antlr.TerminalNode
	Statement() IStatementContext

	// IsElseStmtContext differentiates from other interfaces.
	IsElseStmtContext()
}

type ElseStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseStmtContext() *ElseStmtContext {
	var p = new(ElseStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V4LangGrammarParserRULE_elseStmt
	return p
}

func InitEmptyElseStmtContext(p *ElseStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V4LangGrammarParserRULE_elseStmt
}

func (*ElseStmtContext) IsElseStmtContext() {}

func NewElseStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseStmtContext {
	var p = new(ElseStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = V4LangGrammarParserRULE_elseStmt

	return p
}

func (s *ElseStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseStmtContext) ELSE() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserELSE, 0)
}

func (s *ElseStmtContext) Statement() IStatementContext {
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

func (s *ElseStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElseStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.EnterElseStmt(s)
	}
}

func (s *ElseStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.ExitElseStmt(s)
	}
}

func (s *ElseStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V4LangGrammarVisitor:
		return t.VisitElseStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V4LangGrammarParser) ElseStmt() (localctx IElseStmtContext) {
	localctx = NewElseStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, V4LangGrammarParserRULE_elseStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(301)
		p.Match(V4LangGrammarParserELSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(302)
		p.Statement()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISwitchCaseContext is an interface to support dynamic dispatch.
type ISwitchCaseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext
	AllDcl() []IDclContext
	Dcl(i int) IDclContext

	// IsSwitchCaseContext differentiates from other interfaces.
	IsSwitchCaseContext()
}

type SwitchCaseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitchCaseContext() *SwitchCaseContext {
	var p = new(SwitchCaseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V4LangGrammarParserRULE_switchCase
	return p
}

func InitEmptySwitchCaseContext(p *SwitchCaseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V4LangGrammarParserRULE_switchCase
}

func (*SwitchCaseContext) IsSwitchCaseContext() {}

func NewSwitchCaseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwitchCaseContext {
	var p = new(SwitchCaseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = V4LangGrammarParserRULE_switchCase

	return p
}

func (s *SwitchCaseContext) GetParser() antlr.Parser { return s.parser }

func (s *SwitchCaseContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SwitchCaseContext) AllDcl() []IDclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDclContext); ok {
			len++
		}
	}

	tst := make([]IDclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDclContext); ok {
			tst[i] = t.(IDclContext)
			i++
		}
	}

	return tst
}

func (s *SwitchCaseContext) Dcl(i int) IDclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDclContext); ok {
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

	return t.(IDclContext)
}

func (s *SwitchCaseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchCaseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SwitchCaseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.EnterSwitchCase(s)
	}
}

func (s *SwitchCaseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.ExitSwitchCase(s)
	}
}

func (s *SwitchCaseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V4LangGrammarVisitor:
		return t.VisitSwitchCase(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V4LangGrammarParser) SwitchCase() (localctx ISwitchCaseContext) {
	localctx = NewSwitchCaseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, V4LangGrammarParserRULE_switchCase)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(304)
		p.Match(V4LangGrammarParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(305)
		p.expr(0)
	}
	{
		p.SetState(306)
		p.Match(V4LangGrammarParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(310)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-4418313327768238532) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&23) != 0) {
		{
			p.SetState(307)
			p.Dcl()
		}

		p.SetState(312)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDefaultCaseContext is an interface to support dynamic dispatch.
type IDefaultCaseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllDcl() []IDclContext
	Dcl(i int) IDclContext

	// IsDefaultCaseContext differentiates from other interfaces.
	IsDefaultCaseContext()
}

type DefaultCaseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefaultCaseContext() *DefaultCaseContext {
	var p = new(DefaultCaseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V4LangGrammarParserRULE_defaultCase
	return p
}

func InitEmptyDefaultCaseContext(p *DefaultCaseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V4LangGrammarParserRULE_defaultCase
}

func (*DefaultCaseContext) IsDefaultCaseContext() {}

func NewDefaultCaseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefaultCaseContext {
	var p = new(DefaultCaseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = V4LangGrammarParserRULE_defaultCase

	return p
}

func (s *DefaultCaseContext) GetParser() antlr.Parser { return s.parser }

func (s *DefaultCaseContext) AllDcl() []IDclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDclContext); ok {
			len++
		}
	}

	tst := make([]IDclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDclContext); ok {
			tst[i] = t.(IDclContext)
			i++
		}
	}

	return tst
}

func (s *DefaultCaseContext) Dcl(i int) IDclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDclContext); ok {
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

	return t.(IDclContext)
}

func (s *DefaultCaseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefaultCaseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefaultCaseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.EnterDefaultCase(s)
	}
}

func (s *DefaultCaseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.ExitDefaultCase(s)
	}
}

func (s *DefaultCaseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V4LangGrammarVisitor:
		return t.VisitDefaultCase(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V4LangGrammarParser) DefaultCase() (localctx IDefaultCaseContext) {
	localctx = NewDefaultCaseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, V4LangGrammarParserRULE_defaultCase)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(313)
		p.Match(V4LangGrammarParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(314)
		p.Match(V4LangGrammarParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(318)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-4418313327768238532) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&23) != 0) {
		{
			p.SetState(315)
			p.Dcl()
		}

		p.SetState(320)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArgsContext is an interface to support dynamic dispatch.
type IArgsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpr() []IExprContext
	Expr(i int) IExprContext

	// IsArgsContext differentiates from other interfaces.
	IsArgsContext()
}

type ArgsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgsContext() *ArgsContext {
	var p = new(ArgsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V4LangGrammarParserRULE_args
	return p
}

func InitEmptyArgsContext(p *ArgsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V4LangGrammarParserRULE_args
}

func (*ArgsContext) IsArgsContext() {}

func NewArgsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgsContext {
	var p = new(ArgsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = V4LangGrammarParserRULE_args

	return p
}

func (s *ArgsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgsContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ArgsContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *ArgsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.EnterArgs(s)
	}
}

func (s *ArgsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.ExitArgs(s)
	}
}

func (s *ArgsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V4LangGrammarVisitor:
		return t.VisitArgs(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V4LangGrammarParser) Args() (localctx IArgsContext) {
	localctx = NewArgsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, V4LangGrammarParserRULE_args)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(321)
		p.expr(0)
	}
	p.SetState(326)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == V4LangGrammarParserT__0 {
		{
			p.SetState(322)
			p.Match(V4LangGrammarParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(323)
			p.expr(0)
		}

		p.SetState(328)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V4LangGrammarParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V4LangGrammarParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = V4LangGrammarParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyAll(ctx *ExprContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type MulDivContext struct {
	ExprContext
	op antlr.Token
}

func NewMulDivContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MulDivContext {
	var p = new(MulDivContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *MulDivContext) GetOp() antlr.Token { return s.op }

func (s *MulDivContext) SetOp(v antlr.Token) { s.op = v }

func (s *MulDivContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulDivContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *MulDivContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *MulDivContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.EnterMulDiv(s)
	}
}

func (s *MulDivContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.ExitMulDiv(s)
	}
}

func (s *MulDivContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V4LangGrammarVisitor:
		return t.VisitMulDiv(s)

	default:
		return t.VisitChildren(s)
	}
}

type GetValueIndexContext struct {
	ExprContext
}

func NewGetValueIndexContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GetValueIndexContext {
	var p = new(GetValueIndexContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *GetValueIndexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GetValueIndexContext) ID() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserID, 0)
}

func (s *GetValueIndexContext) L_BRACKET() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserL_BRACKET, 0)
}

func (s *GetValueIndexContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *GetValueIndexContext) R_BRACKET() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserR_BRACKET, 0)
}

func (s *GetValueIndexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.EnterGetValueIndex(s)
	}
}

func (s *GetValueIndexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.ExitGetValueIndex(s)
	}
}

func (s *GetValueIndexContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V4LangGrammarVisitor:
		return t.VisitGetValueIndex(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParensContext struct {
	ExprContext
}

func NewParensContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParensContext {
	var p = new(ParensContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ParensContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParensContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserL_PAREN, 0)
}

func (s *ParensContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ParensContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserR_PAREN, 0)
}

func (s *ParensContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.EnterParens(s)
	}
}

func (s *ParensContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.ExitParens(s)
	}
}

func (s *ParensContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V4LangGrammarVisitor:
		return t.VisitParens(s)

	default:
		return t.VisitChildren(s)
	}
}

type LogicalContext struct {
	ExprContext
	op antlr.Token
}

func NewLogicalContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LogicalContext {
	var p = new(LogicalContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *LogicalContext) GetOp() antlr.Token { return s.op }

func (s *LogicalContext) SetOp(v antlr.Token) { s.op = v }

func (s *LogicalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *LogicalContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *LogicalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.EnterLogical(s)
	}
}

func (s *LogicalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.ExitLogical(s)
	}
}

func (s *LogicalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V4LangGrammarVisitor:
		return t.VisitLogical(s)

	default:
		return t.VisitChildren(s)
	}
}

type StringContext struct {
	ExprContext
}

func NewStringContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringContext {
	var p = new(StringContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *StringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringContext) STRING() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserSTRING, 0)
}

func (s *StringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.EnterString(s)
	}
}

func (s *StringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.ExitString(s)
	}
}

func (s *StringContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V4LangGrammarVisitor:
		return t.VisitString(s)

	default:
		return t.VisitChildren(s)
	}
}

type IntContext struct {
	ExprContext
}

func NewIntContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntContext {
	var p = new(IntContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *IntContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntContext) INT() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserINT, 0)
}

func (s *IntContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.EnterInt(s)
	}
}

func (s *IntContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.ExitInt(s)
	}
}

func (s *IntContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V4LangGrammarVisitor:
		return t.VisitInt(s)

	default:
		return t.VisitChildren(s)
	}
}

type IdentifierContext struct {
	ExprContext
}

func NewIdentifierContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdentifierContext {
	var p = new(IdentifierContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) ID() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserID, 0)
}

func (s *IdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.EnterIdentifier(s)
	}
}

func (s *IdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.ExitIdentifier(s)
	}
}

func (s *IdentifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V4LangGrammarVisitor:
		return t.VisitIdentifier(s)

	default:
		return t.VisitChildren(s)
	}
}

type SetValueIndexContext struct {
	ExprContext
	valor IExprContext
}

func NewSetValueIndexContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SetValueIndexContext {
	var p = new(SetValueIndexContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *SetValueIndexContext) GetValor() IExprContext { return s.valor }

func (s *SetValueIndexContext) SetValor(v IExprContext) { s.valor = v }

func (s *SetValueIndexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SetValueIndexContext) ID() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserID, 0)
}

func (s *SetValueIndexContext) L_BRACKET() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserL_BRACKET, 0)
}

func (s *SetValueIndexContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *SetValueIndexContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *SetValueIndexContext) R_BRACKET() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserR_BRACKET, 0)
}

func (s *SetValueIndexContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserASSIGN, 0)
}

func (s *SetValueIndexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.EnterSetValueIndex(s)
	}
}

func (s *SetValueIndexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.ExitSetValueIndex(s)
	}
}

func (s *SetValueIndexContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V4LangGrammarVisitor:
		return t.VisitSetValueIndex(s)

	default:
		return t.VisitChildren(s)
	}
}

type GetValueMatrixContext struct {
	ExprContext
}

func NewGetValueMatrixContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GetValueMatrixContext {
	var p = new(GetValueMatrixContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *GetValueMatrixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GetValueMatrixContext) ID() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserID, 0)
}

func (s *GetValueMatrixContext) AllL_BRACKET() []antlr.TerminalNode {
	return s.GetTokens(V4LangGrammarParserL_BRACKET)
}

func (s *GetValueMatrixContext) L_BRACKET(i int) antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserL_BRACKET, i)
}

func (s *GetValueMatrixContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *GetValueMatrixContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *GetValueMatrixContext) AllR_BRACKET() []antlr.TerminalNode {
	return s.GetTokens(V4LangGrammarParserR_BRACKET)
}

func (s *GetValueMatrixContext) R_BRACKET(i int) antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserR_BRACKET, i)
}

func (s *GetValueMatrixContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.EnterGetValueMatrix(s)
	}
}

func (s *GetValueMatrixContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.ExitGetValueMatrix(s)
	}
}

func (s *GetValueMatrixContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V4LangGrammarVisitor:
		return t.VisitGetValueMatrix(s)

	default:
		return t.VisitChildren(s)
	}
}

type SliceWithValuesContext struct {
	ExprContext
}

func NewSliceWithValuesContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SliceWithValuesContext {
	var p = new(SliceWithValuesContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *SliceWithValuesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SliceWithValuesContext) L_BRACKET() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserL_BRACKET, 0)
}

func (s *SliceWithValuesContext) R_BRACKET() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserR_BRACKET, 0)
}

func (s *SliceWithValuesContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *SliceWithValuesContext) L_CURLY() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserL_CURLY, 0)
}

func (s *SliceWithValuesContext) SliceValues() ISliceValuesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISliceValuesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISliceValuesContext)
}

func (s *SliceWithValuesContext) R_CURLY() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserR_CURLY, 0)
}

func (s *SliceWithValuesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.EnterSliceWithValues(s)
	}
}

func (s *SliceWithValuesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.ExitSliceWithValues(s)
	}
}

func (s *SliceWithValuesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V4LangGrammarVisitor:
		return t.VisitSliceWithValues(s)

	default:
		return t.VisitChildren(s)
	}
}

type EqualityContext struct {
	ExprContext
	op antlr.Token
}

func NewEqualityContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EqualityContext {
	var p = new(EqualityContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *EqualityContext) GetOp() antlr.Token { return s.op }

func (s *EqualityContext) SetOp(v antlr.Token) { s.op = v }

func (s *EqualityContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualityContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *EqualityContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *EqualityContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.EnterEquality(s)
	}
}

func (s *EqualityContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.ExitEquality(s)
	}
}

func (s *EqualityContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V4LangGrammarVisitor:
		return t.VisitEquality(s)

	default:
		return t.VisitChildren(s)
	}
}

type BooleanContext struct {
	ExprContext
}

func NewBooleanContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BooleanContext {
	var p = new(BooleanContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *BooleanContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanContext) BOOL() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserBOOL, 0)
}

func (s *BooleanContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.EnterBoolean(s)
	}
}

func (s *BooleanContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.ExitBoolean(s)
	}
}

func (s *BooleanContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V4LangGrammarVisitor:
		return t.VisitBoolean(s)

	default:
		return t.VisitChildren(s)
	}
}

type NewInstanceContext struct {
	ExprContext
}

func NewNewInstanceContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NewInstanceContext {
	var p = new(NewInstanceContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *NewInstanceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NewInstanceContext) ID() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserID, 0)
}

func (s *NewInstanceContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserL_PAREN, 0)
}

func (s *NewInstanceContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserR_PAREN, 0)
}

func (s *NewInstanceContext) Args() IArgsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgsContext)
}

func (s *NewInstanceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.EnterNewInstance(s)
	}
}

func (s *NewInstanceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.ExitNewInstance(s)
	}
}

func (s *NewInstanceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V4LangGrammarVisitor:
		return t.VisitNewInstance(s)

	default:
		return t.VisitChildren(s)
	}
}

type SetValueMatrixContext struct {
	ExprContext
}

func NewSetValueMatrixContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SetValueMatrixContext {
	var p = new(SetValueMatrixContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *SetValueMatrixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SetValueMatrixContext) ID() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserID, 0)
}

func (s *SetValueMatrixContext) AllL_BRACKET() []antlr.TerminalNode {
	return s.GetTokens(V4LangGrammarParserL_BRACKET)
}

func (s *SetValueMatrixContext) L_BRACKET(i int) antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserL_BRACKET, i)
}

func (s *SetValueMatrixContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *SetValueMatrixContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *SetValueMatrixContext) AllR_BRACKET() []antlr.TerminalNode {
	return s.GetTokens(V4LangGrammarParserR_BRACKET)
}

func (s *SetValueMatrixContext) R_BRACKET(i int) antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserR_BRACKET, i)
}

func (s *SetValueMatrixContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserASSIGN, 0)
}

func (s *SetValueMatrixContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.EnterSetValueMatrix(s)
	}
}

func (s *SetValueMatrixContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.ExitSetValueMatrix(s)
	}
}

func (s *SetValueMatrixContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V4LangGrammarVisitor:
		return t.VisitSetValueMatrix(s)

	default:
		return t.VisitChildren(s)
	}
}

type ModContext struct {
	ExprContext
}

func NewModContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ModContext {
	var p = new(ModContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ModContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ModContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *ModContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.EnterMod(s)
	}
}

func (s *ModContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.ExitMod(s)
	}
}

func (s *ModContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V4LangGrammarVisitor:
		return t.VisitMod(s)

	default:
		return t.VisitChildren(s)
	}
}

type SliceIndexContext struct {
	ExprContext
}

func NewSliceIndexContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SliceIndexContext {
	var p = new(SliceIndexContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *SliceIndexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SliceIndexContext) SLICE_INDEX() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserSLICE_INDEX, 0)
}

func (s *SliceIndexContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserL_PAREN, 0)
}

func (s *SliceIndexContext) ID() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserID, 0)
}

func (s *SliceIndexContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SliceIndexContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserR_PAREN, 0)
}

func (s *SliceIndexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.EnterSliceIndex(s)
	}
}

func (s *SliceIndexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.ExitSliceIndex(s)
	}
}

func (s *SliceIndexContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V4LangGrammarVisitor:
		return t.VisitSliceIndex(s)

	default:
		return t.VisitChildren(s)
	}
}

type TypeOfContext struct {
	ExprContext
}

func NewTypeOfContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeOfContext {
	var p = new(TypeOfContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *TypeOfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeOfContext) TYPEOF() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserTYPEOF, 0)
}

func (s *TypeOfContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserL_PAREN, 0)
}

func (s *TypeOfContext) ID() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserID, 0)
}

func (s *TypeOfContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserR_PAREN, 0)
}

func (s *TypeOfContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.EnterTypeOf(s)
	}
}

func (s *TypeOfContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.ExitTypeOf(s)
	}
}

func (s *TypeOfContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V4LangGrammarVisitor:
		return t.VisitTypeOf(s)

	default:
		return t.VisitChildren(s)
	}
}

type AddSubContext struct {
	ExprContext
	op antlr.Token
}

func NewAddSubContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddSubContext {
	var p = new(AddSubContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *AddSubContext) GetOp() antlr.Token { return s.op }

func (s *AddSubContext) SetOp(v antlr.Token) { s.op = v }

func (s *AddSubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddSubContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *AddSubContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *AddSubContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.EnterAddSub(s)
	}
}

func (s *AddSubContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.ExitAddSub(s)
	}
}

func (s *AddSubContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V4LangGrammarVisitor:
		return t.VisitAddSub(s)

	default:
		return t.VisitChildren(s)
	}
}

type StructInstanceContext struct {
	ExprContext
}

func NewStructInstanceContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructInstanceContext {
	var p = new(StructInstanceContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *StructInstanceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructInstanceContext) ID() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserID, 0)
}

func (s *StructInstanceContext) L_CURLY() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserL_CURLY, 0)
}

func (s *StructInstanceContext) InstanceValues() IInstanceValuesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstanceValuesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstanceValuesContext)
}

func (s *StructInstanceContext) R_CURLY() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserR_CURLY, 0)
}

func (s *StructInstanceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.EnterStructInstance(s)
	}
}

func (s *StructInstanceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.ExitStructInstance(s)
	}
}

func (s *StructInstanceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V4LangGrammarVisitor:
		return t.VisitStructInstance(s)

	default:
		return t.VisitChildren(s)
	}
}

type RelationalContext struct {
	ExprContext
	op antlr.Token
}

func NewRelationalContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RelationalContext {
	var p = new(RelationalContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *RelationalContext) GetOp() antlr.Token { return s.op }

func (s *RelationalContext) SetOp(v antlr.Token) { s.op = v }

func (s *RelationalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationalContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *RelationalContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *RelationalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.EnterRelational(s)
	}
}

func (s *RelationalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.ExitRelational(s)
	}
}

func (s *RelationalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V4LangGrammarVisitor:
		return t.VisitRelational(s)

	default:
		return t.VisitChildren(s)
	}
}

type CallStmtContext struct {
	ExprContext
}

func NewCallStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CallStmtContext {
	var p = new(CallStmtContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *CallStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *CallStmtContext) AllCall() []ICallContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICallContext); ok {
			len++
		}
	}

	tst := make([]ICallContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICallContext); ok {
			tst[i] = t.(ICallContext)
			i++
		}
	}

	return tst
}

func (s *CallStmtContext) Call(i int) ICallContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICallContext); ok {
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

	return t.(ICallContext)
}

func (s *CallStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.EnterCallStmt(s)
	}
}

func (s *CallStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.ExitCallStmt(s)
	}
}

func (s *CallStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V4LangGrammarVisitor:
		return t.VisitCallStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

type NilContext struct {
	ExprContext
}

func NewNilContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NilContext {
	var p = new(NilContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *NilContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NilContext) NIL() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserNIL, 0)
}

func (s *NilContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.EnterNil(s)
	}
}

func (s *NilContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.ExitNil(s)
	}
}

func (s *NilContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V4LangGrammarVisitor:
		return t.VisitNil(s)

	default:
		return t.VisitChildren(s)
	}
}

type StringJoinContext struct {
	ExprContext
}

func NewStringJoinContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringJoinContext {
	var p = new(StringJoinContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *StringJoinContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringJoinContext) JOIN() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserJOIN, 0)
}

func (s *StringJoinContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserL_PAREN, 0)
}

func (s *StringJoinContext) ID() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserID, 0)
}

func (s *StringJoinContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *StringJoinContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserR_PAREN, 0)
}

func (s *StringJoinContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.EnterStringJoin(s)
	}
}

func (s *StringJoinContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.ExitStringJoin(s)
	}
}

func (s *StringJoinContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V4LangGrammarVisitor:
		return t.VisitStringJoin(s)

	default:
		return t.VisitChildren(s)
	}
}

type FloatContext struct {
	ExprContext
}

func NewFloatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FloatContext {
	var p = new(FloatContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *FloatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserFLOAT, 0)
}

func (s *FloatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.EnterFloat(s)
	}
}

func (s *FloatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.ExitFloat(s)
	}
}

func (s *FloatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V4LangGrammarVisitor:
		return t.VisitFloat(s)

	default:
		return t.VisitChildren(s)
	}
}

type NotContext struct {
	ExprContext
}

func NewNotContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotContext {
	var p = new(NotContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *NotContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *NotContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.EnterNot(s)
	}
}

func (s *NotContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.ExitNot(s)
	}
}

func (s *NotContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V4LangGrammarVisitor:
		return t.VisitNot(s)

	default:
		return t.VisitChildren(s)
	}
}

type AtoiContext struct {
	ExprContext
}

func NewAtoiContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AtoiContext {
	var p = new(AtoiContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *AtoiContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtoiContext) ATOI() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserATOI, 0)
}

func (s *AtoiContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserL_PAREN, 0)
}

func (s *AtoiContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AtoiContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserR_PAREN, 0)
}

func (s *AtoiContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.EnterAtoi(s)
	}
}

func (s *AtoiContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.ExitAtoi(s)
	}
}

func (s *AtoiContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V4LangGrammarVisitor:
		return t.VisitAtoi(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParseFloatContext struct {
	ExprContext
}

func NewParseFloatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParseFloatContext {
	var p = new(ParseFloatContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ParseFloatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParseFloatContext) PARSEFLOAT() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserPARSEFLOAT, 0)
}

func (s *ParseFloatContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserL_PAREN, 0)
}

func (s *ParseFloatContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ParseFloatContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserR_PAREN, 0)
}

func (s *ParseFloatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.EnterParseFloat(s)
	}
}

func (s *ParseFloatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.ExitParseFloat(s)
	}
}

func (s *ParseFloatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V4LangGrammarVisitor:
		return t.VisitParseFloat(s)

	default:
		return t.VisitChildren(s)
	}
}

type AppendContext struct {
	ExprContext
}

func NewAppendContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AppendContext {
	var p = new(AppendContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *AppendContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AppendContext) APPEND() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserAPPEND, 0)
}

func (s *AppendContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserL_PAREN, 0)
}

func (s *AppendContext) ID() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserID, 0)
}

func (s *AppendContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AppendContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserR_PAREN, 0)
}

func (s *AppendContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.EnterAppend(s)
	}
}

func (s *AppendContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.ExitAppend(s)
	}
}

func (s *AppendContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V4LangGrammarVisitor:
		return t.VisitAppend(s)

	default:
		return t.VisitChildren(s)
	}
}

type LenContext struct {
	ExprContext
}

func NewLenContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LenContext {
	var p = new(LenContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *LenContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LenContext) LEN() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserLEN, 0)
}

func (s *LenContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserL_PAREN, 0)
}

func (s *LenContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *LenContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserR_PAREN, 0)
}

func (s *LenContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.EnterLen(s)
	}
}

func (s *LenContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.ExitLen(s)
	}
}

func (s *LenContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V4LangGrammarVisitor:
		return t.VisitLen(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignArithmeticContext struct {
	ExprContext
	op antlr.Token
}

func NewAssignArithmeticContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignArithmeticContext {
	var p = new(AssignArithmeticContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *AssignArithmeticContext) GetOp() antlr.Token { return s.op }

func (s *AssignArithmeticContext) SetOp(v antlr.Token) { s.op = v }

func (s *AssignArithmeticContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignArithmeticContext) ID() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserID, 0)
}

func (s *AssignArithmeticContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AssignArithmeticContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.EnterAssignArithmetic(s)
	}
}

func (s *AssignArithmeticContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.ExitAssignArithmetic(s)
	}
}

func (s *AssignArithmeticContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V4LangGrammarVisitor:
		return t.VisitAssignArithmetic(s)

	default:
		return t.VisitChildren(s)
	}
}

type SliceContext struct {
	ExprContext
}

func NewSliceContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SliceContext {
	var p = new(SliceContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *SliceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SliceContext) L_CURLY() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserL_CURLY, 0)
}

func (s *SliceContext) R_CURLY() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserR_CURLY, 0)
}

func (s *SliceContext) AllSliceValues() []ISliceValuesContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISliceValuesContext); ok {
			len++
		}
	}

	tst := make([]ISliceValuesContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISliceValuesContext); ok {
			tst[i] = t.(ISliceValuesContext)
			i++
		}
	}

	return tst
}

func (s *SliceContext) SliceValues(i int) ISliceValuesContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISliceValuesContext); ok {
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

	return t.(ISliceValuesContext)
}

func (s *SliceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.EnterSlice(s)
	}
}

func (s *SliceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.ExitSlice(s)
	}
}

func (s *SliceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V4LangGrammarVisitor:
		return t.VisitSlice(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignContext struct {
	ExprContext
}

func NewAssignContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignContext {
	var p = new(AssignContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *AssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *AssignContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *AssignContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserASSIGN, 0)
}

func (s *AssignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.EnterAssign(s)
	}
}

func (s *AssignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.ExitAssign(s)
	}
}

func (s *AssignContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V4LangGrammarVisitor:
		return t.VisitAssign(s)

	default:
		return t.VisitChildren(s)
	}
}

type NegateContext struct {
	ExprContext
}

func NewNegateContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NegateContext {
	var p = new(NegateContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *NegateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NegateContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *NegateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.EnterNegate(s)
	}
}

func (s *NegateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.ExitNegate(s)
	}
}

func (s *NegateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V4LangGrammarVisitor:
		return t.VisitNegate(s)

	default:
		return t.VisitChildren(s)
	}
}

type IncrementDecrementContext struct {
	ExprContext
	op antlr.Token
}

func NewIncrementDecrementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IncrementDecrementContext {
	var p = new(IncrementDecrementContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *IncrementDecrementContext) GetOp() antlr.Token { return s.op }

func (s *IncrementDecrementContext) SetOp(v antlr.Token) { s.op = v }

func (s *IncrementDecrementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IncrementDecrementContext) ID() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserID, 0)
}

func (s *IncrementDecrementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.EnterIncrementDecrement(s)
	}
}

func (s *IncrementDecrementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.ExitIncrementDecrement(s)
	}
}

func (s *IncrementDecrementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V4LangGrammarVisitor:
		return t.VisitIncrementDecrement(s)

	default:
		return t.VisitChildren(s)
	}
}

type RuneContext struct {
	ExprContext
}

func NewRuneContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RuneContext {
	var p = new(RuneContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *RuneContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RuneContext) RUNE() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserRUNE, 0)
}

func (s *RuneContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.EnterRune(s)
	}
}

func (s *RuneContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.ExitRune(s)
	}
}

func (s *RuneContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V4LangGrammarVisitor:
		return t.VisitRune(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V4LangGrammarParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *V4LangGrammarParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 36
	p.EnterRecursionRule(localctx, 36, V4LangGrammarParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(448)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 39, p.GetParserRuleContext()) {
	case 1:
		localctx = NewNegateContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(330)
			p.Match(V4LangGrammarParserT__8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(331)
			p.expr(35)
		}

	case 2:
		localctx = NewNotContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(332)
			p.Match(V4LangGrammarParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(333)
			p.expr(34)
		}

	case 3:
		localctx = NewAssignArithmeticContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(334)
			p.Match(V4LangGrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(335)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*AssignArithmeticContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&31457280) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*AssignArithmeticContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(336)
			p.expr(27)
		}

	case 4:
		localctx = NewNewInstanceContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(337)
			p.Match(V4LangGrammarParserT__26)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(338)
			p.Match(V4LangGrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(339)
			p.Match(V4LangGrammarParserL_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(341)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64((_la-9)) & ^0x3f) == 0 && ((int64(1)<<(_la-9))&856055836861267971) != 0 {
			{
				p.SetState(340)
				p.Args()
			}

		}
		{
			p.SetState(343)
			p.Match(V4LangGrammarParserR_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		localctx = NewStructInstanceContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(344)
			p.Match(V4LangGrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(345)
			p.Match(V4LangGrammarParserL_CURLY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(346)
			p.InstanceValues()
		}
		{
			p.SetState(347)
			p.Match(V4LangGrammarParserR_CURLY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		localctx = NewSetValueMatrixContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(349)
			p.Match(V4LangGrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(350)
			p.Match(V4LangGrammarParserL_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(351)
			p.expr(0)
		}
		{
			p.SetState(352)
			p.Match(V4LangGrammarParserR_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(353)
			p.Match(V4LangGrammarParserL_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(354)
			p.expr(0)
		}
		{
			p.SetState(355)
			p.Match(V4LangGrammarParserR_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(356)
			p.Match(V4LangGrammarParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(357)
			p.expr(23)
		}

	case 7:
		localctx = NewGetValueMatrixContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(359)
			p.Match(V4LangGrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(360)
			p.Match(V4LangGrammarParserL_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(361)
			p.expr(0)
		}
		{
			p.SetState(362)
			p.Match(V4LangGrammarParserR_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(363)
			p.Match(V4LangGrammarParserL_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(364)
			p.expr(0)
		}
		p.SetState(366)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 37, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(365)
				p.Match(V4LangGrammarParserR_BRACKET)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case 8:
		localctx = NewSetValueIndexContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(368)
			p.Match(V4LangGrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(369)
			p.Match(V4LangGrammarParserL_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(370)
			p.expr(0)
		}
		{
			p.SetState(371)
			p.Match(V4LangGrammarParserR_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(372)
			p.Match(V4LangGrammarParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(373)

			var _x = p.expr(21)

			localctx.(*SetValueIndexContext).valor = _x
		}

	case 9:
		localctx = NewGetValueIndexContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(375)
			p.Match(V4LangGrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(376)
			p.Match(V4LangGrammarParserL_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(377)
			p.expr(0)
		}
		{
			p.SetState(378)
			p.Match(V4LangGrammarParserR_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 10:
		localctx = NewLenContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(380)
			p.Match(V4LangGrammarParserLEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(381)
			p.Match(V4LangGrammarParserL_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(382)
			p.expr(0)
		}
		{
			p.SetState(383)
			p.Match(V4LangGrammarParserR_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 11:
		localctx = NewStringJoinContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(385)
			p.Match(V4LangGrammarParserJOIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(386)
			p.Match(V4LangGrammarParserL_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(387)
			p.Match(V4LangGrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(388)
			p.Match(V4LangGrammarParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(389)
			p.expr(0)
		}
		{
			p.SetState(390)
			p.Match(V4LangGrammarParserR_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 12:
		localctx = NewAtoiContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(392)
			p.Match(V4LangGrammarParserATOI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(393)
			p.Match(V4LangGrammarParserL_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(394)
			p.expr(0)
		}
		{
			p.SetState(395)
			p.Match(V4LangGrammarParserR_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 13:
		localctx = NewParseFloatContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(397)
			p.Match(V4LangGrammarParserPARSEFLOAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(398)
			p.Match(V4LangGrammarParserL_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(399)
			p.expr(0)
		}
		{
			p.SetState(400)
			p.Match(V4LangGrammarParserR_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 14:
		localctx = NewTypeOfContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(402)
			p.Match(V4LangGrammarParserTYPEOF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(403)
			p.Match(V4LangGrammarParserL_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(404)
			p.Match(V4LangGrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(405)
			p.Match(V4LangGrammarParserR_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 15:
		localctx = NewSliceIndexContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(406)
			p.Match(V4LangGrammarParserSLICE_INDEX)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(407)
			p.Match(V4LangGrammarParserL_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(408)
			p.Match(V4LangGrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(409)
			p.Match(V4LangGrammarParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(410)
			p.expr(0)
		}
		{
			p.SetState(411)
			p.Match(V4LangGrammarParserR_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 16:
		localctx = NewIncrementDecrementContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(413)
			p.Match(V4LangGrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(414)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*IncrementDecrementContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == V4LangGrammarParserT__27 || _la == V4LangGrammarParserT__28) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*IncrementDecrementContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case 17:
		localctx = NewIdentifierContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(415)
			p.Match(V4LangGrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 18:
		localctx = NewRuneContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(416)
			p.Match(V4LangGrammarParserRUNE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 19:
		localctx = NewBooleanContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(417)
			p.Match(V4LangGrammarParserBOOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 20:
		localctx = NewFloatContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(418)
			p.Match(V4LangGrammarParserFLOAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 21:
		localctx = NewStringContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(419)
			p.Match(V4LangGrammarParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 22:
		localctx = NewIntContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(420)
			p.Match(V4LangGrammarParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 23:
		localctx = NewSliceContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(421)
			p.Match(V4LangGrammarParserL_CURLY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(425)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64((_la-9)) & ^0x3f) == 0 && ((int64(1)<<(_la-9))&856055836861267971) != 0 {
			{
				p.SetState(422)
				p.SliceValues()
			}

			p.SetState(427)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(428)
			p.Match(V4LangGrammarParserR_CURLY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 24:
		localctx = NewSliceWithValuesContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(429)
			p.Match(V4LangGrammarParserL_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(430)
			p.Match(V4LangGrammarParserR_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(431)
			p.Type_()
		}
		{
			p.SetState(432)
			p.Match(V4LangGrammarParserL_CURLY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(433)
			p.SliceValues()
		}
		{
			p.SetState(434)
			p.Match(V4LangGrammarParserR_CURLY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 25:
		localctx = NewAppendContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(436)
			p.Match(V4LangGrammarParserAPPEND)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(437)
			p.Match(V4LangGrammarParserL_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(438)
			p.Match(V4LangGrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(439)
			p.Match(V4LangGrammarParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(440)
			p.expr(0)
		}
		{
			p.SetState(441)
			p.Match(V4LangGrammarParserR_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 26:
		localctx = NewNilContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(443)
			p.Match(V4LangGrammarParserNIL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 27:
		localctx = NewParensContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(444)
			p.Match(V4LangGrammarParserL_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(445)
			p.expr(0)
		}
		{
			p.SetState(446)
			p.Match(V4LangGrammarParserR_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(479)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 42, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(477)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 41, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMulDivContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, V4LangGrammarParserRULE_expr)
				p.SetState(450)

				if !(p.Precpred(p.GetParserRuleContext(), 32)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 32)", ""))
					goto errorExit
				}
				{
					p.SetState(451)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*MulDivContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == V4LangGrammarParserT__10 || _la == V4LangGrammarParserT__11) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*MulDivContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(452)
					p.expr(33)
				}

			case 2:
				localctx = NewAddSubContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, V4LangGrammarParserRULE_expr)
				p.SetState(453)

				if !(p.Precpred(p.GetParserRuleContext(), 31)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 31)", ""))
					goto errorExit
				}
				{
					p.SetState(454)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*AddSubContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == V4LangGrammarParserT__8 || _la == V4LangGrammarParserT__12) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*AddSubContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(455)
					p.expr(32)
				}

			case 3:
				localctx = NewRelationalContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, V4LangGrammarParserRULE_expr)
				p.SetState(456)

				if !(p.Precpred(p.GetParserRuleContext(), 30)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 30)", ""))
					goto errorExit
				}
				{
					p.SetState(457)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*RelationalContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&245760) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*RelationalContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(458)
					p.expr(31)
				}

			case 4:
				localctx = NewModContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, V4LangGrammarParserRULE_expr)
				p.SetState(459)

				if !(p.Precpred(p.GetParserRuleContext(), 29)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 29)", ""))
					goto errorExit
				}
				{
					p.SetState(460)
					p.Match(V4LangGrammarParserT__17)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(461)
					p.expr(30)
				}

			case 5:
				localctx = NewEqualityContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, V4LangGrammarParserRULE_expr)
				p.SetState(462)

				if !(p.Precpred(p.GetParserRuleContext(), 28)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 28)", ""))
					goto errorExit
				}
				{
					p.SetState(463)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*EqualityContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == V4LangGrammarParserT__18 || _la == V4LangGrammarParserT__19) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*EqualityContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(464)
					p.expr(29)
				}

			case 6:
				localctx = NewLogicalContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, V4LangGrammarParserRULE_expr)
				p.SetState(465)

				if !(p.Precpred(p.GetParserRuleContext(), 26)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 26)", ""))
					goto errorExit
				}
				{
					p.SetState(466)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*LogicalContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == V4LangGrammarParserT__24 || _la == V4LangGrammarParserT__25) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*LogicalContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(467)
					p.expr(27)
				}

			case 7:
				localctx = NewAssignContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, V4LangGrammarParserRULE_expr)
				p.SetState(468)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
					goto errorExit
				}
				{
					p.SetState(469)
					p.Match(V4LangGrammarParserASSIGN)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(470)
					p.expr(14)
				}

			case 8:
				localctx = NewCallStmtContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, V4LangGrammarParserRULE_expr)
				p.SetState(471)

				if !(p.Precpred(p.GetParserRuleContext(), 33)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 33)", ""))
					goto errorExit
				}
				p.SetState(473)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_alt = 1
				for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
					switch _alt {
					case 1:
						{
							p.SetState(472)
							p.Call()
						}

					default:
						p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
						goto errorExit
					}

					p.SetState(475)
					p.GetErrorHandler().Sync(p)
					_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 40, p.GetParserRuleContext())
					if p.HasError() {
						goto errorExit
					}
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(481)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 42, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICallContext is an interface to support dynamic dispatch.
type ICallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsCallContext differentiates from other interfaces.
	IsCallContext()
}

type CallContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCallContext() *CallContext {
	var p = new(CallContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V4LangGrammarParserRULE_call
	return p
}

func InitEmptyCallContext(p *CallContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V4LangGrammarParserRULE_call
}

func (*CallContext) IsCallContext() {}

func NewCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CallContext {
	var p = new(CallContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = V4LangGrammarParserRULE_call

	return p
}

func (s *CallContext) GetParser() antlr.Parser { return s.parser }

func (s *CallContext) CopyAll(ctx *CallContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *CallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FuncCallContext struct {
	CallContext
}

func NewFuncCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FuncCallContext {
	var p = new(FuncCallContext)

	InitEmptyCallContext(&p.CallContext)
	p.parser = parser
	p.CopyAll(ctx.(*CallContext))

	return p
}

func (s *FuncCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncCallContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserL_PAREN, 0)
}

func (s *FuncCallContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserR_PAREN, 0)
}

func (s *FuncCallContext) Args() IArgsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgsContext)
}

func (s *FuncCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.EnterFuncCall(s)
	}
}

func (s *FuncCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.ExitFuncCall(s)
	}
}

func (s *FuncCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V4LangGrammarVisitor:
		return t.VisitFuncCall(s)

	default:
		return t.VisitChildren(s)
	}
}

type GetContext struct {
	CallContext
}

func NewGetContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GetContext {
	var p = new(GetContext)

	InitEmptyCallContext(&p.CallContext)
	p.parser = parser
	p.CopyAll(ctx.(*CallContext))

	return p
}

func (s *GetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GetContext) ID() antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserID, 0)
}

func (s *GetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.EnterGet(s)
	}
}

func (s *GetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.ExitGet(s)
	}
}

func (s *GetContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V4LangGrammarVisitor:
		return t.VisitGet(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V4LangGrammarParser) Call() (localctx ICallContext) {
	localctx = NewCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, V4LangGrammarParserRULE_call)
	var _la int

	p.SetState(489)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case V4LangGrammarParserL_PAREN:
		localctx = NewFuncCallContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(482)
			p.Match(V4LangGrammarParserL_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(484)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64((_la-9)) & ^0x3f) == 0 && ((int64(1)<<(_la-9))&856055836861267971) != 0 {
			{
				p.SetState(483)
				p.Args()
			}

		}
		{
			p.SetState(486)
			p.Match(V4LangGrammarParserR_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case V4LangGrammarParserT__29:
		localctx = NewGetContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(487)
			p.Match(V4LangGrammarParserT__29)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(488)
			p.Match(V4LangGrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInstanceValuesContext is an interface to support dynamic dispatch.
type IInstanceValuesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext

	// IsInstanceValuesContext differentiates from other interfaces.
	IsInstanceValuesContext()
}

type InstanceValuesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInstanceValuesContext() *InstanceValuesContext {
	var p = new(InstanceValuesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V4LangGrammarParserRULE_instanceValues
	return p
}

func InitEmptyInstanceValuesContext(p *InstanceValuesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V4LangGrammarParserRULE_instanceValues
}

func (*InstanceValuesContext) IsInstanceValuesContext() {}

func NewInstanceValuesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstanceValuesContext {
	var p = new(InstanceValuesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = V4LangGrammarParserRULE_instanceValues

	return p
}

func (s *InstanceValuesContext) GetParser() antlr.Parser { return s.parser }

func (s *InstanceValuesContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(V4LangGrammarParserID)
}

func (s *InstanceValuesContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(V4LangGrammarParserID, i)
}

func (s *InstanceValuesContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *InstanceValuesContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *InstanceValuesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstanceValuesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstanceValuesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.EnterInstanceValues(s)
	}
}

func (s *InstanceValuesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.ExitInstanceValues(s)
	}
}

func (s *InstanceValuesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V4LangGrammarVisitor:
		return t.VisitInstanceValues(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V4LangGrammarParser) InstanceValues() (localctx IInstanceValuesContext) {
	localctx = NewInstanceValuesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, V4LangGrammarParserRULE_instanceValues)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(491)
		p.Match(V4LangGrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(492)
		p.Match(V4LangGrammarParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(493)
		p.expr(0)
	}
	p.SetState(500)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 45, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(494)
				p.Match(V4LangGrammarParserT__0)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(495)
				p.Match(V4LangGrammarParserID)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(496)
				p.Match(V4LangGrammarParserT__6)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(497)
				p.expr(0)
			}

		}
		p.SetState(502)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 45, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(504)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == V4LangGrammarParserT__0 {
		{
			p.SetState(503)
			p.Match(V4LangGrammarParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypeContext is an interface to support dynamic dispatch.
type ITypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsTypeContext differentiates from other interfaces.
	IsTypeContext()
}

type TypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeContext() *TypeContext {
	var p = new(TypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V4LangGrammarParserRULE_type
	return p
}

func InitEmptyTypeContext(p *TypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V4LangGrammarParserRULE_type
}

func (*TypeContext) IsTypeContext() {}

func NewTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeContext {
	var p = new(TypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = V4LangGrammarParserRULE_type

	return p
}

func (s *TypeContext) GetParser() antlr.Parser { return s.parser }
func (s *TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.EnterType(s)
	}
}

func (s *TypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V4LangGrammarListener); ok {
		listenerT.ExitType(s)
	}
}

func (s *TypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V4LangGrammarVisitor:
		return t.VisitType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V4LangGrammarParser) Type_() (localctx ITypeContext) {
	localctx = NewTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, V4LangGrammarParserRULE_type)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(506)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&66571993088) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *V4LangGrammarParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 18:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *V4LangGrammarParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 32)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 31)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 30)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 29)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 28)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 26)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 33)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
