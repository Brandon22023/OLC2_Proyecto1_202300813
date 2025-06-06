// Code generated from V4LangGrammar.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // V4LangGrammar

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by V4LangGrammarParser.
type V4LangGrammarVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by V4LangGrammarParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by V4LangGrammarParser#dcl.
	VisitDcl(ctx *DclContext) interface{}

	// Visit a parse tree produced by V4LangGrammarParser#varDcl.
	VisitVarDcl(ctx *VarDclContext) interface{}

	// Visit a parse tree produced by V4LangGrammarParser#funcDcl.
	VisitFuncDcl(ctx *FuncDclContext) interface{}

	// Visit a parse tree produced by V4LangGrammarParser#params.
	VisitParams(ctx *ParamsContext) interface{}

	// Visit a parse tree produced by V4LangGrammarParser#structDcl.
	VisitStructDcl(ctx *StructDclContext) interface{}

	// Visit a parse tree produced by V4LangGrammarParser#structBody.
	VisitStructBody(ctx *StructBodyContext) interface{}

	// Visit a parse tree produced by V4LangGrammarParser#receiver.
	VisitReceiver(ctx *ReceiverContext) interface{}

	// Visit a parse tree produced by V4LangGrammarParser#sliceDcl.
	VisitSliceDcl(ctx *SliceDclContext) interface{}

	// Visit a parse tree produced by V4LangGrammarParser#sliceValues.
	VisitSliceValues(ctx *SliceValuesContext) interface{}

	// Visit a parse tree produced by V4LangGrammarParser#matrixDcl.
	VisitMatrixDcl(ctx *MatrixDclContext) interface{}

	// Visit a parse tree produced by V4LangGrammarParser#matrixValues.
	VisitMatrixValues(ctx *MatrixValuesContext) interface{}

	// Visit a parse tree produced by V4LangGrammarParser#ExprStmt.
	VisitExprStmt(ctx *ExprStmtContext) interface{}

	// Visit a parse tree produced by V4LangGrammarParser#PrintStmt.
	VisitPrintStmt(ctx *PrintStmtContext) interface{}

	// Visit a parse tree produced by V4LangGrammarParser#BlockStmt.
	VisitBlockStmt(ctx *BlockStmtContext) interface{}

	// Visit a parse tree produced by V4LangGrammarParser#IfStmt.
	VisitIfStmt(ctx *IfStmtContext) interface{}

	// Visit a parse tree produced by V4LangGrammarParser#SwitchStmt.
	VisitSwitchStmt(ctx *SwitchStmtContext) interface{}

	// Visit a parse tree produced by V4LangGrammarParser#ForIncrementStmt.
	VisitForIncrementStmt(ctx *ForIncrementStmtContext) interface{}

	// Visit a parse tree produced by V4LangGrammarParser#ForStmt.
	VisitForStmt(ctx *ForStmtContext) interface{}

	// Visit a parse tree produced by V4LangGrammarParser#ForRangeStmt.
	VisitForRangeStmt(ctx *ForRangeStmtContext) interface{}

	// Visit a parse tree produced by V4LangGrammarParser#BreakStmt.
	VisitBreakStmt(ctx *BreakStmtContext) interface{}

	// Visit a parse tree produced by V4LangGrammarParser#ContinueStmt.
	VisitContinueStmt(ctx *ContinueStmtContext) interface{}

	// Visit a parse tree produced by V4LangGrammarParser#ReturnStmt.
	VisitReturnStmt(ctx *ReturnStmtContext) interface{}

	// Visit a parse tree produced by V4LangGrammarParser#elseIfStmt.
	VisitElseIfStmt(ctx *ElseIfStmtContext) interface{}

	// Visit a parse tree produced by V4LangGrammarParser#elseStmt.
	VisitElseStmt(ctx *ElseStmtContext) interface{}

	// Visit a parse tree produced by V4LangGrammarParser#switchCase.
	VisitSwitchCase(ctx *SwitchCaseContext) interface{}

	// Visit a parse tree produced by V4LangGrammarParser#defaultCase.
	VisitDefaultCase(ctx *DefaultCaseContext) interface{}

	// Visit a parse tree produced by V4LangGrammarParser#args.
	VisitArgs(ctx *ArgsContext) interface{}

	// Visit a parse tree produced by V4LangGrammarParser#MulDiv.
	VisitMulDiv(ctx *MulDivContext) interface{}

	// Visit a parse tree produced by V4LangGrammarParser#GetValueIndex.
	VisitGetValueIndex(ctx *GetValueIndexContext) interface{}

	// Visit a parse tree produced by V4LangGrammarParser#Parens.
	VisitParens(ctx *ParensContext) interface{}

	// Visit a parse tree produced by V4LangGrammarParser#Logical.
	VisitLogical(ctx *LogicalContext) interface{}

	// Visit a parse tree produced by V4LangGrammarParser#String.
	VisitString(ctx *StringContext) interface{}

	// Visit a parse tree produced by V4LangGrammarParser#Int.
	VisitInt(ctx *IntContext) interface{}

	// Visit a parse tree produced by V4LangGrammarParser#Identifier.
	VisitIdentifier(ctx *IdentifierContext) interface{}

	// Visit a parse tree produced by V4LangGrammarParser#SetValueIndex.
	VisitSetValueIndex(ctx *SetValueIndexContext) interface{}

	// Visit a parse tree produced by V4LangGrammarParser#GetValueMatrix.
	VisitGetValueMatrix(ctx *GetValueMatrixContext) interface{}

	// Visit a parse tree produced by V4LangGrammarParser#SliceWithValues.
	VisitSliceWithValues(ctx *SliceWithValuesContext) interface{}

	// Visit a parse tree produced by V4LangGrammarParser#Equality.
	VisitEquality(ctx *EqualityContext) interface{}

	// Visit a parse tree produced by V4LangGrammarParser#Boolean.
	VisitBoolean(ctx *BooleanContext) interface{}

	// Visit a parse tree produced by V4LangGrammarParser#NewInstance.
	VisitNewInstance(ctx *NewInstanceContext) interface{}

	// Visit a parse tree produced by V4LangGrammarParser#SetValueMatrix.
	VisitSetValueMatrix(ctx *SetValueMatrixContext) interface{}

	// Visit a parse tree produced by V4LangGrammarParser#Mod.
	VisitMod(ctx *ModContext) interface{}

	// Visit a parse tree produced by V4LangGrammarParser#SliceIndex.
	VisitSliceIndex(ctx *SliceIndexContext) interface{}

	// Visit a parse tree produced by V4LangGrammarParser#TypeOf.
	VisitTypeOf(ctx *TypeOfContext) interface{}

	// Visit a parse tree produced by V4LangGrammarParser#AddSub.
	VisitAddSub(ctx *AddSubContext) interface{}

	// Visit a parse tree produced by V4LangGrammarParser#StructInstance.
	VisitStructInstance(ctx *StructInstanceContext) interface{}

	// Visit a parse tree produced by V4LangGrammarParser#Relational.
	VisitRelational(ctx *RelationalContext) interface{}

	// Visit a parse tree produced by V4LangGrammarParser#CallStmt.
	VisitCallStmt(ctx *CallStmtContext) interface{}

	// Visit a parse tree produced by V4LangGrammarParser#Nil.
	VisitNil(ctx *NilContext) interface{}

	// Visit a parse tree produced by V4LangGrammarParser#StringJoin.
	VisitStringJoin(ctx *StringJoinContext) interface{}

	// Visit a parse tree produced by V4LangGrammarParser#Float.
	VisitFloat(ctx *FloatContext) interface{}

	// Visit a parse tree produced by V4LangGrammarParser#Not.
	VisitNot(ctx *NotContext) interface{}

	// Visit a parse tree produced by V4LangGrammarParser#Atoi.
	VisitAtoi(ctx *AtoiContext) interface{}

	// Visit a parse tree produced by V4LangGrammarParser#ParseFloat.
	VisitParseFloat(ctx *ParseFloatContext) interface{}

	// Visit a parse tree produced by V4LangGrammarParser#Append.
	VisitAppend(ctx *AppendContext) interface{}

	// Visit a parse tree produced by V4LangGrammarParser#Len.
	VisitLen(ctx *LenContext) interface{}

	// Visit a parse tree produced by V4LangGrammarParser#AssignArithmetic.
	VisitAssignArithmetic(ctx *AssignArithmeticContext) interface{}

	// Visit a parse tree produced by V4LangGrammarParser#Slice.
	VisitSlice(ctx *SliceContext) interface{}

	// Visit a parse tree produced by V4LangGrammarParser#Assign.
	VisitAssign(ctx *AssignContext) interface{}

	// Visit a parse tree produced by V4LangGrammarParser#Negate.
	VisitNegate(ctx *NegateContext) interface{}

	// Visit a parse tree produced by V4LangGrammarParser#IncrementDecrement.
	VisitIncrementDecrement(ctx *IncrementDecrementContext) interface{}

	// Visit a parse tree produced by V4LangGrammarParser#Rune.
	VisitRune(ctx *RuneContext) interface{}

	// Visit a parse tree produced by V4LangGrammarParser#FuncCall.
	VisitFuncCall(ctx *FuncCallContext) interface{}

	// Visit a parse tree produced by V4LangGrammarParser#Get.
	VisitGet(ctx *GetContext) interface{}

	// Visit a parse tree produced by V4LangGrammarParser#instanceValues.
	VisitInstanceValues(ctx *InstanceValuesContext) interface{}

	// Visit a parse tree produced by V4LangGrammarParser#type.
	VisitType(ctx *TypeContext) interface{}
}
