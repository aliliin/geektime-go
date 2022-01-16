package http

import (
	"context"
	"project/internal/app/srv/http/hdl"
	"project/internal/ginex"

	"github.com/gin-gonic/gin"
)

type Initializer struct {
	Hdl    *hdl.Hdl
	Router *gin.Engine
}

func (p *Initializer) Name() string {
	return "http_initializer"
}

func (p *Initializer) IsNeedInit(ctx context.Context) (bool, error) {
	return true, nil
}

func (p *Initializer) Initialize(ctx context.Context) error {
	bv := ginex.HdlVersion("")
	p.Router.HEAD("/v", bv)
	p.Router.GET("/v", bv)

	p.Hdl.Mount(p.Router) // 挂载业务handler

	return nil
}
