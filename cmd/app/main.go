package main

import (
	"fmt"
	"log"

	"fairseller-backend/config"
	"fairseller-backend/internal/app"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	fmt.Println(cfg)
	// Run
	app.Run(cfg)
}
