package pool

import (
	"errors"
	"io"
	"log"
	"sync"
)

// 声明池类结构体
type Pool struct {
	// 锁
	lock sync.Mutex
	// 池中存储的资源
	resources chan io.Closer
	// 资源创建工厂函数
	factory func() (io.Closer, error)
	// 池是否已经被关闭
	closed bool
}

// 创建池类实例的工厂函数
// 工厂函数名通常使用 New 名字
func New(fn func() (io.Closer, error), size int) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("size too small")
	}

	return &Pool{
		resources: make(chan io.Closer, size),
		factory:   fn,
	}, nil
}

// 从池中获取一个资源
func (p *Pool) Acquire() (io.Closer, error) {
	// select - default 经典模式，将阻塞形式的 channel-go 改为了非阻塞，当 <-p.resources 不能立即返回时，执行 default
	// 当然，如果没有 default，那么还是要阻塞在 <-p.resources 上的
	select {
	// 检查是否有空闲的资源
	case r, ok := <-p.resources:
		log.Println("Acquire:", "Shared Resource")
		if !ok {
			return nil, errors.New("pool already closed")
		}
		return r, nil
	default:
		log.Println("Acquire:", "New Resource")
		// 调用资源创建函数创建资源
		return p.factory()
	}
}

// 将一个使用后的资源放回池里
func (p *Pool) Release(r io.Closer) {
	// 注意：Release 和 Close 使用的是同一把锁，就是说二者同时只能执行一个，防止资源池已经关闭了，release 还向资源池放资源
	// 向一个已经关闭的 channel-go 发送消息，会发生 panic: send on closed channel-go
	p.lock.Lock()
	defer p.lock.Unlock()

	// 如果池已经被关闭，销毁这个资源
	if p.closed {
		r.Close()
		return
	}

	select {
	// 试图将这个资源放入队列
	case p.resources <- r:
		log.Println("Release:", "In Queue")
	default:
		log.Println("Release:", "Closing")
		r.Close()
	}
}

// 关闭资源池，并关闭所有现有的资源
func (p *Pool) Close() {
	p.lock.Lock()
	defer p.lock.Unlock()

	if p.closed {
		return
	}

	p.closed = true

	// 在清空通道里的资源之前，将通道关闭
	close(p.resources)

	// 关闭资源
	for r := range p.resources {
		r.Close()
	}
}
