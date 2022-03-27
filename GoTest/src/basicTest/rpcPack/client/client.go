package main

import (
	"basicTest/rpcPack"
	"fmt"
)

func Client() {
	//	1.建立rpc连接
	myClient := rpcPack.NewMyClient("localhost:8080")

	defer myClient.C.Close()

	//	2.调用服务
	var resp string //接收返回值---传出参数

	err := myClient.HelloWorld("rpc", &resp) //更像直接调用本地方法
	if err != nil {
		fmt.Println("call error", err.Error())
		return
	}
	fmt.Println("resp:", resp)
}
func main() {
	Client()
}
