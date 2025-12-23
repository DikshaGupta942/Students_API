package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/DikshaGupta942/student_API/internal/config"
	"github.com/DikshaGupta942/student_API/internal/http/handlers/student"
)

func main() {
	//local config setup
	cfg := config.MustLoad()

	//database setup
	//setup router

	router := http.NewServeMux()
	//start server
	router.HandleFunc("POST /api/student", student.New())
	//w.Write([]byte("Welcome to Student API"))

	server := http.Server{
		Addr:    cfg.Httpserver.Address,
		Handler: router,
	}

	slog.Info("Starting Student API server", slog.String("address", cfg.Httpserver.Address))

	//fmt.Printf("Hello, Student API started at server %s !", cfg.Httpserver.Address)

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {

		err := server.ListenAndServe()

		if err != nil {
			log.Fatal("Failed to start server:", err)
		}
	}()

	<-done

	slog.Info("Server stopped")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Failed to shutdown server", slog.String("error", err.Error())) //"error", err)
	}

	slog.Info("Server exited properly")
	//server.Shutdown()

	//fmt.Printf("Server started at %s\n", cfg.Httpserver.Address)
}
