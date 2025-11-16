# Questions & Answers API Service 
REST API сервис для вопросов и ответов, построенный на Go с использованием PostgreSQL.

### Технологии:
- Backend: Go 1.25
- Database: PostgreSQL
- ORM: GORM
- Router: Chi
- Containerization: Docker & Docker Compose
- Migrations: Goose

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

### Запуск сервиса

1.Клонируйте репозиторий

`git clone <repository-url>`

2.Перейдите в папку с проектом

`cd api_service_questions_and_answers`

3.Создайте и запустите контейнер 

`docker-compose build`
`docker-compose up -d`

Сервис будет доступен по адресу:

`http://localhost:8080`

4.Проверьте работу:

`curl http://localhost:8080/health`

5.Чтобы посмотреть логи контейнера выполните команду:

`docker-compose logs -f app`

5.Чтобы остановить и удалить контейнер выполните команду:

`docker-compose down`

6.Запск тестов:

`go test ./internal/handlers -v`

7.Тестирование через postman:

`импортируйте json файл с коллекцией в postman`

## API Endpoints

### Questions:

- GET `/questions` — список всех вопросов
- POST `/questions` — создать новый вопрос
- GET `/questions/{id}` — получить вопрос и все ответы на него
- DELETE `/questions/{id}` — удалить вопрос (вместе с ответами)

### Answers:

- POST `/questions/{id}/answers` — добавить ответ к вопросу
- GET `/answers/{id}` — получить конкретный ответ
- DELETE `/answers/{id}` — удалить ответ

### Health Check:

- GET `/health` - проверка статуса API
