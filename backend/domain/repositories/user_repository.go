package repositories

import (
	"context"
	"fmt"
	"github.com/minhtuhcmus/coder-push-notes-mono-repo/backend/database/datastore"
	"github.com/minhtuhcmus/coder-push-notes-mono-repo/backend/domain/models"
)

type UserRepository struct{}

var userRepository *UserRepository

func NewUserRepository() *UserRepository {
	if userRepository == nil {
		userRepository = &UserRepository{}
	}
	return userRepository
}

type AuthInfoSelector struct {
	ID             int
	PasswordHashed string
}

func (u *UserRepository) GetUserIDAndPasswordByUsername(ctx context.Context, username string) (int, string, error) {
	var authInfoSelector AuthInfoSelector
	if err := datastore.
		GetDB().
		WithContext(ctx).
		Model(&models.User{}).
		Where(&models.User{Username: username, Active: true}).
		Limit(1).
		Find(&authInfoSelector).Error; err != nil || (authInfoSelector.ID == 0 && authInfoSelector.PasswordHashed == "") {
		return 0, "", fmt.Errorf("error when getting password of user: %s, error %v", username, err)
	}
	return authInfoSelector.ID, authInfoSelector.PasswordHashed, nil
}

func (u *UserRepository) SaveUser(ctx context.Context, user *models.User) error {
	return datastore.GetDB().WithContext(ctx).Create(user).Error
}
