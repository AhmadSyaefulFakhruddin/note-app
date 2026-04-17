package tags

import (
	"time"

	"github.com/google/uuid"
)

type TagDb struct {
	ID        uuid.UUID
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	IsNew     *bool
}

type TagDto struct {
	ID   uuid.UUID `json:"id" binding:"required"`
	Name string    `json:"name" binding:"required"`
}

type CreateMultipleTagsDto struct {
	TagNames []string `json:"tagNames" binding:"required"`
}

type NewTagMinimalDto struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	IsNew *bool     `json:"isNew"`
}

func toNewTagMinimalDTO(db TagDb) NewTagMinimalDto {
	return NewTagMinimalDto{
		ID:    db.ID,
		Name:  db.Name,
		IsNew: db.IsNew,
	}
}

// 10/10: A pure function for a SLICE.
// Notice it doesn't belong to any struct!
func toNewTagMinimalDTOs(dbs []TagDb) []NewTagMinimalDto {
	// Your highly optimized memory allocation!
	res := make([]NewTagMinimalDto, len(dbs))
	for i, db := range dbs {
		res[i] = toNewTagMinimalDTO(db)
	}
	return res
}
