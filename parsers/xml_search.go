package parsers

import (
	"fmt"
	"strings"
)

func SearchInXML(data []Entry, keyword string) {
	var found bool
	for _, entry := range data {
		if strings.Contains(entry.URL, keyword) {
			found = true
			fmt.Println("Found the following information:")
			fmt.Println("===============================")
			fmt.Println("URL:", entry.URL)
			fmt.Println("Username:", entry.Username)
			fmt.Println("Password:", entry.Password)
			fmt.Println("Application:", entry.Application)
			fmt.Println("===============================")
		}
	}

	if !found {
		fmt.Println("No information found for the keyword:", keyword)
	}
}
