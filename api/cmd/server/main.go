package main

import (
	"note-app-api/internal/config"
	"note-app-api/internal/features/notes"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func main() {
	r := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("notblank", config.NotBlank)
	}

	repository := notes.NewRepository()
	noteService := notes.NewService(repository)
	noteHandler := notes.NewHandler(noteService)

	notes.RegisterRoutes(r, noteHandler)

	r.Run()
}
