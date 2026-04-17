package notes

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine, h *Handler) {

	notesGroup := r.Group("/notes")

	{
		notesGroup.GET("", h.GetAllNotes)

		notesGroup.POST("", h.CreateNote)
	}

}
