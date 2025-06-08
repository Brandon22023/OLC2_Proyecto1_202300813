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
    IfScope     []*BaseScope
}

var _ parser.VlangVisitor = &ReplVisitor{} // <-- Esto asegura la interfaz
// Constructor del visitor
func NewReplVisitor() *ReplVisitor {
    global := NewBaseScope("GLOBAL", nil)
	return &ReplVisitor{
		Scope: global,
        GlobalScope: global,
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

func (v *ReplVisitor) VisitBlock(ctx *parser.BlockContext) interface{} {
    //fmt.Println("[DEBUG] Entrando a un nuevo bloque. Scope actual:", v.Scope.Name)
    //entorno global para todas las variables declaradas en el bloque
    if v.Scope != v.GlobalScope {
        newScope := NewBaseScope("block", v.Scope)
        v.Scope = newScope
        defer func()  {
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
			//outputs = append(outputs, fmt.Sprint(val))
			// Detecta si es rune (carácter)
			switch v := val.(type) {
			case int32: // rune en Go es int32
				outputs = append(outputs, string(v))
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
        return variable.Value.Value()
    }
    fmt.Printf("Variable '%s' no encontrada\n", varName)
    return "<undef>"
}

func (v *ReplVisitor) VisitIncredecr(ctx *parser.IncredecrContext) interface{} {
	//fmt.Println("🔼 Incremento/Decremento:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *ReplVisitor) VisitIncremento(ctx *parser.IncrementoContext) interface{} {
	varName := ""
	if idNode, ok := ctx.GetChild(0).(antlr.TerminalNode); ok {
		varName = idNode.GetText()
	} else {
		fmt.Println("Error: no se pudo obtener el ID del incremento/decremento")
		return nil
	}
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

	variable.Value = value.NewIntValue(val + 1)
	return nil
}

func (v *ReplVisitor) VisitDecremento(ctx *parser.DecrementoContext) interface{} {
	varName := ""
	if idNode, ok := ctx.GetChild(0).(antlr.TerminalNode); ok {
		varName = idNode.GetText()
	} else {
		fmt.Println("Error: no se pudo obtener el ID del incremento/decremento")
		return nil
	}
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

	variable.Value = value.NewIntValue(val - 1)
	return nil
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
func (v *ReplVisitor) VisitForCondicionUnica(ctx *parser.ForCondicionUnicaContext) interface{} {
	cond := ctx.Expresion()
	block := ctx.Block()

	for {
		condVal := v.Visit(cond)
		if fmt.Sprint(condVal) != "true" {
			break
		}
		v.Visit(block)
	}
	return nil
}

func (v *ReplVisitor) VisitForClasico(ctx *parser.ForClasicoContext) interface{} {
	fmt.Println(" Entrando a VisitForClasico")

	initStmt := ctx.Stmt(0)
	cond := ctx.Expresion()
	incrStmt := ctx.Stmt(1)
	block := ctx.Block()

	//  Inicialización
	fmt.Println(" Ejecutando inicialización:")
	if initStmt != nil {
		fmt.Println(" initStmt.GetText():", initStmt.GetText())
		v.Visit(initStmt)
	}

	// Verifica si la variable 'z' se declaró correctamente
	zVar := v.Scope.GetVariable("z")
	if zVar != nil {
		fmt.Printf(" z después de inicializar: %v (tipo %s)\n", zVar.Value.Value(), zVar.Type)
	} else {
		fmt.Println(" La variable z no fue declarada en el scope")
	}

	//  Ejecutar ciclo
	iter := 0
	for {
		iter++
		fmt.Printf(" Iteración #%d\n", iter)

		condVal := v.Visit(cond)
		fmt.Printf(" Evaluando condición (%s): %v\n", cond.GetText(), condVal)

		if fmt.Sprint(condVal) != "true" {
			fmt.Println(" Condición falsa, saliendo del ciclo")
			break
		}

		fmt.Println("Condición verdadera, ejecutando bloque:")
		v.Visit(block)

		if incrStmt != nil {
			fmt.Println(" Ejecutando incremento:", incrStmt.GetText())
			v.Visit(incrStmt)

			// Mostrar el nuevo valor de z
			zVar = v.Scope.GetVariable("z")
			if zVar != nil {
				fmt.Printf(" z después de incremento: %v\n", zVar.Value.Value())
			}
		}
	}

	fmt.Println(" Fin del for clásico")
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

func (v *ReplVisitor) VisitSliceCreacion(ctx *parser.SliceCreacionvContext) interface{} {
	tipo := ctx.TipoSlice().TIPO().GetText()
	var elementos []interface{}

	if ctx.ListaExpresiones() != nil {
		for _, expr := range ctx.ListaExpresiones().AllExpresion() {
			val := v.Visit(expr)
			elementos = append(elementos, val)
		}
	}

	switch tipo {
	case "int":
		var list []int
		for _, e := range elementos {
			val, _ := strconv.Atoi(fmt.Sprint(e))
			list = append(list, val)
		}
		return list

	case "float64":
		var list []float64
		for _, e := range elementos {
			val, _ := strconv.ParseFloat(fmt.Sprint(e), 64)
			list = append(list, val)
		}
		return list

	case "string":
		var list []string
		for _, e := range elementos {
			list = append(list, fmt.Sprint(e))
		}
		return list

	case "bool":
		var list []bool
		for _, e := range elementos {
			b := fmt.Sprint(e) == "true"
			list = append(list, b)
		}
		return list

	case "rune":
		var list []rune
		for _, e := range elementos {
			text := fmt.Sprint(e)
			if len(text) > 0 {
				list = append(list, rune(text[0]))
			}
		}
		return list
	}

	return nil
}

// Codigo de los slices funciona el crear slice el de arriba pero el Join da error por tipos invalidos este comentado
func (v *ReplVisitor) VisitLlamadaFuncion(ctx *parser.LlamadaFuncionContext) interface{} {
	if ctx.GetChildCount() < 3 {
		fmt.Println("❌ Error: llamada de función mal formada")
		return nil
	}

	// Obtener el nombre de la función (JOIN, INDEXOF o ID)
	nombreNode := ctx.GetChild(0)
	nombreFunc := ""
	if token, ok := nombreNode.(antlr.TerminalNode); ok {
		nombreFunc = token.GetText()
	} else {
		fmt.Println("❌ Error: no se pudo obtener el nombre de la función")
		return nil
	}

	// Obtener los argumentos
	args := []interface{}{}
	for i := 2; i < ctx.GetChildCount()-1; i += 2 {
		child := ctx.GetChild(i)
		if expr, ok := child.(parser.IExpresionContext); ok {
			val := v.Visit(expr)
			args = append(args, val)
		}
	}

	// Ejecutar la función correspondiente
	switch nombreFunc {
	case "join":
		if len(args) != 2 {
			fmt.Println("❌ Error: join requiere 2 argumentos")
			return ""
		}

		sliceRaw := args[0]
		sep := fmt.Sprint(args[1])

		strSlice, ok := sliceRaw.([]string)
		if !ok {
			fmt.Printf("❌ Error: el primer argumento no es []string, es %T\n", sliceRaw)
			return ""
		}

		return strings.Join(strSlice, sep)

	case "indexOf":
		// TFalta logica para index OF
		return -1

	default:
		fmt.Printf("ℹ Llamando a función '%s' con argumentos: %v\n", nombreFunc, args)
		return nil
	}
}
