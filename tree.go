package typex

import (
	"fmt"
	"strings"
)

// 树节点接口
type treeNode interface {
	GetID() string
	GetPID() string
}

// TreeNode 树形节点
type TreeNode[T any] struct {
	Id       string         `json:"id"`
	Pid      string         `json:"pid"`
	Depth    int            `json:"depth"`
	Data     T              `json:"data"`
	Children []*TreeNode[T] `json:"children"`
}

// Convert2Tree 平铺数组转为树形结构
func Convert2Tree[N treeNode](list []N, root string) []*TreeNode[N] {
	if len(list) == 0 {
		return nil
	}
	// 按照pid进行分组聚合
	origin := make(map[string][]*TreeNode[N])
	for _, item := range list {
		id, pid := item.GetID(), item.GetPID()
		origin[pid] = append(origin[pid], &TreeNode[N]{
			Id:   id,
			Pid:  pid,
			Data: item,
		})
	}
	children, ok := origin[root]
	if !ok {
		return nil
	}
	buildChildren(children, origin, 0)
	return children
}

// buildChildren 构建子节点
func buildChildren[T any](children []*TreeNode[T], origin map[string][]*TreeNode[T], depth int) {
	for _, child := range children {
		child.Depth = depth
		thisChildren := origin[child.Id]
		buildChildren(thisChildren, origin, child.Depth+1)
		child.Children = thisChildren
	}
}

// PrintTree 打印树形结构
func PrintTree[T any](tree []*TreeNode[T], indent string) {
	for _, node := range tree {
		printNode(indent, node)
	}
}

func printNode[T any](indent string, node *TreeNode[T]) {
	fmt.Println(strings.Repeat(indent, node.Depth) + node.Id)
	if node.Children != nil && len(node.Children) > 0 {
		for _, child := range node.Children {
			printNode(indent, child)
		}
	}
}
