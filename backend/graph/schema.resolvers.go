package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/minhtuhcmus/coder-push-notes-mono-repo/backend/graph/generated"
	"github.com/minhtuhcmus/coder-push-notes-mono-repo/backend/graph/model"
)

func (r *mutationResolver) CreateNote(ctx context.Context, newNote model.NewNote) (*model.OverviewNote, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Notes(ctx context.Context) ([]*model.OverviewNote, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Note(ctx context.Context) (*model.OverviewNote, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
