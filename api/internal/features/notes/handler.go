package notes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(s Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) GetNotes(c *gin.Context) {
	notes := h.service.FetchAllNotes()

	response := ApiResponse[[]NoteData]{
		Status:  "success",
		Data:    notes,
		Message: "success to get the notes",
	}

	c.JSON(http.StatusOK, response)
}

func (h *Handler) GetNoteDetail(c *gin.Context) {
	var noteId = c.Param("id")

	note, err := h.service.FindNote(noteId)

	if err != nil {
		response := ApiResponse[any]{
			Status:  "fail",
			Data:    nil,
			Message: fmt.Sprintf("The note id %s is not found", noteId),
		}

		c.JSON(http.StatusNotFound, response)
		return
	}

	response := ApiResponse[NoteData]{
		Status:  "success",
		Data:    note,
		Message: "success to get the note",
	}

	c.JSON(http.StatusOK, response)
}

func (h *Handler) CreateNote(c *gin.Context) {
	var noteData CreateNoteRequest

	if err := c.ShouldBindJSON(&noteData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Validation failed",
			"error":   err.Error(),
		})
		return
	}

	noteId, err := h.service.CreateNewNote(noteData)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "Failed to create note"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "success to create the note", "id": noteId})
}

func (h *Handler) UpdateNote(c *gin.Context) {
	var updateNoteData UpdateNoteRequest
	id := c.Param("id")

	if err := c.ShouldBindJSON(&updateNoteData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "Invalid request body"})
		return
	}

	noteId, err := h.service.UpdateNote(updateNoteData, id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "Failed to update note"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success to update the note", "id": noteId})
}

func (h *Handler) DeleteNote(c *gin.Context) {
	id := c.Param("id")

	err := h.service.DeleteNote(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error(), "message": "Failed to delete note"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success to delete the note", "id": id})
}
