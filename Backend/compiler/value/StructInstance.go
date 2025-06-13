package value

type StructInstance struct {
	StructName string
	Attributes map[string]interface{}
}

func NewStructInstance(name string, attrs map[string]interface{}) *StructInstance {
	return &StructInstance{
		StructName: name,
		Attributes: attrs,
	}
}

// IVOR interface implementation

func (s *StructInstance) Value() interface{} {
	return s
}

func (s *StructInstance) Copy() IVOR {
	newAttrs := make(map[string]interface{})
	for k, v := range s.Attributes {
		newAttrs[k] = v
	}
	return &StructInstance{
		StructName: s.StructName,
		Attributes: newAttrs,
	}
}

func (s *StructInstance) Type() string {
	return "struct_" + s.StructName
}
