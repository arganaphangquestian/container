package todo

import "time"

type TodoModel struct {
	ID        uint64    `json:"id"`
	Title     string    `json:"title"`
	Done      bool      `json:"done"`
	CreatedAt time.Time `json:"created_at"`
}
