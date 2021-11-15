package data

import (
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"week4/internal/model"
)

var (
	ErrNotExist     = errors.New("no such record")
	ErrAlreadyExist = errors.New("recored with same id already exist")
)

type UserDaoInf interface {
	Create(user model.User) (bool, error)
	QueryOne(username string) (model.User, error)
	GetUser(username string) (model.User, error)
}

type InMemoUserDao struct {
	users *Data
}

func NewInMemoUserDao(data *Data) UserDaoInf {
	return &InMemoUserDao{
		users: data,
	}
}

type UserDaoInfo interface {
	GetUser(username string) (model.User, error)
}

type UserDao struct {
	data *Data
}

func NewUserDao(data *Data) UserDaoInfo {
	return &UserDao{data: data}
}

func (dao *InMemoUserDao) Create(user model.User) (bool, error) {
	//for _, u := range dao.users {
	//	if u.UserName == user.UserName {
	//		return false, ErrAlreadyExist
	//	}
	//}
	//dao.users = append(dao.users, user)
	return true, nil
}

func (dao *InMemoUserDao) QueryOne(username string) (model.User, error) {
	//for _, u := range dao.users {
	//	if u.UserName == username {
	//		return u, nil
	//	}
	//}
	return model.User{}, ErrNotExist
}

func (dao *InMemoUserDao) GetUser(username string) (model.User, error) {
	var user model.User
	result := dao.users.db.Where(&model.User{UserName: username}).First(&user)
	if result.RowsAffected == 0 {
		return user, status.Errorf(codes.NotFound, "用户不存在")
	}
	if result.Error != nil {
		return user, result.Error
	}
	return model.User{
		UserName: user.UserName,
		Email:    user.Email,
	}, nil
}


func (dao *UserDao) GetUser(username string) (model.User, error) {
	var user model.User
	result := dao.data.db.Where(&model.User{UserName: username}).First(&user)
	if result.RowsAffected == 0 {
		return user, status.Errorf(codes.NotFound, "用户不存在")
	}
	if result.Error != nil {
		return user, result.Error
	}
	return model.User{
		UserName: user.UserName,
		Email:    user.Email,
	}, nil
}
