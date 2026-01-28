package main

// Основа задания
// Создайте RESTful API для управления списком задач (To-Do List). Реализуйте следующие маршруты:
// GET /tasks: Получить список всех задач.
// POST /tasks: Создать новую задачу

type Task struct {
	Name        string
	Description string
}

type CreateTaskRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func NewTask(name, description string) Task {
	return Task{Name: name, Description: description}

}

//Easy
//Повторить написанные запросы на другом фрейморке и рассказать, почему именно этот фрейморк выбран.
//
//Middle
//Повторить написанные запросы на двух других фрейморках и рассказать, почему именно эти фрейморки выбраны.
//

//Hard
//Повторить написанные запросы на двух других фрейморках и рассказать, почему именно эти фрейморки выбраны, + дополнить своим функционалом (добавить 1-3 запроса).
