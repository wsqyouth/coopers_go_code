// Package filter 服务端和客户端过滤器（拦截器）及链式实现
package filter

import (
	"context"
	"sync"
)

// HandleFunc 过滤器（拦截器）函数接口
type HandleFunc func(ctx context.Context, req interface{}, rsp interface{}) (err error)

// Filter 过滤器（拦截器），根据dispatch处理流程进行上下文拦截处理
type Filter func(ctx context.Context, req interface{}, rsp interface{}, f HandleFunc) (err error)

// NoopFilter 空Filter实现
func NoopFilter(ctx context.Context, req interface{}, rsp interface{}, f HandleFunc) (err error) {
	return f(ctx, req, rsp)
}

// FilterChain 链式过滤器
type FilterChain []Filter

// Handle 链式过滤器递归处理流程
func (fc FilterChain) Handle(ctx context.Context, req interface{}, rsp interface{}, f HandleFunc) (err error) {

	n := len(fc)

	//无Filter,执行空Filter
	if n == 0 {
		return NoopFilter(ctx, req, rsp, f)
	}

	//一个Filter,直接处理
	if n == 1 {
		return fc[0](ctx, req, rsp, f)
	}

	//多个Filter,递归执行
	lastI := n - 1
	return func(ctx context.Context, req interface{}, rsp interface{}, f HandleFunc) error {

		var (
			chainFunc HandleFunc
			curI      int
		)
		chainFunc = func(ctx context.Context, req interface{}, rsp interface{}) error {
			if curI == lastI {
				return f(ctx, req, rsp)
			}
			curI++
			err := fc[curI](ctx, req, rsp, chainFunc)
			curI--
			return err
		}
		return fc[0](ctx, req, rsp, chainFunc)
	}(ctx, req, rsp, f)
}

var (
	serverFilters = make(map[string]Filter)
	clientFilters = make(map[string]Filter)
	lock          = sync.RWMutex{}
)

// Register 通过拦截器名字注册server client拦截器
func Register(name string, serverFilter Filter, clientFilter Filter) {
	lock.Lock()
	serverFilters[name] = serverFilter
	clientFilters[name] = clientFilter
	lock.Unlock()
}

// GetServer 获取server拦截器
func GetServer(name string) Filter {
	lock.RLock()
	f := serverFilters[name]
	lock.RUnlock()
	return f
}

// GetClient 获取client拦截器
func GetClient(name string) Filter {
	lock.RLock()
	f := clientFilters[name]
	lock.RUnlock()
	return f
}
