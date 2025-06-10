// Code generated from parser/Vlang.g4 by ANTLR 4.13.2. DO NOT EDIT.

package compiler // Vlang
import "github.com/antlr4-go/antlr/v4"

// BaseVlangListener is a complete listener for a parse tree produced by VlangParser.
type BaseVlangListener struct{}

var _ VlangListener = &BaseVlangListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseVlangListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseVlangListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseVlangListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseVlangListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterPrograma is called when production programa is entered.
func (s *BaseVlangListener) EnterPrograma(ctx *ProgramaContext) {}

// ExitPrograma is called when production programa is exited.
func (s *BaseVlangListener) ExitPrograma(ctx *ProgramaContext) {}

// EnterFuncMain is called when production funcMain is entered.
func (s *BaseVlangListener) EnterFuncMain(ctx *FuncMainContext) {}

// ExitFuncMain is called when production funcMain is exited.
func (s *BaseVlangListener) ExitFuncMain(ctx *FuncMainContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseVlangListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseVlangListener) ExitBlock(ctx *BlockContext) {}

// EnterDeclaraciones is called when production declaraciones is entered.
func (s *BaseVlangListener) EnterDeclaraciones(ctx *DeclaracionesContext) {}

// ExitDeclaraciones is called when production declaraciones is exited.
func (s *BaseVlangListener) ExitDeclaraciones(ctx *DeclaracionesContext) {}

// EnterVariableDeclaration is called when production variableDeclaration is entered.
func (s *BaseVlangListener) EnterVariableDeclaration(ctx *VariableDeclarationContext) {}

// ExitVariableDeclaration is called when production variableDeclaration is exited.
func (s *BaseVlangListener) ExitVariableDeclaration(ctx *VariableDeclarationContext) {}

// EnterVariableDeclarationImmutable is called when production variableDeclarationImmutable is entered.
func (s *BaseVlangListener) EnterVariableDeclarationImmutable(ctx *VariableDeclarationImmutableContext) {
}

// ExitVariableDeclarationImmutable is called when production variableDeclarationImmutable is exited.
func (s *BaseVlangListener) ExitVariableDeclarationImmutable(ctx *VariableDeclarationImmutableContext) {
}

// EnterVariableCastDeclaration is called when production variableCastDeclaration is entered.
func (s *BaseVlangListener) EnterVariableCastDeclaration(ctx *VariableCastDeclarationContext) {}

// ExitVariableCastDeclaration is called when production variableCastDeclaration is exited.
func (s *BaseVlangListener) ExitVariableCastDeclaration(ctx *VariableCastDeclarationContext) {}

// EnterTipoSlice is called when production tipoSlice is entered.
func (s *BaseVlangListener) EnterTipoSlice(ctx *TipoSliceContext) {}

// ExitTipoSlice is called when production tipoSlice is exited.
func (s *BaseVlangListener) ExitTipoSlice(ctx *TipoSliceContext) {}

// EnterPrintStatement is called when production printStatement is entered.
func (s *BaseVlangListener) EnterPrintStatement(ctx *PrintStatementContext) {}

// ExitPrintStatement is called when production printStatement is exited.
func (s *BaseVlangListener) ExitPrintStatement(ctx *PrintStatementContext) {}

// EnterExpresionStatement is called when production expresionStatement is entered.
func (s *BaseVlangListener) EnterExpresionStatement(ctx *ExpresionStatementContext) {}

// ExitExpresionStatement is called when production expresionStatement is exited.
func (s *BaseVlangListener) ExitExpresionStatement(ctx *ExpresionStatementContext) {}

// EnterControlStatement is called when production controlStatement is entered.
func (s *BaseVlangListener) EnterControlStatement(ctx *ControlStatementContext) {}

// ExitControlStatement is called when production controlStatement is exited.
func (s *BaseVlangListener) ExitControlStatement(ctx *ControlStatementContext) {}

// EnterTransfersentence is called when production transfersentence is entered.
func (s *BaseVlangListener) EnterTransfersentence(ctx *TransfersentenceContext) {}

// ExitTransfersentence is called when production transfersentence is exited.
func (s *BaseVlangListener) ExitTransfersentence(ctx *TransfersentenceContext) {}

// EnterIf_context is called when production if_context is entered.
func (s *BaseVlangListener) EnterIf_context(ctx *If_contextContext) {}

// ExitIf_context is called when production if_context is exited.
func (s *BaseVlangListener) ExitIf_context(ctx *If_contextContext) {}

// EnterFor_context is called when production for_context is entered.
func (s *BaseVlangListener) EnterFor_context(ctx *For_contextContext) {}

// ExitFor_context is called when production for_context is exited.
func (s *BaseVlangListener) ExitFor_context(ctx *For_contextContext) {}

// EnterSwitch_context is called when production switch_context is entered.
func (s *BaseVlangListener) EnterSwitch_context(ctx *Switch_contextContext) {}

// ExitSwitch_context is called when production switch_context is exited.
func (s *BaseVlangListener) ExitSwitch_context(ctx *Switch_contextContext) {}

// EnterWhile_context is called when production while_context is entered.
func (s *BaseVlangListener) EnterWhile_context(ctx *While_contextContext) {}

// ExitWhile_context is called when production while_context is exited.
func (s *BaseVlangListener) ExitWhile_context(ctx *While_contextContext) {}

// EnterBreakStatement is called when production breakStatement is entered.
func (s *BaseVlangListener) EnterBreakStatement(ctx *BreakStatementContext) {}

// ExitBreakStatement is called when production breakStatement is exited.
func (s *BaseVlangListener) ExitBreakStatement(ctx *BreakStatementContext) {}

// EnterContinueStatement is called when production continueStatement is entered.
func (s *BaseVlangListener) EnterContinueStatement(ctx *ContinueStatementContext) {}

// ExitContinueStatement is called when production continueStatement is exited.
func (s *BaseVlangListener) ExitContinueStatement(ctx *ContinueStatementContext) {}

// EnterReturnStatement is called when production returnStatement is entered.
func (s *BaseVlangListener) EnterReturnStatement(ctx *ReturnStatementContext) {}

// ExitReturnStatement is called when production returnStatement is exited.
func (s *BaseVlangListener) ExitReturnStatement(ctx *ReturnStatementContext) {}

// EnterIfDcl is called when production ifDcl is entered.
func (s *BaseVlangListener) EnterIfDcl(ctx *IfDclContext) {}

// ExitIfDcl is called when production ifDcl is exited.
func (s *BaseVlangListener) ExitIfDcl(ctx *IfDclContext) {}

// EnterElseIfDcl is called when production elseIfDcl is entered.
func (s *BaseVlangListener) EnterElseIfDcl(ctx *ElseIfDclContext) {}

// ExitElseIfDcl is called when production elseIfDcl is exited.
func (s *BaseVlangListener) ExitElseIfDcl(ctx *ElseIfDclContext) {}

// EnterElseCondicional is called when production elseCondicional is entered.
func (s *BaseVlangListener) EnterElseCondicional(ctx *ElseCondicionalContext) {}

// ExitElseCondicional is called when production elseCondicional is exited.
func (s *BaseVlangListener) ExitElseCondicional(ctx *ElseCondicionalContext) {}

// EnterForClasico is called when production forClasico is entered.
func (s *BaseVlangListener) EnterForClasico(ctx *ForClasicoContext) {}

// ExitForClasico is called when production forClasico is exited.
func (s *BaseVlangListener) ExitForClasico(ctx *ForClasicoContext) {}

// EnterForCondicionUnica is called when production forCondicionUnica is entered.
func (s *BaseVlangListener) EnterForCondicionUnica(ctx *ForCondicionUnicaContext) {}

// ExitForCondicionUnica is called when production forCondicionUnica is exited.
func (s *BaseVlangListener) ExitForCondicionUnica(ctx *ForCondicionUnicaContext) {}

// EnterAsignacion is called when production asignacion is entered.
func (s *BaseVlangListener) EnterAsignacion(ctx *AsignacionContext) {}

// ExitAsignacion is called when production asignacion is exited.
func (s *BaseVlangListener) ExitAsignacion(ctx *AsignacionContext) {}

// EnterSwitchDcl is called when production switchDcl is entered.
func (s *BaseVlangListener) EnterSwitchDcl(ctx *SwitchDclContext) {}

// ExitSwitchDcl is called when production switchDcl is exited.
func (s *BaseVlangListener) ExitSwitchDcl(ctx *SwitchDclContext) {}

// EnterCaseBlock is called when production caseBlock is entered.
func (s *BaseVlangListener) EnterCaseBlock(ctx *CaseBlockContext) {}

// ExitCaseBlock is called when production caseBlock is exited.
func (s *BaseVlangListener) ExitCaseBlock(ctx *CaseBlockContext) {}

// EnterDefaultBlock is called when production defaultBlock is entered.
func (s *BaseVlangListener) EnterDefaultBlock(ctx *DefaultBlockContext) {}

// ExitDefaultBlock is called when production defaultBlock is exited.
func (s *BaseVlangListener) ExitDefaultBlock(ctx *DefaultBlockContext) {}

// EnterLlamadaFuncion is called when production llamadaFuncion is entered.
func (s *BaseVlangListener) EnterLlamadaFuncion(ctx *LlamadaFuncionContext) {}

// ExitLlamadaFuncion is called when production llamadaFuncion is exited.
func (s *BaseVlangListener) ExitLlamadaFuncion(ctx *LlamadaFuncionContext) {}

// EnterWhileDcl is called when production whileDcl is entered.
func (s *BaseVlangListener) EnterWhileDcl(ctx *WhileDclContext) {}

// ExitWhileDcl is called when production whileDcl is exited.
func (s *BaseVlangListener) ExitWhileDcl(ctx *WhileDclContext) {}

// EnterMultdivmod is called when production multdivmod is entered.
func (s *BaseVlangListener) EnterMultdivmod(ctx *MultdivmodContext) {}

// ExitMultdivmod is called when production multdivmod is exited.
func (s *BaseVlangListener) ExitMultdivmod(ctx *MultdivmodContext) {}

// EnterIncredecr is called when production incredecr is entered.
func (s *BaseVlangListener) EnterIncredecr(ctx *IncredecrContext) {}

// ExitIncredecr is called when production incredecr is exited.
func (s *BaseVlangListener) ExitIncredecr(ctx *IncredecrContext) {}

// EnterOPERADORESLOGICOS is called when production OPERADORESLOGICOS is entered.
func (s *BaseVlangListener) EnterOPERADORESLOGICOS(ctx *OPERADORESLOGICOSContext) {}

// ExitOPERADORESLOGICOS is called when production OPERADORESLOGICOS is exited.
func (s *BaseVlangListener) ExitOPERADORESLOGICOS(ctx *OPERADORESLOGICOSContext) {}

// EnterValorexpr is called when production valorexpr is entered.
func (s *BaseVlangListener) EnterValorexpr(ctx *ValorexprContext) {}

// ExitValorexpr is called when production valorexpr is exited.
func (s *BaseVlangListener) ExitValorexpr(ctx *ValorexprContext) {}

// EnterIgualdad is called when production igualdad is entered.
func (s *BaseVlangListener) EnterIgualdad(ctx *IgualdadContext) {}

// ExitIgualdad is called when production igualdad is exited.
func (s *BaseVlangListener) ExitIgualdad(ctx *IgualdadContext) {}

// EnterSliceCreacionv is called when production sliceCreacionv is entered.
func (s *BaseVlangListener) EnterSliceCreacionv(ctx *SliceCreacionvContext) {}

// ExitSliceCreacionv is called when production sliceCreacionv is exited.
func (s *BaseVlangListener) ExitSliceCreacionv(ctx *SliceCreacionvContext) {}

// EnterLlamadaFuncionExpr is called when production llamadaFuncionExpr is entered.
func (s *BaseVlangListener) EnterLlamadaFuncionExpr(ctx *LlamadaFuncionExprContext) {}

// ExitLlamadaFuncionExpr is called when production llamadaFuncionExpr is exited.
func (s *BaseVlangListener) ExitLlamadaFuncionExpr(ctx *LlamadaFuncionExprContext) {}

// EnterExpdotexp is called when production expdotexp is entered.
func (s *BaseVlangListener) EnterExpdotexp(ctx *ExpdotexpContext) {}

// ExitExpdotexp is called when production expdotexp is exited.
func (s *BaseVlangListener) ExitExpdotexp(ctx *ExpdotexpContext) {}

// EnterRelacionales is called when production relacionales is entered.
func (s *BaseVlangListener) EnterRelacionales(ctx *RelacionalesContext) {}

// ExitRelacionales is called when production relacionales is exited.
func (s *BaseVlangListener) ExitRelacionales(ctx *RelacionalesContext) {}

// EnterCorchetesexpre is called when production corchetesexpre is entered.
func (s *BaseVlangListener) EnterCorchetesexpre(ctx *CorchetesexpreContext) {}

// ExitCorchetesexpre is called when production corchetesexpre is exited.
func (s *BaseVlangListener) ExitCorchetesexpre(ctx *CorchetesexpreContext) {}

// EnterUnario is called when production unario is entered.
func (s *BaseVlangListener) EnterUnario(ctx *UnarioContext) {}

// ExitUnario is called when production unario is exited.
func (s *BaseVlangListener) ExitUnario(ctx *UnarioContext) {}

// EnterParentesisexpre is called when production parentesisexpre is entered.
func (s *BaseVlangListener) EnterParentesisexpre(ctx *ParentesisexpreContext) {}

// ExitParentesisexpre is called when production parentesisexpre is exited.
func (s *BaseVlangListener) ExitParentesisexpre(ctx *ParentesisexpreContext) {}

// EnterIMCPLICIT is called when production IMCPLICIT is entered.
func (s *BaseVlangListener) EnterIMCPLICIT(ctx *IMCPLICITContext) {}

// ExitIMCPLICIT is called when production IMCPLICIT is exited.
func (s *BaseVlangListener) ExitIMCPLICIT(ctx *IMCPLICITContext) {}

// EnterSumres is called when production sumres is entered.
func (s *BaseVlangListener) EnterSumres(ctx *SumresContext) {}

// ExitSumres is called when production sumres is exited.
func (s *BaseVlangListener) ExitSumres(ctx *SumresContext) {}

// EnterAsignacionLUEGO is called when production asignacionLUEGO is entered.
func (s *BaseVlangListener) EnterAsignacionLUEGO(ctx *AsignacionLUEGOContext) {}

// ExitAsignacionLUEGO is called when production asignacionLUEGO is exited.
func (s *BaseVlangListener) ExitAsignacionLUEGO(ctx *AsignacionLUEGOContext) {}

// EnterId is called when production id is entered.
func (s *BaseVlangListener) EnterId(ctx *IdContext) {}

// ExitId is called when production id is exited.
func (s *BaseVlangListener) ExitId(ctx *IdContext) {}

// EnterExpdotexp1 is called when production expdotexp1 is entered.
func (s *BaseVlangListener) EnterExpdotexp1(ctx *Expdotexp1Context) {}

// ExitExpdotexp1 is called when production expdotexp1 is exited.
func (s *BaseVlangListener) ExitExpdotexp1(ctx *Expdotexp1Context) {}

// EnterParametros is called when production parametros is entered.
func (s *BaseVlangListener) EnterParametros(ctx *ParametrosContext) {}

// ExitParametros is called when production parametros is exited.
func (s *BaseVlangListener) ExitParametros(ctx *ParametrosContext) {}

// EnterValores is called when production valores is entered.
func (s *BaseVlangListener) EnterValores(ctx *ValoresContext) {}

// ExitValores is called when production valores is exited.
func (s *BaseVlangListener) ExitValores(ctx *ValoresContext) {}

// EnterValorEntero is called when production valorEntero is entered.
func (s *BaseVlangListener) EnterValorEntero(ctx *ValorEnteroContext) {}

// ExitValorEntero is called when production valorEntero is exited.
func (s *BaseVlangListener) ExitValorEntero(ctx *ValorEnteroContext) {}

// EnterValorDecimal is called when production valorDecimal is entered.
func (s *BaseVlangListener) EnterValorDecimal(ctx *ValorDecimalContext) {}

// ExitValorDecimal is called when production valorDecimal is exited.
func (s *BaseVlangListener) ExitValorDecimal(ctx *ValorDecimalContext) {}

// EnterValorCadena is called when production valorCadena is entered.
func (s *BaseVlangListener) EnterValorCadena(ctx *ValorCadenaContext) {}

// ExitValorCadena is called when production valorCadena is exited.
func (s *BaseVlangListener) ExitValorCadena(ctx *ValorCadenaContext) {}

// EnterValorBooleano is called when production valorBooleano is entered.
func (s *BaseVlangListener) EnterValorBooleano(ctx *ValorBooleanoContext) {}

// ExitValorBooleano is called when production valorBooleano is exited.
func (s *BaseVlangListener) ExitValorBooleano(ctx *ValorBooleanoContext) {}

// EnterValorCaracter is called when production valorCaracter is entered.
func (s *BaseVlangListener) EnterValorCaracter(ctx *ValorCaracterContext) {}

// ExitValorCaracter is called when production valorCaracter is exited.
func (s *BaseVlangListener) ExitValorCaracter(ctx *ValorCaracterContext) {}

// EnterListaExpresiones is called when production listaExpresiones is entered.
func (s *BaseVlangListener) EnterListaExpresiones(ctx *ListaExpresionesContext) {}

// ExitListaExpresiones is called when production listaExpresiones is exited.
func (s *BaseVlangListener) ExitListaExpresiones(ctx *ListaExpresionesContext) {}

// EnterIncremento is called when production incremento is entered.
func (s *BaseVlangListener) EnterIncremento(ctx *IncrementoContext) {}

// ExitIncremento is called when production incremento is exited.
func (s *BaseVlangListener) ExitIncremento(ctx *IncrementoContext) {}

// EnterDecremento is called when production decremento is entered.
func (s *BaseVlangListener) EnterDecremento(ctx *DecrementoContext) {}

// ExitDecremento is called when production decremento is exited.
func (s *BaseVlangListener) ExitDecremento(ctx *DecrementoContext) {}
