package notes

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Service interface {
	FetchAllNotes() ([]Note, error)
	FindNote(id string) (Note, error)
	CreateNewNote(noteData CreateNoteRequest) (string, error)
	UpdateNote(newNoteData UpdateNoteRequest, id string) (noteId string, err error)
	DeleteNote(id string) error
}

type NoteUpdater func(n *Note)

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) FetchAllNotes() ([]Note, error) {
	return s.repo.GetAllNotes()
}

func (s *service) FindNote(id string) (Note, error) {
	return s.repo.GetNoteByID(id)
}

func (s *service) CreateNewNote(noteData CreateNoteRequest) (string, error) {
	now := time.Now()

	noteId := uuid.New()

	s.repo.StoreNewNote(Note{
		ID:         noteId.String(),
		IsPinned:   false,
		IsArchived: false,
		SyncStatus: "synced",
		CreatedAt:  now,
		UpdatedAt:  now,
		Title:      noteData.Title,
		Content:    noteData.Content,
		Folder:     noteData.Folder,
		Tags:       noteData.Tags,
	})

	return noteId.String(), nil
}

func (s *service) UpdateNote(newNoteData UpdateNoteRequest, id string) (string, error) {
	note, err := s.repo.GetNoteByID(id)

	if err != nil {
		return "", err
	}

	var updates []NoteUpdater

	if newNoteData.Title != nil {
		updates = append(updates, updateTitle(newNoteData.Title))
	}

	if newNoteData.Content != nil {
		updates = append(updates, updateContent(newNoteData.Content))
	}

	if newNoteData.Folder != nil {
		updates = append(updates, updateFolder(newNoteData.Folder))
	}

	if newNoteData.Tags != nil {
		updates = append(updates, updateTags(newNoteData.Tags))
	}

	if newNoteData.IsPinned != nil {
		updates = append(updates, updateIsPinned(newNoteData.IsPinned))
	}

	if newNoteData.IsArchived != nil {
		updates = append(updates, updateIsArchived(newNoteData.IsArchived))
	}

	if newNoteData.SyncStatus != nil {
		updates = append(updates, updateSyncStatus(newNoteData.SyncStatus))
	}

	applyUpdate(&note, updates...)

	if len(updates) == 0 {
		return "", fmt.Errorf("at least one field must be provided for update")
	}

	note.UpdatedAt = time.Now()

	noteId, err := s.repo.UpdateNote(note, id)

	if err != nil {
		return "", err
	}

	return noteId, nil
}

func (s *service) DeleteNote(id string) error {
	return s.repo.DeleteNote(id)
}

// helper
func updateTitle(title *string) NoteUpdater {
	return func(n *Note) {
		if title != nil {
			n.Title = *title
		}
	}
}

func updateContent(content *string) NoteUpdater {
	return func(n *Note) {
		if content != nil {
			n.Content = *content
		}
	}
}

func updateFolder(folder *string) NoteUpdater {
	return func(n *Note) {
		if folder != nil {
			n.Folder = *folder
		}
	}
}

func updateTags(tags *[]string) NoteUpdater {
	return func(n *Note) {
		if tags != nil {
			n.Tags = *tags
		}
	}
}
func updateIsPinned(isPinned *bool) NoteUpdater {
	return func(n *Note) {
		if isPinned != nil {
			n.IsPinned = *isPinned
		}
	}
}

func updateIsArchived(isArchived *bool) NoteUpdater {
	return func(n *Note) {
		if isArchived != nil {
			n.IsArchived = *isArchived
		}
	}
}

func updateSyncStatus(syncStatus *string) NoteUpdater {
	return func(n *Note) {
		if syncStatus != nil {
			n.SyncStatus = *syncStatus
		}
	}
}

func applyUpdate(note *Note, updates ...NoteUpdater) {
	for _, update := range updates {
		update(note)
	}
}
