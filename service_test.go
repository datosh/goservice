package goservice

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestService(t *testing.T) {
	req, err := http.NewRequest("GET", "/echo/hello", nil)
	if err != nil {
		t.Fatalf("Unable to build request: %v", err)
	}

	rr := httptest.NewRecorder()
	s := NewService(8080)
	s.router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Wrong status code! Got %d, want %d", status, http.StatusOK)
	}

	expected := "hello"
	got := rr.Body.String()
	if got != expected {
		t.Errorf("Wrong response body! Got '%s', want '%s'", got, expected)
	}
}
