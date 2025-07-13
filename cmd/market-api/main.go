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

	"github.com/fromcode/market-api/internal/config"
)

func main() {
	// load config
	cfg := config.MustLoad()
	// database setup
	// setup router
	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) { //mengatur handler pada url "/"
		w.Write([]byte("Selamat datang di Market API"))
	})
	// setup server

	server := http.Server{ //bawaan package http.Server
		Addr:    cfg.Addr, // mengambil struct HTTPServer `yaml:"address" env-required:"true"`
		Handler: router,
	}

	slog.Info("Server mulai %s", slog.String("address", cfg.Addr))
	fmt.Printf("Server mulai %s", cfg.HTTPServer.Addr) //Print pertanda dimulai server beserta portnya

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

	err := server.Shutdown(ctx)
	if err != nil {
		slog.Error("Gagal mematikan server", slog.String("error", err.Error()))
	}

	slog.Info("Server berhasil dimatikan")

}
