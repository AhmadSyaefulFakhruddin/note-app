package notes

import (
	"net/http"
	"note-app-api/internal/features/response"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(s Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) GetAllNotes(c *gin.Context) {
	notes, err := h.service.GetAllNotes(c.Request.Context())
	if err != nil {
		errorCode, errorRes := response.Error(err)

		c.JSON(errorCode, errorRes)
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse(notes, "Success get the notes"))
}

func (h *Handler) CreateNote(c *gin.Context) {
	var newNoteData CreateNoteRequest

	if err := c.ShouldBindBodyWithJSON(&newNoteData); err != nil {
		errorCode, errorRes := response.Error(err)

		c.JSON(errorCode, errorRes)
		return
	}

	newNote, err := h.service.CreateNote(c.Request.Context(), newNoteData)

	if err != nil {
		errorCode, errorRes := response.Error(err)

		c.JSON(errorCode, errorRes)
		return
	}

	c.JSON(http.StatusCreated, response.SuccessResponse(newNote, "Successfully created new note"))
}
