# 变量variable

## 变量定义

- 完全体：`var name type`
- 类型推断：`var name = value`
- 最简体：`name := value`（仅用于函数内变量，包内变量不行）
- 变量聚合定义：`var( name1=value1 name2=value2 )`

## 内建变量类型

- `bool`
- `string`
- `(u)int、(u)int8、(u)int16、(u)int32、(u)int64`
- `uintptr 指针`
- `byte`
- `rune 字符型，32 位，类比 char`
- `float32、float64`
- `complex32、complex64 复数 i = √-1`

## 类型转换

注意：类型转换必须强制转，转化成的类型可以不带小括号

```go
// 强制类型转换
func triangle()  {
    var a, b int = 3, 4
    var c int
    // float64 和 int 可以不加小括号，也可以加上
    // 开方内建函数定义：func Sqrt(x float64) float64
    c = int(math.Sqrt(float64(a*a + b*b)))
    fmt.Println(c)
}
```

## 常量与枚举

- 常量定义姿势（常量必须有 value）
- 完全体：`const name type = value`
- 后续的使用自动补类型：`const name = value`
- 自定义枚举：`const ( name1=value1 name2=value2 )`
- iota 表达式枚举：`const ( name1=iota表达式 name2 )`

## 区分单引号与双引号

- 单引号 `fmt.Println('a') // 97` 
- 双引号 `fmt.Println("a") // a` 
- 反引号 ```fmt.Println(`a`) // a```
