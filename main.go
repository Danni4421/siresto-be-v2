package main

import (
  "github.com/gofiber/fiber/v2"
  "github.com/gofiber/fiber/v2/middleware/cors"
  "github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
  f := fiber.New()
  f.Use(logger.New())
  f.Use(cors.New(cors.Config{
    AllowOrigins: "*",
    AllowMethods: "GET,POST,PUT,DELETE,OPTIONS,PATCH",
    AllowHeaders: "Origin, Content-Type, Accept, Authorization",
  }))

  f.Get("/", func(c *fiber.Ctx) error {
    return c.SendString("Hello, World!")
  })

  err := f.Listen(":5000")
  
  if err != nil {
    panic(err)
  }
}
