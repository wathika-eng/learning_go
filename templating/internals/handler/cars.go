package handler

import "gorm.io/gorm"

type Cars struct {
	gorm.Model
	ID     int64  `json:"ID"`
	Engine string `json:"engine"`
	Brand  string `json:"brand"`
	Year   int64  `json:"year"`
}
