package graph

import (
	"github.com/minhtuhcmus/coder-push-notes-mono-repo/backend/domain/services"
	"github.com/minhtuhcmus/coder-push-notes-mono-repo/backend/graph/generated"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	authService *services.AuthService
}

func New(
	authService *services.AuthService,
) generated.Config {
	return generated.Config{Resolvers: &Resolver{
		authService: authService,
	}}
}
