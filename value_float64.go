package typex

import (
	"strconv"
)

// NewFloat64 创建浮点数值
func NewFloat64(f ...float64) *Float64 {
	var value float64
	if len(f) > 0 {
		value = f[0]
	}
	return &Float64{
		value:   value,
		notnull: true,
	}
}

// String2Float64 将字符串解析为浮点数值
func String2Float64(s string) float64 {
	f, _ := strconv.ParseFloat(s, 64)
	return f
}

// Float642String 将浮点数值格式化为字符串
func Float642String(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

// Float64 浮点数值
type Float64 struct {
	value   float64
	notnull bool
}

func (x *Float64) Cover(v any) {
	switch value := v.(type) {
	case float64:
		x.value = value
	case float32:
		x.value = float64(value)
	case string:
		x.value = String2Float64(value)
	case Value:
		x.value = value.Float64()
	default:
		x.value = String2Float64(Any2String(value))
	}
}

func (x *Float64) Value(def ...float64) float64 {
	return x.Float64(def...)
}

func (x *Float64) Valid() bool {
	return x != nil && x.notnull
}

func (x *Float64) String(def ...string) string {
	if x.Valid() {
		return Float642String(x.value)
	} else if len(def) > 0 {
		return def[0]
	}
	return ""
}

func (x *Float64) Int(def ...int) int {
	if x.Valid() {
		return int(x.value)
	} else if len(def) > 0 {
		return def[0]
	}
	return 0
}

func (x *Float64) Int64(def ...int64) int64 {
	if x.Valid() {
		return int64(int(x.value))
	} else if len(def) > 0 {
		return def[0]
	}
	return 0
}

func (x *Float64) Float64(def ...float64) float64 {
	if x.Valid() {
		return x.value
	} else if len(def) > 0 {
		return def[0]
	}
	return 0
}

func (x *Float64) Bool(def ...bool) bool {
	if x.Valid() {
		return x.value != 0
	} else if len(def) > 0 {
		return def[0]
	}
	return false
}

func (x *Float64) UnmarshalJSON(bytes []byte) error {
	if str := string(bytes); str != "" && str != "null" {
		if value, err := strconv.ParseFloat(str, 64); err == nil {
			x.value = value
			x.notnull = true
			return nil
		}
	}
	x.notnull = false
	return nil
}

func (x *Float64) MarshalJSON() ([]byte, error) {
	if x.Valid() {
		return []byte(Float642String(x.value)), nil
	}
	return []byte("null"), nil
}
