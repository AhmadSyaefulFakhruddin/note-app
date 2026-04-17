package main

import (
	"context"
	"log"
	"note-app-api/internal/config"
	"note-app-api/internal/database"
	"note-app-api/internal/features/folders"
	"note-app-api/internal/features/notes"
	"note-app-api/internal/features/tags"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func main() {
	r := gin.Default()

	ctx := context.Background()

	dbURL := "postgres://postgres:123qweas@localhost:5432/note_db?sslmode=disable"

	dbPool, err := database.InitPostgres(ctx, dbURL)
	if err != nil {
		log.Fatalf("Critical Error: %s", err.Error())
	}

	defer dbPool.Close()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("notblank", config.NotBlank)
		v.RegisterValidation("validfolder", config.ValidFolder)
	}

	foldersRepository := folders.NewRepository(dbPool)
	foldersService := folders.NewService(foldersRepository)
	foldersHandler := folders.NewHandler(foldersService)
	folders.RegisterRoutes(r, foldersHandler)

	tagsRepository := tags.NewRepository(dbPool)
	tagsService := tags.NewService(tagsRepository)
	tagsHandler := tags.NewHandler(tagsService)
	tags.RegisterRoute(r, tagsHandler)

	notesRepository := notes.NewRepository(dbPool)
	notesService := notes.NewService(notesRepository, foldersService)
	notesHandler := notes.NewHandler(notesService)
	notes.RegisterRoutes(r, notesHandler)

	r.Run()
}
