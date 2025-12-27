package typex

import (
	"strconv"
)

// NewInt64 创建整数值
func NewInt64(v ...int64) *Int64 {
	var value int64
	if len(v) > 0 {
		value = v[0]
	}
	return &Int64{
		value:   value,
		notnull: true,
	}
}

// String2Int64 将字符串解析为整数值
func String2Int64(s string) int64 {
	i, _ := strconv.ParseInt(s, 10, 64)
	return i
}

// Int642String 将整数值格式化为字符串
func Int642String(i int64) string {
	return strconv.FormatInt(i, 10)
}

// Int64 整数值
type Int64 struct {
	value   int64
	notnull bool
}

func (x *Int64) Cover(v any) {
	switch value := v.(type) {
	case int64:
		x.value = value
	case int:
		x.value = int64(value)
	case string:
		x.value = String2Int64(value)
	case Value:
		x.value = value.Int64()
	default:
		x.value = String2Int64(Any2String(value))
	}
}

func (x *Int64) Value(def ...int64) int64 {
	return x.Int64(def...)
}

func (x *Int64) Valid() bool {
	return x != nil && x.notnull
}

func (x *Int64) String(def ...string) string {
	if x.Valid() {
		return Int642String(x.value)
	} else if len(def) > 0 {
		return def[0]
	}
	return ""
}

func (x *Int64) Int(def ...int) int {
	if x.Valid() {
		return int(x.value)
	} else if len(def) > 0 {
		return def[0]
	}
	return 0
}

func (x *Int64) Int64(def ...int64) int64 {
	if x.Valid() {
		return x.value
	} else if len(def) > 0 {
		return def[0]
	}
	return 0
}

func (x *Int64) Float64(def ...float64) float64 {
	if x.Valid() {
		return float64(x.value)
	} else if len(def) > 0 {
		return def[0]
	}
	return 0
}

func (x *Int64) Bool(def ...bool) bool {
	if x.Valid() {
		return x.value != 0
	} else if len(def) > 0 {
		return def[0]
	}
	return false
}

func (x *Int64) UnmarshalJSON(bytes []byte) error {
	if str := string(bytes); str != "" && str != "null" {
		if value, err := strconv.ParseInt(str, 10, 64); err == nil {
			x.value = value
			x.notnull = true
			return nil
		}
	}
	x.notnull = false
	return nil
}

func (x *Int64) MarshalJSON() ([]byte, error) {
	if x.Valid() {
		return []byte(Int642String(x.value)), nil
	}
	return []byte("null"), nil
}
