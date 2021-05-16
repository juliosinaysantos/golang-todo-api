package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/juliosinaysantos/golang-todo-api/internal/database"
	"github.com/juliosinaysantos/golang-todo-api/pkg/utils"
)

func main() {
	dbDriver := "postgres"
	dbSource := utils.GetStringEnv("DATABASE_URL", "postgres://postgres@localhost:5432/go_todo_api?sslmode=disable")

	db, err := database.New(dbDriver, dbSource)
	if err != nil {
		log.Fatalf("Unable database connection, %v\n", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Unable database ping, %v\n", err)
	}

	defer db.Close()

	PORT := utils.GetStringEnv("PORT", "5050")
	app := fiber.New()

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"message": "Hello, world! 👋",
		})
	})

	log.Fatal(app.Listen(":" + PORT))
}
