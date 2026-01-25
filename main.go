package main

import (
	"encoding/json"
	"net/http"

	//"github.com/gin-gonic/gin"
	"github.com/go-chi/chi/v5"
)

//клиент (транспорт) - http (сам запрос) -> createTaskHandler
//DTO (как json запроса/входные данные разложить на структуру) -> CreateTaskRequest struct
//Бизнес логика (функция) -> func NewTask
//модель данных для хранения -> Task struct
//ответ клиенту

var allTasks []Task

func createTaskHandler(w http.ResponseWriter, r *http.Request) {
	var req CreateTaskRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	task := NewTask(req.Name, req.Description)
	allTasks = append(allTasks, task)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(task)

}

func getAllTaskHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(allTasks)
}

//      !FOR GIN FRAME
//func createTaskHandler(c *gin.Context) {
//	var req CreateTaskRequest
//
//	err := c.BindJSON(&req) //BindJSON раскладывает json на структуру, возвращает только ошибку
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	task := NewTask(req.Name, req.Description)
//
//	allTasks = append(allTasks, task)
//	c.JSON(http.StatusCreated, task)
//}
//
//func getAllTaskHandler(c *gin.Context) {
//	c.JSON(http.StatusOK, allTasks)
//}

func main() {

	r2 := chi.NewRouter()
	r2.Post("/tasks", createTaskHandler)
	r2.Get("/tasks", getAllTaskHandler)
	http.ListenAndServe(":8080", r2)

	//r := gin.Default()
	//r.POST("/tasks", createTaskHandler)
	//r.GET("/tasks", getAllTaskHandler)
	//r.Run(":8080")
}
