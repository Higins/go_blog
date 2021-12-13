package struct/Commnet

import "gorm.io/gorm"

type Commenst struct {
	gorm.Model
	BlogId  string `json:"blogid"`
	Comment string `json:"comment"`
}