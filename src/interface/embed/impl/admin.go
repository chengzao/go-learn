package impl

// 定义实现类 Admin，首字母大写表示 public
type Admin struct {
	User // 嵌入类型
	Name string
}

//func (a *Admin) Notify() {
//	fmt.Println("admin", *a)
//}
