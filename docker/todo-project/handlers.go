package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

type TodoHandlers struct {
	todo TodoService
}

type templateData struct {
	Todos []Todo
}

func NewTodoHandlers(todo TodoService) *TodoHandlers {
	if todo == nil {
		panic("todohandlers unimplemented: todo service is nil")
	}

	return &TodoHandlers{todo: todo}
}

func (t *TodoHandlers) Home(w http.ResponseWriter, r *http.Request) error {
	todos, err := t.todo.GetAllTasks()
	templateData := templateData{Todos: todos}
	if err != nil {
		return fmt.Errorf("TodoHandlers.Home.GetAllTasks: %v", err)
	}

	ts, err := template.ParseFiles("./static/index.html")
	if err != nil {
		return fmt.Errorf("TodoHandlers.Home.ParseFiles: %v", err)
	}

	err = ts.Execute(w, templateData)
	if err != nil {
		return fmt.Errorf("TodoHandlers.Home.Executer: %v", err)
	}

	return nil
}

func (t *TodoHandlers) CreateTodo(w http.ResponseWriter, r *http.Request) error {
	err := r.ParseForm()
	if err != nil {
		return fmt.Errorf("TodoHandlers.CreateTodo.ParseForm: %v", err)
	}
	task := r.FormValue("task")
	if task == "" {
		return fmt.Errorf("TodoHandlers.CreateTodo.FormValue: form task is empty")
	}

	err = t.todo.CreateTask(Todo{Task: task, Status: false})
	if err != nil {
		return fmt.Errorf("TodoHandlers.CreateTodo.CreateTask: %v", err)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
	return nil
}

func (t *TodoHandlers) EditTodo(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return nil
	}

	if r.Method == http.MethodPost {
		method := r.PostFormValue("_method")
		if method != http.MethodPut {
			return fmt.Errorf("TodoHandlers.EditTodo.FormValueMethod: method is not put")
		}
		if method == http.MethodPut {
			id := r.PostFormValue("id")
			if id == "" {
				return fmt.Errorf("TodoHandlers.EditTodo.FormValueID: form id is empty")
			}
			idInt, err := strconv.Atoi(id)
			fmt.Println(idInt)
			if err != nil {
				return fmt.Errorf("TodoHandlers.EditTodo.Atoi ID: %v", err)
			}

			err = t.todo.UpdateTaskStatus(idInt)
			if err != nil {
				return fmt.Errorf("TodoHandlers.EditTodo.UpdateTaskStatus: %v", err)
			}
		}
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
	return nil
}

func (t *TodoHandlers) DeleteTodo(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return nil
	}

	if r.Method == http.MethodPost {
		method := r.PostFormValue("_method")
		if method != http.MethodDelete {
			return fmt.Errorf("TodoHandlers.EditTodo.FormValueMethod: method is not put")
		}
		if method == http.MethodDelete {
			id := r.PostFormValue("id")
			if id == "" {
				return fmt.Errorf("TodoHandlers.EditTodo.FormValueID: form id is empty")
			}
			idInt, err := strconv.Atoi(id)
			fmt.Println(idInt)
			if err != nil {
				return fmt.Errorf("TodoHandlers.EditTodo.Atoi ID: %v", err)
			}

			err = t.todo.DeleteTask(idInt)
			if err != nil {
				return fmt.Errorf("TodoHandlers.EditTodo.DeleteTask: %v", err)
			}
		}
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
	return nil
}
