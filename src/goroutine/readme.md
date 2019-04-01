# Coroutine

Goroutine 基于协程 Coroutine，原理总结：

- 如果创建一个 goroutine 并准备运行，这个 goroutine 就会被放到调度器的全局运行队列中。
- 之后，调度器就将这些队列中的 goroutine 分配给一个逻辑处理器，并放到这个逻辑处理器对应的本地运行队列中。
- 本地运行队列中的 goroutine 会一直等待直到自己被分配的逻辑处理器执行

## 使用 go 关键字创建 Goroutine

- 匿名函数实现方式 `go func() {xxx}()`
- 普通函数 funcA 实现方式 `go funcA()`

## 打断正在运行的 Goroutine

- 基于调度器的内部算法，一个正运行的 goroutine 在工作结束前，可以被停止并重新调度。
- 调度器这样做的目的是防止某个 goroutine 长时间占用逻辑处理器。当 goroutine 占用时间过长时，调度器会停止当前正运行的 goroutine，并给其他可运行的 goroutine 运行的机会。

## 设置逻辑处理器数量

```
// 为每个物理处理器分配一个逻辑处理器给调度器使用
runtime.GOMAXPROCS(runtime.NumCPU())
```

## 竞争状态

- 如果两个或者多个 goroutine 在没有互相同步的情况下，访问某个共享的资源，并试图同时读和写这个资源，就处于相互竞争的状态，这种情况被称作竞争状态(race candition)。
- 同一时刻只能有一个 goroutine 对共享资源进行读和写操作

## 锁机制

### 原子类 atomic

### 互斥锁 mutex
