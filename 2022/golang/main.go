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
	snacks := [3]int{ 0, 0, 0 }
	counter, pos := 0, 0
	for _, line := range lines {
		if line != "" {
			i, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			counter += i

			continue
		}
		

		weighSnack(counter, &snacks)	
		counter = 0
		pos++
	}

	fmt.Println(snacks)
	fmt.Println(calculateSnackWeight(snacks))
}

func weighSnack(snack int, snacks *[3]int) {
	if snack > snacks[0] {
		t1, t0 := snacks[1], snacks[0]
		snacks[2] = t1
		snacks[1] = t0
		snacks[0] = snack
	} else if snack > snacks[1] && snack < snacks[0] {
		t1 := snacks[1]
		snacks[2] = t1
		snacks[1] = snack
	} else if snack > snacks[2] && snack < snacks[1] {
		snacks[2] = snack
	}
}

func calculateSnackWeight(snacks [3]int) int {
	return (snacks[0] + snacks[1] + snacks[2])
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
