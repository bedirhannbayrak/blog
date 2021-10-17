package main

import (
	"log"
	"os"

	"github.com/bedirhannbayrak/blog/controllers"
	"github.com/Kamva/mgm/v2"
	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	connectionString := os.Getenv("MONGODB_CONNECTION_STRING")
	if len(connectionString) == 0 {
		connectionString = "mongodb://localhost:27017"
	}

	err := mgm.SetDefaultConfig(nil, "blog", options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	app := fiber.New()

	app.Get("/api/posts", controllers.GetAllPosts)
	app.Get("/api/posts/:id", controllers.GetPostByID)
	app.Post("/api/posts", controllers.CreatePost)

	app.Listen(3000)
}
