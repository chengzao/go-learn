package impl

import "fmt"

// 定义实现类 user
type User struct {
	Name string
}

func (u *User) Notify() {
	fmt.Println("user", *u)
}
