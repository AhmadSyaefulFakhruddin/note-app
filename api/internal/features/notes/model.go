package notes

import (
	"note-app-api/internal/features/folders"
	"note-app-api/internal/features/tags"
	"time"

	"github.com/google/uuid"
)

type NoteDb struct {
	ID         uuid.UUID  `db:"id"`
	Title      string     `db:"title"`
	Content    string     `db:"content"`
	FolderId   *uuid.UUID `db:"folder_id"`
	IsPinned   bool       `db:"is_pinned"`
	IsArchived bool       `db:"is_archived"`
	SyncStatus string     `db:"sync_status"`
	CreatedAt  time.Time  `db:"created_at"`
	UpdatedAt  time.Time  `db:"updated_at"`
}

type NoteDto struct {
	ID         uuid.UUID              `json:"id" binding:"required"`
	Title      string                 `json:"title" binding:"required"`
	Content    string                 `json:"content" binding:"required"`
	Folder     folders.FolderResponse `json:"folder" binding:"required"`
	Tags       []tags.TagDto          `json:"tags" binding:"required"`
	IsPinned   bool                   `json:"isPinned" binding:"required"`
	IsArchived bool                   `json:"isArchived" binding:"required"`
	SyncStatus string                 `json:"syncStatus" binding:"required"`
	CreatedAt  time.Time              `json:"createdAt" binding:"required"`
	UpdatedAt  time.Time              `json:"updatedAt" binding:"required"`
}

type CreateNoteRequest struct {
	Title    string     `json:"title" binding:"required,min=1,notblank"`
	Content  string     `json:"content" binding:"omitempty,notblank"`
	FolderId *uuid.UUID `json:"folderId" binding:"omitempty,notblank"`
}

type UpdateNoteRequest struct {
	Title      *string       `json:"title" binding:"omitempty,min=1,notblank"`
	Content    *string       `json:"content" binding:"omitempty,notblank"`
	Folder     *string       `json:"folder" binding:"omitempty,notblank"`
	Tags       *[]tags.TagDb `json:"tags" binding:"omitempty,notblank"`
	IsPinned   *bool         `json:"isPinned" binding:"omitempty,notblank"`
	IsArchived *bool         `json:"isArchived" binding:"omitempty,notblank"`
}
