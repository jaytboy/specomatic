package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"strconv"
	"time"
)

// SpecServer returns "test" upon request.
func SpecServer(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))

		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("index.gtpl")
		t.Execute(w, token)
	} else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Println(handler.Header)
		//TODO: 	- Check that the header contains "officedocument"
		//			- Unzip the file
		//			- Extract all the goodies and put them in a JSON struct
	}

}
