package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Unimplemented"))
			return
		}

		var user string

		// Mengambil Environment Variable "user"
		user = os.Getenv("user")

		// Kalau user kosong maka ubah jadi World
		if user == "" {
			user = "World"
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("Hello %s", user)))

	}))

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
