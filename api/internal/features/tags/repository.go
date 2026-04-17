package tags

import (
	"context"
	"fmt"
	"note-app-api/internal/database"
	"note-app-api/internal/features/apperr"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	GetTags(c context.Context) ([]TagDb, error)
	CreateMultipleTags(c context.Context, tagNames []string) ([]TagDb, error)
}

type repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetTags(c context.Context) ([]TagDb, error) {
	query := `
		SELECT id, name, created_at, updated_at
		FROM tags
	`

	rows, err := r.db.Query(c, query)

	if err != nil {
		return []TagDb{}, database.HandlePostgresError(err)
	}

	defer rows.Close()

	tagDbs := make([]TagDb, 0)
	for rows.Next() {
		var tagDb TagDb

		if err := rows.Scan(
			&tagDb.ID,
			&tagDb.Name,
			&tagDb.CreatedAt,
			&tagDb.UpdatedAt,
		); err != nil {
			return []TagDb{}, database.HandlePostgresError(err)
		}

		tagDbs = append(tagDbs, tagDb)
	}

	err = rows.Err()

	if err != nil {
		return []TagDb{}, database.HandlePostgresError(err)
	}

	return tagDbs, nil
}

func (r *repository) CreateMultipleTags(c context.Context, tagNames []string) ([]TagDb, error) {
	if len(tagNames) == 0 {
		return []TagDb{}, &apperr.AppError{
			Code:    400,
			Message: "Please input min 1 tag name",
		}
	}

	query := `
		WITH new_tags AS (
			INSERT INTO tags (name)
			SELECT * FROM UNNEST($1::text[])
			ON CONFLICT (name) DO NOTHING
			RETURNING id, name, created_at, updated_at
		)

		SELECT id, name, created_at, updated_at, true as is_new 
		FROM new_tags

		UNION ALL

		SELECT id, name, created_at, updated_at, false as is_new 
		FROM tags
		WHERE name = ANY($1::text[]);
	`

	rows, err := r.db.Query(c, query, tagNames)

	if err != nil {
		return nil, database.HandlePostgresError(err)
	}

	defer rows.Close()

	tagDbs := make([]TagDb, 0, len(tagNames))
	for rows.Next() {
		var tagDb TagDb

		err := rows.Scan(
			&tagDb.ID,
			&tagDb.Name,
			&tagDb.CreatedAt,
			&tagDb.UpdatedAt,
			&tagDb.IsNew,
		)

		if err != nil {
			return nil, database.HandlePostgresError(err)
		}

		tagDbs = append(tagDbs, tagDb)
	}

	if err := rows.Err(); err != nil {
		alreadyExistTags := make([]string, 0)

		for _, tagDb := range tagDbs {
			if tagDb.IsNew != nil && *tagDb.IsNew == false {
				alreadyExistTags = append(alreadyExistTags, tagDb.Name)
			}
		}

		return []TagDb{}, database.HandlePostgresError(err, fmt.Sprintf(`The tag names "%s" already exist`, strings.Join(alreadyExistTags, ", ")))
	}

	return tagDbs, nil
}
