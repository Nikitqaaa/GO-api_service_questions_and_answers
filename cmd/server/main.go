package main

import (
	"api_service_questions_and_answers/internal/config"
	"api_service_questions_and_answers/internal/database"
	"api_service_questions_and_answers/internal/handlers"
	"api_service_questions_and_answers/internal/repositories"
	"api_service_questions_and_answers/internal/route"
	"api_service_questions_and_answers/internal/services"
	"log"
	"net/http"
)

func main() {
	cfg := config.LoadConfig()

	db, err := database.NewDatabase(cfg.DB)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	questionRepo := repositories.NewQuestionRepository(db.DB)
	questionService := services.NewQuestionService(questionRepo)
	questionHandler := handlers.NewQuestionHandler(questionService)

	answerRepo := repositories.NewAnswerRepository(db.DB)
	answerService := services.NewAnswerService(questionRepo, answerRepo)
	answerHandler := handlers.NewAnswerHandler(answerService)

	apiRoute := route.SetupQuestionRoutes(questionHandler, answerHandler)

	serverAddr := cfg.Server.Address + ":" + cfg.Server.Port
	log.Printf("Server starting on %s", serverAddr)

	if err := http.ListenAndServe(serverAddr, apiRoute); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
