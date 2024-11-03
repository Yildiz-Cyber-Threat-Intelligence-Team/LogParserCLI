package parsers

import (
	"encoding/xml"
	"fmt"
	"strings"
)


func SearchInXML(data string, keyword string) {
	decoder := xml.NewDecoder(strings.NewReader(data))
	for {
		token, err := decoder.Token()
		if err != nil {
			break
		}
		switch elem := token.(type) {
		case xml.StartElement:
		
			if strings.Contains(elem.Name.Local, keyword) {
				fmt.Println("Found in element:", elem.Name.Local)
			}
		case xml.CharData:
		
			content := string(elem)
			if strings.Contains(content, keyword) {
				fmt.Println("Found in content:", content)
			}
		}
	}
}

