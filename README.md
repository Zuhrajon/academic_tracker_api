# Academic Tracker API

REST API для учета учебных данных: студентов, предметов, посещаемости и оценок.

Проект написан на Go с использованием Gin, PostgreSQL, Docker Compose и goose migrations. Код разделен на слои:

- `handler` - HTTP-ручки и ответы клиенту;
- `service` - бизнес-логика и валидация;
- `repository` - работа с PostgreSQL;
- `model` - структуры данных.

## Требования

- Go
- Docker и Docker Compose
- goose
- golangci-lint для проверки линтером

## Запуск

Поднять PostgreSQL и приложение:

```bash
make up
```

Команда собирает контейнеры и применяет миграции.

Остановить сервис и удалить volume базы:

```bash
make down
```

По умолчанию API доступен на:

```text
http://localhost:8080
```

PostgreSQL доступен на:

```text
localhost:5432
```

## Миграции

Создать новую миграцию:

```bash
make generate-migration name=create_users_table
```

Применить миграции:

```bash
make migration-up
```

Откатить все миграции:

```bash
make migration-down
```


## Проверки

Unit-тесты:

```bash
make test
```

Линтер:

```bash
make lint
```

Полная проверка:

```bash
make check
```

## Endpoints

### Students

```http
POST   /students
GET    /students
GET    /students/:id
PUT    /students/:id
DELETE /students/:id
```

### Subjects

```http
POST   /subjects
GET    /subjects
GET    /subjects/:id
PUT    /subjects/:id
DELETE /subjects/:id
```

### Attendance

```http
POST   /attendance
GET    /students/:id/attendance
PUT    /attendance/:id
DELETE /attendance/:id
```

### Grades

```http
POST   /grades
GET    /students/:id/grades
PUT    /grades/:id
DELETE /grades/:id
```

## Примеры запросов

Создать студента:

```json
{
  "first_name": "Zuhro",
  "last_name": "Karimova",
  "group_name": "CS-101",
  "course": 1
}
```

Создать предмет:

```json
{
  "subject_name": "Databases",
  "teacher_name": "Hasan",
  "semester": 2
}
```

Создать посещаемость:

```json
{
  "student_id": 1,
  "subject_id": 1,
  "lesson_date": "2026-07-05",
  "status": "present",
  "comment": "on time"
}
```

Создать оценку:

```json
{
  "student_id": 1,
  "subject_id": 1,
  "grade_date": "2026-07-05",
  "grade": 9,
  "comment": "good work"
}
```

## Статусы посещаемости

Допустимые значения поля `status`:

```text
present
absent
late
```

## Оценки

Поле `grade` должно быть в диапазоне от `1` до `10`.
