package main

import (
	"log"

	"github.com/shironxn/go-clean-arch/internal/config"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	// db, err := config.NewDatabase(cfg.Database).Connection()
	// if err != nil {
	// 	panic(err)
	// }

	server := config.NewServer(cfg.Server)
	log.Printf("Server is running on %s...\n", cfg.Server.Port)
	if err := server.Run(); err != nil {
		panic(err)
	}
}
