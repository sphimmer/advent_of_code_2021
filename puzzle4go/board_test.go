package main

import (
	"testing"
)

func TestNewBoard(t *testing.T) {
	entries := []string{"22 13 17 11  0", " 8  2 23  4 24", "21 9 14 16 7", "6 10 3 18 5", "1 12 20 15 19"}
	b := newBoard(entries)

	if len(b.rows) < 5 {
		t.Errorf("Expected 5 rows but got %v", len(b.rows))
	}

	if len(b.rows[0]) < 5 {
		t.Errorf("Expected 5 columns but got %v", len(b.rows[0]))
	}
}

func TestSelectCell(t *testing.T) {
	entries := []string{"22 13 17 11  0", " 8  2 23  4 24", "21 9 14 16 7", "6 10 3 18 5", "1 12 20 15 19"}
	b := newBoard(entries)
	b = b.selectCells("13")
	if !b.rows[0][1].selected {
		t.Errorf("Expected cell selected to be true but got %v", b.rows[0][1].selected)
	}

	if b.rows[0][2].selected {
		t.Errorf("Expected cell selected to be false but got %v", b.rows[0][1].selected)
	}
}

func TestIsRowWinner(t *testing.T) {
	entries := []string{"22 13 17 11  0", " 8  2 23  4 24", "21 9 14 16 7", "6 10 3 18 5", "1 12 20 15 19"}
	b := newBoard(entries)
	b.selectCells("22")
	b.selectCells("13")
	b.selectCells("17")
	b.selectCells("11")
	b.selectCells("0")
	if !b.isWinner() {
		t.Error("expected board to be a winner but was not")
	}
	if b.rows[1][0].selected {
		t.Error("this cell should not be selected")
	}
	b = newBoard(entries)
	b.selectCells("22")
	b.selectCells("9")
	b.selectCells("17")
	b.selectCells("11")
	b.selectCells("0")
	if b.isWinner() {
		t.Error("expected board to not be a winner but was")
	}
}

func TestIsColumnWinner(t *testing.T) {
	entries := []string{"22 13 17 11  0", " 8  2 23  4 24", "21 9 14 16 7", "6 10 3 18 5", "1 12 20 15 19"}
	b := newBoard(entries)
	b.selectCells("22")
	b.selectCells("8")
	b.selectCells("21")
	b.selectCells("6")
	b.selectCells("1")
	if !b.isWinner() {
		t.Error("expected board to be a winner but was not")
	}

	b = newBoard(entries)
	b.selectCells("22")
	b.selectCells("2")
	b.selectCells("21")
	b.selectCells("6")
	b.selectCells("1")
	if b.isWinner() {
		t.Error("expected board to be not winner but was")
	}
}
