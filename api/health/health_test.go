package health

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandlerHealth(t *testing.T) {
	request := httptest.NewRequest("GET", "/healthz", nil)
	response := httptest.NewRecorder()
	handler := http.HandlerFunc(HandlerHealth)

	handler.ServeHTTP(response, request)

	assert.Equal(t, response.Code, http.StatusOK)
	assert.Equal(t, response.Body.String(), `{}`)
}
