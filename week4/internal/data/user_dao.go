package data

import (
	"errors"

	"week4/internal/model"
)

var (
	ErrNotExist     = errors.New("no such record")
	ErrAlreadyExist = errors.New("recored with same id already exist")
)

type UserDaoInf interface {
	Create(user model.User) (bool, error)
	QueryOne(username string) (model.User, error)
}

type InMemoUserDao struct {
	users []model.User
}

func NewInMemoUserDao() UserDaoInf {
	return &InMemoUserDao{
		users: make([]model.User, 0),
	}
}

func (dao *InMemoUserDao) Create(user model.User) (bool, error) {
	for _, u := range dao.users {
		if u.UserName == user.UserName {
			return false, ErrAlreadyExist
		}
	}
	dao.users = append(dao.users, user)
	return true, nil
}

func (dao *InMemoUserDao) QueryOne(username string) (model.User, error) {
	for _, u := range dao.users {
		if u.UserName == username {
			return u, nil
		}
	}
	return model.User{}, ErrNotExist
}
