# go module

## go.mod

有四种指令：module，require，exclude，replace

- `module`：模块名称
- `require`：依赖包列表以及版本
- `exclude`：禁止依赖包列表（仅在当前模块为主模块时生效）
- `replace`：替换依赖包列表 （仅在当前模块为主模块时生效）

## cmd

- `go mod init [module name]` 项目文件夹下初始化模块名
- `go mod tidy` //拉取缺少的模块，移除不用的模块
- `go mod download` //下载依赖包
- `go mod graph` //打印模块依赖图
- `go mod vendor` //将依赖复制到vendor下
- `go mod verify` //校验依赖
- `go mod why` //解释为什么需要依赖
- `go list -m -json all` //依赖详情


## 资源

- [拜拜了，GOPATH君！新版本Golang的包管理入门教程](https://zhuanlan.zhihu.com/p/60703832)
- [Go Modules使用教程](https://segmentfault.com/a/1190000016703769)
