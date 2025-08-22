package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Task struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Done bool `json:"done"`
}

var tasks = []Task{
	{ID: 1, Title: "Lern go", Done: false},
	{ID: 2, Title: "Whrite API", Done: true},
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
}
func main() {
	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "aplication/json")

	    switch r.Method {
    	case "GET":
        	json.NewEncoder(w).Encode(tasks)//	getTasks(w, r)
    	case "POST":
        	createTask(w, r)
    	default:
        	http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
    }
	})

	

	log.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}