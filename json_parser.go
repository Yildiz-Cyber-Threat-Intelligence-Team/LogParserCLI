package parsers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// JSONParser yapısı
type JSONParser struct{}

// Parse fonksiyonu: JSON dosyasını okur ve içerik olarak bir dilim (slice) döner
func (p *JSONParser) Parse(filePath string) []map[string]interface{} {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	var data []map[string]interface{}
	if err := json.Unmarshal(byteValue, &data); err != nil {
		log.Fatalf("Failed to parse JSON: %v", err)
	}

	return data
}
