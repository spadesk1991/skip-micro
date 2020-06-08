package main

import (
	"github.com/opentracing/opentracing-go"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	ocplugin "github.com/micro/go-plugins/wrapper/trace/opentracing"
	"github.com/spadesk1991/skip-micro/app/common"
	"github.com/spadesk1991/skip-micro/app/services/impl"
	"github.com/spadesk1991/skip-micro/app/services/pb"
	"github.com/spadesk1991/skip-micro/app/wrappers"
)

const name = "api.skip.com.server"

func main() {
	reg := etcd.NewRegistry(registry.Addrs("localhost:2379"))
	t, io := common.NewJaegerTracer(name, "localhost:6831")
	defer io.Close()
	opentracing.SetGlobalTracer(t)
	srv := micro.NewService(
		micro.Name(name),
		micro.Registry(reg),
		micro.WrapHandler(wrappers.LogWrapper),
		micro.WrapHandler(ocplugin.NewHandlerWrapper(opentracing.GlobalTracer())),
	)
	pb.RegisterUserServiceHandler(srv.Server(), new(impl.UsersService))
	srv.Init()
	srv.Run()
}
