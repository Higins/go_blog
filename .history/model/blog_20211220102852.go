package model

import ""github.com/jinzhu/gorm""

type Blog struct {
	gorm.Model
	Title string `json:"title"`
	Text  string `json:"text"`
}
