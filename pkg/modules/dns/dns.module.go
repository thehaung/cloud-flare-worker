package dns

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitModule(fiber *fiber.App, db *mongo.Database) {
	dnsService := NewService(db)

	NewDNSController(fiber, dnsService)
}
