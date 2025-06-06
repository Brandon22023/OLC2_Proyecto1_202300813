// Code generated from V4LangGrammar.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // V4LangGrammar

import "github.com/antlr4-go/antlr/v4"

type BaseV4LangGrammarVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseV4LangGrammarVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV4LangGrammarVisitor) VisitDcl(ctx *DclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV4LangGrammarVisitor) VisitVarDcl(ctx *VarDclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV4LangGrammarVisitor) VisitFuncDcl(ctx *FuncDclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV4LangGrammarVisitor) VisitParams(ctx *ParamsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV4LangGrammarVisitor) VisitStructDcl(ctx *StructDclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV4LangGrammarVisitor) VisitStructBody(ctx *StructBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV4LangGrammarVisitor) VisitReceiver(ctx *ReceiverContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV4LangGrammarVisitor) VisitSliceDcl(ctx *SliceDclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV4LangGrammarVisitor) VisitSliceValues(ctx *SliceValuesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV4LangGrammarVisitor) VisitMatrixDcl(ctx *MatrixDclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV4LangGrammarVisitor) VisitMatrixValues(ctx *MatrixValuesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV4LangGrammarVisitor) VisitExprStmt(ctx *ExprStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV4LangGrammarVisitor) VisitPrintStmt(ctx *PrintStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV4LangGrammarVisitor) VisitBlockStmt(ctx *BlockStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV4LangGrammarVisitor) VisitIfStmt(ctx *IfStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV4LangGrammarVisitor) VisitSwitchStmt(ctx *SwitchStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV4LangGrammarVisitor) VisitForIncrementStmt(ctx *ForIncrementStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV4LangGrammarVisitor) VisitForStmt(ctx *ForStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV4LangGrammarVisitor) VisitForRangeStmt(ctx *ForRangeStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV4LangGrammarVisitor) VisitBreakStmt(ctx *BreakStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV4LangGrammarVisitor) VisitContinueStmt(ctx *ContinueStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV4LangGrammarVisitor) VisitReturnStmt(ctx *ReturnStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV4LangGrammarVisitor) VisitElseIfStmt(ctx *ElseIfStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV4LangGrammarVisitor) VisitElseStmt(ctx *ElseStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV4LangGrammarVisitor) VisitSwitchCase(ctx *SwitchCaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV4LangGrammarVisitor) VisitDefaultCase(ctx *DefaultCaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV4LangGrammarVisitor) VisitArgs(ctx *ArgsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV4LangGrammarVisitor) VisitMulDiv(ctx *MulDivContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV4LangGrammarVisitor) VisitGetValueIndex(ctx *GetValueIndexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV4LangGrammarVisitor) VisitParens(ctx *ParensContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV4LangGrammarVisitor) VisitLogical(ctx *LogicalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV4LangGrammarVisitor) VisitString(ctx *StringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV4LangGrammarVisitor) VisitInt(ctx *IntContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV4LangGrammarVisitor) VisitIdentifier(ctx *IdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV4LangGrammarVisitor) VisitSetValueIndex(ctx *SetValueIndexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV4LangGrammarVisitor) VisitGetValueMatrix(ctx *GetValueMatrixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV4LangGrammarVisitor) VisitSliceWithValues(ctx *SliceWithValuesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV4LangGrammarVisitor) VisitEquality(ctx *EqualityContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV4LangGrammarVisitor) VisitBoolean(ctx *BooleanContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV4LangGrammarVisitor) VisitNewInstance(ctx *NewInstanceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV4LangGrammarVisitor) VisitSetValueMatrix(ctx *SetValueMatrixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV4LangGrammarVisitor) VisitMod(ctx *ModContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV4LangGrammarVisitor) VisitSliceIndex(ctx *SliceIndexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV4LangGrammarVisitor) VisitTypeOf(ctx *TypeOfContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV4LangGrammarVisitor) VisitAddSub(ctx *AddSubContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV4LangGrammarVisitor) VisitStructInstance(ctx *StructInstanceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV4LangGrammarVisitor) VisitRelational(ctx *RelationalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV4LangGrammarVisitor) VisitCallStmt(ctx *CallStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV4LangGrammarVisitor) VisitNil(ctx *NilContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV4LangGrammarVisitor) VisitStringJoin(ctx *StringJoinContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV4LangGrammarVisitor) VisitFloat(ctx *FloatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV4LangGrammarVisitor) VisitNot(ctx *NotContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV4LangGrammarVisitor) VisitAtoi(ctx *AtoiContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV4LangGrammarVisitor) VisitParseFloat(ctx *ParseFloatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV4LangGrammarVisitor) VisitAppend(ctx *AppendContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV4LangGrammarVisitor) VisitLen(ctx *LenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV4LangGrammarVisitor) VisitAssignArithmetic(ctx *AssignArithmeticContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV4LangGrammarVisitor) VisitSlice(ctx *SliceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV4LangGrammarVisitor) VisitAssign(ctx *AssignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV4LangGrammarVisitor) VisitNegate(ctx *NegateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV4LangGrammarVisitor) VisitIncrementDecrement(ctx *IncrementDecrementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV4LangGrammarVisitor) VisitRune(ctx *RuneContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV4LangGrammarVisitor) VisitFuncCall(ctx *FuncCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV4LangGrammarVisitor) VisitGet(ctx *GetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV4LangGrammarVisitor) VisitInstanceValues(ctx *InstanceValuesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV4LangGrammarVisitor) VisitType(ctx *TypeContext) interface{} {
	return v.VisitChildren(ctx)
}
