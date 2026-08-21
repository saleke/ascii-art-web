package handlers

import (
	ascii_art "ascii-art-web/ascii-art"
	"bytes"
	"embed"
	"encoding/json"
	"html/template"
	"net/http"
	"os"
	"strings"
)

var TemplatesFS embed.FS
var BannersFS embed.FS

var allowedBanners = map[string]string{
	"acrobat": "acrobat.txt", "graceful": "graceful.txt", "graffiti": "graffiti.txt", "merlin": "merlin.txt",
	"miniwi": "miniwi.txt", "modular": "modular.txt", "ogre": "ogre.txt", "rectangles": "rectangles.txt",
	"shadow": "shadow.txt", "standard": "standard.txt", "temper": "temper.txt", "thinkertoy": "thinkertoy.txt", "train": "train.txt",
}

type PageData struct {
	Input, Output, Banner, Error string
	Banners                      []string
	Project                      ProjectInfo
}

func bannerFilePath(name string) (string, bool) {
	filename, ok := allowedBanners[name]
	if !ok {
		return "", false
	}
	content, err := BannersFS.ReadFile("banners/" + filename)
	if err != nil {
		return "", false
	}
	tmp, err := os.CreateTemp("", "banner-*.txt")
	if err != nil {
		return "", false
	}
	defer tmp.Close()
	if _, err := tmp.Write(content); err != nil {
		os.Remove(tmp.Name())
		return "", false
	}
	return tmp.Name(), true
}

func AsciiHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ascii-art" {
		http.NotFound(w, r)
		return
	}
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)
	if err := r.ParseForm(); err != nil {
		http.Error(w, "invalid form data", http.StatusBadRequest)
		return
	}
	data := PageData{Input: r.FormValue("input"), Banner: strings.ToLower(strings.TrimSpace(r.FormValue("banner"))), Banners: BannerNames(), Project: projectInfo()}
	api := strings.Contains(strings.ToLower(r.Header.Get("Accept")), "application/json")
	if strings.TrimSpace(data.Input) == "" {
		if api {
			writeJSON(w, data, "Please enter some text to generate.", http.StatusBadRequest)
			return
		}
		renderResult(w, data, "Please enter some text to generate.", http.StatusBadRequest)
		return
	}
	if len([]rune(data.Input)) > 500 {
		if api {
			writeJSON(w, data, "Please keep your message under 500 characters.", http.StatusBadRequest)
			return
		}
		renderResult(w, data, "Please keep your message under 500 characters.", http.StatusBadRequest)
		return
	}
	bannerFile, ok := bannerFilePath(data.Banner)
	if !ok {
		if api {
			writeJSON(w, data, "Please choose a supported banner.", http.StatusBadRequest)
			return
		}
		renderResult(w, data, "Please choose a supported banner.", http.StatusBadRequest)
		return
	}
	defer os.Remove(bannerFile)
	output, err := ascii_art.Generate(data.Input, bannerFile)
	if err != nil {
		if api {
			writeJSON(w, data, "The selected banner could not be loaded.", http.StatusNotFound)
			return
		}
		renderResult(w, data, "The selected banner could not be loaded.", http.StatusNotFound)
		return
	}
	data.Output = output
	if api {
		writeJSON(w, data, "", http.StatusOK)
		return
	}
	renderResult(w, data, "", http.StatusOK)
}

func CleanBannerFile(name string) string {
	path, _ := bannerFilePath(strings.ToLower(strings.TrimSpace(name)))
	return path
}

func BannerNames() []string {
	return []string{"standard", "shadow", "thinkertoy", "graffiti", "acrobat", "merlin", "temper", "graceful", "miniwi", "modular", "ogre", "rectangles", "train"}
}

func renderResult(w http.ResponseWriter, data PageData, message string, status int) {
	data.Error = message
	tmpl, err := template.ParseFS(TemplatesFS, "templates/result.html")
	if err != nil {
		http.Error(w, "result template is unavailable", http.StatusInternalServerError)
		return
	}
	var page bytes.Buffer
	if err := tmpl.Execute(&page, data); err != nil {
		http.Error(w, "result template could not be rendered", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(status)
	_, _ = page.WriteTo(w)
}

func writeJSON(w http.ResponseWriter, data PageData, message string, status int) {
	response := struct {
		Output string `json:"output"`
		Banner string `json:"banner"`
		Error  string `json:"error,omitempty"`
	}{Output: data.Output, Banner: data.Banner, Error: message}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(response)
}
