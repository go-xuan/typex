package typex

// NewBool 创建布尔值
func NewBool(v ...bool) *Bool {
	var value bool
	if len(v) > 0 {
		value = v[0]
	}
	return &Bool{
		value:   value,
		notnull: true,
	}
}

// String2Bool 将字符串解析为布尔值
func String2Bool(s string) bool {
	switch s {
	case "t", "T", "true", "TRUE", "True", "yes", "YES", "1", "是":
		return true
	}
	return false
}

// Bool2String 将布尔值格式化为字符串
func Bool2String(b bool) string {
	if b {
		return "true"
	}
	return "false"
}

// Bool 布尔值
type Bool struct {
	value   bool
	notnull bool
}

func (x *Bool) Value(def ...bool) bool {
	return x.Bool(def...)
}

func (x *Bool) Cover(v any) {
	switch value := v.(type) {
	case bool:
		x.value = value
	case string:
		x.value = String2Bool(value)
	case Value:
		x.value = value.Bool()
	default:
		x.value = String2Bool(Any2String(value))
	}
}

func (x *Bool) Valid() bool {
	return x != nil && x.notnull
}

func (x *Bool) String(def ...string) string {
	if x.Valid() {
		return Bool2String(x.value)
	} else if len(def) > 0 {
		return def[0]
	}
	return ""
}

func (x *Bool) Int(def ...int) int {
	if x.Valid() {
		if x.value {
			return 1
		}
	} else if len(def) > 0 {
		return def[0]
	}
	return 0
}

func (x *Bool) Int64(def ...int64) int64 {
	if x.Valid() {
		if x.value {
			return 1
		}
	} else if len(def) > 0 {
		return def[0]
	}
	return 0
}

func (x *Bool) Float64(def ...float64) float64 {
	if x.Valid() {
		if x.value {
			return 1
		}
	} else if len(def) > 0 {
		return def[0]
	}
	return 0
}

func (x *Bool) Bool(def ...bool) bool {
	if x.Valid() {
		return x.value
	} else if len(def) > 0 {
		return def[0]
	}
	return false
}

func (x *Bool) UnmarshalJSON(bytes []byte) error {
	if value := string(bytes); value != "" && value != "null" {
		x.notnull = true
		x.value = String2Bool(value)
		return nil
	}
	x.notnull = false
	return nil
}

func (x *Bool) MarshalJSON() ([]byte, error) {
	if x.Valid() {
		if x.value {
			return []byte("true"), nil
		}
		return []byte("false"), nil
	}
	return []byte("null"), nil
}
