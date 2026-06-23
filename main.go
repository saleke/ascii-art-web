package main

import (
	"ascii-art-web/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/ascii-art-web", handlers.AsciiHandler)

	log.Println("======================= Server starting =======================")

	panic(http.ListenAndServe(":8080", nil))
}
