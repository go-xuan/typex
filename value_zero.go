package typex

// NewZero 创建零值
func NewZero() *Zero {
	return &Zero{}
}

// Zero 零值
type Zero struct{}

func (z *Zero) Valid() bool {
	return false
}

func (z *Zero) String(...string) string {
	return ""
}

func (z *Zero) Int(...int) int {
	return 0
}

func (z *Zero) Int64(...int64) int64 {
	return 0
}

func (z *Zero) Float64(...float64) float64 {
	return 0
}

func (z *Zero) Bool(...bool) bool {
	return false
}
