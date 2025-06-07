package repl

import (
	"compiler/value"

	"github.com/antlr4-go/antlr/v4"
)

type Variable struct {
	Name     string
	Value    value.IVOR
	Type     string
	IsConst  bool
	AllowNil bool
	Token    antlr.Token
	isProp   bool
}
