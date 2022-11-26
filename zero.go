package basic

func IsZero(val any) bool {
	switch val.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return val == 0
	case float64, float32:
		return val == 0.0
	case complex64, complex128:
		return val == 0i
	case string:
		return val == ""
	case bool:
		return val.(bool) == false
	case []byte:
		return len(val.([]byte)) == 0
	default:
		return val == nil
	}
}
