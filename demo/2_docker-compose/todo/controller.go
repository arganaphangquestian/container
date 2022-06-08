package todo

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"
)

type todoController struct {
	service *todoService
}

func NewController(router *mux.Router, db *pgx.Conn) *todoController {
	c := &todoController{
		service: NewService(
			NewRepository(db),
		),
	}
	router.HandleFunc("/", c.getTodos).Methods(http.MethodGet)
	router.HandleFunc("/{id:[0-9]+}", c.getTodo).Methods(http.MethodGet)
	router.HandleFunc("/", c.createTodo).Methods(http.MethodPost)
	router.HandleFunc("/{id:[0-9]+}", c.deleteTodo).Methods(http.MethodDelete)
	return c
}

func (c *todoController) getTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	list := c.service.getTodos()
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "OK",
		"data":    list,
	})
}

func (c *todoController) getTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)["id"]
	id, err := strconv.Atoi(vars)
	if err != nil {
		http.Error(w, "Body invalid", http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "NOT OK",
		})
		return
	}
	todo := c.service.getTodo(id)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "OK",
		"data":    todo,
	})
}

func (c *todoController) createTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body := TodoModel{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Body invalid", http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "NOT OK",
		})
		return
	}
	todo := c.service.createTodo(body)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "OK",
		"data":    todo,
	})
}

func (c *todoController) deleteTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)["id"]
	id, err := strconv.Atoi(vars)
	if err != nil {
		http.Error(w, "Body invalid", http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "NOT OK",
		})
		return
	}
	res := c.service.deleteTodo(id)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "OK",
		"data":    res,
	})
}
