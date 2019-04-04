# 资源管理与错误处理defer and error

## 资源管理defer

- defer 在函数结束时发生调用
- defer 的调用是栈类型 - 先进后出
- defer 通常用于资源关闭 `Open/Close`，`Lock/UnLock` 等
- defer 的调用机制是 `将defer语句加入栈中，当函数结束时（包括正常执行结束 / return / panic 出错结束等），从栈中依次执行 defer`

## 错误简单处理error

- 通过被调用函数的注释查看其可能发生的错误，然后依据错误类型并进行处理；
- 错误处理结束后要 `return`

## panic & recover

### panic

- 停止当前程序运行
- 一直向上返回，执行每一层的 defer
- 如果没有遇见 recover，程序退出

### recover

- 相当于对 panic 的 catch 语句
- 仅在 defer 调用中使用
- 获取 panic 的值
- 如果无法处理，可重新 panic

## 错误统一处理
