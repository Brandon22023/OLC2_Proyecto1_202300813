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

func (v *Variable) TypeValidation() (bool, string) {
    // Si el valor es nil, permitimos si AllowNil está activo
    if v.Value == nil {
        if v.AllowNil {
            return true, ""
        }
        return false, "La variable '" + v.Name + "' no puede ser nil"
    }

    // Comprobamos que el tipo del valor coincida con el tipo declarado
    if v.Value.Type() != v.Type {
        return false, "Tipo incompatible para la variable '" + v.Name + "': se esperaba '" + v.Type + "', pero se obtuvo '" + v.Value.Type() + "'"
    }

    return true, ""
}