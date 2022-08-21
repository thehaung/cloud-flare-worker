package pkg

import (
	"fmt"
	"github.com/gofiber/helmet/v2"
	"github.com/thehaung/cloudflare-worker/config"
	"github.com/thehaung/cloudflare-worker/pkg/utils"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func setupMiddlewares(app *fiber.App) {
	app.Use(helmet.New())
	app.Use(recover.New())
	app.Use(cors.New())
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed, // 1
	}))
	app.Use(etag.New())
	if os.Getenv("ENABLE_LIMITER") != "" {
		app.Use(limiter.New())
	}
	if os.Getenv("ENABLE_LOGGER") != "" {
		app.Use(logger.New())
	}
}

func Create() *fiber.App {
	app := fiber.New(fiber.Config{
		// Override default error handler
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			if e, ok := err.(*utils.Error); ok {
				return ctx.Status(e.Status).JSON(e)
			} else if e, ok := err.(*fiber.Error); ok {
				return ctx.Status(e.Code).JSON(utils.Error{Status: e.Code, Code: "internal-server", Message: e.Message})
			} else {
				return ctx.Status(500).JSON(utils.Error{Status: 500, Code: "internal-server", Message: err.Error()})
			}
		},
	})

	setupMiddlewares(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	return app
}

func Listen(app *fiber.App) error {

	// 404 Handler
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	serverPort := config.GetServerPort()
	return app.Listen(fmt.Sprintf(":%s", serverPort))
}