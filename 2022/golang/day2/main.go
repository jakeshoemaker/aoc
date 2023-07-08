package main

import (
    "fmt"
    "jshoemaker/advent_of_code/helpers"
)

const (
    Rock = 1
    Paper = 2
    Scissors = 3
)

const (
    Lose = 0
    Draw = 3
    Win = 6
)

func main() {
	rpc_rounds, err := helpers.ParseFile("data/day2.txt")
    if err != nil {
        panic(err)
    }

    score := 0
    for _, line := range rpc_rounds {
        evaluateRound(line, &score)
        fmt.Println("score update:")
        fmt.Println(score)
    }
    
    fmt.Println(score)
}

func evaluateRound(line string, my_score *int) {
    elf_move, outcome := rune(line[0]), rune(line[2])

    switch outcome {
    case 'X': // lose
        lose(elf_move, my_score)
    case 'Y': // draw
        draw(elf_move, my_score)
    case 'Z': // win
        win(elf_move, my_score)
    }
}

// 0 for loss | 3 for draw | 6 for win
// X == rock(1) | Y == paper(2) | Z == scissor(3)
func lose(elf_move rune, my_score *int) {
    curr_score := *my_score
    switch elf_move {
    case 'A': // rock
        *my_score = curr_score + Scissors
    case 'B': // paper
        *my_score = curr_score + Rock
    case 'C': // scissors
        *my_score = curr_score + Paper
    default:
        panic("how did we get here")
    }
}

func draw(move rune, my_score *int) {
    curr_score := *my_score
    switch move {
    case 'A': // rock
        *my_score = curr_score + Rock + Draw
    case 'B': // paper
        *my_score = curr_score + Paper + Draw
    case 'C': // scissors
        *my_score = curr_score + Scissors + Draw
    default:
        panic("how did we get here")
    }
}

func win(move rune, my_score *int) {
    curr_score := *my_score
    switch move {
    case 'A': // rock
        *my_score = curr_score + Paper + Win
    case 'B': // paper
        *my_score = curr_score + Scissors + Win
    case 'C': // scissors
        *my_score = curr_score + Rock + Win
    default:
        panic("how did we get here")
    }
}
