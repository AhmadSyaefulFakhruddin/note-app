package folders

import (
	"context"
	"fmt"
	"note-app-api/internal/database"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	GetFolders(ctx context.Context) ([]FolderDb, error)
	CreateFolder(ctx context.Context, newFolder FolderDb) (FolderDb, error)
}

type repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetFolders(ctx context.Context) ([]FolderDb, error) {
	query := `
		SELECT id, name, created_at, updated_at 
		FROM folders
	`

	rows, err := r.db.Query(ctx, query)

	if err != nil {
		return nil, database.HandlePostgresError(err)
	}
	defer rows.Close()

	folders := make([]FolderDb, 0)

	for rows.Next() {
		var folder FolderDb

		if err := rows.Scan(
			&folder.ID,
			&folder.Name,
			&folder.createdAt,
			&folder.updatedAt,
		); err != nil {
			return nil, database.HandlePostgresError(err)
		}

		folders = append(folders, folder)
	}

	if err := rows.Err(); err != nil {
		return nil, database.HandlePostgresError(err)
	}

	return folders, nil
}

func (r *repository) CreateFolder(ctx context.Context, newFolder FolderDb) (FolderDb, error) {
	query := `
		INSERT INTO folders (name)
		VALUES ($1)
		RETURNING id, created_at, updated_at
	`

	err := r.db.QueryRow(ctx, query, newFolder.Name).Scan(
		&newFolder.ID,
		&newFolder.createdAt,
		&newFolder.updatedAt,
	)

	if err != nil {
		return FolderDb{}, database.HandlePostgresError(err, fmt.Sprintf(`The folder name "%s" already taken`, newFolder.Name))
	}

	return newFolder, nil
}
