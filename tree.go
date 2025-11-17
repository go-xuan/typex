package typex

// Convert2Tree 平铺数组转为树形结构
func Convert2Tree[N Node](list []N, root string) []*TreeNode[N] {
	if len(list) == 0 {
		return nil
	}
	data := make(map[string][]*TreeNode[N])
	for _, item := range list {
		id, pid := item.GetID(), item.GetPID()
		data[pid] = append(data[pid], &TreeNode[N]{
			Id:   id,
			Pid:  pid,
			Data: item,
		})
	}
	if tree, ok := data[root]; ok {
		tree = buildChild(tree, data)
		return tree
	}
	return nil
}

type Node interface {
	GetID() string
	GetPID() string
}

// TreeNode 树形节点
type TreeNode[T any] struct {
	Id    string         `json:"id"`
	Pid   string         `json:"pid"`
	Data  T              `json:"data"`
	Child []*TreeNode[T] `json:"child"`
}

// buildChild 构建子节点
func buildChild[T any](nodes []*TreeNode[T], groups map[string][]*TreeNode[T]) []*TreeNode[T] {
	var list []*TreeNode[T]
	if nodes != nil && len(nodes) > 0 {
		for _, node := range nodes {
			var group = groups[node.Id]
			group = buildChild(group, groups)
			list = append(list, &TreeNode[T]{
				Id:    node.Id,
				Pid:   node.Pid,
				Data:  node.Data,
				Child: group,
			})
		}
	}
	return list
}
