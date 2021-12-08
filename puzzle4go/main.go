package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	solvePart1()
	solvePart2()
}

func solvePart2() {
	boards, entries := buildBoardsAndGetEntries()
	winners := playBingo(boards, entries)
	lastCall, err := strconv.Atoi(winners[len(winners)-1].lastCalledNum)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	fmt.Println("Part 2 Answer: ", calculateScore(winners[len(winners)-1].Board, lastCall))
}

func solvePart1() {
	boards, entries := buildBoardsAndGetEntries()
	winners := playBingo(boards, entries)
	lastCall, err := strconv.Atoi(winners[0].lastCalledNum)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	fmt.Println("Part 1 Answer: ", calculateScore(winners[0].Board, lastCall))
}

func calculateScore(b board, lastCall int) int {
	sum := 0
	for _, r := range b.rows {
		for _, c := range r {
			if !c.selected {
				value, err := strconv.Atoi(c.value)
				if err != nil {
					fmt.Printf("Error: %v", err)
				}
				sum += value
			}
		}
	}

	return sum * lastCall
}

func playBingo(boards []board, entries []string) []winner {
	winners := []winner{}
	for _, e := range entries {
		for b := range boards {
			if !winnersContain(winners, b) {
				boards[b].selectCells(e)
				if boards[b].isWinner() {
					winners = append(winners, winner{id: b, Board: boards[b], lastCalledNum: e})
				}
			}
		}
	}
	return winners
}

func parseFile() string {
	bs, err := ioutil.ReadFile("input")
	if err != nil {
		fmt.Printf("error opening file %v", err)
	}
	inputString := string(bs)
	return inputString
}

func buildBoardsAndGetEntries() ([]board, []string) {
	scanner := bufio.NewScanner(strings.NewReader(parseFile()))
	line := 0
	var entries []string
	var boards []board
	boardNum := 0
	var boardContent []string
	for scanner.Scan() {
		content := scanner.Text()
		if line == 0 {
			entries = strings.Split(content, ",")
			line++
			continue
		}

		if content == "" && len(boardContent) == 0 {
			line++
			continue
		} else {
			boardContent = append(boardContent, content)
		}

		if content == "" || len(boardContent) == 5 {
			boards = append(boards, newBoard(boardContent))
			boardContent = []string{}
			boardNum++
		}
		line++
	}
	return boards, entries
}
