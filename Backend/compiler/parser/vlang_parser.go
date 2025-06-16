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
		"", "'fn'", "'main'", "'mut'", "'range'", "'case'", "'default'", "'struct'",
		"'while'", "", "", "'Atoi'", "'parseFloat'", "'typeOf'", "'len'", "'cap'",
		"'append'", "'if'", "'else'", "'for'", "'switch'", "'indexOf'", "'join'",
		"'break'", "'continue'", "'return'", "", "", "", "", "", "'int'", "'float64'",
		"'string'", "'bool'", "'rune'", "'print'", "", "'++'", "'--'", "'+='",
		"'-='", "'+'", "'-'", "'*'", "'/'", "'%'", "'!'", "'||'", "'&&'", "'=='",
		"'!='", "'<='", "'>='", "'<'", "'>'", "'='", "'('", "')'", "'['", "']'",
		"'{'", "'}'", "';'", "':'", "'.'", "','",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "TIPO", "CASTEOS", "ATOI", "PARSEFLOAT",
		"TYPEOF", "LEN", "CAP", "APPEND", "IF", "ELSE", "FOR", "SWITCH", "INDEXOF",
		"JOIN", "BREAK", "CONTINUE", "RETURN", "BOOLEANO", "ENTERO", "DECIMAL",
		"CADENA", "CARACTER", "TIPO_ENTERO", "TIPO_DECIMAL", "TIPO_CADENA",
		"TIPO_BOOLEANO", "TIPO_CHAR", "PRINT", "ID", "INC", "DEC", "SUMAIMPLICITA",
		"RESTOIMPLICITO", "PLUS", "MINUS", "MUL", "DIV", "MOD", "NOT", "OR",
		"AND", "EQ", "NEQ", "LE", "GE", "LT", "GT", "ASSIGN", "LPAREN", "RPAREN",
		"LBRACK", "RBRACK", "LBRACE", "RBRACE", "SEMICOLON", "COLON", "DOT",
		"COMMA", "WS", "LINE_COMMENT", "BLOCK_COMMENT",
	}
	staticData.RuleNames = []string{
		"programa", "funcMain", "funcDcl", "block", "declaraciones", "varDcl",
		"sliceTipo", "sliceInit", "stmt", "sentencias_control", "sentencias_transferencia",
		"ifDcl", "elseIfDcl", "elseCondicional", "forDcl", "asignacion", "switchDcl",
		"caseBlock", "defaultBlock", "llamadaFuncion", "funcCall", "parametrosFormales",
		"parametro", "parametrosReales", "structDcl", "atributosStruct", "atributoStruct",
		"listaAsignaciones", "asignacionStruct", "whileDcl", "expresion", "parametros",
		"valores", "valor", "listaExpresiones", "incredecre",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 69, 555, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 1, 0, 5, 0,
		74, 8, 0, 10, 0, 12, 0, 77, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 91, 8, 2, 1, 2, 1, 2, 3, 2, 95,
		8, 2, 1, 2, 1, 2, 1, 3, 1, 3, 5, 3, 101, 8, 3, 10, 3, 12, 3, 104, 9, 3,
		1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 113, 8, 4, 1, 5, 1, 5,
		1, 5, 3, 5, 118, 8, 5, 1, 5, 1, 5, 3, 5, 122, 8, 5, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 145, 8, 5, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 161,
		8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 3, 7, 169, 8, 7, 1, 7, 1, 7,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 5, 8, 178, 8, 8, 10, 8, 12, 8, 181, 9, 8,
		3, 8, 183, 8, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 189, 8, 8, 1, 9, 1, 9, 1,
		9, 1, 9, 3, 9, 195, 8, 9, 1, 10, 1, 10, 1, 10, 1, 10, 3, 10, 201, 8, 10,
		3, 10, 203, 8, 10, 1, 11, 1, 11, 1, 11, 1, 11, 5, 11, 209, 8, 11, 10, 11,
		12, 11, 212, 9, 11, 1, 11, 1, 11, 5, 11, 216, 8, 11, 10, 11, 12, 11, 219,
		9, 11, 1, 11, 3, 11, 222, 8, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 5,
		12, 229, 8, 12, 10, 12, 12, 12, 232, 9, 12, 1, 12, 1, 12, 1, 13, 1, 13,
		1, 13, 5, 13, 239, 8, 13, 10, 13, 12, 13, 242, 9, 13, 1, 13, 1, 13, 1,
		14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 3, 14, 252, 8, 14, 1, 14, 1, 14,
		1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1,
		14, 1, 14, 3, 14, 268, 8, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16,
		1, 16, 1, 16, 5, 16, 278, 8, 16, 10, 16, 12, 16, 281, 9, 16, 1, 16, 3,
		16, 284, 8, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 5, 17, 292, 8,
		17, 10, 17, 12, 17, 295, 9, 17, 1, 18, 1, 18, 1, 18, 5, 18, 300, 8, 18,
		10, 18, 12, 18, 303, 9, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 5, 19, 310,
		8, 19, 10, 19, 12, 19, 313, 9, 19, 3, 19, 315, 8, 19, 1, 19, 1, 19, 1,
		19, 1, 19, 1, 19, 1, 19, 5, 19, 323, 8, 19, 10, 19, 12, 19, 326, 9, 19,
		3, 19, 328, 8, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 5, 19, 336,
		8, 19, 10, 19, 12, 19, 339, 9, 19, 3, 19, 341, 8, 19, 1, 19, 1, 19, 1,
		19, 1, 19, 1, 19, 1, 19, 5, 19, 349, 8, 19, 10, 19, 12, 19, 352, 9, 19,
		3, 19, 354, 8, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 5, 19, 362,
		8, 19, 10, 19, 12, 19, 365, 9, 19, 3, 19, 367, 8, 19, 1, 19, 3, 19, 370,
		8, 19, 1, 20, 1, 20, 1, 20, 3, 20, 375, 8, 20, 1, 20, 1, 20, 1, 21, 1,
		21, 1, 21, 5, 21, 382, 8, 21, 10, 21, 12, 21, 385, 9, 21, 1, 22, 1, 22,
		1, 22, 1, 23, 1, 23, 1, 23, 5, 23, 393, 8, 23, 10, 23, 12, 23, 396, 9,
		23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 25, 4, 25, 405, 8, 25,
		11, 25, 12, 25, 406, 1, 26, 1, 26, 1, 26, 1, 26, 3, 26, 413, 8, 26, 1,
		27, 1, 27, 1, 27, 5, 27, 418, 8, 27, 10, 27, 12, 27, 421, 9, 27, 1, 28,
		1, 28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 5, 29, 433,
		8, 29, 10, 29, 12, 29, 436, 9, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1,
		30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30,
		1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1,
		30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30,
		1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1,
		30, 1, 30, 1, 30, 1, 30, 5, 30, 489, 8, 30, 10, 30, 12, 30, 492, 9, 30,
		1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 3, 30, 502, 8,
		30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30,
		1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 5, 30, 519, 8, 30, 10, 30, 12, 30, 522,
		9, 30, 1, 31, 1, 31, 1, 31, 5, 31, 527, 8, 31, 10, 31, 12, 31, 530, 9,
		31, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 3, 33, 539, 8, 33,
		1, 34, 1, 34, 1, 34, 5, 34, 544, 8, 34, 10, 34, 12, 34, 547, 9, 34, 1,
		35, 1, 35, 1, 35, 1, 35, 3, 35, 553, 8, 35, 1, 35, 0, 1, 60, 36, 0, 2,
		4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40,
		42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 0, 7, 2, 0,
		43, 43, 47, 47, 1, 0, 40, 41, 1, 0, 44, 46, 1, 0, 42, 43, 1, 0, 50, 51,
		1, 0, 52, 55, 1, 0, 48, 49, 609, 0, 75, 1, 0, 0, 0, 2, 80, 1, 0, 0, 0,
		4, 86, 1, 0, 0, 0, 6, 98, 1, 0, 0, 0, 8, 112, 1, 0, 0, 0, 10, 160, 1, 0,
		0, 0, 12, 162, 1, 0, 0, 0, 14, 166, 1, 0, 0, 0, 16, 188, 1, 0, 0, 0, 18,
		194, 1, 0, 0, 0, 20, 202, 1, 0, 0, 0, 22, 204, 1, 0, 0, 0, 24, 223, 1,
		0, 0, 0, 26, 235, 1, 0, 0, 0, 28, 267, 1, 0, 0, 0, 30, 269, 1, 0, 0, 0,
		32, 273, 1, 0, 0, 0, 34, 287, 1, 0, 0, 0, 36, 296, 1, 0, 0, 0, 38, 369,
		1, 0, 0, 0, 40, 371, 1, 0, 0, 0, 42, 378, 1, 0, 0, 0, 44, 386, 1, 0, 0,
		0, 46, 389, 1, 0, 0, 0, 48, 397, 1, 0, 0, 0, 50, 404, 1, 0, 0, 0, 52, 412,
		1, 0, 0, 0, 54, 414, 1, 0, 0, 0, 56, 422, 1, 0, 0, 0, 58, 426, 1, 0, 0,
		0, 60, 501, 1, 0, 0, 0, 62, 523, 1, 0, 0, 0, 64, 531, 1, 0, 0, 0, 66, 538,
		1, 0, 0, 0, 68, 540, 1, 0, 0, 0, 70, 552, 1, 0, 0, 0, 72, 74, 3, 8, 4,
		0, 73, 72, 1, 0, 0, 0, 74, 77, 1, 0, 0, 0, 75, 73, 1, 0, 0, 0, 75, 76,
		1, 0, 0, 0, 76, 78, 1, 0, 0, 0, 77, 75, 1, 0, 0, 0, 78, 79, 5, 0, 0, 1,
		79, 1, 1, 0, 0, 0, 80, 81, 5, 1, 0, 0, 81, 82, 5, 2, 0, 0, 82, 83, 5, 57,
		0, 0, 83, 84, 5, 58, 0, 0, 84, 85, 3, 6, 3, 0, 85, 3, 1, 0, 0, 0, 86, 87,
		5, 1, 0, 0, 87, 88, 5, 37, 0, 0, 88, 90, 5, 57, 0, 0, 89, 91, 3, 42, 21,
		0, 90, 89, 1, 0, 0, 0, 90, 91, 1, 0, 0, 0, 91, 92, 1, 0, 0, 0, 92, 94,
		5, 58, 0, 0, 93, 95, 5, 9, 0, 0, 94, 93, 1, 0, 0, 0, 94, 95, 1, 0, 0, 0,
		95, 96, 1, 0, 0, 0, 96, 97, 3, 6, 3, 0, 97, 5, 1, 0, 0, 0, 98, 102, 5,
		61, 0, 0, 99, 101, 3, 8, 4, 0, 100, 99, 1, 0, 0, 0, 101, 104, 1, 0, 0,
		0, 102, 100, 1, 0, 0, 0, 102, 103, 1, 0, 0, 0, 103, 105, 1, 0, 0, 0, 104,
		102, 1, 0, 0, 0, 105, 106, 5, 62, 0, 0, 106, 7, 1, 0, 0, 0, 107, 113, 3,
		10, 5, 0, 108, 113, 3, 16, 8, 0, 109, 113, 3, 4, 2, 0, 110, 113, 3, 2,
		1, 0, 111, 113, 3, 48, 24, 0, 112, 107, 1, 0, 0, 0, 112, 108, 1, 0, 0,
		0, 112, 109, 1, 0, 0, 0, 112, 110, 1, 0, 0, 0, 112, 111, 1, 0, 0, 0, 113,
		9, 1, 0, 0, 0, 114, 115, 5, 3, 0, 0, 115, 117, 5, 37, 0, 0, 116, 118, 5,
		9, 0, 0, 117, 116, 1, 0, 0, 0, 117, 118, 1, 0, 0, 0, 118, 121, 1, 0, 0,
		0, 119, 120, 5, 56, 0, 0, 120, 122, 3, 60, 30, 0, 121, 119, 1, 0, 0, 0,
		121, 122, 1, 0, 0, 0, 122, 161, 1, 0, 0, 0, 123, 124, 5, 3, 0, 0, 124,
		125, 5, 37, 0, 0, 125, 161, 3, 12, 6, 0, 126, 127, 5, 37, 0, 0, 127, 128,
		5, 56, 0, 0, 128, 129, 5, 37, 0, 0, 129, 130, 5, 61, 0, 0, 130, 131, 3,
		54, 27, 0, 131, 132, 5, 62, 0, 0, 132, 161, 1, 0, 0, 0, 133, 134, 5, 37,
		0, 0, 134, 135, 5, 56, 0, 0, 135, 136, 3, 12, 6, 0, 136, 137, 3, 14, 7,
		0, 137, 161, 1, 0, 0, 0, 138, 139, 5, 37, 0, 0, 139, 140, 5, 56, 0, 0,
		140, 161, 5, 37, 0, 0, 141, 144, 5, 37, 0, 0, 142, 143, 5, 56, 0, 0, 143,
		145, 3, 60, 30, 0, 144, 142, 1, 0, 0, 0, 144, 145, 1, 0, 0, 0, 145, 161,
		1, 0, 0, 0, 146, 147, 5, 37, 0, 0, 147, 148, 5, 56, 0, 0, 148, 149, 5,
		10, 0, 0, 149, 150, 5, 57, 0, 0, 150, 151, 3, 60, 30, 0, 151, 152, 5, 58,
		0, 0, 152, 161, 1, 0, 0, 0, 153, 154, 5, 37, 0, 0, 154, 155, 5, 59, 0,
		0, 155, 156, 3, 60, 30, 0, 156, 157, 5, 60, 0, 0, 157, 158, 5, 56, 0, 0,
		158, 159, 3, 60, 30, 0, 159, 161, 1, 0, 0, 0, 160, 114, 1, 0, 0, 0, 160,
		123, 1, 0, 0, 0, 160, 126, 1, 0, 0, 0, 160, 133, 1, 0, 0, 0, 160, 138,
		1, 0, 0, 0, 160, 141, 1, 0, 0, 0, 160, 146, 1, 0, 0, 0, 160, 153, 1, 0,
		0, 0, 161, 11, 1, 0, 0, 0, 162, 163, 5, 59, 0, 0, 163, 164, 5, 60, 0, 0,
		164, 165, 5, 9, 0, 0, 165, 13, 1, 0, 0, 0, 166, 168, 5, 61, 0, 0, 167,
		169, 3, 68, 34, 0, 168, 167, 1, 0, 0, 0, 168, 169, 1, 0, 0, 0, 169, 170,
		1, 0, 0, 0, 170, 171, 5, 62, 0, 0, 171, 15, 1, 0, 0, 0, 172, 173, 5, 36,
		0, 0, 173, 182, 5, 57, 0, 0, 174, 179, 3, 60, 30, 0, 175, 176, 5, 66, 0,
		0, 176, 178, 3, 60, 30, 0, 177, 175, 1, 0, 0, 0, 178, 181, 1, 0, 0, 0,
		179, 177, 1, 0, 0, 0, 179, 180, 1, 0, 0, 0, 180, 183, 1, 0, 0, 0, 181,
		179, 1, 0, 0, 0, 182, 174, 1, 0, 0, 0, 182, 183, 1, 0, 0, 0, 183, 184,
		1, 0, 0, 0, 184, 189, 5, 58, 0, 0, 185, 189, 3, 60, 30, 0, 186, 189, 3,
		18, 9, 0, 187, 189, 3, 20, 10, 0, 188, 172, 1, 0, 0, 0, 188, 185, 1, 0,
		0, 0, 188, 186, 1, 0, 0, 0, 188, 187, 1, 0, 0, 0, 189, 17, 1, 0, 0, 0,
		190, 195, 3, 22, 11, 0, 191, 195, 3, 28, 14, 0, 192, 195, 3, 32, 16, 0,
		193, 195, 3, 58, 29, 0, 194, 190, 1, 0, 0, 0, 194, 191, 1, 0, 0, 0, 194,
		192, 1, 0, 0, 0, 194, 193, 1, 0, 0, 0, 195, 19, 1, 0, 0, 0, 196, 203, 5,
		23, 0, 0, 197, 203, 5, 24, 0, 0, 198, 200, 5, 25, 0, 0, 199, 201, 3, 60,
		30, 0, 200, 199, 1, 0, 0, 0, 200, 201, 1, 0, 0, 0, 201, 203, 1, 0, 0, 0,
		202, 196, 1, 0, 0, 0, 202, 197, 1, 0, 0, 0, 202, 198, 1, 0, 0, 0, 203,
		21, 1, 0, 0, 0, 204, 205, 5, 17, 0, 0, 205, 206, 3, 60, 30, 0, 206, 210,
		5, 61, 0, 0, 207, 209, 3, 8, 4, 0, 208, 207, 1, 0, 0, 0, 209, 212, 1, 0,
		0, 0, 210, 208, 1, 0, 0, 0, 210, 211, 1, 0, 0, 0, 211, 213, 1, 0, 0, 0,
		212, 210, 1, 0, 0, 0, 213, 217, 5, 62, 0, 0, 214, 216, 3, 24, 12, 0, 215,
		214, 1, 0, 0, 0, 216, 219, 1, 0, 0, 0, 217, 215, 1, 0, 0, 0, 217, 218,
		1, 0, 0, 0, 218, 221, 1, 0, 0, 0, 219, 217, 1, 0, 0, 0, 220, 222, 3, 26,
		13, 0, 221, 220, 1, 0, 0, 0, 221, 222, 1, 0, 0, 0, 222, 23, 1, 0, 0, 0,
		223, 224, 5, 18, 0, 0, 224, 225, 5, 17, 0, 0, 225, 226, 3, 60, 30, 0, 226,
		230, 5, 61, 0, 0, 227, 229, 3, 8, 4, 0, 228, 227, 1, 0, 0, 0, 229, 232,
		1, 0, 0, 0, 230, 228, 1, 0, 0, 0, 230, 231, 1, 0, 0, 0, 231, 233, 1, 0,
		0, 0, 232, 230, 1, 0, 0, 0, 233, 234, 5, 62, 0, 0, 234, 25, 1, 0, 0, 0,
		235, 236, 5, 18, 0, 0, 236, 240, 5, 61, 0, 0, 237, 239, 3, 8, 4, 0, 238,
		237, 1, 0, 0, 0, 239, 242, 1, 0, 0, 0, 240, 238, 1, 0, 0, 0, 240, 241,
		1, 0, 0, 0, 241, 243, 1, 0, 0, 0, 242, 240, 1, 0, 0, 0, 243, 244, 5, 62,
		0, 0, 244, 27, 1, 0, 0, 0, 245, 246, 5, 19, 0, 0, 246, 247, 3, 30, 15,
		0, 247, 248, 5, 63, 0, 0, 248, 249, 3, 60, 30, 0, 249, 251, 5, 63, 0, 0,
		250, 252, 3, 16, 8, 0, 251, 250, 1, 0, 0, 0, 251, 252, 1, 0, 0, 0, 252,
		253, 1, 0, 0, 0, 253, 254, 3, 6, 3, 0, 254, 268, 1, 0, 0, 0, 255, 256,
		5, 19, 0, 0, 256, 257, 3, 60, 30, 0, 257, 258, 3, 6, 3, 0, 258, 268, 1,
		0, 0, 0, 259, 260, 5, 19, 0, 0, 260, 261, 5, 37, 0, 0, 261, 262, 5, 66,
		0, 0, 262, 263, 5, 37, 0, 0, 263, 264, 5, 56, 0, 0, 264, 265, 5, 4, 0,
		0, 265, 266, 5, 37, 0, 0, 266, 268, 3, 6, 3, 0, 267, 245, 1, 0, 0, 0, 267,
		255, 1, 0, 0, 0, 267, 259, 1, 0, 0, 0, 268, 29, 1, 0, 0, 0, 269, 270, 5,
		37, 0, 0, 270, 271, 5, 56, 0, 0, 271, 272, 3, 60, 30, 0, 272, 31, 1, 0,
		0, 0, 273, 274, 5, 20, 0, 0, 274, 275, 3, 60, 30, 0, 275, 279, 5, 61, 0,
		0, 276, 278, 3, 34, 17, 0, 277, 276, 1, 0, 0, 0, 278, 281, 1, 0, 0, 0,
		279, 277, 1, 0, 0, 0, 279, 280, 1, 0, 0, 0, 280, 283, 1, 0, 0, 0, 281,
		279, 1, 0, 0, 0, 282, 284, 3, 36, 18, 0, 283, 282, 1, 0, 0, 0, 283, 284,
		1, 0, 0, 0, 284, 285, 1, 0, 0, 0, 285, 286, 5, 62, 0, 0, 286, 33, 1, 0,
		0, 0, 287, 288, 5, 5, 0, 0, 288, 289, 3, 60, 30, 0, 289, 293, 5, 64, 0,
		0, 290, 292, 3, 8, 4, 0, 291, 290, 1, 0, 0, 0, 292, 295, 1, 0, 0, 0, 293,
		291, 1, 0, 0, 0, 293, 294, 1, 0, 0, 0, 294, 35, 1, 0, 0, 0, 295, 293, 1,
		0, 0, 0, 296, 297, 5, 6, 0, 0, 297, 301, 5, 64, 0, 0, 298, 300, 3, 8, 4,
		0, 299, 298, 1, 0, 0, 0, 300, 303, 1, 0, 0, 0, 301, 299, 1, 0, 0, 0, 301,
		302, 1, 0, 0, 0, 302, 37, 1, 0, 0, 0, 303, 301, 1, 0, 0, 0, 304, 305, 5,
		21, 0, 0, 305, 314, 5, 57, 0, 0, 306, 311, 3, 60, 30, 0, 307, 308, 5, 66,
		0, 0, 308, 310, 3, 60, 30, 0, 309, 307, 1, 0, 0, 0, 310, 313, 1, 0, 0,
		0, 311, 309, 1, 0, 0, 0, 311, 312, 1, 0, 0, 0, 312, 315, 1, 0, 0, 0, 313,
		311, 1, 0, 0, 0, 314, 306, 1, 0, 0, 0, 314, 315, 1, 0, 0, 0, 315, 316,
		1, 0, 0, 0, 316, 370, 5, 58, 0, 0, 317, 318, 5, 22, 0, 0, 318, 327, 5,
		57, 0, 0, 319, 324, 3, 60, 30, 0, 320, 321, 5, 66, 0, 0, 321, 323, 3, 60,
		30, 0, 322, 320, 1, 0, 0, 0, 323, 326, 1, 0, 0, 0, 324, 322, 1, 0, 0, 0,
		324, 325, 1, 0, 0, 0, 325, 328, 1, 0, 0, 0, 326, 324, 1, 0, 0, 0, 327,
		319, 1, 0, 0, 0, 327, 328, 1, 0, 0, 0, 328, 329, 1, 0, 0, 0, 329, 370,
		5, 58, 0, 0, 330, 331, 5, 37, 0, 0, 331, 340, 5, 57, 0, 0, 332, 337, 3,
		60, 30, 0, 333, 334, 5, 66, 0, 0, 334, 336, 3, 60, 30, 0, 335, 333, 1,
		0, 0, 0, 336, 339, 1, 0, 0, 0, 337, 335, 1, 0, 0, 0, 337, 338, 1, 0, 0,
		0, 338, 341, 1, 0, 0, 0, 339, 337, 1, 0, 0, 0, 340, 332, 1, 0, 0, 0, 340,
		341, 1, 0, 0, 0, 341, 342, 1, 0, 0, 0, 342, 370, 5, 58, 0, 0, 343, 344,
		5, 14, 0, 0, 344, 353, 5, 57, 0, 0, 345, 350, 3, 60, 30, 0, 346, 347, 5,
		66, 0, 0, 347, 349, 3, 60, 30, 0, 348, 346, 1, 0, 0, 0, 349, 352, 1, 0,
		0, 0, 350, 348, 1, 0, 0, 0, 350, 351, 1, 0, 0, 0, 351, 354, 1, 0, 0, 0,
		352, 350, 1, 0, 0, 0, 353, 345, 1, 0, 0, 0, 353, 354, 1, 0, 0, 0, 354,
		355, 1, 0, 0, 0, 355, 370, 5, 58, 0, 0, 356, 357, 5, 16, 0, 0, 357, 366,
		5, 57, 0, 0, 358, 363, 3, 60, 30, 0, 359, 360, 5, 66, 0, 0, 360, 362, 3,
		60, 30, 0, 361, 359, 1, 0, 0, 0, 362, 365, 1, 0, 0, 0, 363, 361, 1, 0,
		0, 0, 363, 364, 1, 0, 0, 0, 364, 367, 1, 0, 0, 0, 365, 363, 1, 0, 0, 0,
		366, 358, 1, 0, 0, 0, 366, 367, 1, 0, 0, 0, 367, 368, 1, 0, 0, 0, 368,
		370, 5, 58, 0, 0, 369, 304, 1, 0, 0, 0, 369, 317, 1, 0, 0, 0, 369, 330,
		1, 0, 0, 0, 369, 343, 1, 0, 0, 0, 369, 356, 1, 0, 0, 0, 370, 39, 1, 0,
		0, 0, 371, 372, 5, 37, 0, 0, 372, 374, 5, 57, 0, 0, 373, 375, 3, 46, 23,
		0, 374, 373, 1, 0, 0, 0, 374, 375, 1, 0, 0, 0, 375, 376, 1, 0, 0, 0, 376,
		377, 5, 58, 0, 0, 377, 41, 1, 0, 0, 0, 378, 383, 3, 44, 22, 0, 379, 380,
		5, 66, 0, 0, 380, 382, 3, 44, 22, 0, 381, 379, 1, 0, 0, 0, 382, 385, 1,
		0, 0, 0, 383, 381, 1, 0, 0, 0, 383, 384, 1, 0, 0, 0, 384, 43, 1, 0, 0,
		0, 385, 383, 1, 0, 0, 0, 386, 387, 5, 37, 0, 0, 387, 388, 5, 9, 0, 0, 388,
		45, 1, 0, 0, 0, 389, 394, 3, 60, 30, 0, 390, 391, 5, 66, 0, 0, 391, 393,
		3, 60, 30, 0, 392, 390, 1, 0, 0, 0, 393, 396, 1, 0, 0, 0, 394, 392, 1,
		0, 0, 0, 394, 395, 1, 0, 0, 0, 395, 47, 1, 0, 0, 0, 396, 394, 1, 0, 0,
		0, 397, 398, 5, 7, 0, 0, 398, 399, 5, 37, 0, 0, 399, 400, 5, 61, 0, 0,
		400, 401, 3, 50, 25, 0, 401, 402, 5, 62, 0, 0, 402, 49, 1, 0, 0, 0, 403,
		405, 3, 52, 26, 0, 404, 403, 1, 0, 0, 0, 405, 406, 1, 0, 0, 0, 406, 404,
		1, 0, 0, 0, 406, 407, 1, 0, 0, 0, 407, 51, 1, 0, 0, 0, 408, 409, 5, 37,
		0, 0, 409, 413, 5, 9, 0, 0, 410, 411, 5, 37, 0, 0, 411, 413, 5, 37, 0,
		0, 412, 408, 1, 0, 0, 0, 412, 410, 1, 0, 0, 0, 413, 53, 1, 0, 0, 0, 414,
		419, 3, 56, 28, 0, 415, 416, 5, 66, 0, 0, 416, 418, 3, 56, 28, 0, 417,
		415, 1, 0, 0, 0, 418, 421, 1, 0, 0, 0, 419, 417, 1, 0, 0, 0, 419, 420,
		1, 0, 0, 0, 420, 55, 1, 0, 0, 0, 421, 419, 1, 0, 0, 0, 422, 423, 5, 37,
		0, 0, 423, 424, 5, 64, 0, 0, 424, 425, 3, 60, 30, 0, 425, 57, 1, 0, 0,
		0, 426, 427, 5, 8, 0, 0, 427, 428, 5, 57, 0, 0, 428, 429, 3, 60, 30, 0,
		429, 430, 5, 58, 0, 0, 430, 434, 5, 59, 0, 0, 431, 433, 3, 8, 4, 0, 432,
		431, 1, 0, 0, 0, 433, 436, 1, 0, 0, 0, 434, 432, 1, 0, 0, 0, 434, 435,
		1, 0, 0, 0, 435, 437, 1, 0, 0, 0, 436, 434, 1, 0, 0, 0, 437, 438, 5, 60,
		0, 0, 438, 59, 1, 0, 0, 0, 439, 440, 6, 30, -1, 0, 440, 441, 7, 0, 0, 0,
		441, 502, 3, 60, 30, 18, 442, 502, 3, 66, 33, 0, 443, 444, 5, 57, 0, 0,
		444, 445, 3, 60, 30, 0, 445, 446, 5, 58, 0, 0, 446, 502, 1, 0, 0, 0, 447,
		448, 5, 59, 0, 0, 448, 449, 3, 60, 30, 0, 449, 450, 5, 60, 0, 0, 450, 502,
		1, 0, 0, 0, 451, 452, 5, 37, 0, 0, 452, 453, 5, 59, 0, 0, 453, 454, 3,
		60, 30, 0, 454, 455, 5, 60, 0, 0, 455, 502, 1, 0, 0, 0, 456, 502, 3, 38,
		19, 0, 457, 458, 5, 37, 0, 0, 458, 459, 5, 65, 0, 0, 459, 460, 5, 37, 0,
		0, 460, 461, 5, 56, 0, 0, 461, 502, 3, 60, 30, 10, 462, 463, 5, 37, 0,
		0, 463, 464, 5, 61, 0, 0, 464, 465, 3, 54, 27, 0, 465, 466, 5, 62, 0, 0,
		466, 502, 1, 0, 0, 0, 467, 502, 5, 37, 0, 0, 468, 502, 3, 70, 35, 0, 469,
		470, 5, 37, 0, 0, 470, 471, 5, 65, 0, 0, 471, 502, 5, 37, 0, 0, 472, 473,
		5, 37, 0, 0, 473, 474, 5, 65, 0, 0, 474, 502, 3, 60, 30, 5, 475, 476, 5,
		37, 0, 0, 476, 477, 5, 9, 0, 0, 477, 478, 5, 56, 0, 0, 478, 502, 3, 60,
		30, 4, 479, 480, 5, 37, 0, 0, 480, 481, 7, 1, 0, 0, 481, 502, 3, 60, 30,
		3, 482, 483, 5, 13, 0, 0, 483, 484, 5, 57, 0, 0, 484, 485, 5, 59, 0, 0,
		485, 490, 3, 60, 30, 0, 486, 487, 5, 66, 0, 0, 487, 489, 3, 60, 30, 0,
		488, 486, 1, 0, 0, 0, 489, 492, 1, 0, 0, 0, 490, 488, 1, 0, 0, 0, 490,
		491, 1, 0, 0, 0, 491, 493, 1, 0, 0, 0, 492, 490, 1, 0, 0, 0, 493, 494,
		5, 60, 0, 0, 494, 495, 5, 58, 0, 0, 495, 502, 1, 0, 0, 0, 496, 497, 5,
		13, 0, 0, 497, 498, 5, 57, 0, 0, 498, 499, 3, 60, 30, 0, 499, 500, 5, 58,
		0, 0, 500, 502, 1, 0, 0, 0, 501, 439, 1, 0, 0, 0, 501, 442, 1, 0, 0, 0,
		501, 443, 1, 0, 0, 0, 501, 447, 1, 0, 0, 0, 501, 451, 1, 0, 0, 0, 501,
		456, 1, 0, 0, 0, 501, 457, 1, 0, 0, 0, 501, 462, 1, 0, 0, 0, 501, 467,
		1, 0, 0, 0, 501, 468, 1, 0, 0, 0, 501, 469, 1, 0, 0, 0, 501, 472, 1, 0,
		0, 0, 501, 475, 1, 0, 0, 0, 501, 479, 1, 0, 0, 0, 501, 482, 1, 0, 0, 0,
		501, 496, 1, 0, 0, 0, 502, 520, 1, 0, 0, 0, 503, 504, 10, 21, 0, 0, 504,
		505, 7, 2, 0, 0, 505, 519, 3, 60, 30, 22, 506, 507, 10, 20, 0, 0, 507,
		508, 7, 3, 0, 0, 508, 519, 3, 60, 30, 21, 509, 510, 10, 19, 0, 0, 510,
		511, 7, 4, 0, 0, 511, 519, 3, 60, 30, 20, 512, 513, 10, 17, 0, 0, 513,
		514, 7, 5, 0, 0, 514, 519, 3, 60, 30, 18, 515, 516, 10, 16, 0, 0, 516,
		517, 7, 6, 0, 0, 517, 519, 3, 60, 30, 17, 518, 503, 1, 0, 0, 0, 518, 506,
		1, 0, 0, 0, 518, 509, 1, 0, 0, 0, 518, 512, 1, 0, 0, 0, 518, 515, 1, 0,
		0, 0, 519, 522, 1, 0, 0, 0, 520, 518, 1, 0, 0, 0, 520, 521, 1, 0, 0, 0,
		521, 61, 1, 0, 0, 0, 522, 520, 1, 0, 0, 0, 523, 528, 3, 60, 30, 0, 524,
		525, 5, 66, 0, 0, 525, 527, 3, 60, 30, 0, 526, 524, 1, 0, 0, 0, 527, 530,
		1, 0, 0, 0, 528, 526, 1, 0, 0, 0, 528, 529, 1, 0, 0, 0, 529, 63, 1, 0,
		0, 0, 530, 528, 1, 0, 0, 0, 531, 532, 3, 66, 33, 0, 532, 65, 1, 0, 0, 0,
		533, 539, 5, 27, 0, 0, 534, 539, 5, 28, 0, 0, 535, 539, 5, 29, 0, 0, 536,
		539, 5, 26, 0, 0, 537, 539, 5, 30, 0, 0, 538, 533, 1, 0, 0, 0, 538, 534,
		1, 0, 0, 0, 538, 535, 1, 0, 0, 0, 538, 536, 1, 0, 0, 0, 538, 537, 1, 0,
		0, 0, 539, 67, 1, 0, 0, 0, 540, 545, 3, 60, 30, 0, 541, 542, 5, 66, 0,
		0, 542, 544, 3, 60, 30, 0, 543, 541, 1, 0, 0, 0, 544, 547, 1, 0, 0, 0,
		545, 543, 1, 0, 0, 0, 545, 546, 1, 0, 0, 0, 546, 69, 1, 0, 0, 0, 547, 545,
		1, 0, 0, 0, 548, 549, 5, 37, 0, 0, 549, 553, 5, 38, 0, 0, 550, 551, 5,
		37, 0, 0, 551, 553, 5, 39, 0, 0, 552, 548, 1, 0, 0, 0, 552, 550, 1, 0,
		0, 0, 553, 71, 1, 0, 0, 0, 53, 75, 90, 94, 102, 112, 117, 121, 144, 160,
		168, 179, 182, 188, 194, 200, 202, 210, 217, 221, 230, 240, 251, 267, 279,
		283, 293, 301, 311, 314, 324, 327, 337, 340, 350, 353, 363, 366, 369, 374,
		383, 394, 406, 412, 419, 434, 490, 501, 518, 520, 528, 538, 545, 552,
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
	VlangParserEOF            = antlr.TokenEOF
	VlangParserT__0           = 1
	VlangParserT__1           = 2
	VlangParserT__2           = 3
	VlangParserT__3           = 4
	VlangParserT__4           = 5
	VlangParserT__5           = 6
	VlangParserT__6           = 7
	VlangParserT__7           = 8
	VlangParserTIPO           = 9
	VlangParserCASTEOS        = 10
	VlangParserATOI           = 11
	VlangParserPARSEFLOAT     = 12
	VlangParserTYPEOF         = 13
	VlangParserLEN            = 14
	VlangParserCAP            = 15
	VlangParserAPPEND         = 16
	VlangParserIF             = 17
	VlangParserELSE           = 18
	VlangParserFOR            = 19
	VlangParserSWITCH         = 20
	VlangParserINDEXOF        = 21
	VlangParserJOIN           = 22
	VlangParserBREAK          = 23
	VlangParserCONTINUE       = 24
	VlangParserRETURN         = 25
	VlangParserBOOLEANO       = 26
	VlangParserENTERO         = 27
	VlangParserDECIMAL        = 28
	VlangParserCADENA         = 29
	VlangParserCARACTER       = 30
	VlangParserTIPO_ENTERO    = 31
	VlangParserTIPO_DECIMAL   = 32
	VlangParserTIPO_CADENA    = 33
	VlangParserTIPO_BOOLEANO  = 34
	VlangParserTIPO_CHAR      = 35
	VlangParserPRINT          = 36
	VlangParserID             = 37
	VlangParserINC            = 38
	VlangParserDEC            = 39
	VlangParserSUMAIMPLICITA  = 40
	VlangParserRESTOIMPLICITO = 41
	VlangParserPLUS           = 42
	VlangParserMINUS          = 43
	VlangParserMUL            = 44
	VlangParserDIV            = 45
	VlangParserMOD            = 46
	VlangParserNOT            = 47
	VlangParserOR             = 48
	VlangParserAND            = 49
	VlangParserEQ             = 50
	VlangParserNEQ            = 51
	VlangParserLE             = 52
	VlangParserGE             = 53
	VlangParserLT             = 54
	VlangParserGT             = 55
	VlangParserASSIGN         = 56
	VlangParserLPAREN         = 57
	VlangParserRPAREN         = 58
	VlangParserLBRACK         = 59
	VlangParserRBRACK         = 60
	VlangParserLBRACE         = 61
	VlangParserRBRACE         = 62
	VlangParserSEMICOLON      = 63
	VlangParserCOLON          = 64
	VlangParserDOT            = 65
	VlangParserCOMMA          = 66
	VlangParserWS             = 67
	VlangParserLINE_COMMENT   = 68
	VlangParserBLOCK_COMMENT  = 69
)

// VlangParser rules.
const (
	VlangParserRULE_programa                 = 0
	VlangParserRULE_funcMain                 = 1
	VlangParserRULE_funcDcl                  = 2
	VlangParserRULE_block                    = 3
	VlangParserRULE_declaraciones            = 4
	VlangParserRULE_varDcl                   = 5
	VlangParserRULE_sliceTipo                = 6
	VlangParserRULE_sliceInit                = 7
	VlangParserRULE_stmt                     = 8
	VlangParserRULE_sentencias_control       = 9
	VlangParserRULE_sentencias_transferencia = 10
	VlangParserRULE_ifDcl                    = 11
	VlangParserRULE_elseIfDcl                = 12
	VlangParserRULE_elseCondicional          = 13
	VlangParserRULE_forDcl                   = 14
	VlangParserRULE_asignacion               = 15
	VlangParserRULE_switchDcl                = 16
	VlangParserRULE_caseBlock                = 17
	VlangParserRULE_defaultBlock             = 18
	VlangParserRULE_llamadaFuncion           = 19
	VlangParserRULE_funcCall                 = 20
	VlangParserRULE_parametrosFormales       = 21
	VlangParserRULE_parametro                = 22
	VlangParserRULE_parametrosReales         = 23
	VlangParserRULE_structDcl                = 24
	VlangParserRULE_atributosStruct          = 25
	VlangParserRULE_atributoStruct           = 26
	VlangParserRULE_listaAsignaciones        = 27
	VlangParserRULE_asignacionStruct         = 28
	VlangParserRULE_whileDcl                 = 29
	VlangParserRULE_expresion                = 30
	VlangParserRULE_parametros               = 31
	VlangParserRULE_valores                  = 32
	VlangParserRULE_valor                    = 33
	VlangParserRULE_listaExpresiones         = 34
	VlangParserRULE_incredecre               = 35
)

// IProgramaContext is an interface to support dynamic dispatch.
type IProgramaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllDeclaraciones() []IDeclaracionesContext
	Declaraciones(i int) IDeclaracionesContext

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

func (s *ProgramaContext) AllDeclaraciones() []IDeclaracionesContext {
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

func (s *ProgramaContext) Declaraciones(i int) IDeclaracionesContext {
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
	p.SetState(75)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&720725682266268042) != 0 {
		{
			p.SetState(72)
			p.Declaraciones()
		}

		p.SetState(77)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(78)
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
		p.SetState(80)
		p.Match(VlangParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(81)
		p.Match(VlangParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(82)
		p.Match(VlangParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(83)
		p.Match(VlangParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(84)
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

// IFuncDclContext is an interface to support dynamic dispatch.
type IFuncDclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	Block() IBlockContext
	ParametrosFormales() IParametrosFormalesContext
	TIPO() antlr.TerminalNode

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
	p.RuleIndex = VlangParserRULE_funcDcl
	return p
}

func InitEmptyFuncDclContext(p *FuncDclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_funcDcl
}

func (*FuncDclContext) IsFuncDclContext() {}

func NewFuncDclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncDclContext {
	var p = new(FuncDclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_funcDcl

	return p
}

func (s *FuncDclContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncDclContext) ID() antlr.TerminalNode {
	return s.GetToken(VlangParserID, 0)
}

func (s *FuncDclContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(VlangParserLPAREN, 0)
}

func (s *FuncDclContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(VlangParserRPAREN, 0)
}

func (s *FuncDclContext) Block() IBlockContext {
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

func (s *FuncDclContext) ParametrosFormales() IParametrosFormalesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParametrosFormalesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParametrosFormalesContext)
}

func (s *FuncDclContext) TIPO() antlr.TerminalNode {
	return s.GetToken(VlangParserTIPO, 0)
}

func (s *FuncDclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncDclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncDclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterFuncDcl(s)
	}
}

func (s *FuncDclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitFuncDcl(s)
	}
}

func (s *FuncDclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitFuncDcl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) FuncDcl() (localctx IFuncDclContext) {
	localctx = NewFuncDclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, VlangParserRULE_funcDcl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(86)
		p.Match(VlangParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(87)
		p.Match(VlangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(88)
		p.Match(VlangParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(90)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == VlangParserID {
		{
			p.SetState(89)
			p.ParametrosFormales()
		}

	}
	{
		p.SetState(92)
		p.Match(VlangParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(94)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == VlangParserTIPO {
		{
			p.SetState(93)
			p.Match(VlangParserTIPO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(96)
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
	p.EnterRule(localctx, 6, VlangParserRULE_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(98)
		p.Match(VlangParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(102)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&720725682266268042) != 0 {
		{
			p.SetState(99)
			p.Declaraciones()
		}

		p.SetState(104)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(105)
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
	FuncDcl() IFuncDclContext
	FuncMain() IFuncMainContext
	StructDcl() IStructDclContext

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

func (s *DeclaracionesContext) FuncDcl() IFuncDclContext {
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

func (s *DeclaracionesContext) FuncMain() IFuncMainContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncMainContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncMainContext)
}

func (s *DeclaracionesContext) StructDcl() IStructDclContext {
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
	p.EnterRule(localctx, 8, VlangParserRULE_declaraciones)
	p.SetState(112)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(107)
			p.VarDcl()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(108)
			p.Stmt()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(109)
			p.FuncDcl()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(110)
			p.FuncMain()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(111)
			p.StructDcl()
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

type StructDirectInitDeclarationContext struct {
	VarDclContext
}

func NewStructDirectInitDeclarationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructDirectInitDeclarationContext {
	var p = new(StructDirectInitDeclarationContext)

	InitEmptyVarDclContext(&p.VarDclContext)
	p.parser = parser
	p.CopyAll(ctx.(*VarDclContext))

	return p
}

func (s *StructDirectInitDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructDirectInitDeclarationContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(VlangParserID)
}

func (s *StructDirectInitDeclarationContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(VlangParserID, i)
}

func (s *StructDirectInitDeclarationContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(VlangParserASSIGN, 0)
}

func (s *StructDirectInitDeclarationContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(VlangParserLBRACE, 0)
}

func (s *StructDirectInitDeclarationContext) ListaAsignaciones() IListaAsignacionesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListaAsignacionesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListaAsignacionesContext)
}

func (s *StructDirectInitDeclarationContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(VlangParserRBRACE, 0)
}

func (s *StructDirectInitDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterStructDirectInitDeclaration(s)
	}
}

func (s *StructDirectInitDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitStructDirectInitDeclaration(s)
	}
}

func (s *StructDirectInitDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitStructDirectInitDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

type SliceEmptyDeclarationContext struct {
	VarDclContext
}

func NewSliceEmptyDeclarationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SliceEmptyDeclarationContext {
	var p = new(SliceEmptyDeclarationContext)

	InitEmptyVarDclContext(&p.VarDclContext)
	p.parser = parser
	p.CopyAll(ctx.(*VarDclContext))

	return p
}

func (s *SliceEmptyDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SliceEmptyDeclarationContext) ID() antlr.TerminalNode {
	return s.GetToken(VlangParserID, 0)
}

func (s *SliceEmptyDeclarationContext) SliceTipo() ISliceTipoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISliceTipoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISliceTipoContext)
}

func (s *SliceEmptyDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterSliceEmptyDeclaration(s)
	}
}

func (s *SliceEmptyDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitSliceEmptyDeclaration(s)
	}
}

func (s *SliceEmptyDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitSliceEmptyDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

type SliceAssignmentContext struct {
	VarDclContext
}

func NewSliceAssignmentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SliceAssignmentContext {
	var p = new(SliceAssignmentContext)

	InitEmptyVarDclContext(&p.VarDclContext)
	p.parser = parser
	p.CopyAll(ctx.(*VarDclContext))

	return p
}

func (s *SliceAssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SliceAssignmentContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(VlangParserID)
}

func (s *SliceAssignmentContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(VlangParserID, i)
}

func (s *SliceAssignmentContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(VlangParserASSIGN, 0)
}

func (s *SliceAssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterSliceAssignment(s)
	}
}

func (s *SliceAssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitSliceAssignment(s)
	}
}

func (s *SliceAssignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitSliceAssignment(s)

	default:
		return t.VisitChildren(s)
	}
}

type VariableDeclarationImmutableContext struct {
	VarDclContext
}

func NewVariableDeclarationImmutableContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VariableDeclarationImmutableContext {
	var p = new(VariableDeclarationImmutableContext)

	InitEmptyVarDclContext(&p.VarDclContext)
	p.parser = parser
	p.CopyAll(ctx.(*VarDclContext))

	return p
}

func (s *VariableDeclarationImmutableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDeclarationImmutableContext) ID() antlr.TerminalNode {
	return s.GetToken(VlangParserID, 0)
}

func (s *VariableDeclarationImmutableContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(VlangParserASSIGN, 0)
}

func (s *VariableDeclarationImmutableContext) Expresion() IExpresionContext {
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

func (s *VariableDeclarationImmutableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterVariableDeclarationImmutable(s)
	}
}

func (s *VariableDeclarationImmutableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitVariableDeclarationImmutable(s)
	}
}

func (s *VariableDeclarationImmutableContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitVariableDeclarationImmutable(s)

	default:
		return t.VisitChildren(s)
	}
}

type VariableCastDeclarationContext struct {
	VarDclContext
}

func NewVariableCastDeclarationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VariableCastDeclarationContext {
	var p = new(VariableCastDeclarationContext)

	InitEmptyVarDclContext(&p.VarDclContext)
	p.parser = parser
	p.CopyAll(ctx.(*VarDclContext))

	return p
}

func (s *VariableCastDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableCastDeclarationContext) ID() antlr.TerminalNode {
	return s.GetToken(VlangParserID, 0)
}

func (s *VariableCastDeclarationContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(VlangParserASSIGN, 0)
}

func (s *VariableCastDeclarationContext) CASTEOS() antlr.TerminalNode {
	return s.GetToken(VlangParserCASTEOS, 0)
}

func (s *VariableCastDeclarationContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(VlangParserLPAREN, 0)
}

func (s *VariableCastDeclarationContext) Expresion() IExpresionContext {
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

func (s *VariableCastDeclarationContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(VlangParserRPAREN, 0)
}

func (s *VariableCastDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterVariableCastDeclaration(s)
	}
}

func (s *VariableCastDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitVariableCastDeclaration(s)
	}
}

func (s *VariableCastDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitVariableCastDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

type SliceInitDeclarationContext struct {
	VarDclContext
}

func NewSliceInitDeclarationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SliceInitDeclarationContext {
	var p = new(SliceInitDeclarationContext)

	InitEmptyVarDclContext(&p.VarDclContext)
	p.parser = parser
	p.CopyAll(ctx.(*VarDclContext))

	return p
}

func (s *SliceInitDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SliceInitDeclarationContext) ID() antlr.TerminalNode {
	return s.GetToken(VlangParserID, 0)
}

func (s *SliceInitDeclarationContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(VlangParserASSIGN, 0)
}

func (s *SliceInitDeclarationContext) SliceTipo() ISliceTipoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISliceTipoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISliceTipoContext)
}

func (s *SliceInitDeclarationContext) SliceInit() ISliceInitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISliceInitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISliceInitContext)
}

func (s *SliceInitDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterSliceInitDeclaration(s)
	}
}

func (s *SliceInitDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitSliceInitDeclaration(s)
	}
}

func (s *SliceInitDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitSliceInitDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
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

type SliceAssignmentIndexContext struct {
	VarDclContext
}

func NewSliceAssignmentIndexContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SliceAssignmentIndexContext {
	var p = new(SliceAssignmentIndexContext)

	InitEmptyVarDclContext(&p.VarDclContext)
	p.parser = parser
	p.CopyAll(ctx.(*VarDclContext))

	return p
}

func (s *SliceAssignmentIndexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SliceAssignmentIndexContext) ID() antlr.TerminalNode {
	return s.GetToken(VlangParserID, 0)
}

func (s *SliceAssignmentIndexContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(VlangParserLBRACK, 0)
}

func (s *SliceAssignmentIndexContext) AllExpresion() []IExpresionContext {
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

func (s *SliceAssignmentIndexContext) Expresion(i int) IExpresionContext {
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

func (s *SliceAssignmentIndexContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(VlangParserRBRACK, 0)
}

func (s *SliceAssignmentIndexContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(VlangParserASSIGN, 0)
}

func (s *SliceAssignmentIndexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterSliceAssignmentIndex(s)
	}
}

func (s *SliceAssignmentIndexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitSliceAssignmentIndex(s)
	}
}

func (s *SliceAssignmentIndexContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitSliceAssignmentIndex(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) VarDcl() (localctx IVarDclContext) {
	localctx = NewVarDclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, VlangParserRULE_varDcl)
	var _la int

	p.SetState(160)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		localctx = NewVariableDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(114)
			p.Match(VlangParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(115)
			p.Match(VlangParserID)
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

		if _la == VlangParserTIPO {
			{
				p.SetState(116)
				p.Match(VlangParserTIPO)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(121)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == VlangParserASSIGN {
			{
				p.SetState(119)
				p.Match(VlangParserASSIGN)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(120)
				p.expresion(0)
			}

		}

	case 2:
		localctx = NewSliceEmptyDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(123)
			p.Match(VlangParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(124)
			p.Match(VlangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(125)
			p.SliceTipo()
		}

	case 3:
		localctx = NewStructDirectInitDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(126)
			p.Match(VlangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(127)
			p.Match(VlangParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(128)
			p.Match(VlangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(129)
			p.Match(VlangParserLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(130)
			p.ListaAsignaciones()
		}
		{
			p.SetState(131)
			p.Match(VlangParserRBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewSliceInitDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(133)
			p.Match(VlangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(134)
			p.Match(VlangParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(135)
			p.SliceTipo()
		}
		{
			p.SetState(136)
			p.SliceInit()
		}

	case 5:
		localctx = NewSliceAssignmentContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(138)
			p.Match(VlangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(139)
			p.Match(VlangParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(140)
			p.Match(VlangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		localctx = NewVariableDeclarationImmutableContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(141)
			p.Match(VlangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(144)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == VlangParserASSIGN {
			{
				p.SetState(142)
				p.Match(VlangParserASSIGN)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(143)
				p.expresion(0)
			}

		}

	case 7:
		localctx = NewVariableCastDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(146)
			p.Match(VlangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(147)
			p.Match(VlangParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(148)
			p.Match(VlangParserCASTEOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(149)
			p.Match(VlangParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(150)
			p.expresion(0)
		}
		{
			p.SetState(151)
			p.Match(VlangParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 8:
		localctx = NewSliceAssignmentIndexContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(153)
			p.Match(VlangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(154)
			p.Match(VlangParserLBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(155)
			p.expresion(0)
		}
		{
			p.SetState(156)
			p.Match(VlangParserRBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(157)
			p.Match(VlangParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(158)
			p.expresion(0)
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

// ISliceTipoContext is an interface to support dynamic dispatch.
type ISliceTipoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACK() antlr.TerminalNode
	RBRACK() antlr.TerminalNode
	TIPO() antlr.TerminalNode

	// IsSliceTipoContext differentiates from other interfaces.
	IsSliceTipoContext()
}

type SliceTipoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySliceTipoContext() *SliceTipoContext {
	var p = new(SliceTipoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_sliceTipo
	return p
}

func InitEmptySliceTipoContext(p *SliceTipoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_sliceTipo
}

func (*SliceTipoContext) IsSliceTipoContext() {}

func NewSliceTipoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SliceTipoContext {
	var p = new(SliceTipoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_sliceTipo

	return p
}

func (s *SliceTipoContext) GetParser() antlr.Parser { return s.parser }

func (s *SliceTipoContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(VlangParserLBRACK, 0)
}

func (s *SliceTipoContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(VlangParserRBRACK, 0)
}

func (s *SliceTipoContext) TIPO() antlr.TerminalNode {
	return s.GetToken(VlangParserTIPO, 0)
}

func (s *SliceTipoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SliceTipoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SliceTipoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterSliceTipo(s)
	}
}

func (s *SliceTipoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitSliceTipo(s)
	}
}

func (s *SliceTipoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitSliceTipo(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) SliceTipo() (localctx ISliceTipoContext) {
	localctx = NewSliceTipoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, VlangParserRULE_sliceTipo)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(162)
		p.Match(VlangParserLBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(163)
		p.Match(VlangParserRBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(164)
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

// ISliceInitContext is an interface to support dynamic dispatch.
type ISliceInitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	ListaExpresiones() IListaExpresionesContext

	// IsSliceInitContext differentiates from other interfaces.
	IsSliceInitContext()
}

type SliceInitContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySliceInitContext() *SliceInitContext {
	var p = new(SliceInitContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_sliceInit
	return p
}

func InitEmptySliceInitContext(p *SliceInitContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_sliceInit
}

func (*SliceInitContext) IsSliceInitContext() {}

func NewSliceInitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SliceInitContext {
	var p = new(SliceInitContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_sliceInit

	return p
}

func (s *SliceInitContext) GetParser() antlr.Parser { return s.parser }

func (s *SliceInitContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(VlangParserLBRACE, 0)
}

func (s *SliceInitContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(VlangParserRBRACE, 0)
}

func (s *SliceInitContext) ListaExpresiones() IListaExpresionesContext {
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

func (s *SliceInitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SliceInitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SliceInitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterSliceInit(s)
	}
}

func (s *SliceInitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitSliceInit(s)
	}
}

func (s *SliceInitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitSliceInit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) SliceInit() (localctx ISliceInitContext) {
	localctx = NewSliceInitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, VlangParserRULE_sliceInit)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(166)
		p.Match(VlangParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(168)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&720725613486366720) != 0 {
		{
			p.SetState(167)
			p.ListaExpresiones()
		}

	}
	{
		p.SetState(170)
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

type TransfersentenceContext struct {
	StmtContext
}

func NewTransfersentenceContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TransfersentenceContext {
	var p = new(TransfersentenceContext)

	InitEmptyStmtContext(&p.StmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*StmtContext))

	return p
}

func (s *TransfersentenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TransfersentenceContext) Sentencias_transferencia() ISentencias_transferenciaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISentencias_transferenciaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISentencias_transferenciaContext)
}

func (s *TransfersentenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterTransfersentence(s)
	}
}

func (s *TransfersentenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitTransfersentence(s)
	}
}

func (s *TransfersentenceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitTransfersentence(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Stmt() (localctx IStmtContext) {
	localctx = NewStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, VlangParserRULE_stmt)
	var _la int

	p.SetState(188)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case VlangParserPRINT:
		localctx = NewPrintStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(172)
			p.Match(VlangParserPRINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(173)
			p.Match(VlangParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(182)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&720725613486366720) != 0 {
			{
				p.SetState(174)
				p.expresion(0)
			}
			p.SetState(179)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == VlangParserCOMMA {
				{
					p.SetState(175)
					p.Match(VlangParserCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(176)
					p.expresion(0)
				}

				p.SetState(181)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(184)
			p.Match(VlangParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VlangParserTYPEOF, VlangParserLEN, VlangParserAPPEND, VlangParserINDEXOF, VlangParserJOIN, VlangParserBOOLEANO, VlangParserENTERO, VlangParserDECIMAL, VlangParserCADENA, VlangParserCARACTER, VlangParserID, VlangParserMINUS, VlangParserNOT, VlangParserLPAREN, VlangParserLBRACK:
		localctx = NewExpresionStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(185)
			p.expresion(0)
		}

	case VlangParserT__7, VlangParserIF, VlangParserFOR, VlangParserSWITCH:
		localctx = NewControlStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(186)
			p.Sentencias_control()
		}

	case VlangParserBREAK, VlangParserCONTINUE, VlangParserRETURN:
		localctx = NewTransfersentenceContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(187)
			p.Sentencias_transferencia()
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
	p.EnterRule(localctx, 18, VlangParserRULE_sentencias_control)
	p.SetState(194)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case VlangParserIF:
		localctx = NewIf_contextContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(190)
			p.IfDcl()
		}

	case VlangParserFOR:
		localctx = NewFor_contextContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(191)
			p.ForDcl()
		}

	case VlangParserSWITCH:
		localctx = NewSwitch_contextContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(192)
			p.SwitchDcl()
		}

	case VlangParserT__7:
		localctx = NewWhile_contextContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(193)
			p.WhileDcl()
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

// ISentencias_transferenciaContext is an interface to support dynamic dispatch.
type ISentencias_transferenciaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSentencias_transferenciaContext differentiates from other interfaces.
	IsSentencias_transferenciaContext()
}

type Sentencias_transferenciaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySentencias_transferenciaContext() *Sentencias_transferenciaContext {
	var p = new(Sentencias_transferenciaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_sentencias_transferencia
	return p
}

func InitEmptySentencias_transferenciaContext(p *Sentencias_transferenciaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_sentencias_transferencia
}

func (*Sentencias_transferenciaContext) IsSentencias_transferenciaContext() {}

func NewSentencias_transferenciaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Sentencias_transferenciaContext {
	var p = new(Sentencias_transferenciaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_sentencias_transferencia

	return p
}

func (s *Sentencias_transferenciaContext) GetParser() antlr.Parser { return s.parser }

func (s *Sentencias_transferenciaContext) CopyAll(ctx *Sentencias_transferenciaContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Sentencias_transferenciaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Sentencias_transferenciaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type BreakStatementContext struct {
	Sentencias_transferenciaContext
}

func NewBreakStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BreakStatementContext {
	var p = new(BreakStatementContext)

	InitEmptySentencias_transferenciaContext(&p.Sentencias_transferenciaContext)
	p.parser = parser
	p.CopyAll(ctx.(*Sentencias_transferenciaContext))

	return p
}

func (s *BreakStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BreakStatementContext) BREAK() antlr.TerminalNode {
	return s.GetToken(VlangParserBREAK, 0)
}

func (s *BreakStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterBreakStatement(s)
	}
}

func (s *BreakStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitBreakStatement(s)
	}
}

func (s *BreakStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitBreakStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

type ContinueStatementContext struct {
	Sentencias_transferenciaContext
}

func NewContinueStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ContinueStatementContext {
	var p = new(ContinueStatementContext)

	InitEmptySentencias_transferenciaContext(&p.Sentencias_transferenciaContext)
	p.parser = parser
	p.CopyAll(ctx.(*Sentencias_transferenciaContext))

	return p
}

func (s *ContinueStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContinueStatementContext) CONTINUE() antlr.TerminalNode {
	return s.GetToken(VlangParserCONTINUE, 0)
}

func (s *ContinueStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterContinueStatement(s)
	}
}

func (s *ContinueStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitContinueStatement(s)
	}
}

func (s *ContinueStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitContinueStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

type ReturnStatementContext struct {
	Sentencias_transferenciaContext
}

func NewReturnStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ReturnStatementContext {
	var p = new(ReturnStatementContext)

	InitEmptySentencias_transferenciaContext(&p.Sentencias_transferenciaContext)
	p.parser = parser
	p.CopyAll(ctx.(*Sentencias_transferenciaContext))

	return p
}

func (s *ReturnStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStatementContext) RETURN() antlr.TerminalNode {
	return s.GetToken(VlangParserRETURN, 0)
}

func (s *ReturnStatementContext) Expresion() IExpresionContext {
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

func (s *ReturnStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterReturnStatement(s)
	}
}

func (s *ReturnStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitReturnStatement(s)
	}
}

func (s *ReturnStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitReturnStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Sentencias_transferencia() (localctx ISentencias_transferenciaContext) {
	localctx = NewSentencias_transferenciaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, VlangParserRULE_sentencias_transferencia)
	p.SetState(202)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case VlangParserBREAK:
		localctx = NewBreakStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(196)
			p.Match(VlangParserBREAK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VlangParserCONTINUE:
		localctx = NewContinueStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(197)
			p.Match(VlangParserCONTINUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VlangParserRETURN:
		localctx = NewReturnStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(198)
			p.Match(VlangParserRETURN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(200)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(199)
				p.expresion(0)
			}

		} else if p.HasError() { // JIM
			goto errorExit
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

// IIfDclContext is an interface to support dynamic dispatch.
type IIfDclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IF() antlr.TerminalNode
	Expresion() IExpresionContext
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllDeclaraciones() []IDeclaracionesContext
	Declaraciones(i int) IDeclaracionesContext
	AllElseIfDcl() []IElseIfDclContext
	ElseIfDcl(i int) IElseIfDclContext
	ElseCondicional() IElseCondicionalContext

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

func (s *IfDclContext) IF() antlr.TerminalNode {
	return s.GetToken(VlangParserIF, 0)
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

func (s *IfDclContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(VlangParserLBRACE, 0)
}

func (s *IfDclContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(VlangParserRBRACE, 0)
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

func (s *IfDclContext) AllElseIfDcl() []IElseIfDclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IElseIfDclContext); ok {
			len++
		}
	}

	tst := make([]IElseIfDclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IElseIfDclContext); ok {
			tst[i] = t.(IElseIfDclContext)
			i++
		}
	}

	return tst
}

func (s *IfDclContext) ElseIfDcl(i int) IElseIfDclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElseIfDclContext); ok {
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

	return t.(IElseIfDclContext)
}

func (s *IfDclContext) ElseCondicional() IElseCondicionalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElseCondicionalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElseCondicionalContext)
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
	p.EnterRule(localctx, 22, VlangParserRULE_ifDcl)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(204)
		p.Match(VlangParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(205)
		p.expresion(0)
	}
	{
		p.SetState(206)
		p.Match(VlangParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(210)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&720725682266268042) != 0 {
		{
			p.SetState(207)
			p.Declaraciones()
		}

		p.SetState(212)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(213)
		p.Match(VlangParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(217)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(214)
				p.ElseIfDcl()
			}

		}
		p.SetState(219)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(221)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == VlangParserELSE {
		{
			p.SetState(220)
			p.ElseCondicional()
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

// IElseIfDclContext is an interface to support dynamic dispatch.
type IElseIfDclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ELSE() antlr.TerminalNode
	IF() antlr.TerminalNode
	Expresion() IExpresionContext
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllDeclaraciones() []IDeclaracionesContext
	Declaraciones(i int) IDeclaracionesContext

	// IsElseIfDclContext differentiates from other interfaces.
	IsElseIfDclContext()
}

type ElseIfDclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseIfDclContext() *ElseIfDclContext {
	var p = new(ElseIfDclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_elseIfDcl
	return p
}

func InitEmptyElseIfDclContext(p *ElseIfDclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_elseIfDcl
}

func (*ElseIfDclContext) IsElseIfDclContext() {}

func NewElseIfDclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseIfDclContext {
	var p = new(ElseIfDclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_elseIfDcl

	return p
}

func (s *ElseIfDclContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseIfDclContext) ELSE() antlr.TerminalNode {
	return s.GetToken(VlangParserELSE, 0)
}

func (s *ElseIfDclContext) IF() antlr.TerminalNode {
	return s.GetToken(VlangParserIF, 0)
}

func (s *ElseIfDclContext) Expresion() IExpresionContext {
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

func (s *ElseIfDclContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(VlangParserLBRACE, 0)
}

func (s *ElseIfDclContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(VlangParserRBRACE, 0)
}

func (s *ElseIfDclContext) AllDeclaraciones() []IDeclaracionesContext {
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

func (s *ElseIfDclContext) Declaraciones(i int) IDeclaracionesContext {
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

func (s *ElseIfDclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseIfDclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElseIfDclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterElseIfDcl(s)
	}
}

func (s *ElseIfDclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitElseIfDcl(s)
	}
}

func (s *ElseIfDclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitElseIfDcl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) ElseIfDcl() (localctx IElseIfDclContext) {
	localctx = NewElseIfDclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, VlangParserRULE_elseIfDcl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(223)
		p.Match(VlangParserELSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(224)
		p.Match(VlangParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(225)
		p.expresion(0)
	}
	{
		p.SetState(226)
		p.Match(VlangParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(230)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&720725682266268042) != 0 {
		{
			p.SetState(227)
			p.Declaraciones()
		}

		p.SetState(232)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(233)
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

// IElseCondicionalContext is an interface to support dynamic dispatch.
type IElseCondicionalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ELSE() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllDeclaraciones() []IDeclaracionesContext
	Declaraciones(i int) IDeclaracionesContext

	// IsElseCondicionalContext differentiates from other interfaces.
	IsElseCondicionalContext()
}

type ElseCondicionalContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseCondicionalContext() *ElseCondicionalContext {
	var p = new(ElseCondicionalContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_elseCondicional
	return p
}

func InitEmptyElseCondicionalContext(p *ElseCondicionalContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_elseCondicional
}

func (*ElseCondicionalContext) IsElseCondicionalContext() {}

func NewElseCondicionalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseCondicionalContext {
	var p = new(ElseCondicionalContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_elseCondicional

	return p
}

func (s *ElseCondicionalContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseCondicionalContext) ELSE() antlr.TerminalNode {
	return s.GetToken(VlangParserELSE, 0)
}

func (s *ElseCondicionalContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(VlangParserLBRACE, 0)
}

func (s *ElseCondicionalContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(VlangParserRBRACE, 0)
}

func (s *ElseCondicionalContext) AllDeclaraciones() []IDeclaracionesContext {
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

func (s *ElseCondicionalContext) Declaraciones(i int) IDeclaracionesContext {
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

func (s *ElseCondicionalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseCondicionalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElseCondicionalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterElseCondicional(s)
	}
}

func (s *ElseCondicionalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitElseCondicional(s)
	}
}

func (s *ElseCondicionalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitElseCondicional(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) ElseCondicional() (localctx IElseCondicionalContext) {
	localctx = NewElseCondicionalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, VlangParserRULE_elseCondicional)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(235)
		p.Match(VlangParserELSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(236)
		p.Match(VlangParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(240)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&720725682266268042) != 0 {
		{
			p.SetState(237)
			p.Declaraciones()
		}

		p.SetState(242)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(243)
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

type ForRangeSliceContext struct {
	ForDclContext
}

func NewForRangeSliceContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForRangeSliceContext {
	var p = new(ForRangeSliceContext)

	InitEmptyForDclContext(&p.ForDclContext)
	p.parser = parser
	p.CopyAll(ctx.(*ForDclContext))

	return p
}

func (s *ForRangeSliceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForRangeSliceContext) FOR() antlr.TerminalNode {
	return s.GetToken(VlangParserFOR, 0)
}

func (s *ForRangeSliceContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(VlangParserID)
}

func (s *ForRangeSliceContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(VlangParserID, i)
}

func (s *ForRangeSliceContext) COMMA() antlr.TerminalNode {
	return s.GetToken(VlangParserCOMMA, 0)
}

func (s *ForRangeSliceContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(VlangParserASSIGN, 0)
}

func (s *ForRangeSliceContext) Block() IBlockContext {
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

func (s *ForRangeSliceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterForRangeSlice(s)
	}
}

func (s *ForRangeSliceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitForRangeSlice(s)
	}
}

func (s *ForRangeSliceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitForRangeSlice(s)

	default:
		return t.VisitChildren(s)
	}
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

func (s *ForCondicionUnicaContext) FOR() antlr.TerminalNode {
	return s.GetToken(VlangParserFOR, 0)
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

func (s *ForClasicoContext) FOR() antlr.TerminalNode {
	return s.GetToken(VlangParserFOR, 0)
}

func (s *ForClasicoContext) Asignacion() IAsignacionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAsignacionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAsignacionContext)
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

func (s *ForClasicoContext) Stmt() IStmtContext {
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
	p.EnterRule(localctx, 28, VlangParserRULE_forDcl)
	var _la int

	p.SetState(267)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		localctx = NewForClasicoContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(245)
			p.Match(VlangParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(246)
			p.Asignacion()
		}
		{
			p.SetState(247)
			p.Match(VlangParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(248)
			p.expresion(0)
		}
		{
			p.SetState(249)
			p.Match(VlangParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(251)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&720725682266267904) != 0 {
			{
				p.SetState(250)
				p.Stmt()
			}

		}
		{
			p.SetState(253)
			p.Block()
		}

	case 2:
		localctx = NewForCondicionUnicaContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(255)
			p.Match(VlangParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(256)
			p.expresion(0)
		}
		{
			p.SetState(257)
			p.Block()
		}

	case 3:
		localctx = NewForRangeSliceContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(259)
			p.Match(VlangParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(260)
			p.Match(VlangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(261)
			p.Match(VlangParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(262)
			p.Match(VlangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(263)
			p.Match(VlangParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(264)
			p.Match(VlangParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(265)
			p.Match(VlangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(266)
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

// IAsignacionContext is an interface to support dynamic dispatch.
type IAsignacionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	ASSIGN() antlr.TerminalNode
	Expresion() IExpresionContext

	// IsAsignacionContext differentiates from other interfaces.
	IsAsignacionContext()
}

type AsignacionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAsignacionContext() *AsignacionContext {
	var p = new(AsignacionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_asignacion
	return p
}

func InitEmptyAsignacionContext(p *AsignacionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_asignacion
}

func (*AsignacionContext) IsAsignacionContext() {}

func NewAsignacionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AsignacionContext {
	var p = new(AsignacionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_asignacion

	return p
}

func (s *AsignacionContext) GetParser() antlr.Parser { return s.parser }

func (s *AsignacionContext) ID() antlr.TerminalNode {
	return s.GetToken(VlangParserID, 0)
}

func (s *AsignacionContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(VlangParserASSIGN, 0)
}

func (s *AsignacionContext) Expresion() IExpresionContext {
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

func (s *AsignacionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsignacionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AsignacionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterAsignacion(s)
	}
}

func (s *AsignacionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitAsignacion(s)
	}
}

func (s *AsignacionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitAsignacion(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Asignacion() (localctx IAsignacionContext) {
	localctx = NewAsignacionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, VlangParserRULE_asignacion)
	p.EnterOuterAlt(localctx, 1)
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
		p.Match(VlangParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(271)
		p.expresion(0)
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
	SWITCH() antlr.TerminalNode
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

func (s *SwitchDclContext) SWITCH() antlr.TerminalNode {
	return s.GetToken(VlangParserSWITCH, 0)
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
	p.EnterRule(localctx, 32, VlangParserRULE_switchDcl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(273)
		p.Match(VlangParserSWITCH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(274)
		p.expresion(0)
	}
	{
		p.SetState(275)
		p.Match(VlangParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(279)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == VlangParserT__4 {
		{
			p.SetState(276)
			p.CaseBlock()
		}

		p.SetState(281)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(283)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == VlangParserT__5 {
		{
			p.SetState(282)
			p.DefaultBlock()
		}

	}
	{
		p.SetState(285)
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
	p.EnterRule(localctx, 34, VlangParserRULE_caseBlock)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(287)
		p.Match(VlangParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(288)
		p.expresion(0)
	}
	{
		p.SetState(289)
		p.Match(VlangParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(293)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&720725682266268042) != 0 {
		{
			p.SetState(290)
			p.Declaraciones()
		}

		p.SetState(295)
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
	p.EnterRule(localctx, 36, VlangParserRULE_defaultBlock)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(296)
		p.Match(VlangParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(297)
		p.Match(VlangParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(301)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&720725682266268042) != 0 {
		{
			p.SetState(298)
			p.Declaraciones()
		}

		p.SetState(303)
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
	LEN() antlr.TerminalNode
	APPEND() antlr.TerminalNode

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

func (s *LlamadaFuncionContext) LEN() antlr.TerminalNode {
	return s.GetToken(VlangParserLEN, 0)
}

func (s *LlamadaFuncionContext) APPEND() antlr.TerminalNode {
	return s.GetToken(VlangParserAPPEND, 0)
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
	p.EnterRule(localctx, 38, VlangParserRULE_llamadaFuncion)
	var _la int

	p.SetState(369)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case VlangParserINDEXOF:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(304)
			p.Match(VlangParserINDEXOF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(305)
			p.Match(VlangParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(314)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&720725613486366720) != 0 {
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

			for _la == VlangParserCOMMA {
				{
					p.SetState(307)
					p.Match(VlangParserCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(308)
					p.expresion(0)
				}

				p.SetState(313)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(316)
			p.Match(VlangParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VlangParserJOIN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(317)
			p.Match(VlangParserJOIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(318)
			p.Match(VlangParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(327)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&720725613486366720) != 0 {
			{
				p.SetState(319)
				p.expresion(0)
			}
			p.SetState(324)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == VlangParserCOMMA {
				{
					p.SetState(320)
					p.Match(VlangParserCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
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
			}

		}
		{
			p.SetState(329)
			p.Match(VlangParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VlangParserID:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(330)
			p.Match(VlangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(331)
			p.Match(VlangParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(340)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&720725613486366720) != 0 {
			{
				p.SetState(332)
				p.expresion(0)
			}
			p.SetState(337)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == VlangParserCOMMA {
				{
					p.SetState(333)
					p.Match(VlangParserCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(334)
					p.expresion(0)
				}

				p.SetState(339)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(342)
			p.Match(VlangParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VlangParserLEN:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(343)
			p.Match(VlangParserLEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(344)
			p.Match(VlangParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(353)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&720725613486366720) != 0 {
			{
				p.SetState(345)
				p.expresion(0)
			}
			p.SetState(350)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == VlangParserCOMMA {
				{
					p.SetState(346)
					p.Match(VlangParserCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(347)
					p.expresion(0)
				}

				p.SetState(352)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(355)
			p.Match(VlangParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VlangParserAPPEND:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(356)
			p.Match(VlangParserAPPEND)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(357)
			p.Match(VlangParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(366)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&720725613486366720) != 0 {
			{
				p.SetState(358)
				p.expresion(0)
			}
			p.SetState(363)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == VlangParserCOMMA {
				{
					p.SetState(359)
					p.Match(VlangParserCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(360)
					p.expresion(0)
				}

				p.SetState(365)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(368)
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

// IFuncCallContext is an interface to support dynamic dispatch.
type IFuncCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	ParametrosReales() IParametrosRealesContext

	// IsFuncCallContext differentiates from other interfaces.
	IsFuncCallContext()
}

type FuncCallContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncCallContext() *FuncCallContext {
	var p = new(FuncCallContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_funcCall
	return p
}

func InitEmptyFuncCallContext(p *FuncCallContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_funcCall
}

func (*FuncCallContext) IsFuncCallContext() {}

func NewFuncCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncCallContext {
	var p = new(FuncCallContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_funcCall

	return p
}

func (s *FuncCallContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncCallContext) ID() antlr.TerminalNode {
	return s.GetToken(VlangParserID, 0)
}

func (s *FuncCallContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(VlangParserLPAREN, 0)
}

func (s *FuncCallContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(VlangParserRPAREN, 0)
}

func (s *FuncCallContext) ParametrosReales() IParametrosRealesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParametrosRealesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParametrosRealesContext)
}

func (s *FuncCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterFuncCall(s)
	}
}

func (s *FuncCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitFuncCall(s)
	}
}

func (s *FuncCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitFuncCall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) FuncCall() (localctx IFuncCallContext) {
	localctx = NewFuncCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, VlangParserRULE_funcCall)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(371)
		p.Match(VlangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(372)
		p.Match(VlangParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(374)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&720725613486366720) != 0 {
		{
			p.SetState(373)
			p.ParametrosReales()
		}

	}
	{
		p.SetState(376)
		p.Match(VlangParserRPAREN)
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

// IParametrosFormalesContext is an interface to support dynamic dispatch.
type IParametrosFormalesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllParametro() []IParametroContext
	Parametro(i int) IParametroContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsParametrosFormalesContext differentiates from other interfaces.
	IsParametrosFormalesContext()
}

type ParametrosFormalesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParametrosFormalesContext() *ParametrosFormalesContext {
	var p = new(ParametrosFormalesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_parametrosFormales
	return p
}

func InitEmptyParametrosFormalesContext(p *ParametrosFormalesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_parametrosFormales
}

func (*ParametrosFormalesContext) IsParametrosFormalesContext() {}

func NewParametrosFormalesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParametrosFormalesContext {
	var p = new(ParametrosFormalesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_parametrosFormales

	return p
}

func (s *ParametrosFormalesContext) GetParser() antlr.Parser { return s.parser }

func (s *ParametrosFormalesContext) AllParametro() []IParametroContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParametroContext); ok {
			len++
		}
	}

	tst := make([]IParametroContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParametroContext); ok {
			tst[i] = t.(IParametroContext)
			i++
		}
	}

	return tst
}

func (s *ParametrosFormalesContext) Parametro(i int) IParametroContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParametroContext); ok {
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

	return t.(IParametroContext)
}

func (s *ParametrosFormalesContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(VlangParserCOMMA)
}

func (s *ParametrosFormalesContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(VlangParserCOMMA, i)
}

func (s *ParametrosFormalesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParametrosFormalesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParametrosFormalesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterParametrosFormales(s)
	}
}

func (s *ParametrosFormalesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitParametrosFormales(s)
	}
}

func (s *ParametrosFormalesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitParametrosFormales(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) ParametrosFormales() (localctx IParametrosFormalesContext) {
	localctx = NewParametrosFormalesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, VlangParserRULE_parametrosFormales)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(378)
		p.Parametro()
	}
	p.SetState(383)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == VlangParserCOMMA {
		{
			p.SetState(379)
			p.Match(VlangParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(380)
			p.Parametro()
		}

		p.SetState(385)
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

// IParametroContext is an interface to support dynamic dispatch.
type IParametroContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	TIPO() antlr.TerminalNode

	// IsParametroContext differentiates from other interfaces.
	IsParametroContext()
}

type ParametroContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParametroContext() *ParametroContext {
	var p = new(ParametroContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_parametro
	return p
}

func InitEmptyParametroContext(p *ParametroContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_parametro
}

func (*ParametroContext) IsParametroContext() {}

func NewParametroContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParametroContext {
	var p = new(ParametroContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_parametro

	return p
}

func (s *ParametroContext) GetParser() antlr.Parser { return s.parser }

func (s *ParametroContext) ID() antlr.TerminalNode {
	return s.GetToken(VlangParserID, 0)
}

func (s *ParametroContext) TIPO() antlr.TerminalNode {
	return s.GetToken(VlangParserTIPO, 0)
}

func (s *ParametroContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParametroContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParametroContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterParametro(s)
	}
}

func (s *ParametroContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitParametro(s)
	}
}

func (s *ParametroContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitParametro(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Parametro() (localctx IParametroContext) {
	localctx = NewParametroContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, VlangParserRULE_parametro)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(386)
		p.Match(VlangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(387)
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

// IParametrosRealesContext is an interface to support dynamic dispatch.
type IParametrosRealesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpresion() []IExpresionContext
	Expresion(i int) IExpresionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsParametrosRealesContext differentiates from other interfaces.
	IsParametrosRealesContext()
}

type ParametrosRealesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParametrosRealesContext() *ParametrosRealesContext {
	var p = new(ParametrosRealesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_parametrosReales
	return p
}

func InitEmptyParametrosRealesContext(p *ParametrosRealesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_parametrosReales
}

func (*ParametrosRealesContext) IsParametrosRealesContext() {}

func NewParametrosRealesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParametrosRealesContext {
	var p = new(ParametrosRealesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_parametrosReales

	return p
}

func (s *ParametrosRealesContext) GetParser() antlr.Parser { return s.parser }

func (s *ParametrosRealesContext) AllExpresion() []IExpresionContext {
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

func (s *ParametrosRealesContext) Expresion(i int) IExpresionContext {
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

func (s *ParametrosRealesContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(VlangParserCOMMA)
}

func (s *ParametrosRealesContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(VlangParserCOMMA, i)
}

func (s *ParametrosRealesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParametrosRealesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParametrosRealesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterParametrosReales(s)
	}
}

func (s *ParametrosRealesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitParametrosReales(s)
	}
}

func (s *ParametrosRealesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitParametrosReales(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) ParametrosReales() (localctx IParametrosRealesContext) {
	localctx = NewParametrosRealesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, VlangParserRULE_parametrosReales)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(389)
		p.expresion(0)
	}
	p.SetState(394)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == VlangParserCOMMA {
		{
			p.SetState(390)
			p.Match(VlangParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(391)
			p.expresion(0)
		}

		p.SetState(396)
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
	LBRACE() antlr.TerminalNode
	AtributosStruct() IAtributosStructContext
	RBRACE() antlr.TerminalNode

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
	p.RuleIndex = VlangParserRULE_structDcl
	return p
}

func InitEmptyStructDclContext(p *StructDclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_structDcl
}

func (*StructDclContext) IsStructDclContext() {}

func NewStructDclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructDclContext {
	var p = new(StructDclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_structDcl

	return p
}

func (s *StructDclContext) GetParser() antlr.Parser { return s.parser }

func (s *StructDclContext) ID() antlr.TerminalNode {
	return s.GetToken(VlangParserID, 0)
}

func (s *StructDclContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(VlangParserLBRACE, 0)
}

func (s *StructDclContext) AtributosStruct() IAtributosStructContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAtributosStructContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAtributosStructContext)
}

func (s *StructDclContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(VlangParserRBRACE, 0)
}

func (s *StructDclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructDclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructDclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterStructDcl(s)
	}
}

func (s *StructDclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitStructDcl(s)
	}
}

func (s *StructDclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitStructDcl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) StructDcl() (localctx IStructDclContext) {
	localctx = NewStructDclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, VlangParserRULE_structDcl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(397)
		p.Match(VlangParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(398)
		p.Match(VlangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(399)
		p.Match(VlangParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(400)
		p.AtributosStruct()
	}
	{
		p.SetState(401)
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

// IAtributosStructContext is an interface to support dynamic dispatch.
type IAtributosStructContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllAtributoStruct() []IAtributoStructContext
	AtributoStruct(i int) IAtributoStructContext

	// IsAtributosStructContext differentiates from other interfaces.
	IsAtributosStructContext()
}

type AtributosStructContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtributosStructContext() *AtributosStructContext {
	var p = new(AtributosStructContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_atributosStruct
	return p
}

func InitEmptyAtributosStructContext(p *AtributosStructContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_atributosStruct
}

func (*AtributosStructContext) IsAtributosStructContext() {}

func NewAtributosStructContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtributosStructContext {
	var p = new(AtributosStructContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_atributosStruct

	return p
}

func (s *AtributosStructContext) GetParser() antlr.Parser { return s.parser }

func (s *AtributosStructContext) AllAtributoStruct() []IAtributoStructContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAtributoStructContext); ok {
			len++
		}
	}

	tst := make([]IAtributoStructContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAtributoStructContext); ok {
			tst[i] = t.(IAtributoStructContext)
			i++
		}
	}

	return tst
}

func (s *AtributosStructContext) AtributoStruct(i int) IAtributoStructContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAtributoStructContext); ok {
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

	return t.(IAtributoStructContext)
}

func (s *AtributosStructContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtributosStructContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtributosStructContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterAtributosStruct(s)
	}
}

func (s *AtributosStructContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitAtributosStruct(s)
	}
}

func (s *AtributosStructContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitAtributosStruct(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) AtributosStruct() (localctx IAtributosStructContext) {
	localctx = NewAtributosStructContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, VlangParserRULE_atributosStruct)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(404)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == VlangParserID {
		{
			p.SetState(403)
			p.AtributoStruct()
		}

		p.SetState(406)
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

// IAtributoStructContext is an interface to support dynamic dispatch.
type IAtributoStructContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAtributoStructContext differentiates from other interfaces.
	IsAtributoStructContext()
}

type AtributoStructContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtributoStructContext() *AtributoStructContext {
	var p = new(AtributoStructContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_atributoStruct
	return p
}

func InitEmptyAtributoStructContext(p *AtributoStructContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_atributoStruct
}

func (*AtributoStructContext) IsAtributoStructContext() {}

func NewAtributoStructContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtributoStructContext {
	var p = new(AtributoStructContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_atributoStruct

	return p
}

func (s *AtributoStructContext) GetParser() antlr.Parser { return s.parser }

func (s *AtributoStructContext) CopyAll(ctx *AtributoStructContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *AtributoStructContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtributoStructContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AtributoPrimitivoContext struct {
	AtributoStructContext
}

func NewAtributoPrimitivoContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AtributoPrimitivoContext {
	var p = new(AtributoPrimitivoContext)

	InitEmptyAtributoStructContext(&p.AtributoStructContext)
	p.parser = parser
	p.CopyAll(ctx.(*AtributoStructContext))

	return p
}

func (s *AtributoPrimitivoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtributoPrimitivoContext) ID() antlr.TerminalNode {
	return s.GetToken(VlangParserID, 0)
}

func (s *AtributoPrimitivoContext) TIPO() antlr.TerminalNode {
	return s.GetToken(VlangParserTIPO, 0)
}

func (s *AtributoPrimitivoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterAtributoPrimitivo(s)
	}
}

func (s *AtributoPrimitivoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitAtributoPrimitivo(s)
	}
}

func (s *AtributoPrimitivoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitAtributoPrimitivo(s)

	default:
		return t.VisitChildren(s)
	}
}

type AtributoStructAnidadoContext struct {
	AtributoStructContext
}

func NewAtributoStructAnidadoContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AtributoStructAnidadoContext {
	var p = new(AtributoStructAnidadoContext)

	InitEmptyAtributoStructContext(&p.AtributoStructContext)
	p.parser = parser
	p.CopyAll(ctx.(*AtributoStructContext))

	return p
}

func (s *AtributoStructAnidadoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtributoStructAnidadoContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(VlangParserID)
}

func (s *AtributoStructAnidadoContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(VlangParserID, i)
}

func (s *AtributoStructAnidadoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterAtributoStructAnidado(s)
	}
}

func (s *AtributoStructAnidadoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitAtributoStructAnidado(s)
	}
}

func (s *AtributoStructAnidadoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitAtributoStructAnidado(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) AtributoStruct() (localctx IAtributoStructContext) {
	localctx = NewAtributoStructContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, VlangParserRULE_atributoStruct)
	p.SetState(412)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 42, p.GetParserRuleContext()) {
	case 1:
		localctx = NewAtributoPrimitivoContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(408)
			p.Match(VlangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(409)
			p.Match(VlangParserTIPO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewAtributoStructAnidadoContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(410)
			p.Match(VlangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(411)
			p.Match(VlangParserID)
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

// IListaAsignacionesContext is an interface to support dynamic dispatch.
type IListaAsignacionesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllAsignacionStruct() []IAsignacionStructContext
	AsignacionStruct(i int) IAsignacionStructContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsListaAsignacionesContext differentiates from other interfaces.
	IsListaAsignacionesContext()
}

type ListaAsignacionesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListaAsignacionesContext() *ListaAsignacionesContext {
	var p = new(ListaAsignacionesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_listaAsignaciones
	return p
}

func InitEmptyListaAsignacionesContext(p *ListaAsignacionesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_listaAsignaciones
}

func (*ListaAsignacionesContext) IsListaAsignacionesContext() {}

func NewListaAsignacionesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListaAsignacionesContext {
	var p = new(ListaAsignacionesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_listaAsignaciones

	return p
}

func (s *ListaAsignacionesContext) GetParser() antlr.Parser { return s.parser }

func (s *ListaAsignacionesContext) AllAsignacionStruct() []IAsignacionStructContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAsignacionStructContext); ok {
			len++
		}
	}

	tst := make([]IAsignacionStructContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAsignacionStructContext); ok {
			tst[i] = t.(IAsignacionStructContext)
			i++
		}
	}

	return tst
}

func (s *ListaAsignacionesContext) AsignacionStruct(i int) IAsignacionStructContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAsignacionStructContext); ok {
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

	return t.(IAsignacionStructContext)
}

func (s *ListaAsignacionesContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(VlangParserCOMMA)
}

func (s *ListaAsignacionesContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(VlangParserCOMMA, i)
}

func (s *ListaAsignacionesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListaAsignacionesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListaAsignacionesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterListaAsignaciones(s)
	}
}

func (s *ListaAsignacionesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitListaAsignaciones(s)
	}
}

func (s *ListaAsignacionesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitListaAsignaciones(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) ListaAsignaciones() (localctx IListaAsignacionesContext) {
	localctx = NewListaAsignacionesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, VlangParserRULE_listaAsignaciones)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(414)
		p.AsignacionStruct()
	}
	p.SetState(419)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == VlangParserCOMMA {
		{
			p.SetState(415)
			p.Match(VlangParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(416)
			p.AsignacionStruct()
		}

		p.SetState(421)
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

// IAsignacionStructContext is an interface to support dynamic dispatch.
type IAsignacionStructContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	COLON() antlr.TerminalNode
	Expresion() IExpresionContext

	// IsAsignacionStructContext differentiates from other interfaces.
	IsAsignacionStructContext()
}

type AsignacionStructContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAsignacionStructContext() *AsignacionStructContext {
	var p = new(AsignacionStructContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_asignacionStruct
	return p
}

func InitEmptyAsignacionStructContext(p *AsignacionStructContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_asignacionStruct
}

func (*AsignacionStructContext) IsAsignacionStructContext() {}

func NewAsignacionStructContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AsignacionStructContext {
	var p = new(AsignacionStructContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_asignacionStruct

	return p
}

func (s *AsignacionStructContext) GetParser() antlr.Parser { return s.parser }

func (s *AsignacionStructContext) ID() antlr.TerminalNode {
	return s.GetToken(VlangParserID, 0)
}

func (s *AsignacionStructContext) COLON() antlr.TerminalNode {
	return s.GetToken(VlangParserCOLON, 0)
}

func (s *AsignacionStructContext) Expresion() IExpresionContext {
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

func (s *AsignacionStructContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsignacionStructContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AsignacionStructContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterAsignacionStruct(s)
	}
}

func (s *AsignacionStructContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitAsignacionStruct(s)
	}
}

func (s *AsignacionStructContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitAsignacionStruct(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) AsignacionStruct() (localctx IAsignacionStructContext) {
	localctx = NewAsignacionStructContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, VlangParserRULE_asignacionStruct)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(422)
		p.Match(VlangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(423)
		p.Match(VlangParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(424)
		p.expresion(0)
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
	p.EnterRule(localctx, 58, VlangParserRULE_whileDcl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(426)
		p.Match(VlangParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(427)
		p.Match(VlangParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(428)
		p.expresion(0)
	}
	{
		p.SetState(429)
		p.Match(VlangParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(430)
		p.Match(VlangParserLBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(434)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&720725682266268042) != 0 {
		{
			p.SetState(431)
			p.Declaraciones()
		}

		p.SetState(436)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(437)
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

type Casteo_paratipoContext struct {
	ExpresionContext
}

func NewCasteo_paratipoContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Casteo_paratipoContext {
	var p = new(Casteo_paratipoContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *Casteo_paratipoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Casteo_paratipoContext) TYPEOF() antlr.TerminalNode {
	return s.GetToken(VlangParserTYPEOF, 0)
}

func (s *Casteo_paratipoContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(VlangParserLPAREN, 0)
}

func (s *Casteo_paratipoContext) Expresion() IExpresionContext {
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

func (s *Casteo_paratipoContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(VlangParserRPAREN, 0)
}

func (s *Casteo_paratipoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterCasteo_paratipo(s)
	}
}

func (s *Casteo_paratipoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitCasteo_paratipo(s)
	}
}

func (s *Casteo_paratipoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitCasteo_paratipo(s)

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

type StructInstanceCreationContext struct {
	ExpresionContext
}

func NewStructInstanceCreationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructInstanceCreationContext {
	var p = new(StructInstanceCreationContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *StructInstanceCreationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructInstanceCreationContext) ID() antlr.TerminalNode {
	return s.GetToken(VlangParserID, 0)
}

func (s *StructInstanceCreationContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(VlangParserLBRACE, 0)
}

func (s *StructInstanceCreationContext) ListaAsignaciones() IListaAsignacionesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListaAsignacionesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListaAsignacionesContext)
}

func (s *StructInstanceCreationContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(VlangParserRBRACE, 0)
}

func (s *StructInstanceCreationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterStructInstanceCreation(s)
	}
}

func (s *StructInstanceCreationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitStructInstanceCreation(s)
	}
}

func (s *StructInstanceCreationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitStructInstanceCreation(s)

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

type StructAttrAssignContext struct {
	ExpresionContext
}

func NewStructAttrAssignContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructAttrAssignContext {
	var p = new(StructAttrAssignContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *StructAttrAssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructAttrAssignContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(VlangParserID)
}

func (s *StructAttrAssignContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(VlangParserID, i)
}

func (s *StructAttrAssignContext) DOT() antlr.TerminalNode {
	return s.GetToken(VlangParserDOT, 0)
}

func (s *StructAttrAssignContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(VlangParserASSIGN, 0)
}

func (s *StructAttrAssignContext) Expresion() IExpresionContext {
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

func (s *StructAttrAssignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterStructAttrAssign(s)
	}
}

func (s *StructAttrAssignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitStructAttrAssign(s)
	}
}

func (s *StructAttrAssignContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitStructAttrAssign(s)

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

type Casteo_paratipo_sliceContext struct {
	ExpresionContext
}

func NewCasteo_paratipo_sliceContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Casteo_paratipo_sliceContext {
	var p = new(Casteo_paratipo_sliceContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *Casteo_paratipo_sliceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Casteo_paratipo_sliceContext) TYPEOF() antlr.TerminalNode {
	return s.GetToken(VlangParserTYPEOF, 0)
}

func (s *Casteo_paratipo_sliceContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(VlangParserLPAREN, 0)
}

func (s *Casteo_paratipo_sliceContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(VlangParserLBRACK, 0)
}

func (s *Casteo_paratipo_sliceContext) AllExpresion() []IExpresionContext {
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

func (s *Casteo_paratipo_sliceContext) Expresion(i int) IExpresionContext {
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

func (s *Casteo_paratipo_sliceContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(VlangParserRBRACK, 0)
}

func (s *Casteo_paratipo_sliceContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(VlangParserRPAREN, 0)
}

func (s *Casteo_paratipo_sliceContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(VlangParserCOMMA)
}

func (s *Casteo_paratipo_sliceContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(VlangParserCOMMA, i)
}

func (s *Casteo_paratipo_sliceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterCasteo_paratipo_slice(s)
	}
}

func (s *Casteo_paratipo_sliceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitCasteo_paratipo_slice(s)
	}
}

func (s *Casteo_paratipo_sliceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitCasteo_paratipo_slice(s)

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

type PARAPRINTSLICEContext struct {
	ExpresionContext
}

func NewPARAPRINTSLICEContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PARAPRINTSLICEContext {
	var p = new(PARAPRINTSLICEContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *PARAPRINTSLICEContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PARAPRINTSLICEContext) ID() antlr.TerminalNode {
	return s.GetToken(VlangParserID, 0)
}

func (s *PARAPRINTSLICEContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(VlangParserLBRACK, 0)
}

func (s *PARAPRINTSLICEContext) Expresion() IExpresionContext {
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

func (s *PARAPRINTSLICEContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(VlangParserRBRACK, 0)
}

func (s *PARAPRINTSLICEContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterPARAPRINTSLICE(s)
	}
}

func (s *PARAPRINTSLICEContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitPARAPRINTSLICE(s)
	}
}

func (s *PARAPRINTSLICEContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitPARAPRINTSLICE(s)

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
	_startState := 60
	p.EnterRecursionRule(localctx, 60, VlangParserRULE_expresion, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(501)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 46, p.GetParserRuleContext()) {
	case 1:
		localctx = NewUnarioContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(440)

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
			p.SetState(441)
			p.expresion(18)
		}

	case 2:
		localctx = NewValorexprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(442)
			p.Valor()
		}

	case 3:
		localctx = NewParentesisexpreContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(443)
			p.Match(VlangParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(444)
			p.expresion(0)
		}
		{
			p.SetState(445)
			p.Match(VlangParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewCorchetesexpreContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(447)
			p.Match(VlangParserLBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(448)
			p.expresion(0)
		}
		{
			p.SetState(449)
			p.Match(VlangParserRBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		localctx = NewPARAPRINTSLICEContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(451)
			p.Match(VlangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(452)
			p.Match(VlangParserLBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(453)
			p.expresion(0)
		}
		{
			p.SetState(454)
			p.Match(VlangParserRBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		localctx = NewLlamadaFuncionExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(456)
			p.LlamadaFuncion()
		}

	case 7:
		localctx = NewStructAttrAssignContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(457)
			p.Match(VlangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(458)
			p.Match(VlangParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(459)
			p.Match(VlangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(460)
			p.Match(VlangParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(461)
			p.expresion(10)
		}

	case 8:
		localctx = NewStructInstanceCreationContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(462)
			p.Match(VlangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(463)
			p.Match(VlangParserLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(464)
			p.ListaAsignaciones()
		}
		{
			p.SetState(465)
			p.Match(VlangParserRBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 9:
		localctx = NewIdContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(467)
			p.Match(VlangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 10:
		localctx = NewIncredecrContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(468)
			p.Incredecre()
		}

	case 11:
		localctx = NewExpdotexp1Context(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(469)
			p.Match(VlangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(470)
			p.Match(VlangParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(471)
			p.Match(VlangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 12:
		localctx = NewExpdotexpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(472)
			p.Match(VlangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(473)
			p.Match(VlangParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(474)
			p.expresion(5)
		}

	case 13:
		localctx = NewAsignacionLUEGOContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(475)
			p.Match(VlangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(476)
			p.Match(VlangParserTIPO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(477)
			p.Match(VlangParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(478)
			p.expresion(4)
		}

	case 14:
		localctx = NewIMCPLICITContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(479)
			p.Match(VlangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(480)

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
			p.SetState(481)
			p.expresion(3)
		}

	case 15:
		localctx = NewCasteo_paratipo_sliceContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(482)
			p.Match(VlangParserTYPEOF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(483)
			p.Match(VlangParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(484)
			p.Match(VlangParserLBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(485)
			p.expresion(0)
		}
		p.SetState(490)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == VlangParserCOMMA {
			{
				p.SetState(486)
				p.Match(VlangParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(487)
				p.expresion(0)
			}

			p.SetState(492)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(493)
			p.Match(VlangParserRBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(494)
			p.Match(VlangParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 16:
		localctx = NewCasteo_paratipoContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(496)
			p.Match(VlangParserTYPEOF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(497)
			p.Match(VlangParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(498)
			p.expresion(0)
		}
		{
			p.SetState(499)
			p.Match(VlangParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(520)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 48, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(518)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 47, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMultdivmodContext(p, NewExpresionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, VlangParserRULE_expresion)
				p.SetState(503)

				if !(p.Precpred(p.GetParserRuleContext(), 21)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 21)", ""))
					goto errorExit
				}
				{
					p.SetState(504)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*MultdivmodContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&123145302310912) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*MultdivmodContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(505)
					p.expresion(22)
				}

			case 2:
				localctx = NewSumresContext(p, NewExpresionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, VlangParserRULE_expresion)
				p.SetState(506)

				if !(p.Precpred(p.GetParserRuleContext(), 20)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 20)", ""))
					goto errorExit
				}
				{
					p.SetState(507)

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
					p.SetState(508)
					p.expresion(21)
				}

			case 3:
				localctx = NewIgualdadContext(p, NewExpresionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, VlangParserRULE_expresion)
				p.SetState(509)

				if !(p.Precpred(p.GetParserRuleContext(), 19)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 19)", ""))
					goto errorExit
				}
				{
					p.SetState(510)

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
					p.SetState(511)
					p.expresion(20)
				}

			case 4:
				localctx = NewRelacionalesContext(p, NewExpresionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, VlangParserRULE_expresion)
				p.SetState(512)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
					goto errorExit
				}
				{
					p.SetState(513)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*RelacionalesContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&67553994410557440) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*RelacionalesContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(514)
					p.expresion(18)
				}

			case 5:
				localctx = NewOPERADORESLOGICOSContext(p, NewExpresionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, VlangParserRULE_expresion)
				p.SetState(515)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
					goto errorExit
				}
				{
					p.SetState(516)

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
					p.SetState(517)
					p.expresion(17)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(522)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 48, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 62, VlangParserRULE_parametros)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(523)
		p.expresion(0)
	}
	p.SetState(528)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == VlangParserCOMMA {
		{
			p.SetState(524)
			p.Match(VlangParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(525)
			p.expresion(0)
		}

		p.SetState(530)
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
	p.EnterRule(localctx, 64, VlangParserRULE_valores)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(531)
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
	p.EnterRule(localctx, 66, VlangParserRULE_valor)
	p.SetState(538)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case VlangParserENTERO:
		localctx = NewValorEnteroContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(533)
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
			p.SetState(534)
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
			p.SetState(535)
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
			p.SetState(536)
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
			p.SetState(537)
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
	p.EnterRule(localctx, 68, VlangParserRULE_listaExpresiones)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(540)
		p.expresion(0)
	}
	p.SetState(545)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == VlangParserCOMMA {
		{
			p.SetState(541)
			p.Match(VlangParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(542)
			p.expresion(0)
		}

		p.SetState(547)
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
	p.EnterRule(localctx, 70, VlangParserRULE_incredecre)
	p.SetState(552)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 52, p.GetParserRuleContext()) {
	case 1:
		localctx = NewIncrementoContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(548)
			p.Match(VlangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(549)
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
			p.SetState(550)
			p.Match(VlangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(551)
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
	case 30:
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
		return p.Precpred(p.GetParserRuleContext(), 21)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 20)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 19)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 17)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 16)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
