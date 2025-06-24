package typex

import (
	"strings"
)

// Trie Trie树
type Trie struct {
	root    *TrieNode      // 树根
	mask    string         // 掩码
	weights map[string]int // 权重
}

// TrieNode Trie树节点
type TrieNode struct {
	children map[string]*TrieNode // 子节点
	end      bool                 // 是否是单词词尾
	terminal bool                 // 是否是分支末端
}

// NewTrie 创建前缀树
func NewTrie() *Trie {
	return &Trie{
		root:    newTrieNode(),
		weights: make(map[string]int),
		mask:    "*",
	}
}

func newTrieNode() *TrieNode {
	return &TrieNode{
		children: make(map[string]*TrieNode),
		end:      false,
		terminal: true,
	}
}

// AddMap 添加
func (t *Trie) AddMap(sentences map[string]int) {
	if len(sentences) > 0 {
		for sentence, weight := range sentences {
			t.Add(sentence, weight)
		}
	}
}

// AddSlice 添加文本
func (t *Trie) AddSlice(sentences []string) {
	if len(sentences) > 0 {
		for _, sentence := range sentences {
			t.Add(sentence)
		}
	}
}

// Add 添加句子
func (t *Trie) Add(sentence string, weight ...int) {
	if len(weight) > 0 {
		t.weights[sentence] = weight[0]
	}
	words := strings.Split(sentence, "")
	node := t.root
	for _, word := range words {
		if _, ok := node.children[word]; !ok {
			node.children[word] = newTrieNode()
		}
		node.terminal = false
		node = node.children[word]
	}
	node.end = true
}

// Update 更新权重
func (t *Trie) Update(word string, weight int) {
	t.weights[word] = weight
}

// Mask 文本打码
func (t *Trie) Mask(sentence string) string {
	var sb = strings.Builder{}
	var node, start = t.root, 0
	var temp *TrieNode
	words := strings.Split(sentence, "")
	for i, item := range strings.Split(sentence, "") {
		if child, ok := node.children[item]; ok {
			node = child
		} else if temp != nil {
			if child, ok = temp.children[item]; ok {
				node = child
				temp = nil
			} else {
				sb.WriteString(strings.Join(words[start:i+1], ""))
				start, node = i+1, t.root
			}
		} else {
			sb.WriteString(strings.Join(words[start:i+1], ""))
			start, node = i+1, t.root
		}
		if node.end {
			if !node.terminal {
				temp = node
			}
			for j := start; j <= i; j++ {
				sb.WriteString(t.mask)
			}
			start, node = i+1, t.root
		}
	}
	sb.WriteString(strings.Join(words[start:], ""))
	return sb.String()
}

// Check 判断是否包含
func (t *Trie) Check(sentence string) (bool, int) {
	var exist, weight = false, 0
	var node, start = t.root, 0
	var sb = strings.Builder{}
	var temp *TrieNode
	words := strings.Split(sentence, "")
	for i, word := range words {
		if child, ok := node.children[word]; ok {
			node = child
		} else if temp != nil {
			if child, ok = temp.children[word]; ok {
				node = child
				temp = nil
			} else {
				sb.Reset()
				start, node = i+1, t.root
			}
		} else {
			sb.Reset()
			start, node = i+1, t.root
		}
		if node.end {
			if !node.terminal {
				temp = node
			}
			exist = true
			sb.WriteString(strings.Join(words[start:i+1], ""))
			if max, ok := t.weights[sb.String()]; ok && max > weight {
				weight = max
			}
			start, node = i+1, t.root
		}
	}
	return exist, weight
}
