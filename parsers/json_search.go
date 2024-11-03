package parsers

import (
	"fmt"
	"strings"
)


func SearchInJSON(data []map[string]interface{}, keyword string) {
	for _, entry := range data {
		if message, ok := entry["message"].(string); ok {
			if containsKeyword(message, keyword) {
				fmt.Printf("Found keyword '%s' in entry: %s\n", keyword, message)
			}
		}
	}
}


func containsKeyword(message, keyword string) bool {
	return strings.Contains(message, keyword)
}
