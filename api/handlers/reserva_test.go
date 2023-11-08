package handlers_test

import (
	"backend/api/handlers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestObtenerPaquetesByUser(t *testing.T) {
	// Mock the HTTP request and response objects
	req, err := http.NewRequest("GET", "/paquetes?id_usuario=snaranjo@utem.cl", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	// Call the function under test
	handlers.ObtenerPaquetesByUser(rr, req)

	// Check the response content type
	if rr.Header().Get("Content-Type") != "application/json" {
		t.Errorf("expected content type %v but got %v", "application/json", rr.Header().Get("Content-Type"))
	}
}
