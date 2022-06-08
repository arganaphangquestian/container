package todo

import (
	"context"
	"demoapp/data"

	"github.com/jackc/pgx/v4"
)

type todoRepository struct {
	queries *data.Queries
}

func NewRepository(db *pgx.Conn) *todoRepository {
	queries := data.New(db)
	return &todoRepository{
		queries: queries,
	}
}

func (r *todoRepository) getTodos() []TodoModel {
	data, err := r.queries.GetTodos(context.Background())
	if err != nil {
		return nil
	}
	var res []TodoModel
	for _, v := range data {
		res = append(res, TodoModel{
			ID:        uint64(v.ID),
			Title:     v.Title,
			Done:      v.Done,
			CreatedAt: v.CreatedAt,
		})
	}
	return res
}

func (r *todoRepository) getTodo(id int) *TodoModel {
	data, err := r.queries.GetTodo(context.Background(), int64(id))
	if err != nil {
		return nil
	}
	return &TodoModel{
		ID:        uint64(data.ID),
		Title:     data.Title,
		Done:      data.Done,
		CreatedAt: data.CreatedAt,
	}
}

func (r *todoRepository) createTodo(todo TodoModel) *TodoModel {
	data, err := r.queries.CreateTodo(context.Background(), todo.Title)
	if err != nil {
		return nil
	}
	return &TodoModel{
		ID:        uint64(data.ID),
		Title:     data.Title,
		Done:      data.Done,
		CreatedAt: data.CreatedAt,
	}
}

func (r *todoRepository) deleteTodo(id int) bool {
	err := r.queries.DeleteAuthor(context.Background(), int64(id))
	if err != nil {
		return false
	}
	return true
}
