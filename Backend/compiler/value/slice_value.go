package value

type SliceValue struct {
	ElementType string
	Elements    []IVOR
}

func NewSliceValue(elementType string, elements []IVOR) *SliceValue {
	return &SliceValue{ElementType: elementType, Elements: elements}
}

// ✅ Requerido para que implemente IVOR
func (s *SliceValue) Type() string {
	return "slice_" + s.ElementType
}

func (s *SliceValue) Value() interface{} {
	return s.Elements
}

func (s *SliceValue) Copy() IVOR {
	copied := make([]IVOR, len(s.Elements))
	copy(copied, s.Elements)
	return &SliceValue{ElementType: s.ElementType, Elements: copied}
}
