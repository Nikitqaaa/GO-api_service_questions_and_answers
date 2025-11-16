package route

import (
	"api_service_questions_and_answers/internal/handlers"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func SetupQuestionRoutes(questionHandler *handlers.QuestionHandler, answerHandler *handlers.AnswerHandler) http.Handler {
	r := chi.NewRouter()

	r.Route("/api", func(r chi.Router) {
		r.Get("/questions", questionHandler.GetQuestions)
		r.Post("/questions", questionHandler.CreateQuestion)
		r.Get("/questions/{id}", questionHandler.GetQuestion)
		r.Delete("/questions/{id}", questionHandler.DeleteQuestion)

		r.Get("/answers/{id}", answerHandler.GetAnswer)
		r.Post("/questions/{id}/answers", answerHandler.CreateAnswer)
		r.Delete("/answers/{id}", answerHandler.DeleteAnswer)
	})

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Health check from %s", r.RemoteAddr)
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
		if err != nil {
			log.Printf("Error encoding response: %v", err)
			return
		}
	})

	return r
}
