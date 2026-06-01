package models

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	UserID uuid.UUID `gorm:"type:uuid;index;not null"`

	ExpiresAt time.Time `gorm:"index;not null"`

	CreatedAt time.Time
	UpdatedAt time.Time

	User User `gorm:"constraint:OnDelete:CASCADE;"`
}

func (s *Session) BeforeCreate(_ any) error {
	if s.ID == uuid.Nil {
		s.ID = uuid.New()
	}
	return nil
}
