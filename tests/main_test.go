package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mshra/renovatioBackend/handlers"
)

const (
	rootEndpointResponse string = "<html><body><a href=\"https://renovatio-design.vercel.app/\">Renovatio </a></body></html>"
)

func TestRoot(t *testing.T) {
	t.Run("testing root endpoint", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/", nil)
		responseRecorder := httptest.NewRecorder()

		handlers.Home(responseRecorder, request)

		response := responseRecorder.Body.String()

		if response != rootEndpointResponse {
			t.Errorf("got: %s\nwant: %s", response, rootEndpointResponse)
		}
	})
}
