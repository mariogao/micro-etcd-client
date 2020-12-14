package main

import (
	"context"
	"fmt"

	pb "github.com/mariogao/micro-etcd-server/proto"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
)

func main() {
	service := micro.NewService(
		micro.Registry(etcd.NewRegistry(registry.Addrs("127.0.0.1:2379"))),
	)
	service.Init()
	req := pb.NewHiService("helloworld", service.Client())
	reqParam := &pb.HelloReq{
		Name: "mario ",
	}
	res, err := req.MyHello(context.Background(), reqParam)
	fmt.Print("res:", res, "err:", err)

}
