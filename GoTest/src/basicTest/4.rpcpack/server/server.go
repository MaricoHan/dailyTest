package main

import (
	"basicTest/4.rpcpack"
	"fmt"
	"log"
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
	err := rpcpack.RegisterService("hello", new(Hello))
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
	defer func(listener net.Listener) {
		if err = listener.Close(); err != nil {
			log.Panicf("failed to close listener: %v", err.Error())
		}
	}(listener)
	//	3.建立连接
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("failed to Accept: ", err.Error())
		return
	}
	defer func(conn net.Conn) {
		if err = conn.Close(); err != nil {
			log.Panicf("failed to close listener: %v", err.Error())
		}
	}(conn)

	fmt.Println("Connect Success...")
	//4.绑定连接
	rpc.ServeConn(conn)
}
func main() {
	Server()
}
