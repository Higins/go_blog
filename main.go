package main

import (
	"os"

	"web/db"
	"web/login"

	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	log "github.com/sirupsen/logrus"
)

func setupRoutes(app *fiber.App) {

	app.Get("/", blogController.Getallblog)

	app.Get("/new", login.AuthRequired(), blogController.NewBlog)
	app.Get("/:blogid/comment", login.AuthRequired() commentController.NewComment)
	app.Post("/login", login.Login)

}
func main() {
	app := fiber.New()
	app.Use(middleware.Logger())
	db.InitDatabase()
	setupRoutes(app)

	app.Listen(3000)
	defer db.DBConn.Close()
}
func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.WarnLevel)
}
