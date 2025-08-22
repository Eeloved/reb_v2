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

func main() {
	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "aplication/json")

		if r.Method == "GET" {
			json.NewEncoder(w).Encode(tasks)
		} else { 
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}		
	})

	log.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}