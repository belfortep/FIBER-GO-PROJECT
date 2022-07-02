package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/google/uuid"
)

//structs y punteros!!!
type User struct {
	Id        string
	FirstName string
	LastName  string
}

func handleUser(c *fiber.Ctx) error {

	user := User{
		FirstName: "Some name",
		LastName:  "Some last name",
	}

	return c.Status(fiber.StatusOK).JSON(user)

}

func handleCreateUser(c *fiber.Ctx) error {

	user := User{}

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	user.Id = uuid.NewString()

	return c.Status(fiber.StatusOK).JSON(user)

}

func main() {

	app := fiber.New()

	//middlewares

	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello")
	})

	//agrupar las rutas
	userGroup := app.Group("/user")

	//muy similar a express en node.js
	userGroup.Get("", handleUser)

	userGroup.Post("", handleCreateUser)

	app.Listen(":3000")
}
