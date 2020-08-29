package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

// Location is the location in the document
type Location struct {
	Level        int
	NumberSelf   int
	NumberParent int
}

// Paragraph contains the raw pharagraph information.
type Paragraph struct {
	Text     string
	Location Location
	Options  []string
	Comment  string
}

//Section contain the section information
type Section struct {
	Title     string
	Paragraph []Paragraph
	Comments  []Comment
}

// Comment is the struct for PStyle OMN
type Comment struct {
	Text      string
	Reference Location
}

func main() {
	b := new(Body)

	data, err := ioutil.ReadFile("../LS Mech MF 2012v - 23 09 93 - Mechanical Sequence of Operations/word/document.xml")
	if err != nil {
		panic(err)
	}

	if err := xml.Unmarshal(extractedBody(data), &b); err != nil {
		panic(err)
	}

	for _, Paragraph
}

func extractedBody(data []byte) []byte {
	leftTrim := bytes.SplitN(data, []byte("<w:body>"), 2)
	rightTrim := bytes.SplitN(leftTrim[1], []byte("</w:body>"), 2)
	comboSplit := [][]byte{[]byte("<w:body>"), rightTrim[0], []byte("</w:body>")}
	return bytes.Join(comboSplit, nil)
}
