package commentController

import (
	"encoding/json"
	"fmt"
	"web/db"
	"web/model"

	"github.com/gofiber/fiber"
	log "github.com/sirupsen/logrus"
)

func Newcomments(c *fiber.Ctx) {
	var commentsRequest model.Commenst
	data := c.Body()
	if err := json.Unmarshal([]byte(data), &commentsRequest); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}

	db := *db.DBConn
	var conmment model.Commenst
	var blogId = c.Params("blogid")
	var comment = commentsRequest.Comment
	conmment.BlogId = blogId
	conmment.Comment = comment
	log.WithFields(log.Fields{
		"blogid":  blogId,
		"Comment": comment}).Info("New comment write")
	db.Create(&conmment)
	c.JSON(conmment)
}
