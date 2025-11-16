package models

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Answer struct {
	ID         int       `json:"id" gorm:"primary_key"`
	QuestionID uint      `json:"question_id" gorm:"not null"`
	UserID     uuid.UUID `json:"user_id" gorm:"not null"`
	Text       string    `json:"text" gorm:"not null"`
	CreatedAt  time.Time `json:"created_at"`
}

func (r *Answer) Validate() error {
	text := strings.TrimSpace(r.Text)
	if text == "" {
		return errors.New("text is required")
	}
	if len(text) < 5 {
		return errors.New("text must be at least 5 characters")
	}
	if len(text) > 1000 {
		return errors.New("text cannot exceed 1000 characters")
	}
	if r.UserID == uuid.Nil {
		return errors.New("user id required")
	}
	return nil
}
