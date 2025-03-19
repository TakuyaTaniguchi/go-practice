package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Todo represents a todo item
type Todo struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description,omitempty"`
	Priority    Priority `json:"priority"`
	Completed   bool     `json:"completed"`
}

// TodoCreate is the model for creating a new todo
type TodoCreate struct {
	Title       string   `json:"title"`
	Description string   `json:"description,omitempty"`
	Priority    Priority `json:"priority,omitempty"`
	Completed   bool     `json:"completed,omitempty"`
}

// Priority represents the priority of a todo item
type Priority string

// Priority constants
const (
	PriorityLow    Priority = "low"
	PriorityMedium Priority = "medium"
	PriorityHigh   Priority = "high"
)

// Global todos storage (in-memory database)
var todos = make(map[string]Todo)

// Main function
func main() {
	router := mux.NewRouter()

	// Routes
	router.HandleFunc("/todos", getTodos).Methods("GET")
	router.HandleFunc("/todos/{id}", getTodo).Methods("GET")
	router.HandleFunc("/todos", createTodo).Methods("POST")
	router.HandleFunc("/todos/{id}", updateTodo).Methods("PUT")
	router.HandleFunc("/todos/{id}", deleteTodo).Methods("DELETE")

	// Start server
	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

// Get all todos
func getTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	todoList := make([]Todo, 0, len(todos))
	for _, todo := range todos {
		todoList = append(todoList, todo)
	}
	
	json.NewEncoder(w).Encode(todoList)
}

// Get a specific todo
func getTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	params := mux.Vars(r)
	id := params["id"]
	
	todo, exists := todos[id]
	if !exists {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Todo not found"})
		return
	}
	
	json.NewEncoder(w).Encode(todo)
}

// Create a new todo
func createTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	var todoCreate TodoCreate
	err := json.NewDecoder(r.Body).Decode(&todoCreate)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request payload"})
		return
	}
	
	if todoCreate.Title == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Title is required"})
		return
	}
	
	// Set default values if not provided
	if todoCreate.Priority == "" {
		todoCreate.Priority = PriorityMedium
	}
	
	// Generate a simple ID (in a real app you would use a better method)
	id := fmt.Sprintf("%d", len(todos)+1)
	
	// Create the todo with the provided data
	todo := Todo{
		ID:          id,
		Title:       todoCreate.Title,
		Description: todoCreate.Description,
		Priority:    todoCreate.Priority,
		Completed:   todoCreate.Completed,
	}
	
	// Save to our in-memory database
	todos[id] = todo
	
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)
}

// Update an existing todo
func updateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	params := mux.Vars(r)
	id := params["id"]
	
	_, exists := todos[id]
	if !exists {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Todo not found"})
		return
	}
	
	var todoUpdate TodoCreate
	err := json.NewDecoder(r.Body).Decode(&todoUpdate)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request payload"})
		return
	}
	
	if todoUpdate.Title == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Title is required"})
		return
	}
	
	// Update the todo
	todo := Todo{
		ID:          id,
		Title:       todoUpdate.Title,
		Description: todoUpdate.Description,
		Priority:    todoUpdate.Priority,
		Completed:   todoUpdate.Completed,
	}
	
	todos[id] = todo
	
	json.NewEncoder(w).Encode(todo)
}

// Delete a todo
func deleteTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	params := mux.Vars(r)
	id := params["id"]
	
	_, exists := todos[id]
	if !exists {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Todo not found"})
		return
	}
	
	delete(todos, id)
	
	w.WriteHeader(http.StatusNoContent)
}