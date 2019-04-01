// 自定义用户异常实现
package exception

// 基于基本类型创建自定义类型
type MyCustomError string

func (e MyCustomError) Error() string {
	return e.Message()
}

func (e MyCustomError) Message() string {
	return string(e)
}
