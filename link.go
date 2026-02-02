package typex

import "sync"

// NewLink 创建链表
func NewLink[T any](value T) *Link[T] {
	node := &linkNode[T]{
		value: value,
	}
	return &Link[T]{
		mutex: new(sync.Mutex),
		size:  1,
		head:  node,
		tail:  node,
	}
}

// 链表节点
type linkNode[T any] struct {
	value T            // 节点数据
	prev  *linkNode[T] // 指向上一个节点
	next  *linkNode[T] // 指向下一个节点
}

type Link[T any] struct {
	mutex *sync.Mutex  // 锁
	size  int          // 链表大小
	head  *linkNode[T] // 头节点
	tail  *linkNode[T] // 尾节点
}

func (list *Link[T]) Size() int {
	return list.size
}

// GetHead 获取头节点数据
func (list *Link[T]) GetHead() (T, bool) {
	if list.head == nil {
		return *new(T), false
	}
	return list.head.value, true
}

// GetTail 获取尾节点数据
func (list *Link[T]) GetTail() (T, bool) {
	if list.tail == nil {
		return *new(T), false
	}
	return list.tail.value, true
}

// Append 追加节点，每次追加一个节点
func (list *Link[T]) Append(value T) *Link[T] {
	list.mutex.Lock()
	defer list.mutex.Unlock()

	node := &linkNode[T]{
		value: value,
	}
	if list.tail == nil {
		list.head = node
		list.tail = node
	} else {
		node.prev = list.tail
		list.tail.next = node
		list.tail = node
	}
	list.size++
	return list
}

// Remove 删除尾节点，每次删除一个节点，若链表为空则不操作
func (list *Link[T]) Remove() *Link[T] {
	list.mutex.Lock()
	defer list.mutex.Unlock()

	if list.tail == nil {
		return list
	}
	if list.tail.prev == nil {
		list.head = nil
		list.tail = nil
	} else {
		list.tail.prev.next = nil
		list.tail = list.tail.prev
	}
	list.size--
	return list
}
