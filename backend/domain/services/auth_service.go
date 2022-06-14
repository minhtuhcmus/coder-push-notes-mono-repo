package services

import (
	"context"
	"github.com/minhtuhcmus/coder-push-notes-mono-repo/backend/domain/models"
	"github.com/minhtuhcmus/coder-push-notes-mono-repo/backend/domain/repositories"
	"github.com/minhtuhcmus/coder-push-notes-mono-repo/backend/graph/model"
)

type AuthService interface {
	SignUp(ctx context.Context, newUser *model.NewUser) (string, error)
	SignIn(ctx context.Context, request *models.AuthRequest) (string, error)
}

type authService struct {
	userRepository *repositories.UserRepository
}

func (a *authService) SignUp(ctx context.Context, newUser *model.NewUser) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (a *authService) SignIn(ctx context.Context, request *models.AuthRequest) (string, error) {
	//TODO implement me
	panic("implement me")
}

func NewAuthService(
	userRepository *repositories.UserRepository,
) AuthService {
	return &authService{
		userRepository: userRepository,
	}
}
