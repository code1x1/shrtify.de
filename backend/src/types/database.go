package types

import (
	"gorm.io/gorm"
)

type ShortUrl struct {
	gorm.Model
	Url  string
	Code string `gorm:"unique"`
	Host string
}
