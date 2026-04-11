package notes

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// ErrNoteNotFound is a sentinel error. Professionals define these at the top
// so the caller can check 'if err == notes.ErrNoteNotFound' without string matching.
var ErrNoteNotFound = errors.New("note not found")

type Repository interface {
	// note
	GetAllNotes() []Note
	GetNoteByID(id uuid.UUID) (Note, error)
	StoreNewNote(newNote Note)
	UpdateNote(newNoteData Note, id uuid.UUID) (noteId uuid.UUID, err error)
	DeleteNote(id uuid.UUID) error

	// tag
	GetAllTags() []Tag
	CreateNewTag(newTag []Tag)
	GetTagsByIds(tagIds []uuid.UUID) (tags []Tag, unknownTagIDs []string)
}

type repository struct {
	mockNote []Note
	mocktag  []Tag
}

func GetMockData() ([]Tag, []Note) {
	// 4 Mock Tags
	tags := []Tag{
		{ID: uuid.MustParse("11111111-1111-1111-1111-111111111111"), Name: "Work"},
		{ID: uuid.MustParse("22222222-2222-2222-2222-222222222222"), Name: "Personal"},
		{ID: uuid.MustParse("33333333-3333-3333-3333-333333333333"), Name: "Ideas"},
		{ID: uuid.MustParse("44444444-4444-4444-4444-444444444444"), Name: "Urgent"},
	}

	// Helper to simulate different times
	now := time.Now()

	// 10 Mock Notes
	notes := []Note{
		{
			ID:         uuid.MustParse("a1a1a1a1-a1a1-a1a1-a1a1-a1a1a1a1a1a1"),
			Title:      "Q3 Project Roadmap",
			Content:    "Drafting the milestones for the upcoming quarter.",
			Folder:     "Projects/baru",
			TagsIds:    []uuid.UUID{tags[0].ID, tags[3].ID}, // Work, Urgent
			IsPinned:   true,
			IsArchived: false,
			SyncStatus: "synced",
			CreatedAt:  now.Add(-72 * time.Hour),
			UpdatedAt:  now.Add(-24 * time.Hour),
		},
		{
			ID:         uuid.MustParse("b2b2b2b2-b2b2-b2b2-b2b2-b2b2b2b2b2b2"),
			Title:      "Grocery List",
			Content:    "Milk, eggs, bread, coffee beans, apples.",
			Folder:     "Home",
			TagsIds:    []uuid.UUID{tags[1].ID}, // Personal
			IsPinned:   false,
			IsArchived: false,
			SyncStatus: "syncing",
			CreatedAt:  now.Add(-2 * time.Hour),
			UpdatedAt:  now.Add(-1 * time.Hour),
		},
		{
			ID:         uuid.MustParse("c3c3c3c3-c3c3-c3c3-c3c3-c3c3c3c3c3c3"),
			Title:      "Startup App Concept",
			Content:    "An app that tracks indoor plant watering schedules.",
			Folder:     "Brainstorming",
			TagsIds:    []uuid.UUID{tags[2].ID}, // Ideas
			IsPinned:   false,
			IsArchived: false,
			SyncStatus: "synced",
			CreatedAt:  now.Add(-120 * time.Hour),
			UpdatedAt:  now.Add(-120 * time.Hour),
		},
		{
			ID:         uuid.MustParse("d4d4d4d4-d4d4-d4d4-d4d4-d4d4d4d4d4d4"),
			Title:      "1-on-1 with Manager",
			Content:    "Discussed career progression and current blockers.",
			Folder:     "Meetings",
			TagsIds:    []uuid.UUID{tags[0].ID}, // Work
			IsPinned:   false,
			IsArchived: false,
			SyncStatus: "synced",
			CreatedAt:  now.Add(-48 * time.Hour),
			UpdatedAt:  now.Add(-47 * time.Hour),
		},
		{
			ID:         uuid.MustParse("e5e5e5e5-e5e5-e5e5-e5e5-e5e5e5e5e5e5"),
			Title:      "Gym Routine V2",
			Content:    "Monday: Push, Tuesday: Pull, Wednesday: Legs.",
			Folder:     "Health",
			TagsIds:    []uuid.UUID{tags[1].ID}, // Personal
			IsPinned:   true,
			IsArchived: false,
			SyncStatus: "synced",
			CreatedAt:  now.Add(-300 * time.Hour),
			UpdatedAt:  now.Add(-10 * time.Hour),
		},
		{
			ID:         uuid.MustParse("f6f6f6f6-f6f6-f6f6-f6f6-f6f6f6f6f6f6"),
			Title:      "Server Outage Post-Mortem",
			Content:    "Database connection pool exhausted. Need to optimize queries.",
			Folder:     "Incidents",
			TagsIds:    []uuid.UUID{tags[0].ID, tags[3].ID}, // Work, Urgent
			IsPinned:   false,
			IsArchived: true,
			SyncStatus: "synced",
			CreatedAt:  now.Add(-800 * time.Hour),
			UpdatedAt:  now.Add(-790 * time.Hour),
		},
		{
			ID:         uuid.MustParse("77777777-7777-7777-7777-777777777777"),
			Title:      "Books to Read in 2026",
			Content:    "1. Dune\n2. Atomic Habits\n3. The Pragmatic Programmer",
			Folder:     "Hobbies",
			TagsIds:    []uuid.UUID{tags[1].ID, tags[2].ID}, // Personal, Ideas
			IsPinned:   false,
			IsArchived: false,
			SyncStatus: "error",
			CreatedAt:  now.Add(-400 * time.Hour),
			UpdatedAt:  now.Add(-400 * time.Hour),
		},
		{
			ID:         uuid.MustParse("88888888-8888-8888-8888-888888888888"),
			Title:      "Gift Ideas for Mom",
			Content:    "New gardening tools, a spa voucher, or a cooking class.",
			Folder:     "Personal",
			TagsIds:    []uuid.UUID{tags[1].ID, tags[2].ID}, // Personal, Ideas
			IsPinned:   false,
			IsArchived: false,
			SyncStatus: "synced",
			CreatedAt:  now.Add(-50 * time.Hour),
			UpdatedAt:  now.Add(-5 * time.Hour),
		},
		{
			ID:         uuid.MustParse("99999999-9999-9999-9999-999999999999"),
			Title:      "Client Feedback - Alpha Corp",
			Content:    "They want the dashboard to load faster and have dark mode.",
			Folder:     "Clients",
			TagsIds:    []uuid.UUID{tags[0].ID}, // Work
			IsPinned:   true,
			IsArchived: false,
			SyncStatus: "synced",
			CreatedAt:  now.Add(-15 * time.Hour),
			UpdatedAt:  now.Add(-15 * time.Hour),
		},
		{
			ID:         uuid.MustParse("00000000-0000-0000-0000-000000000000"),
			Title:      "Old Apartment Lease",
			Content:    "Contract details and deposit information for the 2024 apartment.",
			Folder:     "Archive",
			TagsIds:    []uuid.UUID{tags[1].ID}, // Personal
			IsPinned:   false,
			IsArchived: true,
			SyncStatus: "synced",
			CreatedAt:  now.Add(-10000 * time.Hour),
			UpdatedAt:  now.Add(-9000 * time.Hour),
		},
	}

	return tags, notes
}

func NewRepository() Repository {
	tags, notes := GetMockData()

	return &repository{
		mockNote: notes,
		mocktag:  tags,
	}
}

//#region Notes

func (r *repository) GetAllNotes() []Note {
	return r.mockNote
}

func (r *repository) GetNoteByID(id uuid.UUID) (Note, error) {
	for _, note := range r.mockNote {
		if note.ID == id {
			return note, nil
		}
	}
	return Note{}, fmt.Errorf("note with ID %s not found", id)
}

func (r *repository) StoreNewNote(newNote Note) {
	r.mockNote = append(r.mockNote, newNote)
}

func (r *repository) UpdateNote(updatedNote Note, id uuid.UUID) (uuid.UUID, error) {
	for index := range r.mockNote {
		if r.mockNote[index].ID == id {
			r.mockNote[index] = updatedNote
			return id, nil
		}
	}

	return uuid.Nil, fmt.Errorf("%w: no note exists with ID %s", ErrNoteNotFound, id)
}

func (r *repository) DeleteNote(id uuid.UUID) error {
	for index := range r.mockNote {
		if r.mockNote[index].ID == id {
			r.mockNote = append(r.mockNote[:index], r.mockNote[index+1:]...)
			return nil
		}
	}

	return fmt.Errorf("%w: no note exists with ID %s", ErrNoteNotFound, id)
}

//#endregion Notes

// #region tags
func (r *repository) GetAllTags() []Tag {
	return r.mocktag
}

func (r *repository) CreateNewTag(newTags []Tag) {
	r.mocktag = append(r.mocktag, newTags...)
}

func (r *repository) GetTagsByIds(tagIds []uuid.UUID) (tags []Tag, unknownTagIDs []string) {
	tagMap := make(map[string]Tag)
	for _, tag := range r.mocktag {
		tagMap[tag.ID.String()] = tag
	}

	tags = make([]Tag, 0, len(tagIds))
	unknownTagIDs = make([]string, 0)
	for _, tagID := range tagIds {
		if tag, exist := tagMap[tagID.String()]; exist {
			tags = append(tags, tag)
		} else {
			unknownTagIDs = append(unknownTagIDs, tagID.String())
		}
	}

	return tags, unknownTagIDs
}

// #endregion tags
