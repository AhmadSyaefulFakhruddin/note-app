package notes

import (
	"errors"
	"fmt"
	"time"
)

// ErrNoteNotFound is a sentinel error. Professionals define these at the top
// so the caller can check 'if err == notes.ErrNoteNotFound' without string matching.
var ErrNoteNotFound = errors.New("note not found")

type Repository interface {
	GetAllNotes() ([]Note, error)
	GetNoteByID(id string) (Note, error)
	StoreNewNote(newNote Note)
	UpdateNote(newNoteData Note, id string) (noteId string, err error)
	DeleteNote(id string) error
}

type repository struct {
	mockData []Note
}

func NewRepository() Repository {
	now := time.Now()

	return &repository{
		mockData: []Note{
			{
				ID:         "note-001",
				Title:      "Weekly Sync Meeting",
				Content:    "Discussed Q3 goals and marketing budget allocation.",
				Folder:     "Work",
				Tags:       []string{"meeting", "marketing", "q3"},
				IsPinned:   true,
				IsArchived: false,
				SyncStatus: "synced",
				CreatedAt:  now.Add(-48 * time.Hour),
				UpdatedAt:  now.Add(-1 * time.Hour),
			},
			{
				ID:         "note-002",
				Title:      "Grocery List",
				Content:    "Milk, Eggs, Bread, Coffee beans, Apples.",
				Folder:     "Personal",
				Tags:       []string{"shopping", "errands"},
				IsPinned:   false,
				IsArchived: false,
				SyncStatus: "pending",
				CreatedAt:  now.Add(-5 * time.Hour),
				UpdatedAt:  now,
			},
			{
				ID:         "note-003",
				Title:      "App Idea: Habit Tracker",
				Content:    "A gamified habit tracker that rewards users with virtual pets.",
				Folder:     "Ideas",
				Tags:       []string{"brainstorm", "dev"},
				IsPinned:   true,
				IsArchived: false,
				SyncStatus: "synced",
				CreatedAt:  now.Add(-720 * time.Hour), // roughly 1 month ago
				UpdatedAt:  now.Add(-24 * time.Hour),
			},
			{
				ID:         "note-004",
				Title:      "Tax Documents 2024",
				Content:    "Need to organize W2s and expense receipts for the accountant.",
				Folder:     "Finance",
				Tags:       []string{"taxes", "urgent"},
				IsPinned:   false,
				IsArchived: true,
				SyncStatus: "synced",
				CreatedAt:  now.Add(-8760 * time.Hour), // roughly 1 year ago
				UpdatedAt:  now.Add(-720 * time.Hour),
			},
			{
				ID:         "note-005",
				Title:      "Japan Trip Itinerary",
				Content:    "Days 1-3: Tokyo. Days 4-6: Kyoto. Day 7: Osaka.",
				Folder:     "Travel",
				Tags:       []string{"vacation", "japan"},
				IsPinned:   false,
				IsArchived: false,
				SyncStatus: "failed",
				CreatedAt:  now.Add(-2 * time.Hour),
				UpdatedAt:  now.Add(-2 * time.Hour),
			},
			{
				ID:         "note-006",
				Title:      "Books to Read",
				Content:    "1. Dune\n2. The Pragmatic Programmer\n3. Atomic Habits",
				Folder:     "Hobbies",
				Tags:       []string{"reading", "wishlist"},
				IsPinned:   false,
				IsArchived: false,
				SyncStatus: "synced",
				CreatedAt:  now.Add(-120 * time.Hour),
				UpdatedAt:  now.Add(-100 * time.Hour),
			},
			{
				ID:         "note-007",
				Title:      "Wi-Fi Passwords",
				Content:    "Home: supersecret123\nCafe: coffeetime",
				Folder:     "Important",
				Tags:       []string{"security", "passwords"},
				IsPinned:   true,
				IsArchived: false,
				SyncStatus: "synced",
				CreatedAt:  now.Add(-4000 * time.Hour),
				UpdatedAt:  now.Add(-4000 * time.Hour),
			},
		},
	}
}

func (r *repository) GetAllNotes() ([]Note, error) {
	return r.mockData, nil
}

func (r *repository) GetNoteByID(id string) (Note, error) {
	for _, note := range r.mockData {
		if note.ID == id {
			return note, nil
		}
	}
	return Note{}, fmt.Errorf("note with ID %s not found", id)
}

func (r *repository) StoreNewNote(newNote Note) {
	r.mockData = append(r.mockData, newNote)
}

func (r *repository) UpdateNote(updatedNote Note, id string) (string, error) {
	for index := range r.mockData {
		if r.mockData[index].ID == id {
			r.mockData[index] = updatedNote
			return id, nil
		}
	}

	return "", fmt.Errorf("%w: no note exists with ID %s", ErrNoteNotFound, id)
}

func (r *repository) DeleteNote(id string) error {
	for index := range r.mockData {
		if r.mockData[index].ID == id {
			r.mockData = append(r.mockData[:index], r.mockData[index+1:]...)
			return nil
		}
	}

	return fmt.Errorf("%w: no note exists with ID %s", ErrNoteNotFound, id)
}
