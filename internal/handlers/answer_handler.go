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

type AnswerHandler struct {
	service services.AnswerService
}

func NewAnswerHandler(service services.AnswerService) *AnswerHandler {
	return &AnswerHandler{
		service,
	}
}

// GetAnswer godoc
// @Summary Get answer by ID
// @Description Get a specific answer by its ID
// @Tags answers
// @Accept json
// @Produce json
// @Param id path int true "Answer ID"
// @Success 200 {object} models.Answer
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/answers/{id} [get]
func (h *AnswerHandler) GetAnswer(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id, err := helpers.ExtractIDFromPath(r)
	if err != nil {
		http.Error(w, "Failed to get answer", http.StatusBadRequest)
		return
	}

	answer, err := h.service.GetAnswer(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "Not found", http.StatusBadRequest)
			return
		}
		http.Error(w, "Failed to get answer", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(answer)
	if err != nil {
		log.Printf("Error encoding response: %v", err)
		return
	}
}

// CreateAnswer godoc
// @Summary Create a new answer for a question
// @Description Create a new answer for a specific question
// @Tags answers
// @Accept json
// @Produce json
// @Param question_id path int true "Question ID"
// @Param answer body models.Answer true "Answer object"
// @Success 201 {object} models.Answer
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/questions/{question_id}/answers [post]
func (h *AnswerHandler) CreateAnswer(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id, err := helpers.ExtractIDFromPath(r)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	var req models.Answer

	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := req.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	answer := &models.Answer{
		Text:      req.Text,
		UserID:    req.UserID,
		CreatedAt: time.Now(),
	}

	createdAnswer, err := h.service.CreateAnswer(id, answer)
	if err != nil {
		if err.Error() == "question not found" {
			http.Error(w, "Question not found", http.StatusNotFound)
			return
		}
		log.Printf("Service error creating answer: %v", err)
		http.Error(w, "Failed to create answer", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(createdAnswer)
	if err != nil {
		log.Printf("Error encoding response: %v", err)
		return
	}
}

// DeleteAnswer godoc
// @Summary Delete an answer
// @Description Delete a specific answer by its ID
// @Tags answers
// @Accept json
// @Produce json
// @Param id path int true "Answer ID"
// @Success 204 "No Content"
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/answers/{id} [delete]
func (h *AnswerHandler) DeleteAnswer(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id, err := helpers.ExtractIDFromPath(r)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	err = h.service.DeleteAnswer(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "Not found", http.StatusNotFound)
		}
		http.Error(w, "Failed to deleted answer", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
