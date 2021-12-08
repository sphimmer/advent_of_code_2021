package main

import "strings"

type board struct {
	rows [][]cell
}

func newBoard(input []string) board {
	b := board{rows: [][]cell{}}
	for i, l := range input {
		entries := strings.Split(l, " ")
		b.rows = append(b.rows, []cell{})
		for _, v := range entries {
			if v == "" {
				continue
			}
			b.rows[i] = append(b.rows[i], cell{value: v, selected: false})
		}
	}
	return b
}

func (b board) isWinner() bool {

	for i := range b.rows {
		isColumnWinner := true
		isRowWinner := true
		for l := range b.rows[i] {
			if !b.rows[i][l].selected {
				isRowWinner = false
			}
			if !b.rows[l][i].selected {
				isColumnWinner = false
			}
		}
		if isColumnWinner || isRowWinner {
			return true
		}
	}
	return false
}

func (b board) selectCells(number string) {
	for _, r := range b.rows {
		for i := range r {
			if r[i].value == number {
				r[i].selected = true
			}
		}
	}
}
