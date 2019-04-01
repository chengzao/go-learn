// 自定义用户异常接口
package exception

type UserError interface {
	error // 内嵌类型
	Message() string
}
