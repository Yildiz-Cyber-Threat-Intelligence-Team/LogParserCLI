package parsers

import (
	"fmt"
	"strings"
)

// TXT içindeki satırları tarar ve anahtar kelimenin geçtiği yerleri bulur
func SearchInTXT(data []string, keyword string) {
	var found bool
	for i := 0; i < len(data); i++ {
		// Eğer satırda URL varsa ve anahtar kelime mevcutsa
		if strings.Contains(data[i], "URL:") && strings.Contains(data[i], keyword) {
			found = true
			fmt.Println("Found the following information:")
			fmt.Println("===============================")
			fmt.Println(data[i]) // URL satırını yazdır

			// Username, Password ve Application bilgilerini bul
			if i+1 < len(data) {
				fmt.Println(data[i+1]) // Username
			}
			if i+2 < len(data) {
				fmt.Println(data[i+2]) // Password
			}
			if i+3 < len(data) {
				fmt.Println(data[i+3]) // Application
			}
			fmt.Println("===============================")
		}
	}

	if !found {
		fmt.Println("No information found for the keyword:", keyword)
	}
}
