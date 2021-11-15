package biz

import (
	"context"

	"github.com/pkg/errors"

	"week4/internal/data"
	"week4/internal/model"
)

type UserService struct {
	userDao data.UserDaoInf
}

func NewUserService(dao data.UserDaoInf) *UserService {
	return &UserService{
		userDao: dao,
	}
}

type UserServices struct {
	userDaos data.UserDaoInfo
}

func NewUserDBService(dao data.UserDaoInfo) *UserServices {
	return &UserServices{
		userDaos: dao,
	}
}

func (service *UserService) GetUser(ctx context.Context, username string) (model.User, error) {
	return service.userDao.GetUser(username)
}

func (service *UserService) Register(ctx context.Context, user model.User) (bool, error) {
	return service.userDao.Create(user)
}

func (service *UserService) Login(ctx context.Context, username string, password string) (bool, error) {
	user, err := service.userDao.QueryOne(username)
	if err != nil {
		if errors.Is(err, data.ErrNotExist) {
			return false, errors.Wrap(err, "user not exist")
		} else {
			return false, err
		}
	}
	return user.Password == password, nil
}
