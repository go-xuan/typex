package typex

import (
	"strconv"
	"time"
)

const (
	TimeLayout = "2006-01-02 15:04:05"
	DateLayout = "2006-01-02"
)

// NewTime 创建时间值
func NewTime(v ...time.Time) *Time {
	var value time.Time
	if len(v) > 0 {
		value = v[0]
	}
	return &Time{
		value:   value,
		layout:  TimeLayout,
		notnull: true,
	}
}

// NewDate 创建日期值
func NewDate(v ...time.Time) *Date {
	var value time.Time
	if len(v) > 0 {
		value = v[0]
	}
	y, m, d := value.Date()
	return &Date{Time{
		value:   time.Date(y, m, d, 0, 0, 0, 0, time.Local),
		layout:  DateLayout,
		notnull: true,
	}}
}

// String2Time 将字符串解析为时间值
func String2Time(s string) time.Time {
	if l := len(s); l == len(DateLayout) {
		return String2Date(s)
	}
	parse, _ := time.Parse(TimeLayout, s)
	return parse
}

// Time2String 将时间值格式化为字符串
func Time2String(t time.Time) string {
	return t.Format(TimeLayout)
}

// String2Date 将字符串解析为日期值
func String2Date(s string) time.Time {
	parse, _ := time.Parse(DateLayout, s)
	return parse
}

func Date2String(t time.Time) string {
	return t.Format(DateLayout)
}

type Time struct {
	value   time.Time // 时间值
	layout  string    // 时间格式
	notnull bool      // 是否非空
}

func (x *Time) Value(def ...time.Time) time.Time {
	if x.Valid() {
		return x.value
	} else if len(def) > 0 {
		return def[0]
	}
	return time.Time{}
}

func (x *Time) Valid() bool {
	return x != nil && x.notnull
}

func (x *Time) Cover(v any) {
	switch value := v.(type) {
	case time.Time:
		x.value = value
	case string:
		x.value = String2Time(value)
	case int64:
		x.value = time.Unix(value, 0)
	default:
		x.value = String2Time(Any2String(value))
	}
}

func (x *Time) String(def ...string) string {
	if x.Valid() {
		return x.value.Format(x.layout)
	} else if len(def) > 0 {
		return def[0]
	}
	return ""
}

func (x *Time) Int(def ...int) int {
	if x.Valid() {
		return int(x.value.Unix())
	} else if len(def) > 0 {
		return def[0]
	}
	return 0
}

func (x *Time) Int64(def ...int64) int64 {
	if x.Valid() {
		return x.value.Unix()
	} else if len(def) > 0 {
		return def[0]
	}
	return 0
}

func (x *Time) Float64(def ...float64) float64 {
	if x.Valid() {
		return float64(x.value.Unix())
	} else if len(def) > 0 {
		return def[0]
	}
	return 0
}

func (x *Time) Bool(def ...bool) bool {
	if x.Valid() {
		return !x.value.IsZero()
	} else if len(def) > 0 {
		return def[0]
	}
	return false
}

func (x *Time) UnmarshalJSON(bytes []byte) error {
	if l := len(bytes); l >= 2 && string(bytes) != "null" {
		// 带引号则去掉引号
		if bytes[0] == 34 && bytes[l-1] == 34 {
			bytes = bytes[1 : l-1]
		}
		str := string(bytes)
		if parse, err := time.Parse(TimeLayout, str); err == nil {
			x.value = parse
			x.layout = TimeLayout
			x.notnull = true
			return nil
		}
		if unix, err := strconv.ParseInt(str, 10, 64); err == nil {
			if unix > 1e12 {
				x.value = time.UnixMilli(unix)
			} else {
				x.value = time.Unix(unix, 0)
			}
			x.layout = TimeLayout
			x.notnull = true
			return nil
		}
	}
	x.notnull = false
	return nil
}

func (x *Time) MarshalJSON() ([]byte, error) {
	if x.Valid() {
		var bytes []byte
		bytes = append(bytes, 34)
		bytes = x.value.AppendFormat(bytes, x.layout)
		bytes = append(bytes, 34)
		return bytes, nil
	}
	return []byte("null"), nil
}

type Date struct {
	Time
}

func (x *Date) UnmarshalJSON(bytes []byte) error {
	if l := len(bytes); l >= 2 && string(bytes) != "null" {
		// 带引号则去掉引号
		if bytes[0] == 34 && bytes[l-1] == 34 {
			bytes = bytes[1 : l-1]
		}
		if value, err := time.Parse(DateLayout, string(bytes)); err == nil {
			x.value = value
			x.layout = DateLayout
			x.notnull = true
			return nil
		}
	}
	x.notnull = false
	return nil
}

func (x *Date) MarshalJSON() ([]byte, error) {
	if x.Valid() {
		var bytes []byte
		bytes = append(bytes, 34)
		bytes = x.value.AppendFormat(bytes, x.layout)
		bytes = append(bytes, 34)
		return bytes, nil
	}
	return []byte("null"), nil
}
