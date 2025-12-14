package typex

import (
	"fmt"
	"testing"
)

func TestEnum(t *testing.T) {
	// 声明key为string，value为float64类型的枚举，两种声明方式等价
	sf1 := NewStringEnum[float64]()   // *Enum[string, float64]
	sf2 := NewEnum[string, float64]() // *Enum[string, float64]

	sf1.Add("1", 111.111).
		Add("2", 222.222).
		Add("3", 333.333)

	sf2.Add("1", 111.111).
		Add("2", 222.222).
		Add("3", 333.333)

	fmt.Println(sf1.Get("1") == sf2.Get("1"))
	fmt.Println(sf1.Values())
	sf1.Clear()
	fmt.Println(sf1.Values())

	fmt.Println("=============================")

	// 声明key为int类型，value为任意类型的枚举
	ia := NewEnum[int, any]()
	ia.Add(1, 1111).
		Add(2, "AAA").
		Add(3, 454.33).
		Add(4, true).
		Add(5, sf2.Values())

	// 遍历枚举
	ia.Range(func(k int, v any) bool {
		fmt.Println(k, v)
		return false
	})

}
