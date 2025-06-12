package repl

import (
	"compiler/symbols"
	"compiler/value"
	//"fmt"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

type BaseScope struct {
	name       string
	parent     *BaseScope   /// variable aqui ->
	children   []*BaseScope // hijos deben conocer esa variable
	variables  map[string]*Variable
	functions  map[string]value.IVOR
	structs    map[string]*Struct
	isStruct   bool
	IsMutating bool
}
type Struct struct {
    Name  string
    Token antlr.Token
    // Puedes agregar más campos según tus necesidades
}
// Anadir hijos en Clase

func (s *BaseScope) Name() string {
	return s.name
}

func (s *BaseScope) Parent() *BaseScope {
	return s.parent
}

func (s *BaseScope) Children() []*BaseScope {
	return s.children
}

func (s *BaseScope) ValidType(_type string) bool {

	_, isStructType := s.structs[_type]

	return value.IsPrimitiveType(_type) || isStructType
}

func (s *BaseScope) AddChild(child *BaseScope) {
	s.children = append(s.children, child)
	child.parent = s
}

/*
Clase Aqui xd
*/
func (s *BaseScope) variableExists(variable *Variable) bool {
    _, exists := s.variables[variable.Name]
    return exists
}

func (s *BaseScope) AddVariable(name string, varType string, value value.IVOR, isConst bool, allowNil bool, token antlr.Token) (*Variable, string) {

	variable := &Variable{
		Name:     name,
		Value:    value,
		Type:     varType,
		IsConst:  isConst,
		AllowNil: allowNil,
		Token:    token,
	}

	if s.variableExists(variable) {
		return nil, "La variable " + name + " ya existe"
	}

	typesOk, msg := variable.TypeValidation()

	// even if the variable is not valid, we add it to the scope, (internally it will be nil)
	s.variables[name] = variable

	if !typesOk {
		// report error
		return nil, msg
	}

	return variable, ""
}

func (s *BaseScope) GetVariable(name string) *Variable {
	// verify if is refering to and object/struct function

	initialScope := s

	for {
		if variable, ok := initialScope.variables[name]; ok {

			// verify if is refering to a pointer
			// if variable.Type == value.IVOR_POINTER {
			// 	return variable.Value.(*PointerValue).AssocVariable // pointer of a pointer ?
			// }

			return variable
		}

		if initialScope.parent == nil {
			break
		}

		// -> Hacia el padre
		initialScope = initialScope.parent
	}

	return nil
}

// obj1.obj2.prop1

func (s *BaseScope) AddFunction(name string, function value.IVOR) (bool, string) {
	// check if function already exists

	if _, ok := s.functions[name]; ok {
		return false, "La funcion " + name + " ya existe"
	}

	s.functions[name] = function

	return true, ""
}

func (s *BaseScope) GetFunction(name string) (value.IVOR, string) {

	// verify if is refering to and object/struct function

	initialScope := s

	for {
		if function, ok := initialScope.functions[name]; ok {
			return function, ""
		}

		if initialScope.parent == nil {
			break
		}

		initialScope = initialScope.parent
	}

	return nil, "La funcion " + name + " no existe"
}

func (s *BaseScope) AddStruct(name string, structValue *Struct) (bool, string) {

	if _, ok := s.structs[name]; ok {
		return false, "La estructura " + name + " ya existe"
	}

	s.structs[name] = structValue
	return true, ""
}

func (s *BaseScope) GetStruct(name string) (*Struct, string) {

	initialScope := s

	for {
		if structValue, ok := initialScope.structs[name]; ok {
			return structValue, ""
		}

		if initialScope.parent == nil {
			break
		}

		initialScope = initialScope.parent
	}

	return nil, "La estructura " + name + " no existe"
}

func (s *BaseScope) Reset() {
	s.variables = make(map[string]*Variable)
	s.children = make([]*BaseScope, 0)
	s.functions = make(map[string]value.IVOR)
}

func (s *BaseScope) IsMutatingScope() bool {
	aux := s

	for {
		if aux.IsMutating {
			return true
		}

		if aux.parent == nil {
			break
		}

		aux = aux.parent
	}

	return false
}

func NewGlobalScope() *BaseScope {

	// register built-in functions

	funcs := make(map[string]value.IVOR)

	// for k, v := range DefaultBuiltInFunctions {
	// 	funcs[k] = v
	// }

	return &BaseScope{
		name:      "global",
		variables: make(map[string]*Variable),
		children:  make([]*BaseScope, 0),
		//structs:   make(map[string]*Struct),
		parent:    nil,
		functions: funcs,
	}
}

func NewLocalScope(name string) *BaseScope {
	return &BaseScope{
		name:      name,
		variables: make(map[string]*Variable),
		functions: make(map[string]value.IVOR),
		children:  make([]*BaseScope, 0),
		parent:    nil,
	}
}

func (s *BaseScope) PrintScopeVariables(indent int) {
	//fmt.Printf("-----------------------")
	//prefix := strings.Repeat("  ", indent)
	//fmt.Printf("%sScope: %s\n", prefix, s.Name())

	if len(s.variables) == 0 {
		//fmt.Printf("%s  (no variables)\n", prefix)
	} else {
		//for name, variable := range s.variables {
			//fmt.Printf("%s  - Variable: %s (Type: %s)\n", prefix, name, variable.Type)
		//}
	}

	for _, child := range s.children {
		child.PrintScopeVariables(indent + 1)
	}
	//fmt.Printf("%s-----------------------\n", prefix)
}

func (s *BaseScope) CollectSymbols(tabla *symbols.SymbolTable) {
    scopeName := s.name
    // Busca función contenedora
    funcName := ""
    parent := s.parent
    for parent != nil {
        if strings.HasPrefix(parent.name, "fn_") {
            funcName = parent.name[3:] // Quita "fn_"
            break
        }
        if parent.name == "main" {
            funcName = "main"
            break
        }
        parent = parent.parent
    }
    // Renombra IF y FOR
    if strings.HasPrefix(scopeName, "IF") && funcName != "" {
        scopeName = "IF_" + funcName
    }
    if strings.HasPrefix(scopeName, "FOR") && funcName != "" {
        scopeName = "FOR_" + funcName
    }
    // Variables
    for _, v := range s.variables {
        line, col := 0, 0
        if v.Token != nil {
            line = v.Token.GetLine()
            col = v.Token.GetColumn()
        }
        tabla.AddSymbol(symbols.Symbol{
            ID:      v.Name,
            SymType: "Variable",
            DataType: v.Type,
            Scope:   scopeName,
            Line:    line,
            Column:  col,
        })
    }
    // Funciones
    for fname := range s.functions {
        tabla.AddSymbol(symbols.Symbol{
            ID:      fname,
            SymType: "Función",
            DataType: "function",
            Scope:   scopeName,
            Line:    0,
            Column:  0,
        })
    }
    for _, child := range s.children {
        child.CollectSymbols(tabla)
    }
}