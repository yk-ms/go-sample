package main
import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// TodoItem represents a TODO item
type TodoItem struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

// todoItems is a slice to store TODO items in memory
var todoItems []TodoItem

// getTodoItems returns all TODO items
func getTodoItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todoItems)
}

// createTodoItem creates a new TODO item
func createTodoItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var todoItem TodoItem
	_ = json.NewDecoder(r.Body).Decode(&todoItem)
	todoItem.ID = len(todoItems) + 1
	todoItems = append(todoItems, todoItem)
	json.NewEncoder(w).Encode(todoItem)
}

// getTodoItem returns a TODO item by ID
func getTodoItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	for _, item := range todoItems {
		if item.ID == id {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&TodoItem{})
}

// updateTodoItem updates a TODO item by ID
func updateTodoItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	for index, item := range todoItems {
		if item.ID == id {
			var updatedTodoItem TodoItem
			_ = json.NewDecoder(r.Body).Decode(&updatedTodoItem)
			updatedTodoItem.ID = id
			todoItems[index] = updatedTodoItem
			json.NewEncoder(w).Encode(updatedTodoItem)
			return
		}
	}
	json.NewEncoder(w).Encode(&TodoItem{})
}

// deleteTodoItem deletes a TODO item by ID
func deleteTodoItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	for index, item := range todoItems {
		if item.ID == id {
			todoItems = append(todoItems[:index], todoItems[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(todoItems)
}

func main() {
	// Init router
	r := mux.NewRouter()

	// Route handles & endpoints
	r.HandleFunc("/api/todo", getTodoItems).Methods("GET")
	r.HandleFunc("/api/todo", createTodoItem).Methods("POST")
	r.HandleFunc("/api/todo/{id}", getTodoItem).Methods("GET")
	r.HandleFunc("/api/todo/{id}", updateTodoItem).Methods("PUT")
	r.HandleFunc("/api/todo/{id}", deleteTodoItem).Methods("DELETE")

	// Start server
	log.Fatal(http.ListenAndServe(":8000", r))
}
