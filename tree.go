package typex

// Convert2Tree 平铺数组转为树形结构
func Convert2Tree[T Node](list []T, root string) Tree[T] {
	if len(list) == 0 {
		return nil
	}
	data := make(map[string]Tree[T])
	for _, item := range list {
		id, pid := item.ID(), item.PID()
		data[pid] = append(data[pid], &TreeNode[T]{
			Id:   id,
			Pid:  pid,
			Data: item,
		})
	}
	if result, ok := data[root]; ok {
		result = result.buildChild(data)
		return result
	}
	return nil
}

type Node interface {
	ID() string
	PID() string
}

// TreeNode 树形节点
type TreeNode[T any] struct {
	Id    string  `json:"id"`
	Pid   string  `json:"pid"`
	Data  T       `json:"data"`
	Child Tree[T] `json:"child"`
}

// Tree 树形结构
type Tree[T any] []*TreeNode[T]

// buildChild 构建子节点
func (t Tree[T]) buildChild(data map[string]Tree[T]) Tree[T] {
	var tree Tree[T]
	if t != nil && len(t) > 0 {
		for _, item := range t {
			var child = data[item.Id]
			child = child.buildChild(data)
			tree = append(tree, &TreeNode[T]{
				Id:    item.Id,
				Pid:   item.Pid,
				Data:  item.Data,
				Child: child,
			})
		}
	}
	return tree
}
