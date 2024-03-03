package test

import (
	"github.com/JackD-io/go-model-transformer/model"
	"gorm.io/gorm"
	"time"
)

type User struct {
	model.BaseModel
	Id        uint64 `json:"id" gorm:"primaryKey"`
	Name      string `json:"name"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (b *User) Fields() []string {
	return []string{"Id", "Name", "CreatedAt"}
}
