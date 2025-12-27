package typex

import (
	"fmt"
	"testing"
)

func TestValue(t *testing.T) {
	value := NewValue(123.4)
	value.Cover(1879.3545)
	fmt.Println("string = ", value.String())
	fmt.Println("int = ", value.Int(111))
	fmt.Println("int64 = ", value.Int64())
	fmt.Println("float64 = ", value.Float64())
	fmt.Println("bool = ", value.Bool())
	fmt.Println("string = ", value.String())
}
