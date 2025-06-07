package repl

import (
	"compiler/value"

	"github.com/antlr4-go/antlr/v4"
)

type BaseScope struct {
	Name      string
	Parent    *BaseScope
	Variables map[string]*Variable
}

func NewBaseScope(name string, parent *BaseScope) *BaseScope {
	return &BaseScope{
		Name:      name,
		Parent:    parent,
		Variables: make(map[string]*Variable),
	}
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
