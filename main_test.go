package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestServerAddress(t *testing.T) {
	t.Setenv("PORT", "10000")
	if got := serverAddress(); got != ":10000" {
		t.Fatalf("serverAddress() = %q", got)
	}
}

func TestServerAddressDefaultsTo8080(t *testing.T) {
	t.Setenv("PORT", "")
	if got := serverAddress(); got != ":8080" {
		t.Fatalf("serverAddress() = %q", got)
	}
}

func TestStaticStylesAreServed(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/static/style.css", nil)
	rec := httptest.NewRecorder()
	appHandler().ServeHTTP(rec, req)
	if rec.Code != http.StatusOK || !strings.Contains(rec.Body.String(), ":root") {
		t.Fatalf("static response: status=%d", rec.Code)
	}
}
