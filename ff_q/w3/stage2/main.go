package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var tasks = []Task{}

func loadTasks() error {
	data, err := os.ReadFile("tasks.json")
	if err != nil {
		tasks = []Task{}
		return saveTasks()
	}
	return json.Unmarshal(data, &tasks)
}

func saveTasks() error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile("tasks.json", data, 0644)
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func getTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)                 // get from url
	id, err := strconv.Atoi(params["id"]) // convert string to int
	if err != nil {
		http.Error(w, "Wrong ID", http.StatusBadRequest)
		return
	}

	for _, task := range tasks {
		if task.ID == id {
			json.NewEncoder(w).Encode(task)
			return
		}
	}
	http.Error(w, "Task not found", http.StatusNotFound)
}

func createTask(w http.ResponseWriter, r *http.Request) {
	var newTask Task
	if err := json.NewDecoder(r.Body).Decode(&newTask); err != nil {
		http.Error(w, "Неверный JSON", http.StatusBadRequest)
		return
	}

	if newTask.Title == "" {
		http.Error(w, "Поле 'title' обязательно", http.StatusBadRequest)
		return
	}

	newTask.ID = len(tasks) + 1
	tasks = append(tasks, newTask)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)

	if err := saveTasks(); err != nil {
		http.Error(w, "Не удалось сохранить", http.StatusInternalServerError)
		return
	}
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r) // Получаем параметры из URL
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Неверный ID", http.StatusBadRequest)
		return
	}

	var updatedTask Task
	if err := json.NewDecoder(r.Body).Decode(&updatedTask); err != nil {
		http.Error(w, "Неверный JSON", http.StatusBadRequest)
		return
	}

	for i, task := range tasks {
		if task.ID == id {
			updatedTask.ID = id
			tasks[i] = updatedTask
			json.NewEncoder(w).Encode(updatedTask)
			return
		}
	}

	if err := saveTasks(); err != nil {
		http.Error(w, "Не удалось сохранить", http.StatusInternalServerError)
		return
	}
	http.Error(w, "Task not found", http.StatusNotFound)

}
func deleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r) // Получаем параметры из URL
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Неверный ID", http.StatusBadRequest)
		return
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...) // удаляет элемент по индексу
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	if err := saveTasks(); err != nil {
		http.Error(w, "Не удалось сохранить", http.StatusInternalServerError)
		return
	}
	http.Error(w, "Task not found", http.StatusNotFound)

}

func main() {

	if err := loadTasks(); err != nil {
		log.Fatal("Error loading tasks:", err)
	}
	r := mux.NewRouter() // Вместо http.HandleFunc

	// Маршруты
	r.HandleFunc("/tasks", getTasks).Methods("GET")
	r.HandleFunc("/tasks", createTask).Methods("POST")
	r.HandleFunc("/tasks/{id}", getTask).Methods("GET")
	r.HandleFunc("/tasks/{id}", updateTask).Methods("PUT")
	r.HandleFunc("/tasks/{id}", deleteTask).Methods("DELETE")

	r.Use(mux.CORSMethodMiddleware(r))

	log.Println("Сервер запущен на :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
