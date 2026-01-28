package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//1. GET /tasks: Получить список всех задач.
//2. GET /tasks/:id: Получить информацию о задаче по ее ID.
//3. POST /tasks: Создать новую задачу.
//
//4. PUT /tasks/:id: Обновить информацию о задаче по ее ID.
//
//5. DELETE /tasks/:id: Удалить задачу по ее ID.
//
//Используйте структуры и слайсы для хранения информации о задачах.
//Каждая задача должна иметь уникальный идентификатор (ID), название (Title), описание (Description) и статус (Status), который может быть Новая, В процессе или Завершена.

//клиент (транспорт) - http (сам запрос) -> createTaskHandler
//DTO (как json запроса/входные данные разложить на структуру) -> CreateTaskRequest struct
//Бизнес логика (функция) -> func NewTask
//модель данных для хранения -> Task struct
//ответ клиенту

var allTasks []Task

func createTaskHandler(c *gin.Context) {
	var req CreateUpdateTaskRequest

	err := c.BindJSON(&req) //BindJSON раскладывает json на структуру, возвращает только ошибку
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	s := StatusTask(req.Status)
	if CheckStatus(s) == true {
		task := NewTask(req.ID, req.Title, req.Description, StatusTask(req.Status))

		allTasks = append(allTasks, task)
		c.JSON(http.StatusCreated, task)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid status"})
	}
}

func getAllTaskHandler(c *gin.Context) {
	c.JSON(http.StatusOK, allTasks)
}

func getTaskByID(c *gin.Context) {
	idIn := c.Param("id")
	if idIn != "" {
		for _, task := range allTasks {
			if task.ID == idIn {
				c.JSON(http.StatusOK, task)
			} else {
				c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
			}
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
	}
}

func deleteByID(c *gin.Context) {
	idIn := c.Param("id")
	if idIn != "" {
		for i, _ := range allTasks {
			if allTasks[i].ID == idIn {
				allTasks = append(allTasks[:i], allTasks[i+1:]...)
				c.JSON(http.StatusOK, allTasks)
			} else {
				c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
			}
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
}

func updateById(c *gin.Context) {
	var req CreateUpdateTaskRequest

	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	s := StatusTask(req.Status)
	if CheckStatus(s) == false {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid status"})
		return
	}

	idIn := c.Param("id")
	if idIn != "" {
		for i, _ := range allTasks {
			if allTasks[i].ID == idIn {
				allTasks[i].Title = req.Title
				allTasks[i].Description = req.Description
				allTasks[i].Status = StatusTask(req.Status)
				c.JSON(http.StatusOK, allTasks[i])
			}
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

}

func main() {

	r := gin.Default()
	r.POST("/tasks", createTaskHandler)
	r.GET("/tasks", getAllTaskHandler)
	r.GET("/tasks/:id", getTaskByID)
	r.DELETE("/tasks/:id", deleteByID)
	r.PUT("/tasks/:id", updateById)
	r.Run(":8080")
}
