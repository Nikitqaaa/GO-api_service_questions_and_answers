package repositories

import (
	"api_service_questions_and_answers/internal/models"

	"gorm.io/gorm"
)

type QuestionRepository interface {
	FindAll() ([]*models.Question, error)
	Create(question *models.Question) error
	FindByID(id uint) (*models.Question, error)
	Delete(id uint) error
}

type questionRepository struct {
	database *gorm.DB
}

func NewQuestionRepository(database *gorm.DB) QuestionRepository {
	return &questionRepository{
		database,
	}
}

func (q questionRepository) FindAll() ([]*models.Question, error) {
	var questions []*models.Question
	err := q.database.Find(&questions).Error
	if err != nil {
		return nil, err
	}
	return questions, nil
}

func (q questionRepository) Create(question *models.Question) error {
	return q.database.Create(question).Error
}

func (q questionRepository) FindByID(id uint) (*models.Question, error) {
	var question models.Question
	err := q.database.Preload("Answers").First(&question, id).Error
	if err != nil {
		return nil, err
	}
	return &question, nil
}

func (q questionRepository) Delete(id uint) error {
	return q.database.Delete(&models.Question{}, id).Error
}
