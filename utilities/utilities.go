package utilities

import (
	"encoding/csv"
	"fmt"
	"go/npcGen/configuration"
	"log"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

func ReadCSV(path string) [][]string {
	// Open CSV File
	log.Printf("reading %s file", path)
	f, err := os.OpenFile(path, os.O_RDONLY, 0)
	if err != nil {
		log.Printf("error opening CSV: %s", err)
	}
	defer f.Close()

	// Read CSV File
	reader := csv.NewReader(f)
	data, err := reader.ReadAll()
	if err != nil {
		log.Printf("error reading CSV: %s", err)
	}

	return data
}

func WriteCSV(path string, filename string, data [][]string) {
	// Create CSV file
	location := fmt.Sprintf("%s/%s", path, filename)
	log.Printf("writing data to %s", location)
	file, err := os.Create(location)
	if err != nil {
		log.Print(err)
	}
	defer file.Close()

	// Write to CSV
	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, val := range data {
		var row []string
		row = append(row, val...)
		writer.Write(row)
	}
}

func ReadConfig(path string, config *configuration.Config) error {
	f, err := os.OpenFile(path, os.O_RDONLY, 0)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Parse file based on file extention
	// TODO(wholesomeow): Implement Environment variables and JSON
	switch ext := strings.ToLower(filepath.Ext(path)); ext {
	case ".yaml", ".yml":
		yaml_decoder := yaml.NewDecoder(f)
		err := yaml_decoder.Decode(config)
		if err != nil {
			// TODO(wholesomeow): Figure out better logging/error handling for known things like this
			log.Fatal(err)
		}
	default:
		return fmt.Errorf("file format '%s' not supported by parser", ext)
	}

	return nil
}

func CheckFilePath(path string, required bool) bool {
	// TODO(wholesomeow): Maybe rewrite to remove error handling and make function more flexible
	found := true
	_, err := os.Stat(path)
	if err == nil {
		log.Printf("file %s exists", path)
		return found
	} else if os.IsNotExist(err) {
		if !required {
			log.Printf("file %s doesn't exist in expected location", path)
		}
		log.Fatalf("file %s does not exist", path)
	} else {
		log.Fatalf("file %s stat error: %v", path, err)
	}
	return !found
}
