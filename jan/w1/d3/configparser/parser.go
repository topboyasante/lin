package configparser

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func ParseConfig(filePath string) (map[string]string, error) {
	// open the file
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	config := map[string]string{}

	for scanner.Scan() {
		line := scanner.Text() // Get the line as a string

		// We know that for config files, comments start with a #.
		// we can assume that lines starting with # are comments
		lineRunes := []rune(line)
		firstLetter := string(lineRunes[0])
		if firstLetter == "#" {
			continue
		}

		// We split the key=value text on the first instance we see "=".
		// strings.Cut gives us both the left and right side,
		// which ideally will be the key and the value
		key, value, found := strings.Cut(line, "=")

		// if we don't find the "=" on the line, we know that we have a malformed line
		// we also need to handle cases where key or value may be empty (like key= or =value)
		if !found || key == "" || value == "" {
			log.Printf("parsing was terminated because there was an error reading line '%s'. it is malformed.", line)
			break
		}

		config[key] = value

	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading file: %s", err)
	}

	return config, nil

}
