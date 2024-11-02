package parsers

import (
	"encoding/csv"
	"log"
	"os"
	"strings"
)

type CSVParser struct {
	expectedFields int 
}

func NewCSVParser(expectedFields int) *CSVParser {
	return &CSVParser{expectedFields: expectedFields}
}

func (p *CSVParser) Parse(filePath string) [][]string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Printf("Failed to open file: %s, error: %v", filePath, err)
		return nil 
	}
	defer func() {
		if cerr := file.Close(); cerr != nil {
			log.Printf("Failed to close file: %s, error: %v", filePath, cerr)
		}
	}()

	reader := csv.NewReader(file)
	var validRecords [][]string
	lineNumber := 0

	for {
		record, err := reader.Read()
		lineNumber++

		if err != nil {
			if err.Error() == "EOF" {
				break 
			}
			log.Printf("Failed to read CSV record on line %d: %v", lineNumber, err)
			continue 
		}

		
		if len(record) == 0 || allFieldsEmpty(record) {
			log.Printf("Warning: empty record on line %d", lineNumber)
			continue 
		}

		
		if p.expectedFields > 0 && len(record) != p.expectedFields {
			log.Printf("Warning: record on line %d has wrong number of fields: expected %d, got %d", lineNumber, p.expectedFields, len(record))
			continue 
		}

		validRecords = append(validRecords, record)
	}

	log.Printf("Successfully parsed %d valid records from %s", len(validRecords), filePath)
	return validRecords
}


func allFieldsEmpty(record []string) bool {
	for _, field := range record {
		if strings.TrimSpace(field) != "" {
			return false 
		}
	}
	return true 
}


func (p *CSVParser) FindRowsWithKeyword(records [][]string, keyword string) [][]string {
	var matchingRecords [][]string

	for _, record := range records {
		for _, field := range record {
			if strings.Contains(strings.ToLower(field), strings.ToLower(keyword)) {
				matchingRecords = append(matchingRecords, record)
				break 
			}
		}
	}

	return matchingRecords
}
