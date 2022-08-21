package dns

import (
	"github.com/gofiber/fiber/v2"
)

type IService interface {
	Fetch(ctx *fiber.Ctx) (err error)
}
