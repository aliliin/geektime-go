package model

import (
	"gorm.io/gorm"
	"time"
)

// BaseModel gorm 公共的提取
type BaseModel struct {
	ID          int            `gorm:"primarykey;type:int" json:"id"`
	CreatedAt   time.Time      `gorm:"column:add_time" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"column:update_time" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
	IsDeletedAt bool           `json:"is_deleted_at"`
}
