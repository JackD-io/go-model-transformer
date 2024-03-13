package test

import (
	"github.com/JackDPro/go-model-transformer/model"
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

func (m *User) Fields() []string {
	return []string{"Id", "Name", "CreatedAt"}
}

func (m *User) ToMap() (map[string]interface{}, error) {
	return m.BaseModel.ToMap(m)
}

func (m *User) ToJson() (string, error) {
	return m.BaseModel.ToJson(m)
}
