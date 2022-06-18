package services

import (
	"context"
	"github.com/minhtuhcmus/coder-push-notes-mono-repo/backend/domain/repositories"
	"github.com/minhtuhcmus/coder-push-notes-mono-repo/backend/graph/model"
)

type NoteService interface {
	CreateNote(ctx context.Context, newNote *model.NewNote) (*model.OverviewNote, error)
	GetNoteDetail(ctx context.Context, noteID int) (*model.OverviewNote, error)
	FetchListNote(ctx context.Context, page, size, userID int) (*model.OverviewNote, error)
}

type noteService struct {
	noteRepository *repositories.NoteRepository
}

func (n noteService) CreateNote(ctx context.Context, newNote *model.NewNote) (*model.OverviewNote, error) {
	//TODO implement me
	panic("implement me")
}

func (n noteService) GetNoteDetail(ctx context.Context, noteID int) (*model.OverviewNote, error) {
	//TODO implement me
	panic("implement me")
}

func (n noteService) FetchListNote(ctx context.Context, page, size, userID int) (*model.OverviewNote, error) {
	//TODO implement me
	panic("implement me")
}

func NewNoteService(
	noteRepository *repositories.NoteRepository,
) NoteService {
	return &noteService{
		noteRepository: noteRepository,
	}
}
