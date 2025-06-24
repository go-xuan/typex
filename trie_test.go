package typex

import (
	"fmt"
	"testing"
)

func TestTrie(t *testing.T) {
	tire := NewTrie()
	tire.AddSlice([]string{"你好", "你好世界", "世界你好", "世界你好世界", "你好世界你好"})
	s := tire.Mask("是谁素不相识你好时间世界你好常见的比赛你好世界你好v的角色")
	fmt.Println(s)
}
