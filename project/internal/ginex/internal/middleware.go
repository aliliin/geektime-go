package internal

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// MidHSTS 强制使用https
// 详见 https://zh.wikipedia.org/wiki/HTTP%E4%B8%A5%E6%A0%BC%E4%BC%A0%E8%BE%93%E5%AE%89%E5%85%A8
func MidHSTS(c *gin.Context) {
	if c.Request.Header.Get("X-Forwarded-Proto") == "https" {
		c.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")
	}
}

// MidCors 允许 CORS
func MidCors(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", c.Request.Header.Get("Origin"))
	c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token, wos-auth-session, tid")
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE")
	c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
	c.Header("Access-Control-Allow-Credentials", "true")
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(http.StatusNoContent)
	}
	c.Next()
}
