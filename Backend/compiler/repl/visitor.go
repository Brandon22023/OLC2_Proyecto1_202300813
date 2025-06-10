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
	Scope                   *BaseScope //este servira para obtener el valor de las variables declaradas xd
	GlobalScope             *BaseScope // sera en si el entorno global
	IfScope                 []*BaseScope
	inForLoop               bool // Bandera para rastrear si estamos en un bucle for
}

var _ parser.VlangVisitor = &ReplVisitor{} // <-- Esto asegura la interfaz
// Constructor del visitor
func NewReplVisitor() *ReplVisitor {
	global := NewBaseScope("GLOBAL", nil)
	return &ReplVisitor{
		Scope:       global,
		GlobalScope: global,
		inForLoop:   false,
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
	//fmt.Println("Entrando a VisitPrograma")
	for _, stmt := range ctx.AllFuncMain() {
		//fmt.Printf("Tipo de funcMain: %T\n", stmt)
		v.Visit(stmt)
	}
	return nil
}
func (v *ReplVisitor) VisitDeclaraciones(ctx *parser.DeclaracionesContext) interface{} {
	//fmt.Println("Entrando a VisitDeclaraciones")
	return v.VisitChildren(ctx)
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
	//fmt.Println("Entrando a VisitStmt")
	return v.VisitChildren(ctx)
}

func (v *ReplVisitor) VisitFuncMain(ctx *parser.FuncMainContext) interface{} {
	//fmt.Println("Entrando a VisitFuncMain")
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
	//fmt.Println("[DEBUG] Entrando a un nuevo bloque. Scope actual:", v.Scope.Name)
	//entorno global para todas las variables declaradas en el bloque
	if v.Scope != v.GlobalScope {
		newScope := NewBaseScope("block", v.Scope)
		v.Scope = newScope
		defer func() {
			v.Scope = v.Scope.Parent
		}()
	}
	for _, decl := range ctx.AllDeclaraciones() {
		v.Visit(decl)
	}
	//fmt.Println("[DEBUG] Saliendo del bloque. Scope actual:", v.Scope.Name)
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

	v.Scope.AddVariable(varName, varType, valueObj, false, false, ctx.GetStart())
	if n == 0 {
		fmt.Println(varName, varType, valueObj)
	}
	//fmt.Printf("[DEBUG] Variable '%s' declarada en scope '%s'\n", varName, v.Scope.Name)

	return nil
}

func (v *ReplVisitor) VisitVariableDeclarationImmutable(ctx *parser.VariableDeclarationImmutableContext) interface{} {
	varName := ctx.ID().GetText()

	if ctx.ASSIGN() != nil && ctx.Expresion() != nil {
		variable := v.Scope.GetVariable(varName)
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
			fmt.Printf("SEMANTICO: tipo '%s' no soportado para asignación\n", variable.Type)
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
	val, err := strconv.Atoi(ctx.GetText())
	if err != nil {
		fmt.Printf("Error al convertir entero: %v\n", ctx.GetText())
		return 0
	}
	return val
}

func (v *ReplVisitor) VisitPrintStatement(ctx *parser.PrintStatementContext) interface{} {
	var outputs []string

	// Intenta obtener el número de hijos de tipo expresion
	n := ctx.GetChildCount()
	for i := 0; i < n; i++ {
		expr, ok := ctx.GetChild(i).(parser.IExpresionContext)
		if ok {
			val := v.Visit(expr)
			switch v := val.(type) {
			case int32: // rune en Go es int32
				outputs = append(outputs, string(v))
			case float64:
				// Si es float pero es entero exacto, imprime con .0
				if v == float64(int64(v)) {
					outputs = append(outputs, fmt.Sprintf("%.1f", v))
				} else {
					outputs = append(outputs, fmt.Sprint(v))
				}
			case *value.SliceValue:
				var elems []string
				for _, e := range val.(*value.SliceValue).Elements {
					elems = append(elems, fmt.Sprint(e.Value()))
				}
				outputs = append(outputs, "["+strings.Join(elems, ", ")+"]")

			default:
				outputs = append(outputs, fmt.Sprint(val))
			}
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
	variable := v.Scope.GetVariable(varName)
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
	variable := v.Scope.GetVariable(varName)
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
	variable := v.Scope.GetVariable(varName)
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
	// Recorre todos los hijos usando tu visitor personalizado
	for i := 0; i < node.GetChildCount(); i++ {
		child := node.GetChild(i)
		if childNode, ok := child.(antlr.ParseTree); ok {
			childNode.Accept(v)
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
	variable := v.Scope.GetVariable(varName)
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
	return ctx.GetText()
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
	return v.Visit(ctx.Expresion())
}

func (v *ReplVisitor) VisitIMCPLICIT(ctx *parser.IMCPLICITContext) interface{} {
	varName := ctx.ID().GetText()
	variable := v.Scope.GetVariable(varName)
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

	leftBool := fmt.Sprint(left) == "true"
	rightBool := fmt.Sprint(right) == "true"

	// Solo acepta booleanos
	if (fmt.Sprint(left) != "true" && fmt.Sprint(left) != "false") ||
		(fmt.Sprint(right) != "true" && fmt.Sprint(right) != "false") {
		//fmt.Printf("Error: operador lógico '%s' solo acepta booleanos\n", op)
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
	return v.VisitChildren(ctx)
}

// condicionales
func (v *ReplVisitor) VisitIfDcl(ctx *parser.IfDclContext) interface{} {
	//aqui manejaremos el entorno del if, tene en cuenta que es un scope nuevo
	newScope := NewBaseScope("IF", v.Scope)
	v.Scope = newScope
	v.IfScope = append(v.IfScope, newScope)
	defer func() { v.Scope = v.Scope.Parent }()
	//fmt.Println("[DEBUG] Entrando a VisitIfDcl:", ctx.GetText())
	condVal := v.Visit(ctx.Expresion())
	//fmt.Printf("Valor de la condición del if: %v (tipo %T)\n", condVal, condVal)
	condBool := fmt.Sprint(condVal) == "true"

	if condBool {
		//fmt.Println("Condición verdadera, ejecutando declaraciones del if")

		for _, decl := range ctx.AllDeclaraciones() {
			v.Visit(decl)
		}
		return nil
	}

	// else if recursivo asi que no eliminar
	for _, elseIf := range ctx.AllElseIfDcl() {
		elseIfCond := v.Visit(elseIf.Expresion())
		if fmt.Sprint(elseIfCond) == "true" {
			for _, decl := range elseIf.AllDeclaraciones() {
				v.Visit(decl)
			}
			return nil
		}
	}

	// else
	if ctx.ElseCondicional() != nil {
		for _, decl := range ctx.ElseCondicional().AllDeclaraciones() {
			v.Visit(decl)
		}
	}

	return nil
}

// El For tiene error no repite el ciclo :') pero si lo lee el programa
func (v *ReplVisitor) VisitForClasico(ctx *parser.ForClasicoContext) interface{} {
	newScope := NewBaseScope("FOR_CLASICO", v.Scope)
	v.Scope = newScope
	defer func() {
		v.Scope = v.Scope.Parent
	}()

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

	v.Scope.AddVariable(varName, varType, varVal, false, false, ctx.GetStart())

	for {
		condVal := v.Visit(ctx.Expresion())
		if fmt.Sprint(condVal) != "true" {
			break
		}

		iterScope := NewBaseScope("FOR_ITER", v.Scope)
		v.Scope = iterScope

		for _, decl := range ctx.Block().AllDeclaraciones() {
			res := v.Visit(decl)

			// 🔴 Controlar flujo con break o continue
			if str, ok := res.(string); ok {
				if str == "continue" {
					v.Scope = v.Scope.Parent // salir del iterScope
					break                    // salto a incremento y siguiente iteración
				} else if str == "break" {
					v.Scope = v.Scope.Parent
					return nil // salimos del for
				}
			}
		}

		v.Scope = v.Scope.Parent // salir del iterScope

		if ctx.Stmt() != nil {
			v.Visit(ctx.Stmt()) // Ejecutar incremento (como i++)
		}
	}

	return nil
}

func (v *ReplVisitor) VisitForCondicionUnica(ctx *parser.ForCondicionUnicaContext) interface{} {
	newScope := NewBaseScope("FOR_SIMPLE", v.Scope)
	v.Scope = newScope
	defer func() {
		v.Scope = v.Scope.Parent
	}()

	v.inForLoop = true
	defer func() {
		v.inForLoop = false
	}()

	for {
		condVal := v.Visit(ctx.Expresion())
		if fmt.Sprint(condVal) != "true" {
			break
		}

		iterScope := NewBaseScope("FOR_ITER", v.Scope)
		v.Scope = iterScope

		for _, decl := range ctx.Block().AllDeclaraciones() {
			res := v.Visit(decl)

			if str, ok := res.(string); ok {
				if str == "continue" {
					v.Scope = v.Scope.Parent
					break
				} else if str == "break" {
					v.Scope = v.Scope.Parent
					return nil
				}
			}
		}

		v.Scope = v.Scope.Parent
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
	v.Scope.AddVariable(varName, sliceType, sliceVal, false, false, ctx.GetStart())
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
				elements = append(elements, value.NewStringValue(fmt.Sprint(val)))
			case "bool":
				elements = append(elements, value.NewBoolValue(fmt.Sprint(val) == "true"))
			case "rune":
				r := []rune(fmt.Sprint(val))
				if len(r) > 0 {
					elements = append(elements, value.NewCharValue(r[0]))
				}
			}
		}
	}
	sliceVal := value.NewSliceValue(tipo, elements)
	v.Scope.AddVariable(varName, sliceType, sliceVal, false, false, ctx.GetStart())
	return nil
}

func (v *ReplVisitor) VisitSliceAssignment(ctx *parser.SliceAssignmentContext) interface{} {
	left := ctx.ID(0).GetText()
	right := ctx.ID(1).GetText()

	sliceRight := v.Scope.GetVariable(right)
	if sliceRight == nil || !strings.HasPrefix(sliceRight.Type, "slice_") {
		fmt.Printf("SEMANTICO: '%s' no es un slice válido\n", right)
		return nil
	}

	sliceType := sliceRight.Type
	copyVal := sliceRight.Value.Copy() // ← asegúrate de que el método Copy exista
	v.Scope.AddVariable(left, sliceType, copyVal, false, false, ctx.GetStart())
	return nil
}

// Indexof Slice no funciona imprime nil
func (v *ReplVisitor) VisitLlamadaFuncion(ctx *parser.LlamadaFuncionContext) interface{} {
	funcName := ctx.GetStart().GetText()

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

	default:
		fmt.Printf("SEMANTICO: función '%s' no implementada\n", funcName)
		return nil
	}
}

// casteos
func (v *ReplVisitor) VisitVariableCastDeclaration(ctx *parser.VariableCastDeclarationContext) interface{} {
	varName := ctx.ID().GetText()
	variable := v.Scope.GetVariable(varName)
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
