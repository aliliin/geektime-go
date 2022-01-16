package srv

import (
	"context"
	"github.com/spf13/viper"
	"project/internal/databases"
)

type App struct {
	Configs    *viper.Viper
	Bootloader *Bootloader
	DB         *databases.Data
}

func (p *App) Run() error {
	// Run("里面不指定端口号默认为8080")
	p.Bootloader.HttpInitializer.Router.Run(":" + p.Configs.GetString("http.port"))
	p.afterStart()
	return nil
}

func (p *App) Init() *App {
	err := p.Bootloader.Boot(context.Background())
	if err != nil {
		panic(err)
	}
	return p
}

func (p *App) afterStart() error {
	return nil
}
