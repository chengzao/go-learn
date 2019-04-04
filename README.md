# Learn golang

URL : [https://golang.org](https://golang.org)

## Dev Version

go version: 1.12.1

## Config Path Env

- mac env

```bash
# go 安装目录
export GOROOT=/usr/local/go
export GOBIN=$GOROOT/bin
# go 工作目录
export GOPATH=$HOME/go
export PATH=$PATH:$GOBIN:$GOPATH/bin
```

- windows env

```bash
# 环境变量配置:
GOROOT : C:\Go
GOPATH : E:\godev
GOBIN : %GOPATH%\bin
PATH : ;%GOROOT%\bin;%GOPATH%;%GOBIN%
```

- `$GOPATH` 目录约定有三个子目录

```bash
src 存放源代码（比如：.go .c .h .s等）
pkg 编译后生成的文件（比如：.a）
bin 编译后生成的可执行文件 # 可以把此目录加入到 $PATH 变量中
```

- `go version` 查看go当前的版本
- `go env` 查看当前go的环境变量
- `go list` 列出当前全部安装的package
- `go run file.go` 编译并运行Go程序

## 命令参数

### go build

- `go build`会忽略目录下以`_`或`.`开头的go文件
- `go build [-o 输出名] [-i] [编译标记] [包名]`
- 如果参数为`***.go`文件或文件列表，则编译为一个个单独的包
- 当编译单个main包（文件），则生成可执行文件
- 当编译单个或多个包非主包时，只构建编译包，但丢弃生成的对象`.a`，仅用作检查包可以构建
- 当编译包时，会自动忽略`_test.go`的测试文件
- 如果某个项目文件夹下有多个文件，而你只想编译某个文件，就可在go build之后加上文件名，例如`go build a.go`；go build命令默认会编译当前目录下的所有go文件

- 参数介绍

```bash
-o 指定输出的文件名，可以带上路径，例如 go build -o a/b/c

-i 安装相应的包，编译+go install

-a 完全编译，不理会-i产生的.a文件

-n 仅打印输出build需要的命令，不执行build动作

-p n 开多少核cpu来并行编译，默认为本机CPU核数

-race 同时检测数据竞争状态，只支持 linux/amd64, freebsd/amd64, darwin/amd64 和 windows/amd64

-v 打印出被编译的包名

-work 打印临时工作目录的名称，并在退出时不删除它

-x 同时打印输出执行的命令名, 与-n的结果类似，只是这个会执行build动作

-tags 'tag list' 构建出带tag的版本

-msan 启用与内存消毒器的互操作

-ldflags 'flag list'
    '-s -w': 压缩编译后的体积
    -s: 去掉符号表
    -w: 去掉调试信息，不能gdb调试了
```

### go get

- 第一步是下载源码包，第二步是执行go install
- 参数介绍

```bash
-d 只下载不安装

-f 只有在你包含了-u参数的时候才有效

-fix 在获取源码之后先运行fix，然后再去做其他的事情

-t 同时也下载需要为运行测试所需要的包

-u 强制使用网络去更新包和它的依赖包

-v 显示执行的命令
```

### go clean

- 用来移除当前源码包和关联源码包里面编译生成的文件
- 参数介绍

```bash
-i 清除关联的安装的包和可运行文件，也就是通过go install安装的文件

-n 把需要执行的清除命令打印出来，但是不执行，这样就可以很容易的知道底层是如何运行的

-r 循环的清除在import中引入的包

-x 打印出来执行的详细命令，其实就是-n打印的执行版本
```

### go install

- 命令在内部实际上分成了两步操作：
- 第一步是生成结果文件(`可执行文件`或者`.a`包)，
- 第二步会把编译好的结果移到`$GOPATH/pkg`或者`$GOPATH/bin`
- 参数支持`go build`的编译参数, 一般用`-v`

### go test

- 执行这个命令，会自动读取源码目录下面名为`*_test.go`的文件，生成并运行测试用的可执行文件

- 常用的参数

```bash
-bench regexp 执行相应的benchmarks，例如 -bench=.

-cover 开启测试覆盖率

-run regexp 只运行regexp匹配的函数，例如 -run=Array 那么就执行包含有Array开头的函数

-v 显示测试的详细命令
```

### go tool

- `go tool fix .` 用来修复以前老版本的代码到新版本
- `go tool vet directory|files` 用来分析当前目录的代码是否都是正确的代码

## gofmt

- 使用go fmt命令，其实是调用了gofmt

- 参数介绍

```bash
-l 显示那些需要格式化的文件

-w 把改写后的内容直接写入到文件中，而不是作为结果打印到标准输出

-r 添加形如“a[b:len(a)] -> a[b:]”的重写规则，方便我们做批量替换

-s 简化文件中的代码

-d 显示格式化前后的diff而不是写入文件，默认是false

-e 打印所有的语法错误到标准输出。如果不使用此标记，则只会打印不同行的前10个错误

-cpuprofile 支持调试模式，写入相应的cpufile到指定的文件
```

## godoc

命令行运行`godoc -http=:8080` 启动本地文档，浏览器中打开[127.0.0.1:8080](http://127.0.0.1:8080)

## gopm

- [Github: @gpmgo/gopm](https://github.com/gpmgo/gopm)
- 安装: `go get -v -u github.com/gpmgo/gopm`
- 使用: `gopm get -u -v -g url` && `go install url`

## 资料

- [Go 语言极速入门](https://www.cnblogs.com/java-zhao/p/9942311.html)
- [build-web-application-with-golang](https://github.com/astaxie/build-web-application-with-golang)
- [golang.google.cn](https://golang.google.cn/doc/)
- [goproxy.io](https://goproxy.io)
