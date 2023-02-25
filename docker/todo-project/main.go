package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	db, err := NewMySQL()
	if err != nil {
		panic(err)
	}
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	todoService := NewTodoService(db)
	todoHandler := NewTodoHandlers(todoService)
	router := NewRoutes(todoHandler, infoLog, errorLog)

	mux := router.RegisterRoutes()

	s := http.Server{
		Addr:         ":8000",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		log.Println("Server listening on http://localhost:8000")
		err := s.ListenAndServe()
		if err != nil {
			log.Printf("Error starting server: %v", err)
		}

	}()

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	sig := <-c

	log.Println("Got signal:", sig)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	s.Shutdown(ctx)

}
