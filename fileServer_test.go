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
		want := uploadText

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

const uploadText = `<!DOCTYPE html>
<html>
<head>
<title>Upload File</title>
</head>
<body>
<form enctype="multipart/form-data" action="http://localhost:5000/upload/" method="post">
<input type="file" name="uploadfile" id="">
<input type="hidden" name="token" value="">
<input type="submit" value="upload">
</form>
</body>
</html>`
