package repo

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"project/internal/domain"
	"project/internal/repo/model"
)

//go:generate mockgen -destination=../mocks/mrepo/user.go -package=mrepo . UserRepo

type UserRepo interface {
	MustGet(ctx context.Context, id int) (*domain.User, error)
	Create(ctx context.Context, in *domain.User) error
}

type UserRepoImpl struct {
	db *gorm.DB
}

func (u *UserRepoImpl) MustGet(c context.Context, id int) (*domain.User, error) {
	var o model.User
	if err := u.db.Take(&o, id).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	return o.ToDomain(), nil
}

func (u *UserRepoImpl) Create(c context.Context, in *domain.User) error {
	if err := in.Valid(); err != nil {
		return err
	}
	obj := model.User{}.New(in)
	if err := u.db.Create(obj).Error; err != nil {
		return errors.WithStack(err)
	}
	in.Id = obj.ID
	return nil
}
