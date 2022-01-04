package blogController

import (
	"encoding/json"
	"fmt"
	"web/db"
	"web/helper"
	"web/model"

	"github.com/gofiber/fiber"
	log "github.com/sirupsen/logrus"
)

func Newblog(c *fiber.Ctx) {
	var blogRequest model.Blog
	data := c.Body()
	if err := json.Unmarshal([]byte(data), &blogRequest); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}

	db := *db.DBConn
	var blog model.Blog
	blog.Title = blogRequest.Title
	blog.Text = blogRequest.Text
	log.WithFields(log.Fields{
		"title": blogRequest.Title,
		"text":  blogRequest.Text}).Info("New blog write")
	db.Create(&blog)
	c.JSON(blog)
}

func Getallblog(c *fiber.Ctx) {
	db := *db.DBConn

	blog := make([]model.Blog, 0)
	db.Debug().Scopes(helper.Paginate(c)).Preload("Commenst").Find(&blog)
	c.JSON(blog)
}
