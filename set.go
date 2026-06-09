package typex

import "sync"

// NewSet 创建泛型集合
func NewSet[T comparable]() *Set[T] {
	return &Set[T]{data: make(map[T]struct{})}
}

// Set 泛型集合，并发安全
type Set[T comparable] struct {
	mu   sync.RWMutex
	data map[T]struct{}
}

// Add 添加元素
func (s *Set[T]) Add(v T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[v] = struct{}{}
}

// Has 判断元素是否存在
func (s *Set[T]) Has(v T) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	_, ok := s.data[v]
	return ok
}

// Remove 删除元素
func (s *Set[T]) Remove(v T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.data, v)
}

// Len 返回元素数量
func (s *Set[T]) Len() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.data)
}

// Clear 清空集合
func (s *Set[T]) Clear() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data = make(map[T]struct{})
}

// Range 遍历集合，handle 返回 false 停止遍历
func (s *Set[T]) Range(handle func(v T) bool) {
	s.mu.RLock()
	keys := make([]T, 0, len(s.data))
	for k := range s.data {
		keys = append(keys, k)
	}
	s.mu.RUnlock()
	for _, v := range keys {
		if !handle(v) {
			break
		}
	}
}

// ToSlice 转为切片
func (s *Set[T]) ToSlice() []T {
	s.mu.RLock()
	defer s.mu.RUnlock()
	result := make([]T, 0, len(s.data))
	for k := range s.data {
		result = append(result, k)
	}
	return result
}
