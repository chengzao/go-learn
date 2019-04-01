package main

import "fmt"

type user struct {
	name  string
	email string
}

// 普通的函数定义 "func 方法名(入参) 返回值"
// 自定义类型的函数定义 "func (接收者) 方法名(入参) 返回值"
// 值传递，拷贝一份 user
func (u user) notify() {
	fmt.Println("pass-by-value", u.name, u.email)
	u.email = "0@qq.com"
}

// 传递指针(即地址)，内部改变会影响外部
func (u *user) changeEmail(newEmail string) {
	// 不需要 (*u).email
	u.email = newEmail
}

// 方法
func main() {
	// 1. user类型的值可以用来调用使用值接收者声明的方法
	bill := user{"bill", "1@qq.com"}
	bill.notify()                // {"bill", "1@qq.com"}
	fmt.Println("1", bill.email) // "1@qq.com"

	// 2. 指向 user 类型值的指针也可以用来调用使用值接收者声明的方法
	lisa := &user{"lisa", "2@qq.com"}
	// 等价于 (*lisa).notify()
	// 注意：把 lisa 指针指向的 user 对象复制了一份,"再强调一次，notify 操作的是一个副本，只不过这次操作的是从 lisa 指针指向的值的副本。"
	lisa.notify()                // {"lisa", "2@qq.com"}
	fmt.Println("2", lisa.email) // "0@qq.com"（错）  2@qq.com（对）

	// 3.user 类型的值可以用来调用使用指针接收者声明的方法
	// 等价于 (&bill).changeEmail ("100@qq.com")，注意 changeEmail 接收的是一个指针
	bill.changeEmail("100@qq.com")
	fmt.Println("3", bill.email) // "100@qq.com"

	// 4.指向 user 类型值的指针可以用来调用使用指针接收者声明的方法
	lisa.changeEmail("200@qq.com")
	fmt.Println("4", lisa.email) // "200@qq.com"
}
