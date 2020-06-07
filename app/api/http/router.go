package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spadesk1991/skip-micro/app/api/http/controllers"
	"github.com/spadesk1991/skip-micro/app/api/http/middleware"
)

func NewGinRouter(prodGrpcSrvs ...interface{}) (router *gin.Engine) {
	router = gin.Default()
	router.NoMethod(middleware.HandleNotFound)
	router.NoRoute(middleware.HandleNotFound)
	router.Use(middleware.ErrHandler())
	//for _, s := range prodGrpcSrvs {
	//	router.Use(middleware.AddGRPCServer(s))
	//}
	initRouter(router)
	return
}
func initRouter(router *gin.Engine) {
	ur := router.Group("/user")
	ur.POST("/", controllers.Registry)
}
