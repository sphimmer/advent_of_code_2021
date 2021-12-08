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

	for _, r := range b.rows {
		isRowWinner := true
		for i := range r {
			if !r[i].selected {
				isRowWinner = false
				break
			}
		}
		if isRowWinner {
			return true
		}
	}

	for i := range b.rows {
		isColumnWinner := true
		for l := range b.rows[i] {
			if !b.rows[l][i].selected {
				isColumnWinner = false
			}
		}
		if isColumnWinner {
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
