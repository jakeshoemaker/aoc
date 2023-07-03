package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var lines, err = parseFile()
	if err != nil {
		fmt.Println("error parsing file")
	}

	highest, highest_pos := 0, 1
	counter, pos := 0, 1
	for _, line := range lines {
		if line != "" {
			i, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			counter += i
			continue
		}

		// line ended, increment pos, eval counter, reset
		if counter > highest {
			highest = counter
			highest_pos = pos
		}
		pos++
		counter = 0
	}

	fmt.Println(highest_pos)
	fmt.Println(highest)
}

func parseFile() ([]string, error) {
	file, err := os.Open("numbers.txt")
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
