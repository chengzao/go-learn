# Learn golang

URL : https://golang.org

## Dev Env

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

- `go version` | `go env`

# 资料

- [Go 语言极速入门](https://www.cnblogs.com/java-zhao/p/9942311.html) 
