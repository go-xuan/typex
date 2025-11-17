package typex

// Unique 唯一主键接口
type Unique interface {
	GetKey() string // 获取唯一主键
}
