package helpers

import (
    "bufio"
    "os"
)

func ParseFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var lines []string
	// append lines to our line buffer
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	file.Close()

	return lines, nil
}
