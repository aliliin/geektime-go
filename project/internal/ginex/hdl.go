package ginex

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HdlVersion 显示版本号的handler
func HdlVersion(version string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, fmt.Sprintf("version: %s", version))
	}
}
