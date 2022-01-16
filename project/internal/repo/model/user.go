package model

import (
	"github.com/jinzhu/copier"
	"github.com/quexer/utee"
	"project/internal/domain"
)

type User struct {
	BaseModel
	Openid     string `gorm:"column:openid;type:longtext;NOT NULL" json:"openid"`
	Nickname   string `gorm:"column:nickname;type:longtext;NOT NULL" json:"nickname"`
	Mobile     string `gorm:"column:mobile;type:longtext;NOT NULL" json:"mobile"`
	Logo       string `gorm:"column:logo;type:longtext;NOT NULL" json:"logo"`
	Authorized int    `gorm:"column:authorized;type:tinyint(1);NOT NULL" json:"authorized"`
	UnionId    string `gorm:"column:union_id;type:longtext;NOT NULL" json:"union_id"`
}

func (User) TableName() string {
	return "users"
}

func (User) ModelToDomain(x *User) *domain.User {
	return x.ToDomain()
}

func (User) New(x *domain.User) *User {
	out := &User{}
	utee.Chk(copier.Copy(out, x))
	return out
}

func (p *User) ToDomain() *domain.User {
	out := &domain.User{}
	utee.Chk(copier.Copy(out, p))
	return out
}
