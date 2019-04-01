package impl

import "fmt"

// 定义实现类 Admin，首字母大写表示 public
type Admin struct {
	Name string
}

func (a *Admin) Notify() {
	fmt.Println("admin", *a)
}
