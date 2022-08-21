package dns

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Service struct {
	db *mongo.Database
}

func NewService(db *mongo.Database) IService {
	return &Service{
		db: db,
	}
}

func (s Service) Fetch(ctx *fiber.Ctx) (err error) {
	cursor, err := s.db.Collection("dns").Find(ctx.Context(), bson.D{})
	err = ctx.Status(fiber.StatusOK).JSON(cursor)
	if err != nil {
		return err
	}
	return
}
