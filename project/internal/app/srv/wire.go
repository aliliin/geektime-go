//go:build wireinject
// +build wireinject

package srv

import (
	"github.com/google/wire"
	proto "project/api/v1"
	"project/configs"
	"project/internal/app/srv/grpc"
	"project/internal/app/srv/http"
	"project/internal/databases"
	"project/internal/initializer"
)

func RunSrv() (*App, error) {
	panic(wire.Build(
		configs.NewConfig,
		databases.InitDB,
		initializer.BaseSet,
		// 启动一个服务 grpc http
		wire.Struct(new(App), "*"),
		wire.Struct(new(Bootloader), "*"),

		http.Set,
		grpc.Set,
		proto.Set,
	))
}
