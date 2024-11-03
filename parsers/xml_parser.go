package parsers

import (
	"io/ioutil"
	"log"
	"os"
)


type XMLParser struct{}


func (p *XMLParser) Parse(filePath string) string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	return string(byteValue)
}
