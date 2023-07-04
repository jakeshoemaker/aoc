package main

import (
	"bufio"
	"io"
	"os"
)

func main() {
	rpc_rounds, err := parseFile()
    if err != nil {
        panic(err)
    }

    for _, line in range rpc_rounds {
        // TODO: solve :)
    }
}

// TODO: refactor this
func parseFile() ([]string, error) {
	file, err := os.Open("data/day2.txt")
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
