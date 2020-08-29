package main

import (
	"archive/zip"
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// P accepts a document paragraph structure
type P struct {
	XMLName      xml.Name `xml:"p"`
	Text         string   `xml:",chardata"`
	ParaID       string   `xml:"paraId,attr"`
	TextID       string   `xml:"textId,attr"`
	RsidR        string   `xml:"rsidR,attr"`
	RsidRDefault string   `xml:"rsidRDefault,attr"`
	RsidP        string   `xml:"rsidP,attr"`
	PPr          struct {
		Text   string `xml:",chardata"`
		PStyle struct {
			Text string `xml:",chardata"`
			Val  string `xml:"val,attr"`
		} `xml:"pStyle"`
	} `xml:"pPr"`
	R []struct {
		Text    string `xml:",chardata"`
		RsidRPr string `xml:"rsidRPr,attr"`
		T       struct {
			Text  string `xml:",chardata"`
			Space string `xml:"space,attr"`
		} `xml:"t"`
		RPr struct {
			Text   string `xml:",chardata"`
			RStyle struct {
				Text string `xml:",chardata"`
				Val  string `xml:"val,attr"`
			} `xml:"rStyle"`
		} `xml:"rPr"`
	} `xml:"r"`
}

// Paragraph contains the raw pharagraph information.
type Paragraph struct {
	Text   string
	Parent int
}

//Section contain the section information
type Section struct {
	Title     string
	Paragraph []Paragraph
}

// ProcessDoc processes the document.xml file.
func ProcessDoc(data []byte) (sec *Section, err error) {
	var vs []P
	sec = new(Section)
	p := new(Paragraph)
	paraLevel := make(map[string]int)
	n := 0
	var sb strings.Builder

	b := splitParagraphs(data, "<w:body>", "</w:p>")

	v := new(P)
	for _, s := range b {
		v = &P{}
		err := xml.Unmarshal(s, &v)
		if err != nil {
			fmt.Printf("error: %v", err)
			return nil, err
		}
		vs = append(vs, *v)
	}

	for _, phrase := range vs[0].R {
		sb.WriteString(phrase.T.Text)
	}

	sec.Title = sb.String()
	vs = vs[1:]

	// Creats a hash table of paragraph levels
	for _, paragraph := range vs {
		_, ok := paraLevel[paragraph.PPr.PStyle.Val]
		if !ok {
			paraLevel[paragraph.PPr.PStyle.Val] = n
			n++
		}
	}

	for _, paragraph := range vs {
		sb.Reset()
		for _, phrase := range paragraph.R {
			sb.WriteString(phrase.T.Text)
		}
		p.Parent = paraLevel[paragraph.PPr.PStyle.Val]
		p.Text = sb.String()
		sec.Paragraph = append(sec.Paragraph, *p)
	}

	return sec, nil
}

func splitParagraphs(data []byte, leftText, endTest string) [][]byte {
	a := bytes.SplitAfterN(data, []byte(leftText), 2)
	a = bytes.SplitAfter(a[1], []byte(endTest))
	a = a[:(len(a) - 1)][:]
	return a
}

// UnzipDoc unzips the word doc and saves the document.xml file in a folder
func UnzipDoc(fileName string) {
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
