package store

import (
	"context"

	"github.com/miguelhigueradev/codeflow/services/auth-service/internal/db/models"
	"gorm.io/gorm"
)

type Store struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CreateUser(ctx context.Context, user *models.User) error {
	return s.db.WithContext(ctx).Create(user).Error
}

func (s *Store) FindUserByGithubID(ctx context.Context, githubID int64) (models.User, error) {
	var user models.User
	err := s.db.WithContext(ctx).Where("github_id = ?", githubID).First(&user).Error
	return user, err
}
