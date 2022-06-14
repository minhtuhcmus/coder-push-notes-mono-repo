package controllers

import (
	"context"
	"github.com/minhtuhcmus/coder-push-notes-mono-repo/backend/domain/models"
	"github.com/minhtuhcmus/coder-push-notes-mono-repo/backend/domain/services"
	"github.com/minhtuhcmus/coder-push-notes-mono-repo/backend/graph/model"
)

type authController struct {
	authServices *services.AuthService
}

func (a *authController) SignIn(ctx context.Context, authRequest *models.AuthRequest) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (a *authController) SignUp(ctx context.Context, newUser *model.NewUser) (string, error) {
	//TODO implement me
	panic("implement me")
}

type AuthController interface {
	SignIn(ctx context.Context, authRequest *models.AuthRequest) (string, error)
	SignUp(ctx context.Context, newUser *model.NewUser) (string, error)
}

func NewAuthController(
	authService *services.AuthService,
) AuthController {
	return &authController{
		authServices: authService,
	}
}
