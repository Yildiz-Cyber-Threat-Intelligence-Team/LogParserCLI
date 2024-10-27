package parsers

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"os"
)

type XMLParser struct{}

type Entry struct {
	XMLName     xml.Name `xml:"entry"`
	URL         string   `xml:"URL"`
	Username    string   `xml:"Username"`
	Password    string   `xml:"Password"`
	Application string   `xml:"Application"`
}

func (p *XMLParser) Parse(filePath string) []Entry {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	byteValue, _ := ioutil.ReadAll(file)

	var entries []Entry
	err = xml.Unmarshal(byteValue, &entries)
	if err != nil {
		log.Fatalf("Failed to parse XML: %v", err)
	}

	return entries
}
