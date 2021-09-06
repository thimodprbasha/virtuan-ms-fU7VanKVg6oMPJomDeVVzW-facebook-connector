package model

import (
	"gorm.io/gorm"
)

type Requeststring struct {
	Model   `bson:"-"`
	MID_    string `bson:"_id" gorm:"-" json:"-" xml:"mid_"`
	Token   string `bson:"token" json:"token" xml:"token"`
	Message string `bson:"message" json:"message" xml:"message"`
}

func (Requeststring) TableName() string {
	return "requeststring"
}
func (m *Requeststring) PreloadRequeststring(db *gorm.DB) *gorm.DB {
	return db
}
