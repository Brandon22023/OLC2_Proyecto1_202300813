// Code generated from parser/Vlang.g4 by ANTLR 4.13.2. DO NOT EDIT.

package compiler // Vlang
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

type VlangParser struct {
	*antlr.BaseParser
}

var VlangParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func vlangParserInit() {
	staticData := &VlangParserStaticData
	staticData.LiteralNames = []string{
		"", "'fn'", "'main'", "'mut'", "'case'", "'default'", "'while'", "",
		"", "'len'", "'cap'", "'append'", "'if'", "'else'", "'for'", "'switch'",
		"'indexOf'", "'join'", "", "", "", "", "", "'int'", "'float64'", "'string'",
		"'bool'", "'rune'", "'print'", "", "'++'", "'--'", "'+='", "'-='", "'+'",
		"'-'", "'*'", "'/'", "'%'", "'!'", "'||'", "'&&'", "'=='", "'!='", "'<='",
		"'>='", "'<'", "'>'", "'='", "'('", "')'", "'['", "']'", "'{'", "'}'",
		"';'", "':'", "'.'", "','",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "TIPO", "CONDICIONAL_ETC", "LEN", "CAP",
		"APPEND", "IF", "ELSE", "FOR", "SWITCH", "INDEXOF", "JOIN", "BOOLEANO",
		"ENTERO", "DECIMAL", "CADENA", "CARACTER", "TIPO_ENTERO", "TIPO_DECIMAL",
		"TIPO_CADENA", "TIPO_BOOLEANO", "TIPO_CHAR", "PRINT", "ID", "INC", "DEC",
		"SUMAIMPLICITA", "RESTOIMPLICITO", "PLUS", "MINUS", "MUL", "DIV", "MOD",
		"NOT", "OR", "AND", "EQ", "NEQ", "LE", "GE", "LT", "GT", "ASSIGN", "LPAREN",
		"RPAREN", "LBRACK", "RBRACK", "LBRACE", "RBRACE", "SEMICOLON", "COLON",
		"DOT", "COMMA", "WS", "LINE_COMMENT", "BLOCK_COMMENT",
	}
	staticData.RuleNames = []string{
		"programa", "funcMain", "block", "declaraciones", "varDcl", "tipoSlice",
		"stmt", "sentencias_control", "ifDcl", "forDcl", "switchDcl", "caseBlock",
		"defaultBlock", "llamadaFuncion", "whileDcl", "expresion", "parametros",
		"valores", "valor", "listaExpresiones", "incredecre",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 61, 336, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 1,
		0, 5, 0, 44, 8, 0, 10, 0, 12, 0, 47, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 5, 2, 59, 8, 2, 10, 2, 12, 2, 62, 9, 2, 1,
		2, 1, 2, 1, 3, 1, 3, 3, 3, 68, 8, 3, 1, 4, 1, 4, 1, 4, 3, 4, 73, 8, 4,
		1, 4, 1, 4, 3, 4, 77, 8, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 5, 6, 88, 8, 6, 10, 6, 12, 6, 91, 9, 6, 3, 6, 93, 8, 6, 1, 6,
		1, 6, 1, 6, 3, 6, 98, 8, 6, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 104, 8, 7, 1,
		8, 1, 8, 3, 8, 108, 8, 8, 1, 8, 1, 8, 3, 8, 112, 8, 8, 1, 8, 1, 8, 5, 8,
		116, 8, 8, 10, 8, 12, 8, 119, 9, 8, 1, 8, 1, 8, 1, 8, 1, 8, 5, 8, 125,
		8, 8, 10, 8, 12, 8, 128, 9, 8, 1, 8, 3, 8, 131, 8, 8, 1, 9, 1, 9, 3, 9,
		135, 8, 9, 1, 9, 3, 9, 138, 8, 9, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 144, 8,
		9, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 150, 8, 9, 1, 9, 1, 9, 3, 9, 154, 8, 9,
		1, 9, 1, 9, 3, 9, 158, 8, 9, 1, 10, 1, 10, 1, 10, 1, 10, 5, 10, 164, 8,
		10, 10, 10, 12, 10, 167, 9, 10, 1, 10, 3, 10, 170, 8, 10, 1, 10, 1, 10,
		1, 11, 1, 11, 1, 11, 1, 11, 5, 11, 178, 8, 11, 10, 11, 12, 11, 181, 9,
		11, 1, 12, 1, 12, 1, 12, 5, 12, 186, 8, 12, 10, 12, 12, 12, 189, 9, 12,
		1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 5, 13, 196, 8, 13, 10, 13, 12, 13, 199,
		9, 13, 3, 13, 201, 8, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 5,
		13, 209, 8, 13, 10, 13, 12, 13, 212, 9, 13, 3, 13, 214, 8, 13, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 13, 5, 13, 222, 8, 13, 10, 13, 12, 13, 225,
		9, 13, 3, 13, 227, 8, 13, 1, 13, 3, 13, 230, 8, 13, 1, 14, 1, 14, 1, 14,
		1, 14, 1, 14, 1, 14, 5, 14, 238, 8, 14, 10, 14, 12, 14, 241, 9, 14, 1,
		14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15,
		1, 15, 1, 15, 1, 15, 1, 15, 3, 15, 258, 8, 15, 1, 15, 1, 15, 1, 15, 1,
		15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15,
		1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 3, 15, 280, 8, 15, 1, 15, 1,
		15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15,
		1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 5, 15, 300, 8, 15, 10, 15, 12,
		15, 303, 9, 15, 1, 16, 1, 16, 1, 16, 5, 16, 308, 8, 16, 10, 16, 12, 16,
		311, 9, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 3, 18, 320,
		8, 18, 1, 19, 1, 19, 1, 19, 5, 19, 325, 8, 19, 10, 19, 12, 19, 328, 9,
		19, 1, 20, 1, 20, 1, 20, 1, 20, 3, 20, 334, 8, 20, 1, 20, 0, 1, 30, 21,
		0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36,
		38, 40, 0, 7, 2, 0, 35, 35, 39, 39, 1, 0, 32, 33, 1, 0, 40, 41, 1, 0, 36,
		38, 1, 0, 34, 35, 1, 0, 44, 47, 1, 0, 42, 43, 375, 0, 45, 1, 0, 0, 0, 2,
		50, 1, 0, 0, 0, 4, 56, 1, 0, 0, 0, 6, 67, 1, 0, 0, 0, 8, 69, 1, 0, 0, 0,
		10, 78, 1, 0, 0, 0, 12, 97, 1, 0, 0, 0, 14, 103, 1, 0, 0, 0, 16, 105, 1,
		0, 0, 0, 18, 157, 1, 0, 0, 0, 20, 159, 1, 0, 0, 0, 22, 173, 1, 0, 0, 0,
		24, 182, 1, 0, 0, 0, 26, 229, 1, 0, 0, 0, 28, 231, 1, 0, 0, 0, 30, 279,
		1, 0, 0, 0, 32, 304, 1, 0, 0, 0, 34, 312, 1, 0, 0, 0, 36, 319, 1, 0, 0,
		0, 38, 321, 1, 0, 0, 0, 40, 333, 1, 0, 0, 0, 42, 44, 3, 2, 1, 0, 43, 42,
		1, 0, 0, 0, 44, 47, 1, 0, 0, 0, 45, 43, 1, 0, 0, 0, 45, 46, 1, 0, 0, 0,
		46, 48, 1, 0, 0, 0, 47, 45, 1, 0, 0, 0, 48, 49, 5, 0, 0, 1, 49, 1, 1, 0,
		0, 0, 50, 51, 5, 1, 0, 0, 51, 52, 5, 2, 0, 0, 52, 53, 5, 49, 0, 0, 53,
		54, 5, 50, 0, 0, 54, 55, 3, 4, 2, 0, 55, 3, 1, 0, 0, 0, 56, 60, 5, 53,
		0, 0, 57, 59, 3, 6, 3, 0, 58, 57, 1, 0, 0, 0, 59, 62, 1, 0, 0, 0, 60, 58,
		1, 0, 0, 0, 60, 61, 1, 0, 0, 0, 61, 63, 1, 0, 0, 0, 62, 60, 1, 0, 0, 0,
		63, 64, 5, 54, 0, 0, 64, 5, 1, 0, 0, 0, 65, 68, 3, 8, 4, 0, 66, 68, 3,
		12, 6, 0, 67, 65, 1, 0, 0, 0, 67, 66, 1, 0, 0, 0, 68, 7, 1, 0, 0, 0, 69,
		70, 5, 3, 0, 0, 70, 72, 5, 29, 0, 0, 71, 73, 5, 7, 0, 0, 72, 71, 1, 0,
		0, 0, 72, 73, 1, 0, 0, 0, 73, 76, 1, 0, 0, 0, 74, 75, 5, 48, 0, 0, 75,
		77, 3, 30, 15, 0, 76, 74, 1, 0, 0, 0, 76, 77, 1, 0, 0, 0, 77, 9, 1, 0,
		0, 0, 78, 79, 5, 51, 0, 0, 79, 80, 5, 52, 0, 0, 80, 81, 5, 7, 0, 0, 81,
		11, 1, 0, 0, 0, 82, 83, 5, 28, 0, 0, 83, 92, 5, 49, 0, 0, 84, 89, 3, 30,
		15, 0, 85, 86, 5, 58, 0, 0, 86, 88, 3, 30, 15, 0, 87, 85, 1, 0, 0, 0, 88,
		91, 1, 0, 0, 0, 89, 87, 1, 0, 0, 0, 89, 90, 1, 0, 0, 0, 90, 93, 1, 0, 0,
		0, 91, 89, 1, 0, 0, 0, 92, 84, 1, 0, 0, 0, 92, 93, 1, 0, 0, 0, 93, 94,
		1, 0, 0, 0, 94, 98, 5, 50, 0, 0, 95, 98, 3, 30, 15, 0, 96, 98, 3, 14, 7,
		0, 97, 82, 1, 0, 0, 0, 97, 95, 1, 0, 0, 0, 97, 96, 1, 0, 0, 0, 98, 13,
		1, 0, 0, 0, 99, 104, 3, 16, 8, 0, 100, 104, 3, 18, 9, 0, 101, 104, 3, 20,
		10, 0, 102, 104, 3, 28, 14, 0, 103, 99, 1, 0, 0, 0, 103, 100, 1, 0, 0,
		0, 103, 101, 1, 0, 0, 0, 103, 102, 1, 0, 0, 0, 104, 15, 1, 0, 0, 0, 105,
		107, 5, 8, 0, 0, 106, 108, 5, 49, 0, 0, 107, 106, 1, 0, 0, 0, 107, 108,
		1, 0, 0, 0, 108, 109, 1, 0, 0, 0, 109, 111, 3, 30, 15, 0, 110, 112, 5,
		50, 0, 0, 111, 110, 1, 0, 0, 0, 111, 112, 1, 0, 0, 0, 112, 113, 1, 0, 0,
		0, 113, 117, 5, 53, 0, 0, 114, 116, 3, 6, 3, 0, 115, 114, 1, 0, 0, 0, 116,
		119, 1, 0, 0, 0, 117, 115, 1, 0, 0, 0, 117, 118, 1, 0, 0, 0, 118, 120,
		1, 0, 0, 0, 119, 117, 1, 0, 0, 0, 120, 130, 5, 54, 0, 0, 121, 122, 5, 8,
		0, 0, 122, 126, 5, 53, 0, 0, 123, 125, 3, 6, 3, 0, 124, 123, 1, 0, 0, 0,
		125, 128, 1, 0, 0, 0, 126, 124, 1, 0, 0, 0, 126, 127, 1, 0, 0, 0, 127,
		129, 1, 0, 0, 0, 128, 126, 1, 0, 0, 0, 129, 131, 5, 54, 0, 0, 130, 121,
		1, 0, 0, 0, 130, 131, 1, 0, 0, 0, 131, 17, 1, 0, 0, 0, 132, 134, 5, 8,
		0, 0, 133, 135, 5, 49, 0, 0, 134, 133, 1, 0, 0, 0, 134, 135, 1, 0, 0, 0,
		135, 137, 1, 0, 0, 0, 136, 138, 3, 12, 6, 0, 137, 136, 1, 0, 0, 0, 137,
		138, 1, 0, 0, 0, 138, 139, 1, 0, 0, 0, 139, 140, 5, 55, 0, 0, 140, 141,
		3, 30, 15, 0, 141, 143, 5, 55, 0, 0, 142, 144, 3, 12, 6, 0, 143, 142, 1,
		0, 0, 0, 143, 144, 1, 0, 0, 0, 144, 145, 1, 0, 0, 0, 145, 146, 3, 4, 2,
		0, 146, 158, 1, 0, 0, 0, 147, 149, 5, 8, 0, 0, 148, 150, 5, 49, 0, 0, 149,
		148, 1, 0, 0, 0, 149, 150, 1, 0, 0, 0, 150, 151, 1, 0, 0, 0, 151, 153,
		3, 30, 15, 0, 152, 154, 5, 50, 0, 0, 153, 152, 1, 0, 0, 0, 153, 154, 1,
		0, 0, 0, 154, 155, 1, 0, 0, 0, 155, 156, 3, 4, 2, 0, 156, 158, 1, 0, 0,
		0, 157, 132, 1, 0, 0, 0, 157, 147, 1, 0, 0, 0, 158, 19, 1, 0, 0, 0, 159,
		160, 5, 8, 0, 0, 160, 161, 3, 30, 15, 0, 161, 165, 5, 53, 0, 0, 162, 164,
		3, 22, 11, 0, 163, 162, 1, 0, 0, 0, 164, 167, 1, 0, 0, 0, 165, 163, 1,
		0, 0, 0, 165, 166, 1, 0, 0, 0, 166, 169, 1, 0, 0, 0, 167, 165, 1, 0, 0,
		0, 168, 170, 3, 24, 12, 0, 169, 168, 1, 0, 0, 0, 169, 170, 1, 0, 0, 0,
		170, 171, 1, 0, 0, 0, 171, 172, 5, 54, 0, 0, 172, 21, 1, 0, 0, 0, 173,
		174, 5, 4, 0, 0, 174, 175, 3, 30, 15, 0, 175, 179, 5, 56, 0, 0, 176, 178,
		3, 6, 3, 0, 177, 176, 1, 0, 0, 0, 178, 181, 1, 0, 0, 0, 179, 177, 1, 0,
		0, 0, 179, 180, 1, 0, 0, 0, 180, 23, 1, 0, 0, 0, 181, 179, 1, 0, 0, 0,
		182, 183, 5, 5, 0, 0, 183, 187, 5, 56, 0, 0, 184, 186, 3, 6, 3, 0, 185,
		184, 1, 0, 0, 0, 186, 189, 1, 0, 0, 0, 187, 185, 1, 0, 0, 0, 187, 188,
		1, 0, 0, 0, 188, 25, 1, 0, 0, 0, 189, 187, 1, 0, 0, 0, 190, 191, 5, 16,
		0, 0, 191, 200, 5, 49, 0, 0, 192, 197, 3, 30, 15, 0, 193, 194, 5, 58, 0,
		0, 194, 196, 3, 30, 15, 0, 195, 193, 1, 0, 0, 0, 196, 199, 1, 0, 0, 0,
		197, 195, 1, 0, 0, 0, 197, 198, 1, 0, 0, 0, 198, 201, 1, 0, 0, 0, 199,
		197, 1, 0, 0, 0, 200, 192, 1, 0, 0, 0, 200, 201, 1, 0, 0, 0, 201, 202,
		1, 0, 0, 0, 202, 230, 5, 50, 0, 0, 203, 204, 5, 17, 0, 0, 204, 213, 5,
		49, 0, 0, 205, 210, 3, 30, 15, 0, 206, 207, 5, 58, 0, 0, 207, 209, 3, 30,
		15, 0, 208, 206, 1, 0, 0, 0, 209, 212, 1, 0, 0, 0, 210, 208, 1, 0, 0, 0,
		210, 211, 1, 0, 0, 0, 211, 214, 1, 0, 0, 0, 212, 210, 1, 0, 0, 0, 213,
		205, 1, 0, 0, 0, 213, 214, 1, 0, 0, 0, 214, 215, 1, 0, 0, 0, 215, 230,
		5, 50, 0, 0, 216, 217, 5, 29, 0, 0, 217, 226, 5, 49, 0, 0, 218, 223, 3,
		30, 15, 0, 219, 220, 5, 58, 0, 0, 220, 222, 3, 30, 15, 0, 221, 219, 1,
		0, 0, 0, 222, 225, 1, 0, 0, 0, 223, 221, 1, 0, 0, 0, 223, 224, 1, 0, 0,
		0, 224, 227, 1, 0, 0, 0, 225, 223, 1, 0, 0, 0, 226, 218, 1, 0, 0, 0, 226,
		227, 1, 0, 0, 0, 227, 228, 1, 0, 0, 0, 228, 230, 5, 50, 0, 0, 229, 190,
		1, 0, 0, 0, 229, 203, 1, 0, 0, 0, 229, 216, 1, 0, 0, 0, 230, 27, 1, 0,
		0, 0, 231, 232, 5, 6, 0, 0, 232, 233, 5, 49, 0, 0, 233, 234, 3, 30, 15,
		0, 234, 235, 5, 50, 0, 0, 235, 239, 5, 51, 0, 0, 236, 238, 3, 6, 3, 0,
		237, 236, 1, 0, 0, 0, 238, 241, 1, 0, 0, 0, 239, 237, 1, 0, 0, 0, 239,
		240, 1, 0, 0, 0, 240, 242, 1, 0, 0, 0, 241, 239, 1, 0, 0, 0, 242, 243,
		5, 52, 0, 0, 243, 29, 1, 0, 0, 0, 244, 245, 6, 15, -1, 0, 245, 280, 3,
		36, 18, 0, 246, 247, 5, 49, 0, 0, 247, 248, 3, 30, 15, 0, 248, 249, 5,
		50, 0, 0, 249, 280, 1, 0, 0, 0, 250, 251, 5, 51, 0, 0, 251, 252, 3, 30,
		15, 0, 252, 253, 5, 52, 0, 0, 253, 280, 1, 0, 0, 0, 254, 255, 3, 10, 5,
		0, 255, 257, 5, 53, 0, 0, 256, 258, 3, 38, 19, 0, 257, 256, 1, 0, 0, 0,
		257, 258, 1, 0, 0, 0, 258, 259, 1, 0, 0, 0, 259, 260, 5, 54, 0, 0, 260,
		280, 1, 0, 0, 0, 261, 280, 3, 26, 13, 0, 262, 263, 7, 0, 0, 0, 263, 280,
		3, 30, 15, 13, 264, 280, 5, 29, 0, 0, 265, 280, 3, 40, 20, 0, 266, 267,
		5, 29, 0, 0, 267, 268, 5, 57, 0, 0, 268, 280, 5, 29, 0, 0, 269, 270, 5,
		29, 0, 0, 270, 271, 5, 57, 0, 0, 271, 280, 3, 30, 15, 3, 272, 273, 5, 29,
		0, 0, 273, 274, 5, 7, 0, 0, 274, 275, 5, 48, 0, 0, 275, 280, 3, 30, 15,
		2, 276, 277, 5, 29, 0, 0, 277, 278, 7, 1, 0, 0, 278, 280, 3, 30, 15, 1,
		279, 244, 1, 0, 0, 0, 279, 246, 1, 0, 0, 0, 279, 250, 1, 0, 0, 0, 279,
		254, 1, 0, 0, 0, 279, 261, 1, 0, 0, 0, 279, 262, 1, 0, 0, 0, 279, 264,
		1, 0, 0, 0, 279, 265, 1, 0, 0, 0, 279, 266, 1, 0, 0, 0, 279, 269, 1, 0,
		0, 0, 279, 272, 1, 0, 0, 0, 279, 276, 1, 0, 0, 0, 280, 301, 1, 0, 0, 0,
		281, 282, 10, 12, 0, 0, 282, 283, 7, 2, 0, 0, 283, 300, 3, 30, 15, 13,
		284, 285, 10, 11, 0, 0, 285, 286, 7, 3, 0, 0, 286, 300, 3, 30, 15, 12,
		287, 288, 10, 10, 0, 0, 288, 289, 7, 4, 0, 0, 289, 300, 3, 30, 15, 11,
		290, 291, 10, 9, 0, 0, 291, 292, 7, 5, 0, 0, 292, 300, 3, 30, 15, 10, 293,
		294, 10, 8, 0, 0, 294, 295, 7, 6, 0, 0, 295, 300, 3, 30, 15, 9, 296, 297,
		10, 7, 0, 0, 297, 298, 5, 40, 0, 0, 298, 300, 3, 30, 15, 8, 299, 281, 1,
		0, 0, 0, 299, 284, 1, 0, 0, 0, 299, 287, 1, 0, 0, 0, 299, 290, 1, 0, 0,
		0, 299, 293, 1, 0, 0, 0, 299, 296, 1, 0, 0, 0, 300, 303, 1, 0, 0, 0, 301,
		299, 1, 0, 0, 0, 301, 302, 1, 0, 0, 0, 302, 31, 1, 0, 0, 0, 303, 301, 1,
		0, 0, 0, 304, 309, 3, 30, 15, 0, 305, 306, 5, 58, 0, 0, 306, 308, 3, 30,
		15, 0, 307, 305, 1, 0, 0, 0, 308, 311, 1, 0, 0, 0, 309, 307, 1, 0, 0, 0,
		309, 310, 1, 0, 0, 0, 310, 33, 1, 0, 0, 0, 311, 309, 1, 0, 0, 0, 312, 313,
		3, 36, 18, 0, 313, 35, 1, 0, 0, 0, 314, 320, 5, 19, 0, 0, 315, 320, 5,
		20, 0, 0, 316, 320, 5, 21, 0, 0, 317, 320, 5, 18, 0, 0, 318, 320, 5, 22,
		0, 0, 319, 314, 1, 0, 0, 0, 319, 315, 1, 0, 0, 0, 319, 316, 1, 0, 0, 0,
		319, 317, 1, 0, 0, 0, 319, 318, 1, 0, 0, 0, 320, 37, 1, 0, 0, 0, 321, 326,
		3, 30, 15, 0, 322, 323, 5, 58, 0, 0, 323, 325, 3, 30, 15, 0, 324, 322,
		1, 0, 0, 0, 325, 328, 1, 0, 0, 0, 326, 324, 1, 0, 0, 0, 326, 327, 1, 0,
		0, 0, 327, 39, 1, 0, 0, 0, 328, 326, 1, 0, 0, 0, 329, 330, 5, 29, 0, 0,
		330, 334, 5, 30, 0, 0, 331, 332, 5, 29, 0, 0, 332, 334, 5, 31, 0, 0, 333,
		329, 1, 0, 0, 0, 333, 331, 1, 0, 0, 0, 334, 41, 1, 0, 0, 0, 40, 45, 60,
		67, 72, 76, 89, 92, 97, 103, 107, 111, 117, 126, 130, 134, 137, 143, 149,
		153, 157, 165, 169, 179, 187, 197, 200, 210, 213, 223, 226, 229, 239, 257,
		279, 299, 301, 309, 319, 326, 333,
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

// VlangParserInit initializes any static state used to implement VlangParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewVlangParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func VlangParserInit() {
	staticData := &VlangParserStaticData
	staticData.once.Do(vlangParserInit)
}

// NewVlangParser produces a new parser instance for the optional input antlr.TokenStream.
func NewVlangParser(input antlr.TokenStream) *VlangParser {
	VlangParserInit()
	this := new(VlangParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &VlangParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Vlang.g4"

	return this
}

// VlangParser tokens.
const (
	VlangParserEOF             = antlr.TokenEOF
	VlangParserT__0            = 1
	VlangParserT__1            = 2
	VlangParserT__2            = 3
	VlangParserT__3            = 4
	VlangParserT__4            = 5
	VlangParserT__5            = 6
	VlangParserTIPO            = 7
	VlangParserCONDICIONAL_ETC = 8
	VlangParserLEN             = 9
	VlangParserCAP             = 10
	VlangParserAPPEND          = 11
	VlangParserIF              = 12
	VlangParserELSE            = 13
	VlangParserFOR             = 14
	VlangParserSWITCH          = 15
	VlangParserINDEXOF         = 16
	VlangParserJOIN            = 17
	VlangParserBOOLEANO        = 18
	VlangParserENTERO          = 19
	VlangParserDECIMAL         = 20
	VlangParserCADENA          = 21
	VlangParserCARACTER        = 22
	VlangParserTIPO_ENTERO     = 23
	VlangParserTIPO_DECIMAL    = 24
	VlangParserTIPO_CADENA     = 25
	VlangParserTIPO_BOOLEANO   = 26
	VlangParserTIPO_CHAR       = 27
	VlangParserPRINT           = 28
	VlangParserID              = 29
	VlangParserINC             = 30
	VlangParserDEC             = 31
	VlangParserSUMAIMPLICITA   = 32
	VlangParserRESTOIMPLICITO  = 33
	VlangParserPLUS            = 34
	VlangParserMINUS           = 35
	VlangParserMUL             = 36
	VlangParserDIV             = 37
	VlangParserMOD             = 38
	VlangParserNOT             = 39
	VlangParserOR              = 40
	VlangParserAND             = 41
	VlangParserEQ              = 42
	VlangParserNEQ             = 43
	VlangParserLE              = 44
	VlangParserGE              = 45
	VlangParserLT              = 46
	VlangParserGT              = 47
	VlangParserASSIGN          = 48
	VlangParserLPAREN          = 49
	VlangParserRPAREN          = 50
	VlangParserLBRACK          = 51
	VlangParserRBRACK          = 52
	VlangParserLBRACE          = 53
	VlangParserRBRACE          = 54
	VlangParserSEMICOLON       = 55
	VlangParserCOLON           = 56
	VlangParserDOT             = 57
	VlangParserCOMMA           = 58
	VlangParserWS              = 59
	VlangParserLINE_COMMENT    = 60
	VlangParserBLOCK_COMMENT   = 61
)

// VlangParser rules.
const (
	VlangParserRULE_programa           = 0
	VlangParserRULE_funcMain           = 1
	VlangParserRULE_block              = 2
	VlangParserRULE_declaraciones      = 3
	VlangParserRULE_varDcl             = 4
	VlangParserRULE_tipoSlice          = 5
	VlangParserRULE_stmt               = 6
	VlangParserRULE_sentencias_control = 7
	VlangParserRULE_ifDcl              = 8
	VlangParserRULE_forDcl             = 9
	VlangParserRULE_switchDcl          = 10
	VlangParserRULE_caseBlock          = 11
	VlangParserRULE_defaultBlock       = 12
	VlangParserRULE_llamadaFuncion     = 13
	VlangParserRULE_whileDcl           = 14
	VlangParserRULE_expresion          = 15
	VlangParserRULE_parametros         = 16
	VlangParserRULE_valores            = 17
	VlangParserRULE_valor              = 18
	VlangParserRULE_listaExpresiones   = 19
	VlangParserRULE_incredecre         = 20
)

// IProgramaContext is an interface to support dynamic dispatch.
type IProgramaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllFuncMain() []IFuncMainContext
	FuncMain(i int) IFuncMainContext

	// IsProgramaContext differentiates from other interfaces.
	IsProgramaContext()
}

type ProgramaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramaContext() *ProgramaContext {
	var p = new(ProgramaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_programa
	return p
}

func InitEmptyProgramaContext(p *ProgramaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_programa
}

func (*ProgramaContext) IsProgramaContext() {}

func NewProgramaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramaContext {
	var p = new(ProgramaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_programa

	return p
}

func (s *ProgramaContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramaContext) EOF() antlr.TerminalNode {
	return s.GetToken(VlangParserEOF, 0)
}

func (s *ProgramaContext) AllFuncMain() []IFuncMainContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFuncMainContext); ok {
			len++
		}
	}

	tst := make([]IFuncMainContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFuncMainContext); ok {
			tst[i] = t.(IFuncMainContext)
			i++
		}
	}

	return tst
}

func (s *ProgramaContext) FuncMain(i int) IFuncMainContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncMainContext); ok {
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

	return t.(IFuncMainContext)
}

func (s *ProgramaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterPrograma(s)
	}
}

func (s *ProgramaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitPrograma(s)
	}
}

func (s *ProgramaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitPrograma(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Programa() (localctx IProgramaContext) {
	localctx = NewProgramaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, VlangParserRULE_programa)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(45)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == VlangParserT__0 {
		{
			p.SetState(42)
			p.FuncMain()
		}

		p.SetState(47)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(48)
		p.Match(VlangParserEOF)
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

// IFuncMainContext is an interface to support dynamic dispatch.
type IFuncMainContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	Block() IBlockContext

	// IsFuncMainContext differentiates from other interfaces.
	IsFuncMainContext()
}

type FuncMainContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncMainContext() *FuncMainContext {
	var p = new(FuncMainContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_funcMain
	return p
}

func InitEmptyFuncMainContext(p *FuncMainContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_funcMain
}

func (*FuncMainContext) IsFuncMainContext() {}

func NewFuncMainContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncMainContext {
	var p = new(FuncMainContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_funcMain

	return p
}

func (s *FuncMainContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncMainContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(VlangParserLPAREN, 0)
}

func (s *FuncMainContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(VlangParserRPAREN, 0)
}

func (s *FuncMainContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FuncMainContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncMainContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncMainContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterFuncMain(s)
	}
}

func (s *FuncMainContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitFuncMain(s)
	}
}

func (s *FuncMainContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitFuncMain(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) FuncMain() (localctx IFuncMainContext) {
	localctx = NewFuncMainContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, VlangParserRULE_funcMain)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(50)
		p.Match(VlangParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(51)
		p.Match(VlangParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(52)
		p.Match(VlangParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(53)
		p.Match(VlangParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(54)
		p.Block()
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

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllDeclaraciones() []IDeclaracionesContext
	Declaraciones(i int) IDeclaracionesContext

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(VlangParserLBRACE, 0)
}

func (s *BlockContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(VlangParserRBRACE, 0)
}

func (s *BlockContext) AllDeclaraciones() []IDeclaracionesContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDeclaracionesContext); ok {
			len++
		}
	}

	tst := make([]IDeclaracionesContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDeclaracionesContext); ok {
			tst[i] = t.(IDeclaracionesContext)
			i++
		}
	}

	return tst
}

func (s *BlockContext) Declaraciones(i int) IDeclaracionesContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclaracionesContext); ok {
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

	return t.(IDeclaracionesContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, VlangParserRULE_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(56)
		p.Match(VlangParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(60)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2815334696288584) != 0 {
		{
			p.SetState(57)
			p.Declaraciones()
		}

		p.SetState(62)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(63)
		p.Match(VlangParserRBRACE)
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

// IDeclaracionesContext is an interface to support dynamic dispatch.
type IDeclaracionesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VarDcl() IVarDclContext
	Stmt() IStmtContext

	// IsDeclaracionesContext differentiates from other interfaces.
	IsDeclaracionesContext()
}

type DeclaracionesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclaracionesContext() *DeclaracionesContext {
	var p = new(DeclaracionesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_declaraciones
	return p
}

func InitEmptyDeclaracionesContext(p *DeclaracionesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_declaraciones
}

func (*DeclaracionesContext) IsDeclaracionesContext() {}

func NewDeclaracionesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclaracionesContext {
	var p = new(DeclaracionesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_declaraciones

	return p
}

func (s *DeclaracionesContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclaracionesContext) VarDcl() IVarDclContext {
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

func (s *DeclaracionesContext) Stmt() IStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *DeclaracionesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclaracionesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclaracionesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterDeclaraciones(s)
	}
}

func (s *DeclaracionesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitDeclaraciones(s)
	}
}

func (s *DeclaracionesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitDeclaraciones(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Declaraciones() (localctx IDeclaracionesContext) {
	localctx = NewDeclaracionesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, VlangParserRULE_declaraciones)
	p.SetState(67)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case VlangParserT__2:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(65)
			p.VarDcl()
		}

	case VlangParserT__5, VlangParserCONDICIONAL_ETC, VlangParserINDEXOF, VlangParserJOIN, VlangParserBOOLEANO, VlangParserENTERO, VlangParserDECIMAL, VlangParserCADENA, VlangParserCARACTER, VlangParserPRINT, VlangParserID, VlangParserMINUS, VlangParserNOT, VlangParserLPAREN, VlangParserLBRACK:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(66)
			p.Stmt()
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

// IVarDclContext is an interface to support dynamic dispatch.
type IVarDclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
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
	p.RuleIndex = VlangParserRULE_varDcl
	return p
}

func InitEmptyVarDclContext(p *VarDclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_varDcl
}

func (*VarDclContext) IsVarDclContext() {}

func NewVarDclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarDclContext {
	var p = new(VarDclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_varDcl

	return p
}

func (s *VarDclContext) GetParser() antlr.Parser { return s.parser }

func (s *VarDclContext) CopyAll(ctx *VarDclContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *VarDclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarDclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type VariableDeclarationContext struct {
	VarDclContext
}

func NewVariableDeclarationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VariableDeclarationContext {
	var p = new(VariableDeclarationContext)

	InitEmptyVarDclContext(&p.VarDclContext)
	p.parser = parser
	p.CopyAll(ctx.(*VarDclContext))

	return p
}

func (s *VariableDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDeclarationContext) ID() antlr.TerminalNode {
	return s.GetToken(VlangParserID, 0)
}

func (s *VariableDeclarationContext) TIPO() antlr.TerminalNode {
	return s.GetToken(VlangParserTIPO, 0)
}

func (s *VariableDeclarationContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(VlangParserASSIGN, 0)
}

func (s *VariableDeclarationContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *VariableDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterVariableDeclaration(s)
	}
}

func (s *VariableDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitVariableDeclaration(s)
	}
}

func (s *VariableDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitVariableDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) VarDcl() (localctx IVarDclContext) {
	localctx = NewVarDclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, VlangParserRULE_varDcl)
	var _la int

	localctx = NewVariableDeclarationContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(69)
		p.Match(VlangParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(70)
		p.Match(VlangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(72)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == VlangParserTIPO {
		{
			p.SetState(71)
			p.Match(VlangParserTIPO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(76)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == VlangParserASSIGN {
		{
			p.SetState(74)
			p.Match(VlangParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(75)
			p.expresion(0)
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

// ITipoSliceContext is an interface to support dynamic dispatch.
type ITipoSliceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACK() antlr.TerminalNode
	RBRACK() antlr.TerminalNode
	TIPO() antlr.TerminalNode

	// IsTipoSliceContext differentiates from other interfaces.
	IsTipoSliceContext()
}

type TipoSliceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTipoSliceContext() *TipoSliceContext {
	var p = new(TipoSliceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_tipoSlice
	return p
}

func InitEmptyTipoSliceContext(p *TipoSliceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_tipoSlice
}

func (*TipoSliceContext) IsTipoSliceContext() {}

func NewTipoSliceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TipoSliceContext {
	var p = new(TipoSliceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_tipoSlice

	return p
}

func (s *TipoSliceContext) GetParser() antlr.Parser { return s.parser }

func (s *TipoSliceContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(VlangParserLBRACK, 0)
}

func (s *TipoSliceContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(VlangParserRBRACK, 0)
}

func (s *TipoSliceContext) TIPO() antlr.TerminalNode {
	return s.GetToken(VlangParserTIPO, 0)
}

func (s *TipoSliceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TipoSliceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TipoSliceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterTipoSlice(s)
	}
}

func (s *TipoSliceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitTipoSlice(s)
	}
}

func (s *TipoSliceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitTipoSlice(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) TipoSlice() (localctx ITipoSliceContext) {
	localctx = NewTipoSliceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, VlangParserRULE_tipoSlice)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(78)
		p.Match(VlangParserLBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(79)
		p.Match(VlangParserRBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(80)
		p.Match(VlangParserTIPO)
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

// IStmtContext is an interface to support dynamic dispatch.
type IStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsStmtContext differentiates from other interfaces.
	IsStmtContext()
}

type StmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStmtContext() *StmtContext {
	var p = new(StmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_stmt
	return p
}

func InitEmptyStmtContext(p *StmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_stmt
}

func (*StmtContext) IsStmtContext() {}

func NewStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StmtContext {
	var p = new(StmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_stmt

	return p
}

func (s *StmtContext) GetParser() antlr.Parser { return s.parser }

func (s *StmtContext) CopyAll(ctx *StmtContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *StmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type PrintStatementContext struct {
	StmtContext
}

func NewPrintStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrintStatementContext {
	var p = new(PrintStatementContext)

	InitEmptyStmtContext(&p.StmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*StmtContext))

	return p
}

func (s *PrintStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintStatementContext) PRINT() antlr.TerminalNode {
	return s.GetToken(VlangParserPRINT, 0)
}

func (s *PrintStatementContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(VlangParserLPAREN, 0)
}

func (s *PrintStatementContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(VlangParserRPAREN, 0)
}

func (s *PrintStatementContext) AllExpresion() []IExpresionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpresionContext); ok {
			len++
		}
	}

	tst := make([]IExpresionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpresionContext); ok {
			tst[i] = t.(IExpresionContext)
			i++
		}
	}

	return tst
}

func (s *PrintStatementContext) Expresion(i int) IExpresionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
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

	return t.(IExpresionContext)
}

func (s *PrintStatementContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(VlangParserCOMMA)
}

func (s *PrintStatementContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(VlangParserCOMMA, i)
}

func (s *PrintStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterPrintStatement(s)
	}
}

func (s *PrintStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitPrintStatement(s)
	}
}

func (s *PrintStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitPrintStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

type ControlStatementContext struct {
	StmtContext
}

func NewControlStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ControlStatementContext {
	var p = new(ControlStatementContext)

	InitEmptyStmtContext(&p.StmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*StmtContext))

	return p
}

func (s *ControlStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ControlStatementContext) Sentencias_control() ISentencias_controlContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISentencias_controlContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISentencias_controlContext)
}

func (s *ControlStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterControlStatement(s)
	}
}

func (s *ControlStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitControlStatement(s)
	}
}

func (s *ControlStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitControlStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpresionStatementContext struct {
	StmtContext
}

func NewExpresionStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpresionStatementContext {
	var p = new(ExpresionStatementContext)

	InitEmptyStmtContext(&p.StmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*StmtContext))

	return p
}

func (s *ExpresionStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpresionStatementContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *ExpresionStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterExpresionStatement(s)
	}
}

func (s *ExpresionStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitExpresionStatement(s)
	}
}

func (s *ExpresionStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitExpresionStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Stmt() (localctx IStmtContext) {
	localctx = NewStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, VlangParserRULE_stmt)
	var _la int

	p.SetState(97)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case VlangParserPRINT:
		localctx = NewPrintStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(82)
			p.Match(VlangParserPRINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(83)
			p.Match(VlangParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(92)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2815334427852800) != 0 {
			{
				p.SetState(84)
				p.expresion(0)
			}
			p.SetState(89)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == VlangParserCOMMA {
				{
					p.SetState(85)
					p.Match(VlangParserCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(86)
					p.expresion(0)
				}

				p.SetState(91)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(94)
			p.Match(VlangParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VlangParserINDEXOF, VlangParserJOIN, VlangParserBOOLEANO, VlangParserENTERO, VlangParserDECIMAL, VlangParserCADENA, VlangParserCARACTER, VlangParserID, VlangParserMINUS, VlangParserNOT, VlangParserLPAREN, VlangParserLBRACK:
		localctx = NewExpresionStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(95)
			p.expresion(0)
		}

	case VlangParserT__5, VlangParserCONDICIONAL_ETC:
		localctx = NewControlStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(96)
			p.Sentencias_control()
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

// ISentencias_controlContext is an interface to support dynamic dispatch.
type ISentencias_controlContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSentencias_controlContext differentiates from other interfaces.
	IsSentencias_controlContext()
}

type Sentencias_controlContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySentencias_controlContext() *Sentencias_controlContext {
	var p = new(Sentencias_controlContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_sentencias_control
	return p
}

func InitEmptySentencias_controlContext(p *Sentencias_controlContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_sentencias_control
}

func (*Sentencias_controlContext) IsSentencias_controlContext() {}

func NewSentencias_controlContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Sentencias_controlContext {
	var p = new(Sentencias_controlContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_sentencias_control

	return p
}

func (s *Sentencias_controlContext) GetParser() antlr.Parser { return s.parser }

func (s *Sentencias_controlContext) CopyAll(ctx *Sentencias_controlContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Sentencias_controlContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Sentencias_controlContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type While_contextContext struct {
	Sentencias_controlContext
}

func NewWhile_contextContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *While_contextContext {
	var p = new(While_contextContext)

	InitEmptySentencias_controlContext(&p.Sentencias_controlContext)
	p.parser = parser
	p.CopyAll(ctx.(*Sentencias_controlContext))

	return p
}

func (s *While_contextContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *While_contextContext) WhileDcl() IWhileDclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhileDclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhileDclContext)
}

func (s *While_contextContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterWhile_context(s)
	}
}

func (s *While_contextContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitWhile_context(s)
	}
}

func (s *While_contextContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitWhile_context(s)

	default:
		return t.VisitChildren(s)
	}
}

type Switch_contextContext struct {
	Sentencias_controlContext
}

func NewSwitch_contextContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Switch_contextContext {
	var p = new(Switch_contextContext)

	InitEmptySentencias_controlContext(&p.Sentencias_controlContext)
	p.parser = parser
	p.CopyAll(ctx.(*Sentencias_controlContext))

	return p
}

func (s *Switch_contextContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Switch_contextContext) SwitchDcl() ISwitchDclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISwitchDclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISwitchDclContext)
}

func (s *Switch_contextContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterSwitch_context(s)
	}
}

func (s *Switch_contextContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitSwitch_context(s)
	}
}

func (s *Switch_contextContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitSwitch_context(s)

	default:
		return t.VisitChildren(s)
	}
}

type If_contextContext struct {
	Sentencias_controlContext
}

func NewIf_contextContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *If_contextContext {
	var p = new(If_contextContext)

	InitEmptySentencias_controlContext(&p.Sentencias_controlContext)
	p.parser = parser
	p.CopyAll(ctx.(*Sentencias_controlContext))

	return p
}

func (s *If_contextContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *If_contextContext) IfDcl() IIfDclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfDclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfDclContext)
}

func (s *If_contextContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterIf_context(s)
	}
}

func (s *If_contextContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitIf_context(s)
	}
}

func (s *If_contextContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitIf_context(s)

	default:
		return t.VisitChildren(s)
	}
}

type For_contextContext struct {
	Sentencias_controlContext
}

func NewFor_contextContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *For_contextContext {
	var p = new(For_contextContext)

	InitEmptySentencias_controlContext(&p.Sentencias_controlContext)
	p.parser = parser
	p.CopyAll(ctx.(*Sentencias_controlContext))

	return p
}

func (s *For_contextContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *For_contextContext) ForDcl() IForDclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForDclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForDclContext)
}

func (s *For_contextContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterFor_context(s)
	}
}

func (s *For_contextContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitFor_context(s)
	}
}

func (s *For_contextContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitFor_context(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Sentencias_control() (localctx ISentencias_controlContext) {
	localctx = NewSentencias_controlContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, VlangParserRULE_sentencias_control)
	p.SetState(103)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		localctx = NewIf_contextContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(99)
			p.IfDcl()
		}

	case 2:
		localctx = NewFor_contextContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(100)
			p.ForDcl()
		}

	case 3:
		localctx = NewSwitch_contextContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(101)
			p.SwitchDcl()
		}

	case 4:
		localctx = NewWhile_contextContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(102)
			p.WhileDcl()
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

// IIfDclContext is an interface to support dynamic dispatch.
type IIfDclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllCONDICIONAL_ETC() []antlr.TerminalNode
	CONDICIONAL_ETC(i int) antlr.TerminalNode
	Expresion() IExpresionContext
	AllLBRACE() []antlr.TerminalNode
	LBRACE(i int) antlr.TerminalNode
	AllRBRACE() []antlr.TerminalNode
	RBRACE(i int) antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	AllDeclaraciones() []IDeclaracionesContext
	Declaraciones(i int) IDeclaracionesContext

	// IsIfDclContext differentiates from other interfaces.
	IsIfDclContext()
}

type IfDclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfDclContext() *IfDclContext {
	var p = new(IfDclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_ifDcl
	return p
}

func InitEmptyIfDclContext(p *IfDclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_ifDcl
}

func (*IfDclContext) IsIfDclContext() {}

func NewIfDclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfDclContext {
	var p = new(IfDclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_ifDcl

	return p
}

func (s *IfDclContext) GetParser() antlr.Parser { return s.parser }

func (s *IfDclContext) AllCONDICIONAL_ETC() []antlr.TerminalNode {
	return s.GetTokens(VlangParserCONDICIONAL_ETC)
}

func (s *IfDclContext) CONDICIONAL_ETC(i int) antlr.TerminalNode {
	return s.GetToken(VlangParserCONDICIONAL_ETC, i)
}

func (s *IfDclContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *IfDclContext) AllLBRACE() []antlr.TerminalNode {
	return s.GetTokens(VlangParserLBRACE)
}

func (s *IfDclContext) LBRACE(i int) antlr.TerminalNode {
	return s.GetToken(VlangParserLBRACE, i)
}

func (s *IfDclContext) AllRBRACE() []antlr.TerminalNode {
	return s.GetTokens(VlangParserRBRACE)
}

func (s *IfDclContext) RBRACE(i int) antlr.TerminalNode {
	return s.GetToken(VlangParserRBRACE, i)
}

func (s *IfDclContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(VlangParserLPAREN, 0)
}

func (s *IfDclContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(VlangParserRPAREN, 0)
}

func (s *IfDclContext) AllDeclaraciones() []IDeclaracionesContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDeclaracionesContext); ok {
			len++
		}
	}

	tst := make([]IDeclaracionesContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDeclaracionesContext); ok {
			tst[i] = t.(IDeclaracionesContext)
			i++
		}
	}

	return tst
}

func (s *IfDclContext) Declaraciones(i int) IDeclaracionesContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclaracionesContext); ok {
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

	return t.(IDeclaracionesContext)
}

func (s *IfDclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfDclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfDclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterIfDcl(s)
	}
}

func (s *IfDclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitIfDcl(s)
	}
}

func (s *IfDclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitIfDcl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) IfDcl() (localctx IIfDclContext) {
	localctx = NewIfDclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, VlangParserRULE_ifDcl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(105)
		p.Match(VlangParserCONDICIONAL_ETC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(107)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(106)
			p.Match(VlangParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(109)
		p.expresion(0)
	}
	p.SetState(111)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == VlangParserRPAREN {
		{
			p.SetState(110)
			p.Match(VlangParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(113)
		p.Match(VlangParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(117)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2815334696288584) != 0 {
		{
			p.SetState(114)
			p.Declaraciones()
		}

		p.SetState(119)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(120)
		p.Match(VlangParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(130)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(121)
			p.Match(VlangParserCONDICIONAL_ETC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(122)
			p.Match(VlangParserLBRACE)
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

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2815334696288584) != 0 {
			{
				p.SetState(123)
				p.Declaraciones()
			}

			p.SetState(128)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(129)
			p.Match(VlangParserRBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
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

// IForDclContext is an interface to support dynamic dispatch.
type IForDclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsForDclContext differentiates from other interfaces.
	IsForDclContext()
}

type ForDclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForDclContext() *ForDclContext {
	var p = new(ForDclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_forDcl
	return p
}

func InitEmptyForDclContext(p *ForDclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_forDcl
}

func (*ForDclContext) IsForDclContext() {}

func NewForDclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForDclContext {
	var p = new(ForDclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_forDcl

	return p
}

func (s *ForDclContext) GetParser() antlr.Parser { return s.parser }

func (s *ForDclContext) CopyAll(ctx *ForDclContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ForDclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForDclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ForCondicionUnicaContext struct {
	ForDclContext
}

func NewForCondicionUnicaContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForCondicionUnicaContext {
	var p = new(ForCondicionUnicaContext)

	InitEmptyForDclContext(&p.ForDclContext)
	p.parser = parser
	p.CopyAll(ctx.(*ForDclContext))

	return p
}

func (s *ForCondicionUnicaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForCondicionUnicaContext) CONDICIONAL_ETC() antlr.TerminalNode {
	return s.GetToken(VlangParserCONDICIONAL_ETC, 0)
}

func (s *ForCondicionUnicaContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *ForCondicionUnicaContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ForCondicionUnicaContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(VlangParserLPAREN, 0)
}

func (s *ForCondicionUnicaContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(VlangParserRPAREN, 0)
}

func (s *ForCondicionUnicaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterForCondicionUnica(s)
	}
}

func (s *ForCondicionUnicaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitForCondicionUnica(s)
	}
}

func (s *ForCondicionUnicaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitForCondicionUnica(s)

	default:
		return t.VisitChildren(s)
	}
}

type ForClasicoContext struct {
	ForDclContext
}

func NewForClasicoContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForClasicoContext {
	var p = new(ForClasicoContext)

	InitEmptyForDclContext(&p.ForDclContext)
	p.parser = parser
	p.CopyAll(ctx.(*ForDclContext))

	return p
}

func (s *ForClasicoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForClasicoContext) CONDICIONAL_ETC() antlr.TerminalNode {
	return s.GetToken(VlangParserCONDICIONAL_ETC, 0)
}

func (s *ForClasicoContext) AllSEMICOLON() []antlr.TerminalNode {
	return s.GetTokens(VlangParserSEMICOLON)
}

func (s *ForClasicoContext) SEMICOLON(i int) antlr.TerminalNode {
	return s.GetToken(VlangParserSEMICOLON, i)
}

func (s *ForClasicoContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *ForClasicoContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ForClasicoContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(VlangParserLPAREN, 0)
}

func (s *ForClasicoContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *ForClasicoContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
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

	return t.(IStmtContext)
}

func (s *ForClasicoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterForClasico(s)
	}
}

func (s *ForClasicoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitForClasico(s)
	}
}

func (s *ForClasicoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitForClasico(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) ForDcl() (localctx IForDclContext) {
	localctx = NewForDclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, VlangParserRULE_forDcl)
	var _la int

	p.SetState(157)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		localctx = NewForClasicoContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(132)
			p.Match(VlangParserCONDICIONAL_ETC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(134)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(133)
				p.Match(VlangParserLPAREN)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}
		p.SetState(137)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2815334696288576) != 0 {
			{
				p.SetState(136)
				p.Stmt()
			}

		}
		{
			p.SetState(139)
			p.Match(VlangParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(140)
			p.expresion(0)
		}
		{
			p.SetState(141)
			p.Match(VlangParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(143)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2815334696288576) != 0 {
			{
				p.SetState(142)
				p.Stmt()
			}

		}
		{
			p.SetState(145)
			p.Block()
		}

	case 2:
		localctx = NewForCondicionUnicaContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(147)
			p.Match(VlangParserCONDICIONAL_ETC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(149)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(148)
				p.Match(VlangParserLPAREN)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}
		{
			p.SetState(151)
			p.expresion(0)
		}
		p.SetState(153)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == VlangParserRPAREN {
			{
				p.SetState(152)
				p.Match(VlangParserRPAREN)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(155)
			p.Block()
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

// ISwitchDclContext is an interface to support dynamic dispatch.
type ISwitchDclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CONDICIONAL_ETC() antlr.TerminalNode
	Expresion() IExpresionContext
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllCaseBlock() []ICaseBlockContext
	CaseBlock(i int) ICaseBlockContext
	DefaultBlock() IDefaultBlockContext

	// IsSwitchDclContext differentiates from other interfaces.
	IsSwitchDclContext()
}

type SwitchDclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitchDclContext() *SwitchDclContext {
	var p = new(SwitchDclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_switchDcl
	return p
}

func InitEmptySwitchDclContext(p *SwitchDclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_switchDcl
}

func (*SwitchDclContext) IsSwitchDclContext() {}

func NewSwitchDclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwitchDclContext {
	var p = new(SwitchDclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_switchDcl

	return p
}

func (s *SwitchDclContext) GetParser() antlr.Parser { return s.parser }

func (s *SwitchDclContext) CONDICIONAL_ETC() antlr.TerminalNode {
	return s.GetToken(VlangParserCONDICIONAL_ETC, 0)
}

func (s *SwitchDclContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *SwitchDclContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(VlangParserLBRACE, 0)
}

func (s *SwitchDclContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(VlangParserRBRACE, 0)
}

func (s *SwitchDclContext) AllCaseBlock() []ICaseBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICaseBlockContext); ok {
			len++
		}
	}

	tst := make([]ICaseBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICaseBlockContext); ok {
			tst[i] = t.(ICaseBlockContext)
			i++
		}
	}

	return tst
}

func (s *SwitchDclContext) CaseBlock(i int) ICaseBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICaseBlockContext); ok {
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

	return t.(ICaseBlockContext)
}

func (s *SwitchDclContext) DefaultBlock() IDefaultBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefaultBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDefaultBlockContext)
}

func (s *SwitchDclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchDclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SwitchDclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterSwitchDcl(s)
	}
}

func (s *SwitchDclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitSwitchDcl(s)
	}
}

func (s *SwitchDclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitSwitchDcl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) SwitchDcl() (localctx ISwitchDclContext) {
	localctx = NewSwitchDclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, VlangParserRULE_switchDcl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(159)
		p.Match(VlangParserCONDICIONAL_ETC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(160)
		p.expresion(0)
	}
	{
		p.SetState(161)
		p.Match(VlangParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(165)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == VlangParserT__3 {
		{
			p.SetState(162)
			p.CaseBlock()
		}

		p.SetState(167)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(169)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == VlangParserT__4 {
		{
			p.SetState(168)
			p.DefaultBlock()
		}

	}
	{
		p.SetState(171)
		p.Match(VlangParserRBRACE)
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

// ICaseBlockContext is an interface to support dynamic dispatch.
type ICaseBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expresion() IExpresionContext
	COLON() antlr.TerminalNode
	AllDeclaraciones() []IDeclaracionesContext
	Declaraciones(i int) IDeclaracionesContext

	// IsCaseBlockContext differentiates from other interfaces.
	IsCaseBlockContext()
}

type CaseBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCaseBlockContext() *CaseBlockContext {
	var p = new(CaseBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_caseBlock
	return p
}

func InitEmptyCaseBlockContext(p *CaseBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_caseBlock
}

func (*CaseBlockContext) IsCaseBlockContext() {}

func NewCaseBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CaseBlockContext {
	var p = new(CaseBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_caseBlock

	return p
}

func (s *CaseBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *CaseBlockContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *CaseBlockContext) COLON() antlr.TerminalNode {
	return s.GetToken(VlangParserCOLON, 0)
}

func (s *CaseBlockContext) AllDeclaraciones() []IDeclaracionesContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDeclaracionesContext); ok {
			len++
		}
	}

	tst := make([]IDeclaracionesContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDeclaracionesContext); ok {
			tst[i] = t.(IDeclaracionesContext)
			i++
		}
	}

	return tst
}

func (s *CaseBlockContext) Declaraciones(i int) IDeclaracionesContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclaracionesContext); ok {
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

	return t.(IDeclaracionesContext)
}

func (s *CaseBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CaseBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CaseBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterCaseBlock(s)
	}
}

func (s *CaseBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitCaseBlock(s)
	}
}

func (s *CaseBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitCaseBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) CaseBlock() (localctx ICaseBlockContext) {
	localctx = NewCaseBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, VlangParserRULE_caseBlock)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(173)
		p.Match(VlangParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(174)
		p.expresion(0)
	}
	{
		p.SetState(175)
		p.Match(VlangParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(179)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2815334696288584) != 0 {
		{
			p.SetState(176)
			p.Declaraciones()
		}

		p.SetState(181)
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

// IDefaultBlockContext is an interface to support dynamic dispatch.
type IDefaultBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	COLON() antlr.TerminalNode
	AllDeclaraciones() []IDeclaracionesContext
	Declaraciones(i int) IDeclaracionesContext

	// IsDefaultBlockContext differentiates from other interfaces.
	IsDefaultBlockContext()
}

type DefaultBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefaultBlockContext() *DefaultBlockContext {
	var p = new(DefaultBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_defaultBlock
	return p
}

func InitEmptyDefaultBlockContext(p *DefaultBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_defaultBlock
}

func (*DefaultBlockContext) IsDefaultBlockContext() {}

func NewDefaultBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefaultBlockContext {
	var p = new(DefaultBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_defaultBlock

	return p
}

func (s *DefaultBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *DefaultBlockContext) COLON() antlr.TerminalNode {
	return s.GetToken(VlangParserCOLON, 0)
}

func (s *DefaultBlockContext) AllDeclaraciones() []IDeclaracionesContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDeclaracionesContext); ok {
			len++
		}
	}

	tst := make([]IDeclaracionesContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDeclaracionesContext); ok {
			tst[i] = t.(IDeclaracionesContext)
			i++
		}
	}

	return tst
}

func (s *DefaultBlockContext) Declaraciones(i int) IDeclaracionesContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclaracionesContext); ok {
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

	return t.(IDeclaracionesContext)
}

func (s *DefaultBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefaultBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefaultBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterDefaultBlock(s)
	}
}

func (s *DefaultBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitDefaultBlock(s)
	}
}

func (s *DefaultBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitDefaultBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) DefaultBlock() (localctx IDefaultBlockContext) {
	localctx = NewDefaultBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, VlangParserRULE_defaultBlock)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(182)
		p.Match(VlangParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(183)
		p.Match(VlangParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(187)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2815334696288584) != 0 {
		{
			p.SetState(184)
			p.Declaraciones()
		}

		p.SetState(189)
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

// ILlamadaFuncionContext is an interface to support dynamic dispatch.
type ILlamadaFuncionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INDEXOF() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	AllExpresion() []IExpresionContext
	Expresion(i int) IExpresionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	JOIN() antlr.TerminalNode
	ID() antlr.TerminalNode

	// IsLlamadaFuncionContext differentiates from other interfaces.
	IsLlamadaFuncionContext()
}

type LlamadaFuncionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLlamadaFuncionContext() *LlamadaFuncionContext {
	var p = new(LlamadaFuncionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_llamadaFuncion
	return p
}

func InitEmptyLlamadaFuncionContext(p *LlamadaFuncionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_llamadaFuncion
}

func (*LlamadaFuncionContext) IsLlamadaFuncionContext() {}

func NewLlamadaFuncionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LlamadaFuncionContext {
	var p = new(LlamadaFuncionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_llamadaFuncion

	return p
}

func (s *LlamadaFuncionContext) GetParser() antlr.Parser { return s.parser }

func (s *LlamadaFuncionContext) INDEXOF() antlr.TerminalNode {
	return s.GetToken(VlangParserINDEXOF, 0)
}

func (s *LlamadaFuncionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(VlangParserLPAREN, 0)
}

func (s *LlamadaFuncionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(VlangParserRPAREN, 0)
}

func (s *LlamadaFuncionContext) AllExpresion() []IExpresionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpresionContext); ok {
			len++
		}
	}

	tst := make([]IExpresionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpresionContext); ok {
			tst[i] = t.(IExpresionContext)
			i++
		}
	}

	return tst
}

func (s *LlamadaFuncionContext) Expresion(i int) IExpresionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
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

	return t.(IExpresionContext)
}

func (s *LlamadaFuncionContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(VlangParserCOMMA)
}

func (s *LlamadaFuncionContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(VlangParserCOMMA, i)
}

func (s *LlamadaFuncionContext) JOIN() antlr.TerminalNode {
	return s.GetToken(VlangParserJOIN, 0)
}

func (s *LlamadaFuncionContext) ID() antlr.TerminalNode {
	return s.GetToken(VlangParserID, 0)
}

func (s *LlamadaFuncionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LlamadaFuncionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LlamadaFuncionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterLlamadaFuncion(s)
	}
}

func (s *LlamadaFuncionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitLlamadaFuncion(s)
	}
}

func (s *LlamadaFuncionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitLlamadaFuncion(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) LlamadaFuncion() (localctx ILlamadaFuncionContext) {
	localctx = NewLlamadaFuncionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, VlangParserRULE_llamadaFuncion)
	var _la int

	p.SetState(229)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case VlangParserINDEXOF:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(190)
			p.Match(VlangParserINDEXOF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(191)
			p.Match(VlangParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(200)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2815334427852800) != 0 {
			{
				p.SetState(192)
				p.expresion(0)
			}
			p.SetState(197)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == VlangParserCOMMA {
				{
					p.SetState(193)
					p.Match(VlangParserCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(194)
					p.expresion(0)
				}

				p.SetState(199)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(202)
			p.Match(VlangParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VlangParserJOIN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(203)
			p.Match(VlangParserJOIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(204)
			p.Match(VlangParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(213)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2815334427852800) != 0 {
			{
				p.SetState(205)
				p.expresion(0)
			}
			p.SetState(210)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == VlangParserCOMMA {
				{
					p.SetState(206)
					p.Match(VlangParserCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(207)
					p.expresion(0)
				}

				p.SetState(212)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(215)
			p.Match(VlangParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VlangParserID:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(216)
			p.Match(VlangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(217)
			p.Match(VlangParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(226)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2815334427852800) != 0 {
			{
				p.SetState(218)
				p.expresion(0)
			}
			p.SetState(223)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == VlangParserCOMMA {
				{
					p.SetState(219)
					p.Match(VlangParserCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(220)
					p.expresion(0)
				}

				p.SetState(225)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(228)
			p.Match(VlangParserRPAREN)
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

// IWhileDclContext is an interface to support dynamic dispatch.
type IWhileDclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAREN() antlr.TerminalNode
	Expresion() IExpresionContext
	RPAREN() antlr.TerminalNode
	LBRACK() antlr.TerminalNode
	RBRACK() antlr.TerminalNode
	AllDeclaraciones() []IDeclaracionesContext
	Declaraciones(i int) IDeclaracionesContext

	// IsWhileDclContext differentiates from other interfaces.
	IsWhileDclContext()
}

type WhileDclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhileDclContext() *WhileDclContext {
	var p = new(WhileDclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_whileDcl
	return p
}

func InitEmptyWhileDclContext(p *WhileDclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_whileDcl
}

func (*WhileDclContext) IsWhileDclContext() {}

func NewWhileDclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhileDclContext {
	var p = new(WhileDclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_whileDcl

	return p
}

func (s *WhileDclContext) GetParser() antlr.Parser { return s.parser }

func (s *WhileDclContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(VlangParserLPAREN, 0)
}

func (s *WhileDclContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *WhileDclContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(VlangParserRPAREN, 0)
}

func (s *WhileDclContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(VlangParserLBRACK, 0)
}

func (s *WhileDclContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(VlangParserRBRACK, 0)
}

func (s *WhileDclContext) AllDeclaraciones() []IDeclaracionesContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDeclaracionesContext); ok {
			len++
		}
	}

	tst := make([]IDeclaracionesContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDeclaracionesContext); ok {
			tst[i] = t.(IDeclaracionesContext)
			i++
		}
	}

	return tst
}

func (s *WhileDclContext) Declaraciones(i int) IDeclaracionesContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclaracionesContext); ok {
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

	return t.(IDeclaracionesContext)
}

func (s *WhileDclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileDclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhileDclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterWhileDcl(s)
	}
}

func (s *WhileDclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitWhileDcl(s)
	}
}

func (s *WhileDclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitWhileDcl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) WhileDcl() (localctx IWhileDclContext) {
	localctx = NewWhileDclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, VlangParserRULE_whileDcl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(231)
		p.Match(VlangParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(232)
		p.Match(VlangParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(233)
		p.expresion(0)
	}
	{
		p.SetState(234)
		p.Match(VlangParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(235)
		p.Match(VlangParserLBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(239)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2815334696288584) != 0 {
		{
			p.SetState(236)
			p.Declaraciones()
		}

		p.SetState(241)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(242)
		p.Match(VlangParserRBRACK)
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

// IExpresionContext is an interface to support dynamic dispatch.
type IExpresionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExpresionContext differentiates from other interfaces.
	IsExpresionContext()
}

type ExpresionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpresionContext() *ExpresionContext {
	var p = new(ExpresionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_expresion
	return p
}

func InitEmptyExpresionContext(p *ExpresionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_expresion
}

func (*ExpresionContext) IsExpresionContext() {}

func NewExpresionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpresionContext {
	var p = new(ExpresionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_expresion

	return p
}

func (s *ExpresionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpresionContext) CopyAll(ctx *ExpresionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExpresionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpresionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type MultdivmodContext struct {
	ExpresionContext
	op antlr.Token
}

func NewMultdivmodContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MultdivmodContext {
	var p = new(MultdivmodContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *MultdivmodContext) GetOp() antlr.Token { return s.op }

func (s *MultdivmodContext) SetOp(v antlr.Token) { s.op = v }

func (s *MultdivmodContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultdivmodContext) AllExpresion() []IExpresionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpresionContext); ok {
			len++
		}
	}

	tst := make([]IExpresionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpresionContext); ok {
			tst[i] = t.(IExpresionContext)
			i++
		}
	}

	return tst
}

func (s *MultdivmodContext) Expresion(i int) IExpresionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
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

	return t.(IExpresionContext)
}

func (s *MultdivmodContext) MUL() antlr.TerminalNode {
	return s.GetToken(VlangParserMUL, 0)
}

func (s *MultdivmodContext) DIV() antlr.TerminalNode {
	return s.GetToken(VlangParserDIV, 0)
}

func (s *MultdivmodContext) MOD() antlr.TerminalNode {
	return s.GetToken(VlangParserMOD, 0)
}

func (s *MultdivmodContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterMultdivmod(s)
	}
}

func (s *MultdivmodContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitMultdivmod(s)
	}
}

func (s *MultdivmodContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitMultdivmod(s)

	default:
		return t.VisitChildren(s)
	}
}

type IncredecrContext struct {
	ExpresionContext
}

func NewIncredecrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IncredecrContext {
	var p = new(IncredecrContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *IncredecrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IncredecrContext) Incredecre() IIncredecreContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIncredecreContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIncredecreContext)
}

func (s *IncredecrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterIncredecr(s)
	}
}

func (s *IncredecrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitIncredecr(s)
	}
}

func (s *IncredecrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitIncredecr(s)

	default:
		return t.VisitChildren(s)
	}
}

type OPERADORESLOGICOSContext struct {
	ExpresionContext
	op antlr.Token
}

func NewOPERADORESLOGICOSContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OPERADORESLOGICOSContext {
	var p = new(OPERADORESLOGICOSContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *OPERADORESLOGICOSContext) GetOp() antlr.Token { return s.op }

func (s *OPERADORESLOGICOSContext) SetOp(v antlr.Token) { s.op = v }

func (s *OPERADORESLOGICOSContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OPERADORESLOGICOSContext) AllExpresion() []IExpresionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpresionContext); ok {
			len++
		}
	}

	tst := make([]IExpresionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpresionContext); ok {
			tst[i] = t.(IExpresionContext)
			i++
		}
	}

	return tst
}

func (s *OPERADORESLOGICOSContext) Expresion(i int) IExpresionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
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

	return t.(IExpresionContext)
}

func (s *OPERADORESLOGICOSContext) AND() antlr.TerminalNode {
	return s.GetToken(VlangParserAND, 0)
}

func (s *OPERADORESLOGICOSContext) OR() antlr.TerminalNode {
	return s.GetToken(VlangParserOR, 0)
}

func (s *OPERADORESLOGICOSContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterOPERADORESLOGICOS(s)
	}
}

func (s *OPERADORESLOGICOSContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitOPERADORESLOGICOS(s)
	}
}

func (s *OPERADORESLOGICOSContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitOPERADORESLOGICOS(s)

	default:
		return t.VisitChildren(s)
	}
}

type OrContext struct {
	ExpresionContext
}

func NewOrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OrContext {
	var p = new(OrContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *OrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrContext) AllExpresion() []IExpresionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpresionContext); ok {
			len++
		}
	}

	tst := make([]IExpresionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpresionContext); ok {
			tst[i] = t.(IExpresionContext)
			i++
		}
	}

	return tst
}

func (s *OrContext) Expresion(i int) IExpresionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
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

	return t.(IExpresionContext)
}

func (s *OrContext) OR() antlr.TerminalNode {
	return s.GetToken(VlangParserOR, 0)
}

func (s *OrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterOr(s)
	}
}

func (s *OrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitOr(s)
	}
}

func (s *OrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitOr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ValorexprContext struct {
	ExpresionContext
}

func NewValorexprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValorexprContext {
	var p = new(ValorexprContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *ValorexprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValorexprContext) Valor() IValorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValorContext)
}

func (s *ValorexprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterValorexpr(s)
	}
}

func (s *ValorexprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitValorexpr(s)
	}
}

func (s *ValorexprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitValorexpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type IgualdadContext struct {
	ExpresionContext
	op antlr.Token
}

func NewIgualdadContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IgualdadContext {
	var p = new(IgualdadContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *IgualdadContext) GetOp() antlr.Token { return s.op }

func (s *IgualdadContext) SetOp(v antlr.Token) { s.op = v }

func (s *IgualdadContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IgualdadContext) AllExpresion() []IExpresionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpresionContext); ok {
			len++
		}
	}

	tst := make([]IExpresionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpresionContext); ok {
			tst[i] = t.(IExpresionContext)
			i++
		}
	}

	return tst
}

func (s *IgualdadContext) Expresion(i int) IExpresionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
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

	return t.(IExpresionContext)
}

func (s *IgualdadContext) EQ() antlr.TerminalNode {
	return s.GetToken(VlangParserEQ, 0)
}

func (s *IgualdadContext) NEQ() antlr.TerminalNode {
	return s.GetToken(VlangParserNEQ, 0)
}

func (s *IgualdadContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterIgualdad(s)
	}
}

func (s *IgualdadContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitIgualdad(s)
	}
}

func (s *IgualdadContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitIgualdad(s)

	default:
		return t.VisitChildren(s)
	}
}

type SliceCreacionvContext struct {
	ExpresionContext
}

func NewSliceCreacionvContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SliceCreacionvContext {
	var p = new(SliceCreacionvContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *SliceCreacionvContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SliceCreacionvContext) TipoSlice() ITipoSliceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipoSliceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipoSliceContext)
}

func (s *SliceCreacionvContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(VlangParserLBRACE, 0)
}

func (s *SliceCreacionvContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(VlangParserRBRACE, 0)
}

func (s *SliceCreacionvContext) ListaExpresiones() IListaExpresionesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListaExpresionesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListaExpresionesContext)
}

func (s *SliceCreacionvContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterSliceCreacionv(s)
	}
}

func (s *SliceCreacionvContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitSliceCreacionv(s)
	}
}

func (s *SliceCreacionvContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitSliceCreacionv(s)

	default:
		return t.VisitChildren(s)
	}
}

type LlamadaFuncionExprContext struct {
	ExpresionContext
}

func NewLlamadaFuncionExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LlamadaFuncionExprContext {
	var p = new(LlamadaFuncionExprContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *LlamadaFuncionExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LlamadaFuncionExprContext) LlamadaFuncion() ILlamadaFuncionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILlamadaFuncionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILlamadaFuncionContext)
}

func (s *LlamadaFuncionExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterLlamadaFuncionExpr(s)
	}
}

func (s *LlamadaFuncionExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitLlamadaFuncionExpr(s)
	}
}

func (s *LlamadaFuncionExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitLlamadaFuncionExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpdotexpContext struct {
	ExpresionContext
}

func NewExpdotexpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpdotexpContext {
	var p = new(ExpdotexpContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *ExpdotexpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpdotexpContext) ID() antlr.TerminalNode {
	return s.GetToken(VlangParserID, 0)
}

func (s *ExpdotexpContext) DOT() antlr.TerminalNode {
	return s.GetToken(VlangParserDOT, 0)
}

func (s *ExpdotexpContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *ExpdotexpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterExpdotexp(s)
	}
}

func (s *ExpdotexpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitExpdotexp(s)
	}
}

func (s *ExpdotexpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitExpdotexp(s)

	default:
		return t.VisitChildren(s)
	}
}

type RelacionalesContext struct {
	ExpresionContext
	op antlr.Token
}

func NewRelacionalesContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RelacionalesContext {
	var p = new(RelacionalesContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *RelacionalesContext) GetOp() antlr.Token { return s.op }

func (s *RelacionalesContext) SetOp(v antlr.Token) { s.op = v }

func (s *RelacionalesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelacionalesContext) AllExpresion() []IExpresionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpresionContext); ok {
			len++
		}
	}

	tst := make([]IExpresionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpresionContext); ok {
			tst[i] = t.(IExpresionContext)
			i++
		}
	}

	return tst
}

func (s *RelacionalesContext) Expresion(i int) IExpresionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
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

	return t.(IExpresionContext)
}

func (s *RelacionalesContext) LT() antlr.TerminalNode {
	return s.GetToken(VlangParserLT, 0)
}

func (s *RelacionalesContext) LE() antlr.TerminalNode {
	return s.GetToken(VlangParserLE, 0)
}

func (s *RelacionalesContext) GE() antlr.TerminalNode {
	return s.GetToken(VlangParserGE, 0)
}

func (s *RelacionalesContext) GT() antlr.TerminalNode {
	return s.GetToken(VlangParserGT, 0)
}

func (s *RelacionalesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterRelacionales(s)
	}
}

func (s *RelacionalesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitRelacionales(s)
	}
}

func (s *RelacionalesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitRelacionales(s)

	default:
		return t.VisitChildren(s)
	}
}

type CorchetesexpreContext struct {
	ExpresionContext
}

func NewCorchetesexpreContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CorchetesexpreContext {
	var p = new(CorchetesexpreContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *CorchetesexpreContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CorchetesexpreContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(VlangParserLBRACK, 0)
}

func (s *CorchetesexpreContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *CorchetesexpreContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(VlangParserRBRACK, 0)
}

func (s *CorchetesexpreContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterCorchetesexpre(s)
	}
}

func (s *CorchetesexpreContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitCorchetesexpre(s)
	}
}

func (s *CorchetesexpreContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitCorchetesexpre(s)

	default:
		return t.VisitChildren(s)
	}
}

type UnarioContext struct {
	ExpresionContext
	op antlr.Token
}

func NewUnarioContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnarioContext {
	var p = new(UnarioContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *UnarioContext) GetOp() antlr.Token { return s.op }

func (s *UnarioContext) SetOp(v antlr.Token) { s.op = v }

func (s *UnarioContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnarioContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *UnarioContext) NOT() antlr.TerminalNode {
	return s.GetToken(VlangParserNOT, 0)
}

func (s *UnarioContext) MINUS() antlr.TerminalNode {
	return s.GetToken(VlangParserMINUS, 0)
}

func (s *UnarioContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterUnario(s)
	}
}

func (s *UnarioContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitUnario(s)
	}
}

func (s *UnarioContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitUnario(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParentesisexpreContext struct {
	ExpresionContext
}

func NewParentesisexpreContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParentesisexpreContext {
	var p = new(ParentesisexpreContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *ParentesisexpreContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParentesisexpreContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(VlangParserLPAREN, 0)
}

func (s *ParentesisexpreContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *ParentesisexpreContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(VlangParserRPAREN, 0)
}

func (s *ParentesisexpreContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterParentesisexpre(s)
	}
}

func (s *ParentesisexpreContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitParentesisexpre(s)
	}
}

func (s *ParentesisexpreContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitParentesisexpre(s)

	default:
		return t.VisitChildren(s)
	}
}

type IMCPLICITContext struct {
	ExpresionContext
	op antlr.Token
}

func NewIMCPLICITContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IMCPLICITContext {
	var p = new(IMCPLICITContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *IMCPLICITContext) GetOp() antlr.Token { return s.op }

func (s *IMCPLICITContext) SetOp(v antlr.Token) { s.op = v }

func (s *IMCPLICITContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IMCPLICITContext) ID() antlr.TerminalNode {
	return s.GetToken(VlangParserID, 0)
}

func (s *IMCPLICITContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *IMCPLICITContext) SUMAIMPLICITA() antlr.TerminalNode {
	return s.GetToken(VlangParserSUMAIMPLICITA, 0)
}

func (s *IMCPLICITContext) RESTOIMPLICITO() antlr.TerminalNode {
	return s.GetToken(VlangParserRESTOIMPLICITO, 0)
}

func (s *IMCPLICITContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterIMCPLICIT(s)
	}
}

func (s *IMCPLICITContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitIMCPLICIT(s)
	}
}

func (s *IMCPLICITContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitIMCPLICIT(s)

	default:
		return t.VisitChildren(s)
	}
}

type SumresContext struct {
	ExpresionContext
	op antlr.Token
}

func NewSumresContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SumresContext {
	var p = new(SumresContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *SumresContext) GetOp() antlr.Token { return s.op }

func (s *SumresContext) SetOp(v antlr.Token) { s.op = v }

func (s *SumresContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SumresContext) AllExpresion() []IExpresionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpresionContext); ok {
			len++
		}
	}

	tst := make([]IExpresionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpresionContext); ok {
			tst[i] = t.(IExpresionContext)
			i++
		}
	}

	return tst
}

func (s *SumresContext) Expresion(i int) IExpresionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
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

	return t.(IExpresionContext)
}

func (s *SumresContext) PLUS() antlr.TerminalNode {
	return s.GetToken(VlangParserPLUS, 0)
}

func (s *SumresContext) MINUS() antlr.TerminalNode {
	return s.GetToken(VlangParserMINUS, 0)
}

func (s *SumresContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterSumres(s)
	}
}

func (s *SumresContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitSumres(s)
	}
}

func (s *SumresContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitSumres(s)

	default:
		return t.VisitChildren(s)
	}
}

type AsignacionLUEGOContext struct {
	ExpresionContext
}

func NewAsignacionLUEGOContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AsignacionLUEGOContext {
	var p = new(AsignacionLUEGOContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *AsignacionLUEGOContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsignacionLUEGOContext) ID() antlr.TerminalNode {
	return s.GetToken(VlangParserID, 0)
}

func (s *AsignacionLUEGOContext) TIPO() antlr.TerminalNode {
	return s.GetToken(VlangParserTIPO, 0)
}

func (s *AsignacionLUEGOContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(VlangParserASSIGN, 0)
}

func (s *AsignacionLUEGOContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *AsignacionLUEGOContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterAsignacionLUEGO(s)
	}
}

func (s *AsignacionLUEGOContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitAsignacionLUEGO(s)
	}
}

func (s *AsignacionLUEGOContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitAsignacionLUEGO(s)

	default:
		return t.VisitChildren(s)
	}
}

type IdContext struct {
	ExpresionContext
}

func NewIdContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdContext {
	var p = new(IdContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *IdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdContext) ID() antlr.TerminalNode {
	return s.GetToken(VlangParserID, 0)
}

func (s *IdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterId(s)
	}
}

func (s *IdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitId(s)
	}
}

func (s *IdContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitId(s)

	default:
		return t.VisitChildren(s)
	}
}

type Expdotexp1Context struct {
	ExpresionContext
}

func NewExpdotexp1Context(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expdotexp1Context {
	var p = new(Expdotexp1Context)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *Expdotexp1Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expdotexp1Context) AllID() []antlr.TerminalNode {
	return s.GetTokens(VlangParserID)
}

func (s *Expdotexp1Context) ID(i int) antlr.TerminalNode {
	return s.GetToken(VlangParserID, i)
}

func (s *Expdotexp1Context) DOT() antlr.TerminalNode {
	return s.GetToken(VlangParserDOT, 0)
}

func (s *Expdotexp1Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterExpdotexp1(s)
	}
}

func (s *Expdotexp1Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitExpdotexp1(s)
	}
}

func (s *Expdotexp1Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitExpdotexp1(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Expresion() (localctx IExpresionContext) {
	return p.expresion(0)
}

func (p *VlangParser) expresion(_p int) (localctx IExpresionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpresionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpresionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 30
	p.EnterRecursionRule(localctx, 30, VlangParserRULE_expresion, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(279)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 33, p.GetParserRuleContext()) {
	case 1:
		localctx = NewValorexprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(245)
			p.Valor()
		}

	case 2:
		localctx = NewParentesisexpreContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(246)
			p.Match(VlangParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(247)
			p.expresion(0)
		}
		{
			p.SetState(248)
			p.Match(VlangParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewCorchetesexpreContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(250)
			p.Match(VlangParserLBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(251)
			p.expresion(0)
		}
		{
			p.SetState(252)
			p.Match(VlangParserRBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewSliceCreacionvContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(254)
			p.TipoSlice()
		}
		{
			p.SetState(255)
			p.Match(VlangParserLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(257)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2815334427852800) != 0 {
			{
				p.SetState(256)
				p.ListaExpresiones()
			}

		}
		{
			p.SetState(259)
			p.Match(VlangParserRBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		localctx = NewLlamadaFuncionExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(261)
			p.LlamadaFuncion()
		}

	case 6:
		localctx = NewUnarioContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(262)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*UnarioContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == VlangParserMINUS || _la == VlangParserNOT) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*UnarioContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(263)
			p.expresion(13)
		}

	case 7:
		localctx = NewIdContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(264)
			p.Match(VlangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 8:
		localctx = NewIncredecrContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(265)
			p.Incredecre()
		}

	case 9:
		localctx = NewExpdotexp1Context(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(266)
			p.Match(VlangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(267)
			p.Match(VlangParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(268)
			p.Match(VlangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 10:
		localctx = NewExpdotexpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(269)
			p.Match(VlangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(270)
			p.Match(VlangParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(271)
			p.expresion(3)
		}

	case 11:
		localctx = NewAsignacionLUEGOContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(272)
			p.Match(VlangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(273)
			p.Match(VlangParserTIPO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(274)
			p.Match(VlangParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(275)
			p.expresion(2)
		}

	case 12:
		localctx = NewIMCPLICITContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(276)
			p.Match(VlangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(277)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*IMCPLICITContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == VlangParserSUMAIMPLICITA || _la == VlangParserRESTOIMPLICITO) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*IMCPLICITContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(278)
			p.expresion(1)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(301)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 35, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(299)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 34, p.GetParserRuleContext()) {
			case 1:
				localctx = NewOPERADORESLOGICOSContext(p, NewExpresionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, VlangParserRULE_expresion)
				p.SetState(281)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
					goto errorExit
				}
				{
					p.SetState(282)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*OPERADORESLOGICOSContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == VlangParserOR || _la == VlangParserAND) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*OPERADORESLOGICOSContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(283)
					p.expresion(13)
				}

			case 2:
				localctx = NewMultdivmodContext(p, NewExpresionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, VlangParserRULE_expresion)
				p.SetState(284)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(285)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*MultdivmodContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&481036337152) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*MultdivmodContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(286)
					p.expresion(12)
				}

			case 3:
				localctx = NewSumresContext(p, NewExpresionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, VlangParserRULE_expresion)
				p.SetState(287)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(288)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*SumresContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == VlangParserPLUS || _la == VlangParserMINUS) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*SumresContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(289)
					p.expresion(11)
				}

			case 4:
				localctx = NewRelacionalesContext(p, NewExpresionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, VlangParserRULE_expresion)
				p.SetState(290)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(291)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*RelacionalesContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&263882790666240) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*RelacionalesContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(292)
					p.expresion(10)
				}

			case 5:
				localctx = NewIgualdadContext(p, NewExpresionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, VlangParserRULE_expresion)
				p.SetState(293)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(294)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*IgualdadContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == VlangParserEQ || _la == VlangParserNEQ) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*IgualdadContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(295)
					p.expresion(9)
				}

			case 6:
				localctx = NewOrContext(p, NewExpresionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, VlangParserRULE_expresion)
				p.SetState(296)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(297)
					p.Match(VlangParserOR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(298)
					p.expresion(8)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(303)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 35, p.GetParserRuleContext())
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

// IParametrosContext is an interface to support dynamic dispatch.
type IParametrosContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpresion() []IExpresionContext
	Expresion(i int) IExpresionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsParametrosContext differentiates from other interfaces.
	IsParametrosContext()
}

type ParametrosContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParametrosContext() *ParametrosContext {
	var p = new(ParametrosContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_parametros
	return p
}

func InitEmptyParametrosContext(p *ParametrosContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_parametros
}

func (*ParametrosContext) IsParametrosContext() {}

func NewParametrosContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParametrosContext {
	var p = new(ParametrosContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_parametros

	return p
}

func (s *ParametrosContext) GetParser() antlr.Parser { return s.parser }

func (s *ParametrosContext) AllExpresion() []IExpresionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpresionContext); ok {
			len++
		}
	}

	tst := make([]IExpresionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpresionContext); ok {
			tst[i] = t.(IExpresionContext)
			i++
		}
	}

	return tst
}

func (s *ParametrosContext) Expresion(i int) IExpresionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
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

	return t.(IExpresionContext)
}

func (s *ParametrosContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(VlangParserCOMMA)
}

func (s *ParametrosContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(VlangParserCOMMA, i)
}

func (s *ParametrosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParametrosContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParametrosContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterParametros(s)
	}
}

func (s *ParametrosContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitParametros(s)
	}
}

func (s *ParametrosContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitParametros(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Parametros() (localctx IParametrosContext) {
	localctx = NewParametrosContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, VlangParserRULE_parametros)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(304)
		p.expresion(0)
	}
	p.SetState(309)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == VlangParserCOMMA {
		{
			p.SetState(305)
			p.Match(VlangParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(306)
			p.expresion(0)
		}

		p.SetState(311)
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

// IValoresContext is an interface to support dynamic dispatch.
type IValoresContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Valor() IValorContext

	// IsValoresContext differentiates from other interfaces.
	IsValoresContext()
}

type ValoresContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValoresContext() *ValoresContext {
	var p = new(ValoresContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_valores
	return p
}

func InitEmptyValoresContext(p *ValoresContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_valores
}

func (*ValoresContext) IsValoresContext() {}

func NewValoresContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValoresContext {
	var p = new(ValoresContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_valores

	return p
}

func (s *ValoresContext) GetParser() antlr.Parser { return s.parser }

func (s *ValoresContext) Valor() IValorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValorContext)
}

func (s *ValoresContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValoresContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValoresContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterValores(s)
	}
}

func (s *ValoresContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitValores(s)
	}
}

func (s *ValoresContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitValores(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Valores() (localctx IValoresContext) {
	localctx = NewValoresContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, VlangParserRULE_valores)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(312)
		p.Valor()
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

// IValorContext is an interface to support dynamic dispatch.
type IValorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsValorContext differentiates from other interfaces.
	IsValorContext()
}

type ValorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValorContext() *ValorContext {
	var p = new(ValorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_valor
	return p
}

func InitEmptyValorContext(p *ValorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_valor
}

func (*ValorContext) IsValorContext() {}

func NewValorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValorContext {
	var p = new(ValorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_valor

	return p
}

func (s *ValorContext) GetParser() antlr.Parser { return s.parser }

func (s *ValorContext) CopyAll(ctx *ValorContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ValorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ValorDecimalContext struct {
	ValorContext
}

func NewValorDecimalContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValorDecimalContext {
	var p = new(ValorDecimalContext)

	InitEmptyValorContext(&p.ValorContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValorContext))

	return p
}

func (s *ValorDecimalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValorDecimalContext) DECIMAL() antlr.TerminalNode {
	return s.GetToken(VlangParserDECIMAL, 0)
}

func (s *ValorDecimalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterValorDecimal(s)
	}
}

func (s *ValorDecimalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitValorDecimal(s)
	}
}

func (s *ValorDecimalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitValorDecimal(s)

	default:
		return t.VisitChildren(s)
	}
}

type ValorEnteroContext struct {
	ValorContext
}

func NewValorEnteroContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValorEnteroContext {
	var p = new(ValorEnteroContext)

	InitEmptyValorContext(&p.ValorContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValorContext))

	return p
}

func (s *ValorEnteroContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValorEnteroContext) ENTERO() antlr.TerminalNode {
	return s.GetToken(VlangParserENTERO, 0)
}

func (s *ValorEnteroContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterValorEntero(s)
	}
}

func (s *ValorEnteroContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitValorEntero(s)
	}
}

func (s *ValorEnteroContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitValorEntero(s)

	default:
		return t.VisitChildren(s)
	}
}

type ValorBooleanoContext struct {
	ValorContext
}

func NewValorBooleanoContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValorBooleanoContext {
	var p = new(ValorBooleanoContext)

	InitEmptyValorContext(&p.ValorContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValorContext))

	return p
}

func (s *ValorBooleanoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValorBooleanoContext) BOOLEANO() antlr.TerminalNode {
	return s.GetToken(VlangParserBOOLEANO, 0)
}

func (s *ValorBooleanoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterValorBooleano(s)
	}
}

func (s *ValorBooleanoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitValorBooleano(s)
	}
}

func (s *ValorBooleanoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitValorBooleano(s)

	default:
		return t.VisitChildren(s)
	}
}

type ValorCaracterContext struct {
	ValorContext
}

func NewValorCaracterContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValorCaracterContext {
	var p = new(ValorCaracterContext)

	InitEmptyValorContext(&p.ValorContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValorContext))

	return p
}

func (s *ValorCaracterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValorCaracterContext) CARACTER() antlr.TerminalNode {
	return s.GetToken(VlangParserCARACTER, 0)
}

func (s *ValorCaracterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterValorCaracter(s)
	}
}

func (s *ValorCaracterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitValorCaracter(s)
	}
}

func (s *ValorCaracterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitValorCaracter(s)

	default:
		return t.VisitChildren(s)
	}
}

type ValorCadenaContext struct {
	ValorContext
}

func NewValorCadenaContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValorCadenaContext {
	var p = new(ValorCadenaContext)

	InitEmptyValorContext(&p.ValorContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValorContext))

	return p
}

func (s *ValorCadenaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValorCadenaContext) CADENA() antlr.TerminalNode {
	return s.GetToken(VlangParserCADENA, 0)
}

func (s *ValorCadenaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterValorCadena(s)
	}
}

func (s *ValorCadenaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitValorCadena(s)
	}
}

func (s *ValorCadenaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitValorCadena(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Valor() (localctx IValorContext) {
	localctx = NewValorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, VlangParserRULE_valor)
	p.SetState(319)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case VlangParserENTERO:
		localctx = NewValorEnteroContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(314)
			p.Match(VlangParserENTERO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VlangParserDECIMAL:
		localctx = NewValorDecimalContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(315)
			p.Match(VlangParserDECIMAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VlangParserCADENA:
		localctx = NewValorCadenaContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(316)
			p.Match(VlangParserCADENA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VlangParserBOOLEANO:
		localctx = NewValorBooleanoContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(317)
			p.Match(VlangParserBOOLEANO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VlangParserCARACTER:
		localctx = NewValorCaracterContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(318)
			p.Match(VlangParserCARACTER)
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

// IListaExpresionesContext is an interface to support dynamic dispatch.
type IListaExpresionesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpresion() []IExpresionContext
	Expresion(i int) IExpresionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsListaExpresionesContext differentiates from other interfaces.
	IsListaExpresionesContext()
}

type ListaExpresionesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListaExpresionesContext() *ListaExpresionesContext {
	var p = new(ListaExpresionesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_listaExpresiones
	return p
}

func InitEmptyListaExpresionesContext(p *ListaExpresionesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_listaExpresiones
}

func (*ListaExpresionesContext) IsListaExpresionesContext() {}

func NewListaExpresionesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListaExpresionesContext {
	var p = new(ListaExpresionesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_listaExpresiones

	return p
}

func (s *ListaExpresionesContext) GetParser() antlr.Parser { return s.parser }

func (s *ListaExpresionesContext) AllExpresion() []IExpresionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpresionContext); ok {
			len++
		}
	}

	tst := make([]IExpresionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpresionContext); ok {
			tst[i] = t.(IExpresionContext)
			i++
		}
	}

	return tst
}

func (s *ListaExpresionesContext) Expresion(i int) IExpresionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
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

	return t.(IExpresionContext)
}

func (s *ListaExpresionesContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(VlangParserCOMMA)
}

func (s *ListaExpresionesContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(VlangParserCOMMA, i)
}

func (s *ListaExpresionesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListaExpresionesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListaExpresionesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterListaExpresiones(s)
	}
}

func (s *ListaExpresionesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitListaExpresiones(s)
	}
}

func (s *ListaExpresionesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitListaExpresiones(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) ListaExpresiones() (localctx IListaExpresionesContext) {
	localctx = NewListaExpresionesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, VlangParserRULE_listaExpresiones)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(321)
		p.expresion(0)
	}
	p.SetState(326)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == VlangParserCOMMA {
		{
			p.SetState(322)
			p.Match(VlangParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(323)
			p.expresion(0)
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

// IIncredecreContext is an interface to support dynamic dispatch.
type IIncredecreContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsIncredecreContext differentiates from other interfaces.
	IsIncredecreContext()
}

type IncredecreContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIncredecreContext() *IncredecreContext {
	var p = new(IncredecreContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_incredecre
	return p
}

func InitEmptyIncredecreContext(p *IncredecreContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_incredecre
}

func (*IncredecreContext) IsIncredecreContext() {}

func NewIncredecreContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IncredecreContext {
	var p = new(IncredecreContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_incredecre

	return p
}

func (s *IncredecreContext) GetParser() antlr.Parser { return s.parser }

func (s *IncredecreContext) CopyAll(ctx *IncredecreContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *IncredecreContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IncredecreContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IncrementoContext struct {
	IncredecreContext
}

func NewIncrementoContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IncrementoContext {
	var p = new(IncrementoContext)

	InitEmptyIncredecreContext(&p.IncredecreContext)
	p.parser = parser
	p.CopyAll(ctx.(*IncredecreContext))

	return p
}

func (s *IncrementoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IncrementoContext) ID() antlr.TerminalNode {
	return s.GetToken(VlangParserID, 0)
}

func (s *IncrementoContext) INC() antlr.TerminalNode {
	return s.GetToken(VlangParserINC, 0)
}

func (s *IncrementoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterIncremento(s)
	}
}

func (s *IncrementoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitIncremento(s)
	}
}

func (s *IncrementoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitIncremento(s)

	default:
		return t.VisitChildren(s)
	}
}

type DecrementoContext struct {
	IncredecreContext
}

func NewDecrementoContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DecrementoContext {
	var p = new(DecrementoContext)

	InitEmptyIncredecreContext(&p.IncredecreContext)
	p.parser = parser
	p.CopyAll(ctx.(*IncredecreContext))

	return p
}

func (s *DecrementoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DecrementoContext) ID() antlr.TerminalNode {
	return s.GetToken(VlangParserID, 0)
}

func (s *DecrementoContext) DEC() antlr.TerminalNode {
	return s.GetToken(VlangParserDEC, 0)
}

func (s *DecrementoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterDecremento(s)
	}
}

func (s *DecrementoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitDecremento(s)
	}
}

func (s *DecrementoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitDecremento(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Incredecre() (localctx IIncredecreContext) {
	localctx = NewIncredecreContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, VlangParserRULE_incredecre)
	p.SetState(333)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 39, p.GetParserRuleContext()) {
	case 1:
		localctx = NewIncrementoContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(329)
			p.Match(VlangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(330)
			p.Match(VlangParserINC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewDecrementoContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(331)
			p.Match(VlangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(332)
			p.Match(VlangParserDEC)
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

func (p *VlangParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 15:
		var t *ExpresionContext = nil
		if localctx != nil {
			t = localctx.(*ExpresionContext)
		}
		return p.Expresion_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *VlangParser) Expresion_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 7)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
