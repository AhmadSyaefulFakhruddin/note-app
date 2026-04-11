package notes

import (
	"time"

	"github.com/google/uuid"
)

type ApiResponse[T any] struct {
	Status  string `json:"status"`
	Data    T      `json:"data"`
	Message string `json:"message"`
}

type Tag struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type Note struct {
	ID         uuid.UUID   `json:"id"`
	Title      string      `json:"title"`
	Content    string      `json:"content"`
	Folder     string      `json:"folder"`
	TagsIds    []uuid.UUID `json:"idTags"`
	IsPinned   bool        `json:"isPinned"`
	IsArchived bool        `json:"isArchived"`
	SyncStatus string      `json:"syncStatus"`
	CreatedAt  time.Time   `json:"createdAt"`
	UpdatedAt  time.Time   `json:"updatedAt"`
}

type NoteData struct {
	ID         uuid.UUID `json:"id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	Folder     string    `json:"folder"`
	Tags       []Tag     `json:"tags"`
	IsPinned   bool      `json:"isPinned"`
	IsArchived bool      `json:"isArchived"`
	SyncStatus string    `json:"syncStatus"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

type CreateNoteRequest struct {
	Title       string   `json:"title" binding:"required,min=1,notblank"`
	Content     string   `json:"content" binding:"omitempty,notblank"`
	Folder      string   `json:"folder" binding:"omitempty,notblank"`
	TagsIds     []string `json:"tagIds" binding:"omitempty,notblank"`
	NewTagNames []string `json:"newTagNames" binding:"omitempty,notblank"`
}

type UpdateNoteRequest struct {
	Title      *string `json:"title" binding:"omitempty,min=1,notblank"`
	Content    *string `json:"content" binding:"omitempty,notblank"`
	Folder     *string `json:"folder" binding:"omitempty,notblank"`
	Tags       *[]Tag  `json:"tags" binding:"omitempty,notblank"`
	IsPinned   *bool   `json:"isPinned" binding:"omitempty,notblank"`
	IsArchived *bool   `json:"isArchived" binding:"omitempty,notblank"`
}
