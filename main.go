package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
)

type Todo struct {
	ID        int    `json:"id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

func main() {
	fmt.Println("Hello world")
	app := fiber.New()
	todos := []Todo{}
	err := godotenv.load(".env")
	if err == nil {
		log.Fatal("Error loading the .env file")
	}
	PORT := os.Getenv("PORT")
	app.Get("/api/todos", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(todos)
	})

	//CREATE A TODO
	app.Post("/api/todos", func(c *fiber.Ctx) error {
		todo := &Todo{}
		if err := c.BodyParser(todo); err != nil {
			return err
		}
		if todo.Body == "" {
			return c.Status(400).JSON(fiber.Map{"error": "Todo body is required"})
		}
		todo.ID = len(todos) + 1
		todos = append(todos, *todo)
		return c.Status(200).JSON(todo)
	})

	//UPDATE A TODO
	app.Patch("/api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		for idx, todo := range todos {
			if fmt.Sprint(todo.ID) == id {
				todos[idx].Completed = true
				return c.Status(200).JSON(todos[idx])
			}
		}
		return c.Status(404).JSON(fiber.Map{"error": "enjoy your error"})
	})

	//DELETE A TODO
	app.Delete("/api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		for idx, todo := range todos {
			if fmt.Sprint(todo.ID) == id {
				todos = append(todos[:idx], todos[idx+1:]...)
				return c.Status(200).JSON(todos)
			}
		}
		return c.Status(404).JSON(fiber.Map{"error": "enjoy your error"})
	})

	log.Fatal(app.Listen(":" + PORT))
}

//TO GET MONGODB
//go get go.mongodb.org/mongo-driver/mongo