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
