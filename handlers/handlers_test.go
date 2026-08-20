package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestHomeHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	HomeHandler(rec, req)
	if rec.Code != http.StatusOK || !strings.Contains(rec.Body.String(), "ASCII Art Web") || !strings.Contains(rec.Body.String(), "Aleke Emmanuel Solomon") {
		t.Fatalf("home response: status=%d body=%q", rec.Code, rec.Body.String()[:min(80, len(rec.Body.String()))])
	}
}

func TestHomeHandlerRejectsUnknownRoutes(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/missing", nil)
	rec := httptest.NewRecorder()
	HomeHandler(rec, req)
	if rec.Code != http.StatusNotFound {
		t.Fatalf("status=%d, want 404", rec.Code)
	}
}

func TestHomeHandlerRejectsUnsupportedMethods(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	rec := httptest.NewRecorder()
	HomeHandler(rec, req)
	if rec.Code != http.StatusMethodNotAllowed || rec.Header().Get("Allow") != http.MethodGet {
		t.Fatalf("status=%d allow=%q", rec.Code, rec.Header().Get("Allow"))
	}
}

func TestAsciiHandler(t *testing.T) {
	form := url.Values{"input": {"Hello"}, "banner": {"standard"}}
	req := httptest.NewRequest(http.MethodPost, "/ascii-art", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rec := httptest.NewRecorder()
	AsciiHandler(rec, req)
	if rec.Code != http.StatusOK || !strings.Contains(rec.Body.String(), "ASCII OUTPUT") {
		t.Fatalf("render response: status=%d", rec.Code)
	}
}

func TestAsciiHandlerRejectsInvalidBanner(t *testing.T) {
	form := url.Values{"input": {"Hello"}, "banner": {"../../etc/passwd"}}
	req := httptest.NewRequest(http.MethodPost, "/ascii-art", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rec := httptest.NewRecorder()
	AsciiHandler(rec, req)
	if rec.Code != http.StatusBadRequest {
		t.Fatalf("status=%d, want 400", rec.Code)
	}
}

func TestAsciiHandlerRejectsEmptyInput(t *testing.T) {
	form := url.Values{"input": {"   "}, "banner": {"standard"}}
	req := httptest.NewRequest(http.MethodPost, "/ascii-art", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rec := httptest.NewRecorder()
	AsciiHandler(rec, req)
	if rec.Code != http.StatusBadRequest || !strings.Contains(rec.Body.String(), "Please enter some text") {
		t.Fatalf("status=%d body=%q", rec.Code, rec.Body.String())
	}
}

func TestAllBannersGenerateOutput(t *testing.T) {
	for _, banner := range BannerNames() {
		t.Run(banner, func(t *testing.T) {
			form := url.Values{"input": {"A"}, "banner": {banner}}
			req := httptest.NewRequest(http.MethodPost, "/ascii-art", strings.NewReader(form.Encode()))
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			req.Header.Set("Accept", "application/json")
			rec := httptest.NewRecorder()
			AsciiHandler(rec, req)
			if rec.Code != http.StatusOK || !strings.Contains(rec.Body.String(), `"output"`) {
				t.Fatalf("status=%d body=%q", rec.Code, rec.Body.String())
			}
		})
	}
}

func TestProjectLinksCanBeConfigured(t *testing.T) {
	t.Setenv("PORTFOLIO_URL", "https://portfolio.example")
	t.Setenv("GITHUB_REPOSITORY_URL", "https://github.example/project")
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	HomeHandler(rec, req)
	if !strings.Contains(rec.Body.String(), "https://portfolio.example") || !strings.Contains(rec.Body.String(), "https://github.example/project") {
		t.Fatalf("configured project links were not rendered")
	}
}

func TestAsciiHandlerJSONPreview(t *testing.T) {
	form := url.Values{"input": {"Preview"}, "banner": {"shadow"}}
	req := httptest.NewRequest(http.MethodPost, "/ascii-art", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")
	rec := httptest.NewRecorder()
	AsciiHandler(rec, req)
	if rec.Code != http.StatusOK || !strings.Contains(rec.Header().Get("Content-Type"), "application/json") || !strings.Contains(rec.Body.String(), `"banner":"shadow"`) {
		t.Fatalf("preview response: status=%d content-type=%q body=%q", rec.Code, rec.Header().Get("Content-Type"), rec.Body.String())
	}
}
