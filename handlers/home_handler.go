package handlers

import (
	"bytes"
	"html/template"
	"net/http"
	"path/filepath"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	tmpl, err := template.ParseFiles(filepath.Join(projectRoot(), "templates", "index.html"))
	if err != nil {
		http.Error(w, "home template is unavailable", http.StatusInternalServerError)
		return
	}
	var page bytes.Buffer
	if err := tmpl.Execute(&page, PageData{Banner: "standard", Banners: BannerNames()}); err != nil {
		http.Error(w, "home template could not be rendered", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, _ = page.WriteTo(w)
}
