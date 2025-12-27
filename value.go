package typex

import "time"

// NewValue 创建任意值
func NewValue(v any) Value {
	switch value := v.(type) {
	case int:
		return NewInt(value)
	case int64:
		return NewInt64(value)
	case float64:
		return NewFloat64(value)
	case bool:
		return NewBool(value)
	case string:
		return NewString(value)
	case time.Time:
		return NewTime(value)
	case []byte:
		return NewString(string(value))
	case error:
		return NewString(value.Error())
	case Value:
		return value
	default:
		return NewZero()
	}
}

// Value 任意值
type Value interface {
	Valid() bool                    // 是否有效
	Cover(v any)                    // 覆盖值
	String(def ...string) string    // 转为字符串，若无效则返回默认值
	Int(def ...int) int             // 转为整数，若无效则返回默认值
	Int64(def ...int64) int64       // 转为整数，若无效则返回默认值
	Float64(def ...float64) float64 // 转为浮点数，若无效则返回默认值
	Bool(def ...bool) bool          // 转为布尔值，若无效则返回默认值
}
