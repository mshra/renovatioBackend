package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mshra/renovatioBackend/handlers"
)

func TestRoot(t *testing.T) {
	t.Run("testing root endpoint", func(t *testing.T) {
		const requiredResponse = "<html><body><a href=\"https://renovatio-design.vercel.app/\">Renovatio </a></body></html>"
		request := httptest.NewRequest(http.MethodGet, "/", nil)
		responseRecorder := httptest.NewRecorder()

		handlers.Home(responseRecorder, request)

		response := responseRecorder.Body.String()

		if response != requiredResponse {
			t.Errorf("response got: %s | response wanted: %s", response, requiredResponse)
		}
	})
	t.Run("testing with other http methods than GET", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodPost, "/", nil)
		responseRecorder := httptest.NewRecorder()

		handlers.Home(responseRecorder, request)

		if responseRecorder.Result().Status != "405 Method Not Allowed" {
			t.Errorf("got: %s | want: %s", responseRecorder.Result().Status, http.StatusText(http.StatusMethodNotAllowed))
		}
	})
}
