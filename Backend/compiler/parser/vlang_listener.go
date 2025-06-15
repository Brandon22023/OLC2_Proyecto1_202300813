// Code generated from parser/Vlang.g4 by ANTLR 4.13.2. DO NOT EDIT.

package compiler // Vlang
import "github.com/antlr4-go/antlr/v4"

// VlangListener is a complete listener for a parse tree produced by VlangParser.
type VlangListener interface {
	antlr.ParseTreeListener

	// EnterPrograma is called when entering the programa production.
	EnterPrograma(c *ProgramaContext)

	// EnterFuncMain is called when entering the funcMain production.
	EnterFuncMain(c *FuncMainContext)

	// EnterFuncDcl is called when entering the funcDcl production.
	EnterFuncDcl(c *FuncDclContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterDeclaraciones is called when entering the declaraciones production.
	EnterDeclaraciones(c *DeclaracionesContext)

	// EnterVariableDeclaration is called when entering the variableDeclaration production.
	EnterVariableDeclaration(c *VariableDeclarationContext)

	// EnterSliceEmptyDeclaration is called when entering the sliceEmptyDeclaration production.
	EnterSliceEmptyDeclaration(c *SliceEmptyDeclarationContext)

	// EnterStructDirectInitDeclaration is called when entering the structDirectInitDeclaration production.
	EnterStructDirectInitDeclaration(c *StructDirectInitDeclarationContext)

	// EnterSliceInitDeclaration is called when entering the sliceInitDeclaration production.
	EnterSliceInitDeclaration(c *SliceInitDeclarationContext)

	// EnterSliceAssignment is called when entering the sliceAssignment production.
	EnterSliceAssignment(c *SliceAssignmentContext)

	// EnterVariableDeclarationImmutable is called when entering the variableDeclarationImmutable production.
	EnterVariableDeclarationImmutable(c *VariableDeclarationImmutableContext)

	// EnterVariableCastDeclaration is called when entering the variableCastDeclaration production.
	EnterVariableCastDeclaration(c *VariableCastDeclarationContext)

	// EnterSliceAssignmentIndex is called when entering the sliceAssignmentIndex production.
	EnterSliceAssignmentIndex(c *SliceAssignmentIndexContext)

	// EnterSliceTipo is called when entering the sliceTipo production.
	EnterSliceTipo(c *SliceTipoContext)

	// EnterSliceInit is called when entering the sliceInit production.
	EnterSliceInit(c *SliceInitContext)

	// EnterPrintStatement is called when entering the printStatement production.
	EnterPrintStatement(c *PrintStatementContext)

	// EnterExpresionStatement is called when entering the expresionStatement production.
	EnterExpresionStatement(c *ExpresionStatementContext)

	// EnterControlStatement is called when entering the controlStatement production.
	EnterControlStatement(c *ControlStatementContext)

	// EnterTransfersentence is called when entering the transfersentence production.
	EnterTransfersentence(c *TransfersentenceContext)

	// EnterIf_context is called when entering the if_context production.
	EnterIf_context(c *If_contextContext)

	// EnterFor_context is called when entering the for_context production.
	EnterFor_context(c *For_contextContext)

	// EnterSwitch_context is called when entering the switch_context production.
	EnterSwitch_context(c *Switch_contextContext)

	// EnterWhile_context is called when entering the while_context production.
	EnterWhile_context(c *While_contextContext)

	// EnterBreakStatement is called when entering the breakStatement production.
	EnterBreakStatement(c *BreakStatementContext)

	// EnterContinueStatement is called when entering the continueStatement production.
	EnterContinueStatement(c *ContinueStatementContext)

	// EnterReturnStatement is called when entering the returnStatement production.
	EnterReturnStatement(c *ReturnStatementContext)

	// EnterIfDcl is called when entering the ifDcl production.
	EnterIfDcl(c *IfDclContext)

	// EnterElseIfDcl is called when entering the elseIfDcl production.
	EnterElseIfDcl(c *ElseIfDclContext)

	// EnterElseCondicional is called when entering the elseCondicional production.
	EnterElseCondicional(c *ElseCondicionalContext)

	// EnterForClasico is called when entering the forClasico production.
	EnterForClasico(c *ForClasicoContext)

	// EnterForCondicionUnica is called when entering the forCondicionUnica production.
	EnterForCondicionUnica(c *ForCondicionUnicaContext)

	// EnterForRangeSlice is called when entering the forRangeSlice production.
	EnterForRangeSlice(c *ForRangeSliceContext)

	// EnterAsignacion is called when entering the asignacion production.
	EnterAsignacion(c *AsignacionContext)

	// EnterSwitchDcl is called when entering the switchDcl production.
	EnterSwitchDcl(c *SwitchDclContext)

	// EnterCaseBlock is called when entering the caseBlock production.
	EnterCaseBlock(c *CaseBlockContext)

	// EnterDefaultBlock is called when entering the defaultBlock production.
	EnterDefaultBlock(c *DefaultBlockContext)

	// EnterLlamadaFuncion is called when entering the llamadaFuncion production.
	EnterLlamadaFuncion(c *LlamadaFuncionContext)

	// EnterFuncCall is called when entering the funcCall production.
	EnterFuncCall(c *FuncCallContext)

	// EnterParametrosFormales is called when entering the parametrosFormales production.
	EnterParametrosFormales(c *ParametrosFormalesContext)

	// EnterParametro is called when entering the parametro production.
	EnterParametro(c *ParametroContext)

	// EnterParametrosReales is called when entering the parametrosReales production.
	EnterParametrosReales(c *ParametrosRealesContext)

	// EnterStructDcl is called when entering the structDcl production.
	EnterStructDcl(c *StructDclContext)

	// EnterAtributosStruct is called when entering the atributosStruct production.
	EnterAtributosStruct(c *AtributosStructContext)

	// EnterAtributoPrimitivo is called when entering the atributoPrimitivo production.
	EnterAtributoPrimitivo(c *AtributoPrimitivoContext)

	// EnterAtributoStructAnidado is called when entering the atributoStructAnidado production.
	EnterAtributoStructAnidado(c *AtributoStructAnidadoContext)

	// EnterListaAsignaciones is called when entering the listaAsignaciones production.
	EnterListaAsignaciones(c *ListaAsignacionesContext)

	// EnterAsignacionStruct is called when entering the asignacionStruct production.
	EnterAsignacionStruct(c *AsignacionStructContext)

	// EnterWhileDcl is called when entering the whileDcl production.
	EnterWhileDcl(c *WhileDclContext)

	// EnterMultdivmod is called when entering the multdivmod production.
	EnterMultdivmod(c *MultdivmodContext)

	// EnterCasteo_paratipo is called when entering the casteo_paratipo production.
	EnterCasteo_paratipo(c *Casteo_paratipoContext)

	// EnterIncredecr is called when entering the incredecr production.
	EnterIncredecr(c *IncredecrContext)

	// EnterOPERADORESLOGICOS is called when entering the OPERADORESLOGICOS production.
	EnterOPERADORESLOGICOS(c *OPERADORESLOGICOSContext)

	// EnterStructInstanceCreation is called when entering the structInstanceCreation production.
	EnterStructInstanceCreation(c *StructInstanceCreationContext)

	// EnterValorexpr is called when entering the valorexpr production.
	EnterValorexpr(c *ValorexprContext)

	// EnterIgualdad is called when entering the igualdad production.
	EnterIgualdad(c *IgualdadContext)

	// EnterLlamadaFuncionExpr is called when entering the llamadaFuncionExpr production.
	EnterLlamadaFuncionExpr(c *LlamadaFuncionExprContext)

	// EnterExpdotexp is called when entering the expdotexp production.
	EnterExpdotexp(c *ExpdotexpContext)

	// EnterStructAttrAssign is called when entering the structAttrAssign production.
	EnterStructAttrAssign(c *StructAttrAssignContext)

	// EnterRelacionales is called when entering the relacionales production.
	EnterRelacionales(c *RelacionalesContext)

	// EnterCasteo_paratipo_slice is called when entering the casteo_paratipo_slice production.
	EnterCasteo_paratipo_slice(c *Casteo_paratipo_sliceContext)

	// EnterCorchetesexpre is called when entering the corchetesexpre production.
	EnterCorchetesexpre(c *CorchetesexpreContext)

	// EnterUnario is called when entering the unario production.
	EnterUnario(c *UnarioContext)

	// EnterParentesisexpre is called when entering the parentesisexpre production.
	EnterParentesisexpre(c *ParentesisexpreContext)

	// EnterIMCPLICIT is called when entering the IMCPLICIT production.
	EnterIMCPLICIT(c *IMCPLICITContext)

	// EnterSumres is called when entering the sumres production.
	EnterSumres(c *SumresContext)

	// EnterPARAPRINTSLICE is called when entering the PARAPRINTSLICE production.
	EnterPARAPRINTSLICE(c *PARAPRINTSLICEContext)

	// EnterAsignacionLUEGO is called when entering the asignacionLUEGO production.
	EnterAsignacionLUEGO(c *AsignacionLUEGOContext)

	// EnterId is called when entering the id production.
	EnterId(c *IdContext)

	// EnterExpdotexp1 is called when entering the expdotexp1 production.
	EnterExpdotexp1(c *Expdotexp1Context)

	// EnterParametros is called when entering the parametros production.
	EnterParametros(c *ParametrosContext)

	// EnterValores is called when entering the valores production.
	EnterValores(c *ValoresContext)

	// EnterValorEntero is called when entering the valorEntero production.
	EnterValorEntero(c *ValorEnteroContext)

	// EnterValorDecimal is called when entering the valorDecimal production.
	EnterValorDecimal(c *ValorDecimalContext)

	// EnterValorCadena is called when entering the valorCadena production.
	EnterValorCadena(c *ValorCadenaContext)

	// EnterValorBooleano is called when entering the valorBooleano production.
	EnterValorBooleano(c *ValorBooleanoContext)

	// EnterValorCaracter is called when entering the valorCaracter production.
	EnterValorCaracter(c *ValorCaracterContext)

	// EnterListaExpresiones is called when entering the listaExpresiones production.
	EnterListaExpresiones(c *ListaExpresionesContext)

	// EnterIncremento is called when entering the incremento production.
	EnterIncremento(c *IncrementoContext)

	// EnterDecremento is called when entering the decremento production.
	EnterDecremento(c *DecrementoContext)

	// ExitPrograma is called when exiting the programa production.
	ExitPrograma(c *ProgramaContext)

	// ExitFuncMain is called when exiting the funcMain production.
	ExitFuncMain(c *FuncMainContext)

	// ExitFuncDcl is called when exiting the funcDcl production.
	ExitFuncDcl(c *FuncDclContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitDeclaraciones is called when exiting the declaraciones production.
	ExitDeclaraciones(c *DeclaracionesContext)

	// ExitVariableDeclaration is called when exiting the variableDeclaration production.
	ExitVariableDeclaration(c *VariableDeclarationContext)

	// ExitSliceEmptyDeclaration is called when exiting the sliceEmptyDeclaration production.
	ExitSliceEmptyDeclaration(c *SliceEmptyDeclarationContext)

	// ExitStructDirectInitDeclaration is called when exiting the structDirectInitDeclaration production.
	ExitStructDirectInitDeclaration(c *StructDirectInitDeclarationContext)

	// ExitSliceInitDeclaration is called when exiting the sliceInitDeclaration production.
	ExitSliceInitDeclaration(c *SliceInitDeclarationContext)

	// ExitSliceAssignment is called when exiting the sliceAssignment production.
	ExitSliceAssignment(c *SliceAssignmentContext)

	// ExitVariableDeclarationImmutable is called when exiting the variableDeclarationImmutable production.
	ExitVariableDeclarationImmutable(c *VariableDeclarationImmutableContext)

	// ExitVariableCastDeclaration is called when exiting the variableCastDeclaration production.
	ExitVariableCastDeclaration(c *VariableCastDeclarationContext)

	// ExitSliceAssignmentIndex is called when exiting the sliceAssignmentIndex production.
	ExitSliceAssignmentIndex(c *SliceAssignmentIndexContext)

	// ExitSliceTipo is called when exiting the sliceTipo production.
	ExitSliceTipo(c *SliceTipoContext)

	// ExitSliceInit is called when exiting the sliceInit production.
	ExitSliceInit(c *SliceInitContext)

	// ExitPrintStatement is called when exiting the printStatement production.
	ExitPrintStatement(c *PrintStatementContext)

	// ExitExpresionStatement is called when exiting the expresionStatement production.
	ExitExpresionStatement(c *ExpresionStatementContext)

	// ExitControlStatement is called when exiting the controlStatement production.
	ExitControlStatement(c *ControlStatementContext)

	// ExitTransfersentence is called when exiting the transfersentence production.
	ExitTransfersentence(c *TransfersentenceContext)

	// ExitIf_context is called when exiting the if_context production.
	ExitIf_context(c *If_contextContext)

	// ExitFor_context is called when exiting the for_context production.
	ExitFor_context(c *For_contextContext)

	// ExitSwitch_context is called when exiting the switch_context production.
	ExitSwitch_context(c *Switch_contextContext)

	// ExitWhile_context is called when exiting the while_context production.
	ExitWhile_context(c *While_contextContext)

	// ExitBreakStatement is called when exiting the breakStatement production.
	ExitBreakStatement(c *BreakStatementContext)

	// ExitContinueStatement is called when exiting the continueStatement production.
	ExitContinueStatement(c *ContinueStatementContext)

	// ExitReturnStatement is called when exiting the returnStatement production.
	ExitReturnStatement(c *ReturnStatementContext)

	// ExitIfDcl is called when exiting the ifDcl production.
	ExitIfDcl(c *IfDclContext)

	// ExitElseIfDcl is called when exiting the elseIfDcl production.
	ExitElseIfDcl(c *ElseIfDclContext)

	// ExitElseCondicional is called when exiting the elseCondicional production.
	ExitElseCondicional(c *ElseCondicionalContext)

	// ExitForClasico is called when exiting the forClasico production.
	ExitForClasico(c *ForClasicoContext)

	// ExitForCondicionUnica is called when exiting the forCondicionUnica production.
	ExitForCondicionUnica(c *ForCondicionUnicaContext)

	// ExitForRangeSlice is called when exiting the forRangeSlice production.
	ExitForRangeSlice(c *ForRangeSliceContext)

	// ExitAsignacion is called when exiting the asignacion production.
	ExitAsignacion(c *AsignacionContext)

	// ExitSwitchDcl is called when exiting the switchDcl production.
	ExitSwitchDcl(c *SwitchDclContext)

	// ExitCaseBlock is called when exiting the caseBlock production.
	ExitCaseBlock(c *CaseBlockContext)

	// ExitDefaultBlock is called when exiting the defaultBlock production.
	ExitDefaultBlock(c *DefaultBlockContext)

	// ExitLlamadaFuncion is called when exiting the llamadaFuncion production.
	ExitLlamadaFuncion(c *LlamadaFuncionContext)

	// ExitFuncCall is called when exiting the funcCall production.
	ExitFuncCall(c *FuncCallContext)

	// ExitParametrosFormales is called when exiting the parametrosFormales production.
	ExitParametrosFormales(c *ParametrosFormalesContext)

	// ExitParametro is called when exiting the parametro production.
	ExitParametro(c *ParametroContext)

	// ExitParametrosReales is called when exiting the parametrosReales production.
	ExitParametrosReales(c *ParametrosRealesContext)

	// ExitStructDcl is called when exiting the structDcl production.
	ExitStructDcl(c *StructDclContext)

	// ExitAtributosStruct is called when exiting the atributosStruct production.
	ExitAtributosStruct(c *AtributosStructContext)

	// ExitAtributoPrimitivo is called when exiting the atributoPrimitivo production.
	ExitAtributoPrimitivo(c *AtributoPrimitivoContext)

	// ExitAtributoStructAnidado is called when exiting the atributoStructAnidado production.
	ExitAtributoStructAnidado(c *AtributoStructAnidadoContext)

	// ExitListaAsignaciones is called when exiting the listaAsignaciones production.
	ExitListaAsignaciones(c *ListaAsignacionesContext)

	// ExitAsignacionStruct is called when exiting the asignacionStruct production.
	ExitAsignacionStruct(c *AsignacionStructContext)

	// ExitWhileDcl is called when exiting the whileDcl production.
	ExitWhileDcl(c *WhileDclContext)

	// ExitMultdivmod is called when exiting the multdivmod production.
	ExitMultdivmod(c *MultdivmodContext)

	// ExitCasteo_paratipo is called when exiting the casteo_paratipo production.
	ExitCasteo_paratipo(c *Casteo_paratipoContext)

	// ExitIncredecr is called when exiting the incredecr production.
	ExitIncredecr(c *IncredecrContext)

	// ExitOPERADORESLOGICOS is called when exiting the OPERADORESLOGICOS production.
	ExitOPERADORESLOGICOS(c *OPERADORESLOGICOSContext)

	// ExitStructInstanceCreation is called when exiting the structInstanceCreation production.
	ExitStructInstanceCreation(c *StructInstanceCreationContext)

	// ExitValorexpr is called when exiting the valorexpr production.
	ExitValorexpr(c *ValorexprContext)

	// ExitIgualdad is called when exiting the igualdad production.
	ExitIgualdad(c *IgualdadContext)

	// ExitLlamadaFuncionExpr is called when exiting the llamadaFuncionExpr production.
	ExitLlamadaFuncionExpr(c *LlamadaFuncionExprContext)

	// ExitExpdotexp is called when exiting the expdotexp production.
	ExitExpdotexp(c *ExpdotexpContext)

	// ExitStructAttrAssign is called when exiting the structAttrAssign production.
	ExitStructAttrAssign(c *StructAttrAssignContext)

	// ExitRelacionales is called when exiting the relacionales production.
	ExitRelacionales(c *RelacionalesContext)

	// ExitCasteo_paratipo_slice is called when exiting the casteo_paratipo_slice production.
	ExitCasteo_paratipo_slice(c *Casteo_paratipo_sliceContext)

	// ExitCorchetesexpre is called when exiting the corchetesexpre production.
	ExitCorchetesexpre(c *CorchetesexpreContext)

	// ExitUnario is called when exiting the unario production.
	ExitUnario(c *UnarioContext)

	// ExitParentesisexpre is called when exiting the parentesisexpre production.
	ExitParentesisexpre(c *ParentesisexpreContext)

	// ExitIMCPLICIT is called when exiting the IMCPLICIT production.
	ExitIMCPLICIT(c *IMCPLICITContext)

	// ExitSumres is called when exiting the sumres production.
	ExitSumres(c *SumresContext)

	// ExitPARAPRINTSLICE is called when exiting the PARAPRINTSLICE production.
	ExitPARAPRINTSLICE(c *PARAPRINTSLICEContext)

	// ExitAsignacionLUEGO is called when exiting the asignacionLUEGO production.
	ExitAsignacionLUEGO(c *AsignacionLUEGOContext)

	// ExitId is called when exiting the id production.
	ExitId(c *IdContext)

	// ExitExpdotexp1 is called when exiting the expdotexp1 production.
	ExitExpdotexp1(c *Expdotexp1Context)

	// ExitParametros is called when exiting the parametros production.
	ExitParametros(c *ParametrosContext)

	// ExitValores is called when exiting the valores production.
	ExitValores(c *ValoresContext)

	// ExitValorEntero is called when exiting the valorEntero production.
	ExitValorEntero(c *ValorEnteroContext)

	// ExitValorDecimal is called when exiting the valorDecimal production.
	ExitValorDecimal(c *ValorDecimalContext)

	// ExitValorCadena is called when exiting the valorCadena production.
	ExitValorCadena(c *ValorCadenaContext)

	// ExitValorBooleano is called when exiting the valorBooleano production.
	ExitValorBooleano(c *ValorBooleanoContext)

	// ExitValorCaracter is called when exiting the valorCaracter production.
	ExitValorCaracter(c *ValorCaracterContext)

	// ExitListaExpresiones is called when exiting the listaExpresiones production.
	ExitListaExpresiones(c *ListaExpresionesContext)

	// ExitIncremento is called when exiting the incremento production.
	ExitIncremento(c *IncrementoContext)

	// ExitDecremento is called when exiting the decremento production.
	ExitDecremento(c *DecrementoContext)
}
