package error

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/quexer/utee"
	log "github.com/sirupsen/logrus"
)

const (
	ErrCodeOk           ErrCode = 0
	ErrCodeBadReq       ErrCode = 400
	ErrCodeUnauthorized ErrCode = 401
	ErrCodeForbidden    ErrCode = 403
	ErrCodeNotFound     ErrCode = 404
	ErrCodeInternalErr  ErrCode = 500
)

type ErrCode int

func (p ErrCode) IsOk() bool {
	return p == ErrCodeOk
}

type CusError struct {
	code    ErrCode
	msg     string
	context utee.J
	err     error
}

func (p *CusError) Error() string {
	return fmt.Sprintf("%s:%+v", p.msg, p.err)
}

func (p *CusError) Msg() string {
	return p.msg
}

func (p *CusError) Code() ErrCode {
	return p.code
}

func (p *CusError) Context() utee.J {
	return p.context
}

func IsCusError(err error) (*CusError, bool) {
	if err == nil {
		return nil, false
	}

	e, ok := err.(*CusError)
	if !ok {
		return nil, ok
	}

	return e, true

}

func WrapCusErr(code ErrCode) func(err error, msg string, contexts ...utee.J) error {
	return func(err error, msg string, contexts ...utee.J) error {
		return wrapCusError(code, err, msg, contexts...)
	}
}

func WrapBadRequestCusError(err error, msg string, contexts ...utee.J) error {
	return wrapCusError(ErrCodeBadReq, err, msg, contexts...)
}

func WrapNotFoundCusError(err error, msg string, contexts ...utee.J) error {
	return wrapCusError(ErrCodeNotFound, err, msg, contexts...)
}

func WrapForbiddenCusError(err error, msg string, contexts ...utee.J) error {
	return wrapCusError(ErrCodeForbidden, err, msg, contexts...)
}

func WrapUnauthorizedCusError(err error, msg string, contexts ...utee.J) error {
	return wrapCusError(ErrCodeUnauthorized, err, msg, contexts...)
}

func WrapInternalCusError(err error, msg string, contexts ...utee.J) error {
	return wrapCusError(ErrCodeInternalErr, err, msg, contexts...)
}

func wrapCusError(code ErrCode, err error, msg string, contexts ...utee.J) error {
	if code.IsOk() {
		return nil
	}

	if err == nil {
		err = errors.New(msg)
	}

	e := &CusError{
		code: code,
		msg:  msg,
		err:  err,
	}

	if len(contexts) > 0 {
		e.context = contexts[0]
	}

	return e
}

func GetErrorMsg(err error) string {
	if err == nil {
		return ""
	}

	if e, ok := IsCusError(err); ok {
		return e.Msg()
	}

	return err.Error()
}

func HandleError(c *gin.Context, err error) bool {
	if err == nil {
		return false
	}

	fmt.Printf("%+v", err) // 打印到标准输出，方便查错

	ce, ok := IsCusError(err)
	if !ok {
		// 如果是未包装错误， 现在包装， 下面统一处理
		ce = WrapInternalCusError(err, "服务错误，请稍后重试").(*CusError)
	}

	lg := log.WithError(ce)
	for k, v := range ce.Context() {
		lg = lg.WithField(k, v)
	}

	lg.Errorln(ce.Msg())

	if ce.Code() > 500 {
		c.String(http.StatusInternalServerError, ce.Msg())
		c.Abort()
		return true
	}

	c.String(int(ce.Code()), ce.Msg())
	c.Abort()
	return true
}
