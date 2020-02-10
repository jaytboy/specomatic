package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// SpecServer returns "test" upon request.
func SpecServer(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	t, _ := template.ParseFiles("index.gtpl")
	t.Execute(w, nil)
}
