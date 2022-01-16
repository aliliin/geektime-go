package initializer

import (
	"context"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// Initializer 初始化器
type Initializer interface {
	// Name 初始化器的名称
	Name() string
	// IsNeedInit 是否需要初始化
	IsNeedInit(ctx context.Context) (bool, error)
	// Initialize 初始化数据
	Initialize(ctx context.Context) error
}

// InitAll 初始化传入的 Initializer
func InitAll(ctx context.Context, it ...Initializer) error {
	for _, v := range it {
		name := v.Name()
		need, err := v.IsNeedInit(ctx)
		if err != nil {
			return err
		}

		log := logrus.WithField("name", name)

		if !need {
			log.Println("init not need")
			continue
		}

		err = v.Initialize(ctx)
		if err != nil {
			return errors.WithMessagef(err, "%s init", name)
		}
		log.Println("init completed")
	}
	return nil
}
