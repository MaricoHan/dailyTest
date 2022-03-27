package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type Hello struct {
}

func (h *Hello) HelloWorld(name string, resp *string) error {
	*resp = name + " 你好！"
	return nil
}

func Server() {
	//1.rpc注册服务
	err := rpc.RegisterName("hello", new(Hello))
	if err != nil {
		fmt.Println("failed to Register server", err.Error())
		return
	}
	//	2.设置监听
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("failed to listen: ", err.Error())
		return
	}
	fmt.Println("Start to listen...")
	defer listener.Close()
	//	3.建立连接
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("failed to Accept: ", err.Error())
		return
	}
	defer conn.Close()
	fmt.Println("Connect Success...")
	//4.绑定连接
	rpc.ServeConn(conn)
}
func main() {
	Server()
}
