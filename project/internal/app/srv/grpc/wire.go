package grpc

import (
	"github.com/google/wire"
)

var Set = wire.NewSet(
	wire.Struct(new(Initializer), "*"),
	wire.Struct(new(UserService), "*"),
)
