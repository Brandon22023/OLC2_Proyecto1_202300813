package value

import "strconv"

func ToFloat(v interface{}) float64 {
	switch val := v.(type) {
	case int:
		return float64(val)
	case float64:
		return val
	case string:
		f, _ := strconv.ParseFloat(val, 64)
		return f
	default:
		return 0.0
	}
}
