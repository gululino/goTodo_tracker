package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"sync"
)

// Todo Model

type Todo struct {
	ID    int
	Title string
	Done  bool
}

var (
	todos   []Todo
	nextID  = 1
	todoMux sync.Mutex
)

// templates

var templates = template.Must(template.ParseFiles("index.html"))

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/add", addHandler)
	http.HandleFunc("/toggle", toggleHandler)
	http.HandleFunc("/api/todos", apiTodosHandler)

	log.Println("server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Frontend Handlers
func homeHandler(w http.ResponseWriter, r *http.Request) {
	todoMux.Lock()
	defer todoMux.Unlock()
	templates.ExecuteTemplate(w, "index.html", todos)
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		title := r.FormValue("title")
		if title != "" {
			todoMux.Lock()
			todos = append(todos, Todo{ID: nextID, Title: title})
			nextID++
			todoMux.Unlock()
		}
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func toggleHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)

	todoMux.Lock()
	for i := range todos {
		if todos[i].ID == id {
			todos[i].Done = !todos[i].Done
			break
		}
	}
	todoMux.Unlock()

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// backend API

func apiTodosHandler(w http.ResponseWriter, r *http.Request) {
	todoMux.Lock()
	defer todoMux.Unlock()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}
