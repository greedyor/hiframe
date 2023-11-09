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
		toValue = strconv.FormatFloat(float64(value.(float32)), 'f', -1, 64)
	case float64:
		toValue = strconv.FormatFloat(value.(float64), 'f', -1, 64)
	case int:
		toValue = strconv.Itoa(value.(int))
	case uint:
		toValue = strconv.Itoa(int(value.(uint)))
	case int8:
		toValue = strconv.Itoa(int(value.(int8)))
	case uint8:
		toValue = strconv.Itoa(int(value.(uint8)))
	case int16:
		toValue = strconv.Itoa(int(value.(int16)))
	case uint16:
		toValue = strconv.Itoa(int(value.(uint16)))
	case int32:
		toValue = strconv.Itoa(int(value.(int32)))
	case uint32:
		toValue = strconv.Itoa(int(value.(uint32)))
	case int64:
		toValue = strconv.FormatInt(value.(int64), 10)
	case uint64:
		toValue = strconv.FormatUint(value.(uint64), 10)
	case string:
		toValue = value.(string)
	case []byte:
		toValue = string(value.([]byte))
	case bool:
		toValue = strconv.FormatBool(value.(bool))
	}
	return
}
