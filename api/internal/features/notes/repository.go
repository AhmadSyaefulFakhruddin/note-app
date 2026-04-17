package notes

import (
	"context"
	"errors"
	"note-app-api/internal/database"

	"github.com/jackc/pgx/v5/pgxpool"
)

// ErrNoteNotFound is a sentinel error. Professionals define these at the top
// so the caller can check 'if err == notes.ErrNoteNotFound' without string matching.
var ErrNoteNotFound = errors.New("note not found")

type Repository interface {
	// note
	GetAllNotes(ctx context.Context) ([]NoteDb, error)
	// GetNoteByID(id uuid.UUID) (NoteDb, error)
	StoreNewNote(ctx context.Context, newNote NoteDb) (NoteDb, error)
	// UpdateNote(newNoteData NoteDb, id uuid.UUID) (uuid.UUID, error)
	// DeleteNote(id uuid.UUID) error
}

type repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAllNotes(ctx context.Context) ([]NoteDb, error) {
	query := `
		SELECT id, title, content, folder_id, is_pinned, is_archived, sync_status, created_at, updated_at 
		FROM notes
		ORDER BY created_at DESC
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, database.HandlePostgresError(err)
	}
	defer rows.Close()

	notes := make([]NoteDb, 0)
	for rows.Next() {
		var note NoteDb
		if err := rows.Scan(&note.ID, &note.Title, &note.Content, &note.FolderId, &note.IsPinned, &note.IsArchived, &note.SyncStatus, &note.CreatedAt, &note.UpdatedAt); err != nil {
			return nil, database.HandlePostgresError(err)
		}
		notes = append(notes, note)
	}

	if err := rows.Err(); err != nil {
		return nil, database.HandlePostgresError(err)
	}

	return notes, nil
}

func (r *repository) StoreNewNote(ctx context.Context, newNote NoteDb) (NoteDb, error) {
	query := `
		INSERT INTO notes (title, content, folder_id)
		VALUES ($1, $2, $3)
		RETURNING id, is_pinned, is_archived, sync_status, created_at, updated_at
	`

	err := r.db.QueryRow(ctx, query, newNote.Title, newNote.Content, newNote.FolderId).Scan(
		&newNote.ID,
		&newNote.IsPinned,
		&newNote.IsArchived,
		&newNote.SyncStatus,
		&newNote.CreatedAt,
		&newNote.UpdatedAt,
	)

	if err != nil {
		return NoteDb{}, database.HandlePostgresError(err, "The note is already exist")
	}

	return newNote, nil
}
