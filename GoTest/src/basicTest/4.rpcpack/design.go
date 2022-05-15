// Package rpcpack 该包实现了对rpc的简单封装
// 1.为什么要封装Service方？
// 答：为了让错误尽量出现在编译期，而不是运行期；例如RegisterName函数的入参service为interface{}，
//这样如果开发者传入了错误的参数，那么只有到运行时才会发现错误，而在原来的基础上封装一层，将入参interface{}具体为一个接口，
//那么，开发者就可以在编译期间提早发现问题了。
// 2.为什么要封装Client方？
// 答：（1）同上。
//    （2）以达到更像调用本地方法的效果！

package rpcpack

import (
	"net/rpc"
)

// -------------------定义接口供server使用----------------------------------

// Service 约束service方法
type Service interface {
	HelloWorld(name string, resp *string) error
}

func RegisterService(name string, service Service) error {
	return rpc.RegisterName(name, service)
}

// -----------------封装给client使用，以达到更像调用本地方法----------------------

type MyClient struct {
	C *rpc.Client
}

func NewMyClient(addr string) *MyClient {
	conn, _ := rpc.Dial("tcp", addr)

	return &MyClient{C: conn}
}

// HelloWorld 参照上面接口实现
func (m *MyClient) HelloWorld(name string, resp *string) error {
	return m.C.Call("hello.HelloWorld", name, resp)
}
