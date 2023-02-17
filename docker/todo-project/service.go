package main

import (
	"database/sql"
	"fmt"
)

type TodoService interface {
	CreateTask(Todo) error
	DeleteTask(id int) error
	UpdateTaskStatus(id int) error
	GetAllTasks() ([]Todo, error)
}

type todoService struct {
	DB *sql.DB
}

func NewTodoService(db *sql.DB) *todoService {
	if db == nil {
		panic("todo service is nil")
	}
	return &todoService{db}
}

func (t *todoService) CreateTask(todo Todo) error {
	stmt, err := t.DB.Prepare(`INSERT INTO todos(task, status) VALUES(?, ?)`)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(todo.Task, todo.Status)
	if err != nil {
		return fmt.Errorf("INSERT TODO ERROR: %v", err)
	}

	return nil
}

func (t *todoService) DeleteTask(id int) error {
	stmt, err := t.DB.Prepare(`DELETE from todos WHERE id = ?`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return fmt.Errorf("DELETE TODO: %v", err)
	}

	return nil
}

func (t *todoService) UpdateTaskStatus(id int) error {
	todo, err := t.GetTaskByID(id)
	if err != nil {
		return err
	}
	stmt, err := t.DB.Prepare(`UPDATE todos SET status=? WHERE id = ?`)
	if err != nil {
		return err
	}
	rows, err := stmt.Exec(!todo.GetStatus(), todo.ID)
	if count, _ := rows.RowsAffected(); count != 1 {
		return fmt.Errorf("ROWS AFFECTED 0 ")
	}
	if err != nil {
		return fmt.Errorf("UPDATE TODO: %v", err)
	}

	return nil
}

func (t *todoService) GetTaskByID(id int) (*Todo, error) {
	stmt, err := t.DB.Prepare(`SELECT id, task, status FROM todos WHERE id = ?`)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}

	var todo Todo
	for rows.Next() {
		err := rows.Scan(&todo.ID, &todo.Task, &todo.Status)
		if err != nil {
			return nil, err
		}
	}

	if &todo == nil {
		return nil, fmt.Errorf("Task with id %d not found", id)
	}

	return &todo, nil
}

func (t *todoService) GetAllTasks() ([]Todo, error) {
	rows, err := t.DB.Query(`SELECT id, task, status FROM todos`)
	if err != nil {
		return nil, err
	}

	var todos []Todo
	for rows.Next() {
		var todo Todo
		err := rows.Scan(&todo.ID, &todo.Task, &todo.Status)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	return todos, nil
}
