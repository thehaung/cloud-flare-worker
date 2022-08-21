package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/helmet/v2"
	"github.com/thehaung/cloudflare-worker/configs"
	"time"
)

func FiberMiddleware(app *fiber.App) {
	app.Use(helmet.New())
	app.Use(recover.New())
	app.Use(cors.New())
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed, // 1
	}))
	app.Use(etag.New())

	// Config this for timeout handler global
	app.Server().WriteTimeout = time.Duration(configs.GetTimeDeadlineAPI())
	if configs.IsEnableLimiter() {
		app.Use(limiter.New())
	}

	if configs.IsEnableLogger() {
		app.Use(logger.New())
	}
}
