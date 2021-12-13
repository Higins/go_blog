package blog/models

import "gorm.io/gorm"

type Blog struct {
	gorm.Model
	Title    string `json:"title"`
	Text     string `json:"text"`
	Commenst []Commenst
}
