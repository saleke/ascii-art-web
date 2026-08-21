package main

import (
	"ascii-art-web/handlers"
	"embed"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

//go:embed static
var staticFS embed.FS

func main() {
	server := &http.Server{Addr: serverAddress(), Handler: appHandler(), ReadHeaderTimeout: 5 * time.Second, IdleTimeout: 60 * time.Second}
	log.Printf("ASCII Art Web listening on %s", server.Addr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

func appHandler() http.Handler {
	mux := http.NewServeMux()
	staticContent, err := fs.Sub(staticFS, "static")
	if err != nil {
		log.Fatal(err)
	}
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(staticContent))))
	mux.HandleFunc("/", handlers.HomeHandler)
	mux.HandleFunc("/ascii-art", handlers.AsciiHandler)
	return logging(mux)
}

func serverAddress() string {
	port := strings.TrimPrefix(strings.TrimSpace(os.Getenv("PORT")), ":")
	if port == "" {
		port = "8080"
	}
	return ":" + port
}

func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
