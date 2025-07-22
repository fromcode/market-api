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

	"github.com/fromcode/market-api/internal/config"
	"github.com/fromcode/market-api/internal/http/handlers/market"
	"github.com/fromcode/market-api/internal/storage/sqlite"
)

func main() {
	// load config
	cfg := config.MustLoad()
	// database setup
	storage, err := sqlite.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	slog.Info("storage initialized", slog.String("env", cfg.Env), slog.String("version", "1.0.0"))
	// setup router
	router := http.NewServeMux()

	router.HandleFunc("POST /api/markets", market.New(storage))
	router.HandleFunc("GET /api/markets/{id}", market.GetById(storage))
	// setup server

	server := http.Server{ //bawaan package http.Server
		Addr:    cfg.Addr, // mengambil struct HTTPServer `yaml:"address" env-required:"true"`
		Handler: router,
	}

	slog.Info("Server mulai", slog.String("address", cfg.Addr)) //Print pertanda dimulai server beserta portnya

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe() //Listener pada website
		if err != nil {
			log.Fatal("Gagal mulai server")
		}
	}()

	<-done

	slog.Info("Matikan Server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Gagal mematikan server", slog.String("error", err.Error()))
	}

	slog.Info("Server berhasil dimatikan")
}
