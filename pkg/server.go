package pkg

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/thehaung/cloudflare-worker/configs"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"sync"
)

type HttpServer struct {
	once   sync.Once
	Server *fiber.App
	Router *fiber.Router
	DB     *mongo.Database
}

func NewHttpServer() {
	httpServer := &HttpServer{}

	httpServer.once.Do(func() {
		httpServer.InitEnv()
	})
}

func (s *HttpServer) InitEnv() {
	env := configs.GetEnvironment()

	if err := godotenv.Load(fmt.Sprintf("deployemnts/%s/.env", env)); err != nil {
		log.Fatalf("Error while loading Env: %v", err)
	}
}
