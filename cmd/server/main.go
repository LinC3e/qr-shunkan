package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/LinC3e/shunkan-qr/internal/analytics"
	"github.com/LinC3e/shunkan-qr/internal/config"
	"github.com/LinC3e/shunkan-qr/internal/database"
	"github.com/LinC3e/shunkan-qr/internal/qr"
	"github.com/LinC3e/shunkan-qr/internal/router"
	"github.com/LinC3e/shunkan-qr/internal/utils"
)

func main() {

	cfg := config.Load()

	db, err := database.New(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Error connecting to DB: %v", err)
	}
	defer db.Close()
	utils.InitGeoDB("GeoLite2-Country.mmdb")
	log.Print("Country db.")

	repo := qr.NewRepository(db)
	service := qr.NewService(repo)
	analyticsRepo := analytics.NewRepository(db)
	analyticsService := analytics.NewService(analyticsRepo)

	qrHandler := qr.NewHandler(service, analyticsService)
	r := router.Setup(qrHandler)

	srv := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: r,
	}

	go func() {
		log.Printf("Server running on http://localhost:%s", cfg.Port)

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Shutdown error:", err)
	}

	log.Println("Server stopped")
}
