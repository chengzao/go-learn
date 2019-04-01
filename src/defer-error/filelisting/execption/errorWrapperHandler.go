// 统一错误处理器
package exception

import (
	"log"
	"net/http"
	"os"
)

// 定义一个 function 类型的 type，返回值是 error
type appHandler func(writer http.ResponseWriter, request *http.Request) error

// 输入 appHandler 是一个函数，输出也是一个函数 - 函数式编程
func ErrWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		// 1. 处理业务逻辑
		err := handler(writer, request)
		if err != nil {
			log.Printf("error occured, %s", err) // 2018/11/04 10:10:12 error occured, open abc.txt1: no such file or directory

			// 2. 处理可以抛给用户的错误
			if err, ok := err.(UserError); ok {
				// 将错误写回到 http.ResponseWriter
				http.Error(writer, err.Message(), http.StatusBadRequest)
			}

			// 3. 处理不可以抛给用户的错误
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code) // 浏览器：Not Found
		}
	}
}
