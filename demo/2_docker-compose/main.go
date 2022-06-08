package main

import (
	"context"
	"demoapp/todo"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"
)

var port string

func main() {
	port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	db, err := pgx.Connect(context.Background(), "postgres://user:password@todo_database:5432/todos?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	r := mux.NewRouter()
	_ = todo.NewController(r, db)

	log.Println("Listening... 🚀")
	http.ListenAndServe("0.0.0.0:"+port, r)
}
