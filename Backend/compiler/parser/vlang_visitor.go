// Code generated from parser/Vlang.g4 by ANTLR 4.13.2. DO NOT EDIT.

package compiler // Vlang
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by VlangParser.
type VlangVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by VlangParser#programa.
	VisitPrograma(ctx *ProgramaContext) interface{}

	// Visit a parse tree produced by VlangParser#funcMain.
	VisitFuncMain(ctx *FuncMainContext) interface{}

	// Visit a parse tree produced by VlangParser#funcDcl.
	VisitFuncDcl(ctx *FuncDclContext) interface{}

	// Visit a parse tree produced by VlangParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by VlangParser#declaraciones.
	VisitDeclaraciones(ctx *DeclaracionesContext) interface{}

	// Visit a parse tree produced by VlangParser#variableDeclaration.
	VisitVariableDeclaration(ctx *VariableDeclarationContext) interface{}

	// Visit a parse tree produced by VlangParser#sliceEmptyDeclaration.
	VisitSliceEmptyDeclaration(ctx *SliceEmptyDeclarationContext) interface{}

	// Visit a parse tree produced by VlangParser#structDirectInitDeclaration.
	VisitStructDirectInitDeclaration(ctx *StructDirectInitDeclarationContext) interface{}

	// Visit a parse tree produced by VlangParser#sliceInitDeclaration.
	VisitSliceInitDeclaration(ctx *SliceInitDeclarationContext) interface{}

	// Visit a parse tree produced by VlangParser#sliceAssignment.
	VisitSliceAssignment(ctx *SliceAssignmentContext) interface{}

	// Visit a parse tree produced by VlangParser#variableDeclarationImmutable.
	VisitVariableDeclarationImmutable(ctx *VariableDeclarationImmutableContext) interface{}

	// Visit a parse tree produced by VlangParser#variableCastDeclaration.
	VisitVariableCastDeclaration(ctx *VariableCastDeclarationContext) interface{}

	// Visit a parse tree produced by VlangParser#sliceAssignmentIndex.
	VisitSliceAssignmentIndex(ctx *SliceAssignmentIndexContext) interface{}

	// Visit a parse tree produced by VlangParser#sliceTipo.
	VisitSliceTipo(ctx *SliceTipoContext) interface{}

	// Visit a parse tree produced by VlangParser#sliceInit.
	VisitSliceInit(ctx *SliceInitContext) interface{}

	// Visit a parse tree produced by VlangParser#printStatement.
	VisitPrintStatement(ctx *PrintStatementContext) interface{}

	// Visit a parse tree produced by VlangParser#expresionStatement.
	VisitExpresionStatement(ctx *ExpresionStatementContext) interface{}

	// Visit a parse tree produced by VlangParser#controlStatement.
	VisitControlStatement(ctx *ControlStatementContext) interface{}

	// Visit a parse tree produced by VlangParser#transfersentence.
	VisitTransfersentence(ctx *TransfersentenceContext) interface{}

	// Visit a parse tree produced by VlangParser#if_context.
	VisitIf_context(ctx *If_contextContext) interface{}

	// Visit a parse tree produced by VlangParser#for_context.
	VisitFor_context(ctx *For_contextContext) interface{}

	// Visit a parse tree produced by VlangParser#switch_context.
	VisitSwitch_context(ctx *Switch_contextContext) interface{}

	// Visit a parse tree produced by VlangParser#while_context.
	VisitWhile_context(ctx *While_contextContext) interface{}

	// Visit a parse tree produced by VlangParser#breakStatement.
	VisitBreakStatement(ctx *BreakStatementContext) interface{}

	// Visit a parse tree produced by VlangParser#continueStatement.
	VisitContinueStatement(ctx *ContinueStatementContext) interface{}

	// Visit a parse tree produced by VlangParser#returnStatement.
	VisitReturnStatement(ctx *ReturnStatementContext) interface{}

	// Visit a parse tree produced by VlangParser#ifDcl.
	VisitIfDcl(ctx *IfDclContext) interface{}

	// Visit a parse tree produced by VlangParser#elseIfDcl.
	VisitElseIfDcl(ctx *ElseIfDclContext) interface{}

	// Visit a parse tree produced by VlangParser#elseCondicional.
	VisitElseCondicional(ctx *ElseCondicionalContext) interface{}

	// Visit a parse tree produced by VlangParser#forClasico.
	VisitForClasico(ctx *ForClasicoContext) interface{}

	// Visit a parse tree produced by VlangParser#forCondicionUnica.
	VisitForCondicionUnica(ctx *ForCondicionUnicaContext) interface{}

	// Visit a parse tree produced by VlangParser#forRangeSlice.
	VisitForRangeSlice(ctx *ForRangeSliceContext) interface{}

	// Visit a parse tree produced by VlangParser#asignacion.
	VisitAsignacion(ctx *AsignacionContext) interface{}

	// Visit a parse tree produced by VlangParser#switchDcl.
	VisitSwitchDcl(ctx *SwitchDclContext) interface{}

	// Visit a parse tree produced by VlangParser#caseBlock.
	VisitCaseBlock(ctx *CaseBlockContext) interface{}

	// Visit a parse tree produced by VlangParser#defaultBlock.
	VisitDefaultBlock(ctx *DefaultBlockContext) interface{}

	// Visit a parse tree produced by VlangParser#llamadaFuncion.
	VisitLlamadaFuncion(ctx *LlamadaFuncionContext) interface{}

	// Visit a parse tree produced by VlangParser#funcCall.
	VisitFuncCall(ctx *FuncCallContext) interface{}

	// Visit a parse tree produced by VlangParser#parametrosFormales.
	VisitParametrosFormales(ctx *ParametrosFormalesContext) interface{}

	// Visit a parse tree produced by VlangParser#parametro.
	VisitParametro(ctx *ParametroContext) interface{}

	// Visit a parse tree produced by VlangParser#parametrosReales.
	VisitParametrosReales(ctx *ParametrosRealesContext) interface{}

	// Visit a parse tree produced by VlangParser#structDcl.
	VisitStructDcl(ctx *StructDclContext) interface{}

	// Visit a parse tree produced by VlangParser#atributosStruct.
	VisitAtributosStruct(ctx *AtributosStructContext) interface{}

	// Visit a parse tree produced by VlangParser#atributoPrimitivo.
	VisitAtributoPrimitivo(ctx *AtributoPrimitivoContext) interface{}

	// Visit a parse tree produced by VlangParser#atributoStructAnidado.
	VisitAtributoStructAnidado(ctx *AtributoStructAnidadoContext) interface{}

	// Visit a parse tree produced by VlangParser#listaAsignaciones.
	VisitListaAsignaciones(ctx *ListaAsignacionesContext) interface{}

	// Visit a parse tree produced by VlangParser#asignacionStruct.
	VisitAsignacionStruct(ctx *AsignacionStructContext) interface{}

	// Visit a parse tree produced by VlangParser#whileDcl.
	VisitWhileDcl(ctx *WhileDclContext) interface{}

	// Visit a parse tree produced by VlangParser#multdivmod.
	VisitMultdivmod(ctx *MultdivmodContext) interface{}

	// Visit a parse tree produced by VlangParser#casteo_paratipo.
	VisitCasteo_paratipo(ctx *Casteo_paratipoContext) interface{}

	// Visit a parse tree produced by VlangParser#incredecr.
	VisitIncredecr(ctx *IncredecrContext) interface{}

	// Visit a parse tree produced by VlangParser#OPERADORESLOGICOS.
	VisitOPERADORESLOGICOS(ctx *OPERADORESLOGICOSContext) interface{}

	// Visit a parse tree produced by VlangParser#structInstanceCreation.
	VisitStructInstanceCreation(ctx *StructInstanceCreationContext) interface{}

	// Visit a parse tree produced by VlangParser#valorexpr.
	VisitValorexpr(ctx *ValorexprContext) interface{}

	// Visit a parse tree produced by VlangParser#igualdad.
	VisitIgualdad(ctx *IgualdadContext) interface{}

	// Visit a parse tree produced by VlangParser#llamadaFuncionExpr.
	VisitLlamadaFuncionExpr(ctx *LlamadaFuncionExprContext) interface{}

	// Visit a parse tree produced by VlangParser#expdotexp.
	VisitExpdotexp(ctx *ExpdotexpContext) interface{}

	// Visit a parse tree produced by VlangParser#structAttrAssign.
	VisitStructAttrAssign(ctx *StructAttrAssignContext) interface{}

	// Visit a parse tree produced by VlangParser#relacionales.
	VisitRelacionales(ctx *RelacionalesContext) interface{}

	// Visit a parse tree produced by VlangParser#casteo_paratipo_slice.
	VisitCasteo_paratipo_slice(ctx *Casteo_paratipo_sliceContext) interface{}

	// Visit a parse tree produced by VlangParser#corchetesexpre.
	VisitCorchetesexpre(ctx *CorchetesexpreContext) interface{}

	// Visit a parse tree produced by VlangParser#unario.
	VisitUnario(ctx *UnarioContext) interface{}

	// Visit a parse tree produced by VlangParser#parentesisexpre.
	VisitParentesisexpre(ctx *ParentesisexpreContext) interface{}

	// Visit a parse tree produced by VlangParser#IMCPLICIT.
	VisitIMCPLICIT(ctx *IMCPLICITContext) interface{}

	// Visit a parse tree produced by VlangParser#sumres.
	VisitSumres(ctx *SumresContext) interface{}

	// Visit a parse tree produced by VlangParser#PARAPRINTSLICE.
	VisitPARAPRINTSLICE(ctx *PARAPRINTSLICEContext) interface{}

	// Visit a parse tree produced by VlangParser#asignacionLUEGO.
	VisitAsignacionLUEGO(ctx *AsignacionLUEGOContext) interface{}

	// Visit a parse tree produced by VlangParser#id.
	VisitId(ctx *IdContext) interface{}

	// Visit a parse tree produced by VlangParser#expdotexp1.
	VisitExpdotexp1(ctx *Expdotexp1Context) interface{}

	// Visit a parse tree produced by VlangParser#parametros.
	VisitParametros(ctx *ParametrosContext) interface{}

	// Visit a parse tree produced by VlangParser#valores.
	VisitValores(ctx *ValoresContext) interface{}

	// Visit a parse tree produced by VlangParser#valorEntero.
	VisitValorEntero(ctx *ValorEnteroContext) interface{}

	// Visit a parse tree produced by VlangParser#valorDecimal.
	VisitValorDecimal(ctx *ValorDecimalContext) interface{}

	// Visit a parse tree produced by VlangParser#valorCadena.
	VisitValorCadena(ctx *ValorCadenaContext) interface{}

	// Visit a parse tree produced by VlangParser#valorBooleano.
	VisitValorBooleano(ctx *ValorBooleanoContext) interface{}

	// Visit a parse tree produced by VlangParser#valorCaracter.
	VisitValorCaracter(ctx *ValorCaracterContext) interface{}

	// Visit a parse tree produced by VlangParser#listaExpresiones.
	VisitListaExpresiones(ctx *ListaExpresionesContext) interface{}

	// Visit a parse tree produced by VlangParser#incremento.
	VisitIncremento(ctx *IncrementoContext) interface{}

	// Visit a parse tree produced by VlangParser#decremento.
	VisitDecremento(ctx *DecrementoContext) interface{}
}
