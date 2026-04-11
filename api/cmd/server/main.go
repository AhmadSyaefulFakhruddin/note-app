package main

import (
	"context"
	"log"
	"note-app-api/internal/config"
	"note-app-api/internal/database"
	"note-app-api/internal/features/notes"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func main() {
	r := gin.Default()

	// ask again
	ctx := context.Background()

	dbURL := "postgres://postgres:123qweas@localhost:5432/note_db?sslmode=disable"

	dbPool, err := database.InitPostgres(ctx, dbURL)
	if err != nil {
		log.Fatalf("Critical Error: %s", err.Error())
	}

	defer dbPool.Close()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("notblank", config.NotBlank)
	}

	repository := notes.NewRepository()
	noteService := notes.NewService(repository)
	noteHandler := notes.NewHandler(noteService)

	notes.RegisterRoutes(r, noteHandler)

	r.Run()
}
