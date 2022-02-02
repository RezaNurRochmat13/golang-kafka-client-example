package main

import (
	"golang-kafka-clients/config"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// config.CreateTopicKafka()
	// config.PublishMessageFromKafka()
	// config.ConsumeMessageFromKafka()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Get("/consume", func(c *fiber.Ctx) error {
		config.ConsumeMessageFromKafka()
		return c.SendString("Consume message")
	})

	app.Listen(":3000")
}
