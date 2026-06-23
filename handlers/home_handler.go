package handlers

import (
	"fmt"
	"net/http"
	"os"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	file := "templates/index.html"
	resp, err := os.ReadFile(file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, string(resp))
}
