package typex

import "sync"

// NewStringEnum key为string类型，value为任意类型
func NewStringEnum[T any]() *Enum[string, T] {
	return &Enum[string, T]{
		keys: make([]string, 0),
		data: make(map[string]T),
	}
}

// NewIntEnum key为int类型，value为任意类型
func NewIntEnum[T any]() *Enum[int, T] {
	return &Enum[int, T]{
		keys: make([]int, 0),
		data: make(map[int]T),
	}
}

// NewEnum key为comparable类型，value为任意类型
func NewEnum[K comparable, V any]() *Enum[K, V] {
	return &Enum[K, V]{
		keys: make([]K, 0),
		data: make(map[K]V),
	}
}

// Enum 枚举类
type Enum[K comparable, V any] struct {
	mu   sync.RWMutex // 读写锁
	keys []K          // 保证有序
	data map[K]V      // 存储枚举值
}

func (e *Enum[K, V]) Len() int {
	return len(e.keys)
}

func (e *Enum[K, V]) Clear() {
	e.keys = make([]K, 0)
	e.data = make(map[K]V)
}

func (e *Enum[K, V]) Remove(k K) {
	if len(e.keys) == 0 {
		return
	}
	e.mu.Lock()
	defer e.mu.Unlock()
	delete(e.data, k)
	i := 0 // 使用双指针法删除切片元素
	for _, key := range e.keys {
		if key != k {
			e.keys[i] = key
			i++
		}
	}
	e.keys = e.keys[:i]
}

func (e *Enum[K, V]) Get(k K) V {
	e.mu.Lock()
	defer e.mu.Unlock()
	return e.data[k]
}

func (e *Enum[K, V]) Exist(k K) bool {
	e.mu.Lock()
	defer e.mu.Unlock()
	_, ok := e.data[k]
	return ok
}

func (e *Enum[K, V]) Add(k K, v V) *Enum[K, V] {
	e.mu.Lock()
	defer e.mu.Unlock()
	if _, ok := e.data[k]; !ok {
		e.keys = append(e.keys, k)
	}
	e.data[k] = v
	return e
}

// Keys 返回枚举值的键列表
func (e *Enum[K, V]) Keys() []K {
	e.mu.RLock()
	defer e.mu.RUnlock()
	// 返回切片的副本，避免外部修改内部数据
	ks := make([]K, len(e.keys))
	copy(ks, e.keys)
	return ks
}

// Values 返回枚举值的值列表
func (e *Enum[K, V]) Values() []V {
	keys := e.Keys()
	var values []V
	for _, key := range keys {
		values = append(values, e.data[key])
	}
	return values
}

// Range 遍历枚举值
func (e *Enum[K, V]) Range(fn func(k K, v V) bool) {
	keys := e.Keys()
	for _, key := range keys {
		if fn(key, e.data[key]) {
			break
		}
	}
	return
}

// RangeWithIndex 遍历枚举值
func (e *Enum[K, V]) RangeWithIndex(fn func(i int, k K, v V) bool) {
	keys := e.Keys()
	for i, key := range keys {
		if fn(i, key, e.data[key]) {
			break
		}
	}
	return
}
