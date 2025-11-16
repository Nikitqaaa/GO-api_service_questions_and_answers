package repositories

import (
	"api_service_questions_and_answers/internal/models"

	"gorm.io/gorm"
)

type AnswerRepository interface {
	Create(answer *models.Answer) error
	FindByID(id uint) (*models.Answer, error)
	DeleteByID(id uint) error
}

type answerRepository struct {
	database *gorm.DB
}

func NewAnswerRepository(database *gorm.DB) AnswerRepository {
	return &answerRepository{
		database,
	}
}

func (a answerRepository) Create(answer *models.Answer) error {
	return a.database.Create(answer).Error
}

func (a answerRepository) FindByID(id uint) (*models.Answer, error) {
	var answer models.Answer
	err := a.database.First(&answer, id).Error
	if err != nil {
		return nil, err
	}
	return &answer, nil
}

func (a answerRepository) DeleteByID(id uint) error {
	return a.database.Delete(&models.Answer{}, id).Error
}
