package db

import (
	"fmt"
	"web/model"

	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
)

func InitDatabase() {
	var err error
	DBConn, err = gorm.Open("sqlite3", "blog.db")
	if err != nil {
		fmt.Println("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")
	var blog model.Blog
	var comments model.Commenst

	DBConn.AutoMigrate(&blog, &comments)

}
