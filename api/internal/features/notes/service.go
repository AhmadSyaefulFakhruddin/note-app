package notes

import (
	"context"
	"note-app-api/internal/features/apperr"
	"note-app-api/internal/features/folders"
)

type Service interface {
	GetAllNotes(ctx context.Context) ([]NoteDb, error)
	CreateNote(ctx context.Context, newNoteData CreateNoteRequest) (NoteDb, error)
}

type service struct {
	repo          Repository
	folderService folders.Service
}

func NewService(r Repository, folderService folders.Service) Service {
	return &service{
		repo:          r,
		folderService: folderService,
	}
}

func (s *service) GetAllNotes(ctx context.Context) ([]NoteDb, error) {
	return s.repo.GetAllNotes(ctx)
}

func (s *service) CreateNote(ctx context.Context, newNoteData CreateNoteRequest) (NoteDb, error) {

	note, err := s.repo.StoreNewNote(ctx, NoteDb{
		Title:    newNoteData.Title,
		Content:  newNoteData.Content,
		FolderId: newNoteData.FolderId,
	})

	if err != nil {
		return NoteDb{}, apperr.NewInternal(err)
	}

	return note, nil
}
