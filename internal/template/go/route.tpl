package route

import(
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/logger"
)

func RouterInit()*fiber.App{
    app := fiber.New()
    app.Use(logger.New())


    return app
}