package notes

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Service interface {
	// notes
	FetchAllNotes() []NoteData
	FindNote(id string) (Note, error)
	CreateNewNote(noteData CreateNoteRequest) (uuid.UUID, error)
	UpdateNote(newNoteData UpdateNoteRequest, id string) (uuid.UUID, error)
	DeleteNote(id string) error

	// tags
	CreateNewTag(newTagNames []string) []Tag
}

type NoteUpdater func(n *Note)

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) FetchAllNotes() []NoteData {
	mockNotes := s.repo.GetAllNotes()
	notes := []NoteData{}

	for _, mockNote := range mockNotes {
		noteTags, unknowTagIds := s.repo.GetTagsByIds(mockNote.TagsIds)

		notes = append(notes, NoteData{
			ID:         mockNote.ID,
			Title:      mockNote.Title,
			Content:    mockNote.Content,
			Folder:     mockNote.Folder,
			Tags:       noteTags,
			IsPinned:   mockNote.IsPinned,
			IsArchived: mockNote.IsArchived,
			SyncStatus: mockNote.SyncStatus,
			CreatedAt:  mockNote.CreatedAt,
			UpdatedAt:  mockNote.UpdatedAt,
		})

		fmt.Println(unknowTagIds)
	}

	return notes
}

func (s *service) FindNote(id string) (Note, error) {
	noteID, err := uuid.Parse(id)

	if err != nil {
		return Note{}, fmt.Errorf("invalid note ID format")
	}

	return s.repo.GetNoteByID(noteID)
}

func (s *service) CreateNewNote(noteData CreateNoteRequest) (uuid.UUID, error) {
	now := time.Now()

	noteId := uuid.New()

	newTags := s.CreateNewTag(noteData.NewTagNames)

	noteTag := append(noteData.Tags, newTags...)

	tagIds := getTagIds(noteTag)

	s.repo.StoreNewNote(Note{
		ID:         noteId,
		IsPinned:   false,
		IsArchived: false,
		SyncStatus: "synced",
		CreatedAt:  now,
		UpdatedAt:  now,
		Title:      noteData.Title,
		Content:    noteData.Content,
		Folder:     noteData.Folder,
		TagsIds:    tagIds,
	})

	return noteId, nil
}

func (s *service) CreateNewTag(newTagNames []string) []Tag {
	newTags := make([]Tag, len(newTagNames))

	for i, tagName := range newTagNames {
		newTags[i] = Tag{
			ID:   uuid.New(),
			Name: tagName,
		}
	}

	return newTags
}

func (s *service) UpdateNote(newNoteData UpdateNoteRequest, id string) (uuid.UUID, error) {
	noteID, err := uuid.Parse(id)

	if err != nil {
		return uuid.Nil, fmt.Errorf("invalid note ID format")
	}

	note, err := s.repo.GetNoteByID(noteID)

	if err != nil {
		return uuid.Nil, err
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

	applyUpdate(&note, updates...)

	if len(updates) == 0 {
		return uuid.Nil, fmt.Errorf("at least one field must be provided for update")
	}

	note.UpdatedAt = time.Now()

	noteId, err := s.repo.UpdateNote(note, noteID)

	if err != nil {
		return uuid.Nil, err
	}

	return noteId, nil
}

func (s *service) DeleteNote(id string) error {
	noteID, err := uuid.Parse(id)

	if err != nil {
		return fmt.Errorf("invalid note ID format")
	}

	return s.repo.DeleteNote(noteID)
}

// helper
func getTagIds(tags []Tag) []uuid.UUID {
	ids := make([]uuid.UUID, len(tags))

	for i, tag := range tags {
		ids[i] = tag.ID
	}

	return ids
}

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

func updateTags(tags *[]Tag) NoteUpdater {
	return func(n *Note) {
		if tags != nil {
			tagIds := getTagIds(*tags)

			for i, tag := range *tags {
				tagIds[i] = tag.ID
			}

			n.TagsIds = tagIds
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
