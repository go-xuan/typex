package typex

import "sync"

// NewEnum key为comparable类型，value为任意类型
func NewEnum[K comparable, V any]() *Enum[K, V] {
	return &Enum[K, V]{
		keys: make([]K, 0),
		data: make(map[K]V),
	}
}

// NewStringEnum key为string类型，value为任意类型
func NewStringEnum[V any]() *Enum[string, V] {
	return &Enum[string, V]{
		keys: make([]string, 0),
		data: make(map[string]V),
	}
}

// Enum 枚举类，实现了Collect[K, V]接口
type Enum[K comparable, V any] struct {
	mu   sync.RWMutex // 读写锁
	keys []K          // 保证有序
	data map[K]V      // 数据存储
}

// Put 添加枚举值
func (e *Enum[K, V]) Put(k K, v V) {
	e.Add(k, v)
}

// Get 获取枚举值
func (e *Enum[K, V]) Get(k K) V {
	e.mu.RLock()
	defer e.mu.RUnlock()
	return e.data[k]
}

// Find 判断枚举值是否存在
func (e *Enum[K, V]) Find(k K) (V, bool) {
	e.mu.RLock()
	defer e.mu.RUnlock()
	v, ok := e.data[k]
	return v, ok
}

// Add 添加枚举值
func (e *Enum[K, V]) Add(k K, v V) *Enum[K, V] {
	e.mu.Lock()
	defer e.mu.Unlock()
	if _, ok := e.data[k]; !ok {
		e.keys = append(e.keys, k)
	}
	e.data[k] = v
	return e
}

// Remove 删除枚举值
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

// Len 返回枚举值数量
func (e *Enum[K, V]) Len() int {
	e.mu.RLock()
	defer e.mu.RUnlock()
	return len(e.keys)
}

// Clear 清空枚举值
func (e *Enum[K, V]) Clear() {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.keys = make([]K, 0)
	e.data = make(map[K]V)
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
	e.mu.RLock()
	defer e.mu.RUnlock()
	values := make([]V, len(e.keys))
	for i, key := range e.keys {
		values[i] = e.data[key]
	}
	return values
}

// Range 遍历枚举值，handle 返回 false 停止遍历
func (e *Enum[K, V]) Range(handle func(k K, v V) bool) {
	// 在锁内取快照，锁外执行 handle，避免 handle 回调 Enum 写方法时死锁
	e.mu.RLock()
	keys := make([]K, len(e.keys))
	copy(keys, e.keys)
	vals := make(map[K]V, len(e.data))
	for k, v := range e.data {
		vals[k] = v
	}
	e.mu.RUnlock()
	for _, key := range keys {
		if !handle(key, vals[key]) {
			break
		}
	}
}

// RangeWithIndex 遍历枚举值，包含下标
func (e *Enum[K, V]) RangeWithIndex(handle func(i int, k K, v V) bool) {
	e.mu.RLock()
	keys := make([]K, len(e.keys))
	copy(keys, e.keys)
	vals := make(map[K]V, len(e.data))
	for k, v := range e.data {
		vals[k] = v
	}
	e.mu.RUnlock()
	for i, key := range keys {
		if !handle(i, key, vals[key]) {
			break
		}
	}
}
