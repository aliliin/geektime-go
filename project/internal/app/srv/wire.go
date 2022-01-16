//go:build wireinject
// +build wireinject

package srv

import (
	"github.com/google/wire"
	"project/configs"
)

func RunSrv() (*App, func(), error) {
	panic(wire.Build(
		configs.NewConfig,
		wire.Struct(new(App), "*"),
	))
}
