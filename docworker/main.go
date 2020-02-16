package main

import (
	"archive/zip"
	"io"
	"log"
	"os"
)

const fileName = "061000 - CARPENTRY.docx"

func main() {
	// Unzips docx file
	rd, err := zip.OpenReader("../docs/" + fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer rd.Close()

	if _, err := os.Stat("../docs/word"); os.IsNotExist(err) {
		os.Mkdir("../docs/word", 0755)
	}

	// saves document.xml in folder
	for _, f := range rd.File {
		if f.Name == "word/document.xml" {
			r, err := f.Open()
			if err != nil {
				log.Fatal(err)
			}
			defer r.Close()

			fs, err := os.OpenFile("../docs/"+f.Name, os.O_WRONLY|os.O_CREATE, 0766)
			if err != nil {
				log.Fatal(err)
			}
			io.Copy(fs, r)
			if err := fs.Close(); err != nil {
				log.Fatal(err)
			}
		}
	}
}
