# typex

泛型类型扩展库，提供枚举、链表、树、值类型等数据结构封装，零外部依赖，是整个 go-xuan 体系的唯一基础模块。

## 安装

```bash
go get github.com/go-xuan/typex
```

## 类型概览

### Enum[K, V] — 并发安全的有序泛型 map

保证插入顺序，读写锁保护，支持遍历、快照和链式调用。

```go
import "github.com/go-xuan/typex"

// 创建
enum := typex.NewStringEnum[string]()  // key 为 string
enum := typex.NewEnum[int, *User]()    // key 为任意 comparable 类型

// 写入
enum.Add("key1", "value1")
enum.Put("key2", "value2")   // Put 是 Add 的别名

// 读取
val := enum.Get("key1")              // 返回零值如果不存在
val, ok := enum.Exist("key1")        // 带 ok 的安全读取

// 删除 & 清空
enum.Remove("key1")
enum.Clear()

// 遍历（handle 返回 true 停止，false 继续）
enum.Range(func(k string, v string) bool {
    fmt.Println(k, v)
    return false // 继续遍历
})

// 带索引遍历
enum.RangeWithIndex(func(i int, k string, v string) bool {
    return false
})

// 批量获取
keys := enum.Keys()       // []string，按插入顺序
values := enum.Values()   // []string，按插入顺序
size := enum.Len()
```

### Link[T] — 线程安全的泛型双向链表

支持头尾追加/删除，互斥锁保护。

```go
list := typex.NewLink("first")
list.Append("second").Append("third")
head, _ := list.GetHead()  // "first"
tail, _ := list.GetTail()  // "third"
list.Remove()              // 删除尾节点
size := list.Size()        // 2
```

### TreeNode[T] — 平铺数组转树形结构

将实现了 `GetID()/GetPID()` 的平铺列表转为嵌套树。

```go
type Dept struct {
    Id       string
    ParentId string
    Name     string
}
func (d Dept) GetID() string  { return d.Id }
func (d Dept) GetPID() string { return d.ParentId }

depts := []Dept{{"1", "0", "总部"}, {"2", "1", "研发部"}}
tree := typex.Convert2Tree(depts, "0")  // "0" 为根节点 pid
typex.PrintTree(tree, "  ")            // 打印缩进树
```

### Value — 多态值类型

统一包装 int/int64/float64/bool/string/time.Time，提供带默认值的类型转换。

```go
// 自动推断类型
v := typex.NewValue(42)          // Int
v := typex.NewValue("hello")     // String
v := typex.NewValue(time.Now())  // Time
v := typex.NewValue(nil)         // Zero

// 带默认值的类型转换
v.String("default")     // 无效时返回 "default"
v.Int(-1)               // 无效时返回 -1
v.Int64(0)
v.Float64(0.0)
v.Bool(false)

// 专用构造函数
v := typex.NewInt(100)
v := typex.NewString("text")
v := typex.NewTime(time.Now())
v := typex.NewZero()     // 始终无效的值

v.Valid()                // 判断是否有效
v.Cover(anotherValue)    // 覆盖当前值
```

### Args — 参数收集器

实现 `Collect[string, Value]` 接口，支持 JSON 反序列化。

```go
args := typex.Args{
    "name": typex.NewString("Alice"),
    "age":  typex.NewInt(30),
}
name := args.Get("name").String()
age := args.Get("age").Int()

// JSON 反序列化（自动推断 value 类型）
json.Unmarshal([]byte(`{"key": "val", "num": 42}`), &args)

// 遍历
typex.ArgsRange(args, func(k string, v typex.Value) {
    fmt.Println(k, v.String())
})
```

### Collect[K, V] — 收集器接口

简单的 Put/Get 抽象，`Enum` 和 `Args` 均实现了该接口。

```go
type Collect[K comparable, V any] interface {
    Put(k K, v V)
    Get(k K) V
}

// 类型别名
typex.CollectAny     // Collect[string, any]
typex.CollectString  // Collect[string, string]
typex.CollectValue   // Collect[string, Value]
```

### Set[T] — 并发安全的泛型集合

```go
s := typex.NewSet[string]()
s.Add("a")
s.Add("b")
s.Has("a")       // true
s.Remove("a")
s.Len()          // 1
s.Clear()

// 遍历（handle 返回 false 停止，与 sync.Map.Range 一致）
s.Range(func(v string) bool {
    fmt.Println(v)
    return true // 继续遍历
})

slice := s.ToSlice()  // []string
```

### JSON — JSON 编解码接口

```go
type JSON interface {
    MarshalJSON() ([]byte, error)
    UnmarshalJSON([]byte) error
}
```

## 完整方法速查

| 类型 | 方法 |
|------|------|
| **Enum[K,V]** | `Add`, `Put`, `Get`, `Exist`, `Remove`, `Clear`, `Keys`, `Values`, `Len`, `Range`, `RangeWithIndex` |
| **Link[T]** | `Append`, `Remove`, `GetHead`, `GetTail`, `Size` |
| **TreeNode[T]** | `Convert2Tree(list, root)`, `PrintTree(tree, indent)` |
| **Value** | `Valid`, `Cover`, `String`, `Int`, `Int64`, `Float64`, `Bool` |
| **Args** | `Put`, `Get`, `ArgsRange` |
| **Set[T]** | `Add`, `Has`, `Remove`, `Clear`, `Len`, `Range`, `ToSlice` |
| **Collect** | `Put`, `Get` |
