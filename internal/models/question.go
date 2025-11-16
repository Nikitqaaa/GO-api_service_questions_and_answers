package models

import (
	"errors"
	"strings"
	"time"
)

type Question struct {
	ID        int       `json:"id" gorm:"primary_key"`
	Text      string    `json:"text" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	Answers   []Answer  `json:"answers,omitempty" gorm:"foreignKey:QuestionID;constraint:OnDelete:CASCADE;"`
}

func (r *Question) Validate() error {
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
	return nil
}
