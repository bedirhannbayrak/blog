package controllers

import (
	"github.com/Kamva/mgm/v2"
	"github.com/bedirhannbayrak/blog/models"
	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllPosts(ctx *fiber.Ctx) {
	collection := mgm.Coll(&models.Post{})
	posts := []models.Post{}

	err := collection.SimpleFind(&posts, bson.D{})
	if err != nil {
		ctx.Status(500).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(fiber.Map{
		"ok":    true,
		"posts": posts,
	})
	return
}

func GetPostByID(ctx *fiber.Ctx) {
	id := ctx.Params("id")

	post := &models.Post{}
	collection := mgm.Coll(post)

	err := collection.FindByID(id, post)
	if err != nil {
		ctx.Status(404).JSON(fiber.Map{
			"ok":    false,
			"error": "Post not found.",
		})
	}

	ctx.JSON(fiber.Map{
		"ok":   true,
		"post": post,
	})
}

func CreatePost(ctx *fiber.Ctx) {
	params := new(struct {
		Title   string
		Content string
	})

	ctx.BodyParser(&params)

	post := models.CreatePost(params.Title, params.Content)
	err := mgm.Coll(post).Create(post)
	if err != nil {
		ctx.Status(500).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(fiber.Map{
		"ok":   true,
		"post": post,
	})
}
