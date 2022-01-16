package databases

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

type Data struct {
	mysql *gorm.DB
}

func (p *Data) Gdb(debug ...bool) *gorm.DB {
	if len(debug) > 0 && debug[0] {
		return p.mysql.Debug()
	}
	return p.mysql
}

func InitDB(v *viper.Viper) (*Data, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		v.Sub("mysql").GetString("user"),
		v.Sub("mysql").GetString("password"),
		v.Sub("mysql").GetString("host"),
		v.Sub("mysql").GetInt("port"),
		v.Sub("mysql").GetString("db"),
	)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢查询 SQL 阈值
			Colorful:      true,        // 禁用彩色打印
			//IgnoreRecordNotFoundError: false,
			LogLevel: logger.Info, // Log lever
		},
	)

	client, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 表名是否加 s
		},
	})
	if err != nil {
		return nil, err
	}

	return &Data{
		mysql: client,
	}, nil
}
