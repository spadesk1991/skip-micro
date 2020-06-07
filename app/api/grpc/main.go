package main

import (
	"github.com/spadesk1991/skip-micro/app/services/pb"
	"github.com/spadesk1991/skip-micro/app/wrappers"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	"github.com/spadesk1991/skip-micro/app/services/impl"
)

func main() {
	reg := etcd.NewRegistry(registry.Addrs("localhost:2379"))
	srv := micro.NewService(
		micro.Name("api.skip.com.server"),
		micro.Registry(reg),
		micro.WrapHandler(wrappers.LogWrapper),
	)
	pb.RegisterUserServiceHandler(srv.Server(), new(impl.UsersService))
	srv.Init()
	srv.Run()
}
