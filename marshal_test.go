package typex

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestMarshal(t *testing.T) {
	type Demo struct {
		String *String  `json:"name"`
		Time   *Time    `json:"create_time"`
		Date   *Date    `json:"create_date"`
		Bool   *Bool    `json:"bool"`
		Int    *Int     `json:"int"`
		Int64  *Int64   `json:"int64"`
		Float  *Float64 `json:"float"`
	}

	bytes := []byte(`{"name":null,"create_time": 11111,"create_date":"2024-11-21","bool":1,"int":47826,"int64":23364,"float":57575.138063,"value":123.4}`)
	demo := &Demo{}
	if err := json.Unmarshal(bytes, demo); err != nil {
		panic(err)
	}
	bytes, _ = json.Marshal(demo)
	fmt.Println("反序列化：", string(bytes))

	NewString().Valid()
	demo.String = NewString("111")
	demo.Bool = NewBool(true)
	demo.Date = NewDate(time.Now())
	demo.Time = NewTime(time.Now())
	demo.Int = NewInt(123456)
	demo.Int64 = NewInt64(9999999)
	demo.Float = NewFloat64(123.456)
	bytes, _ = json.Marshal(demo)
	fmt.Println("重新赋值：", string(bytes))
}
