package main

import (
	"log"

	"github.com/shironxn/go-clean-arch/internal/config"
	"github.com/shironxn/go-clean-arch/internal/config/db"
	"github.com/shironxn/go-clean-arch/internal/handler"
	"github.com/shironxn/go-clean-arch/internal/repository"
	"github.com/shironxn/go-clean-arch/internal/service"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	db, err := db.NewDatabase(cfg.Database).Connection()
	if err != nil {
		panic(err)
	}

	authorRepository := repository.NewAuthorRepository(db)
	authorService := service.NewAuthorService(authorRepository)
	authorHandler := handler.NewAuthorHandler(authorService)

	server := config.NewServer(cfg.Server, config.Handler{Author: authorHandler})
	log.Printf("Server is running on %s...\n", cfg.Server.Port)
	if err := server.Run(); err != nil {
		panic(err)
	}
}
