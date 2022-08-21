package dns

import "github.com/gofiber/fiber/v2"

type Controller struct {
	r       *fiber.App
	service IService
}

func NewDNSController(r *fiber.App, service IService) {
	controller := Controller{
		r:       r,
		service: service,
	}
	controller.InitController()
}

func (c *Controller) InitController() {
	dnsRoute := c.r.Group("dns")

	dnsRoute.Get("/", c.service.Fetch)
}
