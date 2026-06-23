package handlers

import (
	ascii_art "ascii-art-web/ascii-art"
	"errors"
	"html/template"
	"net/http"
	"path/filepath"
)

func AsciiHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
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

	data.Input = r.FormValue("input")
	data.BannerFile = r.FormValue("banner")

	data.BannerFile = CleanBannerFile(data.BannerFile)

	data.Output, data.Error = ascii_art.Generate(data.Input, data.BannerFile)
	if data.Error != nil {
		tmpl.Execute(w, data)
		return
	}
	data.Error = errors.New("0")
	tmpl.Execute(w, data)
}

func CleanBannerFile(bannerFile string) string {
	bannerFile = filepath.Join("banners/", bannerFile) + ".txt"
	return bannerFile
}
