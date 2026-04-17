package folders

import (
	"time"

	"github.com/google/uuid"
)

type FolderDb struct {
	ID        uuid.UUID `db:"id"`
	Name      string    `db:"name"`
	createdAt time.Time `db:"created_at"`
	updatedAt time.Time `db:"updated_at"`
}

type CreateFolderRequest struct {
	FolderName string `json:"folderName" binding:"required,notblank,validfolder,min=1"`
}

type FolderResponse struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
