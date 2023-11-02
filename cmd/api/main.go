package main

import (
	"log"

	"github.com/koha90/podkrepizza-api-v1/config"
	"github.com/koha90/podkrepizza-api-v1/internal/app"
)

func main() {
	cfg := config.MustConfig()

	err := app.Run(cfg)
	if err != nil {
		log.Fatal(err)
	}
}
