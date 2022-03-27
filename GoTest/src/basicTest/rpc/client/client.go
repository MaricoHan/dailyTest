package main

import (
	"fmt"
	"net/rpc"
)

func Client() {
	//	1.建立rpc连接
	client, err := rpc.Dial("tcp", "localhost:8080")
	if err != nil {
		return
	}
	defer client.Close()
	//	2.调用服务
	var resp string //接收返回值---传出参数

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
