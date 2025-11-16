package services

import (
	"api_service_questions_and_answers/internal/models"
	"api_service_questions_and_answers/internal/repositories"
	"errors"

	"gorm.io/gorm"
)

type AnswerService interface {
	CreateAnswer(questionId uint, request *models.Answer) (*models.Answer, error)
	GetAnswer(id uint) (*models.Answer, error)
	DeleteAnswer(id uint) error
}

type answerService struct {
	questionRepository repositories.QuestionRepository
	answerRepository   repositories.AnswerRepository
}

func NewAnswerService(
	questionRepository repositories.QuestionRepository,
	answerRepository repositories.AnswerRepository,
) AnswerService {
	return &answerService{
		questionRepository,
		answerRepository,
	}
}

func (a answerService) CreateAnswer(questionId uint, request *models.Answer) (*models.Answer, error) {
	_, err := a.questionRepository.FindByID(questionId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("question not found")
		}
		return nil, err
	}

	answer := &models.Answer{
		QuestionID: questionId,
		UserID:     request.UserID,
		Text:       request.Text,
		CreatedAt:  request.CreatedAt,
	}

	err = a.answerRepository.Create(answer)
	if err != nil {
		return nil, err
	}

	return answer, nil
}

func (a answerService) GetAnswer(id uint) (*models.Answer, error) {
	return a.answerRepository.FindByID(id)
}

func (a answerService) DeleteAnswer(id uint) error {
	return a.answerRepository.DeleteByID(id)
}
