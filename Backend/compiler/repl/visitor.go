package repl

import (
	parser "compiler/parser"
	"compiler/value"
	"fmt"
	"strconv"
	"strings"

	//"log"

	"github.com/antlr4-go/antlr/v4"
)

// Visitor personalizado para recorrer el árbol de sintaxis
type ReplVisitor struct {
	parser.BaseVlangVisitor            // Embebemos el visitor generado por ANTLR
	ScopeTrace *ScopeTrace
	IfScope                 []*BaseScope
	inForLoop               bool                       // Bandera para rastrear si estamos en un bucle for
	functions               map[string]*StoredFunction // Mapa para almacenar funciones definidas
}

type StoredFunction struct {
	ParamNames []string
	ParamTypes []string
	ReturnType string
	Block      parser.IBlockContext
}

var _ parser.VlangVisitor = &ReplVisitor{} // <-- Esto asegura la interfaz
// Constructor del visitor
func NewReplVisitor() *ReplVisitor {
    scopeTrace := NewScopeTrace()
    return &ReplVisitor{
        ScopeTrace: scopeTrace,
        inForLoop:  false,
        functions:  make(map[string]*StoredFunction),
    }
}

/*
func (v *ReplVisitor) Visit(tree antlr.ParseTree) interface{} {

	switch val := tree.(type) {
	case *antlr.ErrorNodeImpl:
		log.Fatal(val.GetText())
		return nil
	default:
		return tree.Accept(v)
	}

}
*/
func (v *ReplVisitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

func (v *ReplVisitor) VisitPrograma(ctx *parser.ProgramaContext) interface{} {
    // 1. Registrar funciones y variables globales
    for _, decl := range ctx.AllDeclaraciones() {
        if decl.FuncDcl() != nil {
            v.VisitFuncDcl(decl.FuncDcl().(*parser.FuncDclContext))
        } else if decl.FuncMain() != nil {
            v.VisitFuncMain(decl.FuncMain().(*parser.FuncMainContext))
        } else {
            v.Visit(decl)
        }
    }
    // 2. Ejecutar SOLO el main
    for _, decl := range ctx.AllDeclaraciones() {
        if decl.FuncMain() != nil {
            return v.VisitFuncMain(decl.FuncMain().(*parser.FuncMainContext))
        }
    }
    return nil
}

func (v *ReplVisitor) VisitDeclaraciones(ctx *parser.DeclaracionesContext) interface{} {
	for i := 0; i < ctx.GetChildCount(); i++ {
        child := ctx.GetChild(i)
        if node, ok := child.(antlr.ParseTree); ok {
            res := node.Accept(v)
            if ret, ok := res.(ReturnValue); ok {
                return ret
            }
        }
    }
    return nil
}
func (v *ReplVisitor) VisitIf_context(ctx *parser.If_contextContext) interface{} {
	return v.Visit(ctx.IfDcl())
}
func (v *ReplVisitor) VisitFor_context(ctx *parser.For_contextContext) interface{} {
	return v.Visit(ctx.ForDcl())
}
func (v *ReplVisitor) VisitSwitch_context(ctx *parser.Switch_contextContext) interface{} {
	return v.Visit(ctx.SwitchDcl())
}
func (v *ReplVisitor) VisitWhile_context(ctx *parser.While_contextContext) interface{} {
	return v.Visit(ctx.WhileDcl())
}

func (v *ReplVisitor) VisitStmt(ctx *parser.StmtContext) interface{} {
	res := v.VisitChildren(ctx)
    if ret, ok := res.(ReturnValue); ok {
        return ret
    }
    return nil
}

func (v *ReplVisitor) VisitFuncMain(ctx *parser.FuncMainContext) interface{} {
    v.ScopeTrace.PushScope("fn_main")
	v.ScopeTrace.AddFunction("main", nil)
    defer v.ScopeTrace.PopScope()
	
    return v.Visit(ctx.Block())
}

// REVISAR
// Para el continue
func (v *ReplVisitor) VisitContinueStatement(ctx *parser.ContinueStatementContext) interface{} {
	if !v.inForLoop {
		fmt.Printf("Error: 'continue' fuera de un bucle for en la línea %d\n", ctx.GetStart().GetLine())
		return nil
	}
	return "continue" // Señal para saltar a la siguiente iteración
}

// Para el brake
func (v *ReplVisitor) VisitBreakStatement(ctx *parser.BreakStatementContext) interface{} {
	if !v.inForLoop {
		fmt.Printf("Error: 'break' fuera de un bucle for en la línea %d\n", ctx.GetStart().GetLine())
		return nil
	}
	return "break" // Señal para salir del bucle
}

// HASTA ACA
func (v *ReplVisitor) VisitBlock(ctx *parser.BlockContext) interface{} {
    // Si el scope actual ya es de función, no hagas push de "block"
    if strings.HasPrefix(v.ScopeTrace.CurrentScope.name, "fn_") || v.ScopeTrace.CurrentScope.name == "main" {
        for _, decl := range ctx.AllDeclaraciones() {
            res := v.Visit(decl)
            if ret, ok := res.(ReturnValue); ok {
                return ret
            }
        }
        return nil
    }
    // Si no, sí crea un nuevo scope "block"
    v.ScopeTrace.PushScope("block")
    defer v.ScopeTrace.PopScope()
    for _, decl := range ctx.AllDeclaraciones() {
        res := v.Visit(decl)
        if _, ok := res.(ReturnValue); ok {
            return res
        }
    }
    return nil
}

// Visitamos declaraciones de variables
func (v *ReplVisitor) VisitVariableDeclaration(ctx *parser.VariableDeclarationContext) interface{} {
	varName := ctx.ID().GetText()
	var varType string
	var valueObj value.IVOR
	var n = 1

	// 1. Detectar tipo explícito y asignar valor por defecto
	if ctx.TIPO() != nil {
		tipoText := ctx.TIPO().GetText()
		switch tipoText {
		case "int":
			varType = value.IVOR_INT
			valueObj = value.NewIntValue(0)
		case "float64":
			varType = value.IVOR_FLOAT
			valueObj = value.NewFloatValue(0.0)
		case "string":
			varType = value.IVOR_STRING
			valueObj = value.NewStringValue("")
		case "bool":
			varType = value.IVOR_BOOL
			valueObj = value.NewBoolValue(false)
		case "rune":
			varType = value.IVOR_CHARACTER
			valueObj = value.NewCharValue('\x00')
		default:
			return nil
		}
	}

	// 2. Si hay valor, evaluarlo y sobreescribir el valor por defecto
	if ctx.Expresion() != nil {
		val := v.Visit(ctx.Expresion())
		switch varType {
		case value.IVOR_INT:
			intVal, _ := strconv.Atoi(fmt.Sprint(val))
			valueObj = value.NewIntValue(intVal)
		case value.IVOR_FLOAT:
			floatVal, _ := strconv.ParseFloat(fmt.Sprint(val), 64)
			valueObj = value.NewFloatValue(floatVal)
		case value.IVOR_STRING:
			valueObj = value.NewStringValue(fmt.Sprint(val))
		case value.IVOR_BOOL:
			boolVal := false
			if fmt.Sprint(val) == "true" {
				boolVal = true
			}
			valueObj = value.NewBoolValue(boolVal)
		case value.IVOR_CHARACTER:
			switch v := val.(type) {
			case int32:
				valueObj = value.NewCharValue(rune(v))
			case string:
				runes := []rune(v)
				if len(runes) > 0 {
					valueObj = value.NewCharValue(runes[0])
				} else {
					valueObj = value.NewCharValue('\x00')
				}
			default:
				valueObj = value.NewCharValue('\x00')
			}
		}
	}

	v.ScopeTrace.AddVariable(varName, varType, valueObj, false, false, ctx.GetStart())
	if n == 0 {
		fmt.Println(varName, varType, valueObj)
	}
	//fmt.Printf("[DEBUG] Variable '%s' declarada en scope '%s'\n", varName, v.Scope.Name)

	return nil
}

func (v *ReplVisitor) VisitVariableDeclarationImmutable(ctx *parser.VariableDeclarationImmutableContext) interface{} {
	varName := ctx.ID().GetText()

	if ctx.ASSIGN() != nil && ctx.Expresion() != nil {
		variable := v.ScopeTrace.GetVariable(varName)
		if variable == nil {
			fmt.Printf("SEMANTICO: variable '%s' no declarada\n", varName)
			return nil
		}
		val := v.Visit(ctx.Expresion())

		switch variable.Type {
		case value.IVOR_INT:
			intVal, ok := val.(int)
			if !ok {
				fmt.Printf("SEMANTICO valor '%v' no es int\n", val)
				return nil
			}
			variable.Value = value.NewIntValue(intVal)
		case value.IVOR_FLOAT:
			floatVal, ok := val.(float64)
			if !ok {
				fmt.Printf("SEMANTICO: valor '%v' no es float64\n", val)
				return nil
			}
			variable.Value = value.NewFloatValue(floatVal)
		case value.IVOR_STRING:
			strVal, ok := val.(string)
			if !ok {
				fmt.Printf("SEMANTICO: valor '%v' no es string\n", val)
				return nil
			}
			variable.Value = value.NewStringValue(strVal)
		case value.IVOR_BOOL:
			boolVal, ok := val.(bool)
			if !ok {
				fmt.Printf("SEMANTICO: valor '%v' no es bool\n", val)
				return nil
			}
			variable.Value = value.NewBoolValue(boolVal)
		case value.IVOR_CHARACTER:
			charVal, ok := val.(rune)
			if !ok {
				fmt.Printf("SEMANTICO: valor '%v' no es rune\n", val)
				return nil
			}
			variable.Value = value.NewCharValue(charVal)
		default:
			// Soporte para slices
			if strings.HasPrefix(variable.Type, "slice_") {
				sliceVal, ok := val.(*value.SliceValue)
				if !ok {
					fmt.Printf("SEMANTICO: valor '%v' no es un slice válido\n", val)
					return nil
				}
				variable.Value = sliceVal
			} else {
				fmt.Printf("SEMANTICO: tipo '%s' no soportado para asignación: \n", variable.Type)
			}
		}
		return nil
	}

	fmt.Printf("SEMANTICO: declaración inválida para '%s'\n", varName)
	return nil
}
func (v *ReplVisitor) VisitValorCadena(ctx *parser.ValorCadenaContext) interface{} {
	//fmt.Println("Entrando a VisitValorCadena:", ctx.GetText())
	text := ctx.GetText()
	if len(text) >= 2 {
		return text[1 : len(text)-1]
	}
	return ""
}

func (v *ReplVisitor) VisitValorEntero(ctx *parser.ValorEnteroContext) interface{} {
	//fmt.Println("[DEBUG] VisitValorEntero:", ctx.GetText())
	val, err := strconv.Atoi(ctx.GetText())
	if err != nil {
		fmt.Printf("Error al convertir entero: %v\n", ctx.GetText())
		return 0
	}
	return val
}

func (v *ReplVisitor) VisitPrintStatement(ctx *parser.PrintStatementContext) interface{} {
	var outputs []string

	for _, expr := range ctx.AllExpresion() {
		val := v.Visit(expr)

		// Desempaquetamos si es ReturnValue
		if ret, ok := val.(ReturnValue); ok {
			val = ret.Value
		}

		switch v := val.(type) {
		case int32:
			outputs = append(outputs, string(v))
		case float64:
			if v == float64(int64(v)) {
				outputs = append(outputs, fmt.Sprintf("%.1f", v))
			} else {
				outputs = append(outputs, fmt.Sprint(v))
			}
		case *value.SliceValue:
			var elems []string
			slice := val.(*value.SliceValue)
			if slice.ElementType == "rune" || slice.ElementType == value.IVOR_CHARACTER {
				for _, e := range slice.Elements {
					if r, ok := e.Value().(rune); ok {
						elems = append(elems, fmt.Sprintf("%c", r))
					} else if i, ok := e.Value().(int32); ok {
						elems = append(elems, fmt.Sprintf("%c", i))
					} else {
						elems = append(elems, fmt.Sprint(e.Value()))
					}
				}
			} else {
				for _, e := range slice.Elements {
					elems = append(elems, fmt.Sprint(e.Value()))
				}
			}
			outputs = append(outputs, "["+strings.Join(elems, ", ")+"]")

		default:
			outputs = append(outputs, fmt.Sprint(val))
		}
	}

	fmt.Println(strings.Join(outputs, " "))
	return nil
}

func (v *ReplVisitor) VisitParentesisexpre(ctx *parser.ParentesisexpreContext) interface{} {
	//fmt.Println("🌀 Paréntesis:", ctx.GetText())
	return v.Visit(ctx.Expresion())
}

func (v *ReplVisitor) VisitCorchetesexpre(ctx *parser.CorchetesexpreContext) interface{} {
	//fmt.Println("🧱 Corchetes:", ctx.GetText())
	return v.Visit(ctx.Expresion())
}

func (v *ReplVisitor) VisitUnario(ctx *parser.UnarioContext) interface{} {
	op := ctx.GetOp().GetText()
	val := v.Visit(ctx.Expresion())

	switch op {
	case "!":
		// Solo booleanos
		if fmt.Sprint(val) == "true" {
			return false
		} else if fmt.Sprint(val) == "false" {
			return true
		} else {
			fmt.Println("Error: operador '!' solo acepta booleanos")
			return false
		}
	case "-":
		// Solo números
		if num, err := strconv.ParseFloat(fmt.Sprint(val), 64); err == nil {
			// Si era int, regresa int; si era float, regresa float
			if strings.Contains(fmt.Sprint(val), ".") {
				return -num
			} else {
				return int(-num)
			}
		} else {
			fmt.Println("Error: operador '-' solo acepta números")
			return 0
		}
	}
	return nil
}

func (v *ReplVisitor) VisitSumres(ctx *parser.SumresContext) interface{} {
	left := v.Visit(ctx.Expresion(0))
	right := v.Visit(ctx.Expresion(1))
	op := ctx.GetOp().GetText()
	var n = 1

	// Si alguno es string, concatenar
	leftStr, leftIsString := left.(string)
	rightStr, rightIsString := right.(string)
	if op == "+" && (leftIsString || rightIsString) {
		return fmt.Sprint(left) + fmt.Sprint(right)
	}

	// Si alguno es decimal, resultado decimal
	lStr := fmt.Sprint(left)
	rStr := fmt.Sprint(right)
	isLeftFloat := strings.Contains(lStr, ".")
	isRightFloat := strings.Contains(rStr, ".")

	if isLeftFloat || isRightFloat {
		leftVal, _ := strconv.ParseFloat(lStr, 64)
		rightVal, _ := strconv.ParseFloat(rStr, 64)
		switch op {
		case "+":
			return leftVal + rightVal
		case "-":
			return leftVal - rightVal
		}
	} else {
		leftVal, _ := strconv.Atoi(lStr)
		rightVal, _ := strconv.Atoi(rStr)
		switch op {
		case "+":
			return leftVal + rightVal
		case "-":
			return leftVal - rightVal
		}
	}

	if n == 0 {
		fmt.Print(leftStr, " ", rightStr, " ", op, "\n")

	}
	return nil
}

func (v *ReplVisitor) VisitMultdivmod(ctx *parser.MultdivmodContext) interface{} {
	left := v.Visit(ctx.Expresion(0))
	right := v.Visit(ctx.Expresion(1))
	op := ctx.GetOp().GetText()

	leftStr := fmt.Sprint(left)
	rightStr := fmt.Sprint(right)

	isLeftFloat := strings.Contains(leftStr, ".")
	isRightFloat := strings.Contains(rightStr, ".")

	// Si alguno es decimal, resultado decimal
	if isLeftFloat || isRightFloat {
		leftVal, _ := strconv.ParseFloat(leftStr, 64)
		rightVal, _ := strconv.ParseFloat(rightStr, 64)
		switch op {
		case "*":
			return leftVal * rightVal
		case "/":
			if rightVal == 0 {
				fmt.Println("Error: División por cero")
				return nil
			}
			return leftVal / rightVal
		case "%":
			// El módulo solo tiene sentido para enteros, pero si hay decimal, puedes retornar NaN o error
			fmt.Println("Error: El operador % no es válido para decimales")
			return nil
		}
	} else {
		leftVal, _ := strconv.Atoi(leftStr)
		rightVal, _ := strconv.Atoi(rightStr)
		switch op {
		case "*":
			return leftVal * rightVal
		case "/":
			if rightVal == 0 {
				fmt.Println("Error: División por cero")
				return nil
			}
			// División entera
			return leftVal / rightVal
		case "%":
			return leftVal % rightVal
		}
	}
	return nil
}

func (v *ReplVisitor) VisitRelacionales(ctx *parser.RelacionalesContext) interface{} {
	//fmt.Println("[DEBUG] Entrando a VisitRelacionales:", ctx.GetText())
	left := v.Visit(ctx.Expresion(0))
	right := v.Visit(ctx.Expresion(1))
	op := ctx.GetOp().GetText()

	//fmt.Printf("Comparing: %v %s %v\n", left, op, right)

	// Comparación de cadenas (lexicográfica)
	leftStr, leftIsString := left.(string)
	rightStr, rightIsString := right.(string)
	if leftIsString && rightIsString {
		switch op {
		case ">":
			return leftStr > rightStr
		case ">=":
			return leftStr >= rightStr
		case "<":
			//fmt.Printf("[DEBUG] Comparando '%s' y '%s' con operador '%s'\n", leftStr, rightStr, op)
			return leftStr < rightStr
		case "<=":
			return leftStr <= rightStr
		}
	}

	// Comparación de números (int o float)
	leftNum, leftNumErr := strconv.ParseFloat(fmt.Sprint(left), 64)
	rightNum, rightNumErr := strconv.ParseFloat(fmt.Sprint(right), 64)
	if leftNumErr == nil && rightNumErr == nil {
		switch op {
		case ">":
			return leftNum > rightNum
		case ">=":
			return leftNum >= rightNum
		case "<":
			return leftNum < rightNum
		case "<=":
			return leftNum <= rightNum
		}
	}

	// Comparación de runes (caracteres)
	if l, ok1 := left.(rune); ok1 {
		if r, ok2 := right.(rune); ok2 {
			switch op {
			case ">":
				return l > r
			case ">=":
				return l >= r
			case "<":
				return l < r
			case "<=":
				return l <= r
			}
		}
	}

	fmt.Printf("Error: comparación relacional inválida entre tipos '%T' y '%T'\n", left, right)
	return false
}

func (v *ReplVisitor) VisitIgualdad(ctx *parser.IgualdadContext) interface{} {
	left := v.Visit(ctx.Expresion(0))
	right := v.Visit(ctx.Expresion(1))
	op := ctx.GetOp().GetText()

	// Si ambos son runes (caracteres)
	if l, ok1 := left.(rune); ok1 {
		if r, ok2 := right.(rune); ok2 {
			//fmt.Printf("[DEBUG rune] Comparando '%c' (%d) y '%c' (%d) con op %s\n", l, l, r, r, op)
			switch op {
			case "==":
				return l == r
			case "!=":
				return l != r
			}
		}
	}

	// Si ambos son string, compara lexicográficamente
	leftStr, leftIsString := left.(string)
	rightStr, rightIsString := right.(string)
	if leftIsString && rightIsString {
		//fmt.Printf("[DEBUG string] Comparando \"%s\" y \"%s\" con op %s\n", leftStr, rightStr, op)
		switch op {
		case "==":
			return leftStr == rightStr
		case "!=":
			return leftStr != rightStr
		}
	}

	// Si ambos son números (int o float), compara numéricamente
	leftNum, leftNumErr := strconv.ParseFloat(fmt.Sprint(left), 64)
	rightNum, rightNumErr := strconv.ParseFloat(fmt.Sprint(right), 64)
	if leftNumErr == nil && rightNumErr == nil {
		//fmt.Printf("[DEBUG num] Comparando %v y %v con op %s\n", leftNum, rightNum, op)
		switch op {
		case "==":
			return leftNum == rightNum
		case "!=":
			return leftNum != rightNum
		}
	}

	// Si ambos son booleanos
	if (fmt.Sprint(left) == "true" || fmt.Sprint(left) == "false") &&
		(fmt.Sprint(right) == "true" || fmt.Sprint(right) == "false") {
		leftBool := fmt.Sprint(left) == "true"
		rightBool := fmt.Sprint(right) == "true"
		//fmt.Printf("[DEBUG bool] Comparando %v y %v con op %s\n", leftBool, rightBool, op)
		switch op {
		case "==":
			return leftBool == rightBool
		case "!=":
			return leftBool != rightBool
		}
	}

	fmt.Printf("Error: comparación inválida entre tipos '%T' y '%T'\n", left, right)
	return false
}

func (v *ReplVisitor) VisitId(ctx *parser.IdContext) interface{} {
	
	varName := ctx.GetText()
	variable := v.ScopeTrace.GetVariable(varName)
	//fmt.Println("[DEBUG] VisitID:", ctx.GetText())
	if variable != nil && variable.Value != nil {
		//fmt.Printf("Accediendo variable '%s', valor: %v\n", varName, variable.Value.Value())
		// 🔍 Si es slice, devolvemos el objeto completo para poder acceder a sus elementos después
		if strings.HasPrefix(variable.Type, "slice_") {
			return variable.Value // devuelve *SliceValue
		}
		return variable.Value.Value()
	}
	fmt.Printf("Variable '%s' no encontrada\n", varName)
	return "<undef>"
}

func (v *ReplVisitor) VisitIncredecr(ctx *parser.IncredecrContext) interface{} {
	//fmt.Println("🔼 Incremento/Decremento:", ctx.GetText())
	child := ctx.GetChild(0)

	switch node := child.(type) {
	case *parser.IncrementoContext:
		return v.VisitIncremento(node)
	case *parser.DecrementoContext:
		return v.VisitDecremento(node)
	default:
		fmt.Println("❌ Error: incremento/decremento no reconocido")
		return nil
	}
}

func (v *ReplVisitor) VisitIncremento(ctx *parser.IncrementoContext) interface{} {
	varName := ctx.ID().GetText()
	variable := v.ScopeTrace.GetVariable(varName)
	if variable == nil {
		fmt.Printf("Error: variable '%s' no declarada\n", varName)
		return nil
	}

	val, ok := variable.Value.Value().(int)
	if !ok {
		fmt.Printf("Error: variable '%s' no es int\n", varName)
		return nil
	}

	oldVal := val
	variable.Value = value.NewIntValue(val + 1)
	return oldVal
}

func (v *ReplVisitor) VisitDecremento(ctx *parser.DecrementoContext) interface{} {
	varName := ctx.ID().GetText()
	variable := v.ScopeTrace.GetVariable(varName)
	if variable == nil {
		fmt.Printf("Error: variable '%s' no declarada\n", varName)
		return nil
	}

	val, ok := variable.Value.Value().(int)
	if !ok {
		fmt.Printf("Error: variable '%s' no es int\n", varName)
		return nil
	}

	oldVal := val
	variable.Value = value.NewIntValue(val - 1)
	return oldVal
}

func (v *ReplVisitor) VisitChildren(node antlr.RuleNode) interface{} {
    for i := 0; i < node.GetChildCount(); i++ {
        child := node.GetChild(i)
        if childNode, ok := child.(antlr.ParseTree); ok {
            res := childNode.Accept(v)
            if ret, ok := res.(ReturnValue); ok {
                return ret
            }
        }
    }
    return nil
}
func (v *ReplVisitor) VisitExpdotexp1(ctx *parser.Expdotexp1Context) interface{} {
	//fmt.Println("📌 Acceso punto ID.ID:", ctx.GetText())
	return nil
}

func (v *ReplVisitor) VisitExpdotexp(ctx *parser.ExpdotexpContext) interface{} {
	//fmt.Println("📌 Acceso punto ID.exp:", ctx.GetText())
	return v.Visit(ctx.Expresion())
}

func (v *ReplVisitor) VisitAsignacionLUEGO(ctx *parser.AsignacionLUEGOContext) interface{} {
	varName := ctx.ID().GetText()
	variable := v.ScopeTrace.GetVariable(varName)
	if variable == nil {
		fmt.Printf("Variable '%s' no declarada\n", varName)
		return nil
	}
	newVal := v.Visit(ctx.Expresion())
	// Aquí podrías hacer conversión de tipo según variable.Type
	switch variable.Type {
	case value.IVOR_INT:
		intVal, _ := strconv.Atoi(fmt.Sprint(newVal))
		variable.Value = value.NewIntValue(intVal)
	case value.IVOR_FLOAT:
		floatVal, _ := strconv.ParseFloat(fmt.Sprint(newVal), 64)
		variable.Value = value.NewFloatValue(floatVal)
	case value.IVOR_STRING:
		variable.Value = value.NewStringValue(fmt.Sprint(newVal))
	case value.IVOR_BOOL:
		boolVal := false
		if fmt.Sprint(newVal) == "true" {
			boolVal = true
		}
		variable.Value = value.NewBoolValue(boolVal)
	case value.IVOR_CHARACTER:
		strVal := fmt.Sprint(newVal)
		if len(strVal) > 0 {
			variable.Value = value.NewCharValue(rune(strVal[0]))
		} else {
			variable.Value = value.NewCharValue('\x00')
		}
	}
	return nil
}
func (v *ReplVisitor) VisitValorexpr(ctx *parser.ValorexprContext) interface{} {
	return v.Visit(ctx.Valor())
}
func (v *ReplVisitor) VisitValorDecimal(ctx *parser.ValorDecimalContext) interface{} {
	val, err := strconv.ParseFloat(ctx.GetText(), 64)
	if err != nil {
		fmt.Printf("Error al convertir decimal: %v\n", ctx.GetText())
		return 0.0
	}
	return val
}

func (v *ReplVisitor) VisitValorBooleano(ctx *parser.ValorBooleanoContext) interface{} {
	text := ctx.GetText()
    return text == "true"
}

func (v *ReplVisitor) VisitValorCaracter(ctx *parser.ValorCaracterContext) interface{} {
	text := ctx.GetText()
	if len(text) >= 3 {
		return rune([]rune(text[1 : len(text)-1])[0])
	}
	return rune(0)
}

// NO ELIMINES ESTA FUNCION SI NO TE CARGAS TODO LITERALMENTE xD
func (v *ReplVisitor) VisitExpresionStatement(ctx *parser.ExpresionStatementContext) interface{} {
    res := v.Visit(ctx.Expresion())
    if ret, ok := res.(ReturnValue); ok {
        return ret
    }
    return res
}
func (v *ReplVisitor) VisitTransfersentence(ctx *parser.TransfersentenceContext) interface{} {
    res := v.VisitChildren(ctx)
    if ret, ok := res.(ReturnValue); ok {
        return ret
    }
    return nil
}

func (v *ReplVisitor) VisitIMCPLICIT(ctx *parser.IMCPLICITContext) interface{} {
	varName := ctx.ID().GetText()
	variable := v.ScopeTrace.GetVariable(varName)
	if variable == nil {
		fmt.Printf("Error: variable '%s' no declarada\n", varName)
		return nil
	}
	val := v.Visit(ctx.Expresion())
	op := ctx.GetOp().GetText()

	switch variable.Type {
	case value.IVOR_INT:
		current := variable.Value.Value().(int)
		add, _ := strconv.Atoi(fmt.Sprint(val))
		if op == "+=" {
			variable.Value = value.NewIntValue(current + add)
		} else if op == "-=" {
			variable.Value = value.NewIntValue(current - add)
		}
	case value.IVOR_FLOAT:
		current := variable.Value.Value().(float64)
		add, _ := strconv.ParseFloat(fmt.Sprint(val), 64)
		if op == "+=" {
			variable.Value = value.NewFloatValue(current + add)
		} else if op == "-=" {
			variable.Value = value.NewFloatValue(current - add)
		}
	default:
		fmt.Printf("Error: operación %s no soportada para tipo %s\n", op, variable.Type)
	}
	return nil
}

func (v *ReplVisitor) VisitOPERADORESLOGICOS(ctx *parser.OPERADORESLOGICOSContext) interface{} {
	left := v.Visit(ctx.Expresion(0))
	right := v.Visit(ctx.Expresion(1))
	op := ctx.GetOp().GetText()

	leftBool, lok := left.(bool)
    rightBool, rok := right.(bool)
    if !lok || !rok {
        fmt.Printf("Error: Operador lógico solo acepta booleanos, recibidos: %T y %T\n", left, right)
        return false
    }
	switch op {
	case "&&":
		return leftBool && rightBool
	case "||":
		return leftBool || rightBool
	}
	return false
}

func (v *ReplVisitor) VisitControlStatement(ctx *parser.ControlStatementContext) interface{} {
    res := v.VisitChildren(ctx)
    if ret, ok := res.(ReturnValue); ok {
        return ret
    }
    return nil
}

// condicionales
func (v *ReplVisitor) VisitIfDcl(ctx *parser.IfDclContext) interface{} {
    v.ScopeTrace.PushScope("IF")
    defer v.ScopeTrace.PopScope()
    condVal := v.Visit(ctx.Expresion())
    condBool := fmt.Sprint(condVal) == "true"

    if condBool {
        for _, decl := range ctx.AllDeclaraciones() {
            res := v.Visit(decl)
            if ret, ok := res.(ReturnValue); ok {
                return ret
            }
        }
        return nil
    }

    for _, elseIf := range ctx.AllElseIfDcl() {
        elseIfCond := v.Visit(elseIf.Expresion())
        if fmt.Sprint(elseIfCond) == "true" {
            for _, decl := range elseIf.AllDeclaraciones() {
                res := v.Visit(decl)
                if ret, ok := res.(ReturnValue); ok {
                    return ret
                }
            }
            return nil
        }
    }

    if ctx.ElseCondicional() != nil {
        for _, decl := range ctx.ElseCondicional().AllDeclaraciones() {
            res := v.Visit(decl)
            if ret, ok := res.(ReturnValue); ok {
                return ret
            }
        }
    }

    return nil
}

// El For tiene error no repite el ciclo :') pero si lo lee el programa
func (v *ReplVisitor) VisitForClasico(ctx *parser.ForClasicoContext) interface{} {
	v.ScopeTrace.PushScope("FOR_CLASICO")
    defer v.ScopeTrace.PopScope()

	v.inForLoop = true // 🔹 Habilita el uso de "continue" y "break"
	defer func() {
		v.inForLoop = false // 🔹 Restaura estado
	}()

	varName := ctx.Asignacion().ID().GetText()
	initVal := v.Visit(ctx.Asignacion().Expresion())

	varType := value.IVOR_INT
	varVal := value.NewIntValue(0)

	switch val := initVal.(type) {
	case int:
		varType = value.IVOR_INT
		varVal = value.NewIntValue(val)
	case float64:
		varType = value.IVOR_FLOAT
		varVal = value.NewFloatValue(val)
	case string:
		varType = value.IVOR_STRING
		varVal = value.NewStringValue(val)
	case bool:
		varType = value.IVOR_BOOL
		varVal = value.NewBoolValue(val)
	case rune:
		varType = value.IVOR_CHARACTER
		varVal = value.NewCharValue(val)
	default:
		fmt.Printf("❌ Tipo no soportado para la variable '%s'\n", varName)
	}

	v.ScopeTrace.AddVariable(varName, varType, varVal, false, false, ctx.GetStart())

	for {
        condVal := v.Visit(ctx.Expresion())
        if fmt.Sprint(condVal) != "true" {
            break
        }

        // ⬇️ Aquí usa PushScope y PopScope para el scope de la iteración
        v.ScopeTrace.PushScope("FOR_ITER")
        for _, decl := range ctx.Block().AllDeclaraciones() {
            res := v.Visit(decl)
            if str, ok := res.(string); ok {
                if str == "continue" {
                    break // salto a incremento y siguiente iteración
                } else if str == "break" {
                    v.ScopeTrace.PopScope()
                    return nil // salimos del for
                }
            }
        }
        v.ScopeTrace.PopScope()

        if ctx.Stmt() != nil {
            v.Visit(ctx.Stmt()) // Ejecutar incremento (como i++)
        }
    }

	return nil
}

func (v *ReplVisitor) VisitForCondicionUnica(ctx *parser.ForCondicionUnicaContext) interface{} {
    v.ScopeTrace.PushScope("FOR_SIMPLE")
    defer v.ScopeTrace.PopScope()

    v.inForLoop = true
    defer func() { v.inForLoop = false }()

    for {
        condVal := v.Visit(ctx.Expresion())
        if fmt.Sprint(condVal) != "true" {
            break
        }

        v.ScopeTrace.PushScope("FOR_ITER")
        for _, decl := range ctx.Block().AllDeclaraciones() {
            res := v.Visit(decl)
            if str, ok := res.(string); ok {
                if str == "continue" {
                    break
                } else if str == "break" {
                    v.ScopeTrace.PopScope()
                    return nil
                }
            }
        }
        v.ScopeTrace.PopScope()
    }

    return nil
}

func (v *ReplVisitor) VisitSwitchDcl(ctx *parser.SwitchDclContext) interface{} {
	//fmt.Println("========== [DEBUG SWITCH_DCL] ==========")
	switchVal := v.Visit(ctx.Expresion())
	valStr := fmt.Sprint(switchVal)

	matchFound := false

	// Buscar coincidencia en los case
	for _, caseCtx := range ctx.AllCaseBlock() {
		caseVal := v.Visit(caseCtx.Expresion())
		if fmt.Sprint(caseVal) == valStr {
			matchFound = true
			for _, decl := range caseCtx.AllDeclaraciones() {
				v.Visit(decl)
			}
			break // solo el primer match se ejecuta
		}
	}

	// Si no se encontró ningún case, ejecutar default
	if !matchFound && ctx.DefaultBlock() != nil {
		for _, decl := range ctx.DefaultBlock().AllDeclaraciones() {
			v.Visit(decl)
		}
	}
	return nil
}

// Slices
func (v *ReplVisitor) VisitSliceEmptyDeclaration(ctx *parser.SliceEmptyDeclarationContext) interface{} {
	varName := ctx.ID().GetText()
	tipo := ctx.SliceTipo().TIPO().GetText()
	sliceType := "slice_" + tipo
	sliceVal := value.NewSliceValue(tipo, []value.IVOR{})
	v.ScopeTrace.AddVariable(varName, sliceType, sliceVal, false, false, ctx.GetStart())
	return nil
}

func (v *ReplVisitor) VisitSliceInitDeclaration(ctx *parser.SliceInitDeclarationContext) interface{} {
	varName := ctx.ID().GetText()
	tipo := ctx.SliceTipo().TIPO().GetText()
	sliceType := "slice_" + tipo

	var elements []value.IVOR
	if ctx.SliceInit().ListaExpresiones() != nil {
		for _, expr := range ctx.SliceInit().ListaExpresiones().AllExpresion() {
			val := v.Visit(expr)
			switch tipo {
			case "int":
				valInt := int(value.ToFloat(val)) // conversión segura
				elements = append(elements, value.NewIntValue(valInt))
			case "float64":
				elements = append(elements, value.NewFloatValue(value.ToFloat(val)))
			case "string":
				// val puede ser string o cualquier cosa convertible a string
				strVal, ok := val.(string)
				if !ok {
					strVal = fmt.Sprint(val)
				}
				//fmt.Printf("Agregando al slice de string: %s\n", strVal)
				elements = append(elements, value.NewStringValue(strVal))
				//fmt.Println("Elementos del slice:", elements)
			case "bool":
				elements = append(elements, value.NewBoolValue(fmt.Sprint(val) == "true"))
			case "rune":
				switch v := val.(type) {
				case int32:
					elements = append(elements, value.NewCharValue(rune(v)))
				case string:
					runes := []rune(v)
					if len(runes) > 0 {
						elements = append(elements, value.NewCharValue(runes[0]))
					}
				case int:
					elements = append(elements, value.NewCharValue(rune(v)))
				default:
					// Si no es ninguno de los anteriores, intenta convertir a string y tomar el primer rune
					runes := []rune(fmt.Sprint(v))
					if len(runes) > 0 {
						elements = append(elements, value.NewCharValue(runes[0]))
					}
				}
			}
		}
	}
	sliceVal := value.NewSliceValue(tipo, elements)
	v.ScopeTrace.AddVariable(varName, sliceType, sliceVal, false, false, ctx.GetStart())
	/*
			por si la misma palabra se quiere cambiar de valor dentro del slice
		    existing := v.ScopeTrace.GetVariable(varName)
		    if existing != nil {
		        existing.Type = sliceType
		        existing.Value = sliceVal
		    } else {
		        v.ScopeTrace.AddVariable(varName, sliceType, sliceVal, false, false, ctx.GetStart())
		    }
	*/
	return nil
}

func (v *ReplVisitor) VisitSliceAssignment(ctx *parser.SliceAssignmentContext) interface{} {
	left := ctx.ID(0).GetText()
	right := ctx.ID(1).GetText()

	sliceRight := v.ScopeTrace.GetVariable(right)
	if sliceRight == nil || !strings.HasPrefix(sliceRight.Type, "slice_") {
		fmt.Printf("SEMANTICO: '%s' no es un slice válido\n", right)
		return nil
	}

	sliceType := sliceRight.Type
	copyVal := sliceRight.Value.Copy() // ← asegúrate de que el método Copy exista
	v.ScopeTrace.AddVariable(left, sliceType, copyVal, false, false, ctx.GetStart())
	return nil
}

// Indexof Slice no funciona imprime nil
func (v *ReplVisitor) VisitLlamadaFuncion(ctx *parser.LlamadaFuncionContext) interface{} {
	funcName := ctx.GetStart().GetText()
    //fmt.Printf("[LLAMADAFUNCION] Llamando a función '%s'\n", funcName)

    // 🔹 Primero verificamos si es una función definida por el usuario
    if userFunc, exists := v.functions[funcName]; exists {
        // Evaluar los argumentos enviados
        argVals := []interface{}{}
        if len(ctx.AllExpresion()) > 0 {
            for _, expr := range ctx.AllExpresion() {
                val := v.Visit(expr)
                //fmt.Printf("[LLAMADAFUNCION]  Argumento #%d: %v (tipo: %T)\n", i, val, val)
                argVals = append(argVals, val)
            }
        }

        // Validar cantidad de parámetros
        if len(argVals) != len(userFunc.ParamNames) {
            fmt.Printf("[LLAMADAFUNCION] ❌ La función '%s' esperaba %d argumentos, recibió %d\n",
                funcName, len(userFunc.ParamNames), len(argVals))
            return nil
        }

        // Crear nuevo scope para la función
        //fmt.Printf("[LLAMADAFUNCION]  PushScope(func_%s)\n", funcName)
        v.ScopeTrace.PushScope("func_" + funcName)
        defer func() {
            //fmt.Printf("[LLAMADAFUNCION]  PopScope(func_%s)\n", funcName)
            v.ScopeTrace.PopScope()
        }()

        // Declarar parámetros en el nuevo scope
        for i, paramName := range userFunc.ParamNames {
            ivorVal := wrapToIVOR(argVals[i], userFunc.ParamTypes[i])
            //fmt.Printf("[LLAMADAFUNCION]  Declarando parámetro '%s' tipo '%s' valor '%v'\n", paramName, userFunc.ParamTypes[i], ivorVal)
            v.ScopeTrace.AddVariable(paramName, userFunc.ParamTypes[i], ivorVal, false, false, ctx.GetStart())
        }

        // Ejecutar el bloque de la función
        //fmt.Printf("[LLAMADAFUNCION]  Ejecutando bloque de la función '%s'\n", funcName)
        res := v.Visit(userFunc.Block)
        //fmt.Printf("[LLAMADAFUNCION]  Resultado bruto del bloque: %v (tipo: %T)\n", res, res)

        // Si tiene retorno
        if ret, ok := res.(ReturnValue); ok {
            //fmt.Printf("[LLAMADAFUNCION]  ReturnValue detectado: %v (tipo: %T)\n", ret.Value, ret.Value)
            return ret.Value
        }
        //fmt.Printf("[LLAMADAFUNCION]  No hubo return explícito, devolviendo nil\n")
        return nil
    }

	// 🔹 Si no es función de usuario, manejamos funciones nativas
	switch funcName {
	case "indexOf":
		args := ctx.AllExpresion()
		if len(args) != 2 {
			fmt.Printf("SEMANTICO: indexOf espera 2 argumentos (slice, valor), recibidos: %d\n", len(args))
			return -1
		}
		sliceArg := v.Visit(args[0])
		searchVal := v.Visit(args[1])
		slice, ok := sliceArg.(*value.SliceValue)
		if !ok {
			fmt.Println("SEMANTICO: el primer argumento de indexOf debe ser un slice")
			return -1
		}
		for i, elem := range slice.Elements {
			if fmt.Sprint(elem.Value()) == fmt.Sprint(searchVal) {
				return i
			}
		}
		return -1

	case "join":
		args := ctx.AllExpresion()
		if len(args) != 2 {
			fmt.Printf("SEMANTICO: join espera 2 argumentos (slice, separador), recibidos: %d\n", len(args))
			return ""
		}
		sliceArg := v.Visit(args[0])
		sepArg := v.Visit(args[1])
		slice, ok := sliceArg.(*value.SliceValue)
		if !ok {
			fmt.Println("SEMANTICO: el primer argumento de join debe ser un slice")
			return ""
		}
		if slice.ElementType != "string" {
			fmt.Println("SEMANTICO: join solo soporta slices de string")
			return ""
		}
		sep, ok := sepArg.(string)
		if !ok {
			fmt.Println("SEMANTICO: el separador de join debe ser un string")
			return ""
		}
		var strElems []string
		for _, elem := range slice.Elements {
			strVal, ok := elem.Value().(string)
			if !ok {
				fmt.Println("SEMANTICO: todos los elementos del slice deben ser string")
				return ""
			}
			strElems = append(strElems, strVal)
		}
		return strings.Join(strElems, sep)

	case "len":
		args := ctx.AllExpresion()
		if len(args) != 1 {
			fmt.Printf("SEMANTICO: len espera 1 argumento (slice), recibidos: %d\n", len(args))
			return -1
		}
		sliceArg := v.Visit(args[0])
		slice, ok := sliceArg.(*value.SliceValue)
		if !ok {
			fmt.Println("SEMANTICO: el argumento de len debe ser un slice")
			return -1
		}
		return len(slice.Elements)

	case "append":
		args := ctx.AllExpresion()
		if len(args) < 2 {
			fmt.Println("SEMANTICO: append espera al menos 2 argumentos (slice, elemento1, ...)")
			return nil
		}
		sliceArg := v.Visit(args[0])
		slice, ok := sliceArg.(*value.SliceValue)
		if !ok {
			fmt.Println("SEMANTICO: el primer argumento de append debe ser un slice")
			return nil
		}
		newElements := make([]value.IVOR, len(slice.Elements))
		copy(newElements, slice.Elements)
		for i := 1; i < len(args); i++ {
			elem := v.Visit(args[i])
			// Normalizamos tipos aquí (idéntico a lo que ya tienes)
			switch e := elem.(type) {
			case int:
				elem = value.NewIntValue(e)
			case float64:
				elem = value.NewFloatValue(e)
			case string:
				elem = value.NewStringValue(e)
			case bool:
				elem = value.NewBoolValue(e)
			case int32:
				elem = value.NewCharValue(e)
			}
			tipoSlice := slice.ElementType
			tipoElem := ""
			switch elem.(type) {
			case *value.IntValue:
				tipoElem = "int"
			case *value.FloatValue:
				tipoElem = "float64"
			case *value.StringValue:
				tipoElem = "string"
			case *value.BoolValue:
				tipoElem = "bool"
			case *value.CharValue:
				tipoElem = "rune"
			default:
				fmt.Printf("SEMANTICO: tipo de elemento no soportado para append: %T\n", elem)
				return nil
			}
			if tipoElem != tipoSlice {
				fmt.Printf("SEMANTICO: tipo del elemento '%s' no coincide con el tipo del slice '%s'\n", tipoElem, tipoSlice)
				return nil
			}
			newElements = append(newElements, elem.(value.IVOR))
		}
		return value.NewSliceValue(slice.ElementType, newElements)

	default:
		fmt.Printf("SEMANTICO: función '%s' no implementada\n", funcName)
		return nil
	}
}

// casteos
func (v *ReplVisitor) VisitVariableCastDeclaration(ctx *parser.VariableCastDeclarationContext) interface{} {
	varName := ctx.ID().GetText()
	variable := v.ScopeTrace.GetVariable(varName)
	if variable == nil {
		fmt.Printf("SEMANTICO: variable '%s' no declarada\n", varName)
		return nil
	}

	castType := ctx.CASTEOS().GetText()
	val := v.Visit(ctx.Expresion())

	switch castType {
	case "Atoi":
		// Solo se puede convertir de string a int
		strVal, ok := val.(string)
		if !ok {
			fmt.Printf("SEMANTICO: Atoi solo puede convertir strings, recibido: %T\n", val)
			return nil
		}
		// Verifica que no sea decimal
		if strings.Contains(strVal, ".") {
			fmt.Printf("SEMANTICO: Atoi no puede convertir decimales: '%s'\n", strVal)
			return nil
		}
		intVal, err := strconv.Atoi(strVal)
		if err != nil {
			fmt.Printf("SEMANTICO: Atoi error al convertir '%s' a int\n", strVal)
			return nil
		}
		// Solo permite si la variable es int
		if variable.Type != value.IVOR_INT {
			fmt.Printf("SEMANTICO: solo se puede asignar int a variable '%s' de tipo %s\n", varName, variable.Type)
			return nil
		}
		variable.Value = value.NewIntValue(intVal)
		return nil
	case "parseFloat":
		// Solo se puede convertir de string a float64
		strVal, ok := val.(string)
		if !ok {
			fmt.Printf("SEMANTICO: parseFloat solo puede convertir strings, recibido: %T\n", val)
			return nil
		}
		floatVal, err := strconv.ParseFloat(strVal, 64)
		if err != nil {
			fmt.Printf("SEMANTICO: parseFloat error al convertir '%s' a float64\n", strVal)
			return nil
		}
		// Solo permite si la variable es float64
		if variable.Type != value.IVOR_FLOAT {
			fmt.Printf("SEMANTICO: solo se puede asignar float64 a variable '%s' de tipo %s\n", varName, variable.Type)
			return nil
		}
		variable.Value = value.NewFloatValue(floatVal)
		return nil
	default:
		fmt.Printf("SEMANTICO: casteo '%s' no soportado\n", castType)
		return nil
	}
}

func (v *ReplVisitor) VisitLlamadaFuncionExpr(ctx *parser.LlamadaFuncionExprContext) interface{} {
	return v.Visit(ctx.LlamadaFuncion())
}

func (v *ReplVisitor) VisitPARAPRINTSLICE(ctx *parser.PARAPRINTSLICEContext) interface{} {
	varName := ctx.ID().GetText()
	index := v.Visit(ctx.Expresion())
	variable := v.ScopeTrace.GetVariable(varName)
	if variable == nil || !strings.HasPrefix(variable.Type, "slice_") {
		fmt.Printf("SEMANTICO: '%s' no es un slice válido\n", varName)
		return nil
	}
	slice := variable.Value.(*value.SliceValue)
	idx, ok := index.(int)
	if !ok || idx < 0 || idx >= len(slice.Elements) {
		fmt.Printf("SEMANTICO: índice fuera de rango para slice '%s'\n", varName)
		return nil
	}
	return slice.Elements[idx].Value()
}

func (v *ReplVisitor) VisitSliceAssignmentIndex(ctx *parser.SliceAssignmentIndexContext) interface{} {
	varName := ctx.ID().GetText()
	index := v.Visit(ctx.Expresion(0))
	newVal := v.Visit(ctx.Expresion(1))

	variable := v.ScopeTrace.GetVariable(varName)
	if variable == nil || !strings.HasPrefix(variable.Type, "slice_") {
		fmt.Printf("SEMANTICO: '%s' no es un slice válido\n", varName)
		return nil
	}
	slice := variable.Value.(*value.SliceValue)
	idx, ok := index.(int)
	if !ok || idx < 0 || idx >= len(slice.Elements) {
		fmt.Printf("SEMANTICO: índice fuera de rango para slice '%s'\n", varName)
		return nil
	}

	// Verifica el tipo del nuevo valor
	tipoSlice := slice.ElementType
	tipoVal := ""
	switch newVal.(type) {
	case int, *value.IntValue:
		tipoVal = "int"
	case float64, *value.FloatValue:
		tipoVal = "float64"
	case string, *value.StringValue:
		tipoVal = "string"
	case bool, *value.BoolValue:
		tipoVal = "bool"
	case int32, *value.CharValue:
		tipoVal = "rune"
	default:
		fmt.Printf("SEMANTICO: tipo de valor no soportado para asignación en slice: %T\n", newVal)
		return nil
	}
	if tipoVal != tipoSlice {
		fmt.Printf("SEMANTICO: tipo del valor '%s' no coincide con el tipo del slice '%s'\n", tipoVal, tipoSlice)
		return nil
	}

	// Normaliza el valor a IVOR si es primitivo
	switch v := newVal.(type) {
	case int:
		newVal = value.NewIntValue(v)
	case float64:
		newVal = value.NewFloatValue(v)
	case string:
		newVal = value.NewStringValue(v)
	case bool:
		newVal = value.NewBoolValue(v)
	case int32:
		newVal = value.NewCharValue(v)
	}

	slice.Elements[idx] = newVal.(value.IVOR)
	return nil
}

// funciones
func (v *ReplVisitor) VisitFuncDcl(ctx *parser.FuncDclContext) interface{} {
	funcName := ctx.ID().GetText()
	v.ScopeTrace.PushScope("fn_" + funcName)
    defer v.ScopeTrace.PopScope()
	

	// Validación de duplicado
	if _, exists := v.functions[funcName]; exists || v.ScopeTrace.GetVariable(funcName) != nil {
		fmt.Printf("SEMANTICO: El nombre '%s' ya está en uso\n", funcName)
		return nil
	}

	// Extraer parámetros
	paramNames := []string{}
	paramTypes := []string{}

	if ctx.ParametrosFormales() != nil {
		for _, param := range ctx.ParametrosFormales().AllParametro() {
			pName := param.ID().GetText()
			pType := param.TIPO().GetText()

			// Duplicados
			for _, existing := range paramNames {
				if existing == pName {
					fmt.Printf("SEMANTICO: Parámetro duplicado '%s'\n", pName)
					return nil
				}
			}

			paramNames = append(paramNames, pName)
			paramTypes = append(paramTypes, pType)
		}
	}

	var returnType string
	if ctx.TIPO() != nil {
		returnType = ctx.TIPO().GetText()
	}

	v.functions[funcName] = &StoredFunction{
		ParamNames: paramNames,
		ParamTypes: paramTypes,
		ReturnType: returnType,
		Block:      ctx.Block(),
	}

	// También agrégala al scope para la tabla de símbolos
    v.ScopeTrace.AddFunction(funcName, nil) // Puedes pasar nil o un objeto si tienes uno
	return nil
}

func (v *ReplVisitor) VisitFuncCall(ctx *parser.FuncCallContext) interface{} {
	funcName := ctx.ID().GetText()
	fmt.Printf("[FUNCALL] Llamando a función '%s'\n", funcName)
	fn, exists := v.functions[funcName]
	if !exists {
		fmt.Printf("SEMANTICO: Función '%s' no existe\n", funcName)
		return nil
	}

	// Evaluar argumentos
	argVals := []interface{}{}
	if ctx.ParametrosReales() != nil {
		for i, expr := range ctx.ParametrosReales().AllExpresion() {
			val := v.Visit(expr)
			fmt.Printf("[FUNCALL]  Argumento #%d: %v (tipo: %T)\n", i, val, val)
			argVals = append(argVals, val)
		}
	}
    fmt.Printf("[FUNCALL]  Argumentos evaluados: %v\n", argVals)
	if len(argVals) != len(fn.ParamNames) {
		fmt.Printf("SEMANTICO: La función '%s' esperaba %d argumentos, recibió %d\n",
			funcName, len(fn.ParamNames), len(argVals))
		return nil
	}

	// Crear nuevo entorno
	fmt.Printf("[FUNCALL]  PushScope(func_%s)\n", funcName)
	v.ScopeTrace.PushScope("func_" + funcName)
	defer v.ScopeTrace.PopScope()

	// Declarar parámetros
	for i, name := range fn.ParamNames {
		val := argVals[i]
		// Aquí puedes convertir val a IVOR si quieres tipo fuerte
		ivorVal := wrapToIVOR(val, fn.ParamTypes[i])
		fmt.Printf("[FUNCALL]  Declarando parámetro '%s' tipo '%s' valor '%v'\n", name, fn.ParamTypes[i], ivorVal)
		v.ScopeTrace.AddVariable(name, fn.ParamTypes[i], ivorVal, false, false, ctx.GetStart())
	}

	// Ejecutar función
	fmt.Printf("[FUNCALL]  Ejecutando bloque de la función '%s'\n", funcName)
	res := v.Visit(fn.Block)
    fmt.Printf("[FUNCALL]  Resultado bruto del bloque: %v (tipo: %T)\n", res, res)
	// Siempre revisa si hay ReturnValue, aunque no tenga tipo explícito
    if ret, ok := res.(ReturnValue); ok {
		fmt.Printf("[FUNCALL]  ReturnValue detectado: %v (tipo: %T)\n", ret.Value, ret.Value)
        return ret.Value
    }
	fmt.Printf("[FUNCALL]  No hubo return explícito, devolviendo nil\n")
	return nil
}

func wrapToIVOR(val interface{}, tipo string) value.IVOR {
	switch tipo {
	case "int":
		intVal, _ := strconv.Atoi(fmt.Sprint(val))
		return value.NewIntValue(intVal)
	case "float64":
		floatVal, _ := strconv.ParseFloat(fmt.Sprint(val), 64)
		return value.NewFloatValue(floatVal)
	case "string":
		return value.NewStringValue(fmt.Sprint(val))
	case "bool":
		return value.NewBoolValue(fmt.Sprint(val) == "true")
	case "rune":
		str := fmt.Sprint(val)
		if len(str) > 0 {
			return value.NewCharValue(rune(str[0]))
		}
		return value.NewCharValue('\x00')
	default:
		fmt.Printf("SEMANTICO: Tipo '%s' no soportado\n", tipo)
		return nil
	}
}

type ReturnValue struct {
	Value interface{}
}

func (v *ReplVisitor) VisitReturnStatement(ctx *parser.ReturnStatementContext) interface{} {
    //fmt.Printf("[RETURN] Encontrado return en línea %d\n", ctx.GetStart().GetLine())
    if ctx.Expresion() != nil {
        val := v.Visit(ctx.Expresion())
        //fmt.Printf("[RETURN]  Valor de retorno: %v (tipo: %T)\n", val, val)
        return ReturnValue{Value: val}
    }
    //fmt.Printf("[RETURN]  Return sin valor (nil)\n")
    return ReturnValue{Value: nil}
}