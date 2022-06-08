package todo

type todoService struct {
	repo *todoRepository
}

func NewService(repo *todoRepository) *todoService {
	return &todoService{
		repo: repo,
	}
}

func (s *todoService) getTodos() []TodoModel {
	return s.repo.getTodos()
}

func (s *todoService) getTodo(id int) *TodoModel {
	return s.repo.getTodo(id)
}

func (s *todoService) createTodo(todo TodoModel) *TodoModel {
	return s.repo.createTodo(todo)
}

func (s *todoService) deleteTodo(id int) bool {
	return s.repo.deleteTodo(id)
}
