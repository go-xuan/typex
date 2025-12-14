package typex

import (
	"strconv"
	"time"
)

func NewTime(v ...time.Time) *Time {
	var x = &Time{notnull: true}
	if len(v) > 0 {
		x.value = v[0]
	} else {
		x.value = time.Now()
	}
	return x
}

func NewDate(v ...time.Time) *Date {
	var x = &Date{notnull: true}
	if len(v) > 0 {
		x.value = time2date(v[0])
	} else {
		x.value = time2date(time.Now())
	}
	return x
}

type Time struct {
	value   time.Time
	notnull bool
}

func (x *Time) UnmarshalJSON(bytes []byte) error {
	if l := len(bytes); l >= 2 && string(bytes) != "null" {
		// 带引号则去掉引号
		if bytes[0] == 34 && bytes[l-1] == 34 {
			bytes = bytes[1 : l-1]
		}
		str := string(bytes)
		if t, err := time.ParseInLocation(`2006-01-02 15:04:05`, str, time.Local); err == nil {
			x.value = t
			x.notnull = true
			return nil
		}
		if unix, err := strconv.ParseInt(str, 10, 64); err == nil {
			if unix > 1e12 {
				x.value = time.UnixMilli(unix)
			} else {
				x.value = time.Unix(unix, 0)
			}
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
		bytes = x.value.AppendFormat(bytes, "2006-01-02 15:04:05")
		bytes = append(bytes, 34)
		return bytes, nil
	}
	return []byte("null"), nil
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

func time2date(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, time.Local)
}

type Date struct {
	value   time.Time
	notnull bool
}

func (x *Date) UnmarshalJSON(bytes []byte) error {
	if l := len(bytes); l >= 2 && string(bytes) != "null" {
		// 带引号则去掉引号
		if bytes[0] == 34 && bytes[l-1] == 34 {
			bytes = bytes[1 : l-1]
		}
		str := string(bytes)
		if t, err := time.ParseInLocation(`2006-01-02`, str, time.Local); err == nil {
			x.value = t
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
		bytes = x.value.AppendFormat(bytes, "2006-01-02")
		bytes = append(bytes, 34)
		return bytes, nil
	}
	return []byte("null"), nil
}

func (x *Date) Value(def ...time.Time) time.Time {
	if x.Valid() {
		return x.value
	} else if len(def) > 0 {
		return time2date(def[0])
	}
	return time.Time{}
}

func (x *Date) Valid() bool {
	return x != nil && x.notnull
}
