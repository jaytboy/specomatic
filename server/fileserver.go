package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

// SpecServer returns "test" upon request.
func SpecServer(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	var fileName string
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

		contentType := handler.Header.Get("Content-Type")
		fileName = handler.Filename
		if strings.Contains(contentType, "officedocument") {
			fmt.Printf("The document %q was uploaded successfully!\n", handler.Filename)
			f, err := os.OpenFile("./docs/"+fileName, os.O_WRONLY|os.O_CREATE, 0644)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer f.Close()
			io.Copy(f, file)
		}
	}

	UnzipDoc(fileName)
	dat, err := ioutil.ReadFile("../docs/word/document.xml")
	if err != nil {
		panic(err)
	}
	sec, err := ProcessDoc(dat)
	if err != nil {
		fmt.Printf("Error processing doc: %v", err)
	}

	fmt.Printf("Processed %q", sec.Title)

}
