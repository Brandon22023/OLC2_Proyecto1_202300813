// Generated from /home/brandon/Escritorio/OLC2_Proyecto1_202300813/Backend/compiler/parser/Vlang.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link VlangParser}.
 */
public interface VlangListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link VlangParser#programa}.
	 * @param ctx the parse tree
	 */
	void enterPrograma(VlangParser.ProgramaContext ctx);
	/**
	 * Exit a parse tree produced by {@link VlangParser#programa}.
	 * @param ctx the parse tree
	 */
	void exitPrograma(VlangParser.ProgramaContext ctx);
	/**
	 * Enter a parse tree produced by {@link VlangParser#funcMain}.
	 * @param ctx the parse tree
	 */
	void enterFuncMain(VlangParser.FuncMainContext ctx);
	/**
	 * Exit a parse tree produced by {@link VlangParser#funcMain}.
	 * @param ctx the parse tree
	 */
	void exitFuncMain(VlangParser.FuncMainContext ctx);
	/**
	 * Enter a parse tree produced by {@link VlangParser#funcDcl}.
	 * @param ctx the parse tree
	 */
	void enterFuncDcl(VlangParser.FuncDclContext ctx);
	/**
	 * Exit a parse tree produced by {@link VlangParser#funcDcl}.
	 * @param ctx the parse tree
	 */
	void exitFuncDcl(VlangParser.FuncDclContext ctx);
	/**
	 * Enter a parse tree produced by {@link VlangParser#block}.
	 * @param ctx the parse tree
	 */
	void enterBlock(VlangParser.BlockContext ctx);
	/**
	 * Exit a parse tree produced by {@link VlangParser#block}.
	 * @param ctx the parse tree
	 */
	void exitBlock(VlangParser.BlockContext ctx);
	/**
	 * Enter a parse tree produced by {@link VlangParser#declaraciones}.
	 * @param ctx the parse tree
	 */
	void enterDeclaraciones(VlangParser.DeclaracionesContext ctx);
	/**
	 * Exit a parse tree produced by {@link VlangParser#declaraciones}.
	 * @param ctx the parse tree
	 */
	void exitDeclaraciones(VlangParser.DeclaracionesContext ctx);
	/**
	 * Enter a parse tree produced by the {@code variableDeclaration}
	 * labeled alternative in {@link VlangParser#varDcl}.
	 * @param ctx the parse tree
	 */
	void enterVariableDeclaration(VlangParser.VariableDeclarationContext ctx);
	/**
	 * Exit a parse tree produced by the {@code variableDeclaration}
	 * labeled alternative in {@link VlangParser#varDcl}.
	 * @param ctx the parse tree
	 */
	void exitVariableDeclaration(VlangParser.VariableDeclarationContext ctx);
	/**
	 * Enter a parse tree produced by the {@code sliceEmptyDeclaration}
	 * labeled alternative in {@link VlangParser#varDcl}.
	 * @param ctx the parse tree
	 */
	void enterSliceEmptyDeclaration(VlangParser.SliceEmptyDeclarationContext ctx);
	/**
	 * Exit a parse tree produced by the {@code sliceEmptyDeclaration}
	 * labeled alternative in {@link VlangParser#varDcl}.
	 * @param ctx the parse tree
	 */
	void exitSliceEmptyDeclaration(VlangParser.SliceEmptyDeclarationContext ctx);
	/**
	 * Enter a parse tree produced by the {@code structDirectInitDeclaration}
	 * labeled alternative in {@link VlangParser#varDcl}.
	 * @param ctx the parse tree
	 */
	void enterStructDirectInitDeclaration(VlangParser.StructDirectInitDeclarationContext ctx);
	/**
	 * Exit a parse tree produced by the {@code structDirectInitDeclaration}
	 * labeled alternative in {@link VlangParser#varDcl}.
	 * @param ctx the parse tree
	 */
	void exitStructDirectInitDeclaration(VlangParser.StructDirectInitDeclarationContext ctx);
	/**
	 * Enter a parse tree produced by the {@code sliceInitDeclaration}
	 * labeled alternative in {@link VlangParser#varDcl}.
	 * @param ctx the parse tree
	 */
	void enterSliceInitDeclaration(VlangParser.SliceInitDeclarationContext ctx);
	/**
	 * Exit a parse tree produced by the {@code sliceInitDeclaration}
	 * labeled alternative in {@link VlangParser#varDcl}.
	 * @param ctx the parse tree
	 */
	void exitSliceInitDeclaration(VlangParser.SliceInitDeclarationContext ctx);
	/**
	 * Enter a parse tree produced by the {@code sliceAssignment}
	 * labeled alternative in {@link VlangParser#varDcl}.
	 * @param ctx the parse tree
	 */
	void enterSliceAssignment(VlangParser.SliceAssignmentContext ctx);
	/**
	 * Exit a parse tree produced by the {@code sliceAssignment}
	 * labeled alternative in {@link VlangParser#varDcl}.
	 * @param ctx the parse tree
	 */
	void exitSliceAssignment(VlangParser.SliceAssignmentContext ctx);
	/**
	 * Enter a parse tree produced by the {@code variableDeclarationImmutable}
	 * labeled alternative in {@link VlangParser#varDcl}.
	 * @param ctx the parse tree
	 */
	void enterVariableDeclarationImmutable(VlangParser.VariableDeclarationImmutableContext ctx);
	/**
	 * Exit a parse tree produced by the {@code variableDeclarationImmutable}
	 * labeled alternative in {@link VlangParser#varDcl}.
	 * @param ctx the parse tree
	 */
	void exitVariableDeclarationImmutable(VlangParser.VariableDeclarationImmutableContext ctx);
	/**
	 * Enter a parse tree produced by the {@code variableCastDeclaration}
	 * labeled alternative in {@link VlangParser#varDcl}.
	 * @param ctx the parse tree
	 */
	void enterVariableCastDeclaration(VlangParser.VariableCastDeclarationContext ctx);
	/**
	 * Exit a parse tree produced by the {@code variableCastDeclaration}
	 * labeled alternative in {@link VlangParser#varDcl}.
	 * @param ctx the parse tree
	 */
	void exitVariableCastDeclaration(VlangParser.VariableCastDeclarationContext ctx);
	/**
	 * Enter a parse tree produced by the {@code sliceAssignmentIndex}
	 * labeled alternative in {@link VlangParser#varDcl}.
	 * @param ctx the parse tree
	 */
	void enterSliceAssignmentIndex(VlangParser.SliceAssignmentIndexContext ctx);
	/**
	 * Exit a parse tree produced by the {@code sliceAssignmentIndex}
	 * labeled alternative in {@link VlangParser#varDcl}.
	 * @param ctx the parse tree
	 */
	void exitSliceAssignmentIndex(VlangParser.SliceAssignmentIndexContext ctx);
	/**
	 * Enter a parse tree produced by {@link VlangParser#sliceTipo}.
	 * @param ctx the parse tree
	 */
	void enterSliceTipo(VlangParser.SliceTipoContext ctx);
	/**
	 * Exit a parse tree produced by {@link VlangParser#sliceTipo}.
	 * @param ctx the parse tree
	 */
	void exitSliceTipo(VlangParser.SliceTipoContext ctx);
	/**
	 * Enter a parse tree produced by {@link VlangParser#sliceInit}.
	 * @param ctx the parse tree
	 */
	void enterSliceInit(VlangParser.SliceInitContext ctx);
	/**
	 * Exit a parse tree produced by {@link VlangParser#sliceInit}.
	 * @param ctx the parse tree
	 */
	void exitSliceInit(VlangParser.SliceInitContext ctx);
	/**
	 * Enter a parse tree produced by the {@code printStatement}
	 * labeled alternative in {@link VlangParser#stmt}.
	 * @param ctx the parse tree
	 */
	void enterPrintStatement(VlangParser.PrintStatementContext ctx);
	/**
	 * Exit a parse tree produced by the {@code printStatement}
	 * labeled alternative in {@link VlangParser#stmt}.
	 * @param ctx the parse tree
	 */
	void exitPrintStatement(VlangParser.PrintStatementContext ctx);
	/**
	 * Enter a parse tree produced by the {@code expresionStatement}
	 * labeled alternative in {@link VlangParser#stmt}.
	 * @param ctx the parse tree
	 */
	void enterExpresionStatement(VlangParser.ExpresionStatementContext ctx);
	/**
	 * Exit a parse tree produced by the {@code expresionStatement}
	 * labeled alternative in {@link VlangParser#stmt}.
	 * @param ctx the parse tree
	 */
	void exitExpresionStatement(VlangParser.ExpresionStatementContext ctx);
	/**
	 * Enter a parse tree produced by the {@code controlStatement}
	 * labeled alternative in {@link VlangParser#stmt}.
	 * @param ctx the parse tree
	 */
	void enterControlStatement(VlangParser.ControlStatementContext ctx);
	/**
	 * Exit a parse tree produced by the {@code controlStatement}
	 * labeled alternative in {@link VlangParser#stmt}.
	 * @param ctx the parse tree
	 */
	void exitControlStatement(VlangParser.ControlStatementContext ctx);
	/**
	 * Enter a parse tree produced by the {@code transfersentence}
	 * labeled alternative in {@link VlangParser#stmt}.
	 * @param ctx the parse tree
	 */
	void enterTransfersentence(VlangParser.TransfersentenceContext ctx);
	/**
	 * Exit a parse tree produced by the {@code transfersentence}
	 * labeled alternative in {@link VlangParser#stmt}.
	 * @param ctx the parse tree
	 */
	void exitTransfersentence(VlangParser.TransfersentenceContext ctx);
	/**
	 * Enter a parse tree produced by the {@code if_context}
	 * labeled alternative in {@link VlangParser#sentencias_control}.
	 * @param ctx the parse tree
	 */
	void enterIf_context(VlangParser.If_contextContext ctx);
	/**
	 * Exit a parse tree produced by the {@code if_context}
	 * labeled alternative in {@link VlangParser#sentencias_control}.
	 * @param ctx the parse tree
	 */
	void exitIf_context(VlangParser.If_contextContext ctx);
	/**
	 * Enter a parse tree produced by the {@code for_context}
	 * labeled alternative in {@link VlangParser#sentencias_control}.
	 * @param ctx the parse tree
	 */
	void enterFor_context(VlangParser.For_contextContext ctx);
	/**
	 * Exit a parse tree produced by the {@code for_context}
	 * labeled alternative in {@link VlangParser#sentencias_control}.
	 * @param ctx the parse tree
	 */
	void exitFor_context(VlangParser.For_contextContext ctx);
	/**
	 * Enter a parse tree produced by the {@code switch_context}
	 * labeled alternative in {@link VlangParser#sentencias_control}.
	 * @param ctx the parse tree
	 */
	void enterSwitch_context(VlangParser.Switch_contextContext ctx);
	/**
	 * Exit a parse tree produced by the {@code switch_context}
	 * labeled alternative in {@link VlangParser#sentencias_control}.
	 * @param ctx the parse tree
	 */
	void exitSwitch_context(VlangParser.Switch_contextContext ctx);
	/**
	 * Enter a parse tree produced by the {@code while_context}
	 * labeled alternative in {@link VlangParser#sentencias_control}.
	 * @param ctx the parse tree
	 */
	void enterWhile_context(VlangParser.While_contextContext ctx);
	/**
	 * Exit a parse tree produced by the {@code while_context}
	 * labeled alternative in {@link VlangParser#sentencias_control}.
	 * @param ctx the parse tree
	 */
	void exitWhile_context(VlangParser.While_contextContext ctx);
	/**
	 * Enter a parse tree produced by the {@code breakStatement}
	 * labeled alternative in {@link VlangParser#sentencias_transferencia}.
	 * @param ctx the parse tree
	 */
	void enterBreakStatement(VlangParser.BreakStatementContext ctx);
	/**
	 * Exit a parse tree produced by the {@code breakStatement}
	 * labeled alternative in {@link VlangParser#sentencias_transferencia}.
	 * @param ctx the parse tree
	 */
	void exitBreakStatement(VlangParser.BreakStatementContext ctx);
	/**
	 * Enter a parse tree produced by the {@code continueStatement}
	 * labeled alternative in {@link VlangParser#sentencias_transferencia}.
	 * @param ctx the parse tree
	 */
	void enterContinueStatement(VlangParser.ContinueStatementContext ctx);
	/**
	 * Exit a parse tree produced by the {@code continueStatement}
	 * labeled alternative in {@link VlangParser#sentencias_transferencia}.
	 * @param ctx the parse tree
	 */
	void exitContinueStatement(VlangParser.ContinueStatementContext ctx);
	/**
	 * Enter a parse tree produced by the {@code returnStatement}
	 * labeled alternative in {@link VlangParser#sentencias_transferencia}.
	 * @param ctx the parse tree
	 */
	void enterReturnStatement(VlangParser.ReturnStatementContext ctx);
	/**
	 * Exit a parse tree produced by the {@code returnStatement}
	 * labeled alternative in {@link VlangParser#sentencias_transferencia}.
	 * @param ctx the parse tree
	 */
	void exitReturnStatement(VlangParser.ReturnStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link VlangParser#ifDcl}.
	 * @param ctx the parse tree
	 */
	void enterIfDcl(VlangParser.IfDclContext ctx);
	/**
	 * Exit a parse tree produced by {@link VlangParser#ifDcl}.
	 * @param ctx the parse tree
	 */
	void exitIfDcl(VlangParser.IfDclContext ctx);
	/**
	 * Enter a parse tree produced by {@link VlangParser#elseIfDcl}.
	 * @param ctx the parse tree
	 */
	void enterElseIfDcl(VlangParser.ElseIfDclContext ctx);
	/**
	 * Exit a parse tree produced by {@link VlangParser#elseIfDcl}.
	 * @param ctx the parse tree
	 */
	void exitElseIfDcl(VlangParser.ElseIfDclContext ctx);
	/**
	 * Enter a parse tree produced by {@link VlangParser#elseCondicional}.
	 * @param ctx the parse tree
	 */
	void enterElseCondicional(VlangParser.ElseCondicionalContext ctx);
	/**
	 * Exit a parse tree produced by {@link VlangParser#elseCondicional}.
	 * @param ctx the parse tree
	 */
	void exitElseCondicional(VlangParser.ElseCondicionalContext ctx);
	/**
	 * Enter a parse tree produced by the {@code forClasico}
	 * labeled alternative in {@link VlangParser#forDcl}.
	 * @param ctx the parse tree
	 */
	void enterForClasico(VlangParser.ForClasicoContext ctx);
	/**
	 * Exit a parse tree produced by the {@code forClasico}
	 * labeled alternative in {@link VlangParser#forDcl}.
	 * @param ctx the parse tree
	 */
	void exitForClasico(VlangParser.ForClasicoContext ctx);
	/**
	 * Enter a parse tree produced by the {@code forCondicionUnica}
	 * labeled alternative in {@link VlangParser#forDcl}.
	 * @param ctx the parse tree
	 */
	void enterForCondicionUnica(VlangParser.ForCondicionUnicaContext ctx);
	/**
	 * Exit a parse tree produced by the {@code forCondicionUnica}
	 * labeled alternative in {@link VlangParser#forDcl}.
	 * @param ctx the parse tree
	 */
	void exitForCondicionUnica(VlangParser.ForCondicionUnicaContext ctx);
	/**
	 * Enter a parse tree produced by the {@code forRangeSlice}
	 * labeled alternative in {@link VlangParser#forDcl}.
	 * @param ctx the parse tree
	 */
	void enterForRangeSlice(VlangParser.ForRangeSliceContext ctx);
	/**
	 * Exit a parse tree produced by the {@code forRangeSlice}
	 * labeled alternative in {@link VlangParser#forDcl}.
	 * @param ctx the parse tree
	 */
	void exitForRangeSlice(VlangParser.ForRangeSliceContext ctx);
	/**
	 * Enter a parse tree produced by {@link VlangParser#asignacion}.
	 * @param ctx the parse tree
	 */
	void enterAsignacion(VlangParser.AsignacionContext ctx);
	/**
	 * Exit a parse tree produced by {@link VlangParser#asignacion}.
	 * @param ctx the parse tree
	 */
	void exitAsignacion(VlangParser.AsignacionContext ctx);
	/**
	 * Enter a parse tree produced by {@link VlangParser#switchDcl}.
	 * @param ctx the parse tree
	 */
	void enterSwitchDcl(VlangParser.SwitchDclContext ctx);
	/**
	 * Exit a parse tree produced by {@link VlangParser#switchDcl}.
	 * @param ctx the parse tree
	 */
	void exitSwitchDcl(VlangParser.SwitchDclContext ctx);
	/**
	 * Enter a parse tree produced by {@link VlangParser#caseBlock}.
	 * @param ctx the parse tree
	 */
	void enterCaseBlock(VlangParser.CaseBlockContext ctx);
	/**
	 * Exit a parse tree produced by {@link VlangParser#caseBlock}.
	 * @param ctx the parse tree
	 */
	void exitCaseBlock(VlangParser.CaseBlockContext ctx);
	/**
	 * Enter a parse tree produced by {@link VlangParser#defaultBlock}.
	 * @param ctx the parse tree
	 */
	void enterDefaultBlock(VlangParser.DefaultBlockContext ctx);
	/**
	 * Exit a parse tree produced by {@link VlangParser#defaultBlock}.
	 * @param ctx the parse tree
	 */
	void exitDefaultBlock(VlangParser.DefaultBlockContext ctx);
	/**
	 * Enter a parse tree produced by {@link VlangParser#llamadaFuncion}.
	 * @param ctx the parse tree
	 */
	void enterLlamadaFuncion(VlangParser.LlamadaFuncionContext ctx);
	/**
	 * Exit a parse tree produced by {@link VlangParser#llamadaFuncion}.
	 * @param ctx the parse tree
	 */
	void exitLlamadaFuncion(VlangParser.LlamadaFuncionContext ctx);
	/**
	 * Enter a parse tree produced by {@link VlangParser#funcCall}.
	 * @param ctx the parse tree
	 */
	void enterFuncCall(VlangParser.FuncCallContext ctx);
	/**
	 * Exit a parse tree produced by {@link VlangParser#funcCall}.
	 * @param ctx the parse tree
	 */
	void exitFuncCall(VlangParser.FuncCallContext ctx);
	/**
	 * Enter a parse tree produced by {@link VlangParser#parametrosFormales}.
	 * @param ctx the parse tree
	 */
	void enterParametrosFormales(VlangParser.ParametrosFormalesContext ctx);
	/**
	 * Exit a parse tree produced by {@link VlangParser#parametrosFormales}.
	 * @param ctx the parse tree
	 */
	void exitParametrosFormales(VlangParser.ParametrosFormalesContext ctx);
	/**
	 * Enter a parse tree produced by {@link VlangParser#parametro}.
	 * @param ctx the parse tree
	 */
	void enterParametro(VlangParser.ParametroContext ctx);
	/**
	 * Exit a parse tree produced by {@link VlangParser#parametro}.
	 * @param ctx the parse tree
	 */
	void exitParametro(VlangParser.ParametroContext ctx);
	/**
	 * Enter a parse tree produced by {@link VlangParser#parametrosReales}.
	 * @param ctx the parse tree
	 */
	void enterParametrosReales(VlangParser.ParametrosRealesContext ctx);
	/**
	 * Exit a parse tree produced by {@link VlangParser#parametrosReales}.
	 * @param ctx the parse tree
	 */
	void exitParametrosReales(VlangParser.ParametrosRealesContext ctx);
	/**
	 * Enter a parse tree produced by {@link VlangParser#structDcl}.
	 * @param ctx the parse tree
	 */
	void enterStructDcl(VlangParser.StructDclContext ctx);
	/**
	 * Exit a parse tree produced by {@link VlangParser#structDcl}.
	 * @param ctx the parse tree
	 */
	void exitStructDcl(VlangParser.StructDclContext ctx);
	/**
	 * Enter a parse tree produced by {@link VlangParser#atributosStruct}.
	 * @param ctx the parse tree
	 */
	void enterAtributosStruct(VlangParser.AtributosStructContext ctx);
	/**
	 * Exit a parse tree produced by {@link VlangParser#atributosStruct}.
	 * @param ctx the parse tree
	 */
	void exitAtributosStruct(VlangParser.AtributosStructContext ctx);
	/**
	 * Enter a parse tree produced by the {@code atributoPrimitivo}
	 * labeled alternative in {@link VlangParser#atributoStruct}.
	 * @param ctx the parse tree
	 */
	void enterAtributoPrimitivo(VlangParser.AtributoPrimitivoContext ctx);
	/**
	 * Exit a parse tree produced by the {@code atributoPrimitivo}
	 * labeled alternative in {@link VlangParser#atributoStruct}.
	 * @param ctx the parse tree
	 */
	void exitAtributoPrimitivo(VlangParser.AtributoPrimitivoContext ctx);
	/**
	 * Enter a parse tree produced by the {@code atributoStructAnidado}
	 * labeled alternative in {@link VlangParser#atributoStruct}.
	 * @param ctx the parse tree
	 */
	void enterAtributoStructAnidado(VlangParser.AtributoStructAnidadoContext ctx);
	/**
	 * Exit a parse tree produced by the {@code atributoStructAnidado}
	 * labeled alternative in {@link VlangParser#atributoStruct}.
	 * @param ctx the parse tree
	 */
	void exitAtributoStructAnidado(VlangParser.AtributoStructAnidadoContext ctx);
	/**
	 * Enter a parse tree produced by {@link VlangParser#listaAsignaciones}.
	 * @param ctx the parse tree
	 */
	void enterListaAsignaciones(VlangParser.ListaAsignacionesContext ctx);
	/**
	 * Exit a parse tree produced by {@link VlangParser#listaAsignaciones}.
	 * @param ctx the parse tree
	 */
	void exitListaAsignaciones(VlangParser.ListaAsignacionesContext ctx);
	/**
	 * Enter a parse tree produced by {@link VlangParser#asignacionStruct}.
	 * @param ctx the parse tree
	 */
	void enterAsignacionStruct(VlangParser.AsignacionStructContext ctx);
	/**
	 * Exit a parse tree produced by {@link VlangParser#asignacionStruct}.
	 * @param ctx the parse tree
	 */
	void exitAsignacionStruct(VlangParser.AsignacionStructContext ctx);
	/**
	 * Enter a parse tree produced by {@link VlangParser#whileDcl}.
	 * @param ctx the parse tree
	 */
	void enterWhileDcl(VlangParser.WhileDclContext ctx);
	/**
	 * Exit a parse tree produced by {@link VlangParser#whileDcl}.
	 * @param ctx the parse tree
	 */
	void exitWhileDcl(VlangParser.WhileDclContext ctx);
	/**
	 * Enter a parse tree produced by the {@code multdivmod}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void enterMultdivmod(VlangParser.MultdivmodContext ctx);
	/**
	 * Exit a parse tree produced by the {@code multdivmod}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void exitMultdivmod(VlangParser.MultdivmodContext ctx);
	/**
	 * Enter a parse tree produced by the {@code casteo_paratipo}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void enterCasteo_paratipo(VlangParser.Casteo_paratipoContext ctx);
	/**
	 * Exit a parse tree produced by the {@code casteo_paratipo}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void exitCasteo_paratipo(VlangParser.Casteo_paratipoContext ctx);
	/**
	 * Enter a parse tree produced by the {@code incredecr}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void enterIncredecr(VlangParser.IncredecrContext ctx);
	/**
	 * Exit a parse tree produced by the {@code incredecr}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void exitIncredecr(VlangParser.IncredecrContext ctx);
	/**
	 * Enter a parse tree produced by the {@code OPERADORESLOGICOS}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void enterOPERADORESLOGICOS(VlangParser.OPERADORESLOGICOSContext ctx);
	/**
	 * Exit a parse tree produced by the {@code OPERADORESLOGICOS}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void exitOPERADORESLOGICOS(VlangParser.OPERADORESLOGICOSContext ctx);
	/**
	 * Enter a parse tree produced by the {@code structInstanceCreation}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void enterStructInstanceCreation(VlangParser.StructInstanceCreationContext ctx);
	/**
	 * Exit a parse tree produced by the {@code structInstanceCreation}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void exitStructInstanceCreation(VlangParser.StructInstanceCreationContext ctx);
	/**
	 * Enter a parse tree produced by the {@code valorexpr}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void enterValorexpr(VlangParser.ValorexprContext ctx);
	/**
	 * Exit a parse tree produced by the {@code valorexpr}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void exitValorexpr(VlangParser.ValorexprContext ctx);
	/**
	 * Enter a parse tree produced by the {@code igualdad}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void enterIgualdad(VlangParser.IgualdadContext ctx);
	/**
	 * Exit a parse tree produced by the {@code igualdad}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void exitIgualdad(VlangParser.IgualdadContext ctx);
	/**
	 * Enter a parse tree produced by the {@code llamadaFuncionExpr}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void enterLlamadaFuncionExpr(VlangParser.LlamadaFuncionExprContext ctx);
	/**
	 * Exit a parse tree produced by the {@code llamadaFuncionExpr}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void exitLlamadaFuncionExpr(VlangParser.LlamadaFuncionExprContext ctx);
	/**
	 * Enter a parse tree produced by the {@code expdotexp}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void enterExpdotexp(VlangParser.ExpdotexpContext ctx);
	/**
	 * Exit a parse tree produced by the {@code expdotexp}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void exitExpdotexp(VlangParser.ExpdotexpContext ctx);
	/**
	 * Enter a parse tree produced by the {@code structAttrAssign}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void enterStructAttrAssign(VlangParser.StructAttrAssignContext ctx);
	/**
	 * Exit a parse tree produced by the {@code structAttrAssign}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void exitStructAttrAssign(VlangParser.StructAttrAssignContext ctx);
	/**
	 * Enter a parse tree produced by the {@code relacionales}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void enterRelacionales(VlangParser.RelacionalesContext ctx);
	/**
	 * Exit a parse tree produced by the {@code relacionales}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void exitRelacionales(VlangParser.RelacionalesContext ctx);
	/**
	 * Enter a parse tree produced by the {@code casteo_paratipo_slice}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void enterCasteo_paratipo_slice(VlangParser.Casteo_paratipo_sliceContext ctx);
	/**
	 * Exit a parse tree produced by the {@code casteo_paratipo_slice}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void exitCasteo_paratipo_slice(VlangParser.Casteo_paratipo_sliceContext ctx);
	/**
	 * Enter a parse tree produced by the {@code corchetesexpre}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void enterCorchetesexpre(VlangParser.CorchetesexpreContext ctx);
	/**
	 * Exit a parse tree produced by the {@code corchetesexpre}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void exitCorchetesexpre(VlangParser.CorchetesexpreContext ctx);
	/**
	 * Enter a parse tree produced by the {@code unario}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void enterUnario(VlangParser.UnarioContext ctx);
	/**
	 * Exit a parse tree produced by the {@code unario}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void exitUnario(VlangParser.UnarioContext ctx);
	/**
	 * Enter a parse tree produced by the {@code parentesisexpre}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void enterParentesisexpre(VlangParser.ParentesisexpreContext ctx);
	/**
	 * Exit a parse tree produced by the {@code parentesisexpre}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void exitParentesisexpre(VlangParser.ParentesisexpreContext ctx);
	/**
	 * Enter a parse tree produced by the {@code IMCPLICIT}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void enterIMCPLICIT(VlangParser.IMCPLICITContext ctx);
	/**
	 * Exit a parse tree produced by the {@code IMCPLICIT}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void exitIMCPLICIT(VlangParser.IMCPLICITContext ctx);
	/**
	 * Enter a parse tree produced by the {@code sumres}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void enterSumres(VlangParser.SumresContext ctx);
	/**
	 * Exit a parse tree produced by the {@code sumres}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void exitSumres(VlangParser.SumresContext ctx);
	/**
	 * Enter a parse tree produced by the {@code PARAPRINTSLICE}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void enterPARAPRINTSLICE(VlangParser.PARAPRINTSLICEContext ctx);
	/**
	 * Exit a parse tree produced by the {@code PARAPRINTSLICE}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void exitPARAPRINTSLICE(VlangParser.PARAPRINTSLICEContext ctx);
	/**
	 * Enter a parse tree produced by the {@code asignacionLUEGO}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void enterAsignacionLUEGO(VlangParser.AsignacionLUEGOContext ctx);
	/**
	 * Exit a parse tree produced by the {@code asignacionLUEGO}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void exitAsignacionLUEGO(VlangParser.AsignacionLUEGOContext ctx);
	/**
	 * Enter a parse tree produced by the {@code id}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void enterId(VlangParser.IdContext ctx);
	/**
	 * Exit a parse tree produced by the {@code id}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void exitId(VlangParser.IdContext ctx);
	/**
	 * Enter a parse tree produced by the {@code expdotexp1}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void enterExpdotexp1(VlangParser.Expdotexp1Context ctx);
	/**
	 * Exit a parse tree produced by the {@code expdotexp1}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void exitExpdotexp1(VlangParser.Expdotexp1Context ctx);
	/**
	 * Enter a parse tree produced by {@link VlangParser#parametros}.
	 * @param ctx the parse tree
	 */
	void enterParametros(VlangParser.ParametrosContext ctx);
	/**
	 * Exit a parse tree produced by {@link VlangParser#parametros}.
	 * @param ctx the parse tree
	 */
	void exitParametros(VlangParser.ParametrosContext ctx);
	/**
	 * Enter a parse tree produced by {@link VlangParser#valores}.
	 * @param ctx the parse tree
	 */
	void enterValores(VlangParser.ValoresContext ctx);
	/**
	 * Exit a parse tree produced by {@link VlangParser#valores}.
	 * @param ctx the parse tree
	 */
	void exitValores(VlangParser.ValoresContext ctx);
	/**
	 * Enter a parse tree produced by the {@code valorEntero}
	 * labeled alternative in {@link VlangParser#valor}.
	 * @param ctx the parse tree
	 */
	void enterValorEntero(VlangParser.ValorEnteroContext ctx);
	/**
	 * Exit a parse tree produced by the {@code valorEntero}
	 * labeled alternative in {@link VlangParser#valor}.
	 * @param ctx the parse tree
	 */
	void exitValorEntero(VlangParser.ValorEnteroContext ctx);
	/**
	 * Enter a parse tree produced by the {@code valorDecimal}
	 * labeled alternative in {@link VlangParser#valor}.
	 * @param ctx the parse tree
	 */
	void enterValorDecimal(VlangParser.ValorDecimalContext ctx);
	/**
	 * Exit a parse tree produced by the {@code valorDecimal}
	 * labeled alternative in {@link VlangParser#valor}.
	 * @param ctx the parse tree
	 */
	void exitValorDecimal(VlangParser.ValorDecimalContext ctx);
	/**
	 * Enter a parse tree produced by the {@code valorCadena}
	 * labeled alternative in {@link VlangParser#valor}.
	 * @param ctx the parse tree
	 */
	void enterValorCadena(VlangParser.ValorCadenaContext ctx);
	/**
	 * Exit a parse tree produced by the {@code valorCadena}
	 * labeled alternative in {@link VlangParser#valor}.
	 * @param ctx the parse tree
	 */
	void exitValorCadena(VlangParser.ValorCadenaContext ctx);
	/**
	 * Enter a parse tree produced by the {@code valorBooleano}
	 * labeled alternative in {@link VlangParser#valor}.
	 * @param ctx the parse tree
	 */
	void enterValorBooleano(VlangParser.ValorBooleanoContext ctx);
	/**
	 * Exit a parse tree produced by the {@code valorBooleano}
	 * labeled alternative in {@link VlangParser#valor}.
	 * @param ctx the parse tree
	 */
	void exitValorBooleano(VlangParser.ValorBooleanoContext ctx);
	/**
	 * Enter a parse tree produced by the {@code valorCaracter}
	 * labeled alternative in {@link VlangParser#valor}.
	 * @param ctx the parse tree
	 */
	void enterValorCaracter(VlangParser.ValorCaracterContext ctx);
	/**
	 * Exit a parse tree produced by the {@code valorCaracter}
	 * labeled alternative in {@link VlangParser#valor}.
	 * @param ctx the parse tree
	 */
	void exitValorCaracter(VlangParser.ValorCaracterContext ctx);
	/**
	 * Enter a parse tree produced by {@link VlangParser#listaExpresiones}.
	 * @param ctx the parse tree
	 */
	void enterListaExpresiones(VlangParser.ListaExpresionesContext ctx);
	/**
	 * Exit a parse tree produced by {@link VlangParser#listaExpresiones}.
	 * @param ctx the parse tree
	 */
	void exitListaExpresiones(VlangParser.ListaExpresionesContext ctx);
	/**
	 * Enter a parse tree produced by the {@code incremento}
	 * labeled alternative in {@link VlangParser#incredecre}.
	 * @param ctx the parse tree
	 */
	void enterIncremento(VlangParser.IncrementoContext ctx);
	/**
	 * Exit a parse tree produced by the {@code incremento}
	 * labeled alternative in {@link VlangParser#incredecre}.
	 * @param ctx the parse tree
	 */
	void exitIncremento(VlangParser.IncrementoContext ctx);
	/**
	 * Enter a parse tree produced by the {@code decremento}
	 * labeled alternative in {@link VlangParser#incredecre}.
	 * @param ctx the parse tree
	 */
	void enterDecremento(VlangParser.DecrementoContext ctx);
	/**
	 * Exit a parse tree produced by the {@code decremento}
	 * labeled alternative in {@link VlangParser#incredecre}.
	 * @param ctx the parse tree
	 */
	void exitDecremento(VlangParser.DecrementoContext ctx);
}