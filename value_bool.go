package typex

import (
	"strconv"
	"strings"
)

// NewBool 创建布尔值
func NewBool(v ...bool) *Bool {
	var x = &Bool{notnull: true}
	if len(v) > 0 && v[0] {
		x.value = true
	}
	return x
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
	x.value = NewValue(v).Bool()
	return
}

func (x *Bool) Valid() bool {
	return x != nil && x.notnull
}

func (x *Bool) String(def ...string) string {
	if x.Valid() {
		return strconv.FormatBool(x.value)
	} else if len(def) > 0 {
		return def[0]
	}
	return ""
}

func (x *Bool) Int(def ...int) int {
	if x.Valid() {
		return 1
	} else if len(def) > 0 {
		return def[0]
	}
	return 0
}

func (x *Bool) Int64(def ...int64) int64 {
	if x.Valid() {
		return 1
	} else if len(def) > 0 {
		return def[0]
	}
	return 0
}

func (x *Bool) Float64(def ...float64) float64 {
	if x.Valid() {
		return 1
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

func boolOf(s string) bool {
	switch strings.ToLower(s) {
	case "1", "true", "是", "yes":
		return true
	}
	return false
}

func (x *Bool) UnmarshalJSON(bytes []byte) error {
	if str := string(bytes); str != "" && str != "null" {
		x.notnull = true
		x.value = boolOf(str)
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
