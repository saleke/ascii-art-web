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
	if rec.Code != http.StatusOK || !strings.Contains(rec.Body.String(), "ASCII Studio") {
		t.Fatalf("home response: status=%d body=%q", rec.Code, rec.Body.String()[:min(80, len(rec.Body.String()))])
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
