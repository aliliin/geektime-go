package srv

import (
	"context"
	"project/internal/app/srv/grpc"
	"project/internal/app/srv/http"
	"project/internal/initializer"
	"project/internal/repo"
)

// Bootloader 本应用的启动引导。 每个应用都把自己关心的初始化器加到这里
type Bootloader struct {
	DbInitializer   *repo.DbInitializer
	HttpInitializer *http.Initializer
	GrpcInitializer *grpc.Initializer
}

func (p *Bootloader) getInitializer() []initializer.Initializer {
	return []initializer.Initializer{
		p.DbInitializer,
		p.HttpInitializer,
		p.GrpcInitializer,
	}
}

func (p *Bootloader) Boot(ctx context.Context) error {
	return initializer.InitAll(ctx, p.getInitializer()...)
}
