package repl

import (
	"compiler/symbols"
	"compiler/value"
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

type BaseScope struct {
	Name      string
	Parent    *BaseScope
	Variables map[string]*Variable
	Children  []*BaseScope
}

func NewBaseScope(name string, parent *BaseScope) *BaseScope {
	scope := &BaseScope{
		Name:      name,
		Parent:    parent,
		Variables: make(map[string]*Variable),
		Children: []*BaseScope{},
	}
	if parent != nil {
		parent.Children = append(parent.Children, scope)
	}
	return scope

}

func (s *BaseScope) variableExists(name string) bool {
	_, exists := s.Variables[name]
	return exists
}

func (s *BaseScope) AddVariable(name string, varType string, val value.IVOR, isConst bool, allowNil bool, token antlr.Token) (*Variable, string) {
	if s.variableExists(name) {
		return nil, "La variable '" + name + "' ya existe en este scope"
	}

	v := &Variable{
		Name:     name,
		Type:     varType,
		Value:    val,
		IsConst:  isConst,
		AllowNil: allowNil,
		Token:    token,
	}

	s.Variables[name] = v
	return v, ""
}

func (s *BaseScope) GetVariable(name string) *Variable {
	scope := s
	for scope != nil {
		if v, ok := scope.Variables[name]; ok {
			return v
		}
		scope = scope.Parent
	}
	return nil
}

func (s *BaseScope) PrintVariables() {
    for _, variable := range s.Variables {
        fmt.Printf("Declarada variable: nombre='%s', tipo='%s', ambito='%s', linea=%d, columna=%d\n",
            variable.Name,
            variable.Type,
            s.Name,
            variable.Token.GetLine(),
            variable.Token.GetColumn(),
        )
    }
}

func (s *BaseScope) PrintAllScopes(level int) {
    prefix := ""
    for i := 0; i < level; i++ {
        prefix += "  "
    }
    fmt.Printf("%s=== Scope: %s ===\n", prefix, s.Name)
    s.PrintVariables()
    // Recorre hijos si los tienes (deberías guardar hijos si quieres recorrerlos)
}

func (s *BaseScope) PrintAllScopesUp() {
    scope := s
    for scope != nil {
        fmt.Printf("=== Scope: %s ===\n", scope.Name)
        scope.PrintVariables()
        scope = scope.Parent
    }
}


// Imprime todas las variables de todos los entornos desde el scope actual hasta el global
func (s *BaseScope) PrintAllEnvironments() {
    scope := s
    for scope != nil {
        fmt.Printf("=== Entorno: %s ===\n", scope.Name)
        for _, variable := range scope.Variables {
            fmt.Printf("  Variable: nombre='%s', tipo='%s', ambito='%s', linea=%d, columna=%d\n",
                variable.Name,
                variable.Type,
                scope.Name,
                variable.Token.GetLine(),
                variable.Token.GetColumn(),
            )
        }
        scope = scope.Parent
    }
}

func (s *BaseScope) PrintAllEnvironmentsRecursive(level int) {
    prefix := ""
    for i := 0; i < level; i++ {
        prefix += "  "
    }
    //fmt.Printf("%s=== Entorno: %s ===\n", prefix, s.Name)
    for _, variable := range s.Variables {
        fmt.Printf("%s  Variable: nombre='%s', tipo='%s', ambito='%s', linea=%d, columna=%d\n",
            prefix, variable.Name, variable.Type, s.Name, variable.Token.GetLine(), variable.Token.GetColumn())
    }
    for _, child := range s.Children {
        child.PrintAllEnvironmentsRecursive(level + 1)
    }
}

func (s *BaseScope) CollectSymbols(st *symbols.SymbolTable) {
    for _, variable := range s.Variables {
        st.AddSymbol(symbols.Symbol{
            ID:       variable.Name,
            SymType:  "Variable",
            DataType: variable.Type,
            Scope:    s.Name,
            Line:     variable.Token.GetLine(),
            Column:   variable.Token.GetColumn(),
        })
    }
    for _, child := range s.Children {
        child.CollectSymbols(st)
    }
}