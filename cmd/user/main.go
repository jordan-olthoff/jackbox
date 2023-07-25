package main

import (
	"embed"
	_ "github.com/lib/pq"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"io/fs"
	api "jackbox/pkg/pkg/user"
	"jackbox/pkg/pkg/userservice"
	"log"
	"net/http"
	"os"
)

//go:embed swagger
var swagger embed.FS

//go:embed frontend/dist
var frontend embed.FS

var secret = []byte("secret")

func main() {
	logger := log.New(os.Stdout, "INFO: ", log.LstdFlags|log.Lshortfile)

	// Create service instance.
	userService, err := userservice.New(secret, logger)
	if err != nil {
		logger.Fatal(err)
	}
	// Create generated server.
	srv, err := api.NewServer(userService, userService)
	if err != nil {
		logger.Fatal(err)
	}
	port := ":8080"
	mux := http.NewServeMux()
	dist, err := fs.Sub(frontend, "frontend/dist")
	if err != nil {
		logger.Fatal(err)
	}
	// Register vue app
	mux.Handle("/", http.FileServer(http.FS(dist)))
	// Register api
	mux.Handle("/api/v1/", http.StripPrefix("/api/v1", srv))
	// Register swagger files.
	mux.Handle("/swagger/", http.FileServer(http.FS(swagger)))
	// Register metrics handler.
	mux.Handle("/metrics/", promhttp.Handler())

	logger.Printf("listening on %v...\n", port)
	if err = http.ListenAndServe(port, mux); err != nil {
		logger.Fatal(err)
	}
}
