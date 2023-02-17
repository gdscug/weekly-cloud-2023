package main

type Todo struct {
	ID     int
	Task   string
	Status bool
}

func NewTodo(task string, status bool) Todo {
	return Todo{
		Task:   task,
		Status: status,
	}
}

func (t Todo) SetDone() bool {
	t.Status = true
	return t.Status
}

func (t Todo) GetStatus() bool {
	return t.Status
}

func (t Todo) IsDone() bool {
	return t.Status == true
}
