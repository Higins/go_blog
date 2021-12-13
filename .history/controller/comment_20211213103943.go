package controller

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber"
)

func newComments(c *fiber.Ctx) {
	var commentsRequest Commenst
	data := c.Body()
	if err := json.Unmarshal([]byte(data), &commentsRequest); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}

	db := DBConn
	var conmment Commenst
	conmment.BlogId = commentsRequest.BlogId
	conmment.Comment = commentsRequest.Comment
	log.WithFields(log.Fields{
		"blogid":  commentsRequest.BlogId,
		"Comment": commentsRequest.Comment}).Info("New comment write")
	db.Create(&conmment)
	c.JSON(conmment)
}
