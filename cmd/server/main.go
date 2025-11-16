package main

// @title Questions and Answers API
// @version 1.0
// @description API Service for managing questions and answers

// @host localhost:8080
// @BasePath /api
// @query.collection.format multi
import (
	"api_service_questions_and_answers/internal/config"
	"api_service_questions_and_answers/internal/database"
	"api_service_questions_and_answers/internal/handlers"
	"api_service_questions_and_answers/internal/repositories"
	"api_service_questions_and_answers/internal/route"
	"api_service_questions_and_answers/internal/services"
	"log"
	"net/http"

	_ "api_service_questions_and_answers/docs"

	"github.com/swaggo/http-swagger"
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

	mux := http.NewServeMux()
	mux.Handle("/swagger/", httpSwagger.WrapHandler)
	mux.Handle("/", apiRoute)

	serverAddr := cfg.Server.Address + ":" + cfg.Server.Port
	log.Printf("Server starting on %s", serverAddr)
	log.Printf("Swagger documentation available at http://localhost:8080/swagger/index.html")

	if err := http.ListenAndServe(serverAddr, mux); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
