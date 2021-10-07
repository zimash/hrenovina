package health

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlerHealth(t *testing.T) {
	request := httptest.NewRequest("GET", "/healthz", nil)
	response := httptest.NewRecorder()
	handler := http.HandlerFunc(HandlerHealth)

	handler.ServeHTTP(response, request)

	if status := response.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{}`
	if response.Body.String() != expected {
		t.Errorf("handler response unexpected body: got %v want %v", response.Body.String(), expected)
	}
}
