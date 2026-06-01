package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	// GitHub OAuth data
	GithubID  int64  `gorm:"uniqueIndex;not null"`
	Username  string `gorm:"not null"`
	Email     *string
	AvatarURL *string

	CreatedAt time.Time
	UpdatedAt time.Time

	Sessions []Session `gorm:"foreignKey:UserID"`
}

func (u *User) BeforeCreate(_ any) error {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return nil
}
