package typex

import (
	"fmt"
	"testing"
)

func TestLinkedList(t *testing.T) {
	list := NewLink(1).
		Append(2).
		Append(3)

	fmt.Println(list.GetTail())

	list.Remove().
		Remove().
		Remove().
		Remove().
		Remove().
		Remove()
	fmt.Println(list.GetTail())
}
