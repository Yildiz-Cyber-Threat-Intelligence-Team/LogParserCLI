package main

import (
	"fmt"
	"stealer-parser/parsers"
	"strings"
)

func main() {
	var filePath, keyword string
	fmt.Println("Enter the file path:")
	fmt.Scanln(&filePath)

	fmt.Println("Enter the search keyword:")
	fmt.Scanln(&keyword)

	if strings.HasSuffix(filePath, ".csv") {
		csvParser := parsers.CSVParser{}
		data := csvParser.Parse(filePath)
		fmt.Println("Searching in CSV data:")
		parsers.SearchInCSV(data, keyword)

	} else if strings.HasSuffix(filePath, ".json") {
		jsonParser := parsers.JSONParser{}
		data := jsonParser.Parse(filePath)
		fmt.Println("Searching in JSON data:")
		SearchInJSON(data, keyword)

	} else if strings.HasSuffix(filePath, ".xml") {
		xmlParser := parsers.XMLParser{}
		data := xmlParser.Parse(filePath)
		fmt.Println("Searching in XML data:")
		SearchInXML(data, keyword)

	} else if strings.HasSuffix(filePath, ".txt") {
		txtParser := parsers.TXTParser{}
		data := txtParser.Parse(filePath)
		fmt.Println("Searching in TXT data:")
		parsers.SearchInTXT(data, keyword)

	} else {
		fmt.Println("Unsupported file format.")
	}
}

// JSON içindeki log mesajlarını tarar ve anahtar kelimenin geçtiği yerleri bulur
func SearchInJSON(data []map[string]interface{}, keyword string) {
	found := false
	for _, entry := range data {
		// Her JSON girdisini terminalde yazdırarak kontrol ediyoruz
		fmt.Printf("Inspecting entry: %v\n", entry)

		if message, ok := entry["message"].(string); ok {
			if strings.Contains(message, keyword) {
				fmt.Printf("Found keyword '%s' in entry: %s\n", keyword, message)
				found = true
			}
		}
	}
	if !found {
		fmt.Println("No matches found for the keyword.")
	}
}

// XML içindeki verilerde anahtar kelimenin geçtiği yerleri bulur
func SearchInXML(data string, keyword string) {
	found := false
	if strings.Contains(data, keyword) {
		fmt.Printf("Found keyword '%s' in XML data: %s\n", keyword, data)
		found = true
	}
	if !found {
		fmt.Println("No matches found for the keyword.")
	}
}
