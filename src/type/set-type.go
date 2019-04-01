package main

import "fmt"

type user struct {
	name  string
	email string
}

// 注意该方法是 user 的方法（接收者为 user）
func (u *user) notify() {
	fmt.Println("notify", *u)
}

type admin struct {
	// 嵌入类型: "要嵌入一个类型，只需要声明这个类型的名字就可以了"
	// 注意：不是 u user, u user 是声明字段
	// user 是外部类型 admin 的内部类型
	user
	level string
}

func main() {
	ad := admin{
		user:  user{"nana", "110@qq.com"},
		level: "super",
	}
	// 我们可以直接访问内部类型的方法
	ad.user.notify() // notify {nana 110@qq.com}
	// 内部类型的方法也被提升到外部类型
	ad.notify() // notify {nana 110@qq.com}
}
