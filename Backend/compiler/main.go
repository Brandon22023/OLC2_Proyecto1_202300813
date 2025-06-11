package main

import (
	"fmt"
	"os"

	compiler "compiler/parser"
	repl "compiler/repl"
	"compiler/symbols"

	//"main/cst"
	"compiler/errors"
	// "main/repl"

	"github.com/antlr4-go/antlr/v4"
)

func main() {
	// inicializamos la aplicacion con fiber

	// Leer código fuente desde archivo o consola
	fmt.Println("Escribe el codigo de Vlang: (Ctrl+D) para terminar")
	inputCode, err := readStdin()
	if err != nil {
		fmt.Println("Error leyendo entrada:", err)
		return
	}

	//startTime := time.Now()

	// Canal para generar CST en paralelo (opcional)
	/*
		resultChannel := make(chan string)
		go func() {
			resultChannel <- cst.CstReport(inputCode)
		}()
	*/

	// 1. Análisis Léxico
	// Para verificar errores
	lexicalErrorListener := errors.NewLexicalErrorListener()
	//
	lexer := compiler.NewVlangLexer(antlr.NewInputStream(inputCode))

	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(lexicalErrorListener)

	// 2. Tokens
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// 3. Parser + errores sintácticos
	// New<Nombre de mi gramatica>(Stream)
	parser := compiler.NewVlangParser(stream)
	parser.BuildParseTrees = true

	syntaxErrorListener := errors.NewSyntaxErrorListener(lexicalErrorListener.ErrorTable)
	parser.RemoveErrorListeners()
	parser.SetErrorHandler(errors.NewCustomErrorStrategy())
	parser.AddErrorListener(syntaxErrorListener)
	fmt.Println("Se termino el analisis sintactico")

	// 4. Árbol sintáctico
	// En tu gramatica tienes el axioma, o simbolo inicial
	// Este es el que deberas agregar como parte del parser.
	arbolito := parser.Programa()
	// Imprime el árbol sintáctico para depuración
	PrintVerticalTree(arbolito, parser.RuleNames)
	// imprimimos los errores de sintaxis y léxicos
	visitor := repl.NewReplVisitor()
	visitor.Visit(arbolito)

	//visitor.GlobalScope.PrintVariables()
	fmt.Println("=== Entornos IF ===")
	for _, ifScope := range visitor.IfScope {
		ifScope.PrintVariables()
	}
	//fmt.Println("aqui imprimire las variables de todos los entornos")
	//visitor.Scope.PrintAllScopesUp()
	//visitor.Scope.PrintAllEnvironments()

	visitor.GlobalScope.PrintAllEnvironmentsRecursive(0)
	// Al final del main:
	tabla := symbols.NewSymbolTable()
	visitor.GlobalScope.CollectSymbols(tabla)
	err = tabla.ToHTML("tabla_simbolos.html")
	if err != nil {
		fmt.Println("Error al escribir la tabla de símbolos:", err)
	} else {
		fmt.Println("Tabla de símbolos generada en tabla_simbolos.html")
	}
}

/*
	func readStdin() (string, error) {
		input, err := os.ReadFile("/dev/stdin")
		return string(input), err
	}
*/
func readStdin() (string, error) {
	input, err := os.ReadFile("/home/brandon/Escritorio/OLC2_Proyecto1_202300813/Backend/compiler/arhivoP.vch")
	//input, err := os.ReadFile("/home/pablo/Escritorio/OLC2_Proyecto1_202300813/Backend/compiler/arhivoP.vch")
	//input, err := os.ReadFile("/home/vboxuser/Documents/OLC2_Proyecto1_202300813/Backend/compiler/arhivoP.vch")
	return string(input), err
}

// Funciones para visualizar nuestro arbol
func PrintVerticalTree(node antlr.Tree, ruleNames []string) {
	printVerticalNode(node, ruleNames, "", true)
}

func printVerticalNode(node antlr.Tree, ruleNames []string, prefix string, isLast bool) {
	connector := "+-- "
	if !isLast {
		connector = "|-- "
	}

	var label string
	switch n := node.(type) {
	case antlr.RuleNode:
		label = ruleNames[n.GetRuleContext().GetRuleIndex()]
	case antlr.TerminalNode:
		label = fmt.Sprintf("\"%s\"", n.GetText())
	default:
		label = fmt.Sprintf("%T", n)
	}

	fmt.Printf("%s%s%s\n", prefix, connector, label)

	// Actualizar el prefijo para los hijos
	childCount := node.GetChildCount()
	for i := 0; i < childCount; i++ {
		child := node.GetChild(i)
		newPrefix := prefix
		if isLast {
			newPrefix += "    "
		} else {
			newPrefix += "|   "
		}
		printVerticalNode(child, ruleNames, newPrefix, i == childCount-1)
	}
}
