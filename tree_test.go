package typex

import (
	"fmt"
	"testing"
)

type Demo struct {
	Id   string
	Pid  string
	Name string
}

func (d Demo) ID() string {
	return d.Id
}

func (d Demo) PID() string {
	return d.Pid
}

func TestTree(t *testing.T) {
	var list = []Demo{
		{Id: "1", Pid: "0", Name: "1"},
		{Id: "2", Pid: "0", Name: "2"},
		{Id: "3", Pid: "0", Name: "3"},
		{Id: "11", Pid: "1", Name: "1-1"},
		{Id: "12", Pid: "1", Name: "1-2"},
		{Id: "13", Pid: "1", Name: "1-3"},
		{Id: "14", Pid: "1", Name: "1-4"},
		{Id: "15", Pid: "1", Name: "1-5"},
		{Id: "21", Pid: "2", Name: "2-1"},
		{Id: "22", Pid: "2", Name: "2-2"},
		{Id: "23", Pid: "2", Name: "2-3"},
		{Id: "31", Pid: "3", Name: "3-1"},
		{Id: "32", Pid: "3", Name: "3-2"},
		{Id: "33", Pid: "3", Name: "3-3"},
		{Id: "34", Pid: "3", Name: "3-4"},
		{Id: "35", Pid: "3", Name: "3-5"},
		{Id: "36", Pid: "3", Name: "3-6"},
		{Id: "111", Pid: "11", Name: "1-1-1"},
		{Id: "112", Pid: "11", Name: "1-1-2"},
		{Id: "113", Pid: "11", Name: "1-1-3"},
		{Id: "121", Pid: "12", Name: "1-2-1"},
		{Id: "122", Pid: "12", Name: "1-2-2"},
		{Id: "123", Pid: "12", Name: "1-2-3"},
		{Id: "131", Pid: "13", Name: "1-3-1"},
		{Id: "132", Pid: "13", Name: "1-3-2"},
		{Id: "133", Pid: "13", Name: "1-3-3"},
	}

	tree := Convert2Tree(list, "0")

	for _, node := range tree {
		printNode("", node)
	}
}

func printNode[T any](s string, node *TreeNode[T]) {
	fmt.Println(s + node.Id)
	if node.Child != nil && len(node.Child) > 0 {
		for _, child := range node.Child {
			printNode(s+"  ", child)
		}
	}
}
