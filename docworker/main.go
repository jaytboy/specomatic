package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
)

const fileName = "061000 - CARPENTRY.docx"

func main() {
	rd, err := zip.OpenReader("./docs/" + fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer rd.Close()

	for _, f := range rd.File {
		if f.Name == "document.xml" {
			fmt.Printf("Contents of %s:\n", f.Name)
			rc, err := f.Open()
			if err != nil {
				log.Fatal(err)
			}
			_, err = io.CopyN(os.Stdout, rc, 68)
			if err != nil {
				log.Fatal(err)
			}
			rc.Close()
			fmt.Println()
		}
	}
}
