package handlers

import (
	"api_service_questions_and_answers/internal/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockAnswerService struct {
	mock.Mock
}

func (m *MockAnswerService) CreateAnswer(questionID uint, answer *models.Answer) (*models.Answer, error) {
	args := m.Called(questionID, answer)
	return args.Get(0).(*models.Answer), args.Error(1)
}

func (m *MockAnswerService) GetAnswer(answerID uint) (*models.Answer, error) {
	args := m.Called(answerID)
	return args.Get(0).(*models.Answer), args.Error(1)
}

func (m *MockAnswerService) DeleteAnswer(answerID uint) error {
	args := m.Called(answerID)
	return args.Error(0)
}

func TestAnswerHandler_CreateAnswer_Success(t *testing.T) {
	mockService := new(MockAnswerService)
	handler := NewAnswerHandler(mockService)

	userID := uuid.MustParse("123e4567-e89b-12d3-a456-426614174000")

	createRequest := models.Answer{
		Text:   "This is a answer",
		UserID: userID,
	}

	expectedAnswer := &models.Answer{
		ID:         1,
		Text:       "This is a answer",
		QuestionID: 1,
		UserID:     userID,
		CreatedAt:  time.Now(),
	}

	mockService.On("CreateAnswer", uint(1), mock.AnythingOfType("*models.Answer")).Return(expectedAnswer, nil)

	requestBody, _ := json.Marshal(createRequest)

	req := httptest.NewRequest("POST", "/questions/1/answers", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler.CreateAnswer(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)

	var response models.Answer
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, expectedAnswer.ID, response.ID)
}

func TestAnswerHandler_CreateAnswer_QuestionNotFound(t *testing.T) {
	mockService := new(MockAnswerService)
	handler := NewAnswerHandler(mockService)

	userID := uuid.MustParse("123e4567-e89b-12d3-a456-426614174000")

	createRequest := models.Answer{
		Text:   "This is a helpful answer",
		UserID: userID,
	}

	mockService.On("CreateAnswer", uint(999), mock.AnythingOfType("*models.Answer")).Return(
		&models.Answer{}, errors.New("question not found"),
	)

	requestBody, _ := json.Marshal(createRequest)

	req := httptest.NewRequest("POST", "/questions/999/answers", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler.CreateAnswer(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code)
}
