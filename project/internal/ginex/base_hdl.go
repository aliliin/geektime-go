package ginex

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	log "github.com/sirupsen/logrus"
	"net/http"
	"project/pkg/errors"
	"strconv"
)

type BaseHdl struct {
}

func GetContext(c *gin.Context) context.Context {
	return nil
}

func (p *BaseHdl) GetContext(c *gin.Context) context.Context {
	return GetContext(c)
}

func (p *BaseHdl) Copy(toValue interface{}, fromValue interface{}) error {
	err := copier.Copy(toValue, fromValue)
	if err != nil {
		return errors.WrapInternalCusError(err, "内部错误")
	}

	return nil
}

func (p *BaseHdl) HandleError(c *gin.Context, err error) bool {
	return errors.HandleError(c, err)
}

func (p *BaseHdl) ParseIntParam(c *gin.Context, key string) (int, bool) {
	s := c.Param(key)
	i, err := strconv.Atoi(s)
	if err != nil {
		log.WithField("key", key).WithField("value", s).Errorln("bad int param")
		c.String(http.StatusBadRequest, "参数错误")
		c.Abort()
		return 0, false
	}

	return i, true
}

func (p *BaseHdl) Valid(c *gin.Context, v Validator) bool {
	if err := v.Valid(); err != nil {
		log.WithError(err).Errorln("valid form error")

		e, ok := errors.IsCusError(err)
		if ok {
			c.String(http.StatusBadRequest, e.Msg())
		} else {
			c.String(http.StatusBadRequest, err.Error())
		}

		c.Abort()
		return false
	}

	return true
}

type Validator interface {
	Valid() error
}
