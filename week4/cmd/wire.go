//+build wireinject

package main

import (
	"week4/config"
	"week4/internal/biz"
	"week4/internal/data"
	"week4/internal/server"

	"github.com/google/wire"
)

func InitServer() (*server.UserServer, error) {
	wire.Build(
		server.NewUserServer,
		server.NewGrpcServer,
		server.NewHttpServer,
		biz.NewUserService,
		//biz.NewUserDBService,
		data.NewInMemoUserDao,
		//data.NewUserDao,
		data.NewDB,
		config.NewConfig,
		config.InitConfig)
	return &server.UserServer{}, nil
}
