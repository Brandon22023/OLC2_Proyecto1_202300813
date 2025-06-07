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
	parser.BaseVlangVisitor // Embebemos el visitor generado por ANTLR
    Scope *BaseScope  //este servira para obtener el valor de las variables declaradas xd
}
var _ parser.VlangVisitor = &ReplVisitor{} // <-- Esto asegura la interfaz
// Constructor del visitor
func NewReplVisitor() *ReplVisitor {
	return &ReplVisitor{
        Scope: NewBaseScope("global", nil),
    }
}
/*func (v *ReplVisitor) Visit(tree antlr.ParseTree) interface{} {

	switch val := tree.(type) {
	case *antlr.ErrorNodeImpl:
		log.Fatal(val.GetText())
		return nil
	default:
		return tree.Accept(v)
	}

}*/
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

func (v *ReplVisitor) VisitStmt(ctx *parser.StmtContext) interface{} {
    //fmt.Println("Entrando a VisitStmt")
    return v.VisitChildren(ctx)
}



func (v *ReplVisitor) VisitFuncMain(ctx *parser.FuncMainContext) interface{} {
    //fmt.Println("Entrando a VisitFuncMain")
    return v.Visit(ctx.Block())
}

func (v *ReplVisitor) VisitBlock(ctx *parser.BlockContext) interface{} {
    //fmt.Println("Entrando a VisitBlock")
    for _, decl := range ctx.AllDeclaraciones() {
        //sfmt.Printf("Tipo de declaración: %T\n", decl)
        v.Visit(decl)
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

    v.Scope.AddVariable(varName, varType, valueObj, false, false, ctx.GetStart())
    if n == 0 {
        fmt.Println(varName, varType, valueObj)
    }
    // Usa valueObj aquí (por ejemplo, agrega al scope)
    // scope.AddVariable(varName, varType, valueObj, false, false, ctx.GetStart())

    return nil
}
func (v *ReplVisitor) VisitValorCadena(ctx *parser.ValorCadenaContext) interface{} {
    //fmt.Println("Entrando a VisitValorCadena:", ctx.GetText())
    text := ctx.GetText()
    if len(text) >= 2 {
        return text[1:len(text)-1]
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
    left := v.Visit(ctx.Expresion(0))
    right := v.Visit(ctx.Expresion(1))
    op := ctx.GetOp().GetText()

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

func (v *ReplVisitor) VisitOr(ctx *parser.OrContext) interface{} {
	//fmt.Println("🚪 Operador OR lógico")
	return v.VisitChildren(ctx)
}

func (v *ReplVisitor) VisitId(ctx *parser.IdContext) interface{} {
    varName := ctx.GetText()
    variable := v.Scope.GetVariable(varName)
    if variable != nil && variable.Value != nil {
        return variable.Value.Value()
    }
    return "<undef>"
}

func (v *ReplVisitor) VisitIncredecr(ctx *parser.IncredecrContext) interface{} {
	//fmt.Println("🔼 Incremento/Decremento:", ctx.GetText())
	return v.VisitChildren(ctx)
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
        return rune([]rune(text[1:len(text)-1])[0])
    }
    return rune(0)
}
//NO ELIMINES ESTA FUNCION SI NO TE CARGAS TODO LITERALMENTE xD 
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
        fmt.Printf("Error: operador lógico '%s' solo acepta booleanos\n", op)
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
//condicionales
func (v *ReplVisitor) VisitIf_context(ctx *parser.If_contextContext) interface{} {
    //fmt.Println("========== [DEBUG IF_CONTEXT] ==========")
    if ctx.IfDcl() != nil {
        return v.Visit(ctx.IfDcl())
    }
    //fmt.Println("[DEBUG IF_CONTEXT] No se encontró ifDcl")
    return nil
}

func (v *ReplVisitor) VisitIfDcl(ctx *parser.IfDclContext) interface{} {
    //fmt.Println("========== [DEBUG IFDCL] INICIO ==========")
    var cond antlr.ParseTree
    var ifBlock []antlr.ParseTree
    var elseBlock []antlr.ParseTree
    var inIfBlock, inElseBlock bool

    ruleNames := ctx.GetParser().GetRuleNames()
    //fmt.Printf("[DEBUG IFDCL] Total hijos del ifDcl: %d\n", ctx.GetChildCount())

    for i := 0; i < ctx.GetChildCount(); i++ {
        child := ctx.GetChild(i)
        //fmt.Printf("[DEBUG IFDCL] Hijo #%d: %T\n", i, child)

        // Busca el primer hijo que sea una regla "expresion"
        if cond == nil {
            if rule, ok := child.(antlr.RuleNode); ok {
                ruleName := ruleNames[rule.GetRuleContext().GetRuleIndex()]
                //fmt.Printf("[DEBUG IFDCL]   Es RuleNode, ruleName: %s\n", ruleName)
                if ruleName == "expresion" {
                    cond = rule.(antlr.ParseTree)
                    //fmt.Printf("[DEBUG IFDCL]   -> Se detectó la condición (expresion) en hijo #%d\n", i)
                }
            } else {
                //fmt.Printf("[DEBUG IFDCL]   No es RuleNode\n")
            }
        }

        // Detectar inicio y fin de bloques IF y ELSE
        var text string
        if t, ok := child.(antlr.TerminalNode); ok {
            text = t.GetText()
            //fmt.Printf("[DEBUG IFDCL]   Es TerminalNode, text: '%s'\n", text)
        } else {
            text = fmt.Sprintf("%T", child)
            //fmt.Printf("[DEBUG IFDCL]   No es TerminalNode, tipo: %s\n", text)
        }

        if text == "{" && !inIfBlock && !inElseBlock {
            inIfBlock = true
            //fmt.Printf("[DEBUG IFDCL]   -> INICIO bloque IF en hijo #%d\n", i)
            continue
        }
        if text == "}" && inIfBlock {
            inIfBlock = false
            //fmt.Printf("[DEBUG IFDCL]   -> FIN bloque IF en hijo #%d\n", i)
            continue
        }
        if text == "else" {
            inElseBlock = true
            //fmt.Printf("[DEBUG IFDCL]   -> INICIO bloque ELSE en hijo #%d\n", i)
            continue
        }
        if text == "{" && inElseBlock {
            //fmt.Printf("[DEBUG IFDCL]   -> INICIO bloque ELSE (llave) en hijo #%d\n", i)
            continue
        }
        if text == "}" && inElseBlock {
            inElseBlock = false
            //fmt.Printf("[DEBUG IFDCL]   -> FIN bloque ELSE en hijo #%d\n", i)
            continue
        }

        // Guardar declaraciones en el bloque correspondiente
        if inIfBlock {
            if pt, ok := child.(parser.IDeclaracionesContext); ok {
                ifBlock = append(ifBlock, pt)
                //fmt.Printf("[DEBUG IFDCL]   -> Agregada declaración a IF en hijo #%d\n", i)
            }
        } else if inElseBlock {
            if pt, ok := child.(parser.IDeclaracionesContext); ok {
                elseBlock = append(elseBlock, pt)
                //fmt.Printf("[DEBUG IFDCL]   -> Agregada declaración a ELSE en hijo #%d\n", i)
            }
        }
    }

    if cond == nil {
        //fmt.Println("[DEBUG IFDCL] No se encontró condición en el if, se omite el bloque.")
        //fmt.Println("========== [DEBUG IFDCL] FIN ==========")
        return nil
    }

    // Evaluar condición
    //fmt.Println("[DEBUG IFDCL] Evaluando condición...")
    condVal := v.Visit(cond)
    //fmt.Printf("[DEBUG IFDCL] Valor de la condición: %v (tipo: %T)\n", condVal, condVal)
    condBool := fmt.Sprint(condVal) == "true"
    //fmt.Printf("[DEBUG IFDCL] ¿Condición evaluada como verdadera?: %v\n", condBool)

    // Ejecutar bloque correspondiente
    if condBool {
        //fmt.Println("[DEBUG IFDCL] Ejecutando bloque IF")
        for idx, decl := range ifBlock {
            var n = 1
            if n == 0{
                fmt.Printf("[DEBUG IFDCL]   Ejecutando declaración IF #%d\n", idx)
            }
            v.Visit(decl)
        }
    } else {
        //fmt.Println("[DEBUG IFDCL] Ejecutando bloque ELSE (si existe)")
        for idx, decl := range elseBlock {
            var n = 1
            if n == 0{
                fmt.Printf("[DEBUG IFDCL]   Ejecutando declaración IF #%d\n", idx)
            }
            v.Visit(decl)
        }
    }
    //fmt.Println("========== [DEBUG IFDCL] FIN ==========")
    return nil
}