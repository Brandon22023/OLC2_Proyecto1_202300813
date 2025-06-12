package value

type StructValue struct {
	Nombre    string
	Atributos map[string]IVOR
}

func NewStructValue(nombre string, definicion map[string]string) *StructValue {
	atributos := make(map[string]IVOR)
	for attr, tipo := range definicion {
		switch tipo {
		case "int":
			atributos[attr] = NewIntValue(0)
		case "float64":
			atributos[attr] = NewFloatValue(0.0)
		case "string":
			atributos[attr] = NewStringValue("")
		case "bool":
			atributos[attr] = NewBoolValue(false)
		case "rune":
			atributos[attr] = NewCharValue('\x00')
		}
	}
	return &StructValue{Nombre: nombre, Atributos: atributos}
}

// Este es el método adicional que pide el compilador:
func (s *StructValue) Type() string {
	return "struct"
}

func (s *StructValue) TypeName() string {
	return "struct_" + s.Nombre
}

func (s *StructValue) Value() interface{} {
	return s.Atributos
}

func (s *StructValue) Copy() IVOR {
	copiaAtributos := make(map[string]IVOR)
	for k, v := range s.Atributos {
		copiaAtributos[k] = v.Copy()
	}
	return &StructValue{
		Nombre:    s.Nombre,
		Atributos: copiaAtributos,
	}
}
