package services

import (
	"api_service_questions_and_answers/internal/models"
	"api_service_questions_and_answers/internal/repositories"
)

type QuestionService interface {
	GetAllQuestions() ([]*models.Question, error)
	CreateQuestion(request *models.Question) (*models.Question, error)
	GetQuestion(id uint) (*models.Question, error)
	DeleteQuestion(id uint) error
}

type questionService struct {
	questionRepository repositories.QuestionRepository
}

func NewQuestionService(
	questionRepository repositories.QuestionRepository,
) QuestionService {
	return &questionService{
		questionRepository,
	}
}

func (q questionService) GetAllQuestions() ([]*models.Question, error) {
	return q.questionRepository.FindAll()
}

func (q questionService) CreateQuestion(question *models.Question) (*models.Question, error) {
	err := q.questionRepository.Create(question)
	if err != nil {
		return nil, err
	}

	return question, nil
}

func (q questionService) GetQuestion(id uint) (*models.Question, error) {
	return q.questionRepository.FindByID(id)
}

func (q questionService) DeleteQuestion(id uint) error {
	return q.questionRepository.Delete(id)
}
