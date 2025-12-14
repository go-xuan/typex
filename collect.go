package typex

import "encoding/json"

type (
	CollectAny    Collect[string, any]    // 任意类型收集器
	CollectString Collect[string, string] // 字符串收集器
	CollectValue  Collect[string, Value]  // 值收集器
)

type Collect[K comparable, V any] interface {
	Put(k K, v V) // 放入值
	Get(k K) V    // 获取值
}

// ArgsRange 遍历参数收集器
func ArgsRange(args Args, handle func(k string, v Value)) {
	for k, v := range args {
		handle(k, v)
	}
}

// Args 参数收集器，隐式实现了 Collect[string, Value]
// 可被json序列化，但是不可被json反序列化，
type Args map[string]Value

func (a Args) Put(k string, v Value) {
	a[k] = v
}

func (a Args) Get(k string) Value {
	if value, ok := a[k]; ok {
		return value
	}
	return NewZero()
}

func (a Args) UnmarshalJSON(data []byte) error {
	var temp map[string]any
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	for k, v := range temp {
		a[k] = NewValue(v)
	}
	return nil
}
