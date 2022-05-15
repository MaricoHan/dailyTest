package main

import (
	"fmt"
	"log"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func Client() {
	//	1.请求建立rpc连接
	client, err := jsonrpc.Dial("tcp", "localhost:8080")
	if err != nil {
		return
	}
	defer func(client *rpc.Client) {
		if err = client.Close(); err != nil {
			log.Panicf("failed to close listener: %v", err.Error())
		}
	}(client)

	//	2.调用服务
	var resp string // 接收返回值---传出参数

	err = client.Call("hello.HelloWorld", "rpc", &resp)
	if err != nil {
		fmt.Println("call error", err.Error())
		return
	}
	fmt.Println("resp:", resp)
}
func main() {
	Client()
}
