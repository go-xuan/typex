package typex

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestReflect(t *testing.T) {
	//var v Value
	//rf := reflect.ValueOf(&v).Elem()
	//fmt.Println("Kind = ", rf.Kind())
	//fmt.Println("NumMethod = ", rf.NumMethod())
	////rf.Set(reflect.ValueOf(true))
	//rf.Set(reflect.ValueOf(NewBool(true)))
	//fmt.Println(v.Bool())

	var b = make(map[string]Value)
	var bytes = []byte(`{"bool":true}`)
	if err := json.Unmarshal(bytes, &b); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(b["bool"])
}

func TestValue(t *testing.T) {
	v := NewFloat64(123.4)
	v.Cover(1)
	fmt.Println("string = ", v.String())
	fmt.Println("int = ", v.Int(111))
	fmt.Println("int64 = ", v.Int64())
	fmt.Println("float64 = ", v.Float64())
	fmt.Println("bool = ", v.Bool())
	fmt.Println("string = ", v.String())
}

func TestMarshal(t *testing.T) {
	demo := &values{
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

func TestUnmarshal(t *testing.T) {
	demo := &values{}
	bytes := []byte(`{"name":null,"time": 11111,"date":"2024-11-21","bool":222,"int":47826,"int64":23364,"float":57575.138063,"value":123.4}`)
	if err := json.Unmarshal(bytes, demo); err != nil {
		panic(err)
	}
	bytes, _ = json.MarshalIndent(demo, "", "  ")
	fmt.Println("反序列化：", string(bytes))
}

type values struct {
	String *String  `json:"string"`
	Time   *Time    `json:"time"`
	Date   *Date    `json:"date"`
	Bool   *Bool    `json:"bool"`
	Int    *Int     `json:"int"`
	Int64  *Int64   `json:"int64"`
	Float  *Float64 `json:"float"`
}
