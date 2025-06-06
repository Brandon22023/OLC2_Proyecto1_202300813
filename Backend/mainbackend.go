package main

import (
	"backend/interpreter" // Importa tu visitor personalizado
	"backend/parser"      // Cambia esto a la ruta real de tu paquete generado por ANTLR
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

func main() {
	// Ejemplo 1: Declaración de variable con asignación
	input1 := `mut x int := 10;`

	// Ejemplo 2: Función simple main con impresión
	input2 := `
	func main() {
		fmt.Println(x);
	}
	`

	// Ejemplo 3: Declaración slice
	input3 := `numbers := []int {1,2,3}`

	// Ejemplo combinado
	input := input1 + "\n" + input2 + "\n" + input3

	// Crear stream de entrada ANTLR
	is := antlr.NewInputStream(input)

	// Crear lexer y parser
	lexer := parser.NewV4LangGrammarLexer(is)
	tokens := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewV4LangGrammarParser(tokens)

	// Parsing del programa raíz
	tree := p.Program()

	// Crear visitor personalizado
	visitor := interpreter.NewV4LangVisitor()

	// Visitar el árbol parseado y obtener resultado
	result := visitor.Visit(tree)

	fmt.Println("Resultado del visitor:")
	fmt.Println(result)
}
