package parsers

import (
	"bufio"
	"log"
	"os"
)

// TXTParser yapısı
type TXTParser struct{}

// Parse fonksiyonu: TXT dosyasını okur ve içerik olarak döner
func (p *TXTParser) Parse(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	return lines
}
