package main

// Основа задания
// Создайте RESTful API для управления списком задач (To-Do List). Реализуйте следующие маршруты:
// GET /tasks: Получить список всех задач.
// POST /tasks: Создать новую задачу

type StatusTask string

const (
	new    StatusTask = "new"
	active StatusTask = "active"
	done   StatusTask = "done"
)

type Task struct {
	ID          string
	Title       string
	Description string
	Status      StatusTask
}

type CreateUpdateTaskRequest struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

func NewTask(id, title, description string, status StatusTask) Task {
	return Task{ID: id, Title: title, Description: description, Status: status}

}

func CheckStatus(status StatusTask) bool {
	switch status {
	case new, active, done:
		return true
	default:
		return false
	}
}
