package notes

import "time"

type Note struct {
	ID         string    `json:"id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	Folder     string    `json:"folder"`
	Tags       []string  `json:"tags"`
	IsPinned   bool      `json:"isPinned"`
	IsArchived bool      `json:"isArchived"`
	SyncStatus string    `json:"syncStatus"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

type CreateNoteRequest struct {
	Title   string   `json:"title" binding:"required,min=1,notblank"`
	Content string   `json:"content" binding:"omitempty,notblank"`
	Folder  string   `json:"folder" binding:"omitempty,notblank"`
	Tags    []string `json:"tags" binding:"omitempty,notblank"`
}

type UpdateNoteRequest struct {
	Title      *string   `json:"title" binding:"omitempty,min=1,notblank"`
	Content    *string   `json:"content" binding:"omitempty,notblank"`
	Folder     *string   `json:"folder" binding:"omitempty,notblank"`
	Tags       *[]string `json:"tags" binding:"omitempty,notblank"`
	IsPinned   *bool     `json:"isPinned" binding:"omitempty,notblank"`
	IsArchived *bool     `json:"isArchived" binding:"omitempty,notblank"`
	SyncStatus *string   `json:"syncStatus" binding:"omitempty,notblank"`
}
