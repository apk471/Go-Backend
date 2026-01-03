package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/apk471/go-crud-api/internal/config"
	"github.com/apk471/go-crud-api/internal/http/handlers/api"
	"github.com/apk471/go-crud-api/internal/storage/sqlite"
)

func main() {
	fmt.Println("Welcome to CRUD API")
	cfg := config.MustLoad()

	router := http.NewServeMux()
	// router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Welcome to CRUD API"))
	// })

	storage, err := sqlite.New(cfg)

	if err != nil{
		log.Fatal(err)	
	}

	slog.Info("storage initialized", slog.String("env", cfg.Env), slog.String("version", "1.0.0"))

	router.HandleFunc("POST /api/users" , api.New(storage))
	router.HandleFunc("GET /api/users/{id}", api.GetById(storage))
	router.HandleFunc("GET /api/users" , api.GetList(storage))


	server := http.Server{
		Addr: cfg.HttpServer.Addr,
		Handler: router,
	}
	// Start the server
	slog.Info("server is running", "address", cfg.HttpServer.Addr)

	done := make(chan os.Signal, 1)
	// Listen for interrupt and terminate signals
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("server error", "error", err)
		}
	}()
	
	<- done

	slog.Info("server is shutting down")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("server shutdown error", "error", err)
	}

	slog.Info("server is shutdown")
}