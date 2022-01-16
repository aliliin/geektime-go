package hdl

import (
	"github.com/gin-gonic/gin"
)

// Hdl web handler聚合，所有业务handler依赖注入到这里
// 如果handler过多， 可考虑把相关的几个handler聚合到一起（共享url 前缀）
type Hdl struct {
	Hello *Users
}

func (p *Hdl) Mount(router gin.IRouter) {
	g := router.Group("/")
	p.Hello.Mount(g.Group("/hello"))
}
