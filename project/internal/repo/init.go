package repo

import (
	"context"
	"project/internal/databases"
	"project/internal/repo/model"

	"github.com/pkg/errors"
)

type DbInitializer struct {
	Db *databases.Data
}

func (p *DbInitializer) Name() string {
	return "db_initializer"
}

func (p *DbInitializer) IsNeedInit(ctx context.Context) (bool, error) {
	return true, nil
}

// Initialize AutoMigrate自动建表
func (p *DbInitializer) Initialize(ctx context.Context) error {
	err := p.Db.Gdb().AutoMigrate(
		&model.User{},
	)
	return errors.WithStack(err)
}
