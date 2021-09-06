package model

import (
	"gorm.io/gorm"
)

type Responsestring struct {
	Model     `bson:"-"`
	MID_      string `bson:"_id" gorm:"-" json:"-" xml:"mid_"`
	Createdid string `bson:"createdid" json:"createdid" xml:"createdid"`
	ReturnData   string `bson:"returndata" json:"returndata" xml:"returndata"`
}

func (Responsestring) TableName() string {
	return "responsestring"
}
func (m *Responsestring) PreloadResponsestring(db *gorm.DB) *gorm.DB {
	return db
}

