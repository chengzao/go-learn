// 核心业务逻辑处理器
package handler

import (
	"defer-error/filelisting/execption"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const PathPrefix = "/list/"

// 实现一个读取文件的 httpServer 处理器
// 假设访问 http://localhost:8888/list/abc.txt
func HandleFileListing(writer http.ResponseWriter, request *http.Request) error {
	// 1. 如果 urlPath 不是以 /list/ 开头的，则自定义用户错误
	if strings.Index(request.URL.Path, PathPrefix) != 0 {
		return exception.MyCustomError("url path need startWith /list/")
	}
	//fmt.Println("path", request.URL.Path)    // /list/abc.txt
	path := request.URL.Path[len(PathPrefix):] // abc.txt 字符串切割，subString

	// 2. 打开文件
	file, err := os.Open(path)
	if err != nil {
		// 遇到错误直接返回，由错误统一处理器进行处理
		return err
	}
	defer file.Close()

	// 3. 读取文件到 byte[]
	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	// 4. 将 byte[] all 写出到响应流
	writer.Write(all)
	return nil
}
