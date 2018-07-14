package iconv

import (
	"fmt"
	"math"
)

func errorInvalid(value interface{}) error {
	return fmt.Errorf("invalid value with type %T: %+#v", value, value)
}

func String(field interface{}) (string, error) {
	v, ok := field.(string)
	if !ok {
		return "", errorInvalid(field)
	}

	return v, nil
}

func Map(field interface{}) (map[string]interface{}, error) {
	v, ok := field.(map[string]interface{})
	if !ok {
		return nil, errorInvalid(field)
	}

	return v, nil
}

func Array(field interface{}) ([]interface{}, error) {
	v, ok := field.([]interface{})
	if !ok {
		return nil, errorInvalid(field)
	}

	return v, nil
}

func MapArray(field interface{}) ([]map[string]interface{}, error) {
	n, ok := field.([]map[string]interface{})
	if ok {
		return n, nil
	}

	raw, ok := field.([]interface{})
	if !ok {
		return nil, errorInvalid(field)
	}

	n = make([]map[string]interface{}, len(raw))
	for i, v := range raw {
		n[i] = v.(map[string]interface{})
	}

	return n, nil
}

func StringArray(field interface{}) ([]string, error) {
	var iraw []interface{}
	sraw, ok := field.([]string)
	if !ok {
		iraw, ok = field.([]interface{})
		if !ok {
			return nil, errorInvalid(field)
		}
	} else {
		return sraw, nil
	}

	s := make([]string, len(iraw))
	for i, v := range iraw {
		s[i] = v.(string)
	}

	return s, nil
}

func Int64(field interface{}) (int64, error) {
	var v int64
	switch fv := field.(type) {
	case int:
		v = int64(fv)
	case int16:
		v = int64(fv)
	case int32:
		v = int64(fv)
	case int64:
		v = fv
	default:
		return 0, errorInvalid(field)
	}

	return v, nil
}

func Int32(field interface{}) (int32, error) {
	var v int32
	switch fv := field.(type) {
	case int:
		if fv > math.MaxInt32 {
			return 0, errorInvalid(field)
		}

		v = int32(fv)
	case int16:
		v = int32(fv)
	case int32:
		v = fv
	case int64:
		if fv > math.MaxInt32 {
			return 0, errorInvalid(field)
		}
	default:
		return 0, errorInvalid(field)
	}

	return v, nil
}

func Int16(field interface{}) (int16, error) {
	var v int16
	switch fv := field.(type) {
	case int:
		if fv > math.MaxInt16 {
			return 0, errorInvalid(field)
		}

		v = int16(fv)
	case int32:
		if fv > math.MaxInt16 {
			return 0, errorInvalid(field)
		}

		v = int16(fv)
	case int64:
		if fv > math.MaxInt16 {
			return 0, errorInvalid(field)
		}

		v = int16(fv)
	case int16:
		v = fv
	default:
		return 0, errorInvalid(field)
	}

	return v, nil
}

func Int(field interface{}) (int, error) {
	var v int
	switch fv := field.(type) {
	case int:
		v = fv
	case int32:
		v = int(fv)
	case int64:
		v = int(fv)
	default:
		return 0, errorInvalid(field)
	}

	return v, nil
}

func Float64(field interface{}) (float64, error) {
	switch v := field.(type) {
	case int64:
		return float64(v), nil
	case int32:
		return float64(v), nil
	case float32:
		return float64(v), nil
	case int:
		return float64(v), nil
	case float64:
		return v, nil
	}

	return 0, errorInvalid(field)
}

func Float32(field interface{}) (float32, error) {
	switch v := field.(type) {
	case int64:
		return float32(v), nil
	case int32:
		return float32(v), nil
	case float64:
		return float32(v), nil
	case int:
		return float32(v), nil
	case float32:
		return v, nil
	}

	return 0, errorInvalid(field)
}

func Bool(field interface{}) (bool, error) {
	v, ok := field.(bool)
	if !ok {
		return false, errorInvalid(field)
	}

	return v, nil
}