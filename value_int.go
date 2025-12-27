package typex

import (
	"strconv"
)

// NewInt 创建整数值
func NewInt(v ...int) *Int {
	var value int
	if len(v) > 0 {
		value = v[0]
	}
	return &Int{
		value:   value,
		notnull: true,
	}
}

// String2Int 将字符串解析为整数值
func String2Int(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

// Int2String 将整数值格式化为字符串
func Int2String(i int) string {
	return strconv.Itoa(i)
}

// Int 整数值
type Int struct {
	value   int
	notnull bool
}

func (x *Int) Cover(v any) {
	switch value := v.(type) {
	case int:
		x.value = value
	case int64:
		x.value = int(value)
	case string:
		x.value = String2Int(value)
	case Value:
		x.value = value.Int()
	default:
		x.value = String2Int(Any2String(value))
	}
}

func (x *Int) Value(def ...int) int {
	return x.Int(def...)
}

func (x *Int) Valid() bool {
	return x != nil && x.notnull
}

func (x *Int) String(def ...string) string {
	if x.Valid() {
		return Int2String(x.value)
	} else if len(def) > 0 {
		return def[0]
	}
	return ""
}

func (x *Int) Int(def ...int) int {
	if x.Valid() {
		return x.value
	} else if len(def) > 0 {
		return def[0]
	}
	return 0
}

func (x *Int) Int64(def ...int64) int64 {
	if x.Valid() {
		return int64(x.value)
	} else if len(def) > 0 {
		return def[0]
	}
	return 0
}

func (x *Int) Float64(def ...float64) float64 {
	if x.Valid() {
		return float64(x.value)
	} else if len(def) > 0 {
		return def[0]
	}
	return 0
}

func (x *Int) Bool(def ...bool) bool {
	if x.Valid() {
		return x.value != 0
	} else if len(def) > 0 {
		return def[0]
	}
	return false
}

func (x *Int) UnmarshalJSON(bytes []byte) error {
	if str := string(bytes); str != "" && str != "null" {
		if value, err := strconv.Atoi(str); err == nil {
			x.value = value
			x.notnull = true
			return nil
		}
	}
	x.notnull = false
	return nil
}

func (x *Int) MarshalJSON() ([]byte, error) {
	if x.Valid() {
		return []byte(Int2String(x.value)), nil
	}
	return []byte("null"), nil
}
