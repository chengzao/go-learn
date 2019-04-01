package main

import "fmt"

// user 类
type user struct {
	name       string
	email      string
	ext        int
	privileged bool
}

// admin 类
type admin struct {
	// 自定义类
	person user
	// 内置类型
	level string
}

func main() {
	// 1. 创建 user 变量，所有属性初始化为其零值
	var bill user
	fmt.Println("bill", bill) // {  0 false}

	// 2. 创建 user 变量，并初始化属性值
	lisa := user{
		name:       "nana",
		email:      "117@qq.com",
		ext:        123,
		privileged: true,
	}
	fmt.Println("lisa", lisa) // {nana 117@qq.com 123 true}
	// 直接使用属性值，属性值的顺序要与 struct 中定义的一致
	lisa2 := user{"nana", "117@qq.com", 123, true}
	fmt.Println("lisa2", lisa2) // {nana 117@qq.com 123 true}

	// 3. 含有自定义类型的 struct 进行初始化
	fred := admin{
		person: user{
			name:       "nana",
			email:      "117@qq.com",
			ext:        123,
			privileged: true,
		},
		level: "super",
	}
	fmt.Println("fred:", fred) // fred: {{nana 117@qq.com 123 true} super}
}
