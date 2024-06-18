package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

type PostID int

type Post struct {
	ID    PostID `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

var posts = map[PostID]Post{
	1: {
		ID:    1,
		Title: "ลองเล่น Fiber",
		Body:  "Fiber ใช่ง่ายๆเลยสำหรับคนเคยเล่น Express มา",
	},
	2: {
		ID:    2,
		Title: "สิ่งที่อยากทำในปี 2021",
		Body:  "อยากลดน้ำหนักให้น้อยกว่านี้",
	},
}

func main() {
	app := fiber.New()

	app.Get("/posts", func(c *fiber.Ctx) error {
		var result []Post
		for _, post := range posts {
			result = append(result, post)
		}
		return c.JSON(result)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Fatalln(app.Listen(fmt.Sprintf(":%v", port)))
}
