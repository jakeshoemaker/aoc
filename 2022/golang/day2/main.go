package main

import (
	"bufio"
    "fmt"
	"os"
)

func main() {
	rpc_rounds, err := parseFile()
    if err != nil {
        panic(err)
    }

    my_score := 0
    for _, line := range rpc_rounds {
        evaluateRound(line, &my_score)
        fmt.Println("score update:")
        fmt.Println(my_score)
    }
    
    fmt.Println(my_score)
}

func evaluateRound(line string, my_score *int) {
    elf_move, my_move := rune(line[0]), rune(line[2])

    switch elf_move {
    case 'A': // rock
        elfPlaysRock(my_move, my_score)
    case 'B': // paper
        elfPlaysPaper(my_move, my_score)
    case 'C': // scissors
        elfPlaysScissors(my_move, my_score)
    }
}

// 0 for loss | 3 for draw | 6 for win
// X == rock(1) | Y == paper(2) | Z == scissor(3)
func elfPlaysRock(my_move rune, my_score *int) {
    curr_score := *my_score
    switch my_move {
    case 'X': // draw
        *my_score = curr_score + 4
    case 'Y': // win
        *my_score = curr_score + 8
    case 'Z': // loss
        *my_score = curr_score + 3
    default:
        panic("how did we get here")
    }
}

func elfPlaysPaper(my_move rune, my_score *int) {
    curr_score := *my_score
    switch my_move {
    case 'X': // loss
        *my_score = curr_score + 1
    case 'Y': // draw
        *my_score = curr_score + 5
    case 'Z': // win
        *my_score = curr_score + 9
    default:
        panic("how did we get here")
    }
}

func elfPlaysScissors(my_move rune, my_score *int) {
    curr_score := *my_score
    switch my_move {
    case 'X': // win
        *my_score = curr_score + 7
    case 'Y': // loss
        *my_score = curr_score + 2
    case 'Z': // draw
        *my_score = curr_score + 6
    default:
        panic("how did we get here")
    }
}

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
