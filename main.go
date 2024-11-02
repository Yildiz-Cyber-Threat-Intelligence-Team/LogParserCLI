package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"stealer-parser/parsers"
	"strings"

	"github.com/nwaples/rardecode"
)

func ExtractZip(src string, dest string) ([]string, error) {
	var extractedFiles []string
	r, err := zip.OpenReader(src)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	for _, f := range r.File {
		fpath := filepath.Join(dest, f.Name)
		extractedFiles = append(extractedFiles, fpath)

		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		if err := os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return nil, err
		}

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return nil, err
		}
		defer outFile.Close()

		rc, err := f.Open()
		if err != nil {
			return nil, err
		}

		_, err = io.Copy(outFile, rc)
		if err != nil {
			return nil, err
		}

		rc.Close()
	}
	return extractedFiles, nil
}

func ExtractRar(src string, dest string) ([]string, error) {
	var extractedFiles []string
	file, err := os.Open(src)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	r, err := rardecode.NewReader(file, "")
	if err != nil {
		return nil, err
	}

	for {
		header, err := r.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		fpath := filepath.Join(dest, header.Name)
		extractedFiles = append(extractedFiles, fpath)

		if header.IsDir {
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		if err := os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return nil, err
		}

		outFile, err := os.Create(fpath)
		if err != nil {
			return nil, err
		}
		defer outFile.Close()

		_, err = io.Copy(outFile, r)
		if err != nil {
			return nil, err
		}
	}
	return extractedFiles, nil
}

func SearchInJSON(data []map[string]interface{}, keyword string) {
	found := false
	for _, entry := range data {
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

func main() {
	var filePath, keyword string
	fmt.Println("Enter the file path:")
	fmt.Scanln(&filePath)

	fmt.Println("Enter the search keyword:")
	fmt.Scanln(&keyword)

	destDir := "./extracted_files"

	if strings.HasSuffix(filePath, ".zip") {
		extractedFiles, err := ExtractZip(filePath, destDir)
		if err != nil {
			log.Fatalf("Failed to extract zip file: %v", err)
		}
		fmt.Println("Extracted files:", extractedFiles)

		for _, extractedFilePath := range extractedFiles {
			searchFile(extractedFilePath, keyword)
		}

	} else if strings.HasSuffix(filePath, ".rar") {
		extractedFiles, err := ExtractRar(filePath, destDir)
		if err != nil {
			log.Fatalf("Failed to extract rar file: %v", err)
		}
		fmt.Println("Extracted files:", extractedFiles)

		for _, extractedFilePath := range extractedFiles {
			searchFile(extractedFilePath, keyword)
		}

	} else {
		searchFile(filePath, keyword)
	}
}

func searchFile(filePath, keyword string) {
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
