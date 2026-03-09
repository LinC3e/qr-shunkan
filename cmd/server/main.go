package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/LinC3e/shunkan-qr/internal/config"
	"github.com/LinC3e/shunkan-qr/internal/database"
	handler "github.com/LinC3e/shunkan-qr/internal/handlers"
	"github.com/LinC3e/shunkan-qr/internal/qr"
	"github.com/LinC3e/shunkan-qr/internal/router"
)

func main() {
	cfg := config.Load()

	db, err := database.New(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Error connecting to DB: %v", err)
	}
	defer db.Close()

	// Repos → Service → Handler
	repo := qr.NewRepository(db)
	service := qr.NewService(repo)
	qrHandler := handler.NewQRHandler(service)

	r := router.Setup(qrHandler)

	srv := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: r,
	}

	// Init server goroutine
	go func() {
		log.Printf("Server on in http://localhost:%s", cfg.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error server init: %v", err)
		}
	}()

	// shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Error in shutdown:", err)
	}

	log.Println("Server Stop. Ok")
}