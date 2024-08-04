package models

import (
	"gorm.io/gorm"
)

type ShortUrl struct {
	gorm.Model
	LongUrl   string `json:"long_url"`
	ShortUrl  string `json:"short_url"`
	ShortCode string `json:"short_code"`
	Visits    uint   `json:"visits"`
}
