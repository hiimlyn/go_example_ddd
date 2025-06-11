package main

import (
	"log"

	"github.com/hiimlyn/go-api/internal/app"
	"github.com/hiimlyn/go-api/internal/config"
)

func main() {
	cfg := config.New()
	app := app.New(cfg)
	log.Fatal(app.Listen(":" + cfg.Port))
}
