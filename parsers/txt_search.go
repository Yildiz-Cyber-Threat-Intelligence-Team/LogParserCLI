package parsers

import (
	"fmt"
	"strings"
)


func SearchInTXT(data []string, keyword string) {
	var found bool
	for i := 0; i < len(data); i++ {
		
		if strings.Contains(data[i], "URL:") && strings.Contains(data[i], keyword) {
			found = true
			fmt.Println("Found the following information:")
			fmt.Println("===============================")
			fmt.Println(data[i])

			
			if i+1 < len(data) {
				fmt.Println(data[i+1]) 
			}
			if i+2 < len(data) {
				fmt.Println(data[i+2]) 
			}
			if i+3 < len(data) {
				fmt.Println(data[i+3]) 
			}
			fmt.Println("===============================")
		}
	}

	if !found {
		fmt.Println("No information found for the keyword:", keyword)
	}
}
