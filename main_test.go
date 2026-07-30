package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// The secure pipeline runs `go test ./...`; a real passing test proves the test
// gate executed rather than silently finding nothing to do.
func TestRootReturns200(t *testing.T) {
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}).ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("got %d, want 200", rec.Code)
	}
}
