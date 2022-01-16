package http

import (
	"github.com/google/wire"
	"project/internal/app/srv/http/hdl"
)

var Set = wire.NewSet(
	wire.Struct(new(Initializer), "*"),
	wire.Struct(new(hdl.Hdl), "*"),
	wire.Struct(new(hdl.Users), "*"),
)
