package main

import (
	"fmt"
	"log"
	"net/http"

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

	fmt.Printf("Server mulai %s", cfg.HTTPServer.Addr) //Print pertanda dimulai server beserta portnya

	err := server.ListenAndServe() //Listener pada untuk website
	if err != nil {
		log.Fatal("Gagal mulai server")
	}

}
