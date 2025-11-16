# Questions & Answers API Service 
REST API сервис для вопросов и ответов, построенный на Go с использованием PostgreSQL.

### Технологии:
- Backend: Go 1.25
- Database: PostgreSQL
- ORM: GORM
- Router: Chi
- Containerization: Docker & Docker Compose
- Migrations: Goose
- Swagger: http-swagger

### Архитектура:
- models - сущности/модели
- repositories - работа с базой данных
- services - бизнес-логика
- handlers - HTTP обработчики
- config - конфигурация
- database - подключение к БД
- route - маршруты
- helpers - вспомогательные функции
- migrations - миграции
- docs - swagger

### Запуск сервиса

1. Клонируйте репозиторий

`git clone <repository-url>`

2. Перейдите в папку с проектом

`cd api_service_questions_and_answers`

3. Создайте и запустите контейнер 

`docker-compose build`
`docker-compose up -d`

Сервис будет доступен по адресу:

`http://localhost:8080`

4. Проверьте работу:

`curl http://localhost:8080/health`

5. Чтобы посмотреть логи контейнера выполните команду:

`docker-compose logs -f app`

6. Чтобы остановить и удалить контейнер выполните команду:

`docker-compose down`

7. Запуск тестов:

`go test ./internal/handlers -v`

8. Тестирование через postman:

`импортируйте json файл с коллекцией в postman`

9. Swagger доступен по маршруту:

`http://localhost:8080/swagger/index.html`

## API Endpoints

### Questions:

- GET `/api/questions` — список всех вопросов
- POST `/api/questions` — создать новый вопрос
- GET `/api/questions/{id}` — получить вопрос и все ответы на него
- DELETE `/api/questions/{id}` — удалить вопрос (вместе с ответами)

### Answers:

- POST `/api/questions/{id}/answers` — добавить ответ к вопросу
- GET `/api/answers/{id}` — получить конкретный ответ
- DELETE `/api/answers/{id}` — удалить ответ

### Health Check:

- GET `/health` - проверка статуса API
