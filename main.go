package main

import (
	"ascii-art-web/handlers"
	"log"
	"net/http"
)

func main() {

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/ascii-art-web", handlers.AsciiHandler)

	log.Println("======================= Server starting =======================")

	panic(http.ListenAndServe(":8080", nil))
}
