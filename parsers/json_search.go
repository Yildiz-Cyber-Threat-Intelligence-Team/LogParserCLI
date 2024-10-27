package parsers

import (
	"fmt"
	"strings"
)

// JSON içindeki log mesajlarını tarar ve anahtar kelimenin geçtiği yerleri bulur
func SearchInJSON(data []map[string]interface{}, keyword string) {
	for _, entry := range data {
		if message, ok := entry["message"].(string); ok {
			if containsKeyword(message, keyword) {
				fmt.Printf("Found keyword '%s' in entry: %s\n", keyword, message)
			}
		}
	}
}

// Anahtar kelime kontrolü
func containsKeyword(message, keyword string) bool {
	return strings.Contains(message, keyword)
}
