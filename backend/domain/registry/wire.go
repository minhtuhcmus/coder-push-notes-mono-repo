//go:build wireinject
// +build wireinject

package registry

import (
	"context"
	"github.com/google/wire"
	"github.com/minhtuhcmus/coder-push-notes-mono-repo/backend/domain/handler"
	"github.com/minhtuhcmus/coder-push-notes-mono-repo/backend/domain/middlewares"
	"github.com/minhtuhcmus/coder-push-notes-mono-repo/backend/domain/repositories"
	"github.com/minhtuhcmus/coder-push-notes-mono-repo/backend/domain/services"
	"net/http"
)

func InitHTTPServer(ctx context.Context) (http.Handler, error) {
	wire.Build(
		handler.NewHTTPServer,
		middlewares.NewMiddleware,

		services.NewAuthService,
		services.NewNoteService,

		repositories.NewUserRepository,
		repositories.NewNoteRepository,
	)
	return nil, nil
}
