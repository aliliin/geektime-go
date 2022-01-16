package ginex

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

type Session struct {
}

func (Session) WebSession(name, salt string, opt sessions.Options) gin.HandlerFunc {
	store := cookie.NewStore([]byte(salt))
	store.Options(opt)

	return sessions.Sessions(name, store)
}

type Kv struct {
	Key   interface{}
	Value interface{}
}

func (Session) Set(c *gin.Context, key, val interface{}) error {
	se := sessions.Default(c)
	se.Set(key, val)
	return se.Save()
}

func (Session) Get(c *gin.Context, key interface{}) interface{} {
	se := sessions.Default(c)
	v := se.Get(key)
	return v
}

func (Session) Delete(c *gin.Context, key interface{}) error {
	se := sessions.Default(c)
	se.Delete(key)
	return se.Save()
}

func (Session) BatchSet(c *gin.Context, kvs ...Kv) error {
	if len(kvs) == 0 {
		return nil
	}

	se := sessions.Default(c)
	for _, kv := range kvs {
		se.Set(kv.Key, kv.Value)
	}

	return se.Save()
}

func (Session) BatchDelete(c *gin.Context, keys ...interface{}) error {
	if len(keys) == 0 {
		return nil
	}

	se := sessions.Default(c)
	for _, k := range keys {
		se.Delete(k)
	}

	return se.Save()
}
