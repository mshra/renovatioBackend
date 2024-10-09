package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mshra/renovatioBackend/handlers"
)

func TestCreateGeneration(t *testing.T) {
	t.Run("testing with non-allowed HTTP method", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/api/createGeneration", nil)

		responseRecorder := httptest.NewRecorder()

		handlers.CreateGeneration(responseRecorder, request)

		if responseRecorder.Result().Status != "405 Method Not Allowed" {
			t.Error("Endpoint allowing other HTTP methods than POST.")
		}
	})
}
