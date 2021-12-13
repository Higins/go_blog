package struct/C

type Commenst struct {
	gorm.Model
	BlogId  string `json:"blogid"`
	Comment string `json:"comment"`
}