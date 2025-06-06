// Code generated from V4LangGrammar.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // V4LangGrammar

import "github.com/antlr4-go/antlr/v4"

// BaseV4LangGrammarListener is a complete listener for a parse tree produced by V4LangGrammarParser.
type BaseV4LangGrammarListener struct{}

var _ V4LangGrammarListener = &BaseV4LangGrammarListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseV4LangGrammarListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseV4LangGrammarListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseV4LangGrammarListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseV4LangGrammarListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseV4LangGrammarListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseV4LangGrammarListener) ExitProgram(ctx *ProgramContext) {}

// EnterDcl is called when production dcl is entered.
func (s *BaseV4LangGrammarListener) EnterDcl(ctx *DclContext) {}

// ExitDcl is called when production dcl is exited.
func (s *BaseV4LangGrammarListener) ExitDcl(ctx *DclContext) {}

// EnterVarDcl is called when production varDcl is entered.
func (s *BaseV4LangGrammarListener) EnterVarDcl(ctx *VarDclContext) {}

// ExitVarDcl is called when production varDcl is exited.
func (s *BaseV4LangGrammarListener) ExitVarDcl(ctx *VarDclContext) {}

// EnterFuncDcl is called when production funcDcl is entered.
func (s *BaseV4LangGrammarListener) EnterFuncDcl(ctx *FuncDclContext) {}

// ExitFuncDcl is called when production funcDcl is exited.
func (s *BaseV4LangGrammarListener) ExitFuncDcl(ctx *FuncDclContext) {}

// EnterParams is called when production params is entered.
func (s *BaseV4LangGrammarListener) EnterParams(ctx *ParamsContext) {}

// ExitParams is called when production params is exited.
func (s *BaseV4LangGrammarListener) ExitParams(ctx *ParamsContext) {}

// EnterStructDcl is called when production structDcl is entered.
func (s *BaseV4LangGrammarListener) EnterStructDcl(ctx *StructDclContext) {}

// ExitStructDcl is called when production structDcl is exited.
func (s *BaseV4LangGrammarListener) ExitStructDcl(ctx *StructDclContext) {}

// EnterStructBody is called when production structBody is entered.
func (s *BaseV4LangGrammarListener) EnterStructBody(ctx *StructBodyContext) {}

// ExitStructBody is called when production structBody is exited.
func (s *BaseV4LangGrammarListener) ExitStructBody(ctx *StructBodyContext) {}

// EnterReceiver is called when production receiver is entered.
func (s *BaseV4LangGrammarListener) EnterReceiver(ctx *ReceiverContext) {}

// ExitReceiver is called when production receiver is exited.
func (s *BaseV4LangGrammarListener) ExitReceiver(ctx *ReceiverContext) {}

// EnterSliceDcl is called when production sliceDcl is entered.
func (s *BaseV4LangGrammarListener) EnterSliceDcl(ctx *SliceDclContext) {}

// ExitSliceDcl is called when production sliceDcl is exited.
func (s *BaseV4LangGrammarListener) ExitSliceDcl(ctx *SliceDclContext) {}

// EnterSliceValues is called when production sliceValues is entered.
func (s *BaseV4LangGrammarListener) EnterSliceValues(ctx *SliceValuesContext) {}

// ExitSliceValues is called when production sliceValues is exited.
func (s *BaseV4LangGrammarListener) ExitSliceValues(ctx *SliceValuesContext) {}

// EnterMatrixDcl is called when production matrixDcl is entered.
func (s *BaseV4LangGrammarListener) EnterMatrixDcl(ctx *MatrixDclContext) {}

// ExitMatrixDcl is called when production matrixDcl is exited.
func (s *BaseV4LangGrammarListener) ExitMatrixDcl(ctx *MatrixDclContext) {}

// EnterMatrixValues is called when production matrixValues is entered.
func (s *BaseV4LangGrammarListener) EnterMatrixValues(ctx *MatrixValuesContext) {}

// ExitMatrixValues is called when production matrixValues is exited.
func (s *BaseV4LangGrammarListener) ExitMatrixValues(ctx *MatrixValuesContext) {}

// EnterExprStmt is called when production ExprStmt is entered.
func (s *BaseV4LangGrammarListener) EnterExprStmt(ctx *ExprStmtContext) {}

// ExitExprStmt is called when production ExprStmt is exited.
func (s *BaseV4LangGrammarListener) ExitExprStmt(ctx *ExprStmtContext) {}

// EnterPrintStmt is called when production PrintStmt is entered.
func (s *BaseV4LangGrammarListener) EnterPrintStmt(ctx *PrintStmtContext) {}

// ExitPrintStmt is called when production PrintStmt is exited.
func (s *BaseV4LangGrammarListener) ExitPrintStmt(ctx *PrintStmtContext) {}

// EnterBlockStmt is called when production BlockStmt is entered.
func (s *BaseV4LangGrammarListener) EnterBlockStmt(ctx *BlockStmtContext) {}

// ExitBlockStmt is called when production BlockStmt is exited.
func (s *BaseV4LangGrammarListener) ExitBlockStmt(ctx *BlockStmtContext) {}

// EnterIfStmt is called when production IfStmt is entered.
func (s *BaseV4LangGrammarListener) EnterIfStmt(ctx *IfStmtContext) {}

// ExitIfStmt is called when production IfStmt is exited.
func (s *BaseV4LangGrammarListener) ExitIfStmt(ctx *IfStmtContext) {}

// EnterSwitchStmt is called when production SwitchStmt is entered.
func (s *BaseV4LangGrammarListener) EnterSwitchStmt(ctx *SwitchStmtContext) {}

// ExitSwitchStmt is called when production SwitchStmt is exited.
func (s *BaseV4LangGrammarListener) ExitSwitchStmt(ctx *SwitchStmtContext) {}

// EnterForIncrementStmt is called when production ForIncrementStmt is entered.
func (s *BaseV4LangGrammarListener) EnterForIncrementStmt(ctx *ForIncrementStmtContext) {}

// ExitForIncrementStmt is called when production ForIncrementStmt is exited.
func (s *BaseV4LangGrammarListener) ExitForIncrementStmt(ctx *ForIncrementStmtContext) {}

// EnterForStmt is called when production ForStmt is entered.
func (s *BaseV4LangGrammarListener) EnterForStmt(ctx *ForStmtContext) {}

// ExitForStmt is called when production ForStmt is exited.
func (s *BaseV4LangGrammarListener) ExitForStmt(ctx *ForStmtContext) {}

// EnterForRangeStmt is called when production ForRangeStmt is entered.
func (s *BaseV4LangGrammarListener) EnterForRangeStmt(ctx *ForRangeStmtContext) {}

// ExitForRangeStmt is called when production ForRangeStmt is exited.
func (s *BaseV4LangGrammarListener) ExitForRangeStmt(ctx *ForRangeStmtContext) {}

// EnterBreakStmt is called when production BreakStmt is entered.
func (s *BaseV4LangGrammarListener) EnterBreakStmt(ctx *BreakStmtContext) {}

// ExitBreakStmt is called when production BreakStmt is exited.
func (s *BaseV4LangGrammarListener) ExitBreakStmt(ctx *BreakStmtContext) {}

// EnterContinueStmt is called when production ContinueStmt is entered.
func (s *BaseV4LangGrammarListener) EnterContinueStmt(ctx *ContinueStmtContext) {}

// ExitContinueStmt is called when production ContinueStmt is exited.
func (s *BaseV4LangGrammarListener) ExitContinueStmt(ctx *ContinueStmtContext) {}

// EnterReturnStmt is called when production ReturnStmt is entered.
func (s *BaseV4LangGrammarListener) EnterReturnStmt(ctx *ReturnStmtContext) {}

// ExitReturnStmt is called when production ReturnStmt is exited.
func (s *BaseV4LangGrammarListener) ExitReturnStmt(ctx *ReturnStmtContext) {}

// EnterElseIfStmt is called when production elseIfStmt is entered.
func (s *BaseV4LangGrammarListener) EnterElseIfStmt(ctx *ElseIfStmtContext) {}

// ExitElseIfStmt is called when production elseIfStmt is exited.
func (s *BaseV4LangGrammarListener) ExitElseIfStmt(ctx *ElseIfStmtContext) {}

// EnterElseStmt is called when production elseStmt is entered.
func (s *BaseV4LangGrammarListener) EnterElseStmt(ctx *ElseStmtContext) {}

// ExitElseStmt is called when production elseStmt is exited.
func (s *BaseV4LangGrammarListener) ExitElseStmt(ctx *ElseStmtContext) {}

// EnterSwitchCase is called when production switchCase is entered.
func (s *BaseV4LangGrammarListener) EnterSwitchCase(ctx *SwitchCaseContext) {}

// ExitSwitchCase is called when production switchCase is exited.
func (s *BaseV4LangGrammarListener) ExitSwitchCase(ctx *SwitchCaseContext) {}

// EnterDefaultCase is called when production defaultCase is entered.
func (s *BaseV4LangGrammarListener) EnterDefaultCase(ctx *DefaultCaseContext) {}

// ExitDefaultCase is called when production defaultCase is exited.
func (s *BaseV4LangGrammarListener) ExitDefaultCase(ctx *DefaultCaseContext) {}

// EnterArgs is called when production args is entered.
func (s *BaseV4LangGrammarListener) EnterArgs(ctx *ArgsContext) {}

// ExitArgs is called when production args is exited.
func (s *BaseV4LangGrammarListener) ExitArgs(ctx *ArgsContext) {}

// EnterMulDiv is called when production MulDiv is entered.
func (s *BaseV4LangGrammarListener) EnterMulDiv(ctx *MulDivContext) {}

// ExitMulDiv is called when production MulDiv is exited.
func (s *BaseV4LangGrammarListener) ExitMulDiv(ctx *MulDivContext) {}

// EnterGetValueIndex is called when production GetValueIndex is entered.
func (s *BaseV4LangGrammarListener) EnterGetValueIndex(ctx *GetValueIndexContext) {}

// ExitGetValueIndex is called when production GetValueIndex is exited.
func (s *BaseV4LangGrammarListener) ExitGetValueIndex(ctx *GetValueIndexContext) {}

// EnterParens is called when production Parens is entered.
func (s *BaseV4LangGrammarListener) EnterParens(ctx *ParensContext) {}

// ExitParens is called when production Parens is exited.
func (s *BaseV4LangGrammarListener) ExitParens(ctx *ParensContext) {}

// EnterLogical is called when production Logical is entered.
func (s *BaseV4LangGrammarListener) EnterLogical(ctx *LogicalContext) {}

// ExitLogical is called when production Logical is exited.
func (s *BaseV4LangGrammarListener) ExitLogical(ctx *LogicalContext) {}

// EnterString is called when production String is entered.
func (s *BaseV4LangGrammarListener) EnterString(ctx *StringContext) {}

// ExitString is called when production String is exited.
func (s *BaseV4LangGrammarListener) ExitString(ctx *StringContext) {}

// EnterInt is called when production Int is entered.
func (s *BaseV4LangGrammarListener) EnterInt(ctx *IntContext) {}

// ExitInt is called when production Int is exited.
func (s *BaseV4LangGrammarListener) ExitInt(ctx *IntContext) {}

// EnterIdentifier is called when production Identifier is entered.
func (s *BaseV4LangGrammarListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production Identifier is exited.
func (s *BaseV4LangGrammarListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterSetValueIndex is called when production SetValueIndex is entered.
func (s *BaseV4LangGrammarListener) EnterSetValueIndex(ctx *SetValueIndexContext) {}

// ExitSetValueIndex is called when production SetValueIndex is exited.
func (s *BaseV4LangGrammarListener) ExitSetValueIndex(ctx *SetValueIndexContext) {}

// EnterGetValueMatrix is called when production GetValueMatrix is entered.
func (s *BaseV4LangGrammarListener) EnterGetValueMatrix(ctx *GetValueMatrixContext) {}

// ExitGetValueMatrix is called when production GetValueMatrix is exited.
func (s *BaseV4LangGrammarListener) ExitGetValueMatrix(ctx *GetValueMatrixContext) {}

// EnterSliceWithValues is called when production SliceWithValues is entered.
func (s *BaseV4LangGrammarListener) EnterSliceWithValues(ctx *SliceWithValuesContext) {}

// ExitSliceWithValues is called when production SliceWithValues is exited.
func (s *BaseV4LangGrammarListener) ExitSliceWithValues(ctx *SliceWithValuesContext) {}

// EnterEquality is called when production Equality is entered.
func (s *BaseV4LangGrammarListener) EnterEquality(ctx *EqualityContext) {}

// ExitEquality is called when production Equality is exited.
func (s *BaseV4LangGrammarListener) ExitEquality(ctx *EqualityContext) {}

// EnterBoolean is called when production Boolean is entered.
func (s *BaseV4LangGrammarListener) EnterBoolean(ctx *BooleanContext) {}

// ExitBoolean is called when production Boolean is exited.
func (s *BaseV4LangGrammarListener) ExitBoolean(ctx *BooleanContext) {}

// EnterNewInstance is called when production NewInstance is entered.
func (s *BaseV4LangGrammarListener) EnterNewInstance(ctx *NewInstanceContext) {}

// ExitNewInstance is called when production NewInstance is exited.
func (s *BaseV4LangGrammarListener) ExitNewInstance(ctx *NewInstanceContext) {}

// EnterSetValueMatrix is called when production SetValueMatrix is entered.
func (s *BaseV4LangGrammarListener) EnterSetValueMatrix(ctx *SetValueMatrixContext) {}

// ExitSetValueMatrix is called when production SetValueMatrix is exited.
func (s *BaseV4LangGrammarListener) ExitSetValueMatrix(ctx *SetValueMatrixContext) {}

// EnterMod is called when production Mod is entered.
func (s *BaseV4LangGrammarListener) EnterMod(ctx *ModContext) {}

// ExitMod is called when production Mod is exited.
func (s *BaseV4LangGrammarListener) ExitMod(ctx *ModContext) {}

// EnterSliceIndex is called when production SliceIndex is entered.
func (s *BaseV4LangGrammarListener) EnterSliceIndex(ctx *SliceIndexContext) {}

// ExitSliceIndex is called when production SliceIndex is exited.
func (s *BaseV4LangGrammarListener) ExitSliceIndex(ctx *SliceIndexContext) {}

// EnterTypeOf is called when production TypeOf is entered.
func (s *BaseV4LangGrammarListener) EnterTypeOf(ctx *TypeOfContext) {}

// ExitTypeOf is called when production TypeOf is exited.
func (s *BaseV4LangGrammarListener) ExitTypeOf(ctx *TypeOfContext) {}

// EnterAddSub is called when production AddSub is entered.
func (s *BaseV4LangGrammarListener) EnterAddSub(ctx *AddSubContext) {}

// ExitAddSub is called when production AddSub is exited.
func (s *BaseV4LangGrammarListener) ExitAddSub(ctx *AddSubContext) {}

// EnterStructInstance is called when production StructInstance is entered.
func (s *BaseV4LangGrammarListener) EnterStructInstance(ctx *StructInstanceContext) {}

// ExitStructInstance is called when production StructInstance is exited.
func (s *BaseV4LangGrammarListener) ExitStructInstance(ctx *StructInstanceContext) {}

// EnterRelational is called when production Relational is entered.
func (s *BaseV4LangGrammarListener) EnterRelational(ctx *RelationalContext) {}

// ExitRelational is called when production Relational is exited.
func (s *BaseV4LangGrammarListener) ExitRelational(ctx *RelationalContext) {}

// EnterCallStmt is called when production CallStmt is entered.
func (s *BaseV4LangGrammarListener) EnterCallStmt(ctx *CallStmtContext) {}

// ExitCallStmt is called when production CallStmt is exited.
func (s *BaseV4LangGrammarListener) ExitCallStmt(ctx *CallStmtContext) {}

// EnterNil is called when production Nil is entered.
func (s *BaseV4LangGrammarListener) EnterNil(ctx *NilContext) {}

// ExitNil is called when production Nil is exited.
func (s *BaseV4LangGrammarListener) ExitNil(ctx *NilContext) {}

// EnterStringJoin is called when production StringJoin is entered.
func (s *BaseV4LangGrammarListener) EnterStringJoin(ctx *StringJoinContext) {}

// ExitStringJoin is called when production StringJoin is exited.
func (s *BaseV4LangGrammarListener) ExitStringJoin(ctx *StringJoinContext) {}

// EnterFloat is called when production Float is entered.
func (s *BaseV4LangGrammarListener) EnterFloat(ctx *FloatContext) {}

// ExitFloat is called when production Float is exited.
func (s *BaseV4LangGrammarListener) ExitFloat(ctx *FloatContext) {}

// EnterNot is called when production Not is entered.
func (s *BaseV4LangGrammarListener) EnterNot(ctx *NotContext) {}

// ExitNot is called when production Not is exited.
func (s *BaseV4LangGrammarListener) ExitNot(ctx *NotContext) {}

// EnterAtoi is called when production Atoi is entered.
func (s *BaseV4LangGrammarListener) EnterAtoi(ctx *AtoiContext) {}

// ExitAtoi is called when production Atoi is exited.
func (s *BaseV4LangGrammarListener) ExitAtoi(ctx *AtoiContext) {}

// EnterParseFloat is called when production ParseFloat is entered.
func (s *BaseV4LangGrammarListener) EnterParseFloat(ctx *ParseFloatContext) {}

// ExitParseFloat is called when production ParseFloat is exited.
func (s *BaseV4LangGrammarListener) ExitParseFloat(ctx *ParseFloatContext) {}

// EnterAppend is called when production Append is entered.
func (s *BaseV4LangGrammarListener) EnterAppend(ctx *AppendContext) {}

// ExitAppend is called when production Append is exited.
func (s *BaseV4LangGrammarListener) ExitAppend(ctx *AppendContext) {}

// EnterLen is called when production Len is entered.
func (s *BaseV4LangGrammarListener) EnterLen(ctx *LenContext) {}

// ExitLen is called when production Len is exited.
func (s *BaseV4LangGrammarListener) ExitLen(ctx *LenContext) {}

// EnterAssignArithmetic is called when production AssignArithmetic is entered.
func (s *BaseV4LangGrammarListener) EnterAssignArithmetic(ctx *AssignArithmeticContext) {}

// ExitAssignArithmetic is called when production AssignArithmetic is exited.
func (s *BaseV4LangGrammarListener) ExitAssignArithmetic(ctx *AssignArithmeticContext) {}

// EnterSlice is called when production Slice is entered.
func (s *BaseV4LangGrammarListener) EnterSlice(ctx *SliceContext) {}

// ExitSlice is called when production Slice is exited.
func (s *BaseV4LangGrammarListener) ExitSlice(ctx *SliceContext) {}

// EnterAssign is called when production Assign is entered.
func (s *BaseV4LangGrammarListener) EnterAssign(ctx *AssignContext) {}

// ExitAssign is called when production Assign is exited.
func (s *BaseV4LangGrammarListener) ExitAssign(ctx *AssignContext) {}

// EnterNegate is called when production Negate is entered.
func (s *BaseV4LangGrammarListener) EnterNegate(ctx *NegateContext) {}

// ExitNegate is called when production Negate is exited.
func (s *BaseV4LangGrammarListener) ExitNegate(ctx *NegateContext) {}

// EnterIncrementDecrement is called when production IncrementDecrement is entered.
func (s *BaseV4LangGrammarListener) EnterIncrementDecrement(ctx *IncrementDecrementContext) {}

// ExitIncrementDecrement is called when production IncrementDecrement is exited.
func (s *BaseV4LangGrammarListener) ExitIncrementDecrement(ctx *IncrementDecrementContext) {}

// EnterRune is called when production Rune is entered.
func (s *BaseV4LangGrammarListener) EnterRune(ctx *RuneContext) {}

// ExitRune is called when production Rune is exited.
func (s *BaseV4LangGrammarListener) ExitRune(ctx *RuneContext) {}

// EnterFuncCall is called when production FuncCall is entered.
func (s *BaseV4LangGrammarListener) EnterFuncCall(ctx *FuncCallContext) {}

// ExitFuncCall is called when production FuncCall is exited.
func (s *BaseV4LangGrammarListener) ExitFuncCall(ctx *FuncCallContext) {}

// EnterGet is called when production Get is entered.
func (s *BaseV4LangGrammarListener) EnterGet(ctx *GetContext) {}

// ExitGet is called when production Get is exited.
func (s *BaseV4LangGrammarListener) ExitGet(ctx *GetContext) {}

// EnterInstanceValues is called when production instanceValues is entered.
func (s *BaseV4LangGrammarListener) EnterInstanceValues(ctx *InstanceValuesContext) {}

// ExitInstanceValues is called when production instanceValues is exited.
func (s *BaseV4LangGrammarListener) ExitInstanceValues(ctx *InstanceValuesContext) {}

// EnterType is called when production type is entered.
func (s *BaseV4LangGrammarListener) EnterType(ctx *TypeContext) {}

// ExitType is called when production type is exited.
func (s *BaseV4LangGrammarListener) ExitType(ctx *TypeContext) {}
