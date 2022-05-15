package main

import (
	"basicTest/6.rpcProtobuf/pb"
	"context"
)

type HelloWorld struct {
}

func (h HelloWorld) Hello(ctx context.Context, cin *pb.Cin) (res *pb.Res, err error) {
	tmp := "hello" + cin.GetName()

	return &pb.Res{Res: tmp}, err
}

var _ pb.HelloWorldServer = HelloWorld{}

func Server() {
	// 注册具体服务
	pb.RegisterHelloWorldServer()
}
