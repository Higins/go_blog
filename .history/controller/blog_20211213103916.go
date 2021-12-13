package controller

import (
	"log"

	"github.com/gofiber/fiber"
)

func getAllBlog(c *fiber.Ctx) {
	db := DBConn
	data, error := db.Table("blogs").Joins("join commensts c on c.blog_id = blogs.id").Select("*").Rows()
	if error != nil {
		log.Panic(error)
	}
	defer data.Close()
	blogNew := Blog{}
	var CommenstItem Commenst
	for data.Next() {

		var err = data.Scan(
			&blogNew.ID,
			&blogNew.CreatedAt,
			&blogNew.UpdatedAt,
			&blogNew.DeletedAt,
			&blogNew.Title,
			&blogNew.Text,
			&CommenstItem.ID,
			&CommenstItem.CreatedAt,
			&CommenstItem.UpdatedAt,
			&CommenstItem.DeletedAt,
			&CommenstItem.BlogId,
			&CommenstItem.Comment)
		if err != nil {
			log.Panic(err)
		}

		blogNew.Commenst = append(blogNew.Commenst, CommenstItem)

	}
	c.JSON(blogNew)
}
