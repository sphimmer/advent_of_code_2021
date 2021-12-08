package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	boards, entries := buildBoardsAndGetEntries()
	winningBoard, lastCalledEntry := playBingo(boards, entries)
	sum := 0
	for _, r := range winningBoard.rows {
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
	lastCall, err := strconv.Atoi(lastCalledEntry)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	fmt.Println("Answer: ", sum*lastCall)

}

func playBingo(boards []board, entries []string) (board, string) {
	winner := board{}
	for i, e := range entries {
		for b := range boards {
			boards[b].selectCells(e)
			if i >= 4 {
				if boards[b].isWinner() {
					return boards[b], e
				}
			}
		}
	}
	return winner, ""
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
