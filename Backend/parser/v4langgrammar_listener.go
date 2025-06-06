// Code generated from V4LangGrammar.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // V4LangGrammar

import "github.com/antlr4-go/antlr/v4"

// V4LangGrammarListener is a complete listener for a parse tree produced by V4LangGrammarParser.
type V4LangGrammarListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterDcl is called when entering the dcl production.
	EnterDcl(c *DclContext)

	// EnterVarDcl is called when entering the varDcl production.
	EnterVarDcl(c *VarDclContext)

	// EnterFuncDcl is called when entering the funcDcl production.
	EnterFuncDcl(c *FuncDclContext)

	// EnterParams is called when entering the params production.
	EnterParams(c *ParamsContext)

	// EnterStructDcl is called when entering the structDcl production.
	EnterStructDcl(c *StructDclContext)

	// EnterStructBody is called when entering the structBody production.
	EnterStructBody(c *StructBodyContext)

	// EnterReceiver is called when entering the receiver production.
	EnterReceiver(c *ReceiverContext)

	// EnterSliceDcl is called when entering the sliceDcl production.
	EnterSliceDcl(c *SliceDclContext)

	// EnterSliceValues is called when entering the sliceValues production.
	EnterSliceValues(c *SliceValuesContext)

	// EnterMatrixDcl is called when entering the matrixDcl production.
	EnterMatrixDcl(c *MatrixDclContext)

	// EnterMatrixValues is called when entering the matrixValues production.
	EnterMatrixValues(c *MatrixValuesContext)

	// EnterExprStmt is called when entering the ExprStmt production.
	EnterExprStmt(c *ExprStmtContext)

	// EnterPrintStmt is called when entering the PrintStmt production.
	EnterPrintStmt(c *PrintStmtContext)

	// EnterBlockStmt is called when entering the BlockStmt production.
	EnterBlockStmt(c *BlockStmtContext)

	// EnterIfStmt is called when entering the IfStmt production.
	EnterIfStmt(c *IfStmtContext)

	// EnterSwitchStmt is called when entering the SwitchStmt production.
	EnterSwitchStmt(c *SwitchStmtContext)

	// EnterForIncrementStmt is called when entering the ForIncrementStmt production.
	EnterForIncrementStmt(c *ForIncrementStmtContext)

	// EnterForStmt is called when entering the ForStmt production.
	EnterForStmt(c *ForStmtContext)

	// EnterForRangeStmt is called when entering the ForRangeStmt production.
	EnterForRangeStmt(c *ForRangeStmtContext)

	// EnterBreakStmt is called when entering the BreakStmt production.
	EnterBreakStmt(c *BreakStmtContext)

	// EnterContinueStmt is called when entering the ContinueStmt production.
	EnterContinueStmt(c *ContinueStmtContext)

	// EnterReturnStmt is called when entering the ReturnStmt production.
	EnterReturnStmt(c *ReturnStmtContext)

	// EnterElseIfStmt is called when entering the elseIfStmt production.
	EnterElseIfStmt(c *ElseIfStmtContext)

	// EnterElseStmt is called when entering the elseStmt production.
	EnterElseStmt(c *ElseStmtContext)

	// EnterSwitchCase is called when entering the switchCase production.
	EnterSwitchCase(c *SwitchCaseContext)

	// EnterDefaultCase is called when entering the defaultCase production.
	EnterDefaultCase(c *DefaultCaseContext)

	// EnterArgs is called when entering the args production.
	EnterArgs(c *ArgsContext)

	// EnterMulDiv is called when entering the MulDiv production.
	EnterMulDiv(c *MulDivContext)

	// EnterGetValueIndex is called when entering the GetValueIndex production.
	EnterGetValueIndex(c *GetValueIndexContext)

	// EnterParens is called when entering the Parens production.
	EnterParens(c *ParensContext)

	// EnterLogical is called when entering the Logical production.
	EnterLogical(c *LogicalContext)

	// EnterString is called when entering the String production.
	EnterString(c *StringContext)

	// EnterInt is called when entering the Int production.
	EnterInt(c *IntContext)

	// EnterIdentifier is called when entering the Identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterSetValueIndex is called when entering the SetValueIndex production.
	EnterSetValueIndex(c *SetValueIndexContext)

	// EnterGetValueMatrix is called when entering the GetValueMatrix production.
	EnterGetValueMatrix(c *GetValueMatrixContext)

	// EnterSliceWithValues is called when entering the SliceWithValues production.
	EnterSliceWithValues(c *SliceWithValuesContext)

	// EnterEquality is called when entering the Equality production.
	EnterEquality(c *EqualityContext)

	// EnterBoolean is called when entering the Boolean production.
	EnterBoolean(c *BooleanContext)

	// EnterNewInstance is called when entering the NewInstance production.
	EnterNewInstance(c *NewInstanceContext)

	// EnterSetValueMatrix is called when entering the SetValueMatrix production.
	EnterSetValueMatrix(c *SetValueMatrixContext)

	// EnterMod is called when entering the Mod production.
	EnterMod(c *ModContext)

	// EnterSliceIndex is called when entering the SliceIndex production.
	EnterSliceIndex(c *SliceIndexContext)

	// EnterTypeOf is called when entering the TypeOf production.
	EnterTypeOf(c *TypeOfContext)

	// EnterAddSub is called when entering the AddSub production.
	EnterAddSub(c *AddSubContext)

	// EnterStructInstance is called when entering the StructInstance production.
	EnterStructInstance(c *StructInstanceContext)

	// EnterRelational is called when entering the Relational production.
	EnterRelational(c *RelationalContext)

	// EnterCallStmt is called when entering the CallStmt production.
	EnterCallStmt(c *CallStmtContext)

	// EnterNil is called when entering the Nil production.
	EnterNil(c *NilContext)

	// EnterStringJoin is called when entering the StringJoin production.
	EnterStringJoin(c *StringJoinContext)

	// EnterFloat is called when entering the Float production.
	EnterFloat(c *FloatContext)

	// EnterNot is called when entering the Not production.
	EnterNot(c *NotContext)

	// EnterAtoi is called when entering the Atoi production.
	EnterAtoi(c *AtoiContext)

	// EnterParseFloat is called when entering the ParseFloat production.
	EnterParseFloat(c *ParseFloatContext)

	// EnterAppend is called when entering the Append production.
	EnterAppend(c *AppendContext)

	// EnterLen is called when entering the Len production.
	EnterLen(c *LenContext)

	// EnterAssignArithmetic is called when entering the AssignArithmetic production.
	EnterAssignArithmetic(c *AssignArithmeticContext)

	// EnterSlice is called when entering the Slice production.
	EnterSlice(c *SliceContext)

	// EnterAssign is called when entering the Assign production.
	EnterAssign(c *AssignContext)

	// EnterNegate is called when entering the Negate production.
	EnterNegate(c *NegateContext)

	// EnterIncrementDecrement is called when entering the IncrementDecrement production.
	EnterIncrementDecrement(c *IncrementDecrementContext)

	// EnterRune is called when entering the Rune production.
	EnterRune(c *RuneContext)

	// EnterFuncCall is called when entering the FuncCall production.
	EnterFuncCall(c *FuncCallContext)

	// EnterGet is called when entering the Get production.
	EnterGet(c *GetContext)

	// EnterInstanceValues is called when entering the instanceValues production.
	EnterInstanceValues(c *InstanceValuesContext)

	// EnterType is called when entering the type production.
	EnterType(c *TypeContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitDcl is called when exiting the dcl production.
	ExitDcl(c *DclContext)

	// ExitVarDcl is called when exiting the varDcl production.
	ExitVarDcl(c *VarDclContext)

	// ExitFuncDcl is called when exiting the funcDcl production.
	ExitFuncDcl(c *FuncDclContext)

	// ExitParams is called when exiting the params production.
	ExitParams(c *ParamsContext)

	// ExitStructDcl is called when exiting the structDcl production.
	ExitStructDcl(c *StructDclContext)

	// ExitStructBody is called when exiting the structBody production.
	ExitStructBody(c *StructBodyContext)

	// ExitReceiver is called when exiting the receiver production.
	ExitReceiver(c *ReceiverContext)

	// ExitSliceDcl is called when exiting the sliceDcl production.
	ExitSliceDcl(c *SliceDclContext)

	// ExitSliceValues is called when exiting the sliceValues production.
	ExitSliceValues(c *SliceValuesContext)

	// ExitMatrixDcl is called when exiting the matrixDcl production.
	ExitMatrixDcl(c *MatrixDclContext)

	// ExitMatrixValues is called when exiting the matrixValues production.
	ExitMatrixValues(c *MatrixValuesContext)

	// ExitExprStmt is called when exiting the ExprStmt production.
	ExitExprStmt(c *ExprStmtContext)

	// ExitPrintStmt is called when exiting the PrintStmt production.
	ExitPrintStmt(c *PrintStmtContext)

	// ExitBlockStmt is called when exiting the BlockStmt production.
	ExitBlockStmt(c *BlockStmtContext)

	// ExitIfStmt is called when exiting the IfStmt production.
	ExitIfStmt(c *IfStmtContext)

	// ExitSwitchStmt is called when exiting the SwitchStmt production.
	ExitSwitchStmt(c *SwitchStmtContext)

	// ExitForIncrementStmt is called when exiting the ForIncrementStmt production.
	ExitForIncrementStmt(c *ForIncrementStmtContext)

	// ExitForStmt is called when exiting the ForStmt production.
	ExitForStmt(c *ForStmtContext)

	// ExitForRangeStmt is called when exiting the ForRangeStmt production.
	ExitForRangeStmt(c *ForRangeStmtContext)

	// ExitBreakStmt is called when exiting the BreakStmt production.
	ExitBreakStmt(c *BreakStmtContext)

	// ExitContinueStmt is called when exiting the ContinueStmt production.
	ExitContinueStmt(c *ContinueStmtContext)

	// ExitReturnStmt is called when exiting the ReturnStmt production.
	ExitReturnStmt(c *ReturnStmtContext)

	// ExitElseIfStmt is called when exiting the elseIfStmt production.
	ExitElseIfStmt(c *ElseIfStmtContext)

	// ExitElseStmt is called when exiting the elseStmt production.
	ExitElseStmt(c *ElseStmtContext)

	// ExitSwitchCase is called when exiting the switchCase production.
	ExitSwitchCase(c *SwitchCaseContext)

	// ExitDefaultCase is called when exiting the defaultCase production.
	ExitDefaultCase(c *DefaultCaseContext)

	// ExitArgs is called when exiting the args production.
	ExitArgs(c *ArgsContext)

	// ExitMulDiv is called when exiting the MulDiv production.
	ExitMulDiv(c *MulDivContext)

	// ExitGetValueIndex is called when exiting the GetValueIndex production.
	ExitGetValueIndex(c *GetValueIndexContext)

	// ExitParens is called when exiting the Parens production.
	ExitParens(c *ParensContext)

	// ExitLogical is called when exiting the Logical production.
	ExitLogical(c *LogicalContext)

	// ExitString is called when exiting the String production.
	ExitString(c *StringContext)

	// ExitInt is called when exiting the Int production.
	ExitInt(c *IntContext)

	// ExitIdentifier is called when exiting the Identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitSetValueIndex is called when exiting the SetValueIndex production.
	ExitSetValueIndex(c *SetValueIndexContext)

	// ExitGetValueMatrix is called when exiting the GetValueMatrix production.
	ExitGetValueMatrix(c *GetValueMatrixContext)

	// ExitSliceWithValues is called when exiting the SliceWithValues production.
	ExitSliceWithValues(c *SliceWithValuesContext)

	// ExitEquality is called when exiting the Equality production.
	ExitEquality(c *EqualityContext)

	// ExitBoolean is called when exiting the Boolean production.
	ExitBoolean(c *BooleanContext)

	// ExitNewInstance is called when exiting the NewInstance production.
	ExitNewInstance(c *NewInstanceContext)

	// ExitSetValueMatrix is called when exiting the SetValueMatrix production.
	ExitSetValueMatrix(c *SetValueMatrixContext)

	// ExitMod is called when exiting the Mod production.
	ExitMod(c *ModContext)

	// ExitSliceIndex is called when exiting the SliceIndex production.
	ExitSliceIndex(c *SliceIndexContext)

	// ExitTypeOf is called when exiting the TypeOf production.
	ExitTypeOf(c *TypeOfContext)

	// ExitAddSub is called when exiting the AddSub production.
	ExitAddSub(c *AddSubContext)

	// ExitStructInstance is called when exiting the StructInstance production.
	ExitStructInstance(c *StructInstanceContext)

	// ExitRelational is called when exiting the Relational production.
	ExitRelational(c *RelationalContext)

	// ExitCallStmt is called when exiting the CallStmt production.
	ExitCallStmt(c *CallStmtContext)

	// ExitNil is called when exiting the Nil production.
	ExitNil(c *NilContext)

	// ExitStringJoin is called when exiting the StringJoin production.
	ExitStringJoin(c *StringJoinContext)

	// ExitFloat is called when exiting the Float production.
	ExitFloat(c *FloatContext)

	// ExitNot is called when exiting the Not production.
	ExitNot(c *NotContext)

	// ExitAtoi is called when exiting the Atoi production.
	ExitAtoi(c *AtoiContext)

	// ExitParseFloat is called when exiting the ParseFloat production.
	ExitParseFloat(c *ParseFloatContext)

	// ExitAppend is called when exiting the Append production.
	ExitAppend(c *AppendContext)

	// ExitLen is called when exiting the Len production.
	ExitLen(c *LenContext)

	// ExitAssignArithmetic is called when exiting the AssignArithmetic production.
	ExitAssignArithmetic(c *AssignArithmeticContext)

	// ExitSlice is called when exiting the Slice production.
	ExitSlice(c *SliceContext)

	// ExitAssign is called when exiting the Assign production.
	ExitAssign(c *AssignContext)

	// ExitNegate is called when exiting the Negate production.
	ExitNegate(c *NegateContext)

	// ExitIncrementDecrement is called when exiting the IncrementDecrement production.
	ExitIncrementDecrement(c *IncrementDecrementContext)

	// ExitRune is called when exiting the Rune production.
	ExitRune(c *RuneContext)

	// ExitFuncCall is called when exiting the FuncCall production.
	ExitFuncCall(c *FuncCallContext)

	// ExitGet is called when exiting the Get production.
	ExitGet(c *GetContext)

	// ExitInstanceValues is called when exiting the instanceValues production.
	ExitInstanceValues(c *InstanceValuesContext)

	// ExitType is called when exiting the type production.
	ExitType(c *TypeContext)
}
