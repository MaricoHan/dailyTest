package main

import (
	"basicTest/4.rpcpack"
	"fmt"
	"log"
	"net/rpc"
)

func Client() {
	//	1.建立rpc连接
	myClient := rpcpack.NewMyClient("localhost:8080")

	defer func(C *rpc.Client) {
		if err := C.Close(); err != nil {
			log.Panicf("failed to close listener: %v", err.Error())
		}
	}(myClient.C)

	//	2.调用服务
	var resp string                            //接收返回值---传出参数
	err := myClient.HelloWorld("2.rpc", &resp) // 更像直接调用本地方法
	if err != nil {
		fmt.Println("call error", err.Error())
		return
	}
	fmt.Println("resp:", resp)
}
func main() {
	Client()
}
