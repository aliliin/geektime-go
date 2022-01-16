package domain

import (
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"github.com/quexer/utee"
)

type User struct {
	Id int // 系统主键

	CreatedAt utee.Tick `validate:"required"`
	UpdatedAt utee.Tick `validate:"required"`

	Openid     string `validate:"required"` // 冗余小程序OPENID
	Nickname   string // 昵称
	Mobile     string // 手机号
	Logo       string // 头像
	Authorized bool   // 用户是否小程序授权
	UnionId    string // union_id
	DeletedAt  int64  // 删除时间
}

func (p *User) Valid() error {
	if err := validator.New().Struct(p); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

type UserList []*User

func (p UserList) FindById(id int) *User {
	for _, v := range p {
		if v.Id == id {
			return v
		}
	}
	return nil
}
