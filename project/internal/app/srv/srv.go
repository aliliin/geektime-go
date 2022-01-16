package srv

import (
	"fmt"
	"github.com/spf13/viper"
)

type App struct {
	Configs *viper.Viper
}

func (p *App) Run() error {
	// Run("里面不指定端口号默认为8080")
	fmt.Println(p.Configs.GetString("http.port"))
	p.afterStart()
	return nil
}

func (p *App) Init() *App {
	return p
}

func (p *App) afterStart() error {
	return nil
}
