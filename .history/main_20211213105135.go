package main

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
	jwtware "github.com/gofiber/jwt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	log "github.com/sirupsen/logrus"
	
)

var (
	DBConn *gorm.DB
)

const jwtSecret = "asecret"

func authRequired() func(ctx *fiber.Ctx) {
	return jwtware.New(jwtware.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) {
			ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		},
		SigningKey: []byte(jwtSecret),
	})
}

func login(ctx *fiber.Ctx) {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var body request
	err := ctx.BodyParser(&body)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse json",
		})
		return
	}

	if body.Email != "admin@admin.com" || body.Password != "123" {
		ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Bad Credentials",
		})
		return
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = "1"
	claims["exp"] = time.Now().Add(time.Hour * 24 * 7) // a week

	s, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		ctx.SendStatus(fiber.StatusInternalServerError)
		return
	}

	ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": s,
		"user": struct {
			Id    int    `json:"id"`
			Email string `json:"email"`
		}{
			Id:    1,
			Email: "admin@admin.hu",
		},
	})
}
func initDatabase() {
	var err error
	DBConn, err = gorm.Open("sqlite3", "blog.db")
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")
	//var blog Blog
	//var comments Commenst
	//DBConn.AutoMigrate(&blog, &comments)

}

func setupRoutes(app *fiber.App) {
	//app.Get("/", getAllBlog)

	//app.Get("/new", authRequired(), newBlog)
	//app.Get("/comment", authRequired(), newComments)
	app.Post("/login", login)

}

func main() {
	app := fiber.New()
	app.Use(middleware.Logger())
	initDatabase()

	setupRoutes(app)
	app.Listen(3000)
	defer DBConn.Close()
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
