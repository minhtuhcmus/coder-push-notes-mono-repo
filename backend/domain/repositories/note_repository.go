package repositories

import (
	"context"
	"fmt"
	"github.com/minhtuhcmus/coder-push-notes-mono-repo/backend/database/datastore"
	"github.com/minhtuhcmus/coder-push-notes-mono-repo/backend/domain/models"
	"github.com/minhtuhcmus/coder-push-notes-mono-repo/backend/graph/model"
)

type NoteRepository struct{}

var noteRepository *NoteRepository

func NewNoteRepository() *NoteRepository {
	if noteRepository == nil {
		noteRepository = &NoteRepository{}
	}
	return noteRepository
}

func (n *NoteRepository) CreateNote(ctx context.Context, newNote *models.Note) error {
	if err := datastore.GetDB().WithContext(ctx).Create(newNote).Error; err != nil {
		return fmt.Errorf("error create newNote %v", err)
	}

	return nil
}

func (n *NoteRepository) GetNoteDetail(ctx context.Context, userID, noteID int) (*models.Note, error) {
	var noteInfo models.Note

	if !n.GetPermissionOnNote(ctx, userID, noteID) {
		return nil, fmt.Errorf("error userID %d do not have permission on noteID %d", userID, noteID)
	}

	if err := datastore.
		GetDB().
		WithContext(ctx).
		Model(&models.Note{}).
		Where(&models.Note{ID: noteID}).Limit(1).Find(&noteInfo).Error; err != nil {
		return nil, fmt.Errorf("error getting note with id: %d detail. %v", noteID, err)
	}

	return &noteInfo, nil
}

func (n *NoteRepository) FetchListNote(ctx context.Context, page, size, userID int) ([]*models.Note, error) {
	var listNote []*models.Note

	if err := datastore.
		GetDB().
		WithContext(ctx).
		Model(&models.Note{}).
		Where(&models.Note{UserID: userID}).
		Offset((page - 1) * size).
		Limit(size).
		Find(&listNote).Error; err != nil {
		return nil, fmt.Errorf("error getting list note %v", err)
	}

	return listNote, nil
}

func (n *NoteRepository) UpdateNote(ctx context.Context, userID, noteID int, newNote *model.NewNote) (*models.Note, error) {
	if !n.GetPermissionOnNote(ctx, userID, noteID) {
		return nil, fmt.Errorf("error userID %d do not have permission on noteID %d", userID, noteID)
	}

	var noteToUpdate *models.Note

	noteToUpdate.ID = noteID
	noteToUpdate.UserID = userID

	if err := datastore.
		GetDB().
		WithContext(ctx).
		First(noteToUpdate).Error; err != nil {
		return nil, fmt.Errorf("error when get note to update noteID: %d, userID %d, error: %v", noteID, userID, err)
	}

	noteToUpdate.Title = newNote.Title
	noteToUpdate.Text = newNote.Text

	if err := datastore.
		GetDB().
		WithContext(ctx).
		Save(noteToUpdate).Error; err != nil {
		return nil, fmt.Errorf("error when save note %v. error: %v", noteToUpdate, err)
	}

	return noteToUpdate, nil
}

func (n *NoteRepository) GetPermissionOnNote(ctx context.Context, userID, noteID int) bool {
	var res bool
	if err := datastore.
		GetDB().
		WithContext(ctx).
		Raw("SELECT EXISTS(SELECT * FROM notes WHERE user_id = ? AND id = ?)", userID, noteID).
		Scan(&res).Error; err != nil {
		return false
	}
	return res
}
