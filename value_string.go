package typex

import (
	"fmt"
	"strconv"
	"time"
)

// NewString 创建字符串值
func NewString(v ...string) *String {
	var value string
	if len(v) > 0 {
		value = v[0]
	}
	return &String{
		value:   value,
		notnull: true,
	}
}

// Any2String 将任意值格式化为字符串
func Any2String(v any) string {
	return fmt.Sprintf("%v", v)
}

// String 字符串值
type String struct {
	value   string
	notnull bool
}

func (x *String) Value(def ...string) string {
	return x.String(def...)
}

func (x *String) Valid() bool {
	return x != nil && x.notnull
}

func (x *String) Cover(v any) {
	switch value := v.(type) {
	case string:
		x.value = value
	case bool:
		x.value = Bool2String(value)
	case float64, float32:
		x.value = fmt.Sprintf("%f", value)
	case int64, int32, int16, int8, int, uint64, uint32, uint16, uint8, uint:
		x.value = fmt.Sprintf("%d", value)
	case []byte:
		x.value = string(value)
	case time.Time:
		x.value = value.Format(TimeLayout)
	case Value:
		x.value = value.String()
	default:
		x.value = Any2String(value)
	}
}

func (x *String) String(def ...string) string {
	if x.Valid() {
		return x.value
	} else if len(def) > 0 {
		return def[0]
	}
	return ""
}

func (x *String) Int(def ...int) int {
	if x.Valid() {
		value, _ := strconv.Atoi(x.value)
		return value
	} else if len(def) > 0 {
		return def[0]
	}
	return 0
}

func (x *String) Int64(def ...int64) int64 {
	if x.Valid() {
		value, _ := strconv.ParseInt(x.value, 10, 64)
		return value
	} else if len(def) > 0 {
		return def[0]
	}
	return 0
}

func (x *String) Float64(def ...float64) float64 {
	if x.Valid() {
		value, _ := strconv.ParseFloat(x.value, 64)
		return value
	} else if len(def) > 0 {
		return def[0]
	}
	return 0
}

func (x *String) Bool(def ...bool) bool {
	if x.Valid() {
		return String2Bool(x.value)
	} else if len(def) > 0 {
		return def[0]
	}
	return false
}

func (x *String) UnmarshalJSON(bytes []byte) error {
	if l := len(bytes); l >= 0 && string(bytes) != "null" {
		// 带引号则去掉引号
		if l > 1 && bytes[0] == 34 && bytes[l-1] == 34 {
			bytes = bytes[1 : l-1]
		}
		x.notnull = true
		x.value = string(bytes)
		return nil
	}
	x.notnull = false
	return nil
}

func (x *String) MarshalJSON() ([]byte, error) {
	if x.Valid() {
		var bytes []byte
		bytes = append(bytes, 34)
		bytes = append(bytes, []byte(x.value)...)
		bytes = append(bytes, 34)
		return bytes, nil
	}
	return []byte("null"), nil
}
