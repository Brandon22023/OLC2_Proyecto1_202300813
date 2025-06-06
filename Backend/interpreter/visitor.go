package interpreter

import (
	"fmt"
	"strings"

	"backend/parser" // Cambia esto a tu path real del paquete generado
)

// V4LangVisitor implementa el visitor generado para la gramática V4LangGrammar
type V4LangVisitor struct {
	*parser.BaseV4LangGrammarVisitor
}

func NewV4LangVisitor() *V4LangVisitor {
	return &V4LangVisitor{BaseV4LangGrammarVisitor: &parser.BaseV4LangGrammarVisitor{}}
}

// VisitProgram visita el nodo raíz 'program'
func (v *V4LangVisitor) VisitProgram(ctx *parser.ProgramContext) interface{} {
	var sb strings.Builder
	for _, dclCtx := range ctx.AllDcl() {
		val := dclCtx.Accept(v)
		sb.WriteString(fmt.Sprintf("%s\n", val))
	}
	return sb.String()
}

// VisitVarDcl procesa declaraciones de variables
func (v *V4LangVisitor) VisitVarDcl(ctx *parser.VarDclContext) interface{} {
	return "Var Declaration: " + ctx.GetText()
}

// VisitFuncDcl procesa declaraciones de funciones
func (v *V4LangVisitor) VisitFuncDcl(ctx *parser.FuncDclContext) interface{} {
	return "Function Declaration: " + ctx.GetText()
}

// VisitStatement procesa sentencias variadas
func (v *V4LangVisitor) VisitStatement(ctx *parser.StatementContext) interface{} {
	return "Statement: " + ctx.GetText()
}
