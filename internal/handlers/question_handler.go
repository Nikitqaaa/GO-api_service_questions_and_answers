package handlers

import (
	"api_service_questions_and_answers/internal/helpers"
	"api_service_questions_and_answers/internal/models"
	"api_service_questions_and_answers/internal/services"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"gorm.io/gorm"
)

type QuestionHandler struct {
	service services.QuestionService
}

func NewQuestionHandler(service services.QuestionService) *QuestionHandler {
	return &QuestionHandler{
		service,
	}
}

// GetQuestions godoc
// @Summary Get all questions
// @Description Get a list of all questions
// @Tags questions
// @Accept json
// @Produce json
// @Success 200 {array} models.Question
// @Failure 500 {object} map[string]string
// @Router /api/questions [get]
func (h *QuestionHandler) GetQuestions(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	question, err := h.service.GetAllQuestions()
	if err != nil {
		http.Error(w, "Failed to get questions", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(question)
	if err != nil {
		log.Printf("Error encoding response: %v", err)
		return
	}
}

// CreateQuestion godoc
// @Summary Create a new question
// @Description Create a new question
// @Tags questions
// @Accept json
// @Produce json
// @Param question body models.Question true "Question object"
// @Success 201 {object} models.Question
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/questions [post]
func (h *QuestionHandler) CreateQuestion(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req models.Question

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := req.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	question := &models.Question{
		Text:      req.Text,
		CreatedAt: time.Now(),
	}

	createdQuestion, err := h.service.CreateQuestion(question)
	if err != nil {
		http.Error(w, "Failed to create question", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(createdQuestion)
	if err != nil {
		log.Printf("Error encoding response: %v", err)
		return
	}
}

// GetQuestion godoc
// @Summary Get question by ID
// @Description Get a specific question by its ID
// @Tags questions
// @Accept json
// @Produce json
// @Param id path int true "Question ID"
// @Success 200 {object} models.Question
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/questions/{id} [get]
func (h *QuestionHandler) GetQuestion(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id, err := helpers.ExtractIDFromPath(r)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	question, err := h.service.GetQuestion(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "Not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Failed to get question", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(question)
	if err != nil {
		log.Printf("Error encoding response: %v", err)
		return
	}
}

// DeleteQuestion godoc
// @Summary Delete a question
// @Description Delete a specific question by its ID
// @Tags questions
// @Accept json
// @Produce json
// @Param id path int true "Question ID"
// @Success 204 "No Content"
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/questions/{id} [delete]
func (h *QuestionHandler) DeleteQuestion(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id, err := helpers.ExtractIDFromPath(r)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	err = h.service.DeleteQuestion(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "Not found", http.StatusNotFound)
		}
		http.Error(w, "Failed to deleted question", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
