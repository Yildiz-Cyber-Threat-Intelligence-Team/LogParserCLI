package parsers

import (
	"fmt"
	"strings"
)


func SearchInCSV(data [][]string, keyword string) {
	for _, row := range data {
		for _, field := range row {
			if strings.Contains(field, keyword) {
				fmt.Println("Found in row:", row)
				break
			}
		}
	}
}
