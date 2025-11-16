package handlers

import (
	"api_service_questions_and_answers/internal/models"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

type MockQuestionService struct {
	mock.Mock
}

func (m *MockQuestionService) GetAllQuestions() ([]*models.Question, error) {
	args := m.Called()
	return args.Get(0).([]*models.Question), args.Error(1)
}

func (m *MockQuestionService) CreateQuestion(question *models.Question) (*models.Question, error) {
	args := m.Called(question)
	return args.Get(0).(*models.Question), args.Error(1)
}

func (m *MockQuestionService) GetQuestion(id uint) (*models.Question, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Question), args.Error(1)
}

func (m *MockQuestionService) DeleteQuestion(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestQuestionHandler_GetQuestions_Success(t *testing.T) {
	mockService := new(MockQuestionService)
	handler := NewQuestionHandler(mockService)

	expectedQuestions := []*models.Question{
		{ID: 1, Text: "First question", CreatedAt: time.Now()},
		{ID: 2, Text: "Second question", CreatedAt: time.Now()},
	}

	mockService.On("GetAllQuestions").Return(expectedQuestions, nil)

	req := httptest.NewRequest("GET", "/questions", nil)
	rr := httptest.NewRecorder()

	handler.GetQuestions(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))

	var response []*models.Question
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	require.NoError(t, err)
	assert.Len(t, response, 2)
	assert.Equal(t, expectedQuestions[0].Text, response[0].Text)
}

func TestQuestionHandler_GetQuestions_ServiceError(t *testing.T) {
	mockService := new(MockQuestionService)
	handler := NewQuestionHandler(mockService)

	mockService.On("GetAllQuestions").Return([]*models.Question{}, errors.New("database error"))

	req := httptest.NewRequest("GET", "/questions", nil)
	rr := httptest.NewRecorder()

	handler.GetQuestions(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}

func TestQuestionHandler_CreateQuestion_Success(t *testing.T) {
	mockService := new(MockQuestionService)
	handler := NewQuestionHandler(mockService)

	createRequest := models.Question{Text: "New question"}
	expectedQuestion := &models.Question{
		ID:        1,
		Text:      "New question",
		CreatedAt: time.Now(),
	}

	mockService.On("CreateQuestion", mock.AnythingOfType("*models.Question")).Return(expectedQuestion, nil)

	requestBody, _ := json.Marshal(createRequest)

	req := httptest.NewRequest("POST", "/questions", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler.CreateQuestion(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))

	var response models.Question
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	require.NoError(t, err)
	assert.Equal(t, expectedQuestion.ID, response.ID)
	assert.Equal(t, expectedQuestion.Text, response.Text)
}

func TestQuestionHandler_CreateQuestion_InvalidJSON(t *testing.T) {
	mockService := new(MockQuestionService)
	handler := NewQuestionHandler(mockService)

	req := httptest.NewRequest("POST", "/questions", bytes.NewBufferString("invalid json"))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler.CreateQuestion(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestQuestionHandler_CreateQuestion_ValidationError(t *testing.T) {
	mockService := new(MockQuestionService)
	handler := NewQuestionHandler(mockService)

	createRequest := models.Question{Text: "Hi"}
	requestBody, _ := json.Marshal(createRequest)

	req := httptest.NewRequest("POST", "/questions", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler.CreateQuestion(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestQuestionHandler_GetQuestion_Success(t *testing.T) {
	mockService := new(MockQuestionService)
	handler := NewQuestionHandler(mockService)

	expectedQuestion := &models.Question{
		ID:        1,
		Text:      "Test question",
		CreatedAt: time.Now(),
	}

	mockService.On("GetQuestion", uint(1)).Return(expectedQuestion, nil)

	req := httptest.NewRequest("GET", "/questions/1", nil)
	rr := httptest.NewRecorder()

	handler.GetQuestion(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var response models.Question
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	require.NoError(t, err)
	assert.Equal(t, expectedQuestion.ID, response.ID)
	assert.Equal(t, expectedQuestion.Text, response.Text)
}

func TestQuestionHandler_GetQuestion_NotFound(t *testing.T) {
	mockService := new(MockQuestionService)
	handler := NewQuestionHandler(mockService)

	mockService.On("GetQuestion", uint(999)).Return(&models.Question{}, gorm.ErrRecordNotFound)

	req := httptest.NewRequest("GET", "/questions/999", nil)
	rr := httptest.NewRecorder()

	handler.GetQuestion(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code)
}

func TestQuestionHandler_DeleteQuestion_Success(t *testing.T) {
	mockService := new(MockQuestionService)
	handler := NewQuestionHandler(mockService)

	mockService.On("DeleteQuestion", uint(1)).Return(nil)

	req := httptest.NewRequest("DELETE", "/questions/1", nil)
	rr := httptest.NewRecorder()

	handler.DeleteQuestion(rr, req)

	assert.Equal(t, http.StatusNoContent, rr.Code)
}

func TestQuestionHandler_DeleteQuestion_NotFound(t *testing.T) {
	mockService := new(MockQuestionService)
	handler := NewQuestionHandler(mockService)

	mockService.On("DeleteQuestion", uint(999)).Return(gorm.ErrRecordNotFound)

	req := httptest.NewRequest("DELETE", "/questions/999", nil)
	rr := httptest.NewRecorder()

	handler.DeleteQuestion(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code)
}

func TestQuestionHandler_MethodNotAllowed(t *testing.T) {
	mockService := new(MockQuestionService)
	handler := NewQuestionHandler(mockService)

	req := httptest.NewRequest("PUT", "/questions/1", nil)
	rr := httptest.NewRecorder()

	handler.GetQuestion(rr, req)

	assert.Equal(t, http.StatusMethodNotAllowed, rr.Code)
}
