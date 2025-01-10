package server

import (
	"github.com/gofiber/fiber/v2"

	"go-view/internal/database"
)

type FiberServer struct {
	*fiber.App

	db database.Service
}

func New() *FiberServer {
	server := &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader: "go-view",
			AppName:      "go-view",
		}),

		db: database.New(),
	}

	return server
}
