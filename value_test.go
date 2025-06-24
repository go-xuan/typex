package typex

import (
	"fmt"
	"testing"
)

func TestValue(t *testing.T) {
	v := NewFloat64(123.4)
	fmt.Println("string = ", v.String())
	fmt.Println("int = ", v.Int(111))
	fmt.Println("int64 = ", v.Int64())
	fmt.Println("float64 = ", v.Float64())
	fmt.Println("bool = ", v.Bool())
	fmt.Println("string = ", v.String())
}
