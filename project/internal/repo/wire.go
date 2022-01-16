package repo

import "github.com/google/wire"

var Set = wire.NewSet(
	wire.Struct(new(DbInitializer), "*"),

	wire.Struct(new(UserRepoImpl), "*"),
	wire.Bind(new(UserRepo), new(*UserRepoImpl)),
)
