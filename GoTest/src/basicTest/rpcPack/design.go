package rpcPack

import "net/rpc"

//-------------------定义接口供server使用----------------------------------

// Service 约束service方法
type Service interface {
	HelloWorld(name string, resp *string) error
}

func RegisterService(name string, service Service) error {
	return rpc.RegisterName(name, service)
}

//-----------------封装给client使用，已达到更像调用本地方法----------------------

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
