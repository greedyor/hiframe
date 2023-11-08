package convert

import (
	"strconv"
)

const HIFRAME_CONVENT_STRING_NULL string = ""

func ToString(value interface{}) (toValue string) {
	// Is null
	if value == nil {
		return HIFRAME_CONVENT_STRING_NULL
	}

	// Is address-of
	if value != &value {
		return ToString(value.(*interface{}))
	}

	// Find the corresponding value type
	switch value.(type) {
	case float32:
		ft := value.(float32)
		toValue = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case float64:
		ft := value.(float64)
		toValue = strconv.FormatFloat(ft, 'f', -1, 64)
	case int:
		it := value.(int)
		toValue = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		toValue = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		toValue = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		toValue = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		toValue = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		toValue = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		toValue = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		toValue = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		toValue = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		toValue = strconv.FormatUint(it, 10)
	case string:
		toValue = value.(string)
	case []byte:
		toValue = string(value.([]byte))
	case bool:
		key_bool := value.(bool)
		if key_bool {
			toValue = "true"
		} else {
			toValue = "false"
		}
	default:
	}
	return
}
