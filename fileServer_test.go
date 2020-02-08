package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETUploadPage(t *testing.T) {
	t.Run("returns an upload page", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/upload", nil)
		response := httptest.NewRecorder()

		SpecServer(response, request)

		got := response.Body.String()
		want := "test"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
