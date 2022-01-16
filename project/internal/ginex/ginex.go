package ginex

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"project/internal/ginex/internal"
)

type GinOption func(*gin.Engine)

// New 创建 gin.Engine, 可指定多个Option
func New(opt ...GinOption) *gin.Engine {
	if !gin.IsDebugging() {
		gin.DisableConsoleColor()
	}

	if err := internal.InitTrans(internal.ZH); err != nil {
		log.Errorln("init trans failed")
	} else {
		log.WithError(err).Infoln("init trans success")
	}

	r := gin.New()
	for _, v := range opt {
		v(r)
	}
	return r
}

// WithRecovery 防止panic
func WithRecovery() GinOption {
	recovery := &internal.GinRecovery{}
	return With(recovery.Recovery())
}

// WithLogger 出错时打打日志
func WithLogger() GinOption {
	return With(internal.Logger())
}

// WithCookieSession CookieSession middleware
// name, cookie name.
// salt, cookie store secret
func WithCookieSession(name, salt string) GinOption {
	store := cookie.NewStore([]byte(salt))
	return With(sessions.Sessions(name, store))
}

// WithStatic 服务静态文件，fileRoot 本地文件路径，默认为 ./public
func WithStatic(fileRoot ...string) GinOption {
	root := "./public"
	if len(fileRoot) > 0 {
		root = fileRoot[0]
	}
	return With(static.ServeRoot("/", root))
}

// WithPprof 启用pprof
func WithPprof() GinOption {
	return func(router *gin.Engine) {
		pprof.Register(router)
	}
}

// WithHSTS 强制使用https
// 详见 https://zh.wikipedia.org/wiki/HTTP%E4%B8%A5%E6%A0%BC%E4%BC%A0%E8%BE%93%E5%AE%89%E5%85%A8
func WithHSTS() GinOption {
	return func(router *gin.Engine) {
		router.Use(internal.MidHSTS)
	}
}

// WithCors 允许 CORS
func WithCors() GinOption {
	return func(router *gin.Engine) {
		router.Use(internal.MidCors)
	}
}

// With 使用任意middleware
func With(fn gin.HandlerFunc) GinOption {
	return func(r *gin.Engine) {
		r.Use(fn)
	}
}
