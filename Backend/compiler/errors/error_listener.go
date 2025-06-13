package errors

import (
	"compiler/repl"
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

type SyntaxErrorListener struct {
	*antlr.DefaultErrorListener
	ErrorTable *repl.ErrorTable
}

func NewSyntaxErrorListener(errorTable *repl.ErrorTable) *SyntaxErrorListener {
	return &SyntaxErrorListener{
		ErrorTable: errorTable,
	}
}

func (l *SyntaxErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{},
    line, column int, msg string, e antlr.RecognitionException) {
    tokenText := ""
    if t, ok := offendingSymbol.(antlr.Token); ok {
        tokenText = t.GetText()
    }
    l.ErrorTable.AddError(
        line,
        column,
        fmt.Sprintf("El token \"%s\" no era el esperado.", tokenText),
        repl.SyntaxError,
    )
}

type LexicalErrorListener struct {
	*antlr.DefaultErrorListener
	ErrorTable *repl.ErrorTable
}

func NewLexicalErrorListener() *LexicalErrorListener {
	return &LexicalErrorListener{
		ErrorTable: repl.NewErrorTable(),
	}
}

func (l *LexicalErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{},
    line, column int, msg string, e antlr.RecognitionException) {

    simbolo := ""
    if t, ok := offendingSymbol.(antlr.Token); ok && t != nil {
        simbolo = t.GetText()
    }

    // Si simbolo sigue vacío, intenta extraerlo del input
    if simbolo == "" {
        if lexer, ok := recognizer.(antlr.Lexer); ok {
            inputStream := lexer.GetInputStream()
            start := inputStream.Index()
            if start < inputStream.Size() {
                simbolo = string(inputStream.GetText(start, start))
            }
        }
    }

    l.ErrorTable.AddError(
        line,
        column,
        fmt.Sprintf("El símbolo \"%s\" no es aceptado en el lenguaje.", simbolo),
        repl.LexicalError,
    )
}
