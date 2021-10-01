package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

type DataFollowers struct {
	Userid    string `json:"userId"`
	Username  string `json:"username"`
	Followers int    `json:"followers"`
}

var followers = []DataFollowers{
	{Userid: "sammy", Username: "SammyShark", Followers: 987},
	{Userid: "jesse", Username: "JesseOctopus", Followers: 432},
	{Userid: "drew", Username: "DrewSquid", Followers: 321},
	{Userid: "jamie", Username: "JamieMantisShrimp", Followers: 654},
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Rahma 👋!")
	})

	app.Get("/follower", getDataFollower)
	app.Get("/follower/:username", getDataFollowerByUsername)
	app.Get("/:userId/detail", getDataFollowerById)

	app.Listen(":" + os.Getenv("PORT"))

}

func getDataFollower(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(followers)
}

func getDataFollowerByUsername(c *fiber.Ctx) error {
	username := c.Params("username")

	for _, follow := range followers {
		if follow.Username == username {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				//"success": true,
				username: fiber.Map{
					"followers": follow.Followers,
				},
			})
		}
	}

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Record not found"})
}

//GetDataFollowerById
func getDataFollowerById(c *fiber.Ctx) error {
	userId := c.Params("userId")

	for _, follow := range followers {
		if follow.Userid == userId {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				//"success": true,
				userId: follow,
			})
		}
	}

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Record not found"})
}
