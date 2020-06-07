package main

import (
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	"github.com/micro/go-micro/web"
)

func main() {
	reg := etcd.NewRegistry(registry.Addrs("localhost:2379"))
	// http 路由
	router := NewGinRouter()
	// web 服务
	webSRV := web.NewService(
		web.Name("api.skip.com.http"),
		web.Address(":3000"),
		web.Registry(reg),
		web.Handler(router),
	)
	_ = webSRV.Init()
	_ = webSRV.Run()
}
