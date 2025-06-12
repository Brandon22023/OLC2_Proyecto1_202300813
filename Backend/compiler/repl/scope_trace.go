package repl

import (
	"compiler/value"

	"github.com/antlr4-go/antlr/v4"
)

/*
Variable_global = 0

	fn fibonacci(n int) int {
		if n <= 1 {
			return n
		}
		return fibonacci(n-1) + fibonacci(n-2)
	}

fn main() {

}
*/
type ScopeTrace struct {
	GlobalScope  *BaseScope
	CurrentScope *BaseScope
	/*
		Lo vamos a ir trasladando con punteros
		en un arbol de scopes hacia arriba.

	*/
}

func (s *ScopeTrace) PushScope(name string) *BaseScope {

	newScope := NewLocalScope(name)
	s.CurrentScope.AddChild(newScope)
	s.CurrentScope = newScope

	return s.CurrentScope
}

func (s *ScopeTrace) PopScope() {
	s.CurrentScope = s.CurrentScope.Parent()
}

func (s *ScopeTrace) Reset() {
	s.CurrentScope = s.GlobalScope
}

func (s *ScopeTrace) AddVariable(name string, varType string, value value.IVOR, isConst bool, allowNil bool, token antlr.Token) (*Variable, string) {
	return s.CurrentScope.AddVariable(name, varType, value, isConst, allowNil, token)
}

func (s *ScopeTrace) GetVariable(name string) *Variable {
	return s.CurrentScope.GetVariable(name)
}

func (s *ScopeTrace) AddFunction(name string, function value.IVOR) (bool, string) {
	return s.CurrentScope.AddFunction(name, function)
}

func (s *ScopeTrace) GetFunction(name string) (value.IVOR, string) {
	return s.CurrentScope.GetFunction(name)
}

func (s *ScopeTrace) IsMutatingEnvironment() bool {
	return s.CurrentScope.IsMutatingScope()
}

func NewScopeTrace() *ScopeTrace {
	globalScope := NewGlobalScope()
	return &ScopeTrace{
		GlobalScope:  globalScope,
		CurrentScope: globalScope,
	}
}

func NewVectorScope() *BaseScope {
	var scope = &BaseScope{
		name:      "vector",
		variables: make(map[string]*Variable),
		children:  make([]*BaseScope, 0),
		functions: make(map[string]value.IVOR),
		parent:    nil,
	}

	// register object built-in functions

	return scope
}

func NewStructScope() *BaseScope {

	newGlobal := NewGlobalScope()

	return &BaseScope{
		name:      "struct",
		variables: make(map[string]*Variable),
		children:  make([]*BaseScope, 0),
		functions: make(map[string]value.IVOR),
		structs:   make(map[string]*Struct),
		parent:    newGlobal,
		isStruct:  true,
	}
}

// * Report

type ReportTable struct {
	GlobalScope ReportScope
}

type ReportScope struct {
	Name        string
	Vars        []ReportSymbol
	Funcs       []ReportSymbol
	Structs     []ReportSymbol
	ChildScopes []ReportScope
}

type ReportSymbol struct {
	Name   string
	Type   string
	Line   int
	Column int
}

func (s *ScopeTrace) Report() ReportTable {
	return ReportTable{
		GlobalScope: s.CurrentScope.Report(),
	}
}

func (s *BaseScope) Report() ReportScope {

	reportScope := ReportScope{
		Name:        s.name,
		Vars:        make([]ReportSymbol, 0),
		Funcs:       make([]ReportSymbol, 0),
		Structs:     make([]ReportSymbol, 0),
		ChildScopes: make([]ReportScope, 0),
	}

	for _, v := range s.variables {

		token := v.Token
		line := 0
		column := 0

		if token != nil {
			line = token.GetLine()
			column = token.GetColumn()
		}

		reportScope.Vars = append(reportScope.Vars, ReportSymbol{
			Name:   v.Name,
			Type:   v.Type,
			Line:   line,
			Column: column,
		})
	}

	// for _, f := range s.functions {
	// 	switch function := f.(type) {
	// 	case *BuiltInFunction:
	// 		reportScope.Funcs = append(reportScope.Funcs, ReportSymbol{
	// 			Name:   function.Name,
	// 			Type:   "Embebida: " + function.Name,
	// 			Line:   0,
	// 			Column: 0,
	// 		})
	// 	case *Function:

	// 		line := 0
	// 		column := 0

	// 		if function.Token != nil {
	// 			line = function.Token.GetLine()
	// 			column = function.Token.GetColumn()
	// 		}

	// 		reportScope.Funcs = append(reportScope.Funcs, ReportSymbol{
	// 			Name:   function.Name,
	// 			Type:   function.ReturnType,
	// 			Line:   line,
	// 			Column: column,
	// 		})
	// 	case *ObjectBuiltInFunction:
	// 		break
	// 	default:
	// 		log.Fatal("Function type not found")
	// 	}
	// }

	for _, v := range s.structs {
		reportScope.Structs = append(reportScope.Structs, ReportSymbol{
			Name:   v.Name,
			Type:   v.Name,
			Line:   v.Token.GetLine(),
			Column: v.Token.GetColumn(),
		})
	}

	for _, v := range s.children {
		reportScope.ChildScopes = append(reportScope.ChildScopes, v.Report())
	}

	return reportScope
}
