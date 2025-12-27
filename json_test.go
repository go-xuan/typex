package typex

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

type Demo struct {
	String *String  `json:"string"`
	Time   *Time    `json:"time"`
	Date   *Date    `json:"date"`
	Bool   *Bool    `json:"bool"`
	Int    *Int     `json:"int"`
	Int64  *Int64   `json:"int64"`
	Float  *Float64 `json:"float"`
}

func TestMarshalJson(t *testing.T) {
	demo := &Demo{
		String: NewString("hello"),
		Time:   NewTime(time.Now()),
		Date:   NewDate(time.Now()),
		Bool:   NewBool(true),
		Int:    NewInt(123456),
		Int64:  NewInt64(9999999),
		Float:  NewFloat64(123.456),
	}
	bytes, _ := json.MarshalIndent(demo, "", "  ")
	fmt.Println("序列化：", string(bytes))
}

func TestUnmarshalJson(t *testing.T) {
	demo := &Demo{}
	bytes := []byte(`{"string":null,"time": 11111,"date":"2024-11-21","bool":222,"int":47826,"int64":23364,"float":57575.138063,"value":123.4}`)
	if err := json.Unmarshal(bytes, demo); err != nil {
		panic(err)
	}
	bytes, _ = json.MarshalIndent(demo, "", "  ")
	fmt.Println("反序列化：", string(bytes))
}
