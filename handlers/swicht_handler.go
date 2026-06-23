package handlers

import (
	ascii_art "ascii-art-web/ascii-art"
	"errors"
	"html/template"
	"net/http"
)

func SwitchHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	data := struct {
		Input      string
		Output     string
		BannerFile string
		Error      error
	}{}

	responseSource := "templates/result.html"

	var tmpl *template.Template
	tmpl, data.Error = template.ParseFiles(responseSource)

	if data.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		tmpl.Execute(w, data)
		return
	}

	data.Input = r.URL.Query().Get("input")
	data.BannerFile = r.URL.Query().Get("banner")

	data.BannerFile = CleanBannerFile(data.BannerFile)

	data.Output, data.Error = ascii_art.Generate(data.Input, data.BannerFile)
	if data.Error != nil {
		tmpl.Execute(w, data)
		return
	}
	data.Error = errors.New("0")
	tmpl.Execute(w, data)
}
