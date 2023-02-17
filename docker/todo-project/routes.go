package main

import (
	"log"
	"net/http"
)

type Routes struct {
	todoHandler *TodoHandlers
	mux         *http.ServeMux
	infoLog     *log.Logger
	errorLog    *log.Logger
}

func NewRoutes(todoHandlers *TodoHandlers, infoLog *log.Logger, errorLog *log.Logger) *Routes {
	return &Routes{
		todoHandler: todoHandlers,
		infoLog:     infoLog,
		errorLog:    errorLog,
	}
}

type handlerFunc func(w http.ResponseWriter, r *http.Request) error

func (ro *Routes) MakeHandlerFunc(handler handlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := handler(w, r)
		if err != nil {
			ro.errorLog.Println(err.Error())
		}
	}
}

func (ro *Routes) todoRoutes() {
	ro.mux.HandleFunc("/", ro.MakeHandlerFunc(ro.todoHandler.Home))
	ro.mux.HandleFunc("/todo/create", ro.MakeHandlerFunc(ro.todoHandler.CreateTodo))
	ro.mux.HandleFunc("/todo/update", ro.MakeHandlerFunc(ro.todoHandler.EditTodo))
	ro.mux.HandleFunc("/todo/delete", ro.MakeHandlerFunc(ro.todoHandler.DeleteTodo))
}

func (ro *Routes) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()
	ro.mux = mux
	ro.todoRoutes()

	return LogRequests(ro.mux, ro.infoLog)
}
