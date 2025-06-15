// Code generated from parser/Vlang.g4 by ANTLR 4.13.2. DO NOT EDIT.

package compiler // Vlang
import "github.com/antlr4-go/antlr/v4"

type BaseVlangVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseVlangVisitor) VisitPrograma(ctx *ProgramaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitFuncMain(ctx *FuncMainContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitFuncDcl(ctx *FuncDclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitDeclaraciones(ctx *DeclaracionesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitVariableDeclaration(ctx *VariableDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitSliceEmptyDeclaration(ctx *SliceEmptyDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitStructDirectInitDeclaration(ctx *StructDirectInitDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitSliceInitDeclaration(ctx *SliceInitDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitSliceAssignment(ctx *SliceAssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitVariableDeclarationImmutable(ctx *VariableDeclarationImmutableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitVariableCastDeclaration(ctx *VariableCastDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitSliceAssignmentIndex(ctx *SliceAssignmentIndexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitSliceTipo(ctx *SliceTipoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitSliceInit(ctx *SliceInitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitPrintStatement(ctx *PrintStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitExpresionStatement(ctx *ExpresionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitControlStatement(ctx *ControlStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitTransfersentence(ctx *TransfersentenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitIf_context(ctx *If_contextContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitFor_context(ctx *For_contextContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitSwitch_context(ctx *Switch_contextContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitWhile_context(ctx *While_contextContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitBreakStatement(ctx *BreakStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitContinueStatement(ctx *ContinueStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitReturnStatement(ctx *ReturnStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitIfDcl(ctx *IfDclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitElseIfDcl(ctx *ElseIfDclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitElseCondicional(ctx *ElseCondicionalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitForClasico(ctx *ForClasicoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitForCondicionUnica(ctx *ForCondicionUnicaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitForRangeSlice(ctx *ForRangeSliceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitAsignacion(ctx *AsignacionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitSwitchDcl(ctx *SwitchDclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitCaseBlock(ctx *CaseBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitDefaultBlock(ctx *DefaultBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitLlamadaFuncion(ctx *LlamadaFuncionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitFuncCall(ctx *FuncCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitParametrosFormales(ctx *ParametrosFormalesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitParametro(ctx *ParametroContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitParametrosReales(ctx *ParametrosRealesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitStructDcl(ctx *StructDclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitAtributosStruct(ctx *AtributosStructContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitAtributoPrimitivo(ctx *AtributoPrimitivoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitAtributoStructAnidado(ctx *AtributoStructAnidadoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitListaAsignaciones(ctx *ListaAsignacionesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitAsignacionStruct(ctx *AsignacionStructContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitWhileDcl(ctx *WhileDclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitMultdivmod(ctx *MultdivmodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitCasteo_paratipo(ctx *Casteo_paratipoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitIncredecr(ctx *IncredecrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitOPERADORESLOGICOS(ctx *OPERADORESLOGICOSContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitStructInstanceCreation(ctx *StructInstanceCreationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitValorexpr(ctx *ValorexprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitIgualdad(ctx *IgualdadContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitLlamadaFuncionExpr(ctx *LlamadaFuncionExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitExpdotexp(ctx *ExpdotexpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitStructAttrAssign(ctx *StructAttrAssignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitRelacionales(ctx *RelacionalesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitCasteo_paratipo_slice(ctx *Casteo_paratipo_sliceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitCorchetesexpre(ctx *CorchetesexpreContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitUnario(ctx *UnarioContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitParentesisexpre(ctx *ParentesisexpreContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitIMCPLICIT(ctx *IMCPLICITContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitSumres(ctx *SumresContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitPARAPRINTSLICE(ctx *PARAPRINTSLICEContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitAsignacionLUEGO(ctx *AsignacionLUEGOContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitId(ctx *IdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitExpdotexp1(ctx *Expdotexp1Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitParametros(ctx *ParametrosContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitValores(ctx *ValoresContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitValorEntero(ctx *ValorEnteroContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitValorDecimal(ctx *ValorDecimalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitValorCadena(ctx *ValorCadenaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitValorBooleano(ctx *ValorBooleanoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitValorCaracter(ctx *ValorCaracterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitListaExpresiones(ctx *ListaExpresionesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitIncremento(ctx *IncrementoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitDecremento(ctx *DecrementoContext) interface{} {
	return v.VisitChildren(ctx)
}
