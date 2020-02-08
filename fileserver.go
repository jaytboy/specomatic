package main

import (
	"fmt"
	"net/http"
)

// SpecServer returns "test" upon request.
func SpecServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "test")
}
