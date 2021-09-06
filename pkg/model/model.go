package model

import "time"

import (
	"gorm.io/gorm"
)

type Model struct {
	ID        int64      `gorm:"primary_key" json:"id" swaggerignore:"true"`
	CreatedAt *time.Time `swaggerignore:"true"`
	UpdatedAt *time.Time `swaggerignore:"true"`
	DeletedAt *time.Time `sql:"index" swaggerignore:"true"`
}

type Tabler interface {
	TableName() string
}

func InitModels(db *gorm.DB) {

}
