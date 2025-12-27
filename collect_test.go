package typex

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestArgsMarshal(t *testing.T) {
	args := Args{
		"bool":   NewBool(false),
		"float":  NewFloat64(123.4),
		"int":    NewInt(111),
		"int64":  NewInt64(1111111111),
		"string": NewString("hello world"),
		"zero":   NewZero(),
	}
	bytes, err := json.MarshalIndent(args, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println("序列化:")
	fmt.Println(string(bytes))
}

func TestArgsUnmarshal(t *testing.T) {
	args := Args{}
	bytes := []byte(`{"bool":true,"float":1354646.4,"int":333333,"int64":444444,"string":"world hello","zero":null}`)
	if err := json.Unmarshal(bytes, &args); err != nil {
		panic(err)
	}
	ArgsRange(args, func(k string, v Value) {
		fmt.Println(k, "=", v.String())
	})
}

func TestCollect(t *testing.T) {
	var collect Collect[string, Value]

	v1 := NewValue("111")
	v2 := NewValue("222")

	collect = Args{}
	collect.Put("1", v1)
	collect.Put("2", v2)
	ArgsRange(collect.(Args), func(k string, v Value) {
		fmt.Println("Args:", k, "==>", v)
	})

	collect = fmtPrint[Value]{}
	collect.Put("1", v1)
	collect.Put("2", v2)

	collect = fmtPrint1{}
	collect.Put("1", v1)
	collect.Put("2", v2)

	var collect2 Collect[string, any]
	collect2 = fmtPrint2{}
	collect2.Put("1", v1)
	collect2.Put("2", v2)
}

// 打印器，Collect[string, any]
type fmtPrint[V any] struct{}

func (p fmtPrint[V]) Put(k string, v V) {
	fmt.Println("fmtPrint[V any]:", k, "==>", v)
}

func (p fmtPrint[V]) Get(_ string) (v V) {
	return
}

// 打印器，未使用范性，但是本质实现了 Collect[string, Value]
type fmtPrint1 struct{}

func (p fmtPrint1) Put(k string, v Value) {
	fmt.Println("fmtPrint1:", k, "==>", v)
}

func (p fmtPrint1) Get(_ string) (v Value) {
	return
}

// 打印器，未使用范性，但是本质实现了 Collect[string, any]
type fmtPrint2 struct{}

func (p fmtPrint2) Put(k string, v any) {
	fmt.Println("fmtPrint2:", k, "==>", v)
}

func (p fmtPrint2) Get(_ string) (v any) {
	return
}
